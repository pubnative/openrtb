package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Video", func() {
	var subject *Video

	BeforeEach(func() { subject = new(Video) })

	Describe("Seq()", func() {
		It("should return '1' as default", func() {
			subject.WithDefaults()
			Expect(subject.Seq()).To(Equal(1))
		})
	})

	Describe("IsBoxingAllowed()", func() {
		It("should return true as default", func() {
			subject.WithDefaults()
			Expect(subject.IsBoxingAllowed()).To(BeTrue())
		})

		It("should return false when set as false", func() {
			Expect(subject.IsBoxingAllowed()).To(BeFalse())
		})
	})

	Describe("Valid()", func() {
		It("should return error messages when attributes missing", func() {
			ok, err := subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: video has no mimes"))

			subject.Mimes = []string{"RAND_KEY"} // With Mimes
			ok, err = subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: video linearity missing"))

			subject.Linearity = 2 // With Linearity
			ok, err = subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: video minduration missing"))

			subject.Minduration = 1 // With Minduration
			ok, err = subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: video maxduration missing"))

			subject.Maxduration = 5 // With Maxduration
			ok, err = subject.Valid()
			Expect(err.Error()).To(Equal("openrtb parse: video protocol missing"))

			subject.Protocol = 1 // With Protocol
			ok, err = subject.Valid()
			Expect(err).NotTo(HaveOccurred())
			Expect(ok).To(BeTrue())
		})
	})

	Describe("WithDefaults()", func() {
		It("should return object with default values", func() {
			subject.WithDefaults()
			Expect(subject.Sequence).To(Equal(1))
			Expect(subject.Boxingallowed).To(Equal(1))
		})
	})
})
