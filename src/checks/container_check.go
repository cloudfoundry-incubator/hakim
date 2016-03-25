package checks

import (
	"errors"

	"github.com/cloudfoundry-incubator/garden"
	gclient "github.com/cloudfoundry-incubator/garden/client"
	gconnection "github.com/cloudfoundry-incubator/garden/client/connection"
)

const (
	localLogonFailure              = "Logon failure: the user has not been granted the requested logon type at this computer"
	operatorFriendlyFailureMessage = "Logon failure: the user created by Containerizer has not been granted the requested logon type at this computer. Local accounts require permissions to logon locally."
)

func ContainerCheck(gardenAddr string) error {
	client := gclient.New(gconnection.New("tcp", gardenAddr))
	container, err := client.Create(garden.ContainerSpec{})
	if container != nil {
		defer client.Destroy(container.Handle())
	}

	if err != nil {
		if err.Error() == localLogonFailure {
			return errors.New("Failed to create container\n" + operatorFriendlyFailureMessage)
		} else {
			return errors.New("Failed to create container\n" + err.Error())
		}
	}
	return nil
}
