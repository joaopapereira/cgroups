package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hierarchy", func() {
	Describe("New hierarchy", func() {
		It("Creates a new hierarchy", func() {
			hierarchy := NewHierarchy("/some/path")
			Expect(hierarchy.pathToHierarchy).To(Equal("/some/path"))
			Expect(hierarchy.root.children).To(HaveLen(0))
		})
	})
})
