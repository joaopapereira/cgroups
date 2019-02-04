package core

import (
	"fmt"
	"os"
	"strings"
)

type FileStructureReader interface {
	ReadFolder(path string, onEachFolderAction OnFolderActioner)
}

func NewHierarchyLoader(reader FileStructureReader) *HierarchyLoader {
	return &HierarchyLoader{
		fileStructureReader: reader,
	}
}

type HierarchyLoader struct {
	OnFolderActioner
	fileStructureReader     FileStructureReader
	hierarchyOnFolderAction OnFolderActioner
	hierarchy               Hierarchy
}

func (loader HierarchyLoader) HierarchyLoad(path string) Hierarchy {
	loader.hierarchy = NewHierarchy(path)
	loader.fileStructureReader.ReadFolder(path, loader)
	return loader.hierarchy
}
func (loader HierarchyLoader) OnFolderAction(path string, fileInformation os.FileInfo, err error) error {
	if path == loader.hierarchy.pathToHierarchy {
		return nil
	}
	relativePath := path[len(loader.hierarchy.pathToHierarchy):]
	splitPath := strings.Split(relativePath, "/")
	currentHierarchyNode := loader.hierarchy.root
	for _, currentFolder := range splitPath {
		foundMatch := false
		for _, possibleMatch := range currentHierarchyNode.children {
			if possibleMatch.name == currentFolder {
				currentHierarchyNode = possibleMatch
				foundMatch = true
			}
		}
		if !foundMatch {
			newNode := &HierarchyNode{
				name: currentFolder,
			}
			currentHierarchyNode.children = append(currentHierarchyNode.children, newNode)
			currentHierarchyNode = newNode
		}
	}
	return nil
}

type LoadHierarchyOnFolderAction struct {
	rootOfHierarchy Hierarchy
}

func (s LoadHierarchyOnFolderAction) InsideFolder(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return err
}

func (s LoadHierarchyOnFolderAction) LoadInformationFromFile(path string, err error) error {
	return nil
}
