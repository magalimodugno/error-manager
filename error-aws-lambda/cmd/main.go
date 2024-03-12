package main

import (
	"context"
	"encoding/json"
	stdErrors "errors"
	"fmt"

	"error-aws-lambda/lib/errors"

	"github.com/Bancar/uala-bis-go-dependencies/v2/logger"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID string `json:"id"`
}

var (
	ErrSimpleError      = errors.New("simple error...")
	ErrResourceNotFound = errors.New("resource not found with specified id")
	ErrDetailError      = errors.New("error with detail")
)

func main() {
	r := &Request{ID: "http"}

	lambda.Start(EventHandler(nil, r))
}

func EventHandler(ctx context.Context, req *Request) (e error) {
	defer func() {
		log := logger.New(ctx).WithField("account_id", "1234")

		if e != nil {
			log.Error(e.Error())
			data, _ := json.Marshal(e)
			e = &errors.ErrorMessage{
				Message: string(data),
			}

			fmt.Printf("%+v\n", e) //only for local debug
			return
		}

		log.Info("successfully ...")

	}()

	return Service(ctx, req)

}

func Service(ctx context.Context, req *Request) error {
	switch req.ID {
	case "ok":
		return nil
	case "simple":
		return ErrSimpleError
	case "http":
		err := ErrResourceNotFound.Code(404).
			ErrMsg(stdErrors.New("dynamodb not found error")).
			Detail("origin: dynamodb_client").
			Detail(fmt.Sprintf("code: %d", 1001))
		return err
	case "detail":
		err := ErrDetailError.Code(500).Detail("detail n1").Detail("detail n3")
		return err
	}

	return nil
}
