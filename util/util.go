package util

import (
	"errors"
	"log"
	"strings"

	"github.com/go-cmd/cmd"
)

// RunCmd : Run external commands in sync. Return stdout[0].
func RunCmd(logger *log.Logger, cmdName string, args ...string) cmd.Status {
	installSnap := cmd.NewCmd(cmdName, args...)
	finalStatus := <-installSnap.Start() // block and wait
	// logger.Print(finalStatus.Cmd)
	logger.Print(finalStatus)
	return finalStatus
}

// GetNameserver : Get Nameserver of the current ENV
func GetNameserver(logger *log.Logger) (string, error) {
	retStatus := RunCmd(logger, "nslookup", "google.com")
	if retStatus.Exit != 0 {
		return "", errors.New("nslookup return non-zero")
	} else if len(retStatus.Stdout) < 2 {
		return "", errors.New("len(stdout) < 2")
	}
	response := strings.Fields(retStatus.Stdout[0])
	if len(response) < 2 {
		return "", errors.New("Error in parsing results")
	}
	return response[1], nil
}
