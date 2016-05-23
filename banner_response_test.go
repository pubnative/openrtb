package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BannerResponse", func() {
	It("should parse banner adm", func() {
		actual, err := ParseBannerAdmBytes(testFixtures.simpleBannerResponse)

		Expect(err).NotTo(HaveOccurred())
		Expect(actual.ClickUrl).To(Equal("http://tracking.network.example.com/click"))
		Expect(actual.ImageUrl).To(Equal("http://cdn.image.example.com/image.jpg"))
	})
})
