package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type UppercaseRequest struct {
	S string `json:"s"`
}

type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	V int `json:"v"`
}

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)

		v, err := svc.Uppercase(req.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}

		return UppercaseResponse{v, ""}, nil
	}

}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)

		v := svc.Count(req.S)

		return CountResponse{v}, nil
	}
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UppercaseRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CountRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
