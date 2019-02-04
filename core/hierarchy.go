package core

type HierarchyNode struct {
	children []*HierarchyNode
	name     string
}

type Hierarchy struct {
	pathToHierarchy string
	root            *HierarchyNode
}

func NewHierarchy(rootFolderPath string) Hierarchy {
	return Hierarchy{
		pathToHierarchy: rootFolderPath,
		root: &HierarchyNode{
			children: []*HierarchyNode{},
			name: rootFolderPath,
		},
	}
}
