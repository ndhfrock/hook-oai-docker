package oai

import (
	"log"
	"os"
	"snap-hook-for-docker/util"
)

// StartMme : Start MME as a daemon
func startMme(logger *log.Logger) {
	// Init mme
	logger.Print("Init mme")
	util.RunCmd(logger, "/snap/bin/oai-cn.mme-init")
	hostname, _ := os.Hostname()
	// Configure oai-mme
	logger.Print("Configure mme.conf")
	sedCommand := "56s/ubuntu/" + hostname + "/g"
	util.RunCmd(logger, "sed", "-i", sedCommand, "/var/snap/oai-cn/current/mme.conf")
	// Configure interface name
	util.RunCmd(logger, "sed", "-i", "153s/lo/eth0/g", "/var/snap/oai-cn/current/mme.conf")
	// Get eth0 ip and replace the default one
	eth0IP, _ := util.GetInterfaceIP(logger, "eth0")
	sedCommand = "154s:\".*;:\"" + eth0IP + "/24\";:g"
	util.RunCmd(logger, "sed", "-i", sedCommand, "/var/snap/oai-cn/current/mme.conf")
	// Replace MNC
	util.RunCmd(logger, "sed", "-i", "78s/93/95/g", "/var/snap/oai-cn/current/mme.conf")
	util.RunCmd(logger, "sed", "-i", "87s/93/95/g", "/var/snap/oai-cn/current/mme.conf")
	// Replace Identity
	logger.Print("Configure mme_fd.conf")
	sedCommand = "4s/ubuntu/" + hostname + "/g"
	util.RunCmd(logger, "sed", "-i", sedCommand, "/var/snap/oai-cn/current/mme_fd.conf")
	sedCommand = "103s/ubuntu/" + hostname + "/g"
	util.RunCmd(logger, "sed", "-i", sedCommand, "/var/snap/oai-cn/current/mme_fd.conf")
	// oai-cn.mme-start
	logger.Print("start mme as daemon")
	util.RunCmd(logger, "/snap/bin/oai-cn.mme-start")
}
