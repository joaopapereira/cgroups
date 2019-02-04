package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

type EmptyHierarchy struct {
	OnFolderActioner
}

func (emptyHierarchy *EmptyHierarchy) OnFolderAction(path string, fileInformation os.FileInfo, err error) error {
	return nil
}

type FileStructureReadStub struct {
	folderStructure []string
}

func (reader *FileStructureReadStub) ReadFolder(path string, onEachFolderAction OnFolderActioner) {
	var err error
	for _, currentPath := range reader.folderStructure {
		err := onEachFolderAction.OnFolderAction(currentPath, nil, err)
		if err != nil {
			panic(err)
		}
	}
}

var _ = Describe("HierarchyLoad", func() {
	var (
		hierarchyLoad *HierarchyLoader
	)
	BeforeEach(func() {
	})

	Describe("When the hierarchy is empty", func() {
		It("returns a hierarchy without children", func() {
			hierarchyLoad = NewHierarchyLoader(&FileStructureReadStub{
				folderStructure: []string{
					"/some/specific/path",
				},
			})
			hierarchy := hierarchyLoad.HierarchyLoad("/some/specific/path")
			Expect(hierarchy.pathToHierarchy).To(Equal("/some/specific/path"))
			Expect(hierarchy.root.children).To(BeEmpty())
		})
	})

	Describe("When the hierarchy as a single empty folder", func() {
		var (
			hierarchy Hierarchy
		)
		BeforeEach(func() {
			hierarchyLoad = NewHierarchyLoader(&FileStructureReadStub{
				folderStructure: []string{
					"/some/specific/path",
					"/some/specific/path/subdir",
				},
			})
			hierarchy = hierarchyLoad.HierarchyLoad("/some/specific/path")
		})

		It("as the correct path to the hierarchy", func() {
			Expect(hierarchy.pathToHierarchy).To(Equal("/some/specific/path"))
		})

		It("returns a hierarchy one child", func() {
			Expect(hierarchy.root.children).To(Not(BeEmpty()))
		})
	})

	Describe("When the hierarchy as a single folder with one file", func() {
		var (
			hierarchy Hierarchy
		)
		BeforeEach(func() {
			hierarchyLoad = NewHierarchyLoader(&FileStructureReadStub{
				folderStructure: []string{
					"/some/specific/path",
					"/some/specific/path/subdir",
					"/some/specific/path/subdir/filename",
				},
			})
			hierarchy = hierarchyLoad.HierarchyLoad("/some/specific/path")
		})

		It("as the correct path to the hierarchy", func() {
			Expect(hierarchy.pathToHierarchy).To(Equal("/some/specific/path"))
		})

		It("returns a hierarchy one child", func() {
			Expect(hierarchy.root.children).To(Not(BeEmpty()))
		})
	})
})
