#!/bin/bash

set -ev

# # # Init Ledger
docker exec -it cli peer chaincode invoke -o orderer2.example.com:7050 -n postap -c '{"Args":["initLedger"]}' -C posta-channel

# # # Query all parsels 
docker exec -it cli peer chaincode invoke -o orderer2.example.com:7050 -n postap -c '{"Args":["queryAllParsels"]}' -C posta-channel

# # # Query first parsel
docker exec -it cli peer chaincode invoke -o orderer2.example.com:7050 -n postap -c '{"Args":["queryParsel","1"]}' -C posta-channel

# # # Query parsel by sender
docker exec -it cli peer chaincode invoke -o orderer2.example.com:7050 -n postap -c '{"Args":["querySender","Ben"]}' -C posta-channel

# # # Delivery first parsel in list (only once)
docker exec -it cli peer chaincode invoke -o orderer2.example.com:7050 -n postap -c '{"Args":["deliveryParsel","1",""]}' -C posta-channel


