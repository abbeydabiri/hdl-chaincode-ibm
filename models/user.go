package models

import (
	"encoding/json"
	//"errors"
	"github.com/abbeydabiri/hdl-chaincode-ibm/utils"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//UserBase data structure
type UserBase struct {
	UserID        int       `json:"userID"`
	UserType      string    `json:"usertype"`
	AccountStatus string    `json:"accountstatus"`
	UserCategory  string    `json:"usercategory"`
	FirstName     string    `json:"firstname"`
	LastName      string    `json:"lastname"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Password      string    `json:"password"`
	RegDate       string    `json:"regdate"`
	Created       time.Time `json:"created"`
	Createdby     string    `json:"createdby"`
}

//User struct for chain state
type User struct {
	ObjectType string   `json:"docType"` // default is 'USEROO'
	UserID     string   `json:"userID"`
	Data       UserBase `json:"data"` // composition
}

//PutState Write asset state to ledger
func (asset *User) PutState(stub shim.ChaincodeStubInterface) pb.Response {

	// Marshal the struct to []byte
	b, err := json.Marshal(asset)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: asset.UserID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}
	// Write key-value to ledger
	err = stub.PutState(asset.UserID, b)
	if err != nil {
		cErr := &utils.ChainError{FCN: utils.UserW, KEY: asset.UserID, CODE: utils.CODEGENEXCEPTION, ERR: err}
		return shim.Error(cErr.Error())
	}

	// Emit transaction event for listeners
	txID := stub.GetTxID()
	stub.SetEvent((asset.UserID + utils.UserW + txID), nil)
	r := utils.Response{Code: utils.CODEALLAOK, Message: asset.UserID, Payload: nil}
	return shim.Success((r.FormatResponse()))
}

//GetState Read asset state from the ledger
func (asset *User) GetState(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAsset(stub, asset.UserID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}

// GetHistory Read asset history from the ledger
func (asset *User) GetHistory(stub shim.ChaincodeStubInterface) pb.Response {
	obj, cErr := utils.QueryAssetHistory(stub, asset.UserID)
	if cErr != nil {
		return shim.Error(cErr.Error())
	}
	r := utils.Response{Code: utils.CODEALLAOK, Message: "OK", Payload: obj}
	return shim.Success((r.FormatResponse()))
}
