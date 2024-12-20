package main

import (
	"encoding/json"
	"os"
)

type Storange[T any] struct {
	FileName string
}

func NewStorange[T any](fileName string) *Storange[T] {
	return &Storange[T]{FileName: fileName}
}

func (s *Storange[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storange[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}