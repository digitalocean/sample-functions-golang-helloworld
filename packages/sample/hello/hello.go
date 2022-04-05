package main

import (
	"errors"
	"fmt"
)

type Request struct {
	Location string `json:"location"`
}

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string `json:"body"`
}

func Main(in Request) (*Response, error) {
	if in.Location == "" {
		return nil, errors.New("location must be passed")
	}

	loc := fmt.Sprintf("Hello, from %s", in.Location)

	return &Response{
		Body: loc,
	}, nil
}
