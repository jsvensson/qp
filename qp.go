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
			return nil, fmt.Errorf("missing required parameter '%s'", p)
		}
	}

	return vs, nil
}

// RequiredParam returns a map of query parameters that are required to be in the query.
// Will only contain the first occurrence of a query parameter.
// Returns an error if any required parameter is omitted.
func RequiredParam(r *http.Request, keys ...string) (map[string]string, error) {
	qs := r.URL.Query()
	vs := make(map[string]string)

	for _, p := range keys {
		switch v, ok := qs[p]; ok {
		case true:
			vs[p] = v[0]
		case false:
			return nil, fmt.Errorf("missing required parameter '%s'", p)
		}
	}

	return vs, nil
}

// Params returns a slice of query parameter values with the given query key.
// Returns a nil slice if the query parameter was not found.
func Params(r *http.Request, key string) []string {
	qs := r.URL.Query()

	if q, ok := qs[key]; ok {
		return q
	}

	return nil
}
