package core

import (
	"fmt"
	"os"
	"path/filepath"
)

type OnFolderActioner interface {
	OnFolderAction(path string, fileInformation os.FileInfo, err error) error
}

func ReadFileStructure(path string, onEachFolderAction OnFolderActioner) {
	filepath.Walk(path, onEachFolderAction.OnFolderAction)
}

type PrintOnFolderAction struct {}
func (s PrintOnFolderAction) InsideFolder(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return err
}