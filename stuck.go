package oneid

type PWD struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CheckServiceClient struct {
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type RegisterMobileConfirm struct {
	MobileRefID string `json:"mobile_ref_id"`
	OTP         string `json:"otp"`
	FlgTerm     string `json:"flg_term"`
}

type LoginMobileConfirm struct {
	MobileNo string `json:"mobile_no"`
	OTP      string `json:"otp"`
}

type ResponseLoginMobile struct {
	Result string `json:"result"`
	Data   struct {
		Otp     string `json:"otp"`
		Refcode string `json:"refcode"`
	} `json:"data"`
	ErrorMessage string `json:"errorMessage"`
	Code         int    `json:"code"`
}

type ResponseLoginMobileConfirm struct {
	TokenType      string `json:"token_type"`
	ExpiresIn      int    `json:"expires_in"`
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	ExpirationDate string `json:"expiration_date"`
	AccountId      string `json:"account_id"`
	Result         string `json:"result"`
	Username       string `json:"username"`
	LoginBy        string `json:"login_by"`
}

type ResponseCheckUserLogin struct {
	Result       string `json:"result"`
	Data         string `json:"data"`
	ErrorMessage string `json:"errorMessage"`
	Code         int    `json:"code"`
}

type ResponseServiceCitizenInfo struct {
	Result       string `json:"result"`
	Data         string `json:"data"`
	ErrorMessage string `json:"errorMessage"`
	Code         int    `json:"code"`
}

type SharedToken struct {
	SharedToken string `json:"shared_token"`
}

type BodyWithSharedToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	SharedToken  string `json:"shared_token"`
	Refcode      string `json:"refcode"`
}

type ResponseRegisterMobile struct {
	Result string `json:"result"`
	Data   struct {
		MobileNo    string `json:"mobile_no"`
		MobileRefId string `json:"mobile_ref_id"`
		Otp         string `json:"otp"`
		Refcode     string `json:"refcode"`
	} `json:"data"`
	ErrorMessage string `json:"errorMessage"`
	Code         int    `json:"code"`
}

type ResponseRegisterMobileConfirm struct {
	Result string `json:"result"`
	Data   struct {
		AccountId string `json:"account_id"`
		Email     string `json:"email"`
		MobileNo  string `json:"mobile_no"`
		OneChat   string `json:"one_chat"`
		OneBox    string `json:"one_box"`
		OneChatId string `json:"one_chat_id"`
		NotiChat  string `json:"noti_chat"`
		Login     struct {
			TokenType      string `json:"token_type"`
			ExpiresIn      int    `json:"expires_in"`
			AccessToken    string `json:"access_token"`
			RefreshToken   string `json:"refresh_token"`
			ExpirationDate string `json:"expiration_date"`
			AccountId      string `json:"account_id"`
			Result         string `json:"result"`
			Username       string `json:"username"`
			LoginBy        string `json:"login_by"`
		} `json:"login"`
	} `json:"data"`
	ErrorMessage string `json:"errorMessage"`
	Code         int    `json:"code"`
}

