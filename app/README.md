## Hyperledger Fabric Sample Application - Parsel Service

This application demonstrates the creation and transfer of parsel operation between customers using Hyperledger Fabric body  as supply chain.
In this exercise we will set up the minimum number of nodes required to develop chaincode. It has a single peer and a single organization.

if getting error about running ./startFabric.sh permission

chmod a+x startFabric.sh


## Basic Network Config
Note that this basic configuration uses pre-generated certificates and key material, and also has predefined transactions to initialize a channel named "mychannel".

To regenerate this material, simply run generate.sh.

To start the network, run start.sh. To stop it, run stop.sh To completely remove all incriminating evidence of the network on your system, run teardown.sh.


## Acknowledgements
This code is based on code written by the Hyperledger Fabric community. Source code can be found here: (https://github.com/hyperledger/fabric-samples).

The scripts are inspired by first-network and fabcar of Hyperledger Fabric samples.
