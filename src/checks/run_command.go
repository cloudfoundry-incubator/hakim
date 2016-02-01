package checks

import (
	"bytes"
	"os/exec"
)

func runCommand(cmd string) (string, string, error) {
	execCmd := exec.Command("powershell", "-noprofile", "-noninteractive", "-command", cmd)

	var stdout bytes.Buffer
	execCmd.Stdout = &stdout

	var stderr bytes.Buffer
	execCmd.Stderr = &stderr

	err := execCmd.Run()
	return stdout.String(), stderr.String(), err
}
