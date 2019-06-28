package oai

import (
	"fmt"
	"oai-snap-in-docker/internal/pkg/util"
	"time"
)

func startENB(OaiObj Oai) {
	c := OaiObj.Conf
	enbConf := c.ConfigurationPathofRAN + "enb.band7.tm1.50PRB.usrpb210.conf"
	mmeDomain := c.MmeDomainName
	// Replace MNC
	// Get MMC MNC from conf
	sedCommand := "18s:\"plmn_list.*;\":\"plmn_list = ( { mcc = " + c.MCC + "; mnc = " + c.MNC + "; mnc_length = 2; } );\":g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	// Get mme ip
	mmeIP, err := util.GetIPFromDomain(OaiObj.Logger, mmeDomain)
	if err != nil {
		fmt.Print(err)
		mmeIP = "10.10.10.10"
	}
	sedCommand = "s:\"eutra_band*\":\"      eutra_band              			      = " + c.EutraBand + ";\":g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "s:\"downlink_frequency*\":\"      downlink_frequency      			      = " + c.DownlinkFrequency + ";\":g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "s:\"uplink_frequency_offset*\":\"      uplink_frequency_offset 			      = " + c.UplinkFrequencyOffset + ";\":g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "175s:\".*;:\"" + mmeIP + "\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	// Replace interface
	util.RunCmd(OaiObj.Logger, "sed", "-i", "191s/eno1/eth0/g", enbConf)
	util.RunCmd(OaiObj.Logger, "sed", "-i", "193s/eno1/eth0/g", enbConf)
	// Replace enb IP
	eth0IP, _ := util.GetInterfaceIP(OaiObj.Logger, "eth0")
	sedCommand = "192s:\".*;:\"" + eth0IP + "/23\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "194s:\".*;:\"" + eth0IP + "/23\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "197s:\".*;:\"" + eth0IP + "/24\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	// Start enb
	OaiObj.Logger.Print("Start enb daemon")
	for {
		retStatus := util.RunCmd(OaiObj.Logger, "/snap/bin/oai-ran.enb-start")
		if retStatus.Complete == true {
			break
		}
		OaiObj.Logger.Print("Start enb failed, try again later")
		time.Sleep(1 * time.Second)
	}

}
