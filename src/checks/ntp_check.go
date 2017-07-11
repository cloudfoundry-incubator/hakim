package checks

import "errors"

func NtpCheck() error {
	stdout, _, err := RunCommand("w32tm /resync")
	if err != nil {
		return errors.New(`
There was an error detecting ntp synchronization on your machine.
An accurate system clock is essential for internal Cloud Foundry metric reports.

Please configure your NTP settings, if not already done.
We recommend that your firewall have outbound rules set for UDP on port 123.
In addition, ensure that your 'DnsCache' service is running.  ` + "Error: \n\n" + stdout)
	}

	return nil
}
