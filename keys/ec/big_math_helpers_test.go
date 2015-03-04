package ec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("big math helper functions", func() {
	Describe("#Gt", func() {
		It("", func() {
			Expect(Gt(B("1"), 0)).To(Equal(true))
			Expect(Gt(B("1123123123123891023801928309128309182309812309128309182093281039281321"), 5)).To(Equal(true))
			Expect(Gt(B("0"), 1)).To(Equal(false))
			Expect(Gt(B("0"), 0)).To(Equal(false))
		})
	})

	Describe("#Lt", func() {
		It("", func() {
			Expect(Gt(B("1"), 0)).To(Equal(true))
			Expect(Lt(B("1"), 0)).To(Equal(false))
			Expect(Lt(B("1123123123123891023801928309128309182309812309128309182093281039281321"), 5)).To(Equal(false))
			Expect(Lt(B("0"), 1)).To(Equal(true))
			Expect(Lt(B("0"), 0)).To(Equal(false))
		})
	})

	Describe("#Eq", func() {
		It("", func() {

			Expect(Eq(B("1"), 0)).To(Equal(false))
			Expect(Eq(B("1123123123123891023801928309128309182309812309128309182093281039281321"), 5)).To(Equal(false))
			Expect(Eq(B("0"), 1)).To(Equal(false))
			Expect(Eq(B("0"), 0)).To(Equal(true))
			Expect(Eq(B("1"), 1)).To(Equal(true))
		})
	})
})


