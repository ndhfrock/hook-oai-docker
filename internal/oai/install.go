package oai

import (
	"os"
	"os/exec"
	"time"

	"github.com/hook-oai-docker/internal/pkg/util"
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// installSnapCore : Install Core
func installSnapCore(OaiObj Oai) {
	//Install Core
	OaiObj.Logger.Print("Installing core")
	ret, err := util.CheckSnapPackageExist(OaiObj.Logger, "core")
	if err != nil {
		OaiObj.Logger.Print(err)
	}
	//Loop until package is installed
	if !ret {
		util.RunCmd(OaiObj.Logger, "snap", "install", "core", "--channel=edge")
		//Snapfail := false
		for ret == false {
			util.RunCmd(OaiObj.Logger, "snap", "install", "core", "--channel=edge")
			OaiObj.Logger.Print("Wait for snapd being ready")
			time.Sleep(5 * time.Second)
			ret, err = util.CheckSnapPackageExist(OaiObj.Logger, "core")
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

// installOaicn : Install oai-cn snap
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

// installOairan : Install oai-ran snap
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
	time.Sleep(10 * time.Second)

}

// installOairanSlicing : Install oai-ran from samuel repository branch "flexran_LTE_slicing_integration"
func installOairanSlicing(OaiObj Oai) {
	//if repo has not been cloned
	if _, err := os.Stat("/LTE_Mac_scheduler_with_network_slicing"); os.IsNotExist(err) {
		// Install oai-ran snap
		OaiObj.Logger.Print("Installing oai-ran")
		// Clone the given repository to the given directory
		Info("git clone https://gitlab.com/changshengliusamuel/LTE_Mac_scheduler_with_network_slicing.git")

		_, err := git.PlainClone("/LTE_Mac_scheduler_with_network_slicing", false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: "nadhifrock",
				Password: "Scruzers97",
			},
			URL:           "https://gitlab.com/changshengliusamuel/LTE_Mac_scheduler_with_network_slicing.git",
			ReferenceName: plumbing.NewBranchReferenceName("flexran_LTE_slicing_integration"),
			Progress:      os.Stdout,
		})

		CheckIfError(err)
		if err != nil {
			OaiObj.Logger.Print(err)
		}

		//build enb
		OaiObj.Logger.Print("Building eNB with USRP...")
		out, err := exec.Command("root/buildenb.sh").Output()

		if err != nil {
			OaiObj.Logger.Print("Err", err)
		} else {
			OaiObj.Logger.Print("Building the eNB...")
			OaiObj.Logger.Print("OUT:", string(out))
		}
	}
	//if repo has been cloned
	if _, err := os.Stat("/LTE_Mac_scheduler_with_network_slicing"); !os.IsNotExist(err) {
		OaiObj.Logger.Print("eNB already cloned and built")
	}
}

// installFlexRAN : Install FlexRAN snap
func installFlexRAN(OaiObj Oai) {
	// Install FlexRAN snap
	OaiObj.Logger.Print("Installing FlexRAN")
	ret, err := util.CheckSnapPackageExist(OaiObj.Logger, "flexran")
	if err != nil {
		OaiObj.Logger.Print(err)
	}
	if !ret {
		util.RunCmd(OaiObj.Logger, "snap", "install", "flexran", "--channel=edge", "--devmode")
	}
	//Wait a moment, cn is not ready yet !
	OaiObj.Logger.Print("Wait 5 seconds... OK now flexran should be ready")
	time.Sleep(5 * time.Second)

}

// installMEC : Install LL-MEC snap
func installMEC(OaiObj Oai) {
	// Install LL-MEC snap
	OaiObj.Logger.Print("Installing LL-MEC")
	ret, err := util.CheckSnapPackageExist(OaiObj.Logger, "ll-mec")
	if err != nil {
		OaiObj.Logger.Print(err)
	}
	if !ret {
		util.RunCmd(OaiObj.Logger, "snap", "install", "ll-mec", "--channel=edge", "--devmode")
	}
	//Wait a moment, cn is not ready yet !
	OaiObj.Logger.Print("Wait 5 seconds... OK now ll-mec should be ready")
	time.Sleep(5 * time.Second)

}
