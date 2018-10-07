docker exec -it cli peer chaincode invoke -o orderer2.example.com:7050 -n postap -c '{"Args":["queryAllParsels"]}' -C posta-channel
