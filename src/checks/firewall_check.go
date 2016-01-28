package checks

import (
	"errors"
	"strings"
)

const CheckFirewallSetting = `
netsh advfirewall show currentprofile state | Select-String -Quiet ON
`

func FirewallCheck() error {
	stdout, _, err := runCommand(CheckFirewallSetting)
	if err != nil {
		return err
	} else if strings.TrimSpace(stdout) != "True" {
		return errors.New("Windows firewall service is not enabled. The Windows firewall is required in order to enforce Application Security Group rules. Running without the firewall is possible, but strongly not recommended.")
	}
	return nil
}
