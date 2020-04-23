package models

import (
	"encoding/json"
	//"errors"
	"github.com/abbeydabiri/hdl-chaincode-ibm/utils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//LoanRatingBase data strucuture
type LoanRatingBase struct {
	RatingID   int       `json:"ratingID"`
	LoanID     int       `json:"loanID"`
	Rating     float64   `json:"rating"`
	RatingDesc string    `json:"ratingdesc"`
	Created    time.Time `json:"created"`
	Createdby  string    `json:"createdby"`
}

// LoanRating struct for chain state
type LoanRating struct {
	ObjectType string         `json:"docType"` // default is 'LONRAT'
	RatingID   string         `json:"ratingID"`
	Data       LoanRatingBase `json:"data"` // composition
}

//PutState Write asset state to ledger
func (asset *LoanRating) PutState(stub shim.ChaincodeStubInterface) pb.Response {

	// Marshal the struct to []byte
	b, err := json.Marshal(asset)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.LoanRatingW, KEY: asset.RatingID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}
	// Write key-value to ledger
	err = stub.PutState(asset.RatingID, b)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.LoanRatingW, KEY: asset.RatingID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}

	// Emit transaction event for listeners
	txID := stub.GetTxID()
	stub.SetEvent((asset.RatingID + utils.LoanRatingW + txID), nil)
	r := utils.Response{Code: utils.CODEALLAOK, Message: asset.RatingID, Payload: nil}
	return shim.Success((r.FormatResponse()))
}

//GetState Read asset state from the ledger
func (asset *LoanRating) GetState(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAsset(stub, asset.RatingID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}


// GetHistory Read asset history from the ledger
func (asset *LoanRating) GetHistory(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAssetHistory(stub, asset.RatingID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}