package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

type LlmResponse struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Response           string    `json:"response"`
	Done               bool      `json:"done"`
	Context            []int     `json:"context"`
	TotalDuration      int64     `json:"total_duration"`
	LoadDuration       int64     `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int64     `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int64     `json:"eval_duration"`
}

func promptLLm(prompt string) LlmResponse {

	var llmResponse LlmResponse

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{
			"model": "tinydolphin",
			"prompt": "` + prompt + `",
			"stream": false
		  }`).
		SetResult(&LlmResponse{}).
		Post("http://localhost:11434/api/generate")

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(resp.Body(), &llmResponse)

	return llmResponse
}
