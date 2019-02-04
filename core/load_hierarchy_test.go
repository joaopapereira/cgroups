package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

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

type ShouldAddToHierarchyStub struct {
	ShouldAddToHierarchy
	ShouldAddAnswers []bool
	currentResponse  int
}

func (addToHierarchy *ShouldAddToHierarchyStub) ShouldAdd(currentParentNode *HierarchyNode, fullPath string, fileInfo os.FileInfo) bool {
	addToHierarchy.currentResponse++
	return addToHierarchy.ShouldAddAnswers[addToHierarchy.currentResponse-1]
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
			},
				&ShouldAddToHierarchyStub{
					ShouldAddAnswers: []bool{true, true},
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
			},
				&ShouldAddToHierarchyStub{
					ShouldAddAnswers: []bool{true, true},
				})
			hierarchy = hierarchyLoad.HierarchyLoad("/some/specific/path")
		})

		It("as the correct path to the hierarchy", func() {
			Expect(hierarchy.pathToHierarchy).To(Equal("/some/specific/path"))
		})

		It("returns a hierarchy one child", func() {
			Expect(hierarchy.root.children).To(Not(BeEmpty()))
		})

		It("set the full folder name for the root", func() {
			Expect(hierarchy.root.name).To(Equal("/some/specific/path"))
		})
	})

	Describe("When the hierarchy as a single folder with one file and file should not be added", func() {
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
			},
				&ShouldAddToHierarchyStub{
					ShouldAddAnswers: []bool{true, false},
				})
			hierarchy = hierarchyLoad.HierarchyLoad("/some/specific/path")
		})


		It("as the correct path to the hierarchy", func() {
			Expect(hierarchy.root.children[0].name).To(Equal("subdir"))
		})

		It("returns a hierarchy one child", func() {
			Expect(hierarchy.root.children[0].children).To(BeEmpty())
		})
	})
})
