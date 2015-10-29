package checks

import (
	"errors"
	"strings"

	"github.com/mitchellh/go-ps"
)

func ProcessCheck(processes []ps.Process, requiredProcessNames []string) error {

	for _, process := range processes {
		processName := strings.ToUpper(process.Executable())

		for i, requiredProcessName := range requiredProcessNames {
			if strings.Contains(processName, strings.ToUpper(requiredProcessName)) {
				requiredProcessNames = append(requiredProcessNames[:i], requiredProcessNames[i+1:]...)
			}
		}
	}

	if len(requiredProcessNames) > 0 {
		return errors.New("The following processes are not running: " + strings.Join(requiredProcessNames, ", "))
	}
	return nil
}
