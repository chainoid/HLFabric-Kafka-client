#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error.
set -e

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

starttime=$(date +%s)

docker-compose -f ./network-config/docker-compose-kafka.yml down
docker-compose -f ./network-config/docker-compose-cli.yml down


docker-compose -f ./network-config/docker-compose-kafka.yml up -d
docker-compose -f ./network-config/docker-compose-cli.yml up -d


# wait for Hyperledger Fabric to start
# incase of errors when running later commands, issue export FABRIC_START_TIMEOUT=<larger number>
#export FABRIC_START_TIMEOUT=5
#echo ${FABRIC_START_TIMEOUT}
#sleep ${FABRIC_START_TIMEOUT}


printf "\nTotal execution time : $(($(date +%s) - starttime)) secs ...\n\n"
printf "\nStart client with: node server.js\n\n"	

