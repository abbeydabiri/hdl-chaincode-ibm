package models

import (
	"encoding/json"
	//"errors"
	"github.com/abbeydabiri/hdl-chaincode-ibm/utils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// PermissionsBase data structure
type PermissionsBase struct {
	PermID     int       `json:"permID"`
	PermRoleID int       `json:"permroleID"`
	PermName   string    `json:"permname"`
	PermModule string    `json:"permmodule"`
	Created    time.Time `json:"created"`
	Createdby  string    `json:"createdby"`
}

//Permissions struct for chain state
type Permissions struct {
	ObjectType string          `json:"docType"` // default is 'PERMSN'
	PermID     string          `json:"permID"`
	Data       PermissionsBase `json:"data"` // composition
}

//PutState Write asset state to ledger
func (asset *Permissions) PutState(stub shim.ChaincodeStubInterface) pb.Response {

	// Marshal the struct to []byte
	b, err := json.Marshal(asset)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.PermissionsW, KEY: asset.PermID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}
	// Write key-value to ledger
	err = stub.PutState(asset.PermID, b)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.PermissionsW, KEY: asset.PermID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}

	// Emit transaction event for listeners
	txID := stub.GetTxID()
	stub.SetEvent((asset.PermID + utils.PermissionsW + txID), nil)
	r := utils.Response{Code: utils.CODEALLAOK, Message: asset.PermID, Payload: nil}
	return shim.Success((r.FormatResponse()))
}

//GetState Read asset state from the ledger
func (asset *Permissions) GetState(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAsset(stub, asset.PermID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}


// GetHistory Read asset history from the ledger
func (asset *Permissions) GetHistory(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAssetHistory(stub, asset.PermID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}