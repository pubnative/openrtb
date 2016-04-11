package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Impression", func() {
	var subject *Impression

	BeforeEach(func() {
		subject = new(Impression)
	})

	It("should validate", func() {
		b := &Banner{}
		v := &Video{Mimes: []string{"MIME_123"}}
		v.Linearity = 1
		v.Minduration = 1
		v.Maxduration = 5
		v.Protocol = 1

		ok, err := subject.Valid()
		Expect(err).To(HaveOccurred())
		Expect(ok).To(BeFalse())

		subject.Id = "CODE_12"
		subject.Banner = b
		ok, err = subject.Valid()
		Expect(err).NotTo(HaveOccurred())
		Expect(ok).To(BeTrue())

		subject.Video = v
		ok, err = subject.Valid()
		Expect(err).To(HaveOccurred())
		Expect(ok).To(BeFalse())

		subject.Banner = nil
		ok, err = subject.Valid()
		Expect(err).NotTo(HaveOccurred())
		Expect(ok).To(BeTrue())
	})

	It("should have accessors", func() {
		Expect(subject.IsSecure()).To(BeFalse())
	})

	It("should have defaults", func() {
		subject.WithDefaults()
		Expect(subject.Bidfloorcur).To(Equal("USD"))
	})
})
