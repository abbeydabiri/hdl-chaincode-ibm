package models

import (
	"encoding/json"
	//"errors"
	"github.com/abbeydabiri/hdl-chaincode-ibm/utils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//LoanDoc data structure
type LoanDocBase struct {
	DocID     int       `json:"docID"`
	LoanID    int       `json:"loanID"`
	DocName   string    `json:"docname"`
	DocDesc   string    `json:"docdesc"`
	DocLink   string    `json:"doclink"`
	Created   time.Time `json:"created"`
	Createdby string    `json:"createdby"`
}

//LoanDoc struct for chain state
type LoanDoc struct {
	ObjectType string      `json:"docType"` // default is 'LONDOC'
	DocID      string      `json:"docID"`
	Data       LoanDocBase `json:"data"` // composition
}

//PutState Write asset state to ledger
func (asset *LoanDoc) PutState(stub shim.ChaincodeStubInterface) pb.Response {

	// Marshal the struct to []byte
	b, err := json.Marshal(asset)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.LoanDocW, KEY: asset.DocID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}
	// Write key-value to ledger
	err = stub.PutState(asset.DocID, b)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.LoanDocW, KEY: asset.DocID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}

	// Emit transaction event for listeners
	txID := stub.GetTxID()
	stub.SetEvent((asset.DocID + utils.LoanDocW + txID), nil)
	r := utils.Response{Code: utils.CODEALLAOK, Message: asset.DocID, Payload: nil}
	return shim.Success((r.FormatResponse()))
}

//GetState Read asset state from the ledger
func (asset *LoanDoc) GetState(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAsset(stub, asset.DocID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}



// GetHistory Read asset history from the ledger
func (asset *LoanDoc) GetHistory(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAssetHistory(stub, asset.DocID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}