package main

import (
	"context"
	"error-aws-lambda/lib/errors"
	"github.com/Bancar/uala-bis-go-dependencies/v2/logger"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID string `json:"id"`
}

var (
	ErrSimpleError = errors.New("simple error...")
	ErrHTTPError   = errors.New("http error")
	ErrDetailError = errors.New("error with detail")
)

func main() {
	r := &Request{ID: "detail"}

	lambda.Start(EventHandler(nil, r))
}

func EventHandler(ctx context.Context, req *Request) error {
	log := logger.New(ctx)
	err := Service(ctx, req)
	if err != nil {
		log.Error(err.Error())
	}
	return nil
}

func Service(ctx context.Context, req *Request) error {
	switch req.ID {
	case "ok":
		return nil
	case "simple":
		return ErrSimpleError
	case "http":
		err := ErrHTTPError.Code(400)
		return err
	case "detail":
		err := ErrDetailError.Code(500).Detail("detail n1").Detail("detail n3")
		return err
	}

	return nil
}
