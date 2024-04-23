package main

import (
	"log"

	"github.com/go-resty/resty/v2"
)

func promptLLm() *resty.Response {
	client := resty.New()

	resp, err := client.R().Get("http://127.0.0.1:11434/api/tags")

	if err != nil {
		log.Fatal(err)
	}

	return resp
}
