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

// NftOffChainWithdrawalRequestV3 model.NftOffChainWithdrawalRequestV3.description
type NftOffChainWithdrawalRequestV3 struct {
	// exchange address
	Exchange string `json:"exchange"`
	// account ID
	AccountId int64 `json:"accountId"`
	// account owner address
	Owner  string             `json:"owner"`
	Token  NftTokenAmountInfo `json:"token"`
	MaxFee TokenAmountInfo    `json:"maxFee"`
	// offchain ID
	StorageId int64 `json:"storageId"`
	// Timestamp for order to become invalid
	ValidUntil int32 `json:"validUntil"`
	// min gas for on-chain withdraw, Loopring exchange allocates gas for each distribution, but people can also assign this min gas, so Loopring have to allocate higher gas value for this specific distribution, 0 means let loopring choose the reasonable gas
	MinGas *int32 `json:"minGas,omitempty"`
	// withdraw to address
	To string `json:"to"`
	// extra data for complex withdraw mode, normally none
	ExtraData          *string             `json:"extraData,omitempty"`
	CounterFactualInfo *CounterFactualInfo `json:"counterFactualInfo,omitempty"`
	// eddsa signature
	EddsaSignature *string `json:"eddsaSignature,omitempty"`
	// ecdsa signature
	EcdsaSignature *string `json:"ecdsaSignature,omitempty"`
	// An approved hash string which was already submitted on eth mainnet
	HashApproved *string `json:"hashApproved,omitempty"`
}

// NewNftOffChainWithdrawalRequestV3 instantiates a new NftOffChainWithdrawalRequestV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftOffChainWithdrawalRequestV3(exchange string, accountId int64, owner string, token NftTokenAmountInfo, maxFee TokenAmountInfo, storageId int64, validUntil int32, to string) *NftOffChainWithdrawalRequestV3 {
	this := NftOffChainWithdrawalRequestV3{}
	this.Exchange = exchange
	this.AccountId = accountId
	this.Owner = owner
	this.Token = token
	this.MaxFee = maxFee
	this.StorageId = storageId
	this.ValidUntil = validUntil
	this.To = to
	return &this
}

// NewNftOffChainWithdrawalRequestV3WithDefaults instantiates a new NftOffChainWithdrawalRequestV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftOffChainWithdrawalRequestV3WithDefaults() *NftOffChainWithdrawalRequestV3 {
	this := NftOffChainWithdrawalRequestV3{}
	return &this
}

// GetExchange returns the Exchange field value
func (o *NftOffChainWithdrawalRequestV3) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *NftOffChainWithdrawalRequestV3) SetExchange(v string) {
	o.Exchange = v
}

// GetAccountId returns the AccountId field value
func (o *NftOffChainWithdrawalRequestV3) GetAccountId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetAccountIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NftOffChainWithdrawalRequestV3) SetAccountId(v int64) {
	o.AccountId = v
}

// GetOwner returns the Owner field value
func (o *NftOffChainWithdrawalRequestV3) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *NftOffChainWithdrawalRequestV3) SetOwner(v string) {
	o.Owner = v
}

// GetToken returns the Token field value
func (o *NftOffChainWithdrawalRequestV3) GetToken() NftTokenAmountInfo {
	if o == nil {
		var ret NftTokenAmountInfo
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetTokenOk() (*NftTokenAmountInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *NftOffChainWithdrawalRequestV3) SetToken(v NftTokenAmountInfo) {
	o.Token = v
}

// GetMaxFee returns the MaxFee field value
func (o *NftOffChainWithdrawalRequestV3) GetMaxFee() TokenAmountInfo {
	if o == nil {
		var ret TokenAmountInfo
		return ret
	}

	return o.MaxFee
}

// GetMaxFeeOk returns a tuple with the MaxFee field value
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetMaxFeeOk() (*TokenAmountInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxFee, true
}

// SetMaxFee sets field value
func (o *NftOffChainWithdrawalRequestV3) SetMaxFee(v TokenAmountInfo) {
	o.MaxFee = v
}

// GetStorageId returns the StorageId field value
func (o *NftOffChainWithdrawalRequestV3) GetStorageId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetStorageIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageId, true
}

// SetStorageId sets field value
func (o *NftOffChainWithdrawalRequestV3) SetStorageId(v int64) {
	o.StorageId = v
}

// GetValidUntil returns the ValidUntil field value
func (o *NftOffChainWithdrawalRequestV3) GetValidUntil() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetValidUntilOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidUntil, true
}

