package oai

import (
	"oai-snap-in-docker/internal/pkg/util"
)

// StartSpgw : Start SPGW as a daemon
func startSpgw(OaiObj Oai) {
	spgwConf := OaiObj.Conf.ConfigurationPathofCN + "spgw.conf"
	spgwBin := OaiObj.Conf.SnapBinaryPath + "oai-cn.spgw"
	// Init spgw
	OaiObj.Logger.Print("Init spgw")
	util.RunCmd(OaiObj.Logger, spgwBin+"-init")
	// Configure oai-spgw
	OaiObj.Logger.Print("Configure spgw.conf")
	//Set up interface
	util.RunCmd(OaiObj.Logger, "sed", "-i", "30s/lo/eth0/g", spgwConf)
	//Get interface IP and configure the spgw.conf
	eth0IP, _ := util.GetInterfaceIP(OaiObj.Logger, "eth0")
	sedCommand := "31s:\".*;:\"" + eth0IP + "/24\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, spgwConf)
	//Get the nameserver from conf
	sedCommand = "101s:\".*;:\"" + OaiObj.Conf.DNS + "\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, spgwConf)
	// oai-cn.spgw-start
	OaiObj.Logger.Print("start spgw as daemon")
	util.RunCmd(OaiObj.Logger, spgwBin+"-start")
}
