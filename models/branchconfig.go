package models

import (
	"encoding/json"
	//"errors"
	"github.com/abbeydabiri/hdl-chaincode-ibm/utils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// BranchConfigBase data structure
type BranchConfigBase struct {
	ConfigID   int       `json:"configID"`
	BankID     int       `json:"bankID"`
	ConfigName string    `json:"configname"`
	ConfigDesc string    `json:"configdesc"`
	Item       string    `json:"item"`
	Value      string    `json:"value"`
	Created    time.Time `json:"created"`
	Createdby  string    `json:"createdby"`
}

//BranchConfig struct for chain state
type BranchConfig struct {
	ObjectType string           `json:"docType"` // default is 'BNKCFG'
	ConfigID   string           `json:"configID"`
	Data       BranchConfigBase `json:"data"` // composition
}

//PutState Write asset state to ledger
func (asset *BranchConfig) PutState(stub shim.ChaincodeStubInterface) pb.Response {
	
	// Marshal the struct to []byte
	b, err := json.Marshal(asset)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.BranchConfigW, KEY: asset.ConfigID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}
	// Write key-value to ledger
	err = stub.PutState(asset.ConfigID, b)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.BranchConfigW, KEY: asset.ConfigID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}

	// Emit transaction event for listeners
	txID := stub.GetTxID()
	stub.SetEvent((asset.ConfigID + utils.BranchConfigW + txID), nil)
	r := utils.Response{Code: utils.CODEALLAOK, Message: asset.ConfigID, Payload: nil}
	return shim.Success((r.FormatResponse()))
}

//GetState Read asset state from the ledger
func (asset *BranchConfig) GetState(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAsset(stub, asset.ConfigID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}


// GetHistory Read asset history from the ledger
func (asset *BranchConfig) GetHistory(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAssetHistory(stub, asset.ConfigID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}