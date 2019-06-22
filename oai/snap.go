package oai

import (
	"log"
	"os"
	"snap-hook-for-docker/util"
	"strings"
	"time"
)

// InstallSnapCore : Install Core
func installSnapCore(logger *log.Logger) {
	//Install Core
	logger.Print("Installing core")
	ret, err := util.CheckSnapPackageExist(logger, "core")
	if err != nil {
		logger.Print(err)
	}
	//Loop until package is installed
	if !ret {
		retStatus := util.RunCmd(logger, "snap", "install", "core", "--channel=edge")
		Snapfail := false
		for {
			if len(retStatus.Stderr[0]) > 0 {
				Snapfail = strings.Contains(retStatus.Stderr[0], "error")
			}
			if Snapfail {
				logger.Print("Wait for snapd being ready")
				time.Sleep(1 * time.Second)
				retStatus = util.RunCmd(logger, "snap", "install", "core", "--channel=edge")
			} else {
				logger.Print("snapd is ready and core is installed")
				break
			}
		}
	}

	// Install hello-world
	logger.Print("Installing hello-world")
	ret, err = util.CheckSnapPackageExist(logger, "hello-world")
	if err != nil {
		logger.Print(err)
	}
	if !ret {
		util.RunCmd(logger, "snap", "install", "hello-world")
	}

}

// InstallOaicn : Install oai-cn snap
func installOaicn(logger *log.Logger) {
	logger.Print("Configure hostname before installing ")
	// Copy hosts
	util.RunCmd(logger, "cp", "/etc/hosts", "./hosts_new")
	hostname, _ := os.Hostname()
	fullDomainName := "1s/^/127.0.0.1 " + hostname + ".openair4G.eur " + hostname + " hss\\n127.0.0.1 " + hostname + ".openair4G.eur " + hostname + " mme \\n/"
	util.RunCmd(logger, "sed", "-i", fullDomainName, "./hosts_new")
	// Replace hosts
	util.RunCmd(logger, "cp", "-f", "./hosts_new", "/etc/hosts")
	// Install oai-cn snap
	logger.Print("Installing oai-cn")
	ret, err := util.CheckSnapPackageExist(logger, "oai-cn")
	if err != nil {
		logger.Print(err)
	}
	if !ret {
		util.RunCmd(logger, "snap", "install", "oai-cn", "--channel=edge", "--devmode")
	}

}

// InstallOairan : Install oai-ran snap
func installOairan(logger *log.Logger) {
	// Install oai-ran snap
	logger.Print("Installing oai-ran")
	ret, err := util.CheckSnapPackageExist(logger, "oai-ran")
	if err != nil {
		logger.Print(err)
	}
	if !ret {
		util.RunCmd(logger, "snap", "install", "oai-ran", "--channel=edge", "--devmode")
	}

}
