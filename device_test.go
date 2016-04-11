package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Device", func() {
	var subject *Device

	BeforeEach(func() {
		subject = new(Device)
	})

	It("should have accessors", func() {
		Expect(subject.IsDnt()).To(BeFalse())
		Expect(subject.IsJs()).To(BeFalse())
	})
})
