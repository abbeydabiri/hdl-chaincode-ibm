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

func TestRoleWR(t *testing.T) {
	fmt.Println("Entering TestRoleLogic")

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
		[][]byte{[]byte(utils.RoleW),
			[]byte("1"),
			[]byte("Test User Category"),
			[]byte("Test User Type"),
			[]byte("Test Role Name"),
			[]byte("Test Role Desc"),
		})
	// assert.EqualValues(shim.OK, writeResp.GetStatus(), writeResp.GetMessage())

	if shim.OK != writeResp.GetStatus() {

		println(writeResp.GetMessage())

		return
	}

	testID := "1"
	readResp := stub.MockInvoke(uid,
		[][]byte{[]byte(utils.RoleR),
			[]byte(testID),
		})
	// assert.EqualValues(shim.OK, readResp.GetStatus(), readResp.GetMessage())

	if shim.OK != readResp.GetStatus() {

		println(readResp.GetMessage())

		return
	}

	var ccResp struct {
		Code    string      `json:"code"`
		Role    string      `json:"message"`
		Payload models.Role `json:"payload"`
	}
	if err := json.Unmarshal(readResp.GetPayload(), &ccResp); err != nil {
		panic(err)
	}

	// assert.Equal(testID, ccResp.Payload.RoleID, "Retrieved Role ID mismatch")
	// assert.Equal(utils.ROLEOO, ccResp.Payload.ObjectType, "Retrieved Object Type mismatch")
}
