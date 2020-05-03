// Code generated by kok; DO NOT EDIT.
// github.com/RussellLuo/kok

package profilesvc

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(svc Service) chi.Router {
	r := chi.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(errorEncoder),
	}

	r.Method(
		"DELETE", "/profiles/{profileID}/addresses/{addressID}",
		kithttp.NewServer(
			MakeEndpointOfDeleteAddress(svc),
			decodeDeleteAddressRequest,
			encodeGenericResponse,
			options...,
		),
	)
	r.Method(
		"DELETE", "/profiles/{id}",
		kithttp.NewServer(
			MakeEndpointOfDeleteProfile(svc),
			decodeDeleteProfileRequest,
			encodeGenericResponse,
			options...,
		),
	)
	r.Method(
		"GET", "/profiles/{profileID}/addresses/{addressID}",
		kithttp.NewServer(
			MakeEndpointOfGetAddress(svc),
			decodeGetAddressRequest,
			encodeGenericResponse,
			options...,
		),
	)
	r.Method(
		"GET", "/profiles/{id}/addresses",
		kithttp.NewServer(
			MakeEndpointOfGetAddresses(svc),
			decodeGetAddressesRequest,
			encodeGenericResponse,
			options...,
		),
	)
	r.Method(
		"GET", "/profiles/{id}",
		kithttp.NewServer(
			MakeEndpointOfGetProfile(svc),
			decodeGetProfileRequest,
			encodeGenericResponse,
			options...,
		),
	)
	r.Method(
		"PATCH", "/profiles/{id}",
		kithttp.NewServer(
			MakeEndpointOfPatchProfile(svc),
			decodePatchProfileRequest,
			encodeGenericResponse,
			options...,
		),
	)
	r.Method(
		"POST", "/profiles/{profileID}/addresses",
		kithttp.NewServer(
			MakeEndpointOfPostAddress(svc),
			decodePostAddressRequest,
			encodeGenericResponse,
			options...,
		),
	)
	r.Method(
		"POST", "/profiles",
		kithttp.NewServer(
			MakeEndpointOfPostProfile(svc),
			decodePostProfileRequest,
			encodeGenericResponse,
			options...,
		),
	)
	r.Method(
		"PUT", "/profiles/{id}",
		kithttp.NewServer(
			MakeEndpointOfPutProfile(svc),
			decodePutProfileRequest,
			encodeGenericResponse,
			options...,
		),
	)

	return r
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	// err2code (signature: func(error) int) must be provided in this package,
	// to transform a business error to an HTTP code!
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

type errorWrapper struct {
	Error string `json:"error"`
}

func decodeDeleteAddressRequest(_ context.Context, r *http.Request) (interface{}, error) {
	profileID := chi.URLParam(r, "profileID")

	addressID := chi.URLParam(r, "addressID")

	return &DeleteAddressRequest{
		ProfileID: profileID,
		AddressID: addressID,
	}, nil
}

func decodeDeleteProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")

	return &DeleteProfileRequest{
		Id: id,
	}, nil
}

func decodeGetAddressRequest(_ context.Context, r *http.Request) (interface{}, error) {
	profileID := chi.URLParam(r, "profileID")

	addressID := chi.URLParam(r, "addressID")

	return &GetAddressRequest{
		ProfileID: profileID,
		AddressID: addressID,
	}, nil
}

func decodeGetAddressesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")

	return &GetAddressesRequest{
		Id: id,
	}, nil
}

func decodeGetProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")

	return &GetProfileRequest{
		Id: id,
	}, nil
}

func decodePatchProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")

	var body struct {
		Profile Profile `json:"profile"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &PatchProfileRequest{
		Id:      id,
		Profile: body.Profile,
	}, nil
}

func decodePostAddressRequest(_ context.Context, r *http.Request) (interface{}, error) {
	profileID := chi.URLParam(r, "profileID")

	var body struct {
		A Address `json:"a"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &PostAddressRequest{
		ProfileID: profileID,
		A:         body.A,
	}, nil
}

func decodePostProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Profile Profile `json:"profile"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &PostProfileRequest{
		Profile: body.Profile,
	}, nil
}

func decodePutProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, "id")

	var body struct {
		Profile Profile `json:"profile"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &PutProfileRequest{
		Id:      id,
		Profile: body.Profile,
	}, nil
}

func encodeGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}