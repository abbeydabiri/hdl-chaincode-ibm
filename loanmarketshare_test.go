package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/abbeydabiri/hdl-chaincode-ibm/models"
	"github.com/abbeydabiri/hdl-chaincode-ibm/utils"

	// "github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	// "github.com/stretchr/testify/assert"
)

func TestLoanMarketShareWR(t *testing.T) {
	fmt.Println("Entering TestLoanMarketShareLogic")

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
		[][]byte{[]byte(utils.LoanMarketShareW),
			[]byte("1"),
			[]byte("Test Title Holder"),
			[]byte("99.99"),
			[]byte("99.99"),
			[]byte("Test Statutes"),
			[]byte("99.99"),
			[]byte("Test Status"),
		})
	// assert.EqualValues(shim.OK, writeResp.GetStatus(), writeResp.GetMessage())

	if shim.OK != writeResp.GetStatus() {

		println(writeResp.GetMessage())

		return
	}

	testID := "1"
	readResp := stub.MockInvoke(uid,
		[][]byte{[]byte(utils.LoanMarketShareR),
			[]byte(testID),
		})
	// assert.EqualValues(shim.OK, readResp.GetStatus(), readResp.GetMessage())

	if shim.OK != readResp.GetStatus() {

		println(readResp.GetMessage())

		return
	}

	var ccResp struct {
		Code    string                 `json:"code"`
		Message string                 `json:"message"`
		Payload models.LoanMarketShare `json:"payload"`
	}
	if err := json.Unmarshal(readResp.GetPayload(), &ccResp); err != nil {
		panic(err)
	}

	// assert.Equal(testID, ccResp.Payload.ShareID, "Retrieved Share ID mismatch")
	// assert.Equal(utils.LONMRK, ccResp.Payload.ObjectType, "Retrieved Object Type mismatch")
}
