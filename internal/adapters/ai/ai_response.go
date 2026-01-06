package ai

import (
	"encoding/json"
	"fmt"

	"www.github.com/fummbly/ai-dash/internal/adapters/http"
	"www.github.com/fummbly/ai-dash/internal/domain"
)

type AIResponseEndpoint struct {
	URL string
}

func NewAIResponseEndpoint(url string) *AIResponseEndpoint {
	return &AIResponseEndpoint{
		url,
	}
}

func (ai *AIResponseEndpoint) Generate(res chan domain.Response) error {
	generateURL := fmt.Sprintf("%s%s", ai.URL, "/generate")

	data := make(chan []byte)

	defer close(res)

	postData := `{
		"model": "gemma3:1b",
		"prompt": "Why is the sky blue?"
	}`

	// TODO change this to be a nested function in a go func
	go http.StreamPost(generateURL, "application/json", postData, data)

	var response domain.Response

	for {
		select {
		case _, ok := <-res:
			if !ok {
				return nil
			}
		case byteData, ok := <-data:
			if !ok {
				return nil
			}

			err := json.Unmarshal(byteData, &response)
			if err != nil {
				fmt.Printf("Failed to unmarshal response data: %v", err)

				return err
			}
			res <- response
		}
	}
}
