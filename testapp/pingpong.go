package testapp

import "net/http"

func  init() {
	http.HandleFunc("/test/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	})
}
