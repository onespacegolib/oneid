package oneid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	requests "github.com/onespacegolib/http-requests"
	"strconv"
)

var (
	resRequest requests.Response
)

type (
	BearerContext interface {
		GetAccount(*ResponseGetAccount) Context

		GetSharedToken(token *ResponseGetSharedToken) Context

		GetServiceCitizenInfo(*ResponseServiceCitizenInfo) Context

		Business() BusinessContext
	}

	BusinessContext interface {
		All(businesses *ResponseBusinesses) Context
		Departments(bizId string, departments *ResponseDepartments) Context
		DepartmentsByTaxId(taxId string, departments *ResponseDepartments) Context
		Department(deptId string, bizId string, department *ResponseDepartment) Context
		Roles(bizId string, roles *ResponseRoles) Context
		RolesByTaxId(taxId string, roles *ResponseRoles) Context
		Role(roleId string, bizId string, roles *ResponseRole) Context
		DepartmentAndRoles(bizId string, deptRoles *ResponseDepartmentAndRoles) Context
		DepartmentAndRole(deptRoleId string, bizId string, deptRole *ResponseDepartmentAndRole) Context
	}

	Context interface {
		Error() error
		Ping() Context
		LoginPWD(PWD, *ResponseLoginPwdOne) Context
		LoginWithSharedToken(sharedToken string, pwd *ResponseLoginPwdOne) Context
		LoginWithOneChatSharedToken(oneChatToken string, pwd *ResponseLoginPwdOne) Context
		LoginWithRefreshToken(refreshToken string, pwd *ResponseLoginPwdOne) Context
		LoginMobile(mobileNo string, res *ResponseLoginMobile) Context
		LoginMobileConfirm(mobileConfirm *LoginMobileConfirm, res *ResponseLoginMobileConfirm) Context
		RegisterMobile(mobileNo string, res *ResponseRegisterMobile) Context
		RegisterMobileConfirm(mobileConfirm RegisterMobileConfirm, res *ResponseRegisterMobileConfirm) Context
		CheckUserLogin(userLogin string, res *ResponseCheckUserLogin) Context
		CheckServiceClient(CheckServiceClient) bool
		Bearer(accessToken string) BearerContext
	}

	oneAuth struct {
		clientId     string
		clientSecret string
		refCode      string
		redirectPath string
	}
	context struct {
		host string
		oneAuth
		err    error
		bearer string
	}
)

type ResBase struct {
	Result       []byte      `json:"result"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"`
	ResponseCode int         `json:"responseCode"`
}

func (c *context) CheckServiceClient(client CheckServiceClient) bool {
	if err := requests.Call().Post(requests.Params{
		URL: c.apiEndpoint(APIEndpointOAuthPWD),
		HEADERS: map[string]string{
			echo.HeaderContentType: "application/json",
		},
		BODY: c.jsonBody(map[string]interface{}{
			"grant_type":    "password",
			"client_id":     client.ClientID,
			"client_secret": client.ClientSecret,
			"username":      `mocking03193`,
			"password":      `mocking03193`,
		}),
		TIMEOUT: 30,
	}, &resRequest).Error(); err != nil {
		return false
	}
	var resBase ResBase
	if err := json.Unmarshal(resRequest.Result, &resBase); err != nil {
		return false
	}

	if resBase.ErrorMessage == `Client authentication failed` {
		return false
	}

	if resBase.ErrorMessage == `The user credentials were incorrect.` {
		return true
	}
	return false
}

func (c *context) Business() BusinessContext {
	return c
}

func (c *context) Bearer(token string) BearerContext {
	c.bearer = `Bearer ` + token
	return c
}

func (c *context) errorOneID(result []byte) Context {
	var resError ResponseErrorOneID
	if err := json.Unmarshal(result, &resError); err != nil {
		c.err = err
		return c
	}
	c.err = fmt.Errorf(strconv.FormatInt(resError.Responsecode, 10) + " " + resError.Errormessage)
	return c
}

func (c *context) apiEndpoint(endpoint string) string {
	return c.host + endpoint
}

func (c *context) jsonBody(i interface{}) *bytes.Buffer {
	jByte, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(jByte)
}

func (c *context) handleRes(res requests.Response, i interface{}) Context {
	if res.Code != 200 {
		return c.errorOneID(res.Result)
	}
	if err := json.Unmarshal(res.Result, i); err != nil {
		c.err = err
		return c
	}
	c.err = nil
	return c
}

func (c *context) Error() error {
	if c.err != nil {
		return c.err
	}
	return nil
}
