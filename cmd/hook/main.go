package main

// This APP is made for installing snaps and Mosaic 5G in docker
// This APP also handle all the configurations inside the docker
// Author: Kevin Hsi-Ping Hsu & Nadhif Muhammad Rekoputra
import (
	"flag"
	"fmt"

	"github.com/hook-oai-docker/internal/oai"
)

const (
	logPath = "/root/hook.log"
	//logPath = "/home/nadhif/mosaic5g/store/docker-oai-snap/dockers/build/hook.log"
	confPath = "/root/config/conf.yaml"
	//confPath = "/home/nadhif/mosaic5g/store/docker-oai-snap/dockers/build/conf.yaml"
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
	installCN := flag.Bool("installCN", false, "a bool")
	installRAN := flag.Bool("installRAN", false, "a bool")
	installRANSlicing := flag.Bool("installRANSlicing", false, "a bool")
	installHSS := flag.Bool("installHSS", false, "a bool")
	installMME := flag.Bool("installMME", false, "a bool")
	installSPGW := flag.Bool("installSPGW", false, "a bool")
	installFlexRAN := flag.Bool("installFlexRAN", false, "a bool")
	installMEC := flag.Bool("installMEC", false, "a bool")
	installDroneStore := flag.Bool("installDroneStore", false, "a bool")
	installRRMKPIStore := flag.Bool("installRRMKPIStore", false, "a bool")
	flag.Parse()
	// Decide actions based on flags
	if *installCN {
		oai.InstallSnap(OaiObj)
		oai.InstallCN(OaiObj)
		oai.StartCN(OaiObj)
	} else if *installRAN {
		oai.InstallSnap(OaiObj)
		oai.InstallRAN(OaiObj)
		oai.StartENB(OaiObj)
	} else if *installRANSlicing {
		oai.InstallRANSlicing(OaiObj)
		oai.InstallSnap(OaiObj)
		oai.StartENBSlicing(OaiObj)
	} else if *installHSS {
		oai.InstallSnap(OaiObj)
		oai.InstallCN(OaiObj)
		oai.StartHSS(OaiObj)
	} else if *installMME {
		oai.InstallSnap(OaiObj)
		oai.InstallCN(OaiObj)
		oai.StartMME(OaiObj)
	} else if *installSPGW {
		oai.InstallSnap(OaiObj)
		oai.InstallCN(OaiObj)
		oai.StartSPGW(OaiObj)
	} else if *installFlexRAN {
		oai.InstallSnap(OaiObj)
		oai.InstallFlexRAN(OaiObj)
		oai.StartFlexRAN(OaiObj)
	} else if *installMEC {
		oai.InstallSnap(OaiObj)
		oai.InstallMEC(OaiObj)
	} else if *installDroneStore {
		oai.InstallStore(OaiObj)
		oai.InstallSnap(OaiObj)
		oai.StartDrone(OaiObj)
	} else if *installRRMKPIStore {
		oai.InstallStore(OaiObj)
		oai.InstallSnap(OaiObj)
		oai.StartRRMKPI(OaiObj)
	} else {
		fmt.Println("This should only be executed in container!!")
		return
	}

	// Give a hello when program ends
	OaiObj.Logger.Print("End of hook")
	OaiObj.Clean()
}
