package commonapi

import (
	"net/http"
	"strings"

	"github.com/ONBUFF-IP-TOKEN/baseapp/base"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/onbuff-membership/rest_server/controllers/auth"
	"github.com/ONBUFF-IP-TOKEN/onbuff-membership/rest_server/controllers/context"
	"github.com/ONBUFF-IP-TOKEN/onbuff-membership/rest_server/controllers/resultcode"
	"github.com/ONBUFF-IP-TOKEN/onbuff-membership/rest_server/model"
	"github.com/ONBUFF-IP-TOKEN/onbuff-membership/rest_server/token"
	"github.com/labstack/echo"
)

func PostLogin(c echo.Context) error {
	params := context.NewLoginParam()
	if err := c.Bind(params); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}

	if err := params.CheckValidate(); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	resp := new(base.BaseResponse)
	// 1. verify sign check
	if !token.VerifySign(params.WalletAuth.WalletAddr, params.WalletAuth.Message, params.WalletAuth.Sign) {
		// invalid sign info
		resp.SetReturn(resultcode.Result_Auth_InvalidLoginInfo)
		return c.JSON(http.StatusOK, resp)
	}

	//1.1 가입정보 존재 확인
	member, err := model.GetDB().GetExistMember(params.WalletAuth.WalletAddr)
	if err != nil {
		resp.SetReturn(resultcode.Result_DBError)
		return c.JSON(http.StatusOK, resp)
	}
	if len(member.Email) == 0 && len(member.WalletAddr) == 0 {
		log.Error("PostLogin not member : ", member.WalletAddr)
		resp.SetReturn(resultcode.Result_Auth_NotMember)
		return c.JSON(http.StatusOK, resp)
	}
	if member.ActivateState == context.Member_Activate_State_Blocked {
		log.Error("PostLogin blocked member : ", member.WalletAddr)
		resp.SetReturn(resultcode.Result_Auth_BlockedMember)
		return c.JSON(http.StatusOK, resp)
	}
	if member.ActivateState == context.Member_Activate_State_Withdraw {
		log.Error("PostLogin withdraw member : ", member.WalletAddr)
		resp.SetReturn(resultcode.Result_Auth_Withdraw)
		return c.JSON(http.StatusOK, resp)
	}

	// 2. redis duplicate check
	if authInfo, err := model.GetDB().GetAuthInfo(params.WalletAuth.WalletAddr); err == nil {
		// redis에 기존 정보가 있다면 기존에 발급된 토큰으로 응답한다.
		resp.Success()
		resp.Value = context.LoginResponse{
			AuthToken:  authInfo.AuthToken,
			ExpireDate: authInfo.ExpireDate,

			Email:      authInfo.Email,
			NickName:   authInfo.NickName,
			ProfileImg: authInfo.ProfileImg,
		}
	} else {
		// 3. create auth token
		authToken, expireDate, err := auth.GetIAuth().EncryptJwt(params.WalletAuth.WalletAddr)
		if err != nil {
			resp.SetReturn(resultcode.Result_Auth_DontEncryptJwt)
		} else {
			resp.Success()
			resp.Value = context.LoginResponse{
				AuthToken:  authToken,
				ExpireDate: expireDate,

				Email:      member.Email,
				NickName:   member.NickName,
				ProfileImg: member.ProfileImg,
			}
			// 3. redis save
			authInfo := model.AuthInfo{
				AuthToken:  authToken,
				ExpireDate: expireDate,
				WalletAuth: params.WalletAuth,

				Email:      member.Email,
				NickName:   member.NickName,
				ProfileImg: member.ProfileImg,
			}
			err = model.GetDB().SetAuthInfo(&authInfo)
			if err != nil {
				return base.BaseJSONInternalServerError(c, err)
			}
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func PostMemberRegister(c echo.Context) error {
	params := context.NewRegisterMember()
	if err := c.Bind(params); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}

	if err := params.CheckValidate(); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	resp := new(base.BaseResponse)
	// 1. verify sign check
	if !token.VerifySign(params.WalletAuth.WalletAddr, params.WalletAuth.Message, params.WalletAuth.Sign) {
		// invalid sign info
		resp.SetReturn(resultcode.Result_Auth_InvalidLoginInfo)
		return c.JSON(http.StatusOK, resp)
	}
	//2. 가입정보 존재 확인
	member, err := model.GetDB().GetExistMember(params.WalletAuth.WalletAddr)
	if err != nil {
		resp.SetReturn(resultcode.Result_DBError)
		return c.JSON(http.StatusOK, resp)
	}
	if len(member.Email) != 0 || len(member.WalletAddr) != 0 {
		resp.SetReturn(resultcode.Result_Auth_ExistMember)
		return c.JSON(http.StatusOK, resp)
	}
	// 계정 활성화
	params.ActivateState = context.Member_Activate_State_Normal

	if _, err := model.GetDB().InsertMember(params); err != nil {
		resp.SetReturn(resultcode.Result_DBError)
		return c.JSON(http.StatusOK, resp)
	}

	resp.Success()
	return c.JSON(http.StatusOK, resp)
}

// auth 정보 정상 확인
func PostVerifyAuthToken(c echo.Context) error {
	params := context.NewVerifyAuthToken()
	if err := c.Bind(params); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}

	if err := params.CheckValidate(); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	resp := new(base.BaseResponse)
	// 1. verify sign check
	walletAddr, isValid := auth.GetIAuth().IsValidAuthToken(params.AuthToken)
	if !isValid {
		// auth token 오류 리턴
		log.Error("VerifyAuthToken invalid jwt : ", params.AuthToken, " walletaddr:", params.WalletAddr)
		resp.SetReturn(resultcode.Result_Auth_InvalidJwt)
		return c.JSON(http.StatusOK, resp)
	}

	// 2. 주소 일치 확인
	if !strings.EqualFold(*walletAddr, params.WalletAddr) {
		log.Error("VerifyAuthToken not equal walletaddr :", walletAddr, " : ", params.WalletAddr)
		resp.SetReturn(resultcode.Result_Auth_InvalidJwt)
		return c.JSON(http.StatusOK, resp)
	}

	// 3. 가입정보 존재 확인
	member, err := model.GetDB().GetExistMember(params.WalletAddr)
	if err != nil {
		resp.SetReturn(resultcode.Result_DBError)
		return c.JSON(http.StatusOK, resp)
	}
	if len(member.Email) == 0 && len(member.WalletAddr) == 0 {
		log.Error("VerifyAuthToken not member : ", params.WalletAddr)
		resp.SetReturn(resultcode.Result_Auth_NotMember)
		return c.JSON(http.StatusOK, resp)
	}
	if member.ActivateState == context.Member_Activate_State_Blocked {
		log.Error("VerifyAuthToken blocked member : ", params.WalletAddr)
		resp.SetReturn(resultcode.Result_Auth_BlockedMember)
		return c.JSON(http.StatusOK, resp)
	}
	if member.ActivateState == context.Member_Activate_State_Withdraw {
		log.Error("VerifyAuthToken withdraw member : ", params.WalletAddr)
		resp.SetReturn(resultcode.Result_Auth_Withdraw)
		return c.JSON(http.StatusOK, resp)
	}

	resp.Success()
	resp.Value = member
	return c.JSON(http.StatusOK, resp)
}

