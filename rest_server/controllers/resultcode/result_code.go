package resultcode

const (
	Result_Success              = 0
	Result_RequireWalletAddress = 12000

	Result_DBError        = 13000
	Result_DBNotExistItem = 13001

	Result_TokenError               = 14000
	Result_TokenERC721CreateError   = 14001
	Result_TokenERC721BurnError     = 14002
	Result_TokenERC721TransferError = 14003

	Result_Auth_RequireMessage    = 20000
	Result_Auth_RequireSign       = 20001
	Result_Auth_InvalidLoginInfo  = 20002
	Result_Auth_DontEncryptJwt    = 20003
	Result_Auth_InvalidJwt        = 20004
	Result_Auth_InvalidWalletType = 20005
	Result_Auth_NotMember         = 20006
	Result_Auth_RequireEmailInfo  = 20007
	Result_Auth_ExistMember       = 20008
	Result_Auth_RequireAuthToken  = 20009
	Result_Auth_RequireNickName   = 20010
	Result_Auth_BlockedMember     = 20011
	Result_Auth_Withdraw          = 20012
)

var ResultCodeText = map[int]string{
	Result_Success:              "success",
	Result_RequireWalletAddress: "Wallet address is required",

	Result_DBError:        "Internal DB error",
	Result_DBNotExistItem: "Not exist item",

	Result_TokenError:               "Internal Token error",
	Result_TokenERC721CreateError:   "ERC721 create error",
	Result_TokenERC721BurnError:     "ERC721 burn error",
	Result_TokenERC721TransferError: "ERC721 transfer error",

	Result_Auth_RequireMessage:    "Message is required",
	Result_Auth_RequireSign:       "Sign info is required",
	Result_Auth_InvalidLoginInfo:  "Invalid login info",
	Result_Auth_DontEncryptJwt:    "Auth token create fail",
	Result_Auth_InvalidJwt:        "Invalid jwt token",
	Result_Auth_InvalidWalletType: "Invalid wallet type",
	Result_Auth_NotMember:         "Not member",
	Result_Auth_RequireEmailInfo:  "Required email",
	Result_Auth_ExistMember:       "Exist member",
	Result_Auth_RequireAuthToken:  "Required auth token",
	Result_Auth_RequireNickName:   "Required nickname",
	Result_Auth_BlockedMember:     "Blocked Member",
	Result_Auth_Withdraw:          "Withdraw member",
}
