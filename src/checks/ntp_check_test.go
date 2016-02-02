package checks_test

import (
	. "checks"

	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("NTP check", func() {
	var (
		originalNtpServer string
		invalidNtpServer  = "example.com"
	)

	BeforeEach(func() {
		stdout, _, err := RunCommand("w32tm /query /source")
		originalNtpServer = strings.Split(stdout, ",")[0]
		Expect(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		_, _, err := RunCommand(fmt.Sprintf("w32tm /config /update /manualpeerlist:\"%s\"", originalNtpServer))
		Expect(err).ShouldNot(HaveOccurred())
		_, _, err = RunCommand("w32tm /resync")
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("returns success when ntp settings are valid", func() {
		Expect(NtpCheck()).To(Succeed())
	})

	It("returns an error when ntp settings are incorrect", func() {
		_, _, err := RunCommand(fmt.Sprintf("w32tm /config /update /syncfromflags:manual /manualpeerlist:\"%s\"", invalidNtpServer))
		Expect(err).ShouldNot(HaveOccurred())

		err = NtpCheck()
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("There was an error detecting ntp"))
	})
})
