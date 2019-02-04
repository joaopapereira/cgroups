package core_test

import (
	"fmt"
	"github.com/joaopapereira/cgroups/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"log"
	"os"
)

type OnFolderActionMock struct {
	expectedFolderName []string
	currentIteration   int
}

func (self *OnFolderActionMock) OnFolderAction(path string, fileInformation os.FileInfo, err error) error {
	Expect(path).To(Equal(self.expectedFolderName[self.currentIteration]), fmt.Sprintf("Call %d: path parameter are incorrect", self.currentIteration))
	self.currentIteration++
	return err
}

func NewOnFolderActionMock(expectedFolderNames []string) core.OnFolderActioner {
	return &OnFolderActionMock{
		expectedFolderNames,
		0,
	}
}

var _ = Describe("ReadStructure", func() {
	Describe("ReadFileStructure", func() {
		var (
			rootDirectory string
			subdir1 string
			subdir2 string
			subdir11 string
			subdir21 string
			subdir211 string
		)
		BeforeSuite(func() {
			rootDirectory = CreateDirectory("", "root")
			subdir1 = CreateDirectory(rootDirectory, "subdir1.")
			subdir11 = CreateDirectory(subdir1, "subdir1.1.")

			subdir2 = CreateDirectory(rootDirectory, "subdir2.")
			subdir21 = CreateDirectory(subdir2, "subdir2.1.")
			subdir211 = CreateDirectory(subdir21, "subdir2.1.1")
		})
		AfterSuite(func() {
			defer os.RemoveAll(rootDirectory)
		})

		It("Check that all the folders are reached", func() {
			structureReader := core.NewDefaultFileStructureRead()
			structureReader.ReadFolder(rootDirectory, NewOnFolderActionMock(
				[]string{
					rootDirectory,
					subdir1,
					subdir11,
					subdir2,
					subdir21,
					subdir211,
				}))
		})
	})
})

func CreateDirectory(path string, prefix string) string {
	dir, err := ioutil.TempDir(path, prefix)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
