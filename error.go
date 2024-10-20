package problem

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
)

type (
	ErrorTitle interface {
		~string
		error
		StatusCode() int
	}
	Error[T ErrorTitle] struct {
		Title          T       `json:"title"`
		Message        string  `json:"message"`
		Internal       []error `json:"-"`
		Details        []any   `json:"details"`
		HttpStatusCode int     `json:"statusCode"`
	}
)

func New[T ErrorTitle](title T) Error[T] {
	return Error[T]{
		Title:          title,
		HttpStatusCode: title.StatusCode(),
	}
}

func (e Error[T]) Error() string {
	presentation := struct {
		Title          T       `json:"title"`
		Message        string  `json:"message"`
		Internal       []error `json:"internalError,omitempty"`
		Details        []any   `json:"details,omitempty"`
		HttpStatusCode int     `json:"statusCode"`
	}(e)

	b, err := json.Marshal(presentation)
	if err != nil {
		return fmt.Errorf("%s: %s. details=(%+v), internalErrors=(%+v)", e.Title, e.Message, e.Details, errors.Join(e.Internal...)).Error()
	}
	return string(b)
}

func (e Error[T]) Unwrap() []error {
	return append(e.Internal, e.Title)
}

func (e Error[T]) LogValue() slog.Value {
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

func (e Error[T]) WithTitle(title T) Error[T] {
	e.Title = title
	return e
}

func (e Error[T]) WithDetails(details ...any) Error[T] {
	e.Details = append(e.Details, details...)
	return e
}

func (e Error[T]) WithInternalErrors(errors ...error) Error[T] {
	e.Internal = append(e.Internal, errors...)
	return e
}

func (e Error[T]) WithMessage(message string) Error[T] {
	e.Message = message
	return e
}
