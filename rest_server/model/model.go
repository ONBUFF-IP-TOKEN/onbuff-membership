package model

import "github.com/ONBUFF-IP-TOKEN/onbuff-membership/rest_server/controllers/context"

type AuthInfo struct {
	AuthToken  string            `json:"auth_token"`
	ExpireDate int64             `json:"expire_date"`
	WalletAuth context.LoginAuth `json:"wallet_auth"`

	Email          string                 `json:"email" validate:"required"`
	NickName       string                 `json:"nickname" validate:"required"`
	ProfileImg     string                 `json:"profile_img" validate:"required"`
	TermsOfService context.TermsOfService `json:"terms_of_service"`
}
