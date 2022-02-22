package oneid

const (
	// APIEndpointRegister ใช้สำหรับสมัครบัญชีในการเข้าใช้งานระบบ https://api-id.one.th/#/1.%20API%20Register%20(Citizen)/register_api
	APIEndpointRegister = `/api/citizen/register`
	// APIEndpointOAuthPWD ใช้สำหรับ Login โดยใช้ Username/Password https://api-id.one.th/#/2.%20API%20Login%20PWD%20(Citizen)/citizenRegister
	APIEndpointOAuthPWD = `/api/oauth/getpwd`
	// APIEndpointOAuthCode (Step 1) ใช้สำหรับ Login Authorization Code เพื่อเก็บ Code https://api-id.one.th/#/3.%20API%20Login%20Authorization%20Code%20(Citizen)/oauthtoken
	APIEndpointOAuthCode = `/api/oauth/getcode`
	// APIEndpointOAuthToken (Step 2) ใช้ Code เพื่อเอา Access Token https://api-id.one.th/#/3.%20API%20Login%20Authorization%20Code%20(Citizen)/oauthtoken
	APIEndpointOAuthToken = `/oauth/token`
	// APIEndpointOAuthSharedToken ใช้สำหรับ Login โดยใช้ Shared Token ที่ได้จาก Service อื่น https://api-id.one.th/#/4.%20API%20Login%20Shared%20Token%20(Citizen)/oauth%40shared-token
	APIEndpointOAuthSharedToken = `/api/oauth/shared-token`
	// APIEndpointServiceSharedToken ใช้สำหรับแชร์ Token ระหว่าง Service https://api-id.one.th/#/5.%20API%20Get%20Shared%20Token%20(Citizen)/v2%40service%40shared-token
	APIEndpointServiceSharedToken = `/api/v2/service/shared-token`

	// APIEndpointAccount ใช้สำหรับดึงข้อมูลผู้ใช้งาน https://api-id.one.th/#/16.%09API%20Get%20Data%20(Citizen)/account
	APIEndpointGetAccount = `/api/account`

	// APIEndpointGetRefreshToken Login โดยใช้ Refresh Token https://api-id.one.th/#/6.%20API%20Refresh%20Token%20(Citizen)/oauth%40get_refresh_token
	APIEndpointGetRefreshToken = `/api/oauth/get_refresh_token`

	// APIEndpointRegisterMobile (Step 1) ใช้สำหรับสมัคร ONE ID ด้วยเบอร์โทรศัพท์ https://api-id.one.th/#/17.%09API%20Mobile%20Register%20(Citizen)/register%40mobile
	APIEndpointRegisterMobile = `/api/register/mobile`

	// APIEndpointRegisterMobileConfirm (Step 2) ใช้สำหรับตรวจสอบ OTP เพื่อยืนยันการสมัครผ่านเบอร์โทรศัพท์ https://api-id.one.th/#/17.%09API%20Mobile%20Register%20(Citizen)/register%40mobile%40confirm
	APIEndpointRegisterMobileConfirm = `/api/register/mobile/confirm`

	//APIEndpointLoginMobile (Step 1) ใช้ส่ง OTP สำหรับการ Login https://api-id.one.th/#/18.%09API%20Mobile%20Login%20(Citizen)/oauth%40otp
	APIEndpointLoginMobile = `/api/oauth/otp`

	//APIEndpointLoginMobileConfirm (Step 2) ใช้ยืนยัน OTP สำหรับการ Login https://api-id.one.th/#/18.%09API%20Mobile%20Login%20(Citizen)/oauth%40otp%40confirm
	APIEndpointLoginMobileConfirm = `/api/oauth/otp/confirm`

	//APIEndpointCheckUserLogin ใช้สำหรับ Check User Login ของ ONE ID สามารถใช้ Username เบอร์โทรศัพท์ หรือ อีเมลได้ https://api-id.one.th/#/15.%09API%20Check%20Data%20(Citizen)/check%40userlogin
	APIEndpointCheckUserLogin = `/api/check/userlogin`

	//APIEndpointServiceCitizenInfo ใช้สำหรับดึงข้อมูลของผู้ใช้ ตาม Scope ของ Access Token https://api-id.one.th/#/16.%09API%20Get%20Data%20(Citizen)/v2%40service%40citizen%40info
	APIEndpointServiceCitizenInfo = `/api/v2/service/citizen/info`
	/**
	API Get Account (Business)
	*/

	//
	APIEndpointAllBranchBusiness  = `/api/v3/business/service/get-all-branch-business`
	APIEndpointAccountPaginate  = `/api/v3/business/service/account/paginate`
	APIEndpointAccount  = `/api/v3/business/service/account/:account_id`
	APIEndpointBusinesses  = `/api/v3/business/service/business`
	APIEndpointDepartments = `/api/v3/business/service/department`
	APIEndpointDepartment  = `/api/v3/business/service/department/:dept_id`

	APIEndpointRoles = `/api/v3/business/service/role`
	APIEndpointRole  = `/api/v3/business/service/role/:role_id`

	APIEndpontDepartmentAndRoles = `/api/v3/business/service/dept-role`
	APIEndpontDepartmentAndRole  = `/api/v3/business/service/dept-role/:dept_role_id`
	/**
	--------------------------------------------------------------------------------------------------------------------
	*/
)
