package bpaygo

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type bpay struct {
	endpoint    string
	username    string
	password    string
	loginObject *BpayLoginData
}

type Bpay interface {
	CustomerRegister(input BpayCustomerRegisterRequest) (BpayCustomerRegisterResponse, error)
	CustomerLogin(input BpayCustomerLoginRequest) (BpayCustomerLoginResponse, error)
	CustomerCheck(input BpayCustomerCheckRequest) (BpayCustomerCheckResponse, error)

	GroupCreate(input BpayGroupCreateRequest, customerId int) (BpayGroupCreateResponse, error)
	GroupEdit(input BpayGroupEditRequest, id string, customerId int) (BpayGroupEditResponse, error)
	GroupList(input BpayGroupListRequest, customerId int) (BpayGroupListResponse, error)
	GroupAddBills(input BpayGroupAddBillsRequest, id string, customerId int) (BpayGroupAddBillsResponse, error)
	GroupBills(id string, customerId int) (BpayGroupBillsResponse, error)

	ConstantAimagHot() ([]BpayConstantData, error)
	ConstantSumDuureg(aimagHotId int) ([]BpayConstantData, error)
	ConstantBagKhoroo(aimagHotId, sumDuuregId int) ([]BpayConstantData, error)
	ConstantBair(aimagHotId, sumDuuregId, bagKhorooId int) ([]BpayConstantData, error)

	FindAddress(aimagId, sumId, khorooId, bairNum, haalgaNum, customerId int) (BpayFindAddressResponse, error)
	FindCid(cid string, customerId int) (BpayFindResponse, error)
	FindElectric(userId string, customerId int) (BpayFindResponse, error)
	FindUnivision(custNo string, customerId int) (BpayFindResponse, error)
	FindSkymedia(billerUserId string, customerId int) (BpayFindResponse, error)
	FindOnlineBiller(billerUserId string, customerId int) (BpayFindResponse, error)

	InvoiceCreate(input BpayInvoiceCreateRequest, customerId int) (BpayInvoiceResponse, error)
	InvoiceGroupCreate(groupId string, customerId int) (BpayInvoiceResponse, error)
	InvoiceTransactionCreate(input BpayInvoiceTransactionCreateRequest, customerId int) (BpayInvoiceTransactionCreateResponse, error)
	BillCheck(invoiceId string) (BpayBillCheckResponse, error)
}

func New(endpoint, username, password string) Bpay {
	return &bpay{
		endpoint:    endpoint,
		username:    username,
		password:    password,
		loginObject: nil,
	}
}

func (b *bpay) CustomerRegister(input BpayCustomerRegisterRequest) (BpayCustomerRegisterResponse, error) {
	res, err := b.httpRequest(input, BpayCustomerRegister, "", 0)
	if err != nil {
		return BpayCustomerRegisterResponse{}, err
	}
	var response BpayCustomerRegisterResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayCustomerRegisterResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) CustomerLogin(input BpayCustomerLoginRequest) (BpayCustomerLoginResponse, error) {
	res, err := b.httpRequest(input, BpayCustomerLogin, "", 0)
	if err != nil {
		return BpayCustomerLoginResponse{}, err
	}
	var response BpayCustomerLoginResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayCustomerLoginResponse{}, errors.New(response.ResponseMsg)
	}

	return response, nil
}

