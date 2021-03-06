package models

import (
	"encoding/json"
	//"errors"
	"github.com/abbeydabiri/hdl-chaincode-ibm/utils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//LoanBuyerBase data structure
type LoanBuyerBase struct {
	LoanBuyerID   int       `json:"loanbuyerID"`
	LoanBuyerName string    `json:"loanbuyername"`
	BuyerCategory string    `json:"buyercategory"`
	AdminUserID   int       `json:"adminuserID"`
	HqAddress     string    `json:"hqaddress"`
	Location      string    `json:"location"`
	LocationLat   string    `json:"locationlat"`
	LocationLong  string    `json:"locationlong"`
	Created       time.Time `json:"created"`
	Createdby     string    `json:"createdby"`
}

//LoanBuyer struct for chain state
type LoanBuyer struct {
	ObjectType  string        `json:"docType"` // default is 'LONBUY'
	LoanBuyerID string        `json:"loanbuyerID"`
	Data        LoanBuyerBase `json:"data"` // composition
}

//PutState Write asset state to ledger
func (asset *LoanBuyer) PutState(stub shim.ChaincodeStubInterface) pb.Response {
	
	// Marshal the struct to []byte
	b, err := json.Marshal(asset)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: asset.LoanBuyerID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}
	// Write key-value to ledger
	err = stub.PutState(asset.LoanBuyerID, b)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.LoanBuyerW, KEY: asset.LoanBuyerID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}

	// Emit transaction event for listeners
	txID := stub.GetTxID()
	stub.SetEvent((asset.LoanBuyerID + utils.LoanBuyerW + txID), nil)
	r := utils.Response{Code: utils.CODEALLAOK, Message: asset.LoanBuyerID, Payload: nil}
	return shim.Success((r.FormatResponse()))
}

//GetState Read asset state from the ledger
func (asset *LoanBuyer) GetState(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAsset(stub, asset.LoanBuyerID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}



// GetHistory Read asset history from the ledger
func (asset *LoanBuyer) GetHistory(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAssetHistory(stub, asset.LoanBuyerID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}