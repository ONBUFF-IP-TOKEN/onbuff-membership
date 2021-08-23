package context

import (
	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/onbuff-membership/rest_server/controllers/resultcode"
)

const (
	Member_Activate_State_Normal   = 0
	Member_Activate_State_Blocked  = 1
	Member_Activate_State_Withdraw = 2
)

// member
type Member struct {
	Id            int64  `json:"id" validate:"required"`
	WalletAddr    string `json:"wallet_address" validate:"required"`
	Email         string `json:"email" validate:"required"`
	WalletType    string `json:"wallet_type" validate:"required"`
	CreateTs      int64  `json:"create_ts" validate:"required"`
	NickName      string `json:"nickname" validate:"required"`
	ProfileImg    string `json:"profile_img" validate:"required"`
	ActivateState int64  `json:"activate_state" validate:"required"`
}

func NewMember() *Member {
	return new(Member)
}

/////////////////////////

// register
type RegisterMember struct {
	WalletType    string    `json:"wallet_type" validate:"required"`
	WalletAuth    LoginAuth `json:"wallet_auth" validate:"required"`
	Email         string    `json:"email" validate:"required"`
	NickName      string    `json:"nickname" validate:"required"`
	ProfileImg    string    `json:"profile_img" validate:"required"`
	ActivateState int64     `json:"activate_state" validate:"required"`
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
	if len(o.NickName) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Auth_RequireNickName)
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

	Email      string `json:"email" validate:"required"`
	NickName   string `json:"nickname" validate:"required"`
	ProfileImg string `json:"profile_img" validate:"required"`
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

// nickname, email duplicate check
type MemberDuplicateCheck struct {
	Email    string `query:"email" validate:"required"`
	NickName string `query:"nickname" validate:"required"`
}

func NewMemberDuplicateCheck() *MemberDuplicateCheck {
	return new(MemberDuplicateCheck)
}

func (o *MemberDuplicateCheck) CheckValidate() *base.BaseResponse {
	if len(o.Email) == 0 && len(o.NickName) == 0 {
		return base.MakeBaseResponse(resultcode.Result_Auth_RequireEmailorNickName)
	}
	return nil
}

/////////////////////////
