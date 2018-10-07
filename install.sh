# to be executed on start
docker exec -it cli peer chaincode install -n postap -p github.com/chaincode -v v0.1
docker exec -it cli peer chaincode instantiate -o orderer0.example.com:7050 -C posta-channel -n postap github.com/chaincode -v v0.1 -c '{"Args": []}'