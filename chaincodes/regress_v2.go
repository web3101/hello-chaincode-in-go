package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/sajari/regression"
)

type RegressionChaincode struct {
	// is lower level high:.contract.api
}

func (t *RegressionChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *RegressionChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "trainRegression" {
		return t.trainRegression(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"trainRegression\"")
}

func (t *RegressionChaincode) trainRegression(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args)%3 != 0 {
		return shim.Error("Incorrect number of arguments. Expecting multiple of 3")
	}

	r := new(regression.Regression)
	r.SetObserved("Murders per annum per 1,000,000 inhabitants")
	r.SetVar(0, "Inhabitants")
	r.SetVar(1, "Percent with incomes below $5000")
	r.SetVar(2, "Percent unemployed")

	for i := 0; i < len(args); i += 3 {
		inhabitants, _ := strconv.ParseFloat(args[i], 64)
		incomeBelow5000, _ := strconv.ParseFloat(args[i+1], 64)
		unemployed, _ := strconv.ParseFloat(args[i+2], 64)
		r.Train(regression.DataPoint(inhabitants, []float64{incomeBelow5000, unemployed}))
	}

	r.Run()

	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("Regression:\n%s\n", r)

	regressionResult, _ := json.Marshal(r)
	return shim.Success(regressionResult)
}

func main() {
	err := shim.Start(new(RegressionChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
