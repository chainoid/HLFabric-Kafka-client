#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error.
set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

starttime=$(date +%s)

docker-compose -f ./network-config/docker-compose-kafka.yml down
docker-compose -f ./network-config/docker-compose-cli.yml down
#docker-compose -f ./network-config/docker-compose-couchdb.yml down

docker-compose -f ./network-config/docker-compose-kafka.yml up -d
docker-compose -f ./network-config/docker-compose-cli.yml up -d
#docker-compose -f ./network-config/docker-compose-couchdb.yml up -d

# wait for Hyperledger Fabric to start
# incase of errors when running later commands, issue export FABRIC_START_TIMEOUT=<larger number>
export FABRIC_START_TIMEOUT=15
echo ${FABRIC_START_TIMEOUT}
sleep ${FABRIC_START_TIMEOUT}

# # # Create the channel
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/var/hyperledger/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel create -o orderer0.example.com:7050 -c posta-channel -f /var/hyperledger/configs/channel.tx
# # # Join peer0.org1.example.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/var/hyperledger/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel join -b posta-channel.block

printf "\nTotal execution time : $(($(date +%s) - starttime)) secs ...\n\n"
printf "\nNext step is install chaincode:  ./install.sh\n\n"