package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AltNativeResponse", func() {
	var iptr = func(i int) *int { return &i }

	var subject *AltNativeResponse

	BeforeEach(func() {
		subject = new(AltNativeResponse)
	})

	It("should have defaults", func() {
		subject.SetAssets(&ResponseAsset{})
		subject.WithDefaults()

		Expect(*subject.Ver).To(Equal(1))
		Expect(*subject.Assets[0].Required).To(Equal(0))
	})

	It("should parse native adm", func() {
		resp, err := ParseAltNativeAdmBytes(testFixtures.simpleAltNativeResponse)

		Expect(err).NotTo(HaveOccurred())
		nativeAdm := AltNativeAdm{
			&AltNativeResponse{
				Ver: iptr(1),
				Assets: []ResponseAsset{
					ResponseAsset{
						Id:       iptr(1),
						Required: iptr(0),
						Title:    &testFixtures.simpleTitle,
						Link:     &testFixtures.simpleLink,
					},
					ResponseAsset{
						Id:       iptr(2),
						Required: iptr(0),
						Data:     &testFixtures.simpleData,
					},
					ResponseAsset{
						Id:       iptr(3),
						Required: iptr(0),
						Img:      &testFixtures.simpleImg,
					},
					ResponseAsset{
						Id:       iptr(4),
						Required: iptr(0),
						Data:     &testFixtures.installData,
						Link:     &testFixtures.simpleLink,
					},
				},
				Link:        &testFixtures.fullLink,
				Imptrackers: []string{"http: //a.com/a", "http: //b.com/b"},
			},
		}

		Expect(resp).To(Equal(&nativeAdm))
	})
})
