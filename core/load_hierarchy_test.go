package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)



var _ = Describe("LoadHierarchy", func() {
	Describe("When the hierarchy is empty", func() {
		It("returns a hierarchy without children", func() {
			hierarchy := LoadHierarchy("/some/specific/path")
			Expect(hierarchy.pathToHierarchy).To(Equal("/some/specific/path"))
			Expect(hierarchy.root.children).To(BeEmpty())
		})
	})
	Describe("When the hierarchy as a single folder", func() {
		It("returns a hierarchy one child", func() {

			hierarchy := LoadHierarchy("/some/specific/path")
			Expect(hierarchy.pathToHierarchy).To(Equal("/some/specific/path"))
			Expect(hierarchy.root.children).To(BeEmpty())
		})
	})
})
