package main

import (
	"encoding/json"
	"fmt"
	"testing"
	// "time"

	"github.com/abbeydabiri/hdl-chaincode-ibm/models"
	"github.com/abbeydabiri/hdl-chaincode-ibm/utils"

	// "github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	// "github.com/stretchr/testify/assert"
)

func TestBankWR(t *testing.T) {
	fmt.Println("Entering TestBankLogic")

	// assert := assert.New(t)
	// Instantiate mockStub using HDLChaincode as the target chaincode to unit test
	stub := shim.NewMockStub("TestStub", new(HDLChaincode))
	//Verify stub is available
	// assert.NotNil\(stub, "Stub is nil, Test stub creation failed"\)

	if stub == nil {

		println("Stub is nil, Test stub creation failed")

		return
	}

	uid := "test-uid"

	writeResp := stub.MockInvoke(uid,
		[][]byte{[]byte(utils.BankW),
			[]byte("2"),
			[]byte("Test Bank 2"),
			[]byte("Test Hq 2"),
			[]byte("Test BankCategory 2"),
			[]byte("2"),
			[]byte("Test Location 2"),
			[]byte("Test LocationLat 2"),
			[]byte("Test LocationLong 2"),
		})
	// assert.EqualValues(shim.OK, writeResp.GetStatus(), "logic.BankW write test Bank to state failed.")
	if shim.OK != writeResp.GetStatus() {

		println(writeResp.GetMessage())

		return
	}

	testID := "2"
	readResp := stub.MockInvoke(uid,
		[][]byte{[]byte(utils.BankR),
			[]byte(testID),
		})
	// assert.EqualValues(shim.OK, readResp.GetStatus(), "logic.BankR read test Bank from state failed.")

	if shim.OK != readResp.GetStatus() {

		println(readResp.GetMessage())

		return
	}

	var ccResp struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Payload models.Bank `json:"payload"`
	}
	if err := json.Unmarshal(readResp.GetPayload(), &ccResp); err != nil {
		panic(err)
	}

	// assert.Equal(testID, ccResp.Payload.BankID, "Retrieved Bank ID mismatch")
	// assert.Equal(utils.BANKOO, ccResp.Payload.ObjectType, "Retrieved Object Type mismatch")
}
