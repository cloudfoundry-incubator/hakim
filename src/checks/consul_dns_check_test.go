package checks_test

import (
	. "checks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Consul DNS checks", func() {
	It("returns an error when host is not known", func() {
		err := ConsulDnsCheck("host-non-existing.")
		Expect(err).To(MatchError("Failed to resolve consul host host-non-existing.\nlookup host-non-existing.: getaddrinfow: No such host is known."))
	})

	It("does not return an error when host is known", func() {
		Expect(ConsulDnsCheck("localhost")).To(Succeed())
	})
})
