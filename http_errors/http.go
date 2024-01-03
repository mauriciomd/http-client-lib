package http_errors

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// Bad Request
// {"message":"required field is missing","code":1032,"extras":{"failed_fields":["id","replicas","image","ports"]}}

type NotFoundResource struct {
	Id       uuid.UUID
	Resource string
}

type ResponseError struct {
	StatusCode int
	Message    string `json:"message"`
	Code       int    `json:"code"`
}

type BadRequestError struct {
	StatusCode int
	response
}

type response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Extras  extras `json:"extras"`
}

type extras struct {
	FailedField []string `json:"failed_fields"`
}

func NewNotFound(id uuid.UUID, resource string) NotFoundResource {
	return NotFoundResource{id, resource}
}

func (n NotFoundResource) Error() string {
	return fmt.Sprintf("%s with ID=%s not found", n.Resource, n.Id)
}

func FromHTTPResponse(res *http.Response) ResponseError {
	var r response
	body, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &r); err != nil {
		return ResponseError{}
	}

	return ResponseError{
		StatusCode: res.StatusCode,
		Message:    r.Message,
		Code:       r.Code,
	}
}

func FromBadRequest(res *http.Response) BadRequestError {
	var r response
	body, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &r); err != nil {
		return BadRequestError{}
	}

	return BadRequestError{
		StatusCode: res.StatusCode,
		response: response{
			Message: r.Message,
			Code:    r.Code,
			Extras:  r.Extras,
		},
	}
}

func (err ResponseError) Error() string {
	return fmt.Sprintf("StatusCode=%d, Message=%s", err.StatusCode, err.Message)
}

func (err BadRequestError) Error() string {
	return fmt.Sprintf("StatusCode=%d, Message=%s, Fields=%s", err.StatusCode, err.Message, strings.Join(err.Extras.FailedField, ", "))
}
