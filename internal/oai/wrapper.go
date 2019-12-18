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

//InstallRANSlicing is a wrapper for installing OAI RAN
func InstallRANSlicing(OaiObj Oai) {

	// Install oai-ran snap
	installOairanSlicing(OaiObj)
}

//StartENB is a wrapper for configuring and starting OAI RAN services
func StartENB(OaiObj Oai) {
	startENB(OaiObj)
}

//StartENBSlicing is a wrapper for configuring and starting OAI RAN services
func StartENBSlicing(OaiObj Oai) {
	startENBSlicing(OaiObj)
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

//InstallStore is a wrapper for installing Store
func InstallStore(OaiObj Oai) {

	// Install store mosaic 5g
	installStore(OaiObj)
}

//StartDrone start drone app from store
func StartDrone(OaiObj Oai) {

	// Start drone store app
	startDrone(OaiObj)
}

//StartRRMKPI start drone app from store
func StartRRMKPI(OaiObj Oai) {

	// Start drone store app
	startRRMKPI(OaiObj)
}
