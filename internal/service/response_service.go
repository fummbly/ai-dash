package service

import "www.github.com/fummbly/ai-dash/internal/domain"

type ResponseService struct {
	endpoint domain.ResponseInterface
}

func NewResponseService(endpoint domain.ResponseInterface) *ResponseService {
	return &ResponseService{
		endpoint: endpoint,
	}
}

func (s *ResponseService) Generate(res chan domain.Response) error {
	return s.endpoint.Generate(res)
}
