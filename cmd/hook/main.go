package main

// This APP is made for installing snaps in docker
import (
	"flag"
	"log"
	"os"
	"snap-hook-for-docker/oai"
)

func main() {
	// Initialize Log
	logFile, err := os.Create("/root/hook.log")
	if err != nil {
		log.Fatal("Cannot open file", err)
		os.Exit(-1)
	}
	defer logFile.Close()
	logger := log.New(logFile, "[Debug] ", log.Lshortfile)

	// Parse input flags
	cn := flag.Bool("cn", false, "a bool")
	ran := flag.Bool("ran", false, "a bool")
	snapCNOnly := flag.Bool("snapCNOnly", false, "a bool")
	flag.Parse()

	// Decide actions based on flags
	if *cn {
		initCN(logger)
	} else if *ran {
		initRAN(logger)
	} else if *snapCNOnly {
		initSnapCNOnly(logger)
	}

	// Give a hello when program ends
	logger.Print("End of hook")
}

func initSnapCNOnly(logger *log.Logger) {
	// Install Snap Core
	oai.InstallSnap(logger)
	// Install oai-cn snap
	oai.InstallOaicn(logger)
}

func initCN(logger *log.Logger) {
	// Install Snap Core
	oai.InstallSnap(logger)
	// Install oai-cn snap
	oai.InstallOaicn(logger)
	// Start HSS
	oai.StartHss(logger)
	// Start MME
	oai.StartMme(logger)
	// Start SPGW
	oai.StartSpgw(logger)
}

func initRAN(logger *log.Logger) {
	// Install Snap Core
	oai.InstallSnap(logger)
	// Install oai-ran snap
	oai.InstallOairan(logger)
}
