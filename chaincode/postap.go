// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
 */

package main

/* Imports  
* 4 utility libraries for handling bytes, reading and writing JSON, 
formatting, and string manipulation  
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts  
*/ 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Parsel structure, with several properties.  
Structure tags are used by encoding/json library
*/
type Parsel struct {
	Sender string `json:"sender"`
	SenderTS string `json:"senderTS"`
	SenderBranch  string `json:"senderBranch"`
	
	Receiver string `json:"receiver"`
	ReceiverTS string `json:"receiverTS"`
	ReceiverBranch string `json:"receiverBranch"`
}


/*
 * The Init method *
 called when the Smart Contract "posta-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function 
 -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "posta-chaincode"
 The app also specifies the specific smart contract function to call with args
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "queryParsel" {
		return s.queryParsel(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "acceptParsel" {
		return s.acceptParsel(APIstub, args)
	} else if function == "queryAllParsels" {
		return s.queryAllParsels(APIstub)
	} else if function == "deliveryParsel" {
		return s.deliveryParsel(APIstub, args)
	} else if function == "querySender" {
		return s.querySender(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The queryParsel method *
 Used to view the records of one particular parsel
 It takes one argument -- the key for the parsel in question
 */
func (s *SmartContract) queryParsel(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	parselAsBytes, _ := APIstub.GetState(args[0])
	if parselAsBytes == nil {
		return shim.Error("Could not locate parsel")
	}
	return shim.Success(parselAsBytes)
}

/*
 * The initLedger method *
Will add test data (5 parsels)to our network
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	parsel := []Parsel{
		Parsel{Sender: "Alex",  SenderTS: time.Now().Format(time.RFC1123Z),  SenderBranch: "001", Receiver: "Miriam", ReceiverTS: "", ReceiverBranch: "101"},
		Parsel{Sender: "Ben",   SenderTS: time.Now().Format(time.RFC1123Z),  SenderBranch: "002", Receiver: "Elvis",  ReceiverTS: "", ReceiverBranch: "102"},
		Parsel{Sender: "Charly",SenderTS: time.Now().Format(time.RFC1123Z),  SenderBranch: "003", Receiver: "Elvis",  ReceiverTS: "", ReceiverBranch: "103"},
		Parsel{Sender: "Dan",   SenderTS: time.Now().Format(time.RFC1123Z),  SenderBranch: "004", Receiver: "Alex",   ReceiverTS: "", ReceiverBranch: "104"},
		Parsel{Sender: "Elvis", SenderTS: time.Now().Format(time.RFC1123Z),  SenderBranch: "005", Receiver: "Alex",   ReceiverTS: "", ReceiverBranch: "105"},
		Parsel{Sender: "Sega",  SenderTS: time.Now().Format(time.RFC1123Z),  SenderBranch: "005", Receiver: "Mary",   ReceiverTS: "", ReceiverBranch: "105"},
	}

	i := 0
	for i < len(parsel) {
		fmt.Println("i is ", i)
		parselAsBytes, _ := json.Marshal(parsel[i])
		APIstub.PutState(strconv.Itoa(i+1), parselAsBytes)
		fmt.Println("Added", parsel[i])
		i = i + 1
	}

	return shim.Success(nil)
}


/*
 * The acceptParsel method *
   In the Post office would use to accept parsel from sender person. 
   This method takes in five arguments (attributes to be saved in the ledger). 
 */

  func (s *SmartContract) acceptParsel(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	var parsel = Parsel{ Sender: args[1], SenderBranch: args[2], SenderTS: time.Now().Format(time.RFC1123Z), Receiver: args[4], ReceiverBranch: args[5], ReceiverTS: args[6] }

	parselAsBytes, _ := json.Marshal(parsel)
	err := APIstub.PutState(args[0], parselAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record new parsel: %s", args[0]))
	}

	return shim.Success(nil)
}


/*
 * The queryAllParsels method *
allows for assessing all the records added to the ledger(all parsels in the delivery system)
This method does not take any arguments. Returns JSON string containing results. 
 */
func (s *SmartContract) queryAllParsels(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "9999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllParsels:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


/*
 * The querySender method *
allows for assessing all the records from selected sender

Returns JSON string containing results. 
 */
func (s *SmartContract) querySender(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	startKey := "0"
	endKey := "9999"
		
	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
			
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
				
		// Create an object
		parsel := Parsel{}
		// Unmarshal record to parsel
		json.Unmarshal(queryResponse.Value, &parsel)
				
		// Add only filtered ny sender records
		if parsel.Sender == args[0] { 
		
		// Add comma before array members,suppress it for the first array member
			if bArrayMemberAlreadyWritten == true {
				buffer.WriteString(",")
			}
						
			buffer.WriteString("{\"Key\":")
			buffer.WriteString("\"")
			buffer.WriteString(queryResponse.Key)
			buffer.WriteString("\"")

			buffer.WriteString(", \"Record\":")
			// Record is a JSON object, so we write as-is
			buffer.WriteString(string(queryResponse.Value))
			buffer.WriteString("}")
			bArrayMemberAlreadyWritten = true
	    }
	}
	buffer.WriteString("]")
	
	if bArrayMemberAlreadyWritten == false {
		return shim.Error("No parsels for sender")
	}

	fmt.Printf("- querySender:\n%s\n", buffer.String())
		
	return shim.Success(buffer.Bytes())
}


/*
 * The deliveryParsel method *
The data in the world state can be updated with who has possession. 
This function takes in 2 arguments, parsel id and timestamp of delivery. 
 */
func (s *SmartContract) deliveryParsel(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	parselAsBytes, _ := APIstub.GetState(args[0])
	if parselAsBytes == nil {
		return shim.Error("Could not locate parsel")
	}
	parsel := Parsel{}

	json.Unmarshal(parselAsBytes, &parsel)
	// Normally check that the specified argument is a valid holder of parsel
	// we are skipping this check for this example
	parsel.ReceiverTS = time.Now().Format(time.RFC1123Z)

	parselAsBytes, _ = json.Marshal(parsel)
	err := APIstub.PutState(args[0], parselAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change status of parsel: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * main function *
calls the Start function 
The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
