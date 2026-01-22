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

func (ai *AIResponseEndpoint) Generate(res chan domain.Response, question string) error {
	generateURL := fmt.Sprintf("%s%s", ai.URL, "/generate")

	data := make(chan []byte)

	defer close(res)

	message := domain.GenerateMessage{
		Model:  "gemma3:1b",
		Prompt: question,
	}

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonMessage))

	// TODO change this to be a nested function in a go func
	go http.StreamPost(generateURL, "application/json", string(jsonMessage), data)

	var response domain.Response

	for {
		select {
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
