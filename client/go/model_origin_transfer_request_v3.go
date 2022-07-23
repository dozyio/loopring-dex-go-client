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

// OriginTransferRequestV3 Submit internal transfer params
type OriginTransferRequestV3 struct {
	// exchange address
	Exchange string `json:"exchange"`
	// payer account ID
	PayerId int64 `json:"payerId"`
	// payer account address
	PayerAddr string `json:"payerAddr"`
	// payee account ID
	PayeeId int64 `json:"payeeId"`
	// payer account address
	PayeeAddr string        `json:"payeeAddr"`
	Token     TokenVolumeV3 `json:"token"`
	MaxFee    TokenVolumeV3 `json:"maxFee"`
	// offchain Id
	StorageId int64 `json:"storageId"`
	// Timestamp for order to become invalid
	ValidUntil         int32               `json:"validUntil"`
	CounterFactualInfo *CounterFactualInfo `json:"counterFactualInfo,omitempty"`
	// eddsa signature
	EddsaSignature *string `json:"eddsaSignature,omitempty"`
	// ecdsa signature
	EcdsaSignature *string `json:"ecdsaSignature,omitempty"`
	// An approved hash string which was already submitted on eth mainnet
	HashApproved *string `json:"hashApproved,omitempty"`
	// transfer memo
	Memo *string `json:"memo,omitempty"`
	// A user-defined id
	ClientId *string `json:"clientId,omitempty"`
	// field.OriginTransferRequestV3.payPayeeUpdateAccount
	PayPayeeUpdateAccount *bool `json:"payPayeeUpdateAccount,omitempty"`
}

// NewOriginTransferRequestV3 instantiates a new OriginTransferRequestV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginTransferRequestV3(exchange string, payerId int64, payerAddr string, payeeId int64, payeeAddr string, token TokenVolumeV3, maxFee TokenVolumeV3, storageId int64, validUntil int32) *OriginTransferRequestV3 {
	this := OriginTransferRequestV3{}
	this.Exchange = exchange
	this.PayerId = payerId
	this.PayerAddr = payerAddr
	this.PayeeId = payeeId
	this.PayeeAddr = payeeAddr
	this.Token = token
	this.MaxFee = maxFee
	this.StorageId = storageId
	this.ValidUntil = validUntil
	return &this
}

// NewOriginTransferRequestV3WithDefaults instantiates a new OriginTransferRequestV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginTransferRequestV3WithDefaults() *OriginTransferRequestV3 {
	this := OriginTransferRequestV3{}
	return &this
}

// GetExchange returns the Exchange field value
func (o *OriginTransferRequestV3) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *OriginTransferRequestV3) SetExchange(v string) {
	o.Exchange = v
}

// GetPayerId returns the PayerId field value
func (o *OriginTransferRequestV3) GetPayerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PayerId
}

// GetPayerIdOk returns a tuple with the PayerId field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetPayerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayerId, true
}

// SetPayerId sets field value
func (o *OriginTransferRequestV3) SetPayerId(v int64) {
	o.PayerId = v
}

// GetPayerAddr returns the PayerAddr field value
func (o *OriginTransferRequestV3) GetPayerAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayerAddr
}

// GetPayerAddrOk returns a tuple with the PayerAddr field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetPayerAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayerAddr, true
}

// SetPayerAddr sets field value
func (o *OriginTransferRequestV3) SetPayerAddr(v string) {
	o.PayerAddr = v
}

// GetPayeeId returns the PayeeId field value
func (o *OriginTransferRequestV3) GetPayeeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PayeeId
}

// GetPayeeIdOk returns a tuple with the PayeeId field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetPayeeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayeeId, true
}

// SetPayeeId sets field value
func (o *OriginTransferRequestV3) SetPayeeId(v int64) {
	o.PayeeId = v
}

// GetPayeeAddr returns the PayeeAddr field value
func (o *OriginTransferRequestV3) GetPayeeAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayeeAddr
}

// GetPayeeAddrOk returns a tuple with the PayeeAddr field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetPayeeAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayeeAddr, true
}

// SetPayeeAddr sets field value
func (o *OriginTransferRequestV3) SetPayeeAddr(v string) {
	o.PayeeAddr = v
}

