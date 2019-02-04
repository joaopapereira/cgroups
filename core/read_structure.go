package core

import (
	"fmt"
	"os"
	"path/filepath"
)

type OnFolderActioner interface {
	OnFolderAction(path string, fileInformation os.FileInfo, err error) error
}

func NewFileStructureRead(reader func(string, OnFolderActioner)) *FileStructureRead {
	return &FileStructureRead{
		fileStructureReader: reader,
	}
}

func NewDefaultFileStructureRead() *FileStructureRead {
	return &FileStructureRead{
		fileStructureReader: ReadFileStructure,
	}
}

type FileStructureRead struct {
	fileStructureReader func(path string, onEachFolderAction OnFolderActioner)
}

func (fileStructureRead *FileStructureRead) ReadFolder(path string, onEachFolderAction OnFolderActioner) {
	ReadFileStructure(path, onEachFolderAction)
}

func ReadFileStructure(path string, onEachFolderAction OnFolderActioner) {
	filepath.Walk(path, onEachFolderAction.OnFolderAction)
}

type PrintOnFolderAction struct{}

func (s PrintOnFolderAction) InsideFolder(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return err
}
