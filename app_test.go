package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	var subject *App

	BeforeEach(func() {
		subject = new(App)
	})

	It("should have accessors", func() {
		Expect(subject.IsPrivacyPolicy()).To(BeFalse())
		Expect(subject.IsPaid()).To(BeFalse())
	})
})
