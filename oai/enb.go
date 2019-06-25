package oai

import (
	"fmt"
	"log"
	"snap-hook-for-docker/util"
	"time"
)

const (
	confPath = "/var/snap/oai-ran/current/enb.band7.tm1.50PRB.usrpb210.conf"
	oaicn    = "oaicn"
)

func startENB(logger *log.Logger) {
	// Replace MNC
	util.RunCmd(logger, "sed", "-i", "18s/93/95/g", confPath)
	// Replace oaicn ip
	mme, err := util.GetIPFromDomain(logger, "oaicn")
	if err != nil {
		fmt.Print(err)
		mme = "oaicn"
	}
	sedCommand := "175s:\".*;:\"" + mme + "\";:g"
	util.RunCmd(logger, "sed", "-i", sedCommand, confPath)
	// Replace interface
	util.RunCmd(logger, "sed", "-i", "191s/eno1/eth0/g", confPath)
	util.RunCmd(logger, "sed", "-i", "193s/eno1/eth0/g", confPath)
	// Replace enb IP
	eth0IP, _ := util.GetInterfaceIP(logger, "eth0")
	sedCommand = "192s:\".*;:\"" + eth0IP + "/23\";:g"
	util.RunCmd(logger, "sed", "-i", sedCommand, confPath)
	sedCommand = "194s:\".*;:\"" + eth0IP + "/23\";:g"
	util.RunCmd(logger, "sed", "-i", sedCommand, confPath)
	sedCommand = "197s:\".*;:\"" + eth0IP + "/24\";:g"
	util.RunCmd(logger, "sed", "-i", sedCommand, confPath)
	// Start enb
	logger.Print("Start enb daemon")
	for {
		retStatus := util.RunCmd(logger, "/snap/bin/oai-ran.enb-start")
		if retStatus.Complete == true {
			break
		}
		logger.Print("Start enb failed, try again later")
		time.Sleep(1 * time.Second)
	}

}
