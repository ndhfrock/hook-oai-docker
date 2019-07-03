package oai

import (
	"errors"
	"fmt"
	"oai-snap-in-docker/internal/pkg/util"
	"time"
)

func startENB(OaiObj Oai) error {
	c := OaiObj.Conf
	enbConf := c.ConfigurationPathofRAN + "enb.band7.tm1.50PRB.usrpb210.conf"
	mmeDomain := c.MmeDomainName
	// Replace MCC
	sedCommand := "s/mcc =.[^;]*/mcc = " + c.MCC + "/g"
	OaiObj.Logger.Print(sedCommand)
	retStatus := util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	if retStatus.Exit != 0 {
		return errors.New("Set MCC in " + enbConf + " failed")
	}
	OaiObj.Logger.Print("Replace MNC")
	//Replace MNC
	sedCommand = "s/mnc =.[^;]*/mnc = " + c.MNC + "/g"
	OaiObj.Logger.Print(sedCommand)
	retStatus = util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	if retStatus.Exit != 0 {
		return errors.New("Set MNC in " + enbConf + " failed")
	}
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
	// Set up FlexRAN
	if OaiObj.Conf.FlexRAN == true {
		// Get flexRAN ip
		var flexranIP string
		OaiObj.Logger.Print("Configure FlexRAN Parameters")
		flexranIP, err = util.GetIPFromDomain(OaiObj.Logger, c.FlexRANDomainName)
		if err != nil {
			OaiObj.Logger.Print(err)
			OaiObj.Logger.Print("Getting IP of FlexRAN failed, try again later")
		}
		sedCommand = "s:FLEXRAN_ENABLED.*;:FLEXRAN_ENABLED=        \"yes\";:g"
		util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
		sedCommand = "s:FLEXRAN_INTERFACE_NAME.*;:FLEXRAN_INTERFACE_NAME= \"eth0\";:g"
		util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
		sedCommand = "s:FLEXRAN_IPV4_ADDRESS.*;:FLEXRAN_IPV4_ADDRESS   = \"" + flexranIP + "\";:g"
		util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	} else {
		OaiObj.Logger.Print("Disable FlexRAN Feature")
		sedCommand = "s:FLEXRAN_ENABLED.*;:FLEXRAN_ENABLED=        \"no\";:g"
		util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	}
	// Start enb
	if OaiObj.Conf.Test == false {
		OaiObj.Logger.Print("Start enb daemon")
		for {
			retStatus := util.RunCmd(OaiObj.Logger, "/snap/bin/oai-ran.enb-start")
			if len(retStatus.Stderr) == 0 {
				break
			}
			OaiObj.Logger.Print("Start enb failed, try again later")
			time.Sleep(1 * time.Second)
		}
	}
	return nil
}
