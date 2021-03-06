package context

import (
	"strings"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/onbuff-membership/rest_server/controllers/resultcode"
)

const (
	Member_Activate_State_Normal   = 0
	Member_Activate_State_Blocked  = 1
	Member_Activate_State_Withdraw = 2
)

type TermsOfService struct {
	ServiceAgree string `json:"service_agree"`
	PrivacyAgree string `json:"privacy_agree"`
}

// member
type Member struct {
	Id             int64          `json:"id" validate:"required"`
	WalletAddr     string         `json:"wallet_address" validate:"required"`
	Email          string         `json:"email" validate:"required"`
	WalletType     string         `json:"wallet_type" validate:"required"`
	CreateTs       int64          `json:"create_ts" validate:"required"`
	NickName       string         `json:"nickname" validate:"required"`
	ProfileImg     string         `json:"profile_img" validate:"required"`
	ActivateState  int64          `json:"activate_state" validate:"required"`
	TermsOfService TermsOfService `json:"terms_of_service"`
}

func NewMember() *Member {
	return new(Member)
}

/////////////////////////

// register
type RegisterMember struct {
	WalletType     string         `json:"wallet_type" validate:"required"`
	WalletAuth     LoginAuth      `json:"wallet_auth" validate:"required"`
	Email          string         `json:"email" validate:"required"`
	NickName       string         `json:"nickname" validate:"required"`
	ProfileImg     string         `json:"profile_img" validate:"required"`
	ActivateState  int64          `json:"activate_state" validate:"required"`
	TermsOfService TermsOfService `json:"terms_of_service"`
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
	if !strings.EqualFold(o.TermsOfService.ServiceAgree, "true") {
		return base.MakeBaseResponse(resultcode.Result_Auc_Bid_RequireServiceAgree)
	}
	if !strings.EqualFold(o.TermsOfService.PrivacyAgree, "true") {
		return base.MakeBaseResponse(resultcode.Result_Auc_Bid_RequirePrivacyAgree)
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

	Email          string         `json:"email" validate:"required"`
	NickName       string         `json:"nickname" validate:"required"`
	ProfileImg     string         `json:"profile_img" validate:"required"`
	TermsOfService TermsOfService `json:"terms_of_service"`
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

//withdraw memeber
type MemberWithdraw struct {
	WalletAddr string `query:"wallet_address" validate:"required"`
}

func NewMemberWithdraw() *MemberWithdraw {
	return new(MemberWithdraw)
}

func (o *MemberWithdraw) CheckValidate() *base.BaseResponse {
	if len(o.WalletAddr) == 0 {
		return base.MakeBaseResponse(resultcode.Result_RequireWalletAddress)
	}
	return nil
}

/////////////////////////

//remove memeber
type MemberRemove struct {
	WalletAddr string `query:"wallet_address" validate:"required"`
}

func NewMemberRemove() *MemberRemove {
	return new(MemberRemove)
}

func (o *MemberRemove) CheckValidate() *base.BaseResponse {
	if len(o.WalletAddr) == 0 {
		return base.MakeBaseResponse(resultcode.Result_RequireWalletAddress)
	}
	return nil
}

/////////////////////////

//get memeber list
type MemberList struct {
	PageInfo
}

func NewMemberList() *MemberList {
	return new(MemberList)
}

func (o *MemberList) CheckValidate() *base.BaseResponse {
	if o.PageOffset < 0 {
		return base.MakeBaseResponse(resultcode.Result_RequireValidPageOffset)
	}
	if o.PageSize <= 0 {
		return base.MakeBaseResponse(resultcode.Result_RequireValidPageSize)
	}
	return nil
}

type ResponseMemberList struct {
	PageInfo PageInfoResponse `json:"page_info"`
	Members  []Member         `json:"members"`
}

/////////////////////////
