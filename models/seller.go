package models

import (
	"encoding/json"
	//"errors"
	"github.com/abbeydabiri/hdl-chaincode-ibm/utils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//SellerBase data structure
type SellerBase struct {
	SellerID   int       `json:"sellerID"`
	UserID     int       `json:"userID"`
	SellerType string    `json:"sellertype"`
	Details    string    `json:"details"`
	RegDate    string    `json:"regdate"`
	Created    time.Time `json:"created"`
	Createdby  string    `json:"createdby"`
}

//Seller struct for chain state
type Seller struct {
	ObjectType string     `json:"docType"` // default is 'SELLER'
	SellerID   string     `json:"sellerID"`
	Data       SellerBase `json:"data"` // composition
}

//PutState Write asset state to ledger
func (asset *Seller) PutState(stub shim.ChaincodeStubInterface) pb.Response {
	
	// Marshal the struct to []byte
	b, err := json.Marshal(asset)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.SellerW, KEY: asset.SellerID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}
	// Write key-value to ledger
	err = stub.PutState(asset.SellerID, b)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.SellerW, KEY: asset.SellerID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}

	// Emit transaction event for listeners
	txID := stub.GetTxID()
	stub.SetEvent((asset.SellerID + utils.SellerW + txID), nil)
	r := utils.Response{Code: utils.CODEALLAOK, Message: asset.SellerID, Payload: nil}
	return shim.Success((r.FormatResponse()))
}

//GetState Read asset state from the ledger
func (asset *Seller) GetState(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAsset(stub, asset.SellerID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}


// GetHistory Read asset history from the ledger
func (asset *Seller) GetHistory(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAssetHistory(stub, asset.SellerID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}
