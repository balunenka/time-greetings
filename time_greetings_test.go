package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Greetings method", func() {
	It("should return greeting", func() {
		Expect(GetGreetings(1)).To(Equal("Good night!"))

	})
})
