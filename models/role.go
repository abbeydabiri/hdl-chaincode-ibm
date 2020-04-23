package models

import (
	"encoding/json"
	//"errors"
	"github.com/abbeydabiri/hdl-chaincode-ibm/utils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//RoleBase data structure
type RoleBase struct {
	RoleID       int       `json:"roleID"`
	UserCategory string    `json:"usercategory"`
	UserType     string    `json:"usertype"`
	RoleName     string    `json:"rolename"`
	RoleDesc     string    `json:"roledesc"`
	Created      time.Time `json:"created"`
	Createdby    string    `json:"createdby"`
}

//Role struct for chain state
type Role struct {
	ObjectType string   `json:"docType"` // default is 'ROLEOO'
	RoleID     string   `json:"roleID"`
	Data       RoleBase `json:"data"` // composition
}

//PutState Write asset state to ledger
func (asset *Role) PutState(stub shim.ChaincodeStubInterface) pb.Response {

	// Marshal the struct to []byte
	b, err := json.Marshal(asset)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.RoleW, KEY: asset.RoleID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}
	// Write key-value to ledger
	err = stub.PutState(asset.RoleID, b)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.RoleW, KEY: asset.RoleID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}

	// Emit transaction event for listeners
	txID := stub.GetTxID()
	stub.SetEvent((asset.RoleID + utils.RoleW + txID), nil)
	r := utils.Response{Code: utils.CODEALLAOK, Message: asset.RoleID, Payload: nil}
	return shim.Success((r.FormatResponse()))
}

//GetState Read asset state from the ledger
func (asset *Role) GetState(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAsset(stub, asset.RoleID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}


// GetHistory Read asset history from the ledger
func (asset *Role) GetHistory(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAssetHistory(stub, asset.RoleID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}