type ResponseLoginPwdOne struct {
	AccessToken    string `json:"access_token"`
	AccountId      string `json:"account_id"`
	ExpirationDate string `json:"expiration_date"`
	ExpiresIn      int    `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	Result         string `json:"result"`
	TokenType      string `json:"token_type"`
	Username       string `json:"username"`
}

type ResponseGetSharedToken struct {
	Result string `json:"result"`
	Data   struct {
		SharedToken string `json:"shared_token"`
		IssueAt     string `json:"issue_at"`
		ExpiredAt   string `json:"expired_at"`
	} `json:"data"`
	ErrorMessage interface{} `json:"errorMessage"`
	Code         int         `json:"code"`
}

type ResponseGetAccount struct {
	ID                 string `json:"id"`
	FirstNameTh        string `json:"first_name_th"`
	LastNameTh         string `json:"last_name_th"`
	FirstNameEng       string `json:"first_name_eng"`
	LastNameEng        string `json:"last_name_eng"`
	AccountTitleTh     string `json:"account_title_th"`
	AccountTitleEng    string `json:"account_title_eng"`
	IDCardType         string `json:"id_card_type"`
	IDCardNum          string `json:"id_card_num"`
	HashIDCardNum      string `json:"hash_id_card_num"`
	AccountCategory    string `json:"account_category"`
	AccountSubCategory string `json:"account_sub_category"`
	ThaiEmail          string `json:"thai_email"`
	ThaiEmail2         string `json:"thai_email2"`
	ThaiEmail3         string `json:"thai_email3"`
	StatusCd           string `json:"status_cd"`
	BirthDate          string `json:"birth_date"`
	StatusDt           string `json:"status_dt"`
	RegisterDt         string `json:"register_dt"`
	AddressID          string `json:"address_id"`
	CreatedAt          string `json:"created_at"`
	CreatedBy          string `json:"created_by"`
	UpdatedAt          string `json:"updated_at"`
	UpdatedBy          string `json:"updated_by"`
	Reason             string `json:"reason"`
	TelNo              string `json:"tel_no"`
	NameOnDocumentTh   string `json:"name_on_document_th"`
	NameOnDocumentEng  string `json:"name_on_document_eng"`
	BlockchainFlg      string `json:"blockchain_flg"`
	Mobile             []struct {
		ID        string `json:"id"`
		MobileNo  string `json:"mobile_no"`
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
		UpdatedAt string `json:"updated_at"`
		UpdatedBy string `json:"updated_by"`
		DeletedAt string `json:"deleted_at"`
		Pivot     struct {
			AccountID             string `json:"account_id"`
			MobileID              string `json:"mobile_id"`
			CreatedAt             string `json:"created_at"`
			UpdatedAt             string `json:"updated_at"`
			StatusCd              string `json:"status_cd"`
			PrimaryFlg            string `json:"primary_flg"`
			MobileConfirmFlg      string `json:"mobile_confirm_flg"`
			MobileConfirmDt       string `json:"mobile_confirm_dt"`
			MobileLoginConfirmFlg string `json:"mobile_login_confirm_flg"`
			MobileLoginConfirmDt  string `json:"mobile_login_confirm_dt"`
		} `json:"pivot"`
	} `json:"mobile"`
	Email []struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
		UpdatedAt string `json:"updated_at"`
		UpdatedBy string `json:"updated_by"`
		DeletedAt string `json:"deleted_at"`
		Pivot     struct {
			AccountID            string `json:"account_id"`
			EmailID              string `json:"email_id"`
			CreatedAt            string `json:"created_at"`
			UpdatedAt            string `json:"updated_at"`
			StatusCd             string `json:"status_cd"`
			PrimaryFlg           string `json:"primary_flg"`
			EmailConfirmFlg      string `json:"email_confirm_flg"`
			EmailConfirmDt       string `json:"email_confirm_dt"`
			EmailLoginConfirmFlg string `json:"email_login_confirm_flg"`
			EmailLoginConfirmDt  string `json:"email_login_confirm_dt"`
		} `json:"pivot"`
	} `json:"email"`
	Address []struct {
		ID           string      `json:"id"`
		HouseNo      string      `json:"house_no"`
		MooBan       interface{} `json:"moo_ban"`
		BuildingName interface{} `json:"building_name"`
		Street       string      `json:"street"`
		Soi          string      `json:"soi"`
		ZipcodeID    string      `json:"zipcode_id"`
		CreatedAt    string      `json:"created_at"`
		CreatedBy    string      `json:"created_by"`
		UpdatedAt    string      `json:"updated_at"`
		UpdatedBy    string      `json:"updated_by"`
		DeletedAt    interface{} `json:"deleted_at"`
		RoomNo       interface{} `json:"room_no"`
		Floor        interface{} `json:"floor"`
		FaxNumber    interface{} `json:"fax_number"`
		Type         string      `json:"type"`
		HouseCode    interface{} `json:"house_code"`
		MooNo        interface{} `json:"moo_no"`
		DataDetail   interface{} `json:"data_detail"`
		CountryID    string      `json:"country_id"`
		Yaek         string      `json:"yaek"`
		Pivot        struct {
			AccountID  string      `json:"account_id"`
			AddressID  string      `json:"address_id"`
			CreatedAt  string      `json:"created_at"`
			UpdatedAt  string      `json:"updated_at"`
			StatusCd   string      `json:"status_cd"`
			PrimaryFlg string      `json:"primary_flg"`
			Type       interface{} `json:"type"`
		} `json:"pivot"`
		Country struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			Alpha2    string `json:"alpha_2"`
			Alpha3    string `json:"alpha_3"`
			Numeric   string `json:"numeric"`
			Iso       string `json:"iso"`
			CreatedAt string `json:"created_at"`
			CreatedBy string `json:"created_by"`
			UpdatedAt string `json:"updated_at"`
			UpdatedBy string `json:"updated_by"`
		} `json:"country"`
		Zipcode struct {
			ID              string `json:"id"`
			Amphoe          string `json:"amphoe"`
			Tambon          string `json:"tambon"`
			Province        string `json:"province"`
			Zipcode         string `json:"zipcode"`
			SubdistrictCode string `json:"subdistrict_code"`
			DistrictCode    string `json:"district_code"`
			ProvinceCode    string `json:"province_code"`
			CreatedAt       string `json:"created_at"`
			CreatedBy       string `json:"created_by"`
			UpdatedAt       string `json:"updated_at"`
			UpdatedBy       string `json:"updated_by"`
		} `json:"zipcode"`
	} `json:"address"`
	AccountAttribute []string `json:"account_attribute"`
	Status           string   `json:"status"`
	LastUpdate       string   `json:"last_update"`
}

type ResponseErrorOneID struct {
	Result       string      `json:"result"`
	Data         interface{} `json:"data"`
	Errormessage string      `json:"errorMessage"`
	Responsecode int64       `json:"responseCode"`
}

type ResponseDepartmentAndRole struct {
	Code int `json:"code"`
	Data struct {
		BizId      string      `json:"biz_id"`
		CreatedAt  string      `json:"created_at"`
		CreatedBy  string      `json:"created_by"`
		DeletedAt  interface{} `json:"deleted_at"`
		Department struct {
			CreatedAt    string      `json:"created_at"`
			CreatedBy    string      `json:"created_by"`
			DeletedAt    interface{} `json:"deleted_at"`
			DeptName     string      `json:"dept_name"`
			DeptPosition string      `json:"dept_position"`
			Id           string      `json:"id"`
			ParentDeptId string      `json:"parent_dept_id"`
			UpdatedAt    string      `json:"updated_at"`
			UpdatedBy    string      `json:"updated_by"`
		} `json:"department"`
		DeptId           string        `json:"dept_id"`
		DeptRoleLevel    string        `json:"dept_role_level"`
		DeptRolePosition string        `json:"dept_role_position"`
		HasAccount       []interface{} `json:"has_account"`
		HasPermission    []interface{} `json:"has_permission"`
		Id               string        `json:"id"`
		ParentDeptRoleId string        `json:"parent_dept_role_id"`
		Role             struct {
			CreatedAt string      `json:"created_at"`
			CreatedBy string      `json:"created_by"`
			DeletedAt interface{} `json:"deleted_at"`
			Id        string      `json:"id"`
			RoleLevel string      `json:"role_level"`
			RoleName  string      `json:"role_name"`
			UpdatedAt string      `json:"updated_at"`
			UpdatedBy string      `json:"updated_by"`
		} `json:"role"`
		RoleId    string `json:"role_id"`
		Status    string `json:"status"`
		UpdatedAt string `json:"updated_at"`
		UpdatedBy string `json:"updated_by"`
	} `json:"data"`
	ErrorMessage interface{} `json:"errorMessage"`
	Result       string      `json:"result"`
}

type ResponseDepartmentAndRoles struct {
	Result string `json:"result"`
	Data   []struct {
		Id               string      `json:"id"`
		BizId            string      `json:"biz_id"`
		DeptId           string      `json:"dept_id"`
		RoleId           string      `json:"role_id"`
		DeptRolePosition string      `json:"dept_role_position"`
		DeptRoleLevel    string      `json:"dept_role_level"`
		ParentDeptRoleId *string     `json:"parent_dept_role_id"`
		Status           string      `json:"status"`
		CreatedAt        string      `json:"created_at"`
		CreatedBy        string      `json:"created_by"`
		UpdatedAt        string      `json:"updated_at"`
		UpdatedBy        string      `json:"updated_by"`
		DeletedAt        interface{} `json:"deleted_at"`
		Department       struct {
			Id           string      `json:"id"`
			DeptName     string      `json:"dept_name"`
			DeptPosition string      `json:"dept_position"`
			ParentDeptId *string     `json:"parent_dept_id"`
			CreatedAt    string      `json:"created_at"`
			CreatedBy    string      `json:"created_by"`
			UpdatedAt    string      `json:"updated_at"`
			UpdatedBy    string      `json:"updated_by"`
			DeletedAt    interface{} `json:"deleted_at"`
		} `json:"department"`
		Role struct {
			Id        string      `json:"id"`
			RoleName  string      `json:"role_name"`
			RoleLevel string      `json:"role_level"`
			CreatedAt string      `json:"created_at"`
			CreatedBy string      `json:"created_by"`
			UpdatedAt string      `json:"updated_at"`
			UpdatedBy string      `json:"updated_by"`
			DeletedAt interface{} `json:"deleted_at"`
		} `json:"role"`
	} `json:"data"`
	ErrorMessage interface{} `json:"errorMessage"`
	Code         int         `json:"code"`
}

type ResponseRole struct {
	Result string `json:"result"`
	Data   struct {
		Id            string      `json:"id"`
		RoleName      string      `json:"role_name"`
		RoleLevel     string      `json:"role_level"`
		CreatedAt     string      `json:"created_at"`
		CreatedBy     string      `json:"created_by"`
		UpdatedAt     string      `json:"updated_at"`
		UpdatedBy     string      `json:"updated_by"`
		DeletedAt     interface{} `json:"deleted_at"`
		HasDepartment []struct {
			Id               string      `json:"id"`
			BizId            string      `json:"biz_id"`
			DeptId           string      `json:"dept_id"`
			RoleId           string      `json:"role_id"`
			DeptRolePosition string      `json:"dept_role_position"`
			DeptRoleLevel    string      `json:"dept_role_level"`
			ParentDeptRoleId string      `json:"parent_dept_role_id"`
			Status           string      `json:"status"`
			CreatedAt        string      `json:"created_at"`
			CreatedBy        string      `json:"created_by"`
			UpdatedAt        string      `json:"updated_at"`
			UpdatedBy        string      `json:"updated_by"`
			DeletedAt        interface{} `json:"deleted_at"`
			Department       struct {
				Id           string      `json:"id"`
				DeptName     string      `json:"dept_name"`
				DeptPosition string      `json:"dept_position"`
				ParentDeptId string      `json:"parent_dept_id"`
				CreatedAt    string      `json:"created_at"`
				CreatedBy    string      `json:"created_by"`
				UpdatedAt    string      `json:"updated_at"`
				UpdatedBy    string      `json:"updated_by"`
				DeletedAt    interface{} `json:"deleted_at"`
			} `json:"department"`
		} `json:"has_department"`
		HasAccount []struct {
			Id         string      `json:"id"`
			AccountId  string      `json:"account_id"`
			BizId      string      `json:"biz_id"`
			DeptRoleId interface{} `json:"dept_role_id"`
			DeptId     interface{} `json:"dept_id"`
			RoleId     string      `json:"role_id"`
			Status     string      `json:"status"`
			SpId       interface{} `json:"sp_id"`
			StartAt    interface{} `json:"start_at"`
			ExpiredAt  interface{} `json:"expired_at"`
			CreatedAt  string      `json:"created_at"`
			CreatedBy  string      `json:"created_by"`
			UpdatedAt  string      `json:"updated_at"`
			UpdatedBy  string      `json:"updated_by"`
			DeletedAt  interface{} `json:"deleted_at"`
			Account    struct {
				Id                  string      `json:"id"`
				FirstNameTh         string      `json:"first_name_th"`
				MiddleNameTh        interface{} `json:"middle_name_th"`
				LastNameTh          string      `json:"last_name_th"`
				FirstNameEng        string      `json:"first_name_eng"`
				MiddleNameEng       interface{} `json:"middle_name_eng"`
				LastNameEng         string      `json:"last_name_eng"`
				SpecialTitleNameTh  interface{} `json:"special_title_name_th"`
				SpecialTitleNameEng interface{} `json:"special_title_name_eng"`
				AccountTitleTh      string      `json:"account_title_th"`
				AccountTitleEng     string      `json:"account_title_eng"`
				IdCardType          string      `json:"id_card_type"`
				IdCardNum           string      `json:"id_card_num"`
				HashIdCardNum       string      `json:"hash_id_card_num"`
				AccountCategory     string      `json:"account_category"`
				AccountSubCategory  string      `json:"account_sub_category"`
				ThaiEmail           string      `json:"thai_email"`
				ThaiEmail2          string      `json:"thai_email2"`
				ThaiEmail3          interface{} `json:"thai_email3"`
				StatusCd            string      `json:"status_cd"`
				BirthDate           string      `json:"birth_date"`
				StatusDt            string      `json:"status_dt"`
				RegisterDt          string      `json:"register_dt"`
				AddressId           interface{} `json:"address_id"`
				CreatedAt           string      `json:"created_at"`
				CreatedBy           string      `json:"created_by"`
				UpdatedAt           string      `json:"updated_at"`
				UpdatedBy           string      `json:"updated_by"`
				Reason              interface{} `json:"reason"`
				TelNo               interface{} `json:"tel_no"`
				NameOnDocumentTh    interface{} `json:"name_on_document_th"`
				NameOnDocumentEng   interface{} `json:"name_on_document_eng"`
				BlockchainFlg       string      `json:"blockchain_flg"`
			} `json:"account"`
			Employee interface{} `json:"employee"`
		} `json:"has_account"`
	} `json:"data"`
	ErrorMessage interface{} `json:"errorMessage"`
	Code         int         `json:"code"`
}

type ResponseRoles struct {
	Result string `json:"result"`
	Data   []struct {
		Id        string      `json:"id"`
		RoleName  string      `json:"role_name"`
		RoleLevel string      `json:"role_level"`
		CreatedAt string      `json:"created_at"`
		CreatedBy string      `json:"created_by"`
		UpdatedAt string      `json:"updated_at"`
		UpdatedBy string      `json:"updated_by"`
		DeletedAt interface{} `json:"deleted_at"`
	} `json:"data"`
	ErrorMessage interface{} `json:"errorMessage"`
	Code         int         `json:"code"`
}

type ResponseDepartment struct {
	Result string `json:"result"`
	Data   struct {
		Id           string        `json:"id"`
		DeptName     string        `json:"dept_name"`
		DeptPosition string        `json:"dept_position"`
		ParentDeptId interface{}   `json:"parent_dept_id"`
		CreatedAt    string        `json:"created_at"`
		CreatedBy    string        `json:"created_by"`
		UpdatedAt    string        `json:"updated_at"`
		UpdatedBy    string        `json:"updated_by"`
		DeletedAt    interface{}   `json:"deleted_at"`
		HasRole      []interface{} `json:"has_role"`
		HasAccount   []interface{} `json:"has_account"`
	} `json:"data"`
	ErrorMessage interface{} `json:"errorMessage"`
	Code         int         `json:"code"`
}

type ResponseDepartments struct {
	Result string `json:"result"`
	Data   []struct {
		Id           string      `json:"id"`
		DeptName     string      `json:"dept_name"`
		DeptPosition string      `json:"dept_position"`
		ParentDeptId *string     `json:"parent_dept_id"`
		CreatedAt    string      `json:"created_at"`
		CreatedBy    string      `json:"created_by"`
		UpdatedAt    string      `json:"updated_at"`
		UpdatedBy    string      `json:"updated_by"`
		DeletedAt    interface{} `json:"deleted_at"`
	} `json:"data"`
	ErrorMessage interface{} `json:"errorMessage"`
	Code         int         `json:"code"`
}

type ResponseBusinesses struct {
	Result string `json:"result"`
	Data   []struct {
		Id                string      `json:"id"`
		NameTh            string      `json:"name_th"`
		NameEng           string      `json:"name_eng"`
		TitleTh           string      `json:"title_th"`
		TitleEng          string      `json:"title_eng"`
		TaxId             string      `json:"tax_id"`
		HashTaxId         string      `json:"hash_tax_id"`
		TelNo             *string     `json:"tel_no"`
		NameOnDocumentTh  *string     `json:"name_on_document_th"`
		NameOnDocumentEng *string     `json:"name_on_document_eng"`
		Onemail           string      `json:"onemail"`
		Status            string      `json:"status"`
		Type              string      `json:"type"`
		Avartar           interface{} `json:"avartar"`
		CreatedAt         string      `json:"created_at"`
		UpdatedAt         string      `json:"updated_at"`
		CreatedBy         string      `json:"created_by"`
		UpdatedBy         string      `json:"updated_by"`
		DeletedAt         interface{} `json:"deleted_at"`
		Business          []struct {
			Id         string      `json:"id"`
			AccountId  string      `json:"account_id"`
			BizId      string      `json:"biz_id"`
			DeptRoleId *string     `json:"dept_role_id"`
			DeptId     *string     `json:"dept_id"`
			RoleId     string      `json:"role_id"`
			Status     string      `json:"status"`
			SpId       *string     `json:"sp_id"`
			StartAt    *string     `json:"start_at"`
			ExpiredAt  *string     `json:"expired_at"`
			CreatedAt  string      `json:"created_at"`
			CreatedBy  string      `json:"created_by"`
			UpdatedAt  string      `json:"updated_at"`
			UpdatedBy  string      `json:"updated_by"`
			DeletedAt  interface{} `json:"deleted_at"`
			DeptRole   *struct {
				Id               string      `json:"id"`
				BizId            string      `json:"biz_id"`
				DeptId           string      `json:"dept_id"`
				RoleId           string      `json:"role_id"`
				DeptRolePosition string      `json:"dept_role_position"`
				DeptRoleLevel    string      `json:"dept_role_level"`
				ParentDeptRoleId *string     `json:"parent_dept_role_id"`
				Status           string      `json:"status"`
				CreatedAt        string      `json:"created_at"`
				CreatedBy        string      `json:"created_by"`
				UpdatedAt        string      `json:"updated_at"`
				UpdatedBy        string      `json:"updated_by"`
				DeletedAt        interface{} `json:"deleted_at"`
			} `json:"dept_role"`
			Department *struct {
				Id           string      `json:"id"`
				DeptName     string      `json:"dept_name"`
				DeptPosition string      `json:"dept_position"`
				ParentDeptId *string     `json:"parent_dept_id"`
				CreatedAt    string      `json:"created_at"`
				CreatedBy    string      `json:"created_by"`
				UpdatedAt    string      `json:"updated_at"`
				UpdatedBy    string      `json:"updated_by"`
				DeletedAt    interface{} `json:"deleted_at"`
			} `json:"department"`
			Role struct {
				Id        string      `json:"id"`
				RoleName  string      `json:"role_name"`
				RoleLevel string      `json:"role_level"`
				CreatedAt string      `json:"created_at"`
				CreatedBy string      `json:"created_by"`
				UpdatedAt string      `json:"updated_at"`
				UpdatedBy string      `json:"updated_by"`
				DeletedAt interface{} `json:"deleted_at"`
			} `json:"role"`
		} `json:"business"`
	} `json:"data"`
	ErrorMessage interface{} `json:"errorMessage"`
	Code         int         `json:"code"`
}

type ResponseLoginAccount struct {
	Result string `json:"result"`
	Data   []struct {
		ID                  string `json:"id"`
		FirstNameTh         string `json:"first_name_th"`
		MiddleNameTh        string `json:"middle_name_th"`
		LastNameTh          string `json:"last_name_th"`
		FirstNameEng        string `json:"first_name_eng"`
		MiddleNameEng       string `json:"middle_name_eng"`
		LastNameEng         string `json:"last_name_eng"`
		SpecialTitleNameTh  string `json:"special_title_name_th"`
		SpecialTitleNameEng string `json:"special_title_name_eng"`
		AccountTitleTh      string `json:"account_title_th"`
		AccountTitleEng     string `json:"account_title_eng"`
		IDCardType          string `json:"id_card_type"`
		IDCardNum           string `json:"id_card_num"`
		HashIDCardNum       string `json:"hash_id_card_num"`
		AccountCategory     string `json:"account_category"`
		AccountSubCategory  string `json:"account_sub_category"`
		ThaiEmail           string `json:"thai_email"`
		ThaiEmail2          string `json:"thai_email2"`
		ThaiEmail3          string `json:"thai_email3"`
		StatusCd            string `json:"status_cd"`
		BirthDate           string `json:"birth_date"`
		RegisterDt          string `json:"register_dt"`
		UpdatedAt           string `json:"updated_at"`
		TelNo               string `json:"tel_no"`
		NameOnDocumentTh    string `json:"name_on_document_th"`
		NameOnDocumentEng   string `json:"name_on_document_eng"`
		LastUpdate          string `json:"last_update"`
		MobileNo            string `json:"mobile_no"`
		Email               string `json:"email"`
		BranchName          string `json:"branch_name"`
		BranchNo            string `json:"branch_no"`
		BlockchainFlg       string `json:"blockchain_flg"`
		Address struct {
			HouseNo                string      `json:"house_no"`
			MooBan                 interface{} `json:"moo_ban"`
			BuildingName           interface{} `json:"building_name"`
			Street                 string      `json:"street"`
			Soi                    string      `json:"soi"`
			RoomNo                 interface{} `json:"room_no"`
			Floor                  interface{} `json:"floor"`
			FaxNumber              interface{} `json:"fax_number"`
			Type                   string      `json:"type"`
			HouseCode              interface{} `json:"house_code"`
			MooNo                  interface{} `json:"moo_no"`
			DataDetail             interface{} `json:"data_detail"`
			Yaek                   string      `json:"yaek"`
			MappingAccountId       string      `json:"mapping_account_id"`
			MappingAddressId       string      `json:"mapping_address_id"`
			MappingCreatedAt       string      `json:"mapping_created_at"`
			MappingUpdatedAt       string      `json:"mapping_updated_at"`
			MappingStatusCd        string      `json:"mapping_status_cd"`
			MappingPrimaryFlg      string      `json:"mapping_primary_flg"`
			CountryName            string      `json:"country_name"`
			CountryAlpha2          string      `json:"country_alpha2"`
			CountryAlpha3          string      `json:"country_alpha3"`
			CountryNumeric         string      `json:"country_numeric"`
			CountryIso             string      `json:"country_iso"`
			ZipcodeAmphoe          string      `json:"zipcode_amphoe"`
			ZipcodeTambon          string      `json:"zipcode_tambon"`
			ZipcodeProvince        string      `json:"zipcode_province"`
			ZipcodeZipcode         string      `json:"zipcode_zipcode"`
			ZipcodeSubdistrictCode string      `json:"zipcode_subdistrict_code"`
			ZipcodeDistrictCode    string      `json:"zipcode_district_code"`
			ZipcodeProvinceCode    string      `json:"zipcode_province_code"`
		} `json:"address"`
	} `json:"data"`
	ErrorMessage interface{} `json:"errorMessage"`
	Code         int         `json:"code"`
}

type ResponseAccountPaginate struct {
	Result string `json:"result"`
	Data   struct {
		AccountDetail []struct {
			ID                  string `json:"id"`
			FirstNameTh         string `json:"first_name_th"`
			MiddleNameTh        string `json:"middle_name_th"`
			LastNameTh          string `json:"last_name_th"`
			FirstNameEng        string `json:"first_name_eng"`
			MiddleNameEng       string `json:"middle_name_eng"`
			LastNameEng         string `json:"last_name_eng"`
			SpecialTitleNameTh  string `json:"special_title_name_th"`
			SpecialTitleNameEng string `json:"special_title_name_eng"`
			AccountTitleTh      string `json:"account_title_th"`
			AccountTitleEng     string `json:"account_title_eng"`
			IDCardType          string `json:"id_card_type"`
			IDCardNum           string `json:"id_card_num"`
			HashIDCardNum       string `json:"hash_id_card_num"`
			AccountCategory     string `json:"account_category"`
			AccountSubCategory  string `json:"account_sub_category"`
			ThaiEmail           string `json:"thai_email"`
			ThaiEmail2          string `json:"thai_email2"`
			ThaiEmail3          string `json:"thai_email3"`
			StatusCd            string `json:"status_cd"`
			BirthDate           string `json:"birth_date"`
			StatusDt            string `json:"status_dt"`
			RegisterDt          string `json:"register_dt"`
			AddressID           string `json:"address_id"`
			CreatedAt           string `json:"created_at"`
			CreatedBy           string `json:"created_by"`
			UpdatedAt           string `json:"updated_at"`
			UpdatedBy           string `json:"updated_by"`
			Reason              string `json:"reason"`
			TelNo               string `json:"tel_no"`
			NameOnDocumentTh    string `json:"name_on_document_th"`
			NameOnDocumentEng   string `json:"name_on_document_eng"`
			BlockchainFlg       string `json:"blockchain_flg"`
			HasEmployeeDetail   struct {
				ID         string `json:"id"`
				Email      string `json:"email"`
				EmployeeId string `json:"employee_id"`
				MobileNo   string `json:"mobile_no"`
				AccountId  string `json:"account_id"`
				BizId      string `json:"biz_id"`
				CreatedAt  string `json:"created_at"`
				CreatedBy  string `json:"created_by"`
				UpdatedAt  string `json:"updated_at"`
				UpdatedBy  string `json:"updated_by"`
				DeletedAt  string `json:"deleted_at"`
				NickName   string `json:"nick_name"`
				ProfileImg string `json:"profile_img"`
			} `json:"has_employee_detail"`
		} `json:"account_detail"`
	} `json:"data"`
	ErrorMessage interface{} `json:"errorMessage"`
	Code         int         `json:"code"`
}

type ResponseAccount struct {
	Result       string      `json:"result"`
	Data         struct {
		ID                  string `json:"id"`
		FirstNameTh         string `json:"first_name_th"`
		MiddleNameTh        string `json:"middle_name_th"`
		LastNameTh          string `json:"last_name_th"`
		FirstNameEng        string `json:"first_name_eng"`
		MiddleNameEng       string `json:"middle_name_eng"`
		LastNameEng         string `json:"last_name_eng"`
		SpecialTitleNameTh  string `json:"special_title_name_th"`
		SpecialTitleNameEng string `json:"special_title_name_eng"`
		AccountTitleTh      string `json:"account_title_th"`
		AccountTitleEng     string `json:"account_title_eng"`
		IDCardType          string `json:"id_card_type"`
		IDCardNum           string `json:"id_card_num"`
		HashIDCardNum       string `json:"hash_id_card_num"`
		AccountCategory     string `json:"account_category"`
		AccountSubCategory  string `json:"account_sub_category"`
		ThaiEmail           string `json:"thai_email"`
		ThaiEmail2          string `json:"thai_email2"`
		ThaiEmail3          string `json:"thai_email3"`
		StatusCd            string `json:"status_cd"`
		BirthDate           string `json:"birth_date"`
		StatusDt            string `json:"status_dt"`
		RegisterDt          string `json:"register_dt"`
		AddressID           string `json:"address_id"`
		CreatedAt           string `json:"created_at"`
		CreatedBy           string `json:"created_by"`
		UpdatedAt           string `json:"updated_at"`
		UpdatedBy           string `json:"updated_by"`
		Reason              string `json:"reason"`
		TelNo               string `json:"tel_no"`
		NameOnDocumentTh    string `json:"name_on_document_th"`
		NameOnDocumentEng   string `json:"name_on_document_eng"`
		BlockchainFlg       string `json:"blockchain_flg"`
		HasEmployeeDetail   struct {
			ID         string `json:"id"`
			Email      string `json:"email"`
			EmployeeId string `json:"employee_id"`
			MobileNo   string `json:"mobile_no"`
			AccountId  string `json:"account_id"`
			BizId      string `json:"biz_id"`
			CreatedAt  string `json:"created_at"`
			CreatedBy  string `json:"created_by"`
			UpdatedAt  string `json:"updated_at"`
			UpdatedBy  string `json:"updated_by"`
			DeletedAt  string `json:"deleted_at"`
			NickName   string `json:"nick_name"`
			ProfileImg string `json:"profile_img"`
		} `json:"has_employee_detail"`
		HasDepartment []struct {
			Id               string      `json:"id"`
			AccountId        string      `json:"account_id"`
			BizId            string      `json:"biz_id"`
			DeptRoleId       string      `json:"dept_role_id"`
			DeptId           string      `json:"dept_id"`
			RoleId           string      `json:"role_id"`
			Status           string      `json:"status"`
			SpId             string      `json:"sp_id"`
			StartAt          interface{} `json:"start_at"`
			ExpiredAt        interface{} `json:"expired_at"`
			CreatedAt        string      `json:"created_at"`
			CreatedBy        string      `json:"created_by"`
			UpdatedAt        string      `json:"updated_at"`
			UpdatedBy        string      `json:"updated_by"`
			DeletedAt        interface{} `json:"deleted_at"`
			Department       struct {
				Id           string      `json:"id"`
				DeptName     string      `json:"dept_name"`
				DeptPosition string      `json:"dept_position"`
				ParentDeptId string      `json:"parent_dept_id"`
				CreatedAt    string      `json:"created_at"`
				CreatedBy    string      `json:"created_by"`
				UpdatedAt    string      `json:"updated_at"`
				UpdatedBy    string      `json:"updated_by"`
				DeletedAt    interface{} `json:"deleted_at"`
			} `json:"department"`
			DeptRole   *struct {
				Id               string      `json:"id"`
				BizId            string      `json:"biz_id"`
				DeptId           string      `json:"dept_id"`
				RoleId           string      `json:"role_id"`
				DeptRolePosition string      `json:"dept_role_position"`
				DeptRoleLevel    string      `json:"dept_role_level"`
				ParentDeptRoleId *string     `json:"parent_dept_role_id"`
				Status           string      `json:"status"`
				CreatedAt        string      `json:"created_at"`
				CreatedBy        string      `json:"created_by"`
				UpdatedAt        string      `json:"updated_at"`
				UpdatedBy        string      `json:"updated_by"`
				DeletedAt        interface{} `json:"deleted_at"`
			} `json:"dept_role"`
		} `json:"has_department"`
		HasRole      []struct {
			Id               string      `json:"id"`
			AccountId        string      `json:"account_id"`
			BizId            string      `json:"biz_id"`
			DeptRoleId       string      `json:"dept_role_id"`
			DeptId           string      `json:"dept_id"`
			RoleId           string      `json:"role_id"`
			Status           string      `json:"status"`
			SpId             string      `json:"sp_id"`
			StartAt          interface{} `json:"start_at"`
			ExpiredAt        interface{} `json:"expired_at"`
			CreatedAt        string      `json:"created_at"`
			CreatedBy        string      `json:"created_by"`
			UpdatedAt        string      `json:"updated_at"`
			UpdatedBy        string      `json:"updated_by"`
			DeletedAt        interface{} `json:"deleted_at"`
			Role struct {
				Id        string      `json:"id"`
				RoleName  string      `json:"role_name"`
				RoleLevel string      `json:"role_level"`
				CreatedAt string      `json:"created_at"`
				CreatedBy string      `json:"created_by"`
				UpdatedAt string      `json:"updated_at"`
				UpdatedBy string      `json:"updated_by"`
				DeletedAt interface{} `json:"deleted_at"`
			} `json:"role"`
			Department       struct {
				Id           string      `json:"id"`
				DeptName     string      `json:"dept_name"`
				DeptPosition string      `json:"dept_position"`
				ParentDeptId string      `json:"parent_dept_id"`
				CreatedAt    string      `json:"created_at"`
				CreatedBy    string      `json:"created_by"`
				UpdatedAt    string      `json:"updated_at"`
				UpdatedBy    string      `json:"updated_by"`
				DeletedAt    interface{} `json:"deleted_at"`
			} `json:"department"`
			DeptRole   *struct {
				Id               string      `json:"id"`
				BizId            string      `json:"biz_id"`
				DeptId           string      `json:"dept_id"`
				RoleId           string      `json:"role_id"`
				DeptRolePosition string      `json:"dept_role_position"`
				DeptRoleLevel    string      `json:"dept_role_level"`
				ParentDeptRoleId *string     `json:"parent_dept_role_id"`
				Status           string      `json:"status"`
				CreatedAt        string      `json:"created_at"`
				CreatedBy        string      `json:"created_by"`
				UpdatedAt        string      `json:"updated_at"`
				UpdatedBy        string      `json:"updated_by"`
				DeletedAt        interface{} `json:"deleted_at"`
			} `json:"dept_role"`
			Permission []interface{} `json:"permission"`
		} `json:"has_role"`
		HasPermission []interface{} `json:"has_permission"`
	} `json:"data"`
	Errormessage string      `json:"errorMessage"`
	Responsecode int64       `json:"responseCode"`
}
