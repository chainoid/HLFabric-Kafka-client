## Hyperledger Fabric Sample Application - Parsel Service

This application demonstrates the creation and transfer of parsel operation between customers using Hyperledger Fabric body  as supply chain.

## Basic Network Config
Note that this basic configuration uses pre-generated certificates and key material, and also has predefined transactions to initialize a channel named "posta-channel".

For install required node packages:

app$ npm install

For register Admin and User into certification center:

app$ node registerAdmin.js
app$ node registerUser.js

After network started (please see the top level README.md) verify that all containers are up:

app$ docker ps

And finally run the client application:

app$ node server.js

The UI client will available on the link:

http://localhost:9000


## Acknowledgements
This code is based on code written by the Hyperledger Fabric community. Source code can be found here: (https://github.com/hyperledger/fabric-samples).

The scripts are inspired by first-network and fabcar of Hyperledger Fabric samples.
