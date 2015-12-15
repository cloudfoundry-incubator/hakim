package checks

import (
	"errors"
	"net"
)

func ConsulDnsCheck(checkHost string) error {
	_, err := net.LookupHost(checkHost)
	if err != nil {
		return errors.New("Failed to resolve consul host " + checkHost + "\n" + err.Error())
	}
	return err
}
