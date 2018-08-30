package http_loopback_server

import (
	"net/http"
	"log"
	"encoding/json"
	"bytes"
)

func Serve(addr string) {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	a := &Answer{
		Host:          request.Host,
		Path:          request.RequestURI,
		Verb:          request.Method,
		ContentLength: request.ContentLength,
		Headers:       request.Header,
	}
	writer.Header().Set("Content-Type", "application/json")

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	enc.Encode(a)

	log.Println(buf)

	writer.Write(buf.Bytes())
}

type Answer struct {
	Host          string              `json:"host"`
	Path          string              `json:"path"`
	Verb          string              `json:"verb"`
	ContentLength int64               `json:"content_length"`
	Headers       map[string][]string `json:"headers"`
}
