package main

// This APP is made for installing snaps in docker
import (
	"log"
	"os"

	"github.com/go-cmd/cmd"
)

func main() {
	// time.Sleep(10 * time.Second)
	// Initialize Log
	logFile, err := os.Create("/root/hook.log")
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "[Debug] ", log.Lshortfile)
	//Install Core
	logger.Print("Installing core")
	installSnap := cmd.NewCmd("snap", "install", "core", "--channel=edge")
	finalStatus := <-installSnap.Start() // block and wait
	logger.Print(finalStatus)
	// Install hello-world
	logger.Print("Installing hello-world")
	installSnap = cmd.NewCmd("snap", "install", "hello-world")
	finalStatus = <-installSnap.Start() // block and wait
	logger.Print(finalStatus)
	// Configure hostname
	logger.Print("Configure hostname before installing ")
	installSnap = cmd.NewCmd("cp", "/etc/hosts", "./hosts_new")
	finalStatus = <-installSnap.Start() // block and wait
	logger.Print(finalStatus)
	installSnap = cmd.NewCmd("sed", "-i", "1s/^/127.0.0.1 ubuntu.openair4G.eur ubuntu hss\\n127.0.0.1 ubuntu.openair4G.eur ubuntu mme \\n/", "./hosts_new")
	finalStatus = <-installSnap.Start() // block and wait
	logger.Print(finalStatus)
	installSnap = cmd.NewCmd("cp", "-f", "./hosts_new", "/etc/hosts")
	finalStatus = <-installSnap.Start() // block and wait
	logger.Print(finalStatus)
	// Install oai-cn snap
	logger.Print("Installing oai-cn")
	installSnap = cmd.NewCmd("snap", "install", "oai-cn", "--channel=edge", "--devmode")
	finalStatus = <-installSnap.Start() // block and wait
	logger.Print(finalStatus)
	// Configure oai-hss
	logger.Print("Configure oai-hss")
	installSnap = cmd.NewCmd("sed", "-i", "-r", "31s/.{1}//", "/var/snap/oai-cn/28/hss.conf")
	finalStatus = <-installSnap.Start() // block and wait
	logger.Print(finalStatus)
	installSnap = cmd.NewCmd("sed", "-i", "30s/^/#/", "/var/snap/oai-cn/28/hss.conf")
	finalStatus = <-installSnap.Start() // block and wait
	logger.Print(finalStatus)
	installSnap = cmd.NewCmd("sed", "-i", "s/127.0.0.1/mysql/g", "/var/snap/oai-cn/28/hss.conf")
	finalStatus = <-installSnap.Start() // block and wait
	logger.Print(finalStatus)
	// Init hss
	logger.Print("Init hss")
	installSnap = cmd.NewCmd("/snap/bin/oai-cn.hss-init")
	finalStatus = <-installSnap.Start() // block and wait
	logger.Print(finalStatus)
	// oai-cn.hss-start
	logger.Print("start hss as daemon")
	installSnap = cmd.NewCmd("/snap/bin/oai-cn.hss-start")
	finalStatus = <-installSnap.Start() // block and wait
	logger.Print(finalStatus)
}

func runCmd(logger *log.Logger, cmd string, args ...string) {

}
