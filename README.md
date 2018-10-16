# HLFabric-Kafka-client

Hyperledger Fabric with Kafka Consensus: network configuration and UI client example.

## Introduction

The Hyperledger Fabric has introduced Kafka as it’s primary consensus mechanism among the orderers. While in development, for testing purposes, a solo config orderer is used. However, in production, you need to have multiple orderer nodes set up to have fail proof systems. In the case of a hardware/software failure, this is what will rescue you from critical situations. Kafka helps implement this easily. To understand in detail how all of this works, refer to this article [TODO]. 


## Assumptions

You already know how to work with Hyperledger Fabric and deploy chaincode to the network in development. You are also comfortable with docker and docker-compose. To get comfortable with these topics, refer to these [articles](https://www.skcript.com/svr/blockchain/).


## Network Architecture

We have a simple network configuration
- 2 Certificate Authorities.
- 3 Orderers.
- 3 Orderers.
- 2 Organizations.
- 4 peers, 2 for each organization.
- 4 Kafka broker instances.
- 3 Zookeper instances.

## Technical Documentation

### Folder Structure
```
./
├── app
│   ├── client
│   │   ├── app.js
│   │   ├── favicon.png
│   │   └── index.html
│   ├── controller.js
│   ├── package.json
│   ├── package-lock.json
│   ├── README.md
│   ├── registerAdmin.js
│   ├── registerUser.js
│   ├── routes.js
│   ├── server.js
│   └── src
│       ├── acceptParsel.js
│       ├── deliveryParsel.js
│       ├── queryAllParsels.js
│       ├── queryParsel.js
│       └── querySender.js
├── bin
│   ├── configtxgen
│   ├── configtxlator
│   ├── cryptogen
│   ├── fabric-ca-client
│   ├── get-docker-images.sh
│   ├── orderer
│   └── peer
├── chaincode
│   └── postap.go
├── cleanimages.sh
├── generate.sh
├── install.sh
├── LICENSE
├── network-config
│   ├── configtx.yaml
│   ├── crypto-config.yaml
│   ├── docker-compose-base.yml
│   ├── docker-compose-cli.yml
│   ├── docker-compose-couchdb.yml
│   └── docker-compose-kafka.yml
├── queryAllParsels.sh
├── README.md
├── start.sh
├── stop.sh
├── teardown.sh
├── test.sh
└── update.sh

6 directories, 41 files

```
### Description

- `app/` tiny client written in node.js.
- `bin/` contains all the binaries required for generating crypto material.
- `chaincode/` contains the postap chaincode, easy parsel delivery example.
- `network-config/` contains all the configuration `yaml` files for the network.
- `generate.sh` will generate all the crypto-material required for the network to run.
- `install.sh` will install and instantiate the chaincode.
- `start.sh` will start all the containers in docker-compose files.
- `stop.sh` will stop and remove all docker containers, use with caution.
- `teardown.sh` will kill containers and remove images generated by the network. 
- `test.sh` is used to test an invoke query.
- `update.sh` is used to install and upgrade chaincode.

### Steps

- Make sure you've grabbed all the [prerequisites](http://hyperledger-fabric.readthedocs.io/en/release-1.2/prereqs.html) and [samples](http://hyperledger-fabric.readthedocs.io/en/release-1.2/samples.html#) for running hyperledger fabric.
- In your `$GOPATH`, make sure you have the hyperledger fabric source.
```bash
cd $GOPATH/src/github.com/
mkdir hyperledger
cd hyperledger
git clone https://github.com/hyperledger/fabric.git
```
- Clone this repository and enter the directory.
```bash
git clone https://github.com/chainoid/HLFabric-Kafka-client.git
cd Kafka-Fabric-Network
```
NOTE : Give exec permissions to the shell scripts
```chmod 777 ./script-name.sh```
- Generate the crypto-material
```bash
./generate.sh
```
- Start the network
```bash
./start.sh
```
- Install and Instantiate the chaincode
```bash
./install.sh
```
- Verify if all the docker containers are running
```bash
docker ps 
```
- Test a sample invoke command
```bash
./test.sh
```
- Stop every docker container and kill all the containers
```bash
./stop.sh
```
- For Instruction how to run a client see /app/README.md
```bash
cd app
cat README.md
```
