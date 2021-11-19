package oneid

import (
	requests "git.onespace.co.th/osgolib/http-requests"
	"github.com/labstack/echo/v4"
)

func (c *context) RegisterMobile(mobileNo string, res *ResponseRegisterMobile) Context {
	if err := requests.Call().Post(requests.Params{
		URL: c.apiEndpoint(APIEndpointRegisterMobile),
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

//
func (c *context) RegisterMobileConfirm(mobileConfirm RegisterMobileConfirm, res *ResponseRegisterMobileConfirm) Context {
	if err := requests.Call().Post(requests.Params{
		URL: c.apiEndpoint(APIEndpointRegisterMobileConfirm),
		HEADERS: map[string]string{
			echo.HeaderContentType: "application/json",
		},
		BODY: c.jsonBody(map[string]interface{}{
			"client_id":     c.clientId,
			"client_secret": c.clientSecret,
			"refcode":       c.refCode,
			"mobile_ref_id": mobileConfirm.MobileRefID,
			"otp":           mobileConfirm.OTP,
			"flg_term":      mobileConfirm.FlgTerm,
		}),
		TIMEOUT: 30,
	}, &resRequest).Error(); err != nil {
		c.err = err
		return c
	}
	return c.handleRes(resRequest, &res)
}
