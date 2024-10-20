package problem

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
)

type (
	ErrorTitle interface {
		~string
		error
		StatusCode() StatusCode
	}
	Error struct {
		Title          string     `json:"title"`
		Message        string     `json:"message"`
		Internal       []error    `json:"-"`
		Details        []any      `json:"details"`
		HttpStatusCode StatusCode `json:"statusCode"`
	}
	StatusCode int
	Details    []any
)

func New[T ErrorTitle](title T) Error {
	return Error{
		Title:          string(title),
		HttpStatusCode: title.StatusCode(),
	}
}

func (e Error) Error() string {
	presentation := struct {
		Title          string     `json:"title"`
		Message        string     `json:"message"`
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
	return append(e.Internal, convertToTitle(int(e.HttpStatusCode), e.Title), e.HttpStatusCode, Details(e.Details))
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
		values = append(values, slog.Any("internalErrors", e.Internal))
	}
	return slog.GroupValue(
		values...,
	)
}

func (e Error) WithTitle(title string) Error {
	e.Title = title
	return e
}

func (e Error) WithDetails(details ...any) Error {
	e.Details = append(e.Details, details...)
	return e
}

func (e Error) WithInternalErrors(errors ...error) Error {
	e.Internal = append(e.Internal, errors...)
	return e
}

func (e Error) WithMessage(message string) Error {
	e.Message = message
	return e
}

func (s StatusCode) Error() string {
	return strconv.Itoa(int(s))
}

func (s Details) Error() string {
	b, _ := json.Marshal(s)
	return string(b)
}
