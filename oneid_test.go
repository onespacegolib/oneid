package oneid

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func printStructJson(target interface{}) {
	fooByte, _ := json.MarshalIndent(&target, "", "\t")
	fmt.Println(string(fooByte))
}

var oneId = Init(`https://testoneid.inet.co.th`, `514`, `rYqWkb6pQtsqqJarHeHRMdACYBY5V63MyiQMCzdq`, `Q8JLO3`, ``)

// Ping() Context
func TestPing(t *testing.T) {
	if err := oneId.Ping().Error(); err != nil {
		panic(err)
	}
}

func TestLoginMobile(t *testing.T) {
	mobileNo := "0843546967" //เบอร์ KNOT
	var res ResponseLoginMobile
	if err := oneId.LoginMobile(mobileNo, &res).Error(); err != nil {
		panic(err)
	}
	fmt.Printf("%#v", res)
}

func TestLoginMobileConfirm(t *testing.T) {
	mobileNo := "0843546967" //เบอร์ KNOT
	var res ResponseLoginMobile
	if err := oneId.LoginMobile(mobileNo, &res).Error(); err != nil {
		panic(err)
	}
	//fmt.Printf("%#v", res)
	printStructJson(res)
	mobileConfirm := LoginMobileConfirm{
		MobileNo: "0843546967",
		OTP:      res.Data.Otp,
	}
	var resConfirm ResponseLoginMobileConfirm
	if err := oneId.LoginMobileConfirm(&mobileConfirm, &resConfirm).Error(); err != nil {
		panic(err)
	}
	//fmt.Printf("%#v", resConfirm)
	printStructJson(resConfirm)
}

func TestContext_CheckUserLogin(t *testing.T) {
	userLogin := "0843546967" //เบอร์ KNOT
	var res ResponseCheckUserLogin
	if err := oneId.CheckUserLogin(userLogin, &res).Error(); err != nil {
		panic(err)
	}
	printStructJson(res)
}

func TestLoginWithPWD(t *testing.T) {
	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	fmt.Printf("%#v", res)
	assert.Equal(t, pwd.Username, res.Username)
}

func TestLoginWithShared(t *testing.T) {

	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var login ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &login).Error(); err != nil {
		panic(err)
	}

	var token ResponseGetSharedToken
	if err := oneId.Bearer(login.AccessToken).GetSharedToken(&token).Error(); err != nil {
		printStructJson(err)
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginWithSharedToken(token.Data.SharedToken, &res).Error(); err != nil {
		printStructJson(err)
	}
	printStructJson(res)
	fmt.Println(res)
}

func TestLoginWithOneChatShared(t *testing.T) {
	sharedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdG9rZW4iOiJleUowZVhBaU9pSktWMVFpTENKaGJHY2lPaUpTVXpJMU5pSXNJbXAwYVNJNklqRmtObVl6WmpkbE0yVmhaREJtWkRKaU5HWmpZVEl3T1RnMVlqWXlOamN3T1dFM05UbG1NRFl4TWpSak1tVTNPVEUyTmpVeE9URXpOalJtWXpVME16RmhPRGcyWkRVeFlXTmxPREF4TURkbUlpd2lhMmxrSWpvaUluMC5leUpoZFdRaU9pSTNPU0lzSW1wMGFTSTZJakZrTm1ZelpqZGxNMlZoWkRCbVpESmlOR1pqWVRJd09UZzFZall5Tmpjd09XRTNOVGxtTURZeE1qUmpNbVUzT1RFMk5qVXhPVEV6TmpSbVl6VTBNekZoT0RnMlpEVXhZV05sT0RBeE1EZG1JaXdpYVdGMElqb3hOakk0TnpRNU5UazBMQ0p1WW1ZaU9qRTJNamczTkRrMU9UUXNJbVY0Y0NJNk1UWXlPRGMwT1RnNU5Dd2ljM1ZpSWpvaU9ERXhPREEwTnpnNU1qUTRJaXdpZFhObGNtNWhiV1VpT2lKdGNHMXBibWx0WVd3aWZRLnJXRXhSZUNrVS1IVmhzdmphYTNfUHdGQ2JWdkFhWGF0WlZxRmlfa3FIaGxCR0p5NHNmV2Q2T0h1QW5Pb0ZNNGZwcWJTWE9jdzY0MXhXZFAxbTZGSFlfU0N5NkxiNWZmS0U4Y1BOWEpNQjNSVThWVmllOW9vSWZoZV91anFnQWZXWVc2SlZDRmJXQ1N0UFJkdmw2QjNlRE92Z3hmQWNvd2t4Vlh6RVV3bUo2eGJPa2xTSnphWi14Y21MZV9hYW9RM3BBMHlpN04wVGRnZUItQy05eWRDYTRiWEJvTXVCS18wZXFoWmZHNDFDcGVpRTFyNVNhVzBPdnlZNW9iRTlLRlNMcENUazhiMzJDMG44Yi1EYnNTZTNCVUNWanlvU0NRRGZVM0piaUtVRUVYUmt2T3d1X0VVMllzOHJWcUZzZEQ1YzkxVHlBNy1uM2U3THU4bmhydDB5USJ9.fv6CFJk5_WusAybBPribVLJiptK-u8ms1pUMOHMUjws"
	var res ResponseLoginPwdOne
	if err := oneId.LoginWithOneChatSharedToken(sharedToken, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, res.AccountId, "811804789248")
}

