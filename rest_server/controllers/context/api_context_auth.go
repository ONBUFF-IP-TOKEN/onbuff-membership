package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/onbuff-membership/rest_server/controllers/resultcode"
)

// member
type Member struct {
	Id         int64  `json:"id" validate:"required"`
	WalletAddr string `json:"wallet_address" validate:"required"`
	Email      string `json:"email" validate:"required"`
	WalletType string `json:"wallet_type" validate:"required"`
	CreateTs   int64  `json:"create_ts" validate:"required"`
}

func NewMember() *Member {
	return new(Member)
}

/////////////////////////

// register
type RegisterMember struct {
	WalletType string    `json:"wallet_type" validate:"required"`
	WalletAuth LoginAuth `json:"wallet_auth" validate:"required"`
	Email      string    `json:"email" validate:"required"`
}

func NewRegisterMember() *RegisterMember {
	return new(RegisterMember)
}

func (o *RegisterMember) CheckValidate() *base.BaseResponse {
	if len(o.WalletType) == 0 && (Wallet_type_metamask != o.WalletType) {
		return base.MakeBaseResponse(resultcode.Result_Auth_InvalidWalletType)
	}
	if len(o.WalletAuth.WalletAddr) == 0 {
		return base.MakeBaseResponse(resultcode.Result_RequireWalletAddress)
	}
	if len(o.WalletAuth.Message) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Auth_RequireMessage)
	}
	if len(o.WalletAuth.Sign) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Auth_RequireSign)
	}
	if len(o.Email) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Auth_RequireEmailInfo)
	}
	return nil
}

/////////////////////////

///////////// API ///////////////////////////////
// login
type LoginAuth struct {
	WalletAddr string `json:"wallet_address" validate:"required"`
	Message    string `json:"message" validate:"required"`
	Sign       string `json:"sign" validate:"required"`
}
type LoginParam struct {
	WalletType string    `json:"wallet_type" validate:"required"`
	WalletAuth LoginAuth `json:"wallet_auth" validate:"required"`
}

func NewLoginParam() *LoginParam {
	return new(LoginParam)
}

func (o *LoginParam) CheckValidate() *base.BaseResponse {
	if len(o.WalletType) == 0 && (Wallet_type_metamask != o.WalletType) {
		return base.MakeBaseResponse(resultcode.Result_Auth_InvalidWalletType)
	}
	if len(o.WalletAuth.WalletAddr) == 0 {
		return base.MakeBaseResponse(resultcode.Result_RequireWalletAddress)
	}
	if len(o.WalletAuth.Message) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Auth_RequireMessage)
	}
	if len(o.WalletAuth.Sign) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Auth_RequireSign)
	}
	return nil
}

type LoginResponse struct {
	AuthToken  string `json:"auth_token" validate:"required"`
	ExpireDate int64  `json:"expire_date" validate:"required"`
}

/////////////////////////

// verify auth token
type VerifyAuthToken struct {
	WalletAddr string `json:"wallet_address" validate:"required"`
	AuthToken  string `json:"auth_token" validate:"required"`
}

func NewVerifyAuthToken() *VerifyAuthToken {
	return new(VerifyAuthToken)
}

func (o *VerifyAuthToken) CheckValidate() *base.BaseResponse {
	if len(o.WalletAddr) == 0 {
		return base.MakeBaseResponse(resultcode.Result_RequireWalletAddress)
	}
	if len(o.AuthToken) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Auth_RequireAuthToken)
	}
	return nil
}

/////////////////////////
