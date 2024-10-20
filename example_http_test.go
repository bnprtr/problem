package problem_test

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/bnprtr/problem"
)

func Example() {
	_ = http.ListenAndServe(":8080", http.HandlerFunc(MyHandler))
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	if err := handle(r); err != nil {
		var statusCode problem.StatusCode
		if !errors.As(err, &statusCode) {
			// unknown error!
			err = InternalServerError.New("unexpected error occurred.").AddInternalErrors(err)
			statusCode = InternalServerError.StatusCode()
		}
		switch {
		case statusCode < 500:
			slog.WarnContext(r.Context(), "handler returned error", "error", err)
		default:
			slog.ErrorContext(r.Context(), "handler returned error", "error", err)
		}
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(int(statusCode))
		_ = json.NewEncoder(w).Encode(err)
	}
}

func handle(r *http.Request) error {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON.New(err.Error()).AddInternalErrors(err).Stack()
	}

	if err := validate(req); err != nil {
		panic(err)
	}
	return nil
}
