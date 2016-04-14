package checks

import (
	"errors"
	"strings"

	"github.com/cloudfoundry-incubator/garden"
	gclient "github.com/cloudfoundry-incubator/garden/client"
	gconnection "github.com/cloudfoundry-incubator/garden/client/connection"
	"github.com/mitchellh/go-ps"
)

const (
	logonFailure      = "Logon failure: the user has not been granted the requested logon type at this computer"
	localLogonMessage = "Logon failure: the user created by Containerizer has not been granted the requested logon type at this computer. Local accounts require permissions to logon locally."
	batchLogonMessage = "Logon failure: the user created by Containerizer has not been granted the requested logon type at this computer. Local accounts require permissions to logon as a batch user."
)

func ContainerCheck(gardenAddr string, processes []ps.Process) error {
	var errMsg string
	stdout, _, err := RunCommand(`
$proc = Get-CimInstance Win32_Process -Filter "name = 'containerizer.exe'"
$result = Invoke-CimMethod -InputObject $proc -MethodName GetOwner
$result.User
`)
	if err != nil {
		return err
	}
	if strings.HasPrefix(stdout, "SYSTEM") {
		errMsg = batchLogonMessage
	} else {
		errMsg = localLogonMessage
	}

	client := gclient.New(gconnection.New("tcp", gardenAddr))
	container, err := client.Create(garden.ContainerSpec{})
	if container != nil {
		defer client.Destroy(container.Handle())
	}

	if err != nil {
		if err.Error() == logonFailure {
			return errors.New("Failed to create container\n" + errMsg)
		} else {
			return errors.New("Failed to create container\n" + err.Error())
		}
	}
	return nil
}
