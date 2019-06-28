package oai

//InstallSnap is a wrapper function for installSnapCore
func InstallSnap(OaiObj Oai) {
	// Install Snap Core
	installSnapCore(OaiObj)
}

//InstallCN is a wrapper for installing OAI CN
func InstallCN(OaiObj Oai) {

	// Install oai-cn snap
	installOaicn(OaiObj)

}

// StartCN is a wrapper for configuring and starting OAI CN services
func StartCN(OaiObj Oai) {
	// Start HSS
	startHss(OaiObj)
	// Start MME
	startMme(OaiObj)
	// Start SPGW
	startSpgw(OaiObj)
}

//InstallRAN is a wrapper for installing OAI RAN
func InstallRAN(OaiObj Oai) {

	// Install oai-ran snap
	installOairan(OaiObj)
}

//StartENB is a wrapper for configuring and starting OAI RAN services
func StartENB(OaiObj Oai) {
	startENB(OaiObj)
}
