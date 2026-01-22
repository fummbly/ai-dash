// Package domain/response is for outlining the responses from the ai API
package domain

import "time"

type Response struct {
	Model      string    `json:"model"`
	CreatedAt  time.Time `json:"created_at"`
	Response   string    `json:"response"`
	Done       bool      `json:"done"`
	DoneReason string    `json:"done_reason"`
	Context    []int     `json:"context"`
}

type GenerateMessage struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	System string `json:"system,omitempty"`
	Think  string `json:"think,omitempty"`
}

type ResponseInterface interface {
	Generate(res chan Response, question string) error
}
