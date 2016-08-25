package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AltNativeResponse", func() {
	var subject AltNativeResponse

	BeforeEach(func() {
		subject = AltNativeResponse{}
	})

	It("should have defaults", func() {
		subject.WithDefaults()
		Expect(subject.Ver).To(Equal(1))
	})

	It("should parse native adm", func() {
		resp, err := ParseAltNativeAdmBytes(testFixtures.simpleAltNativeResponse)

		Expect(err).NotTo(HaveOccurred())
		nativeAdm := AltNativeAdm{
			Native: &AltNativeResponse{
				Ver: 1,
				Assets: []ResponseAsset{
					{
						Id:       1,
						Required: 0,
						Title:    &testFixtures.simpleTitle,
						Link:     &testFixtures.simpleLink,
					},
					{
						Id:       2,
						Required: 0,
						Data:     &testFixtures.simpleData,
					},
					{
						Id:       3,
						Required: 0,
						Img:      &testFixtures.simpleImg,
					},
					{
						Id:       4,
						Required: 0,
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

	It("returns an error when response doesn't contain native node", func() {
		_, err := ParseAltNativeAdmBytes(testFixtures.nullLiteralResponse)
		Expect(err).To(Equal(ErrBlankNativeResponse))
	})
})
