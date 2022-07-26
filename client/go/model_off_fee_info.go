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

// OffFeeInfo Off-chain fee info charged by loopring exchange.
type OffFeeInfo struct {
	// fee token
	Token string `json:"token"`
	// fee amount
	Fee string `json:"fee"`
}

// NewOffFeeInfo instantiates a new OffFeeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffFeeInfo(token string, fee string) *OffFeeInfo {
	this := OffFeeInfo{}
	this.Token = token
	this.Fee = fee
	return &this
}

// NewOffFeeInfoWithDefaults instantiates a new OffFeeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffFeeInfoWithDefaults() *OffFeeInfo {
	this := OffFeeInfo{}
	return &this
}

// GetToken returns the Token field value
func (o *OffFeeInfo) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *OffFeeInfo) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *OffFeeInfo) SetToken(v string) {
	o.Token = v
}

// GetFee returns the Fee field value
func (o *OffFeeInfo) GetFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *OffFeeInfo) GetFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *OffFeeInfo) SetFee(v string) {
	o.Fee = v
}

func (o OffFeeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["fee"] = o.Fee
	}
	return json.Marshal(toSerialize)
}

type NullableOffFeeInfo struct {
	value *OffFeeInfo
	isSet bool
}

func (v NullableOffFeeInfo) Get() *OffFeeInfo {
	return v.value
}

func (v *NullableOffFeeInfo) Set(val *OffFeeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOffFeeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOffFeeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffFeeInfo(val *OffFeeInfo) *NullableOffFeeInfo {
	return &NullableOffFeeInfo{value: val, isSet: true}
}

func (v NullableOffFeeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffFeeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
