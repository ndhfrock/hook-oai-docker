package oai

import (
	"log"
	"os"
	"snap-hook-for-docker/util"
)

// StartHss : Start HSS as a daemon
func StartHss(logger *log.Logger) {
	hostname, _ := os.Hostname()
	// Configure oai-hss
	logger.Print("Configure hss.conf")
	util.RunCmd(logger, "sed", "-i", "-r", "31s/.{1}//", "/var/snap/oai-cn/28/hss.conf")
	util.RunCmd(logger, "sed", "-i", "30s/^/#/", "/var/snap/oai-cn/28/hss.conf")
	util.RunCmd(logger, "sed", "-i", "s/127.0.0.1/mysql/g", "/var/snap/oai-cn/28/hss.conf")
	//Replace Identity
	logger.Print("Configure hss_fd.conf")
	identity := hostname + ".openair4G.eur"
	syntax := "s/ubuntu.openair4G.eur/" + identity + "/g"
	util.RunCmd(logger, "sed", "-i", syntax, "/var/snap/oai-cn/28/hss_fd.conf")
	// Init hss
	logger.Print("Init hss")
	util.RunCmd(logger, "/snap/bin/oai-cn.hss-init")
	// oai-cn.hss-start
	logger.Print("start hss as daemon")
	util.RunCmd(logger, "/snap/bin/oai-cn.hss-start")
}
