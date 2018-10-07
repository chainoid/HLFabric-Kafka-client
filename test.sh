#!/bin/bash

set -ev
docker exec -it cli peer chaincode invoke -o orderer2.example.com:7050 -n postap -c '{"Args":["initLedger"]}' -C posta-channel
