package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
)

// Response struct to have consistent Response structure for all chaincode invokes
type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Payload []byte `json:"payload"`
}

// FormatResponse - present response in json format
func (resp *Response) FormatResponse() []byte {
	var buffer bytes.Buffer
	buffer.WriteString("{\"code\":")
	buffer.WriteString("\"")
	buffer.WriteString(resp.Code)
	buffer.WriteString("\",")
	buffer.WriteString("\"message\":")
	buffer.WriteString("\"")
	buffer.WriteString(resp.Message)
	if resp.Payload == nil {
		buffer.WriteString("\"}")
	} else {
		buffer.WriteString("\",")
		buffer.WriteString("\"payload\":")
		buffer.Write(resp.Payload)
		buffer.WriteString("}")
	}
	return buffer.Bytes()
}

// ChainError - custom error
type ChainError struct {
	FCN  string // function/method
	KEY  string // associated KEY, if any
	CODE string // Error CODE
	ERR  error  // Error.
}

func (e *ChainError) Error() string {
	return e.FCN + " " + e.KEY + " " + e.CODE + ": " + e.ERR.Error()
}

//CheckAsset - check if an asset ( with a key) is already available on the ledger
func CheckAsset(stub shim.ChaincodeStubInterface, assetID string) (bool, *ChainError) {

	assetBytes, err := stub.GetState(assetID)

	if err != nil {
		e := &ChainError{FCN: "CheckAsset", KEY: assetID, CODE: CODEGENEXCEPTION, ERR: err}
		return false, e
	} else if assetBytes != nil {
		e := &ChainError{FCN: "CheckAsset", KEY: assetID, CODE: CODEAlRDEXIST, ERR: errors.New("Asset with key already exists")}
		return true, e
	}
	return false, nil
}

//QueryAsset - return query state from the ledger
func QueryAsset(stub shim.ChaincodeStubInterface, assetID string) ([]byte, *ChainError) {

	assetBytes, err := stub.GetState(assetID)

	if err != nil {
		e := &ChainError{FCN: "QueryAsset", KEY: assetID, CODE: CODEGENEXCEPTION, ERR: err}
		return nil, e
	} else if assetBytes == nil {
		e := &ChainError{FCN: "QueryAsset", KEY: assetID, CODE: CODENOTFOUND, ERR: errors.New("Asset ID not found")}
		return nil, e
	}
	return assetBytes, nil
}

//QueryAssetHistory - return query state from the ledger
func QueryAssetHistory(stub shim.ChaincodeStubInterface, assetID string) ([]byte, *ChainError) {

	iter, err := stub.GetHistoryForKey(assetID)
	if err != nil {
		e := &ChainError{FCN: "QueryAssetHistory", KEY: assetID, CODE: CODENOTFOUND, ERR: err}
        return nil, e
	}
	defer func() { _ = iter.Close() }()

	type historyStruct struct {
		TxID string `json:"txID"`
		Timestamp time.Time  `json:"timestamp"`
		Record []byte  `json:"record"`
	}
	
	var history []historyStruct
	for iter.HasNext() {
		state, err := iter.Next()
		if err != nil {
			e := &ChainError{FCN: "QueryAssetHistory", KEY: assetID, CODE: CODENOTFOUND, ERR: err}
        	return nil, e
		}

		entry := historyStruct{
			TxID:      state.GetTxId(),
			Timestamp: time.Unix(state.GetTimestamp().GetSeconds(), 0),
			Record:     state.Value,
		}

		history = append(history, entry)
	}

	historyBytes, errNew := json.Marshal(history)
	if errNew != nil {
		e := &ChainError{FCN: "QueryAssetHistory", KEY: assetID, CODE: CODENOTFOUND, ERR: errNew}
        return nil, e
	}

	return historyBytes, nil
}

//GetCallerID - reterive caller id from ECert
func GetCallerID(stub shim.ChaincodeStubInterface) (string, *ChainError) {
	id, err := cid.New(stub)
	callerID, err := id.GetID()
	if err != nil {
		e := &ChainError{FCN: "GetCallerID", KEY: "", CODE: CODEGENEXCEPTION, ERR: err}
		return "", e
	}
	// decode the returned base64 string
	data, err := base64.StdEncoding.DecodeString(callerID)
	if err != nil {
		e := &ChainError{FCN: "GetCallerID", KEY: "", CODE: CODEGENEXCEPTION, ERR: err}
		return "", e
	}
	l := strings.Split(string(data), "::")
	return l[1], nil
}
