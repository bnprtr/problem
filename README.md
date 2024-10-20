# problem

An RFC 7807 compatible error library for Go.

## Overview

This library provides a structured way to handle HTTP errors in Go applications, compatible with [RFC 7807](https://tools.ietf.org/html/rfc7807). It defines custom error types for various HTTP status codes and provides methods to create and manipulate these errors.

## Features

- Custom error types for HTTP status codes.
- Methods to create new errors with messages.
- JSON serialization of errors.
- Logging support.

## Installation

To install the library, run:

```sh
go get -u github.com/bnprtr/problem
```

## Usage

To use `problem`, you should define error constants for the types of your errors. The constant should by typed as one of problem's Status types:

```go
package main

import (
  "errors"
  "net/http"
  "github.com/bnprtr/problem"
)

const (
  InvalidInput problem.StatusBadRequest = "invalid input"
  InvalidJSON problem.StatusBadRequest = "invalid json"
  InternalServerError problem.StatusInternalServerError = "internal server error"
)

type Request struct {
 Name string
 URL  string
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
 if err := handle(r); err != nil {
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

func (r *http.Request) error {
 var req Request
 if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
  return InvalidJSON.New(err.Error()).WithInternalErrors(err)
 }

 if err := validate(req); err != nil {
  panic(err)
 }
 return nil
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
```
