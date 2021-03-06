package oneid

import (
	"github.com/labstack/echo/v4"
	requests "github.com/onespacegolib/http-requests"
)

func (c *context) GetAccount(account *ResponseGetAccount) Context {
	if err := requests.Call().Get(requests.Params{
		URL:  c.apiEndpoint(APIEndpointGetAccount),
		BODY: nil,
		HEADERS: map[string]string{
			echo.HeaderContentType:   "application/json",
			echo.HeaderAuthorization: c.bearer,
		},
		TIMEOUT: 5,
	}, &resRequest).Error(); err != nil {
		c.err = err
		return c
	}
	return c.handleRes(resRequest, &account)
}

func (c *context) GetServiceCitizenInfo(info *ResponseServiceCitizenInfo) Context {

	if err := requests.Call().Get(requests.Params{
		URL:  c.apiEndpoint(APIEndpointServiceCitizenInfo),
		BODY: nil,
		HEADERS: map[string]string{
			echo.HeaderAccept:        "application/json",
			echo.HeaderAuthorization: c.bearer,
		},
		TIMEOUT: 5,
	}, &resRequest).Error(); err != nil {
		c.err = err
		return c
	}
	return c.handleRes(resRequest, &info)
}
