package oneid

import (
	"encoding/json"
	"fmt"
	requests "git.onespace.co.th/osgolib/http-requests"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/url"
)

func (c *context) LoginPWD(pwd PWD, res *ResponseLoginPwdOne) Context {
	if err := requests.Call().Post(requests.Params{
		URL: c.apiEndpoint(APIEndpointOAuthPWD),
		HEADERS: map[string]string{
			echo.HeaderContentType: "application/json",
		},
		BODY: c.jsonBody(map[string]interface{}{
			"grant_type":    "password",
			"client_id":     c.clientId,
			"client_secret": c.clientSecret,
			"username":      pwd.Username,
			"password":      pwd.Password,
		}),
		TIMEOUT: 30,
	}, &resRequest).Error(); err != nil {
		c.err = err
		return c
	}
	return c.handleRes(resRequest, &res)
}

func (c *context) LoginWithSharedToken(sharedToken string, pwd *ResponseLoginPwdOne) Context {
	if err := requests.Call().Post(requests.Params{
		URL: c.apiEndpoint(APIEndpointOAuthSharedToken),
		HEADERS: map[string]string{
			echo.HeaderContentType: "application/json",
		},
		BODY: c.jsonBody(map[string]interface{}{
			"client_id":     c.clientId,
			"client_secret": c.clientSecret,
			"shared_token":  sharedToken,
			"refcode":       c.refCode,
		}),
		TIMEOUT: 30,
	}, &resRequest).Error(); err != nil {
		c.err = err
		return c
	}
	return c.handleRes(resRequest, &pwd)
}

func (c *context) LoginWithOneChatSharedToken(oneChatToken string, pwd *ResponseLoginPwdOne) Context {
	token, err := jwt.Parse(oneChatToken, nil)
	if err == nil {
		c.err = err
		return c
	}
	if token == nil {
		c.err = fmt.Errorf("invalid token")
		return c
	}

	var decode struct {
		AccessToken string `json:"access_token"`
	}

	jsonString, _ := json.Marshal(token.Claims)
	_ = json.Unmarshal(jsonString, &decode)

	return c.LoginWithSharedToken(decode.AccessToken, pwd)
}

func (c *context) LoginWithRefreshToken(refreshToken string, pwd *ResponseLoginPwdOne) Context {
	if err := requests.Call().Post(requests.Params{
		URL: c.apiEndpoint(APIEndpointGetRefreshToken),
		HEADERS: map[string]string{
			echo.HeaderContentType: "application/json",
		},
		BODY: c.jsonBody(map[string]interface{}{
			"grant_type":    `refresh_token`,
			"client_id":     c.clientId,
			"client_secret": c.clientSecret,
			"refresh_token": refreshToken,
		}),
		TIMEOUT: 30,
	}, &resRequest).Error(); err != nil {
		c.err = err
		return c
	}
	return c.handleRes(resRequest, &pwd)
}

func (c *context) LoginMobile(mobileNo string, res *ResponseLoginMobile) Context {
	if err := requests.Call().Post(requests.Params{
		URL: c.apiEndpoint(APIEndpointLoginMobile),
		HEADERS: map[string]string{
			echo.HeaderContentType: "application/json",
		},
		BODY: c.jsonBody(map[string]interface{}{
			"client_id":     c.clientId,
			"client_secret": c.clientSecret,
			"refcode":       c.refCode,
			"mobile_no":     mobileNo,
		}),
		TIMEOUT: 30,
	}, &resRequest).Error(); err != nil {
		c.err = err
		return c
	}
	return c.handleRes(resRequest, &res)
}

func (c *context) LoginMobileConfirm(mobileConfirm *LoginMobileConfirm, res *ResponseLoginMobileConfirm) Context {
	if err := requests.Call().Post(requests.Params{
		URL: c.apiEndpoint(APIEndpointLoginMobileConfirm),
		HEADERS: map[string]string{
			echo.HeaderContentType: "application/json",
		},
		BODY: c.jsonBody(map[string]interface{}{
			"client_id":     c.clientId,
			"client_secret": c.clientSecret,
			"refcode":       c.refCode,
			"mobile_no":     mobileConfirm.MobileNo,
			"otp":           mobileConfirm.OTP,
		}),
		TIMEOUT: 30,
	}, &resRequest).Error(); err != nil {
		c.err = err
		return c
	}
	return c.handleRes(resRequest, &res)
}

func (c *context) CheckUserLogin(userLogin string, res *ResponseCheckUserLogin) Context {
	base, _ := url.Parse(c.apiEndpoint(APIEndpointCheckUserLogin))

	query := url.Values{}
	query.Add(`client_id`, c.clientId)
	query.Add(`secret_key`, c.clientSecret)
	query.Add(`user_login`, userLogin)

	base.RawQuery = query.Encode()

	if err := requests.Call().Get(requests.Params{
		URL: base.String(),
		HEADERS: map[string]string{
			echo.HeaderContentType: "application/json",
		},
		BODY:    nil,
		TIMEOUT: 30,
	}, &resRequest).Error(); err != nil {
		c.err = err
		return c
	}
	return c.handleRes(resRequest, &res)
}
