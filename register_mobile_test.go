package oneid

import (
	"fmt"
	"testing"
)

func TestRegisterMobile(t *testing.T) {

	mobileNo := "89495952"
	var register ResponseRegisterMobile
	if err := oneId.RegisterMobile(mobileNo, &register).Error(); err != nil {
		panic(err)
	}

	printStructJson(register)
	fmt.Println(register)
}

func TestRegisterMobileConfirm(t *testing.T) {

	mobileNo := "843325464"
	var register ResponseRegisterMobile
	if err := oneId.RegisterMobile(mobileNo, &register).Error(); err != nil {
		panic(err)
	}

	printStructJson(register)
	mobileConfirm := RegisterMobileConfirm{
		MobileRefID: register.Data.MobileRefId,
		OTP:         register.Data.Otp,
		FlgTerm:     "Y",
	}
	var registerConfirm ResponseRegisterMobileConfirm
	if err := oneId.RegisterMobileConfirm(mobileConfirm, &registerConfirm).Error(); err != nil {
		panic(err)
	}

	printStructJson(registerConfirm)
}
