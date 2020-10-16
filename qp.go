package qp

import (
	"fmt"
	"net/http"
)

// RequiredParams returns a map of query parameters that are required to be in the query.
// Returns an error if any required parameter is omitted.
func RequiredParams(r *http.Request, keys ...string) (map[string][]string, error) {
	qs := r.URL.Query()
	vs := make(map[string][]string)

	for _, p := range keys {
		switch v, ok := qs[p]; ok {
		case true:
			vs[p] = v
		case false:
			return nil, fmt.Errorf("missing required query parameter '%s'", p)
		}
	}

	return vs, nil
}

// RequiredParam returns a map of query parameters that are required to be in the query.
// The map will only contain the first occurrence of a query parameter.
// Returns an error if any required parameter is omitted.
func RequiredParam(r *http.Request, keys ...string) (map[string]string, error) {
	qs := r.URL.Query()
	vs := make(map[string]string)

	for _, p := range keys {
		switch v, ok := qs[p]; ok {
		case true:
			vs[p] = v[0]
		case false:
			return nil, fmt.Errorf("missing required query parameter '%s'", p)
		}
	}

	return vs, nil
}

// Params returns a map of query parameter values with the given query keys.
func Params(r *http.Request, keys ...string) map[string][]string {
	qs := r.URL.Query()
	var values map[string][]string

	for _, p := range keys {
		if v, ok := qs[p]; ok {
			if values == nil {
				values = make(map[string][]string)
			}

			values[p] = v
		}
	}

	return values
}

func Param(r *http.Request, keys ...string) map[string]string {
	qs := r.URL.Query()
	var values map[string]string

	for _, ps := range keys {
		if p, ok := qs[ps]; ok {
			if values == nil {
				values = make(map[string]string)
			}

			values[ps] = p[0]
		}
	}

	return values
}
