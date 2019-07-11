package main

// This APP is made for installing snaps in docker and
// handle the configurations
// Author: Kevin Hsi-Ping Hsu
import (
	"flag"
	"fmt"
	"oai-snap-in-docker/internal/oai"
)

const (
	logPath  = "/root/hook.log"
	confPath = "/root/config/conf.yaml"
)

func main() {
	// Initialize oai struct
	OaiObj := oai.Oai{}
	err := OaiObj.Init(logPath, confPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse input flags
	runCN := flag.Bool("runCN", false, "a bool")
	runENB := flag.Bool("runENB", false, "a bool")
	installCN := flag.Bool("installCN", false, "a bool")
	installRAN := flag.Bool("installRAN", false, "a bool")
	installFlexRAN := flag.Bool("installFlexRAN", false, "a bool")
	installMEC := flag.Bool("installMEC", false, "a bool")
	flag.Parse()
	//Install snap core
	oai.InstallSnap(OaiObj)
	// Decide actions based on flags
	if *installCN {
		oai.InstallCN(OaiObj)
		oai.StartCN(OaiObj)
	} else if *runCN {
		oai.StartCN(OaiObj)
	} else if *runENB {
		oai.StartENB(OaiObj)
	} else if *installRAN {
		oai.InstallRAN(OaiObj)
		oai.StartENB(OaiObj)
	} else if *installFlexRAN {
		oai.InstallFlexRAN(OaiObj)
		oai.StartFlexRAN(OaiObj)
	} else if *installMEC {
		oai.InstallMEC(OaiObj)
	} else {
		fmt.Println("This should only be executed in container!!")
		return
	}

	// Give a hello when program ends
	OaiObj.Logger.Print("End of hook")
	OaiObj.Clean()
}
