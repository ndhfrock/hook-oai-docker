package oai

import (
	"errors"
	"fmt"
	"os/exec"
	"time"

	"github.com/hook-oai-docker/internal/pkg/util"
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
	mmeIP = ""
	// Loop until oairan could connect to oaicn
	for mmeIP == "" {
		mmeIP, err = util.GetIPFromDomain(OaiObj.Logger, mmeDomain)
		OaiObj.Logger.Print("MME cannot be connected, trying again in 15 seconds")
		time.Sleep(15 * time.Second)
	}
	OaiObj.Logger.Print("Connected to OAICN !")
	sedCommand = "s:eutra_band.*:      eutra_band              			      = " + c.EutraBand + ";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "s:downlink_frequency.*:      downlink_frequency      			      = " + c.DownlinkFrequency + ";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "s:uplink_frequency_offset.*:      uplink_frequency_offset 			      = " + c.UplinkFrequencyOffset + ";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "175s:\".*;:\"" + mmeIP + "\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	// Get Outbound IP
	outIP := util.GetOutboundIP()
	outInterface, err := util.GetInterfaceByIP(outIP)
	if err != nil {
		OaiObj.Logger.Print(err)
	}
	OaiObj.Logger.Print("Outbound Interfacea and IP is ", outInterface, " ", outIP)
	// Replace interface
	sedCommand = "s/eno1/" + outInterface + "/g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	// Replace enb IP
	sedCommand = "192s:\".*;:\"" + outIP + "/23\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "194s:\".*;:\"" + outIP + "/23\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "197s:\".*;:\"" + outIP + "/24\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "s:parallel_config.*;:parallel_config    = \"PARALLEL_SINGLE_THREAD\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	// Set up FlexRAN
	if OaiObj.Conf.FlexRAN == true {
		// Get flexRAN ip
		var flexranIP string
		OaiObj.Logger.Print("Configure FlexRAN Parameters")
		flexranIP, err = util.GetIPFromDomain(OaiObj.Logger, c.FlexRANDomainName)
		// Loop until oairan could connect to flexran
		for flexranIP == "" {
			flexranIP, err = util.GetIPFromDomain(OaiObj.Logger, c.FlexRANDomainName)
			OaiObj.Logger.Print("FlexRAN cannot be connected, trying again in 5 seconds")
			time.Sleep(5 * time.Second)
		}
		sedCommand = "s:FLEXRAN_ENABLED.*;:FLEXRAN_ENABLED        = \"yes\";:g"
		util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
		sedCommand = "s:FLEXRAN_INTERFACE_NAME.*;:FLEXRAN_INTERFACE_NAME = \"eth0\";:g"
		util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
		sedCommand = "s:FLEXRAN_IPV4_ADDRESS.*;:FLEXRAN_IPV4_ADDRESS   = \"" + flexranIP + "\";:g"
		util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	} else {
		OaiObj.Logger.Print("Disable FlexRAN Feature")
		sedCommand = "s:FLEXRAN_ENABLED.*;:FLEXRAN_ENABLED        = \"no\";:g"
		util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	}
	// Start enb
	if OaiObj.Conf.Test == false {
		time.Sleep(30 * time.Second)
		util.RunCmd(OaiObj.Logger, "/snap/bin/oai-ran.enb-stop")
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

func startENBSlicing(OaiObj Oai) error {
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
	mmeIP = ""
	// Loop until oairan could connect to oaicn
	for mmeIP == "" {
		mmeIP, err = util.GetIPFromDomain(OaiObj.Logger, mmeDomain)
		OaiObj.Logger.Print("MME cannot be connected, trying again in 15 seconds")
		time.Sleep(15 * time.Second)
	}
	sedCommand = "s:eutra_band.*:      eutra_band              			      = " + c.EutraBand + ";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "s:downlink_frequency.*:      downlink_frequency      			      = " + c.DownlinkFrequency + ";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "s:uplink_frequency_offset.*:      uplink_frequency_offset 			      = " + c.UplinkFrequencyOffset + ";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "175s:\".*;:\"" + mmeIP + "\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	// Get Outbound IP
	outIP := util.GetOutboundIP()
	outInterface, err := util.GetInterfaceByIP(outIP)
	if err != nil {
		OaiObj.Logger.Print(err)
	}
	OaiObj.Logger.Print("Outbound Interfacea and IP is ", outInterface, " ", outIP)
	// Replace interface
	sedCommand = "s/eno1/" + outInterface + "/g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	// Replace enb IP
	sedCommand = "192s:\".*;:\"" + outIP + "/23\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "194s:\".*;:\"" + outIP + "/23\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "197s:\".*;:\"" + outIP + "/24\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	sedCommand = "s:parallel_config.*;:parallel_config    = \"PARALLEL_SINGLE_THREAD\";:g"
	util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	// Set up FlexRAN
	if OaiObj.Conf.FlexRAN == true {
		// Get flexRAN ip
		var flexranIP string
		OaiObj.Logger.Print("Configure FlexRAN Parameters")
		flexranIP, err = util.GetIPFromDomain(OaiObj.Logger, c.FlexRANDomainName)
		// Loop until oairan could connect to flexran
		for flexranIP == "" {
			flexranIP, err = util.GetIPFromDomain(OaiObj.Logger, c.FlexRANDomainName)
			OaiObj.Logger.Print("FlexRAN cannot be connected, trying again in 5 seconds")
			time.Sleep(5 * time.Second)
		}
		sedCommand = "s:FLEXRAN_ENABLED.*;:FLEXRAN_ENABLED        = \"yes\";:g"
		util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
		//sedCommand = "s:FLEXRAN_INTERFACE_NAME.*;:FLEXRAN_INTERFACE_NAME = \"eth0\";:g"
		sedCommand = "s:FLEXRAN_INTERFACE_NAME.*;:FLEXRAN_INTERFACE_NAME = \"eth0\";:g"
		util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
		sedCommand = "s:FLEXRAN_IPV4_ADDRESS.*;:FLEXRAN_IPV4_ADDRESS   = \"" + flexranIP + "\";:g"
		util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	} else {
		OaiObj.Logger.Print("Disable FlexRAN Feature")
		sedCommand = "s:FLEXRAN_ENABLED.*;:FLEXRAN_ENABLED        = \"no\";:g"
		util.RunCmd(OaiObj.Logger, "sed", "-i", sedCommand, enbConf)
	}
	// Start enb
	if OaiObj.Conf.Test == false {
		time.Sleep(30 * time.Second)
		OaiObj.Logger.Print("Building and Configuring eNB Finish, Starting eNB")
		out, err := exec.Command("root/run.sh").Output()

		if err != nil {
			fmt.Println("Err", err)
		} else {
			fmt.Println("OUT:", string(out))
		}
		time.Sleep(1 * time.Second)

	}
	return nil
}