// GetToken returns the Token field value
func (o *OriginTransferRequestV3) GetToken() TokenVolumeV3 {
	if o == nil {
		var ret TokenVolumeV3
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetTokenOk() (*TokenVolumeV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *OriginTransferRequestV3) SetToken(v TokenVolumeV3) {
	o.Token = v
}

// GetMaxFee returns the MaxFee field value
func (o *OriginTransferRequestV3) GetMaxFee() TokenVolumeV3 {
	if o == nil {
		var ret TokenVolumeV3
		return ret
	}

	return o.MaxFee
}

// GetMaxFeeOk returns a tuple with the MaxFee field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetMaxFeeOk() (*TokenVolumeV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxFee, true
}

// SetMaxFee sets field value
func (o *OriginTransferRequestV3) SetMaxFee(v TokenVolumeV3) {
	o.MaxFee = v
}

// GetStorageId returns the StorageId field value
func (o *OriginTransferRequestV3) GetStorageId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetStorageIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageId, true
}

// SetStorageId sets field value
func (o *OriginTransferRequestV3) SetStorageId(v int64) {
	o.StorageId = v
}

// GetValidUntil returns the ValidUntil field value
func (o *OriginTransferRequestV3) GetValidUntil() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetValidUntilOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidUntil, true
}

// SetValidUntil sets field value
func (o *OriginTransferRequestV3) SetValidUntil(v int32) {
	o.ValidUntil = v
}

// GetCounterFactualInfo returns the CounterFactualInfo field value if set, zero value otherwise.
func (o *OriginTransferRequestV3) GetCounterFactualInfo() CounterFactualInfo {
	if o == nil || o.CounterFactualInfo == nil {
		var ret CounterFactualInfo
		return ret
	}
	return *o.CounterFactualInfo
}

// GetCounterFactualInfoOk returns a tuple with the CounterFactualInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetCounterFactualInfoOk() (*CounterFactualInfo, bool) {
	if o == nil || o.CounterFactualInfo == nil {
		return nil, false
	}
	return o.CounterFactualInfo, true
}

// HasCounterFactualInfo returns a boolean if a field has been set.
func (o *OriginTransferRequestV3) HasCounterFactualInfo() bool {
	if o != nil && o.CounterFactualInfo != nil {
		return true
	}

	return false
}

// SetCounterFactualInfo gets a reference to the given CounterFactualInfo and assigns it to the CounterFactualInfo field.
func (o *OriginTransferRequestV3) SetCounterFactualInfo(v CounterFactualInfo) {
	o.CounterFactualInfo = &v
}

// GetEddsaSignature returns the EddsaSignature field value if set, zero value otherwise.
func (o *OriginTransferRequestV3) GetEddsaSignature() string {
	if o == nil || o.EddsaSignature == nil {
		var ret string
		return ret
	}
	return *o.EddsaSignature
}

// GetEddsaSignatureOk returns a tuple with the EddsaSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetEddsaSignatureOk() (*string, bool) {
	if o == nil || o.EddsaSignature == nil {
		return nil, false
	}
	return o.EddsaSignature, true
}

// HasEddsaSignature returns a boolean if a field has been set.
func (o *OriginTransferRequestV3) HasEddsaSignature() bool {
	if o != nil && o.EddsaSignature != nil {
		return true
	}

	return false
}

// SetEddsaSignature gets a reference to the given string and assigns it to the EddsaSignature field.
func (o *OriginTransferRequestV3) SetEddsaSignature(v string) {
	o.EddsaSignature = &v
}

// GetEcdsaSignature returns the EcdsaSignature field value if set, zero value otherwise.
func (o *OriginTransferRequestV3) GetEcdsaSignature() string {
	if o == nil || o.EcdsaSignature == nil {
		var ret string
		return ret
	}
	return *o.EcdsaSignature
}

// GetEcdsaSignatureOk returns a tuple with the EcdsaSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetEcdsaSignatureOk() (*string, bool) {
	if o == nil || o.EcdsaSignature == nil {
		return nil, false
	}
	return o.EcdsaSignature, true
}

// HasEcdsaSignature returns a boolean if a field has been set.
func (o *OriginTransferRequestV3) HasEcdsaSignature() bool {
	if o != nil && o.EcdsaSignature != nil {
		return true
	}

	return false
}

