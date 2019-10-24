<<<<<<< HEAD
# oai-snap-in-docker

[![Go Report Card](https://goreportcard.com/badge/github.com/tig4605246/snap-hook-for-docker)](https://goreportcard.com/report/github.com/tig4605246/snap-hook-for-docker)

This project includes:
c
md :
- hook: For installing snap and configure files inside docker
- gen: For building new version of OAI snap Docker container

internal/oai
- enb.go : configure and start oai-ran
- flexran.go : configure and start flexran
- hss.go : configure and start hss
- mme.go : configure and start mme
- oai.go : store and log the oai conf file (gotten from common)
- snap.go : install snap core, and mosaic5g snap
- spgw.co : configure and start spgw
- wrapper.go : wrapper for the function

pkg :
- common : map the config file
- util : all required functions for autoconfiguration

## Directory Structure Reference

[Project layout](https://github.com/golang-standards/project-layout)
=======
# hook-oai-docker
hook file for openairinterface automated build inside docker container
>>>>>>> 6e7cb0aa636d16eaeca38a04b2fbc0a0648016c4
