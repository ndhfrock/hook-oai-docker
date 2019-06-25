package oai

import (
	"log"
	"os"
	"snap-hook-for-docker/util"
	"strings"
)

// StartHss : Start HSS as a daemon
func startHss(logger *log.Logger) {
	hostname, _ := os.Hostname()
	// Configure oai-hss
	logger.Print("Configure hss.conf")
	//util.RunCmd(logger, "sed", "-i", "-r", "31s/.{1}//", "/var/snap/oai-cn/current/hss.conf")
	//util.RunCmd(logger, "sed", "-i", "30s/^/#/", "/var/snap/oai-cn/current/hss.conf")
	util.RunCmd(logger, "sed", "-i", "s/127.0.0.1/mysql/g", "/var/snap/oai-cn/current/hss.conf")
	//Replace Identity
	logger.Print("Configure hss_fd.conf")
	identity := hostname + ".openair4G.eur"
	syntax := "s/ubuntu.openair4G.eur/" + identity + "/g"
	util.RunCmd(logger, "sed", "-i", syntax, "/var/snap/oai-cn/current/hss_fd.conf")
	// Init hss
	logger.Print("Init hss")
	retStatus := util.RunCmd(logger, "/snap/bin/oai-cn.hss-init")
	for {
		fail := false
		for i := 0; i < len(retStatus.Stderr); i++ {
			if strings.Contains(retStatus.Stderr[i], "ERROR") {
				logger.Println("Init error, re-run again")
				fail = true
			}
		}
		if fail {
			retStatus = util.RunCmd(logger, "/snap/bin/oai-cn.hss-init")
		} else {
			break
		}
	}
	// oai-cn.hss-start
	logger.Print("start hss as daemon")
	util.RunCmd(logger, "/snap/bin/oai-cn.hss-start")
}
