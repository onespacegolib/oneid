package oneid

import (
	"github.com/labstack/echo/v4"
	requests "github.com/onespacegolib/http-requests"
	"net/url"
	"strings"
)

func (c *context) All(businesses *ResponseBusinesses) Context {
	if err := requests.Call().Get(requests.Params{
		URL:  c.apiEndpoint(APIEndpointBusinesses),
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
	return c.handleRes(resRequest, &businesses)
}

func (c *context) Departments(bizId string, departments *ResponseDepartments) Context {
	base, _ := url.Parse(c.apiEndpoint(APIEndpointDepartments))

	query := url.Values{}
	query.Add(`biz_id`, bizId)

	base.RawQuery = query.Encode()

	if err := requests.Call().Get(requests.Params{
		URL:  base.String(),
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
	return c.handleRes(resRequest, &departments)
}

func (c *context) Department(deptId string, bizId string, department *ResponseDepartment) Context {
	base, _ := url.Parse(strings.Replace(
		c.apiEndpoint(APIEndpointDepartment), `:dept_id`, deptId, -1),
	)

	query := url.Values{}
	query.Add(`biz_id`, bizId)

	base.RawQuery = query.Encode()

	if err := requests.Call().Get(requests.Params{
		URL:  base.String(),
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
	return c.handleRes(resRequest, &department)
}

func (c *context) Roles(bizId string, roles *ResponseRoles) Context {
	base, _ := url.Parse(c.apiEndpoint(APIEndpointRoles))

	query := url.Values{}
	query.Add(`biz_id`, bizId)

	base.RawQuery = query.Encode()

	if err := requests.Call().Get(requests.Params{
		URL:  base.String(),
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
	return c.handleRes(resRequest, &roles)
}

func (c *context) Role(roleId string, bizId string, role *ResponseRole) Context {
	base, _ := url.Parse(strings.Replace(
		c.apiEndpoint(APIEndpointRole), `:role_id`, roleId, -1),
	)

	query := url.Values{}
	query.Add(`biz_id`, bizId)

	base.RawQuery = query.Encode()

	if err := requests.Call().Get(requests.Params{
		URL:  base.String(),
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
	return c.handleRes(resRequest, &role)
}

func (c *context) DepartmentAndRoles(bizId string, deptRole *ResponseDepartmentAndRoles) Context {
	base, _ := url.Parse(c.apiEndpoint(APIEndpontDepartmentAndRoles))

	query := url.Values{}
	query.Add(`biz_id`, bizId)

	base.RawQuery = query.Encode()

	if err := requests.Call().Get(requests.Params{
		URL:  base.String(),
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
	return c.handleRes(resRequest, &deptRole)
}

func (c *context) DepartmentAndRole(deptRoleId string, bizId string, deptRole *ResponseDepartmentAndRole) Context {
	base, _ := url.Parse(strings.Replace(
		c.apiEndpoint(APIEndpontDepartmentAndRole), `:dept_role_id`, deptRoleId, -1),
	)

	query := url.Values{}
	query.Add(`biz_id`, bizId)

	base.RawQuery = query.Encode()

	if err := requests.Call().Get(requests.Params{
		URL:  base.String(),
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
	return c.handleRes(resRequest, &deptRole)
}
