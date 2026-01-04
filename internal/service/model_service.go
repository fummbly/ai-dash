// Package service for implementing code for business logic
package service

import "www.github.com/fummbly/ai-dash/internal/domain"

type ModelService struct {
	endpoint domain.ModelInterface
}

func NewModelService(endpoint domain.ModelInterface) *ModelService {
	return &ModelService{
		endpoint: endpoint,
	}
}

func (s *ModelService) List() (domain.Models, error) {
	return s.endpoint.ListModels()
}
