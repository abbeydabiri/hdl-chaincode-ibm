package models

import (
	"encoding/json"
	//"errors"
	"github.com/abbeydabiri/hdl-chaincode-ibm/utils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//TransactionBase data structure
type TransactionBase struct {
	TxnID        int       `json:"txnID"`
	TxnDate      string    `json:"txndate"`
	BuyerID      int       `json:"buyerID"`
	UserID       int       `json:"userID"`
	Repayment    float64   `json:"repayment"`
	Amount       float64   `json:"amount"`
	InterestRate float64   `json:"interestrate"`
	Outstanding  float64   `json:"outstanding"`
	DueDate      string    `json:"duedate"`
	Bank         string    `json:"bank"`
	LoanStatus   string    `json:"loanstatus"`
	Created      time.Time `json:"created"`
	Createdby    string    `json:"createdby"`
}

//Transactions struct for chain state
type Transactions struct {
	ObjectType string          `json:"docType"` // default is 'TRNSAC'
	TxnID      string          `json:"txnID"`
	Data       TransactionBase `json:"data"` // composition
}

//PutState Write asset state to ledger
func (asset *Transactions) PutState(stub shim.ChaincodeStubInterface) pb.Response {
	
	// Marshal the struct to []byte
	b, err := json.Marshal(asset)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.TransactionW, KEY: asset.TxnID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}
	// Write key-value to ledger
	err = stub.PutState(asset.TxnID, b)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.TransactionW, KEY: asset.TxnID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}

	// Emit transaction event for listeners
	txID := stub.GetTxID()
	stub.SetEvent((asset.TxnID + utils.TransactionW + txID), nil)
	r := utils.Response{Code: utils.CODEALLAOK, Message: asset.TxnID, Payload: nil}
	return shim.Success((r.FormatResponse()))
}

//GetState Read asset state from the ledger
func (asset *Transactions) GetState(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAsset(stub, asset.TxnID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}


// GetHistory Read asset history from the ledger
func (asset *Transactions) GetHistory(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAssetHistory(stub, asset.TxnID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}
