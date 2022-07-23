/*
LightCone 2.0 API Documentation

LightCone DEX function interpretation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loopring

import (
	"encoding/json"
)

// OriginTransferRequest model.SubmitOriginTransferRequest.description
type OriginTransferRequest struct {
	// field.SubmitOriginTransferRequest.exchange
	Exchange string `json:"exchange"`
	// field.SubmitOriginTransferRequest.payerId
	PayerId int64 `json:"payerId"`
	// field.SubmitOriginTransferRequest.payerAddr
	PayerAddr string `json:"payerAddr"`
	// field.SubmitOriginTransferRequest.payeeId
	PayeeId int64 `json:"payeeId"`
	// field.SubmitOriginTransferRequest.payeeAddr
	PayeeAddr string `json:"payeeAddr"`
	// field.SubmitOriginTransferRequest.token
	Token int32 `json:"token"`
	// field.SubmitOriginTransferRequest.amount
	Amount string `json:"amount"`
	// field.SubmitOriginTransferRequest.feeToken
	FeeToken int32 `json:"feeToken"`
	// field.SubmitOriginTransferRequest.maxFeeAmount
	MaxFeeAmount string `json:"maxFeeAmount"`
	// field.SubmitOriginTransferRequest.storageId
	StorageId int64 `json:"storageId"`
	// field.SubmitOriginTransferRequest.validUntil
	ValidUntil         int32               `json:"validUntil"`
	CounterFactualInfo *CounterFactualInfo `json:"counterFactualInfo,omitempty"`
	// field.SubmitOriginTransferRequest.eddsaSig
	EddsaSig *string `json:"eddsaSig,omitempty"`
	// field.SubmitOriginTransferRequest.ecdsaSig
	EcdsaSig *string `json:"ecdsaSig,omitempty"`
	// field.SubmitOriginTransferRequest.memo
	Memo *string `json:"memo,omitempty"`
	// field.SubmitOriginTransferRequest.clientId
	ClientId *string `json:"clientId,omitempty"`
}

// NewOriginTransferRequest instantiates a new OriginTransferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginTransferRequest(exchange string, payerId int64, payerAddr string, payeeId int64, payeeAddr string, token int32, amount string, feeToken int32, maxFeeAmount string, storageId int64, validUntil int32) *OriginTransferRequest {
	this := OriginTransferRequest{}
	this.Exchange = exchange
	this.PayerId = payerId
	this.PayerAddr = payerAddr
	this.PayeeId = payeeId
	this.PayeeAddr = payeeAddr
	this.Token = token
	this.Amount = amount
	this.FeeToken = feeToken
	this.MaxFeeAmount = maxFeeAmount
	this.StorageId = storageId
	this.ValidUntil = validUntil
	return &this
}

// NewOriginTransferRequestWithDefaults instantiates a new OriginTransferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginTransferRequestWithDefaults() *OriginTransferRequest {
	this := OriginTransferRequest{}
	return &this
}

// GetExchange returns the Exchange field value
func (o *OriginTransferRequest) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *OriginTransferRequest) SetExchange(v string) {
	o.Exchange = v
}

// GetPayerId returns the PayerId field value
func (o *OriginTransferRequest) GetPayerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PayerId
}

// GetPayerIdOk returns a tuple with the PayerId field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetPayerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayerId, true
}

// SetPayerId sets field value
func (o *OriginTransferRequest) SetPayerId(v int64) {
	o.PayerId = v
}

// GetPayerAddr returns the PayerAddr field value
func (o *OriginTransferRequest) GetPayerAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayerAddr
}

// GetPayerAddrOk returns a tuple with the PayerAddr field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetPayerAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayerAddr, true
}

// SetPayerAddr sets field value
func (o *OriginTransferRequest) SetPayerAddr(v string) {
	o.PayerAddr = v
}

// GetPayeeId returns the PayeeId field value
func (o *OriginTransferRequest) GetPayeeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PayeeId
}

// GetPayeeIdOk returns a tuple with the PayeeId field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetPayeeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayeeId, true
}

// SetPayeeId sets field value
func (o *OriginTransferRequest) SetPayeeId(v int64) {
	o.PayeeId = v
}

// GetPayeeAddr returns the PayeeAddr field value
func (o *OriginTransferRequest) GetPayeeAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayeeAddr
}

// GetPayeeAddrOk returns a tuple with the PayeeAddr field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetPayeeAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayeeAddr, true
}

// SetPayeeAddr sets field value
func (o *OriginTransferRequest) SetPayeeAddr(v string) {
	o.PayeeAddr = v
}

// GetToken returns the Token field value
func (o *OriginTransferRequest) GetToken() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetTokenOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *OriginTransferRequest) SetToken(v int32) {
	o.Token = v
}

// GetAmount returns the Amount field value
func (o *OriginTransferRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *OriginTransferRequest) SetAmount(v string) {
	o.Amount = v
}

// GetFeeToken returns the FeeToken field value
func (o *OriginTransferRequest) GetFeeToken() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FeeToken
}

// GetFeeTokenOk returns a tuple with the FeeToken field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetFeeTokenOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeToken, true
}

// SetFeeToken sets field value
func (o *OriginTransferRequest) SetFeeToken(v int32) {
	o.FeeToken = v
}

// GetMaxFeeAmount returns the MaxFeeAmount field value
func (o *OriginTransferRequest) GetMaxFeeAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxFeeAmount
}

// GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetMaxFeeAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxFeeAmount, true
}

// SetMaxFeeAmount sets field value
func (o *OriginTransferRequest) SetMaxFeeAmount(v string) {
	o.MaxFeeAmount = v
}

// GetStorageId returns the StorageId field value
func (o *OriginTransferRequest) GetStorageId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetStorageIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageId, true
}

// SetStorageId sets field value
func (o *OriginTransferRequest) SetStorageId(v int64) {
	o.StorageId = v
}

// GetValidUntil returns the ValidUntil field value
func (o *OriginTransferRequest) GetValidUntil() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetValidUntilOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidUntil, true
}

// SetValidUntil sets field value
func (o *OriginTransferRequest) SetValidUntil(v int32) {
	o.ValidUntil = v
}

// GetCounterFactualInfo returns the CounterFactualInfo field value if set, zero value otherwise.
func (o *OriginTransferRequest) GetCounterFactualInfo() CounterFactualInfo {
	if o == nil || o.CounterFactualInfo == nil {
		var ret CounterFactualInfo
		return ret
	}
	return *o.CounterFactualInfo
}

// GetCounterFactualInfoOk returns a tuple with the CounterFactualInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetCounterFactualInfoOk() (*CounterFactualInfo, bool) {
	if o == nil || o.CounterFactualInfo == nil {
		return nil, false
	}
	return o.CounterFactualInfo, true
}

// HasCounterFactualInfo returns a boolean if a field has been set.
func (o *OriginTransferRequest) HasCounterFactualInfo() bool {
	if o != nil && o.CounterFactualInfo != nil {
		return true
	}

	return false
}

// SetCounterFactualInfo gets a reference to the given CounterFactualInfo and assigns it to the CounterFactualInfo field.
func (o *OriginTransferRequest) SetCounterFactualInfo(v CounterFactualInfo) {
	o.CounterFactualInfo = &v
}

// GetEddsaSig returns the EddsaSig field value if set, zero value otherwise.
func (o *OriginTransferRequest) GetEddsaSig() string {
	if o == nil || o.EddsaSig == nil {
		var ret string
		return ret
	}
	return *o.EddsaSig
}

// GetEddsaSigOk returns a tuple with the EddsaSig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetEddsaSigOk() (*string, bool) {
	if o == nil || o.EddsaSig == nil {
		return nil, false
	}
	return o.EddsaSig, true
}

// HasEddsaSig returns a boolean if a field has been set.
func (o *OriginTransferRequest) HasEddsaSig() bool {
	if o != nil && o.EddsaSig != nil {
		return true
	}

	return false
}

// SetEddsaSig gets a reference to the given string and assigns it to the EddsaSig field.
func (o *OriginTransferRequest) SetEddsaSig(v string) {
	o.EddsaSig = &v
}

// GetEcdsaSig returns the EcdsaSig field value if set, zero value otherwise.
func (o *OriginTransferRequest) GetEcdsaSig() string {
	if o == nil || o.EcdsaSig == nil {
		var ret string
		return ret
	}
	return *o.EcdsaSig
}

// GetEcdsaSigOk returns a tuple with the EcdsaSig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetEcdsaSigOk() (*string, bool) {
	if o == nil || o.EcdsaSig == nil {
		return nil, false
	}
	return o.EcdsaSig, true
}

// HasEcdsaSig returns a boolean if a field has been set.
func (o *OriginTransferRequest) HasEcdsaSig() bool {
	if o != nil && o.EcdsaSig != nil {
		return true
	}

	return false
}

// SetEcdsaSig gets a reference to the given string and assigns it to the EcdsaSig field.
func (o *OriginTransferRequest) SetEcdsaSig(v string) {
	o.EcdsaSig = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *OriginTransferRequest) GetMemo() string {
	if o == nil || o.Memo == nil {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetMemoOk() (*string, bool) {
	if o == nil || o.Memo == nil {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *OriginTransferRequest) HasMemo() bool {
	if o != nil && o.Memo != nil {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *OriginTransferRequest) SetMemo(v string) {
	o.Memo = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OriginTransferRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginTransferRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OriginTransferRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OriginTransferRequest) SetClientId(v string) {
	o.ClientId = &v
}

func (o OriginTransferRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["exchange"] = o.Exchange
	}
	if true {
		toSerialize["payerId"] = o.PayerId
	}
	if true {
		toSerialize["payerAddr"] = o.PayerAddr
	}
	if true {
		toSerialize["payeeId"] = o.PayeeId
	}
	if true {
		toSerialize["payeeAddr"] = o.PayeeAddr
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["feeToken"] = o.FeeToken
	}
	if true {
		toSerialize["maxFeeAmount"] = o.MaxFeeAmount
	}
	if true {
		toSerialize["storageId"] = o.StorageId
	}
	if true {
		toSerialize["validUntil"] = o.ValidUntil
	}
	if o.CounterFactualInfo != nil {
		toSerialize["counterFactualInfo"] = o.CounterFactualInfo
	}
	if o.EddsaSig != nil {
		toSerialize["eddsaSig"] = o.EddsaSig
	}
	if o.EcdsaSig != nil {
		toSerialize["ecdsaSig"] = o.EcdsaSig
	}
	if o.Memo != nil {
		toSerialize["memo"] = o.Memo
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	return json.Marshal(toSerialize)
}

type NullableOriginTransferRequest struct {
	value *OriginTransferRequest
	isSet bool
}

func (v NullableOriginTransferRequest) Get() *OriginTransferRequest {
	return v.value
}

func (v *NullableOriginTransferRequest) Set(val *OriginTransferRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginTransferRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginTransferRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginTransferRequest(val *OriginTransferRequest) *NullableOriginTransferRequest {
	return &NullableOriginTransferRequest{value: val, isSet: true}
}

func (v NullableOriginTransferRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginTransferRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
