package oai

import (
	"errors"
	"oai-snap-in-docker/internal/pkg/util"
	"os"
)

// StartMme : Start MME as a daemon
func startMme(OaiObj Oai) error {
	c := OaiObj.Conf
	mmeConf := c.ConfigurationPathofCN + "mme.conf"
	mmeFdConf := c.ConfigurationPathofCN + "mme_fd.conf"
	mmeBin := c.SnapBinaryPath + "oai-cn.mme"
	// Init mme
	OaiObj.Logger.Print("Init mme")
	retStatus := util.RunCmd(OaiObj.Logger, mmeBin+"-init")
	if retStatus.Exit != 0 {
		return errors.New("mme init failed ")
	}
	hostname, _ := os.Hostname()
	// Configure oai-mme
	OaiObj.Logger.Print("Configure mme.conf")
	sedCommand := "56s/ubuntu/" + hostname + "/g"
	retStatus = util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, mmeConf)
	if retStatus.Exit != 0 {
		return errors.New("Set mme domain name in " + mmeConf + " failed")
	}
	// Configure interface name
	retStatus = util.RunCmd(OaiObj.Logger, "sed", "-i", "153s/lo/eth0/g", mmeConf)
	if retStatus.Exit != 0 {
		return errors.New("Set interface name in " + mmeConf + " failed")
	}
	// Get eth0 ip and replace the default one
	eth0IP, _ := util.GetInterfaceIP(OaiObj.Logger, "eth0")
	sedCommand = "154s:\".*;:\"" + eth0IP + "/24\";:g"
	retStatus = util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, mmeConf)
	if retStatus.Exit != 0 {
		return errors.New("Set interface IP in " + mmeConf + " failed")
	}
	// Replace GUMMEI
	OaiObj.Logger.Print("Replace MNC")
	sedCommand = "s:\"MNC=\\\"93\\\"\":\"MNC=\\\"" + c.MNC + "\\\" \":g"
	retStatus = util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, mmeConf)
	if retStatus.Exit != 0 {
		return errors.New("Set GUMMEI in " + mmeConf + " failed")
	}
	OaiObj.Logger.Print("Replace MCC")
	//Replace MCC
	sedCommand = "s:\"{MCC=\\\"208\\\"\":\"{MCC=\\\"" + c.MCC + "\\\" \":g"
	retStatus = util.RunCmd(OaiObj.Logger, "sed", "-i", "87s/93/95/g", mmeConf)
	if retStatus.Exit != 0 {
		return errors.New("Set TAI in " + mmeConf + " failed")
	}
	// Replace Identity
	OaiObj.Logger.Print("Configure mme_fd.conf")
	sedCommand = "4s/ubuntu/" + hostname + "/g"
	retStatus = util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, mmeFdConf)
	if retStatus.Exit != 0 {
		return errors.New("Set Identity in " + mmeFdConf + " failed")
	}
	sedCommand = "103s/ubuntu/" + hostname + "/g"
	retStatus = util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, mmeFdConf)
	if retStatus.Exit != 0 {
		return errors.New("Set hostname in " + mmeFdConf + " failed")
	}
	// oai-cn.mme-start
	OaiObj.Logger.Print("start mme as daemon")
	util.RunCmd(OaiObj.Logger, mmeBin+"-start")
	return nil
}
