package checks_test

import (
	. "checks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NTP Checks", func() {
	It("succeeds when ntp work can resync", func() {
		Expect(NtpCheck()).To(Succeed())
	})

	It("fails when NTP service is off", func() {
		_, _, err := RunCommand("Set-Service W32Time -Status Stopped")
		Expect(err).ToNot(HaveOccurred())
		defer RunCommand("Set-Service W32Time -Status Running")

		err = NtpCheck()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("The service has not been started"))
	})
})
