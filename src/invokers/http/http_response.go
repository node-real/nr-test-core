package http

import "net/http"

type Response struct {
	Code   int
	Body   string
	Time   int64
	Header http.Header
}
