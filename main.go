package main

import (
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Lambda!"))
	})
	http.HandleFunc("/challenge", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write(body)
	})

	lambda.Start(httpadapter.NewV2(http.DefaultServeMux).ProxyWithContext)
}
