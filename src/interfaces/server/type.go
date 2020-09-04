package server

import (
	"context"
	"net/http"
)

type HttpParams interface {
	ByName(string) string
}

type HttpHandler func(context.Context, *http.Request, HttpParams) (interface{}, error)

type HttpMiddleware func(context.Context, []HttpMiddleware, *http.Request, HttpParams) (interface{}, error)

type HttpServer interface {
	GET(string, HttpHandler)
	POST(string, HttpHandler)
	PUT(string, HttpHandler)
	DELETE(string, HttpHandler)
	OPTIONS(string, HttpHandler)

	ServeHTTP(http.ResponseWriter, *http.Request)
	ServeFiles(string, http.FileSystem)

	AddMiddleware(HttpMiddleware)
}
