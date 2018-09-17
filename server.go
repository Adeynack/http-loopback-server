package http_loopback_server

import (
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Serve(addr string) {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("ERROR reading body: %+v", err)
	}

	a := &Answer{
		Host:          request.Host,
		Path:          request.RequestURI,
		Verb:          request.Method,
		ContentLength: request.ContentLength,
		Headers:       request.Header,
		Body:          string(bodyBytes),
	}

	writer.Header().Set("Content-Type", "application/json")
	jsonEnc := json.NewEncoder(writer)
	jsonEnc.Encode(a)

	logExchange(a)
}

func logExchange(a *Answer) {
	// Convert to YAML
	yamlBuf := new(bytes.Buffer)
	yamlBuf.WriteRune('\n')
	yamlEnc := yaml.NewEncoder(yamlBuf)
	yamlEnc.Encode(a)

	// Indent the output (each line)
	sb := strings.Builder{}
	var line string
	var err error
	for line, err = yamlBuf.ReadString('\n'); err == nil; line, err = yamlBuf.ReadString('\n') {
		sb.WriteString("    ")
		sb.WriteString(line)
	}
	if err != nil && err != io.EOF {
		log.Printf("error indenting output: %s", err)
	}

	log.Println(sb.String())
}

type Answer struct {
	Host          string              `json:"host"`
	Path          string              `json:"path"`
	Verb          string              `json:"verb"`
	ContentLength int64               `json:"content_length"`
	Headers       map[string][]string `json:"headers"`
	Body          string              `json:"body"`
}
