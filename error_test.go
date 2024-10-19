package problem

import (
	"errors"
	"testing"
)

const (
	BadRequest StatusBadRequest = "bad request"
	Forbidden  StatusForbidden  = "forbidden"
)

func TestError(t *testing.T) {
	err := BadRequest.New("Got orange. Wanted duck.").WithDetails("orange", struct{ Count int }{3})

	t.Run("errors.Is can extract the Title", func(t *testing.T) {
		if !errors.Is(err, BadRequest) {
			t.Errorf("expected: %+v, got: %+v", BadRequest, err)
		}

		if errors.Is(err, Forbidden) {
			t.Errorf("expected: %+v, got: %+v", BadRequest, Forbidden)
		}
	})

	t.Run("errors.Is can extract an internal error", func(t *testing.T) {
		internal := errors.New("something bad")
		err := err.WithInternalErrors(internal)
		if !errors.Is(err, internal) {
			t.Errorf("expected internal error: %+v, got: %+v", internal, err)
		}
		if !errors.Is(err, BadRequest) {
			t.Errorf("expected title: %+v, got: %+v", BadRequest, err)
		}
	})
}

func TestErrorLogValue(t *testing.T) {
	err := New(BadRequest)
	t.Run("", func(t *testing.T) {
		_ = err
	})
}
