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

// SubmitOrderResponseItemV3 Submit order response detail
type SubmitOrderResponseItemV3 struct {
	// Order hash of submit order response
	Hash string `json:"hash"`
	// The clientOrderId of the submitted order
	ClientOrderId string `json:"clientOrderId"`
	// Order status of submit order response
	Status string `json:"status"`
	// Idempotent of submit order response, submit same order again when order was UNKNOWN or WAIT_FREEZE_BALANCE in relayer, idempotent will be true
	IsIdempotent bool `json:"isIdempotent"`
	// field.SubmitOffChainResponseItem.accountId
	AccountId int64 `json:"accountId"`
	// field.SubmitOffChainResponseItem.tokenId
	Tokens []map[string]interface{} `json:"tokens"`
	// field.SubmitOffChainResponseItem.storageId
	StorageId int64 `json:"storageId"`
}

// NewSubmitOrderResponseItemV3 instantiates a new SubmitOrderResponseItemV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitOrderResponseItemV3(hash string, clientOrderId string, status string, isIdempotent bool, accountId int64, tokens []map[string]interface{}, storageId int64) *SubmitOrderResponseItemV3 {
	this := SubmitOrderResponseItemV3{}
	this.Hash = hash
	this.ClientOrderId = clientOrderId
	this.Status = status
	this.IsIdempotent = isIdempotent
	this.AccountId = accountId
	this.Tokens = tokens
	this.StorageId = storageId
	return &this
}

// NewSubmitOrderResponseItemV3WithDefaults instantiates a new SubmitOrderResponseItemV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitOrderResponseItemV3WithDefaults() *SubmitOrderResponseItemV3 {
	this := SubmitOrderResponseItemV3{}
	return &this
}

// GetHash returns the Hash field value
func (o *SubmitOrderResponseItemV3) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderResponseItemV3) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *SubmitOrderResponseItemV3) SetHash(v string) {
	o.Hash = v
}

// GetClientOrderId returns the ClientOrderId field value
func (o *SubmitOrderResponseItemV3) GetClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderResponseItemV3) GetClientOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientOrderId, true
}

// SetClientOrderId sets field value
func (o *SubmitOrderResponseItemV3) SetClientOrderId(v string) {
	o.ClientOrderId = v
}

// GetStatus returns the Status field value
func (o *SubmitOrderResponseItemV3) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderResponseItemV3) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SubmitOrderResponseItemV3) SetStatus(v string) {
	o.Status = v
}

// GetIsIdempotent returns the IsIdempotent field value
func (o *SubmitOrderResponseItemV3) GetIsIdempotent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsIdempotent
}

// GetIsIdempotentOk returns a tuple with the IsIdempotent field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderResponseItemV3) GetIsIdempotentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsIdempotent, true
}

// SetIsIdempotent sets field value
func (o *SubmitOrderResponseItemV3) SetIsIdempotent(v bool) {
	o.IsIdempotent = v
}

// GetAccountId returns the AccountId field value
func (o *SubmitOrderResponseItemV3) GetAccountId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderResponseItemV3) GetAccountIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *SubmitOrderResponseItemV3) SetAccountId(v int64) {
	o.AccountId = v
}

// GetTokens returns the Tokens field value
func (o *SubmitOrderResponseItemV3) GetTokens() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderResponseItemV3) GetTokensOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tokens, true
}

// SetTokens sets field value
func (o *SubmitOrderResponseItemV3) SetTokens(v []map[string]interface{}) {
	o.Tokens = v
}

// GetStorageId returns the StorageId field value
func (o *SubmitOrderResponseItemV3) GetStorageId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderResponseItemV3) GetStorageIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageId, true
}

// SetStorageId sets field value
func (o *SubmitOrderResponseItemV3) SetStorageId(v int64) {
	o.StorageId = v
}

func (o SubmitOrderResponseItemV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["isIdempotent"] = o.IsIdempotent
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["tokens"] = o.Tokens
	}
	if true {
		toSerialize["storageId"] = o.StorageId
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitOrderResponseItemV3 struct {
	value *SubmitOrderResponseItemV3
	isSet bool
}

func (v NullableSubmitOrderResponseItemV3) Get() *SubmitOrderResponseItemV3 {
	return v.value
}

func (v *NullableSubmitOrderResponseItemV3) Set(val *SubmitOrderResponseItemV3) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitOrderResponseItemV3) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitOrderResponseItemV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitOrderResponseItemV3(val *SubmitOrderResponseItemV3) *NullableSubmitOrderResponseItemV3 {
	return &NullableSubmitOrderResponseItemV3{value: val, isSet: true}
}

func (v NullableSubmitOrderResponseItemV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitOrderResponseItemV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}