package services

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type TomlService struct{}

func NewTomlService() *TomlService {
	return &TomlService{}
}

func (ts *TomlService) AddExtension(fileName string) string {
	return fileName + ".toml"
}

func (ts *TomlService) CreateFile(data interface{}, path string, fileName string) error {
	tomlData, err := toml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal toml: %w", err)
	}

	fullPathFile := filepath.Join(path, ts.AddExtension(fileName))
	tempFile := fullPathFile + ".tmp"

	file, err := os.Create(tempFile)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}

	if _, err = file.Write(tomlData); err != nil {
		file.Close()
		os.Remove(tempFile)
		return fmt.Errorf("failed to write toml data: %w", err)
	}
	file.Close()

	if err := os.Rename(tempFile, fullPathFile); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to atomically save toml file: %w", err)
	}

	return nil
}

func (ts *TomlService) ReadFile(path string, fileName string) ([]byte, error) {
	fullPathFile := filepath.Join(path, ts.AddExtension(fileName))
	data, err := os.ReadFile(fullPathFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return data, nil
}

func (ts *TomlService) ParseFile(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}
