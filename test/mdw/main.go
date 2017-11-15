package main

import (
	"context"
	"fmt"
	"net/http"
)

const requestIDKey = "rid"

func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
	reqID := req.Header.Get("X-Request-ID")
	if reqID == "" {
		reqID = "0"
	}
	return context.WithValue(ctx, requestIDKey, reqID)
}

func requestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

func middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := newContextWithRequestID(req.Context(), req)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func h(w http.ResponseWriter, req *http.Request) {
	reqID := requestIDFromContext(req.Context())
	fmt.Fprintln(w, "Request ID: ", reqID)
	return
}

func main() {
	http.Handle("/", middleWare(http.HandlerFunc(h)))
	http.ListenAndServe(":9201", nil)
}
