package ec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("big math helper functions", func() {
	Describe("#gt", func() {
		It("", func() {
			Expect(gt(W("1"), 0)).To(Equal(true))
			Expect(gt(W("1123123123123891023801928309128309182309812309128309182093281039281321"), 5)).To(Equal(true))
			Expect(gt(W("0"), 1)).To(Equal(false))
			Expect(gt(W("0"), 0)).To(Equal(false))
		})
	})

	Describe("#lt", func() {
		It("", func() {
			Expect(lt(W("1"), 0)).To(Equal(false))
			Expect(lt(W("1123123123123891023801928309128309182309812309128309182093281039281321"), 5)).To(Equal(false))
			Expect(lt(W("0"), 1)).To(Equal(true))
			Expect(lt(W("0"), 0)).To(Equal(false))
		})
	})

	Describe("#eq", func() {
		It("", func() {

			Expect(eq(W("1"), 0)).To(Equal(false))
			Expect(eq(W("1123123123123891023801928309128309182309812309128309182093281039281321"), 5)).To(Equal(false))
			Expect(eq(W("0"), 1)).To(Equal(false))
			Expect(eq(W("0"), 0)).To(Equal(true))
			Expect(eq(W("1"), 1)).To(Equal(true))
		})
	})
})
