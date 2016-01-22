package checks

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func runCommand(cmd string) (string, string, error) {
	file, err := ioutil.TempFile(os.TempDir(), "hakim")
	if err != nil {
		return "", "", err
	}
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(cmd))
	if err != nil {
		return "", "", err
	}

	err = file.Close()
	if err != nil {
		return "", "", err
	}

	err = os.Rename(file.Name(), file.Name()+".ps1")
	if err != nil {
		return "", "", err
	}

	execCmd := exec.Command("powershell", "-noprofile", "-noninteractive", file.Name()+".ps1")

	execCmd.Stdin = bytes.NewBufferString(cmd)

	stdout := bytes.NewBufferString("")
	execCmd.Stdout = stdout

	stderr := bytes.NewBufferString("")
	execCmd.Stderr = stderr

	err = execCmd.Run()

	stdoutStr := string(stdout.Bytes())
	stderrStr := string(stderr.Bytes())

	execCmd.Wait()

	return stdoutStr, stderrStr, err
}

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
