package bpaygo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/techpartners-asia/bpay-go/utils"
)

var (
	// Login
	BpayLogin = utils.API{
		Url:    "/users/api/v1/user/oauth/token",
		Method: http.MethodPost,
	}

	//Customer
	BpayCustomerRegister = utils.API{
		Url:    "/payment/api/v1/customer/register",
		Method: http.MethodPost,
	}
	BpayCustomerLogin = utils.API{
		Url:    "/payment/api/v1/customer/login",
		Method: http.MethodPost,
	}
	BpayCustomerCheck = utils.API{
		Url:    "/payment/api/v1/customer/check",
		Method: http.MethodPost,
	}

	// Group
	BpayGroupCreate = utils.API{
		Url:    "/payment/api/v1/group/create",
		Method: http.MethodPost,
	}
	BpayGroupEdit = utils.API{
		Url:    "/payment/api/v1/group/update/",
		Method: http.MethodPost,
	}
	BpayGroupList = utils.API{
		Url:    "/payment/api/v1/group/list",
		Method: http.MethodPost,
	}
	BpayGroupAddBills = utils.API{
		Url:    "/payment/api/v1/group/add/bills/",
		Method: http.MethodPost,
	}
	BpayGroupBills = utils.API{
		Url:    "/payment/api/v1/group/bills/",
		Method: http.MethodGet,
	}

	// Constants
	BpayConstantAimagHot = utils.API{
		Url:    "/constant/Constant/aimaghot",
		Method: http.MethodGet,
	}
	BpayConstantSumDuureg = utils.API{
		Url:    "/constant/Constant/sumDuureg/",
		Method: http.MethodGet,
	}
	BpayConstantBagKhoroo = utils.API{
		Url:    "/constant/Constant/khoroo/",
		Method: http.MethodGet,
	}
	BpayConstantBair = utils.API{
		Url:    "/constant/Constant/bair/",
		Method: http.MethodGet,
	}

	// Find
	BpayFindAddress = utils.API{
		Url:    "/search/api/v1/Search/FindAddress",
		Method: http.MethodGet,
	}
	BpayFindCid = utils.API{
		Url:    "/search/api/v1/Search/FindCid?Cid={{cid}}",
		Method: http.MethodGet,
	}
	BpayFindElectric = utils.API{
		Url:    "/search/api/v1/Search/FindElictric?UserId={{userId}}",
		Method: http.MethodGet,
	}
	BpayFindUnivision = utils.API{
		Url:    "/search/api/v1/Search/FindUnivision?Custno={{custNo}}",
		Method: http.MethodGet,
	}
	BpayFindSkymedia = utils.API{
		Url:    "/search/api/v1/Search/FindSkymedia?BillerUserId={{billerUserId}}",
		Method: http.MethodGet,
	}
	BpayFindOnlineBiller = utils.API{
		Url:    "/search/api/v1/Search/FindOnlineBiller?BillerUserId={{billerUserId}}",
		Method: http.MethodGet,
	}

	// Invoice
	BpayCreateInvoice = utils.API{
		Url:    "/payment/api/v1/invoice/create",
		Method: http.MethodPost,
	}
	BpayInvoiceGroupCreate = utils.API{
		Url:    "/payment/api/v1/invoice/group/create/",
		Method: http.MethodGet,
	}
	BpayinvoiceTransactionCreate = utils.API{
		Url:    "/payment/api/v1/invoice/transaction/create",
		Method: http.MethodPost,
	}
	BpayBillCheck = utils.API{
		Url:    "/payment/api/v1/merchant/bill/check/",
		Method: http.MethodPost,
	}
)

func (b *bpay) auth() (authRes BpayLoginData, err error) {
	if b.loginObject != nil {
		expireInA := time.Unix(int64(b.loginObject.ExpiresIn), 0)
		expireInB := expireInA.Add(time.Duration(-12) * time.Hour)
		now := time.Now()
		if now.Before(expireInB) {
			authRes = *b.loginObject
			err = nil
			return
		}
	}

	body := &BpayLoginRequest{
		Username: b.username,
		Password: b.password,
	}
	requestByte, _ := json.Marshal(body)
	requestBody := bytes.NewReader(requestByte)

	url := b.endpoint + BpayLogin.Url
	req, err := http.NewRequest(BpayLogin.Method, url, requestBody)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return authRes, fmt.Errorf("%s-QPay auth response: %s", time.Now().Format(utils.TimeFormatYYYYMMDDHHMMSS), res.Status)
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return authRes, err
	}
	var resp BpayLoginResponse
	if err := json.Unmarshal(responseBody, &resp); err != nil {
		return authRes, err
	}
	authRes = resp.Data
	return authRes, nil
}

func (b *bpay) httpRequest(body interface{}, api utils.API, urlExt string, customerId int) (response []byte, err error) {

	authObj, authErr := b.auth()
	if authErr != nil {
		err = authErr
		return
	}

	b.loginObject = &authObj

	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	req, _ := http.NewRequest(api.Method, b.endpoint+api.Url+urlExt, requestBody)
	if customerId != 0 {
		userIDstr := strconv.Itoa(customerId)
		req.Header.Add("userId", userIDstr)
	}

	req.Header.Add("Content-Type", utils.HttpContent)
	req.Header.Add("Authorization", "Bearer "+b.loginObject.AccessToken)

	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		return nil, errors.New(string(res.Status))
	}
	defer res.Body.Close()
	response, _ = io.ReadAll(res.Body)
	return
}
