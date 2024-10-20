package problem

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	BadRequest StatusBadRequest = "bad request"
	Forbidden  StatusForbidden  = "forbidden"
)

func TestError(t *testing.T) {
	err := BadRequest.New("Got orange. Wanted duck.").AddDetails("orange", struct{ Count int }{3})

	t.Run("errors.Is can extract the Title", func(t *testing.T) {
		t.Logf("error: %+v", err)
		if !errors.Is(err, BadRequest) {
			t.Errorf("expected: %+v, got: %+v", BadRequest, err)
		}

		if errors.Is(err, Forbidden) {
			t.Errorf("expected: %+v, got: %+v", BadRequest, Forbidden)
		}
	})

	t.Run("errors.Is can extract an internal error", func(t *testing.T) {
		internal := errors.New("something bad")
		err := err.AddInternalErrors(internal)
		if !errors.Is(err, internal) {
			t.Errorf("expected internal error: %+v, got: %+v", internal, err)
		}
		if !errors.Is(err, BadRequest) {
			t.Errorf("expected title: %+v, got: %+v", BadRequest, err)
		}
	})
}

type errorLogEntry struct {
	Title          string `json:"title"`
	Message        string `json:"message"`
	Stack          string `json:"stacktrace"`
	Details        []any  `json:"details"`
	InternalErrors []any  `json:"internalErrors"`
}

type logEntry struct {
	Level   string        `json:"level"`
	Message string        `json:"msg"`
	Error   errorLogEntry `json:"err"`
}

func TestErrorLogValue(t *testing.T) {
	err := BadRequest.New("Got orange. Wanted duck.")
	t.Run("the error is logged with only a title and message", func(t *testing.T) {
		out := bytes.NewBuffer(nil)
		log := slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{}))
		log.Error("error", "err", err)
		t.Log(out.String())
		var got logEntry
		_ = json.Unmarshal(out.Bytes(), &got)
		expected := logEntry{
			Level:   "ERROR",
			Message: "error",
			Error: errorLogEntry{
				Title:   string(BadRequest),
				Message: "Got orange. Wanted duck.",
			},
		}

		assert.Equal(t, expected, got)
	})

	t.Run("the error is logged with details and internal errors", func(t *testing.T) {
		out := bytes.NewBuffer(nil)
		log := slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{}))
		log.Error("error", "err",
			err.
				Stack().
				AddInternalErrors(
					errors.New("1"),
					errors.New("2"),
				).
				AddDetails(
					NewFieldDetail("f1", "", "empty"),
					NewFieldDetail("f2", "aa", "length"),
				),
		)
		t.Log(out.String())
		var got logEntry
		_ = json.Unmarshal(out.Bytes(), &got)
		expected := logEntry{
			Level:   "ERROR",
			Message: "error",
			Error: errorLogEntry{
				Title:   string(BadRequest),
				Message: "Got orange. Wanted duck.",
				InternalErrors: []any{
					"1", "2",
				},
				Details: []any{
					map[string]any{
						"field":  "f1",
						"value":  "",
						"reason": "empty",
					},
					map[string]any{
						"field":  "f2",
						"value":  "aa",
						"reason": "length",
					},
				},
			},
		}
		assert.NotEmpty(t, got.Error.Stack)
		got.Error.Stack = ""
		assert.Equal(t, expected, got)
	})

	t.Run("The StackTraced title captures a stack trace", func(t *testing.T) {
		err := StackTraced[StatusForbidden]("access denied").New("You shall not pass")
		out := bytes.NewBuffer(nil)
		log := slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{}))
		log.Error("error", "err", err)
		t.Log(out.String())
		var got logEntry
		_ = json.Unmarshal(out.Bytes(), &got)

		expected := logEntry{
			Level:   "ERROR",
			Message: "error",
			Error: errorLogEntry{
				Title:   "access denied",
				Message: "You shall not pass",
			},
		}

		assert.NotEmpty(t, got.Error.Stack)
		got.Error.Stack = ""
		assert.Equal(t, expected, got)
	})
}
