package oai

import (
	"log"
	"os"
	"snap-hook-for-docker/util"
)

// InstallSnap : Install Core
func InstallSnap(logger *log.Logger) {
	//Install Core
	logger.Print("Installing core")
	util.RunCmd(logger, "snap", "install", "core", "--channel=edge")
	// Install hello-world
	logger.Print("Installing hello-world")
	util.RunCmd(logger, "snap", "install", "hello-world")
}

// InstallOaicn : Install oai-cn snap
func InstallOaicn(logger *log.Logger) {
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
	util.RunCmd(logger, "snap", "install", "oai-cn", "--channel=edge", "--devmode")
}

// InstallOairan : Install oai-ran snap
func InstallOairan(logger *log.Logger) {
	// Install oai-ran snap
	logger.Print("Installing oai-ran")
	util.RunCmd(logger, "snap", "install", "oai-ran", "--channel=edge", "--devmode")
}
