package checks

import (
	"errors"
	"strings"

	"github.com/mitchellh/go-ps"
)

func ProcessCheck(processes []ps.Process, requiredProcessNames []string) error {
	missingProcesses := []string{}

	for _, requiredProcessName := range requiredProcessNames {
		missing := true
		for _, process := range processes {
			if strings.Contains(strings.ToUpper(process.Executable()), strings.ToUpper(requiredProcessName)) {
				missing = false
				break
			}
		}

		if missing {
			missingProcesses = append(missingProcesses, requiredProcessName)
		}
	}

	if len(missingProcesses) > 0 {
		return errors.New("The following processes are not running: " + strings.Join(missingProcesses, ", "))
	}
	return nil
}
