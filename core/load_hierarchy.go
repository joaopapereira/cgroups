package core

func LoadHierarchy(path string) Hierarchy {
	hierarchy := NewHierarchy(path)
	return hierarchy
}
