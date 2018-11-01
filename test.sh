#!/bin/bash

set -ev

# # # Init Ledger
docker exec -it cli peer chaincode invoke -o orderer1.example.com:7050 -n postap -c '{"Args":["initLedger"]}' -C posta-channel

# # # Query all parsels 
docker exec -it cli peer chaincode invoke -o orderer2.example.com:7050 -n postap -c '{"Args":["queryAllParsels"]}' -C posta-channel




