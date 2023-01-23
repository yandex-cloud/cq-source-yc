package utils

import "net/http"

type InterceptFunc func(*http.Request) error

type InterceptTransport struct {
	T http.RoundTripper
	f InterceptFunc
}

func (it *InterceptTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	err := it.f(req)
	if err != nil {
		return nil, err
	}
	return it.T.RoundTrip(req)
}

func NewInterceptTransport(T http.RoundTripper, f InterceptFunc) *InterceptTransport {
	if T == nil {
		T = http.DefaultTransport
	}
	return &InterceptTransport{T, f}
}
