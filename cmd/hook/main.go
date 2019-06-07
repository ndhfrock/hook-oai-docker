package main

// This APP is made for installing snaps in docker
import (
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
	initAllInOne(logger)
	logger.Print("End of hook")

}

func initAllInOne(logger *log.Logger) {
	// Install oai-cn snap
	oai.InstallSnap(logger)
	// Start HSS
	oai.StartHss(logger)
	// Start MME
	oai.StartMme(logger)
	// Start SPGW
	oai.StartSpgw(logger)
}
