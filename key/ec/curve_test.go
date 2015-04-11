package ec

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Curve", func() {

	expectModularInverse := func(a, b string) {
		var title = fmt.Sprintf("calculates the inverse of %s to be %s", a, b)

		It(title, func() {
			curve := NewCurve()
			r := curve.ModularInverse(W(a))
			Expect(r.String()).To(Equal(b))
		})
	}

	Describe("#ModularInverse", func() {
		expectModularInverse("1", "1")
		expectModularInverse("2", "57896044618658097711785492504343953926634992332820282019728792003954417335832")

		expectModularInverse(
			"103588105881134521297392950363556582805271693099531759694119316285034955766421",
			"45900550174020825671461307549774270965382867892451705830695672428573753728672",
		)
	})
})
