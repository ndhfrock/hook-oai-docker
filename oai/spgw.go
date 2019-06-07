package oai

import (
	"log"
	"snap-hook-for-docker/util"
)

// StartSpgw : Start SPGW as a daemon
func StartSpgw(logger *log.Logger) {
	sedCommand := "101s/8.8.8.8/8.8.8.8/g"
	// Init spgw
	logger.Print("Init spgw")
	util.RunCmd(logger, "/snap/bin/oai-cn.spgw-init")
	// Configure oai-spgw
	logger.Print("Configure spgw.conf")
	util.RunCmd(logger, "sed", "-i", "30s/eth0/lo/g", "/var/snap/oai-cn/28/spgw.conf")
	util.RunCmd(logger, "sed", "-i", "31s/192.168.11.17/127.0.1.10/g", "/var/snap/oai-cn/28/spgw.conf")
	//Get the nameserver of current env
	nameserver, err := util.GetNameserver(logger)
	if err != nil {
		logger.Print("Failed to fetch nameserver, using default(8.8.8.8)")
	} else {
		sedCommand = "101s/8.8.8.8/" + nameserver + "/g"
	}
	util.RunCmd(logger, "sed", "-i", sedCommand, "/var/snap/oai-cn/28/spgw.conf")

	// oai-cn.spgw-start
	logger.Print("start spgw as daemon")
	util.RunCmd(logger, "/snap/bin/oai-cn.spgw-start")
}
