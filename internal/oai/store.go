package oai

import (
	"fmt"
	"os/exec"

	"github.com/hook-oai-docker/internal/pkg/util"
)

// startDrone : Start Drone StoreApp
func startDrone(OaiObj Oai) {
	// Get Outbound IP
	outIP := util.GetOutboundIP()
	outInterface, err := util.GetInterfaceByIP(outIP)
	if err != nil {
		OaiObj.Logger.Print(err)
	}
	OaiObj.Logger.Print("Outbound Interfacea and IP is ", outInterface, " ", outIP)

	//Starting drone app
	OaiObj.Logger.Print("Starting Drone App")
	out, err := exec.Command("python", "store/sdk/frontend/drone/drone.py", "--port=8080", "--address="+outIP).Output()

	if err != nil {
		fmt.Println("Err", err)
	} else {
		fmt.Println("OUT:", string(out))
	}
}

// startRRMKPI : Start RRMKPI StoreApp
func startRRMKPI(OaiObj Oai) {
	// Get Outbound IP
	outIP := util.GetOutboundIP()
	outInterface, err := util.GetInterfaceByIP(outIP)
	if err != nil {
		OaiObj.Logger.Print(err)
	}
	OaiObj.Logger.Print("Outbound Interfacea and IP is ", outInterface, " ", outIP)

	// Get flexRAN ip
	c := OaiObj.Conf
	var flexranIP string
	OaiObj.Logger.Print("Configure FlexRAN Parameters")
	flexranIP, err = util.GetIPFromDomain(OaiObj.Logger, c.FlexRANDomainName)

	//start rrm_kpi app
	OaiObj.Logger.Print("Starting RRM_KPI App")
	out, err := exec.Command("python", "store/sdk/rrm_kpi_app.py", "--port=9999", "--url=http://"+flexranIP, "--app-url=http://"+outIP, "--app-port=8082").Output()

	if err != nil {
		fmt.Println("Err", err)
	} else {
		fmt.Println("OUT:", string(out))
	}
}
