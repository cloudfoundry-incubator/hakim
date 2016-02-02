package checks

import (
	"errors"
	"strings"
)

const CheckNTPSetting = "w32tm /resync"
const failureMessage = "The computer did not resync"

func NtpCheck() error {
	stdout, _, err := RunCommand(CheckNTPSetting)
	if err != nil {
		return err
	} else if strings.Contains(stdout, failureMessage) {
		return errors.New(
			`
There was an error detecting ntp synchronization on your machine.
An accurate system clock is essential for internal Cloud Foundry metric reports.

Please configure your NTP settings, if not already done.
We recommend that your firewall have inbound and outbound rules set for UDP on port 123.
In addition, ensure that your 'DnsCache' service is running.
			`)
	} else {
		return nil
	}
}
