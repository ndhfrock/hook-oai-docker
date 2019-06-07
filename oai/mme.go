package oai

import (
	"log"
	"os"
	"snap-hook-for-docker/util"
)

// StartMme : Start MME as a daemon
func StartMme(logger *log.Logger) {
	// Init mme
	logger.Print("Init mme")
	util.RunCmd(logger, "/snap/bin/oai-cn.mme-init")
	hostname, _ := os.Hostname()
	// Configure oai-mme
	logger.Print("Configure mme.conf")
	sedCommand := "56s/ubuntu/" + hostname + "/g"
	util.RunCmd(logger, "sed", "-i", sedCommand, "/var/snap/oai-cn/28/mme.conf")
	util.RunCmd(logger, "sed", "-i", "153s/eth0/lo/g", "/var/snap/oai-cn/28/mme.conf")
	util.RunCmd(logger, "sed", "-i", "154s/192.168.11.17/127.0.1.10/g", "/var/snap/oai-cn/28/mme.conf")
	//Replace Identity
	logger.Print("Configure mme_fd.conf")
	sedCommand = "4s/ubuntu/" + hostname + "/g"
	util.RunCmd(logger, "sed", "-i", sedCommand, "/var/snap/oai-cn/28/mme_fd.conf")
	sedCommand = "103s/ubuntu/" + hostname + "/g"
	util.RunCmd(logger, "sed", "-i", sedCommand, "/var/snap/oai-cn/28/mme_fd.conf")
	// oai-cn.mme-start
	logger.Print("start mme as daemon")
	util.RunCmd(logger, "/snap/bin/oai-cn.mme-start")
}
