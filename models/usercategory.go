package models

import (
	"encoding/json"
	//"errors"
	"github.com/abbeydabiri/hdl-chaincode-ibm/utils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//UserCategoryBase data structure
type UserCategoryBase struct {
	CatID          int       `json:"catID"`
	CatName        string    `json:"catname"`
	CatDescription string    `json:"catdescription"`
	Created        time.Time `json:"created"`
	Createdby      string    `json:"createdby"`
}

//UserCategory struct for chain state
type UserCategory struct {
	ObjectType string           `json:"docType"` // default is 'USECAT'
	CatID      string           `json:"catID"`
	Data       UserCategoryBase `json:"data"` // composition
}

//PutState Write asset state to ledger
func (asset *UserCategory) PutState(stub shim.ChaincodeStubInterface) pb.Response {
	
	// Marshal the struct to []byte
	b, err := json.Marshal(asset)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.UserCategoryW, KEY: asset.CatID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}
	// Write key-value to ledger
	err = stub.PutState(asset.CatID, b)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.UserCategoryW, KEY: asset.CatID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}

	// Emit transaction event for listeners
	txID := stub.GetTxID()
	stub.SetEvent((asset.CatID + utils.UserCategoryW + txID), nil)
	r := utils.Response{Code: utils.CODEALLAOK, Message: asset.CatID, Payload: nil}
	return shim.Success((r.FormatResponse()))
}

//GetState Read asset state from the ledger
func (asset *UserCategory) GetState(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAsset(stub, asset.CatID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}

// GetHistory Read asset history from the ledger
func (asset *UserCategory) GetHistory(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAssetHistory(stub, asset.CatID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}
