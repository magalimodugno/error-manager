package main

import (
	"context"
	stdErrors "errors"
	"fmt"

	"error-aws-lambda/lib/errors"

	"github.com/Bancar/uala-bis-go-dependencies/v2/logger"
)

type Request struct {
	ErrorID string `json:"error_id"`
}

var (
	ErrSimpleError      = errors.New("invalid account status")
	ErrResourceNotFound = errors.New("resource not found with specified id")
	ErrService          = errors.New("couldn't retrieve information")
)

func main() {
	err := EventHandler(nil, &Request{ErrorID: "service"})
	if err != nil {
		fmt.Printf("\n%s\n", err.Error())
	}
	return
	//lambda.Start(EventHandler(nil, r))
}

func EventHandler(ctx context.Context, req *Request) (e error) {
	defer func() {
		log := logger.New(ctx).WithField("account_id", "1234")

		if e != nil {
			log.Error(e.Error())

			e = errors.AsErrorMessage(e)
			return
		}

		log.Info("successfully ...")

	}()

	return Service(ctx, req)

}

func Service(ctx context.Context, req *Request) error {
	switch req.ErrorID {
	case "ok":
		return nil
	case "simple":
		return ErrSimpleError
	case "http":
		return ErrResourceNotFound.Status(404).
			Err(stdErrors.New("dynamodb not found error"))
	case "service":
		return ErrService.
			Err(stdErrors.New("service error")).
			Detail("code: ...").Detail("origin: ...")
	}

	return nil
}