// SetValidUntil sets field value
func (o *NftOffChainWithdrawalRequestV3) SetValidUntil(v int32) {
	o.ValidUntil = v
}

// GetMinGas returns the MinGas field value if set, zero value otherwise.
func (o *NftOffChainWithdrawalRequestV3) GetMinGas() int32 {
	if o == nil || o.MinGas == nil {
		var ret int32
		return ret
	}
	return *o.MinGas
}

// GetMinGasOk returns a tuple with the MinGas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetMinGasOk() (*int32, bool) {
	if o == nil || o.MinGas == nil {
		return nil, false
	}
	return o.MinGas, true
}

// HasMinGas returns a boolean if a field has been set.
func (o *NftOffChainWithdrawalRequestV3) HasMinGas() bool {
	if o != nil && o.MinGas != nil {
		return true
	}

	return false
}

// SetMinGas gets a reference to the given int32 and assigns it to the MinGas field.
func (o *NftOffChainWithdrawalRequestV3) SetMinGas(v int32) {
	o.MinGas = &v
}

// GetTo returns the To field value
func (o *NftOffChainWithdrawalRequestV3) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *NftOffChainWithdrawalRequestV3) SetTo(v string) {
	o.To = v
}

// GetExtraData returns the ExtraData field value if set, zero value otherwise.
func (o *NftOffChainWithdrawalRequestV3) GetExtraData() string {
	if o == nil || o.ExtraData == nil {
		var ret string
		return ret
	}
	return *o.ExtraData
}

// GetExtraDataOk returns a tuple with the ExtraData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetExtraDataOk() (*string, bool) {
	if o == nil || o.ExtraData == nil {
		return nil, false
	}
	return o.ExtraData, true
}

// HasExtraData returns a boolean if a field has been set.
func (o *NftOffChainWithdrawalRequestV3) HasExtraData() bool {
	if o != nil && o.ExtraData != nil {
		return true
	}

	return false
}

// SetExtraData gets a reference to the given string and assigns it to the ExtraData field.
func (o *NftOffChainWithdrawalRequestV3) SetExtraData(v string) {
	o.ExtraData = &v
}

// GetCounterFactualInfo returns the CounterFactualInfo field value if set, zero value otherwise.
func (o *NftOffChainWithdrawalRequestV3) GetCounterFactualInfo() CounterFactualInfo {
	if o == nil || o.CounterFactualInfo == nil {
		var ret CounterFactualInfo
		return ret
	}
	return *o.CounterFactualInfo
}

// GetCounterFactualInfoOk returns a tuple with the CounterFactualInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetCounterFactualInfoOk() (*CounterFactualInfo, bool) {
	if o == nil || o.CounterFactualInfo == nil {
		return nil, false
	}
	return o.CounterFactualInfo, true
}

// HasCounterFactualInfo returns a boolean if a field has been set.
func (o *NftOffChainWithdrawalRequestV3) HasCounterFactualInfo() bool {
	if o != nil && o.CounterFactualInfo != nil {
		return true
	}

	return false
}

// SetCounterFactualInfo gets a reference to the given CounterFactualInfo and assigns it to the CounterFactualInfo field.
func (o *NftOffChainWithdrawalRequestV3) SetCounterFactualInfo(v CounterFactualInfo) {
	o.CounterFactualInfo = &v
}

// GetEddsaSignature returns the EddsaSignature field value if set, zero value otherwise.
func (o *NftOffChainWithdrawalRequestV3) GetEddsaSignature() string {
	if o == nil || o.EddsaSignature == nil {
		var ret string
		return ret
	}
	return *o.EddsaSignature
}

