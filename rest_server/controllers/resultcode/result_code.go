package resultcode

const (
	Result_Success                = 0
	Result_RequireWalletAddress   = 12000
	Result_RequireTitle           = 12001
	Result_RequireTokenType       = 12002
	Result_RequireThumbnailUrl    = 12003
	Result_RequireValidTokenPrice = 12004
	Result_RequireValidExpireDate = 12005
	Result_RequireCreator         = 12006
	Result_RequireValidItemId     = 12007
	Result_RequireValidPageOffset = 12008
	Result_RequireValidPageSize   = 12009
	Result_RequireDescription     = 12010
	Result_InvalidWalletAddress   = 12011
	Result_RequiredPurchaseTxHash = 12012
	Result_RequireEmailInfo       = 12013

	Result_Product_RequiredTitle        = 12500
	Result_Product_RequiredThumbnailUrl = 12501
	Result_Product_RequiredProductType  = 12502
	Result_Product_RequiredTokenType    = 12503
	Result_Product_RequiredCreator      = 12504
	Result_Product_RequiredDesc         = 12505
	Result_Product_RequireQuantityTotal = 12506
	Result_Product_RequireVaildId       = 12507
	Result_Product_RequireValidState    = 12508
	Result_Product_NotOnSale            = 12509
	Result_Product_LackOfQuantity       = 12510

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
)

var ResultCodeText = map[int]string{
	Result_Success:                "success",
	Result_RequireWalletAddress:   "Wallet address is required",
	Result_RequireTitle:           "Item title is required",
	Result_RequireTokenType:       "Token type is required",
	Result_RequireThumbnailUrl:    "Thumbnail url is required",
	Result_RequireValidTokenPrice: "Valid token price is required",
	Result_RequireValidExpireDate: "Valid expire date is required",
	Result_RequireCreator:         "Creator is required",
	Result_RequireValidItemId:     "Valid item id is required",
	Result_RequireValidPageOffset: "Valid page offset is required",
	Result_RequireValidPageSize:   "Valid page size is required",
	Result_RequireDescription:     "Description is required",
	Result_InvalidWalletAddress:   "Invalid Wallet Address",
	Result_RequiredPurchaseTxHash: "Require purchase tx hash info",
	Result_RequireEmailInfo:       "Require email info",

	Result_Product_RequiredTitle:        "Require product title",
	Result_Product_RequiredThumbnailUrl: "Require product thumbnail url",
	Result_Product_RequiredProductType:  "Require product type info",
	Result_Product_RequiredTokenType:    "Require token type",
	Result_Product_RequiredCreator:      "Require creator info",
	Result_Product_RequiredDesc:         "Require description info",
	Result_Product_RequireQuantityTotal: "Require total quantity",
	Result_Product_RequireVaildId:       "Require valid product id",
	Result_Product_RequireValidState:    "Require valid product state",
	Result_Product_NotOnSale:            "Not on sale",
	Result_Product_LackOfQuantity:       "Lack of Quantity",

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
}
