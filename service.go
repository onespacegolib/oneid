package oneid

import (
	"github.com/labstack/echo/v4"
	requests "github.com/onespacegolib/http-requests"
)

func (c *context) GetSharedToken(token *ResponseGetSharedToken) Context {
	if err := requests.Call().Get(requests.Params{
		URL:  c.apiEndpoint(APIEndpointServiceSharedToken),
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
	return c.handleRes(resRequest, &token)
}
