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

// StartHSS is a wrapper for startHss
func StartHSS(OaiObj Oai) {
	// Start HSS
	startHss(OaiObj)
}

// StartMME is a wrapper for startMme
func StartMME(OaiObj Oai) {
	// Start Mme
	startMme(OaiObj)
}

// StartSPGW is a wrapper for startSpgw
func StartSPGW(OaiObj Oai) {
	// Start Mme
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

//InstallFlexRAN is a wrapper for installing FlexRAN
func InstallFlexRAN(OaiObj Oai) {

	// Install flexran snap
	installFlexRAN(OaiObj)
}

//StartFlexRAN is a wrapper for installing FlexRAN
func StartFlexRAN(OaiObj Oai) {

	// start FlexRAN
	startFlexRAN(OaiObj)
}

//InstallMEC is a wrapper for installing LL-MEC
func InstallMEC(OaiObj Oai) {

	// Install ll-mec snap
	installMEC(OaiObj)
}
