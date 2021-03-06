
# hook-oai-docker
hook file for openairinterface automated build and configuration inside docker container

[![Go Report Card](https://goreportcard.com/badge/github.com/tig4605246/snap-hook-for-docker)](https://goreportcard.com/report/github.com/ndhfrock/hook-oai-docker)

This project includes:
cmd :
- hook: For installing snap and configure files inside docker
- gen: For building new version of OAI snap Docker container

internal/oai
- enb.go : configure and start oai-ran
- flexran.go : configure and start flexran
- hss.go : configure and start hss
- install.go : install snap core, mosaic5g snap, and other oai component
- mme.go : configure and start mme
- oai.go : store and log the oai conf file (gotten from common)
- spgw.co : configure and start spgw
- store.go : configure and start mosaic5g store app
- wrapper.go : wrapper for the function

pkg :
- common : map the config file
- util : all required functions for autoconfiguration

## Golang Directory Structure Reference

[Golang Project layout](https://github.com/golang-standards/project-layout)