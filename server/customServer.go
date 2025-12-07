package server

import (
	"fmt"
	"net/http"
)

type CustomHttpServer struct{}

func ServeHttp(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		panic(err)
	}
}
