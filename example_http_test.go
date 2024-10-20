package problem_test

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bnprtr/problem"
)

func Example(w http.ResponseWriter, r *http.Request) {
	if err := handler(r); err != nil {
		var statusCode problem.StatusCode
		if !errors.As(err, &statusCode) {
			// unknown error!
			err = InternalServerError.New("unexpected error occurred.").WithInternalErrors(err)
			statusCode = InternalServerError.StatusCode()
		}
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(int(statusCode))
		json.NewEncoder(w).Encode(err)
	}
}

func handler(r *http.Request) error {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return InvalidJSON.New(err.Error()).WithInternalErrors(err)
	}

	if err := validate(req); err != nil {
		panic(err)
	}
	return nil
}
