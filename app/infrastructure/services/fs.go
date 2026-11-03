package services

import (
	"os"
	"path/filepath"

	"github.com/sneuder/filesystem"
)

type FsService struct {
}

func NewFsService() *FsService {
	return &FsService{}
}

func (fs *FsService) ReadTextFile(path string, fileName string) (string, error) {
	filePath := fs.JoinPaths([]string{path, fileName})
	content, err := filesystem.Read(filePath)
	return content, err
}

func (fs *FsService) CreateDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}

func (fs *FsService) RemoveDirectory(path string) error {
	return os.RemoveAll(path)
}

func (fs *FsService) ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (fs *FsService) GetHomeDirectory() string {
	return fs.GetHomeDir()
}

func (fs *FsService) GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = os.Getenv("HOME")
		if homeDir == "" {
			homeDir = "."
		}
	}

	return homeDir
}

func (fs *FsService) GetFolderPath(folders []string) string {
	return filepath.Join(fs.GetHomeDir(), fs.JoinPaths(folders))
}

func (fs *FsService) JoinPaths(paths []string) string {
	return filepath.Join(paths...)
}

func (fs *FsService) ExistFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
