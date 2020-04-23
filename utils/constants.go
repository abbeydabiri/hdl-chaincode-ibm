package utils

// Constants
const (
	// Error codes
	CODEALLAOK              string = "000" // Success
	CODENOTFOUND            string = "101" // resource not found
	CODEUNKNOWNINVOKE       string = "102" // Unknown invoke
	CODEUNPROCESSABLEENTITY string = "103" // Invalid input
	CODEGENEXCEPTION        string = "201" // Unknown exception
	CODEAlRDEXIST           string = "202" // Not unique
	CODENOTALLWD            string = "104" // Operation not allowed

	//Asset Object/Doctypes
	BANKOO string = "BANKOO" //bitmask is 00
	BNKBRA string = "BNKBRA" //bitmask is 01
	BNKCFG string = "BNKCFG" //bitmask is 02
	BUYERO string = "BUYERO" //bitmask is 03
	LOANOO string = "LOANOO" //bitmask is 04
	LONBUY string = "LONBUY" //bitmask is 05
	LONDOC string = "LONDOC" //bitmask is 06
	LONMRK string = "LONMRK" //bitmask is 07
	LONRAT string = "LONRAT" //bitmask is 08
	MESSAG string = "MESSAG" //bitmask is 09
	PERMSN string = "PERMSN" //bitmask is 10
	PRPRTY string = "PRPRTY" //bitmask is 11
	ROLEOO string = "ROLEOO" //bitmask is 12
	SELLER string = "SELLER" //bitmask is 13
	TRNSAC string = "TRNSAC" //bitmask is 14
	USEROO string = "USEROO" //bitmask is 15
	USECAT string = "USECAT" //bitmask is 16

	// Range index name - to perform range queries
	INDXNM string = "bitmask~txnID~id"

	FIXEDPT int32 = 4 // All currency values rounded off to 4 decimals i.e. 0.0000

	//function names for read and write

	Init string = "Init"

	BankR string = "BankR"
	BankH string = "BankH"
	BankW string = "BankW"

	BankBranchR string = "BankBranchR"
	BankBranchH string = "BankBranchH"
	BankBranchW string = "BankBranchW"

	BranchConfigR string = "BranchConfigR"
	BranchConfigH string = "BranchConfigH"
	BranchConfigW string = "BranchConfigW"

	BuyerR string = "BuyerR"
	BuyerH string = "BuyerH"
	BuyerW string = "BuyerW"

	LoanR string = "LoanR"
	LoanH string = "LoanH"
	LoanW string = "LoanW"

	LoanBuyerR string = "LoanBuyerR"
	LoanBuyerH string = "LoanBuyerH"
	LoanBuyerW string = "LoanBuyerW"

	LoanDocR string = "LoanDocR"
	LoanDocH string = "LoanDocH"
	LoanDocW string = "LoanDocW"

	LoanMarketShareR string = "LoanMarketShareR"
	LoanMarketShareH string = "LoanMarketShareH"
	LoanMarketShareW string = "LoanMarketShareW"

	LoanRatingR string = "LoanRatingR"
	LoanRatingH string = "LoanRatingH"
	LoanRatingW string = "LoanRatingW"

	MessageR string = "MessageR"
	MessageH string = "MessageH"
	MessageW string = "MessageW"

	PermissionsR string = "PermissionsR"
	PermissionsH string = "PermissionsH"
	PermissionsW string = "PermissionsW"

	PropertyR string = "PropertyR"
	PropertyH string = "PropertyH"
	PropertyW string = "PropertyW"

	RoleR string = "RoleR"
	RoleH string = "RoleH"
	RoleW string = "RoleW"

	SellerR string = "SellerR"
	SellerH string = "SellerH"
	SellerW string = "SellerW"

	TransactionR string = "TransactionR"
	TransactionH string = "TransactionH"
	TransactionW string = "TransactionW"

	UserR string = "UserR"
	UserH string = "UserH"
	UserW string = "UserW"

	UserCategoryR string = "UserCategoryR"
	UserCategoryH string = "UserCategoryH"
	UserCategoryW string = "UserCategoryW"
)
