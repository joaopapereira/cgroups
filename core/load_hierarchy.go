package core

import (
	"fmt"
	"os"
	"strings"
)

type FileStructureReader interface {
	ReadFolder(path string, onEachFolderAction OnFolderActioner)
}

func NewHierarchyLoader(reader FileStructureReader, shouldAdd ShouldAddToHierarchy) *HierarchyLoader {
	return &HierarchyLoader{
		fileStructureReader: reader,
		addToHierarchy:      shouldAdd,
	}
}

type ShouldAddToHierarchy interface {
	ShouldAdd(*HierarchyNode, string, os.FileInfo) bool
}

type HierarchyLoader struct {
	OnFolderActioner
	fileStructureReader     FileStructureReader
	hierarchyOnFolderAction OnFolderActioner
	hierarchy               Hierarchy
	addToHierarchy          ShouldAddToHierarchy
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
		if len(currentFolder) == 0 {
			continue
		}
		foundMatch := false
		for _, possibleMatch := range currentHierarchyNode.children {
			if possibleMatch.name == currentFolder {
				currentHierarchyNode = possibleMatch
				foundMatch = true
			}
		}

		if !foundMatch && loader.addToHierarchy.ShouldAdd(currentHierarchyNode, path, fileInformation) {
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
