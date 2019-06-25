package oai

import (
	"log"
	"snap-hook-for-docker/util"
)

// StartSpgw : Start SPGW as a daemon
func startSpgw(logger *log.Logger) {
	// Init spgw
	logger.Print("Init spgw")
	util.RunCmd(logger, "/snap/bin/oai-cn.spgw-init")
	// Configure oai-spgw
	logger.Print("Configure spgw.conf")
	util.RunCmd(logger, "sed", "-i", "30s/lo/eth0/g", "/var/snap/oai-cn/current/spgw.conf")
	eth0IP, _ := util.GetInterfaceIP(logger, "eth0")
	sedCommand := "31s:\".*;:\"" + eth0IP + "/24\";:g"
	util.RunCmd(logger, "sed", "-i", sedCommand, "/var/snap/oai-cn/current/spgw.conf")
	//Get the nameserver of current env
	// nameserver, err := util.GetNameserver(logger)
	// if err != nil {
	// 	logger.Print("Failed to fetch nameserver, using default(8.8.8.8)")
	// } else {
	// 	sedCommand := "101s:\".*;:\"" + "192.168.106.12" + "\";:g"
	// 	util.RunCmd(logger, "sed", "-i", sedCommand, "/var/snap/oai-cn/current/spgw.conf")
	// }
	sedCommand = "101s:\".*;:\"" + "192.168.106.12" + "\";:g"
	util.RunCmd(logger, "sed", "-i", sedCommand, "/var/snap/oai-cn/current/spgw.conf")

	// oai-cn.spgw-start
	logger.Print("start spgw as daemon")
	util.RunCmd(logger, "/snap/bin/oai-cn.spgw-start")
}
