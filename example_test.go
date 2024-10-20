package problem_test

import (
	"fmt"
	"net/url"

	"github.com/bnprtr/problem"
)

const (
	InvalidJSON         problem.StatusBadRequest          = "invalid json"
	InternalServerError problem.StatusInternalServerError = "internal server error"
	InvalidInput        problem.StatusBadRequest          = "invalid input"
)

type Request struct {
	Name string
	URL  string
}

func ExampleError() {
	request := Request{
		Name: "",
		URL:  "%+0",
	}
	if err := validate(request); err != nil {
		fmt.Println(err.Error())
	}
	// Output: {"title":"invalid input","message":"request input failed validation","internalError":[{"Op":"parse","URL":"%+0","Err":"%+0"}],"details":[{"field":"Name","value":"","reason":"required"},{"field":"URL","value":"%+0","reason":"parse \"%+0\": invalid URL escape \"%+0\""}],"statusCode":400}
}

func validate(r Request) error {
	err := InvalidInput.New("request input failed validation")
	var failed bool
	if r.Name == "" {
		failed = true
		err = err.WithDetails(problem.NewFieldDetail("Name", r.Name, "required"))
	}
	if r.URL != "" {
		_, e := url.Parse(r.URL)
		if e != nil {
			failed = true
			err = err.WithInternalErrors(e)
			err = err.WithDetails(problem.NewFieldDetail("URL", r.URL, e.Error()))
		}
	}
	if failed {
		return err
	}
	return nil
}
