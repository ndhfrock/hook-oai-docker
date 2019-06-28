package oai

import (
	"oai-snap-in-docker/internal/pkg/util"
	"os"
	"strings"
	"time"
)

// InstallSnapCore : Install Core
func installSnapCore(OaiObj Oai) {
	//Install Core
	OaiObj.Logger.Print("Installing core")
	ret, err := util.CheckSnapPackageExist(OaiObj.Logger, "core")
	if err != nil {
		OaiObj.Logger.Print(err)
	}
	//Loop until package is installed
	if !ret {
		retStatus := util.RunCmd(OaiObj.Logger, "snap", "install", "core", "--channel=edge")
		Snapfail := false
		for {
			if len(retStatus.Stderr[0]) > 0 {
				Snapfail = strings.Contains(retStatus.Stderr[0], "error")
			}
			if Snapfail {
				OaiObj.Logger.Print("Wait for snapd being ready")
				time.Sleep(1 * time.Second)
				retStatus = util.RunCmd(OaiObj.Logger, "snap", "install", "core", "--channel=edge")
			} else {
				OaiObj.Logger.Print("snapd is ready and core is installed")
				break
			}
		}
	}

	// Install hello-world
	OaiObj.Logger.Print("Installing hello-world")
	ret, err = util.CheckSnapPackageExist(OaiObj.Logger, "hello-world")
	if err != nil {
		OaiObj.Logger.Print(err)
	}
	if !ret {
		util.RunCmd(OaiObj.Logger, "snap", "install", "hello-world")
	}

}

// InstallOaicn : Install oai-cn snap
func installOaicn(OaiObj Oai) {
	OaiObj.Logger.Print("Configure hostname before installing ")
	// Copy hosts
	util.RunCmd(OaiObj.Logger, "cp", "/etc/hosts", "./hosts_new")
	hostname, _ := os.Hostname()
	fullDomainName := "1s/^/127.0.0.1 " + hostname + ".openair4G.eur " + hostname + " hss\\n127.0.0.1 " + hostname + ".openair4G.eur " + hostname + " mme \\n/"
	util.RunCmd(OaiObj.Logger, "sed", "-i", fullDomainName, "./hosts_new")
	// Replace hosts
	util.RunCmd(OaiObj.Logger, "cp", "-f", "./hosts_new", "/etc/hosts")
	// Install oai-cn snap
	OaiObj.Logger.Print("Installing oai-cn")
	ret, err := util.CheckSnapPackageExist(OaiObj.Logger, "oai-cn")
	if err != nil {
		OaiObj.Logger.Print(err)
	}
	if !ret {
		util.RunCmd(OaiObj.Logger, "snap", "install", "oai-cn", "--channel=edge", "--devmode")
	}

}

// InstallOairan : Install oai-ran snap
func installOairan(OaiObj Oai) {
	// Install oai-ran snap
	OaiObj.Logger.Print("Installing oai-ran")
	ret, err := util.CheckSnapPackageExist(OaiObj.Logger, "oai-ran")
	if err != nil {
		OaiObj.Logger.Print(err)
	}
	if !ret {
		util.RunCmd(OaiObj.Logger, "snap", "install", "oai-ran", "--channel=edge", "--devmode")
	}
	//Wait a moment, cn is not ready yet !
	OaiObj.Logger.Print("Wait 5 seconds... OK now cn should be ready")
	time.Sleep(5 * time.Second)

}
