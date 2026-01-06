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

type ResponseInterface interface {
	Generate(res chan Response) error
}