func TestGetAccount(t *testing.T) {
	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, pwd.Username, res.Username)
	var resAcc ResponseGetAccount
	if err := oneId.Bearer(res.AccessToken).GetAccount(&resAcc).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, resAcc.ID, "811804789248")
}

func TestGetServiceCitizenInfo(t *testing.T) {
	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, pwd.Username, res.Username)
	var resInfo ResponseServiceCitizenInfo
	if err := oneId.Bearer(res.AccessToken).GetServiceCitizenInfo(&resInfo).Error(); err != nil {
		panic(err)
	}
	printStructJson(resInfo)
	//assert.Equal(t, resAcc, "811804789248")
}

func TestContext_GetSharedToken(t *testing.T) {
	pwd := PWD{
		Username: "devonespace01",
		Password: "0neSpace",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, pwd.Username, res.Username)
	//fmt.Printf(`%#v`, res.AccessToken)
	var shared ResponseGetSharedToken
	if err := oneId.Bearer(res.AccessToken).GetSharedToken(&shared).Error(); err != nil {
		panic(err)
	}
	print(shared.Data.SharedToken)
	assert.NotEmpty(t, shared.Data.SharedToken)
}

func TestContext_AllBranch(t *testing.T) {
	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, pwd.Username, res.Username)

	var allBranchBusiness ResponseAllBranchBusiness
	if err := oneId.Bearer(res.AccessToken).Business().AllBranch(`3467976236700`, &allBranchBusiness).Error(); err != nil {
		//printStructJson(err)
	}

	printStructJson(allBranchBusiness)
}

func TestContext_AccountPaginate(t *testing.T) {
	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, pwd.Username, res.Username)

	var accountPaginate ResponseAccountPaginate
	if err := oneId.Bearer(res.AccessToken).Business().AccountPaginate(`1605248760`, `1000000`, `1`, &accountPaginate).Error(); err != nil {
		//printStructJson(err)
	}

	printStructJson(accountPaginate)
}

func TestContext_Businesses(t *testing.T) {
	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, pwd.Username, res.Username)

	var businesses ResponseBusinesses
	if err := oneId.Bearer(res.AccessToken).Business().All(&businesses).Error(); err != nil {
		panic(err)
	}

	printStructJson(businesses)
}

func TestContext_Department(t *testing.T) {
	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, pwd.Username, res.Username)

	var department ResponseDepartment
	if err := oneId.Bearer(res.AccessToken).Business().Department(`17a6e230-8097-11eb-9a78-eda68aa3bede` ,`1605248760`, &department).Error(); err != nil {
		//printStructJson(err)
	}

	printStructJson(department)
}

func TestContext_Departments(t *testing.T) {
	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, pwd.Username, res.Username)

	var departments ResponseDepartments
	if err := oneId.Bearer(res.AccessToken).Business().Departments(`1605248760`, &departments).Error(); err != nil {
		//printStructJson(err)
	}

	printStructJson(departments)
}

func TestContext_Roles(t *testing.T) {
	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, pwd.Username, res.Username)

	var roles ResponseRoles
	if err := oneId.Bearer(res.AccessToken).Business().Roles(`1605248760`, &roles).Error(); err != nil {
		//printStructJson(err)
	}

	printStructJson(roles)
}

func TestContext_Role(t *testing.T) {
	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, pwd.Username, res.Username)

	var role ResponseRole
	if err := oneId.Bearer(res.AccessToken).Business().Role(`18bf52e0-2579-11eb-b8ee-71157c15cce4`, `1605248760`, &role).Error(); err != nil {
		//printStructJson(err)
	}

	printStructJson(role)
}

func TestContext_DepartmentAndRoles(t *testing.T) {
	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, pwd.Username, res.Username)

	var deptRoles ResponseDepartmentAndRoles
	if err := oneId.Bearer(res.AccessToken).Business().DepartmentAndRoles(`1605248760`, &deptRoles).Error(); err != nil {
		//printStructJson(err)
	}

	printStructJson(deptRoles)
}

func TestContext_DepartmentAndRole(t *testing.T) {
	pwd := PWD{
		Username: "mpminimal",
		Password: "mp12345678",
	}
	var res ResponseLoginPwdOne
	if err := oneId.LoginPWD(pwd, &res).Error(); err != nil {
		panic(err)
	}
	assert.Equal(t, pwd.Username, res.Username)

	var deptRole ResponseDepartmentAndRole
	if err := oneId.Bearer(res.AccessToken).Business().DepartmentAndRole(`6e2c1940-817e-11eb-9a66-4d220331cbad`, `1605248760`, &deptRole).Error(); err != nil {
		//printStructJson(err)
	}

	printStructJson(deptRole)
}

func TestContext_CheckServiceClient(t *testing.T) {
	t.Run(`Success`, func(t *testing.T) {
		request := CheckServiceClient{
			ClientID: `514`,
			ClientSecret: `rYqWkb6pQtsqqJarHeHRMdACYBY5V63MyiQMCzdq`,
		}
		ok := oneId.CheckServiceClient(request)
		assert.True(t, ok)
	})
	t.Run(`Fake`, func(t *testing.T) {
		request := CheckServiceClient{
			ClientID: `51433`,
			ClientSecret: `rYqWkb6pQtsqqJarHeHRMdACYBY5V63MyiQMCzdq`,
		}
		ok := oneId.CheckServiceClient(request)
		assert.False(t, ok)
	})
}
