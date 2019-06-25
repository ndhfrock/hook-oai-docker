package main

// This APP is made for installing snaps in docker and
// handle the configurations
// Author: Hsi-Ping Hsu
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
	//Install snap core
	oai.InstallSnap(logger)
	// Decide actions based on flags
	if *cn {
		oai.InstallCN(logger)
		oai.StartCN(logger)
	} else if *ran {
		oai.InstallRAN(logger)
		oai.StartENB(logger)
	} else if *snapCNOnly {
		oai.InstallCN(logger)
	}

	// Give a hello when program ends
	logger.Print("End of hook")
}
