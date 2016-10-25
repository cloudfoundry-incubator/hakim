package main

import (
	"checks"
	"flag"
	"log"
	"os"
	"time"

	"github.com/mitchellh/go-ps"
)

func init() {
	log.SetOutput(os.Stdout)
}

var (
	gardenAddr        = flag.String("gardenAddr", "localhost:9241", "Garden host and port (typically localhost:9241)")
	loopDelay         = flag.Int("loop", -1, "Repeat checks with specified delay (in seconds) between loops")
	requiredProcesses = []string{"consul.exe", "containerizer.exe", "garden-windows.exe", "rep.exe", "metron.exe"}
	bbsConsulHost     = "bbs.service.cf.internal"
)

func main() {
	flag.Parse()
	if *loopDelay >= 0 {
		for {
			runHakim()
			time.Sleep(time.Duration(*loopDelay) * time.Second)
		}
	} else {
		runHakim()
	}
}

func runHakim() {
	log.Print("Start Hakim...\r\n")

	var errs []error

	processes, err := ps.Processes()
	if err == nil {
		errs = append(errs, checks.ProcessCheck(processes, requiredProcesses))
	}
	errs = append(errs, err)

	errs = append(errs, checks.ContainerCheck(*gardenAddr, processes))
	errs = append(errs, checks.ConsulDnsCheck(bbsConsulHost))
	errs = append(errs, checks.FairShareCpuCheck())
	errs = append(errs, checks.FirewallCheck())
	errs = append(errs, checks.NtpCheck())

	hasErr := false
	for _, err := range errs {
		if err != nil {
			log.Print(err)
			hasErr = true
		}
	}
	if hasErr {
		log.Fatal("Hakim Failed to verify all components\r\n")
	}

	log.Print("Finished Hakim...\r\n")
}
