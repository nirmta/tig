package services

import (
	"log"
	"os"
	"os/exec"
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

func (fs *FsService) CreateDirectory(path string) {
	exec.Command("mkdir", "-p", path).Run()
}

func (fs *FsService) RemoveDirectory(path string) {
	exec.Command("rm", "-r", path).Run()
}

func (fs *FsService) GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
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
