package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type JsonService struct {
}

func NewJsonService() *JsonService {
	return &JsonService{}
}

func (js *JsonService) AddExtension(fileName string) string {
	return fileName + ".json"
}

func (js *JsonService) CreateFile(data interface{}, path string, fileName string) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	fullPathFile := filepath.Join(path, js.AddExtension(fileName))
	file, err := os.Create(fullPathFile)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write(jsonData); err != nil {
		return err
	}

	return nil
}

func (js *JsonService) ReadFile(path string, fileName string) ([]byte, error) {
	fullPathFile := filepath.Join(path, js.AddExtension(fileName))
	data, err := os.ReadFile(fullPathFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return data, nil
}

func (js *JsonService) ParseFile(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
