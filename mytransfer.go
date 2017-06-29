/*
Edit by Li Yicong
8/9/2016
*/

package main

import (
	/*"encoding/json"*/
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"github.com/hyperledger/fabric/core/crypto/primitives"
	//"github.com/op/go-logging"

)

//simple Chaincode implementation
type SimpleChaincode struct {

}

//Init is a no-op
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	var A,B string  //Entities
	var Aval, Bval int  //Asset holdings
	var err error

	if len(args) != 4 {
		return nil, errors.New("there must be 4 arguments")
	}

	//Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("the second value should be an integer")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return nil, errors.New("the fourth value should be an integer")
	}

	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	//Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

//Invoke has three functions
//put - takes two arguments, a key and value and stores them in the state
//remove - takes one argument, a key, and removes the key and value from the state
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte,error) {
	
	/*var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value
	var transfersth string
	var err error*/
	switch function {
	case "put":
		if len(args) < 2 {
			return nil, errors.New("there must be two arguments, a key and value, eg: ['car','Lina']")
		}
		key := args[0]
		value := args[1]

		err := stub.PutState(key, []byte(value))
		if err != nil {
			fmt.Printf("Error putting state %s", err)
			return nil, fmt.Errorf("put operation failed. Error updating state: %s", err)
		}
		return nil, nil

	case "remove":
		if len(args) < 1 {
			return nil, errors.New("there must be one argument, a key, eg: ['car']")
		}
		key := args[0]

		err := stub.DelState(key)
		if err!= nil {
			return nil, fmt.Errorf("remove operation failed. Error updating state: %s", err)
		}
		return nil, nil
/*
	case "transfer":
		if len(args) != 3 {
			return nil, errors.New("there must be three arguments")
		}

		var A, B string     //Entities
		var Aval, Bval int  //Asset holdings
		var X int           //Transaction value
		var asset string
		var assetholder string
		var err error

		A = args[0]
		B = args[1]

		//Get the state from the ledger
		//TODO: will be nice to have GetAllState call to ledger
		Avalbytes, err := stub.GetState(A)
		if err != nil {
			return nil, errors.New("failed to get state")
		}
		if Avalbytes == nil {
			return nil, errors.New("Entity not found")
		}
		Aval, _ = strconv.Atoi(string(Avalbytes))

		Bvalbytes, err := stub.GetState(B)
		if err != nil {
			return nil, errors.New("failed to get state")
		}
		if Bvalbytes == nil {
			return nil, errors.New("Entity not found")
		}
		Bval, _ = strconv.Atoi(string(Bvalbytes))

		//Perform the execution
		X, err = strconv.Atoi("20")
		Aval = Aval + X
		Bval = Bval - X
		fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

		//Write the state back to the ledger
		err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
		if err != nil {
			return nil, err
		}

		err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
		if err != nil {
			return nil, err
		}

		err = stub.PutState(asset, []byte(B))
		if err != nil {
			return nil, err
		}

		return nil, nil
*/
/*
	case "transfer":
		//var A, B string    // Entities
		//var Aval, Bval int // Asset holdings
		//var X int          // Transaction value
		//var transfersth string
		//var err error

		if len(args) != 3 {
			return nil, errors.New("Incorrect number of arguments. Expecting 3")
		}

		A = args[0]
		B = args[1]

		// Get the state from the ledger
		// TODO: will be nice to have a GetAllState call to ledger
		Avalbytes, err := stub.GetState(A)
		if err != nil {
			return nil, errors.New("Failed to get state")
		}
		if Avalbytes == nil {
			return nil, errors.New("Entity not found")
		}
		Aval, _ = strconv.Atoi(string(Avalbytes))

		Bvalbytes, err := stub.GetState(B)
		if err != nil {
			return nil, errors.New("Failed to get state")
		}
		if Bvalbytes == nil {
			return nil, errors.New("Entity not found")
		}
		Bval, _ = strconv.Atoi(string(Bvalbytes))

		// Perform the execution
		//X, err = strconv.Atoi(args[2])
		//transfersth, err = args[2]
		Aval = Aval + 20
		Bval = Bval - 20
		fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

		// Write the state back to the ledger
		err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
		if err != nil {
			return nil, err
		}

		err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
		if err != nil {
			return nil, err
		}

		return nil, nil
*/
	default:
		return nil, errors.New("Unsupported operation")
	}
}

func (t *SimpleChaincode) Transfer(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

		var A, B string    // Entities
		var Aval, Bval int // Asset holdings
		var X int          // Transaction value
		var transfersth string
		var err error

		if len(args) != 3 {
			return nil, errors.New("Incorrect number of arguments. Expecting 3")
		}

		A = args[0]
		B = args[1]

		// Get the state from the ledger
		// TODO: will be nice to have a GetAllState call to ledger
		Avalbytes, err := stub.GetState(A)
		if err != nil {
			return nil, errors.New("Failed to get state")
		}
		if Avalbytes == nil {
			return nil, errors.New("Entity not found")
		}
		Aval, _ = strconv.Atoi(string(Avalbytes))

		Bvalbytes, err := stub.GetState(B)
		if err != nil {
			return nil, errors.New("Failed to get state")
		}
		if Bvalbytes == nil {
			return nil, errors.New("Entity not found")
		}
		Bval, _ = strconv.Atoi(string(Bvalbytes))

		// Perform the execution
		//X, err = strconv.Atoi(args[2])
		//transfersth, err = args[2]
		Aval = Aval + 20
		Bval = Bval - 20
		fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

		// Write the state back to the ledger
		err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
		if err != nil {
			return nil, err
		}

		err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
		if err != nil {
			return nil, err
		}

		return nil, nil
}


// maybe there is something wrong
//Query callback representing the query of a chaincode

func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	switch function {

	case "get":
		if len(args) < 1 {
			return nil, errors.New("get operation must include one argument, a key")
		}
		key := args[0]
		value, err := stub.GetState(key)
		if err != nil {
			return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
		}
		return value, nil
/*
	case "query":
		var A string   //Entities
		var err error

		if len(args) != 1 {
			return nil, errors.New("there should be one argument")
		}

		A = args[0]

		//get the state of ledger
		Avalbytes, err := stub.GetState(A)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
			return nil, errors.New(jsonResp)
		}

		if Avalbytes == nil {
			jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
			return nil, errors.New(jsonResp)
		}

		jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
		fmt.Printf("Query Response:%s\n", jsonResp)
		return Avalbytes, nil

*/
	case "keys":

		keysIter, err := stub.RangeQueryState("", "")
		if err != nil {
			return nil, fmt.Errorf("keys operation failed. Error accessing state: %s", err)
		}
		defer keysIter.Close()

		var keys []string
		for keysIter.HasNext() {
			key, _, iterErr := keysIter.Next()
			if iterErr != nil {
				return nil, fmt.Errorf("keys operation failed. Error accessing state: %s", err)
			}
			keys = append(keys, key)
		}

		jsonKeys, err := json.Marshal(keys)
		if err != nil {
			return nil, fmt.Errorf("keys operation failed. Error marshaling JSON: %s", err)
		}

		return jsonKeys, nil

	default:
		return nil, errors.New("Unsupported operation")
	}
}


func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}