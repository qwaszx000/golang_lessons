package main

import (
	"fmt"
	"net/http"
)

func connection_handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Printf("New http connection\nLocation: %s\n\n", req.RequestURI)

	resp.Header().Add("X-Test-Header", "xyz")
	resp.Write([]byte("<h1>Hello</h1>"))
}

func start_http_server() {
	//http.Handle("/", http.HandlerFunc(connection_handler))

	http.ListenAndServe("127.0.0.1:8080", http.HandlerFunc(connection_handler))
}
