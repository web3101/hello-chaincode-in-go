package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type HelloWorldContract struct {
	contractapi.Contract
}

func (c *HelloWorldContract) HelloWorld(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("Hello, World!")

	return nil
}

func main() {
	sampleContract := new(HelloWorldContract)
	cc, err := contractapi.NewChaincode(sampleContract)
	if err != nil {
		fmt.Printf("Error creating HelloWorld chaincode: %s", err.Error())
		return
	}

	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting HelloWorld chaincode: %s", err.Error())
	}
}
