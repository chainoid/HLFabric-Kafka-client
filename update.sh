#!/bin/sh
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
CHANNEL_NAME=posta-channel

rv="v$(shuf -i 200-10000 -n 1)"

docker exec -it cli peer chaincode install -n postap -p github.com/chaincode -v "$rv"

docker exec -it cli peer chaincode upgrade -o orderer.example.com:7050 -C $CHANNEL_NAME -n postap github.com/chaincode -v "$rv" -c '{"Args": ["queryAllParsels"]}'