// SetEcdsaSignature gets a reference to the given string and assigns it to the EcdsaSignature field.
func (o *OriginTransferRequestV3) SetEcdsaSignature(v string) {
	o.EcdsaSignature = &v
}

// GetHashApproved returns the HashApproved field value if set, zero value otherwise.
func (o *OriginTransferRequestV3) GetHashApproved() string {
	if o == nil || o.HashApproved == nil {
		var ret string
		return ret
	}
	return *o.HashApproved
}

// GetHashApprovedOk returns a tuple with the HashApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetHashApprovedOk() (*string, bool) {
	if o == nil || o.HashApproved == nil {
		return nil, false
	}
	return o.HashApproved, true
}

// HasHashApproved returns a boolean if a field has been set.
func (o *OriginTransferRequestV3) HasHashApproved() bool {
	if o != nil && o.HashApproved != nil {
		return true
	}

	return false
}

// SetHashApproved gets a reference to the given string and assigns it to the HashApproved field.
func (o *OriginTransferRequestV3) SetHashApproved(v string) {
	o.HashApproved = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *OriginTransferRequestV3) GetMemo() string {
	if o == nil || o.Memo == nil {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetMemoOk() (*string, bool) {
	if o == nil || o.Memo == nil {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *OriginTransferRequestV3) HasMemo() bool {
	if o != nil && o.Memo != nil {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *OriginTransferRequestV3) SetMemo(v string) {
	o.Memo = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OriginTransferRequestV3) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OriginTransferRequestV3) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OriginTransferRequestV3) SetClientId(v string) {
	o.ClientId = &v
}

// GetPayPayeeUpdateAccount returns the PayPayeeUpdateAccount field value if set, zero value otherwise.
func (o *OriginTransferRequestV3) GetPayPayeeUpdateAccount() bool {
	if o == nil || o.PayPayeeUpdateAccount == nil {
		var ret bool
		return ret
	}
	return *o.PayPayeeUpdateAccount
}

// GetPayPayeeUpdateAccountOk returns a tuple with the PayPayeeUpdateAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginTransferRequestV3) GetPayPayeeUpdateAccountOk() (*bool, bool) {
	if o == nil || o.PayPayeeUpdateAccount == nil {
		return nil, false
	}
	return o.PayPayeeUpdateAccount, true
}

// HasPayPayeeUpdateAccount returns a boolean if a field has been set.
func (o *OriginTransferRequestV3) HasPayPayeeUpdateAccount() bool {
	if o != nil && o.PayPayeeUpdateAccount != nil {
		return true
	}

	return false
}

// SetPayPayeeUpdateAccount gets a reference to the given bool and assigns it to the PayPayeeUpdateAccount field.
func (o *OriginTransferRequestV3) SetPayPayeeUpdateAccount(v bool) {
	o.PayPayeeUpdateAccount = &v
}

func (o OriginTransferRequestV3) MarshalJSON() ([]byte, error) {
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
		toSerialize["maxFee"] = o.MaxFee
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
	if o.EddsaSignature != nil {
		toSerialize["eddsaSignature"] = o.EddsaSignature
	}
	if o.EcdsaSignature != nil {
		toSerialize["ecdsaSignature"] = o.EcdsaSignature
	}
	if o.HashApproved != nil {
		toSerialize["hashApproved"] = o.HashApproved
	}
	if o.Memo != nil {
		toSerialize["memo"] = o.Memo
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.PayPayeeUpdateAccount != nil {
		toSerialize["payPayeeUpdateAccount"] = o.PayPayeeUpdateAccount
	}
	return json.Marshal(toSerialize)
}

type NullableOriginTransferRequestV3 struct {
	value *OriginTransferRequestV3
	isSet bool
}

func (v NullableOriginTransferRequestV3) Get() *OriginTransferRequestV3 {
	return v.value
}

func (v *NullableOriginTransferRequestV3) Set(val *OriginTransferRequestV3) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginTransferRequestV3) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginTransferRequestV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginTransferRequestV3(val *OriginTransferRequestV3) *NullableOriginTransferRequestV3 {
	return &NullableOriginTransferRequestV3{value: val, isSet: true}
}

func (v NullableOriginTransferRequestV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginTransferRequestV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
