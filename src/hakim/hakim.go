package main

import (
	"checks"
	"flag"
	"log"

	"github.com/mitchellh/go-ps"
)

var requiredProcesses = []string{"consul.exe", "containerizer.exe", "garden-windows.exe", "rep.exe", "metron.exe"}

var bbsCosulHost = "bbs.service.cf.internal"

func main() {
	gardenAddr := flag.String("gardenAddr", "localhost:9241", "Garden host and port (typically localhost:9241)")
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

	err = checks.ConsulDnsCheck(bbsCosulHost)
	if err != nil {
		log.Fatal(err)
	}
}