func (b *bpay) CustomerCheck(input BpayCustomerCheckRequest) (BpayCustomerCheckResponse, error) {
	res, err := b.httpRequest(input, BpayCustomerCheck, "", 0)
	if err != nil {
		return BpayCustomerCheckResponse{}, err
	}
	var response BpayCustomerCheckResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayCustomerCheckResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

// Group
func (b *bpay) GroupCreate(input BpayGroupCreateRequest, customerId int) (BpayGroupCreateResponse, error) {
	res, err := b.httpRequest(input, BpayGroupCreate, "", customerId)
	if err != nil {
		return BpayGroupCreateResponse{}, err
	}
	var response BpayGroupCreateResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayGroupCreateResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) GroupEdit(input BpayGroupEditRequest, id string, customerId int) (BpayGroupEditResponse, error) {
	res, err := b.httpRequest(input, BpayGroupEdit, id, customerId)
	if err != nil {
		return BpayGroupEditResponse{}, err
	}
	var response BpayGroupEditResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayGroupEditResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) GroupList(input BpayGroupListRequest, customerId int) (BpayGroupListResponse, error) {
	res, err := b.httpRequest(input, BpayGroupList, "", customerId)
	if err != nil {
		return BpayGroupListResponse{}, err
	}
	var response BpayGroupListResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayGroupListResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) GroupAddBills(input BpayGroupAddBillsRequest, id string, customerId int) (BpayGroupAddBillsResponse, error) {
	res, err := b.httpRequest(input, BpayGroupAddBills, id, customerId)
	if err != nil {
		return BpayGroupAddBillsResponse{}, err
	}
	var response BpayGroupAddBillsResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayGroupAddBillsResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) GroupBills(id string, customerId int) (BpayGroupBillsResponse, error) {
	res, err := b.httpRequest(nil, BpayGroupBills, id, customerId)
	if err != nil {
		return BpayGroupBillsResponse{}, err
	}
	var response BpayGroupBillsResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayGroupBillsResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

// Constants
func (b *bpay) ConstantAimagHot() ([]BpayConstantData, error) {
	res, err := b.httpRequest(nil, BpayConstantAimagHot, "", 0)
	if err != nil {
		return nil, err
	}
	var response []BpayConstantData
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (b *bpay) ConstantSumDuureg(aimagHotId int) ([]BpayConstantData, error) {
	aimagHotIdstr := strconv.Itoa(aimagHotId)
	res, err := b.httpRequest(nil, BpayConstantSumDuureg, aimagHotIdstr, 0)
	if err != nil {
		return nil, err
	}
	var response []BpayConstantData
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (b *bpay) ConstantBagKhoroo(aimagHotId, sumDuuregId int) ([]BpayConstantData, error) {
	aimagHotIdstr := strconv.Itoa(aimagHotId)
	sumDuuregIdstr := strconv.Itoa(sumDuuregId)
	res, err := b.httpRequest(nil, BpayConstantBagKhoroo, aimagHotIdstr+"/"+sumDuuregIdstr, 0)
	if err != nil {
		return nil, err
	}
	var response []BpayConstantData
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (b *bpay) ConstantBair(aimagHotId, sumDuuregId, bagKhorooId int) ([]BpayConstantData, error) {
	aimagHotIdstr := strconv.Itoa(aimagHotId)
	sumDuuregIdstr := strconv.Itoa(sumDuuregId)
	bagKhorooIdstr := strconv.Itoa(bagKhorooId)
	res, err := b.httpRequest(nil, BpayConstantBair, aimagHotIdstr+"/"+sumDuuregIdstr+"/"+bagKhorooIdstr, 0)
	if err != nil {
		return nil, err
	}
	var response []BpayConstantData
	if err := json.Unmarshal(res, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Find

func (b *bpay) FindAddress(aimagId, sumId, khorooId, bairNum, haalgaNum, customerId int) (BpayFindAddressResponse, error) {
	query := fmt.Sprintf("?AimagId=%d&SumId=%d&KhorooId=%d&BairNum=%d&XaalgaNum=%d", aimagId, sumId, khorooId, bairNum, haalgaNum)
	res, err := b.httpRequest(nil, BpayFindAddress, query, customerId)
	if err != nil {
		return BpayFindAddressResponse{}, err
	}
	var response BpayFindAddressResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayFindAddressResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) FindCid(cId string, customerId int) (BpayFindResponse, error) {
	res, err := b.httpRequest(nil, BpayFindCid, "Cid="+cId, customerId)
	if err != nil {
		return BpayFindResponse{}, err
	}
	var response BpayFindResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayFindResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) FindElectric(userId string, customerId int) (BpayFindResponse, error) {
	res, err := b.httpRequest(nil, BpayFindElectric, "UserId="+userId, customerId)
	if err != nil {
		return BpayFindResponse{}, err
	}
	var response BpayFindResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayFindResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) FindUnivision(custNo string, customerId int) (BpayFindResponse, error) {
	res, err := b.httpRequest(nil, BpayFindElectric, "Custno="+custNo, customerId)
	if err != nil {
		return BpayFindResponse{}, err
	}
	var response BpayFindResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayFindResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) FindSkymedia(billerUserId string, customerId int) (BpayFindResponse, error) {
	res, err := b.httpRequest(nil, BpayFindSkymedia, "BillerUserId="+billerUserId, customerId)
	if err != nil {
		return BpayFindResponse{}, err
	}
	var response BpayFindResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayFindResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) FindOnlineBiller(billerUserId string, customerId int) (BpayFindResponse, error) {
	res, err := b.httpRequest(nil, BpayFindOnlineBiller, "BillerUserId="+billerUserId, customerId)
	if err != nil {
		return BpayFindResponse{}, err
	}
	var response BpayFindResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayFindResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

// Invoice
func (b *bpay) InvoiceCreate(input BpayInvoiceCreateRequest, customerId int) (BpayInvoiceResponse, error) {
	res, err := b.httpRequest(input, BpayCreateInvoice, "", customerId)
	if err != nil {
		return BpayInvoiceResponse{}, err
	}
	var response BpayInvoiceResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayInvoiceResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) InvoiceGroupCreate(groupId string, customerId int) (BpayInvoiceResponse, error) {
	res, err := b.httpRequest(nil, BpayInvoiceGroupCreate, groupId, customerId)
	if err != nil {
		return BpayInvoiceResponse{}, err
	}
	var response BpayInvoiceResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayInvoiceResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) InvoiceTransactionCreate(input BpayInvoiceTransactionCreateRequest, customerId int) (BpayInvoiceTransactionCreateResponse, error) {
	res, err := b.httpRequest(input, BpayinvoiceTransactionCreate, "", customerId)
	if err != nil {
		return BpayInvoiceTransactionCreateResponse{}, err
	}
	var response BpayInvoiceTransactionCreateResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayInvoiceTransactionCreateResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}

func (b *bpay) BillCheck(invoiceId string) (BpayBillCheckResponse, error) {
	res, err := b.httpRequest(nil, BpayBillCheck, invoiceId, 0)
	if err != nil {
		return BpayBillCheckResponse{}, err
	}
	var response BpayBillCheckResponse
	json.Unmarshal(res, &response)
	if !response.ResponseCode {
		return BpayBillCheckResponse{}, errors.New(response.ResponseMsg)
	}
	return response, nil
}