// GetEddsaSignatureOk returns a tuple with the EddsaSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetEddsaSignatureOk() (*string, bool) {
	if o == nil || o.EddsaSignature == nil {
		return nil, false
	}
	return o.EddsaSignature, true
}

// HasEddsaSignature returns a boolean if a field has been set.
func (o *NftOffChainWithdrawalRequestV3) HasEddsaSignature() bool {
	if o != nil && o.EddsaSignature != nil {
		return true
	}

	return false
}

// SetEddsaSignature gets a reference to the given string and assigns it to the EddsaSignature field.
func (o *NftOffChainWithdrawalRequestV3) SetEddsaSignature(v string) {
	o.EddsaSignature = &v
}

// GetEcdsaSignature returns the EcdsaSignature field value if set, zero value otherwise.
func (o *NftOffChainWithdrawalRequestV3) GetEcdsaSignature() string {
	if o == nil || o.EcdsaSignature == nil {
		var ret string
		return ret
	}
	return *o.EcdsaSignature
}

// GetEcdsaSignatureOk returns a tuple with the EcdsaSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetEcdsaSignatureOk() (*string, bool) {
	if o == nil || o.EcdsaSignature == nil {
		return nil, false
	}
	return o.EcdsaSignature, true
}

// HasEcdsaSignature returns a boolean if a field has been set.
func (o *NftOffChainWithdrawalRequestV3) HasEcdsaSignature() bool {
	if o != nil && o.EcdsaSignature != nil {
		return true
	}

	return false
}

// SetEcdsaSignature gets a reference to the given string and assigns it to the EcdsaSignature field.
func (o *NftOffChainWithdrawalRequestV3) SetEcdsaSignature(v string) {
	o.EcdsaSignature = &v
}

// GetHashApproved returns the HashApproved field value if set, zero value otherwise.
func (o *NftOffChainWithdrawalRequestV3) GetHashApproved() string {
	if o == nil || o.HashApproved == nil {
		var ret string
		return ret
	}
	return *o.HashApproved
}

// GetHashApprovedOk returns a tuple with the HashApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftOffChainWithdrawalRequestV3) GetHashApprovedOk() (*string, bool) {
	if o == nil || o.HashApproved == nil {
		return nil, false
	}
	return o.HashApproved, true
}

// HasHashApproved returns a boolean if a field has been set.
func (o *NftOffChainWithdrawalRequestV3) HasHashApproved() bool {
	if o != nil && o.HashApproved != nil {
		return true
	}

	return false
}

// SetHashApproved gets a reference to the given string and assigns it to the HashApproved field.
func (o *NftOffChainWithdrawalRequestV3) SetHashApproved(v string) {
	o.HashApproved = &v
}

func (o NftOffChainWithdrawalRequestV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["exchange"] = o.Exchange
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["owner"] = o.Owner
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
	if o.MinGas != nil {
		toSerialize["minGas"] = o.MinGas
	}
	if true {
		toSerialize["to"] = o.To
	}
	if o.ExtraData != nil {
		toSerialize["extraData"] = o.ExtraData
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
	return json.Marshal(toSerialize)
}

type NullableNftOffChainWithdrawalRequestV3 struct {
	value *NftOffChainWithdrawalRequestV3
	isSet bool
}

func (v NullableNftOffChainWithdrawalRequestV3) Get() *NftOffChainWithdrawalRequestV3 {
	return v.value
}

func (v *NullableNftOffChainWithdrawalRequestV3) Set(val *NftOffChainWithdrawalRequestV3) {
	v.value = val
	v.isSet = true
}

func (v NullableNftOffChainWithdrawalRequestV3) IsSet() bool {
	return v.isSet
}

func (v *NullableNftOffChainWithdrawalRequestV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftOffChainWithdrawalRequestV3(val *NftOffChainWithdrawalRequestV3) *NullableNftOffChainWithdrawalRequestV3 {
	return &NullableNftOffChainWithdrawalRequestV3{value: val, isSet: true}
}

func (v NullableNftOffChainWithdrawalRequestV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftOffChainWithdrawalRequestV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
