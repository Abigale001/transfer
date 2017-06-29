/*
Edited by Li Yicong
2016/8/20
*/
package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state to the ledger
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
/*
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	var A string    // Entities
	var Aval int // Asset holdings
	var err error
	var length int
	var i int

	if len(args)%2 != 0 || len(args) == 0 {
		return nil, errors.New("Incorrect number of arguments.")
	}

	length = len(args)

	for i := 0; i < length; i += 2 {
		A = args[i]
		Aval, err = strconv.Atoi(args[i+1])
		if err != nil {
			return nil, errors.New("Expecting integer value for asset holding")
		}

		fmt.Printf("asdf")
		err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
		if err != nil {
			return nil, err
		}

	}
	return nil, nil
}
*/
// give salary and late
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function == "delete"{
		// Deletes an entity from its state
		return t.delete(stub, args)
	}
	switch function {
	case "giveSalary":
		var A string    // Entities
		var Aval int // Asset holdings
		var err error

		if len(args) > 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting 1")
		}

		A = args[0]

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

		// Perform the execution

		Aval = Aval + 5000
		fmt.Printf("You have got your salary")
		// Write the state back to the ledger
		err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
		if err != nil {
			return nil, err
		}

		return nil, nil

	case "beLate":
		var A string    // Entities
		var Aval int // Asset holdings
		var X int
		var err error

		if len(args) < 2 {
			return nil, errors.New("Incorrect number of arguments. Expecting 2")
		}

		A = args[0]

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

		// Perform the execution
		X, err = strconv.Atoi(args[1])
		if X > 0 && X <= 5 {
			Aval = Aval - 50
		}
		if X > 5 && X <= 15 {
			Aval = Aval - 75
		}
		if X > 15 {
			Aval = Aval - 200
		}
		
		fmt.Printf("You are late")
		// Write the state back to the ledger
		err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
		if err != nil {
			return nil, err
		}

		return nil, nil

	case "workOvertime":
		var A string    // Entities
		var Aval int // Asset holdings
		var X int
		var err error

		if len(args) < 2 {
			return nil, errors.New("Incorrect number of arguments. Expecting 2")
		}

		A = args[0]

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

		// Perform the execution
		X, err = strconv.Atoi(args[1])
		if X > 0 && X <= 2 {
			Aval = Aval + 20
		}
		if X > 2 && X <= 4 {
			Aval = Aval + 50
		}
		if X > 4 && X <= 8 {
			Aval = Aval + 120
		}
		if X > 8 {
			Aval = Aval + 200
		}
		
		fmt.Printf("You are late")
		// Write the state back to the ledger
		err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
		if err != nil {
			return nil, err
		}

		return nil, nil

	case "getSalary":
		var A string    // Entities
		var Aval int // Asset holdings
		var X int
		var err error

		if len(args) < 2 {
			return nil, errors.New("Incorrect number of arguments. Expecting 2")
		}

		A = args[0]

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

		// Perform the execution

		X, err = strconv.Atoi(args[1])
		Aval = Aval - X
		fmt.Printf("YOU HAVE GOT YOUR MONEY")
		// Write the state back to the ledger
		err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
		if err != nil {
			return nil, err
		}

		return nil, nil

	default:
		return nil, errors.New("Unsupported operation")
	}
	
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}
	fmt.Printf("asdf")
	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var A string // Entities
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
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
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
