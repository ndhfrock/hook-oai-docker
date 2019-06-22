package oai

import "log"

//InstallSnap is a wrapper function for installSnapCore
func InstallSnap(logger *log.Logger) {
	// Install Snap Core
	installSnapCore(logger)
}

//InstallCN is a wrapper for installing OAI CN
func InstallCN(logger *log.Logger) {

	// Install oai-cn snap
	installOaicn(logger)

}

// StartCN is a wrapper for configuring and starting OAI CN services
func StartCN(logger *log.Logger) {
	// Start HSS
	startHss(logger)
	// Start MME
	startMme(logger)
	// Start SPGW
	startSpgw(logger)
}

//InstallRAN is a wrapper for installing OAI RAN
func InstallRAN(logger *log.Logger) {

	// Install oai-ran snap
	installOairan(logger)
}

//StartRAN is a wrapper for configuring and starting OAI RAN services
func StartRAN(logger *log.Logger) {
}