func PutMemberUpdate(c echo.Context) error {
	params := context.NewMember()
	if err := c.Bind(params); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}

	resp := new(base.BaseResponse)
	//1. 가입정보 존재 확인
	member, err := model.GetDB().GetExistMember(params.WalletAddr)
	if err != nil {
		resp.SetReturn(resultcode.Result_DBError)
		return c.JSON(http.StatusOK, resp)
	}

	if len(member.Email) == 0 || len(member.WalletAddr) == 0 {
		resp.SetReturn(resultcode.Result_Auth_NotMember)
		return c.JSON(http.StatusOK, resp)
	}

	if _, err := model.GetDB().UpdateMember(params); err != nil {
		resp.SetReturn(resultcode.Result_DBError)
		return c.JSON(http.StatusOK, resp)
	}

	resp.Success()
	return c.JSON(http.StatusOK, resp)
}

func GetMemberDuplicateCheck(c echo.Context) error {
	params := context.NewMemberDuplicateCheck()
	if err := c.Bind(params); err != nil {
		log.Error(err)
		return base.BaseJSONInternalServerError(c, err)
	}

	if err := params.CheckValidate(); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	resp := new(base.BaseResponse)
	// 1. 가입정보 존재 확인
	member, err := model.GetDB().GetExistMemberByNickEmail(params.NickName, params.Email)
	if err != nil {
		resp.SetReturn(resultcode.Result_DBError)
		return c.JSON(http.StatusOK, resp)
	}
	if len(member.WalletAddr) != 0 {
		log.Error("GetMemberDuplicateCheck exist member : ", params.Email, "  ", params.NickName)
		resp.SetReturn(resultcode.Result_Auth_ExistMember)
		return c.JSON(http.StatusOK, resp)
	}

	resp.Success()
	return c.JSON(http.StatusOK, resp)
}
