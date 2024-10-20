package problem

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"runtime"
	"strconv"
)

type (
	ErrorTitle interface {
		~string
		error
		StatusCode() StatusCode
		New(string) Error
	}

	// Error represents a problem as defined by RFC 7807. It is go error value that can hold internal errors.
	// Internal errors, the Title, and the HTTP Status Code are unwrapped allowing for error matching using
	// errors.Is and errors.As. The Error properties can be set using a builder pattern:
	//
	//     err := problem.New("title").WithMessage("message").WithDetails("details").WithInternalErrors(internalErrs...)
	//
	// The error is also a loggable value using slog:
	//
	// 		slog.InfoContext(ctx, "something bad happened", "error", err)
	//
	// JSON Serialization of the error will not render out the internal errors. The internal errors are only stored for
	// internal use and are not intended to be exposed outside of the closed system. External details should be added as
	// Details of the problem.
	Error struct {
		// Title is a short, matchable string of the error. It is recommended for the title to be human readable and
		// representative of the error condition.
		//
		// example: "Missing Required Field"
		//
		// Titles are best defined as constants in your codebase to ensure consistency and to allow for easy matching.
		// Use the "Status" types to map the title to an HTTP status code.
		//
		// example: const MissingFieldRequired problem.HTTPStatusBadRequest = "Missing Required Field"
		//
		// This allows for easy matching of error types or status codes in your codebase:
		//
		//   // check if the error is a specific title
		//   if errors.Is(err, MissingFieldRequired) {
		// 		// handle missing field error
		//   }
		//
		//   // check if the error is a specific status code
		//   if errors.Is(err, problem.HTTPStatusBadRequest) {
		// 		// handle bad request error
		//   }
		//
		//   // get the status code of the error
		//   var status problem.StatusCode
		//   if errors.As(err, &status) {
		//   	// handle status code
		//   }
		//
		//   // get the title of the error
		//   var title problem.Title
		//   if errors.As(err, &title) {
		//   	// handle title
		//   }
		//
		Title string `json:"title"`

		// Message is a human-readable explanation specific to this occurrence of the problem. It is recommended to be
		// short and to the point. It is not intended to be a complete description of the error. The message may
		// displayed on the client it so it should not be overly technical. Technical details for debugging or advanced
		// behvaior response should be added to the Details instead.
		// example: "The request body is missing the required field 'name'."
		Message string `json:"message"`

		// Stacktrace is the stacktrace of the error. It is not intended to be exposed outside of the closed system.
		// The stacktrace is stored for internal use and debugging. The stacktrace is not included in the JSON
		// serialization of the error. The stacktrace is not unwrapped when the error is unwrapped.
		// The stacktrace is not automatically captured. It must be manually set by invoking the Stack() method
		// when the error is created.
		Stacktrace string `json:"-"`

		// Internal is a list of internal errors that caused the problem. These errors are not intended to be exposed
		// outside of the closed system. They are stored for internal use and debugging. The internal errors are
		// unwrapped when the error is unwrapped, allowing for easy matching and logging.
		Internal []error `json:"-"`

		// Details is a list of additional details that may be useful for debugging or advanced behavior. The details
		// are intended to be exposed externally and are included in the JSON serialization of the error.
		// It is recommended to use a struct for the details to allow for easy serialization and deserialization.
		// It is also recommended to use standard format for different types of details. This package provides a
		// standard FieldDetail type for common field errors.
		Details []any `json:"details"`

		// HttpStatusCode is the HTTP status code that should be returned with the error. The status code is used to
		// determine the response status code when the error is returned to the client. The status code is unwrapped
		// when the error is unwrapped, allowing for simple matching and serialization of errors returned in a centrale
		// error handler in HTTP handlers.
		//
		// example:
		//
		//   var statusCode problem.StatusCode
		//   if !errors.As(err, &statusCode) {
		//   	// unknown error, return internal server error
		//   	statusCode = problem.HTTPStatusInternalServerError
		//      // ovveride the error and wrap in an internal server error
		//  	err = problem.New(problem.StatusInternalServerError("Internal Server Error")).WithInternalErrors(err)
		//   }
		//	 w.Header().Set("Content-Type", "application/problem+json")
		//	 w.WriteHeader(int(statusCode))
		//	 json.NewEncoder(w).Encode(err)
		HttpStatusCode StatusCode `json:"statusCode"`
	}
	StatusCode int
	Title      string
	Details    []any

	// StackTraced can be used to extend the Status Title types so the Title.New() always captures the stack
	// trace. Establish your title consts with StackTraced[ErrorTitle] and every New calls from the title will
	// automatically capture the stacktrace.
	StackTraced[T ErrorTitle] string
)

