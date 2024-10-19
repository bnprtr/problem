package problem

import (
	"errors"
	"fmt"
	"log/slog"
)

type (
	Title      = ErrorTitle
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
	return fmt.Errorf("%s: %s. %v, %v", e.Title, e.Message, e.Details, errors.Join(e.Internal...)).Error()
}

func (e Error[T]) Unwrap() []error {
	return append(e.Internal, e.Title)
}

func (e Error[T]) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("title", string(e.Title)),
		slog.String("message", string(e.Message)),
		slog.Any("details", e.Details),
		slog.Any("internalErrors", e.Internal),
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
