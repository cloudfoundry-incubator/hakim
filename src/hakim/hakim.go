package main

import (
	"checks"
	"flag"
	"log"

	"github.com/mitchellh/go-ps"
)

var (
	gardenAddr        = flag.String("gardenAddr", "localhost:9241", "Garden host and port (typically localhost:9241)")
	requiredProcesses = []string{"consul.exe", "containerizer.exe", "garden-windows.exe", "rep.exe", "metron.exe"}
	bbsConsulHost     = "bbs.service.cf.internal"
)

func main() {
	flag.Parse()

	processes, err := ps.Processes()
	if err != nil {
		panic(err)
	}

	err = checks.ProcessCheck(processes, requiredProcesses)
	if err != nil {
		log.Fatal(err)
	}

	err = checks.ContainerCheck(*gardenAddr)
	if err != nil {
		log.Fatal(err)
	}

	err = checks.ConsulDnsCheck(bbsConsulHost)
	if err != nil {
		log.Fatal(err)
	}

	err = checks.FairShareCpuCheck()
	if err != nil {
		log.Fatal(err)
	}

	err = checks.FirewallCheck()
	if err != nil {
		log.Fatal(err)
	}

	err = checks.NtpCheck()
	if err != nil {
		log.Fatal(err)
	}
}
