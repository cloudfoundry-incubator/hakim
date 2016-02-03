package checks

import (
	"errors"
	"strings"
)

func NtpCheck() error {
	stdout, _, _ := RunCommand("w32tm /resync")
	if strings.Contains(stdout, "The command completed successfully") {
		return nil
	} else {
		return errors.New(`
There was an error detecting ntp synchronization on your machine.
An accurate system clock is essential for internal Cloud Foundry metric reports.

Please configure your NTP settings, if not already done.
We recommend that your firewall have inbound and outbound rules set for UDP on port 123.
In addition, ensure that your 'DnsCache' service is running.  ` + "Error: \n\n" + stdout)
	}
}
