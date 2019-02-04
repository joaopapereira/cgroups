package information

import (
	"fmt"
	"github.com/joaopapereira/cgroups/core"
	"os"
)

type CgroupInformation struct {
	core.HierarchyNode
	Pids []int
}

func Load() {
	hierarchyLoader := core.NewHierarchyLoader(
		core.NewDefaultFileStructureRead(),
		&IsCgroupFolder{},
	)
	hierarchyLoader.HierarchyLoad("/sys/fs/cgroup")
}

type IsCgroupFolder struct {
}

func (isCgroupFolder *IsCgroupFolder) ShouldAdd(currentNode *core.HierarchyNode, fullPath string, fileInformation os.FileInfo) bool {
	if fileInformation.IsDir() {
		fmt.Printf("Found a CGroup in: %s\n", fullPath)
		return true
	}
	return false
}