// New creates a new Error with the given title.
func New[T ErrorTitle](title T) Error {
	return Error{
		Title:          string(title),
		HttpStatusCode: title.StatusCode(),
	}
}

// Error returns a string representation of the Error. It is rendered as a JSON object, including all internal errors.
func (e Error) Error() string {
	presentation := struct {
		Title          string     `json:"title"`
		Message        string     `json:"message"`
		Stacktrace     string     `json:"stacktrace,omitempty"`
		Internal       []error    `json:"internalError,omitempty"`
		Details        []any      `json:"details,omitempty"`
		HttpStatusCode StatusCode `json:"statusCode"`
	}(e)

	b, err := json.Marshal(presentation)
	if err != nil {
		return fmt.Errorf("%s: %s. details=(%+v), internalErrors=(%+v)", e.Title, e.Message, e.Details, errors.Join(e.Internal...)).Error()
	}
	return string(b)
}

func (e Error) Unwrap() []error {
	title, stackedTitle := convertToTitle(int(e.HttpStatusCode), e.Title)
	return append(
		e.Internal,
		title,
		stackedTitle,
		e.HttpStatusCode,
		Details(e.Details),
		Title(e.Title),
	)
}

func (e Error) LogValue() slog.Value {
	values := []slog.Attr{
		slog.String("title", string(e.Title)),
		slog.String("message", string(e.Message)),
	}
	if len(e.Details) > 0 {
		values = append(values, slog.Any("details", e.Details))
	}
	if len(e.Internal) > 0 {
		errors := make([]string, len(e.Internal))
		for i, err := range e.Internal {
			errors[i] = err.Error()
		}
		values = append(values, slog.Any("internalErrors", errors))
	}
	if e.Stacktrace != "" {
		values = append(values, slog.String("stacktrace", e.Stacktrace))
	}
	return slog.GroupValue(
		values...,
	)
}

// SetTitle overrides the title of the error. This is not recommended as the title should be a constant in your codebase.
func (e Error) SetTitle(title string) Error {
	e.Title = title
	return e
}

// AddDetails adds additional details to the error. The details are intended to be exposed externally and are included
// in the JSON serialization of the error.
func (e Error) AddDetails(details ...any) Error {
	e.Details = append(e.Details, details...)
	return e
}

// AddInternalErrors adds internal errors to the error. The internal errors are not intended to be exposed outside of
// the closed system. They are stored for internal use and debugging. The internal errors are unwrapped when the error
// is unwrapped, allowing for easy matching and logging.
func (e Error) AddInternalErrors(errors ...error) Error {
	e.Internal = append(e.Internal, errors...)
	return e
}

// SetMessage overrides the message of the error. The message is a human-readable explanation specific to this occurrence
func (e Error) SetMessage(message string) Error {
	e.Message = message
	return e
}

// Stack captures the stack trace from the point where it is called and stores it in the Stacktrace property.
func (e Error) Stack() Error {
	const skipFrames = 2 // Skip the first 2 frames to exclude runtime.Callers and Stack
	return e.captureStack(skipFrames)
}

func (e Error) captureStack(skipFrames int) Error {
	stackBuf := make([]uintptr, 32)
	stackSize := runtime.Callers(skipFrames, stackBuf)
	stackBuf = stackBuf[:stackSize]

	stackTrace := make([]byte, 0, 1024)
	for _, pc := range stackBuf {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line := fn.FileLine(pc)
		stackTrace = append(stackTrace, fmt.Sprintf("%s:%d\n", file, line)...)
	}

	e.Stacktrace = string(stackTrace)
	return e
}

func (s StatusCode) Error() string {
	return strconv.Itoa(int(s))
}

func (s Details) Error() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (t Title) Error() string {
	return string(t)
}

func (e StackTraced[T]) Error() string {
	return string(e)
}

func (e StackTraced[T]) New(message string) Error {
	return T(e).New(message).captureStack(-4)
}

func (e StackTraced[T]) StatusCode() StatusCode {
	return T(e).StatusCode()
}
