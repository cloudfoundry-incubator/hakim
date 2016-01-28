package checks

import (
	"errors"
	"strings"
)

const CheckFairShareCpuSetting = `
(gwmi win32_terminalservicesetting -N "root\cimv2\terminalservices").enableDFSS
`

func FairShareCpuCheck() error {
	stdout, _, err := runCommand(CheckFairShareCpuSetting)
	if err != nil {
		return err
	} else if strings.TrimSpace(stdout) != "0" {
		return errors.New("Fair Share CPU Scheduling must be disabled")
	}
	return nil
}
