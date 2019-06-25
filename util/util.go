package util

import (
	"errors"
	"log"
	"net"
	"strings"

	"github.com/go-cmd/cmd"
)

// RunCmd will run external commands in sync. Return stdout[0].
func RunCmd(logger *log.Logger, cmdName string, args ...string) cmd.Status {
	installSnap := cmd.NewCmd(cmdName, args...)
	finalStatus := <-installSnap.Start() // block and wait
	// logger.Print(finalStatus.Cmd)
	logger.Print(finalStatus)
	return finalStatus
}

// GetNameserver will get Nameserver of the current ENV
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

//CheckSnapPackageExist will return if this package is already exist or not
func CheckSnapPackageExist(logger *log.Logger, packageName string) (bool, error) {
	if len(packageName) <= 0 {
		return false, errors.New("Input package name is empty")
	}
	retStatus := RunCmd(logger, "snap", "list")
	if retStatus.Exit != 0 {
		return false, errors.New("snap list return non-zero")
	}
	for i := 0; i < len(retStatus.Stdout); i++ {
		if strings.Contains(retStatus.Stdout[i], packageName) {
			logger.Println("Package: ", packageName, " Exist")
			return true, nil
		}

	}
	logger.Println("Package: ", packageName, " does not Exist")
	return false, nil
}

//GetInterfaceIP will get the ip of the interface. If failed, it'll return a default (127.0.1.10) value
func GetInterfaceIP(logger *log.Logger, interfaceName string) (string, error) {
	ret := RunCmd(logger, "ifconfig", interfaceName)
	if ret.Exit != 0 {
		return "127.0.1.10", errors.New("Fail to run ifconfig")
	}
	if len(ret.Stdout) <= 0 {
		return "127.0.1.10", errors.New("Fail to get result")
	}
	i := 0
	space := " "
	for {
		if ret.Stdout[1][27+i+1] == space[0] {
			break
		}
		i++
	}
	return ret.Stdout[1][20 : 27+i+1], nil
}

//GetIPFromDomain will get the IP of the domain
func GetIPFromDomain(logger *log.Logger, domain string) (string, error) {
	addr, err := net.LookupHost(domain)
	if err != nil {
		logger.Print("Failed to get IP from domain")
		return "", err
	}
	return addr[0], nil
}
