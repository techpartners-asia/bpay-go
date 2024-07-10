package bpaygo

type (

	// login request and response
	BpayLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	BpayLoginResponse struct {
		BpayResponse
		Data BpayLoginData `json:"data"`
	}
	BpayLoginData struct {
		TokenType       string `json:"tokenType"`       //
		RefreshToken    string `json:"refreshToken"`    //
		ExpiresIn       int64  `json:"expiresIn"`       //
		AccessToken     string `json:"accessToken"`     //
		UserId          int64  `json:"userId"`          //
		RoleId          int64  `json:"roleId"`          //
		JTI             int64  `json:"jti"`             //
		PaymentMethodID int64  `json:"paymentMethodId"` //
		Username        string `json:"username"`        //
	}

	// Customer request and response
	BpayCustomerRegisterRequest struct {
		UserID string `json:"userId"`
		Email  string `json:"email"`
	}
	BpayCustomerRegisterResponse struct {
		BpayResponse
		Data string `json:"data"` // Хэрэглэгчийн BPAY рүү нэвтрэхэд таних код
	}

	BpayCustomerLoginRequest struct {
		UserID   string `json:"userId"`
		BpayCOde string `json:"bpayCode"`
	}
	BpayCustomerLoginResponse struct {
		BpayResponse
	}
	BpayCustomerCheckRequest struct {
		UserID string `json:"userId"`
	}
	BpayCustomerCheckResponse struct {
		BpayResponse
		Data string `json:"data"` // Хэрэглэгчийн BPAY рүү нэвтрэхэд таних код``
	}

	BpayConstantData struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	// Group request and response
	BpayGroupCreateRequest struct {
		Name string `json:"name"`
	}
	BpayGroupCreateResponse struct {
		BpayResponse
	}
	BpayGroupEditRequest struct {
		Name string `json:"name"`
	}
	BpayGroupEditResponse struct {
		BpayResponse
	}

	BpayGroupListRequest struct {
		PageNo  int64             `json:"pageNo"`
		PerPage int64             `json:"perPage"`
		Sort    string            `json:"sort"`
		FIlter  []BpayGroupFilter `json:"filter"`
	}
	BpayGroupFilter struct {
		FieldName string `json:"fieldName"`
		Operation string `json:"operation"`
		Value     string `json:"value"`
		FieldType string `json:"fieldType"`
	}

	BpayGroupListResponse struct {
		BpayResponse
		Data []BpayGroupData `json:"data"`
	}
	BpayGroupData struct {
		ID         int64  `json:"id"`
		Name       string `json:"name"`
		CustomerID int64  `json:"customerId"`
	}
	BpayGroupAddBillsRequest struct {
		BillIds []int64 `json:"billIds"`
	}
	BpayGroupAddBillsResponse struct {
		BpayResponse
	}

	BpayGroupBillsResponse struct {
		BpayResponse
		Data []BpayBillData `json:"data"`
	}

	// Find request and response

	BpayFindAddressResponse struct {
		BpayResponse
		Data []BpayAddressData `json:"data"`
	}
	BpayAddressData struct {
		Name    string `gorm:"column:name" json:"name"`
		CID     string `gorm:"column:cid" json:"cid"`
		Address string `gorm:"column:address" json:"address"`
		Count   int64  `gorm:"column:count" json:"count"`
	}

	BpayFindResponse struct {
		BpayResponse
		Data []BpayFindData `json:"data"`
	}
	BpayFindData struct {
		ID          int64          `json:"id"`
		Name        string         `json:"name"`
		Code        string         `json:"code"`
		TotalAmount float64        `json:"totalAmount"`
		ProviderID  int64          `json:"providerId"`
		BIlls       []BpayBillData `json:"bills"`
	}
	BpayBillData struct {
		ID          int64   `gorm:"column:id" json:"id"`
		BillID      string  `gorm:"column:bill_id" json:"billId"`
		Code        string  `gorm:"column:code" json:"code"`                // Хэрэглэгчийн CID код
		BillAmount  float64 `gorm:"column:bill_amount" json:"billAmount"`   // Төлөх дүн
		LossAmount  float64 `gorm:"column:loss_amount" json:"lossAmount"`   // Алдангийн дүн
		TotalAmount float64 `gorm:"column:total_amount" json:"totalAmount"` // Нэхэмжилсэн дүн
		PaidAmount  float64 `gorm:"column:paid_amount" json:"paidAmount"`   // Төлбөл зохих дүн
		Year        int64   `gorm:"column:year" json:"year"`
		Month       int64   `gorm:"column:month" json:"month"`
		Name        string  `gorm:"column:name" json:"name"`
		OrgTypeID   int64   `gorm:"column:org_type_id" json:"orgTypeId"`
		OrgName     string  `gorm:"column:org_name" json:"orgName"`
		ProviderID  int64   `gorm:"column:provider_id" json:"providerId"`
		CustomerID  int64   `gorm:"column:customer_id" json:"customerId"`
		StatusID    int64   `gorm:"column:status_id" json:"statusId"`
	}

	// Invoice request and response
	BpayInvoiceCreateRequest struct {
		BillIDs []int64 `json:"billIds"`
	}
	BpayInvoiceResponse struct {
		BpayResponse
		ID          int64          `json:"id"`
		TotalAmount float64        `json:"totalAmount"`
		CustomerID  int64          `json:"customerId"`
		StatusID    int64          `json:"statusId"`
		BIlls       []BpayBillData `json:"bills"`
	}

	BpayInvoiceTransactionCreateRequest struct {
		InvoiceID int64  `json:"invoiceId"`
		IsOrg     bool   `json:"isOrg"`
		VatInfo   string `json:"vatInfo"` // Company Register
	}
	BpayInvoiceTransactionCreateResponse struct {
		BpayResponse
		InvoiceID    string        `json:"invoiceId"`
		QrText       string        `json:"qr_text"`
		QrImage      string        `json:"qr_image"`
		QpayShrotUrl string        `json:"qPay_shortUrl"`
		Urls         []BpayUrlData `json:"urls"`
	}
	BpayUrlData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Logo        string `json:"logo"`
		Link        string `json:"link"`
	}

	BpayBillCheckResponse struct {
		BpayResponse
		Status       string `json:"status"`
		StatusCode   Status `json:"statusCode"`
		StatusSystem string `json:"statusSystem"`
	}

	BpayResponse struct {
		ResponseCode bool   `json:"responseCode"`
		ResponseMsg  string `json:"responseMsg"`
	}
)
