package checks_test

import (
	. "checks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Consul DNS checks", func() {
	It("returns an error when host is not known", func() {
		err := ConsulDnsCheck("non-existent.example.com.")
		Expect(err.Error()).To(ContainSubstring("Failed to resolve consul host non-existent.example.com."))
	})

	It("does not return an error when host is known", func() {
		Expect(ConsulDnsCheck("localhost")).To(Succeed())
	})
})
