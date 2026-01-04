// Package domain for describing models and interfaces for use
package domain

import "time"

type Model struct {
	Name       string    `json:"name"`
	Model      string    `json:"model"`
	ModifiedAt time.Time `json:"modified_at"`
	Size       int       `json:"size"`
	Digest     string    `json:"digest"`
	Details    struct {
		ParentModel       string   `json:"parent_model"`
		Format            string   `json:"format"`
		Family            string   `json:"family"`
		Families          []string `json:"families"`
		ParameterSize     string   `json:"parameter_size"`
		QuantizationLevel string   `json:"quantization_level"`
	} `json:"details"`
}

type Models struct {
	Models []Model `json:"models"`
}

type ModelInterface interface {
	ListModels() (Models, error)
	ListRunningModels() ([]*Model, error)
	PullModel(name string) error
	ModelDetails(name string) (*Model, error)
	DeleteModel(name string) error
}
