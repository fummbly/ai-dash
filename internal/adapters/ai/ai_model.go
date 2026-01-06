// Package ai for interacting with ai backend server
package ai

import (
	"encoding/json"
	"fmt"

	"www.github.com/fummbly/ai-dash/internal/adapters/http"
	"www.github.com/fummbly/ai-dash/internal/domain"
)

type AIModelEndpoint struct {
	URL string
}

func NewAIModelEnpoint(url string) *AIModelEndpoint {
	return &AIModelEndpoint{
		URL: url,
	}
}

func (ai *AIModelEndpoint) ListModels() (domain.Models, error) {
	allModelsURL := fmt.Sprintf("%s%s", ai.URL, "/tags")

	data, err := http.BasicGet(allModelsURL)
	if err != nil {
		return domain.Models{}, err
	}

	var models domain.Models

	err = json.Unmarshal(data, &models)
	if err != nil {
		fmt.Printf("Failed to unmarshal json: %v err; %v", string(data), err)

		return domain.Models{}, err
	}

	return models, nil
}

func (ai *AIModelEndpoint) ListRunningModels() ([]*domain.Model, error) {
	return []*domain.Model{}, nil
}

func (ai *AIModelEndpoint) PullModel(name string) error {
	return nil
}

func (ai *AIModelEndpoint) ModelDetails(name string) (*domain.Model, error) {
	return &domain.Model{}, nil
}

func (ai *AIModelEndpoint) DeleteModel(name string) error {
	return nil
}
