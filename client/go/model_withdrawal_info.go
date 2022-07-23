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

// WithdrawalInfo struct for WithdrawalInfo
type WithdrawalInfo struct {
	Recipient      string `json:"recipient"`
	FastStatus     string `json:"fastStatus"`
	DistributeHash string `json:"distributeHash"`
}

// NewWithdrawalInfo instantiates a new WithdrawalInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawalInfo(recipient string, fastStatus string, distributeHash string) *WithdrawalInfo {
	this := WithdrawalInfo{}
	this.Recipient = recipient
	this.FastStatus = fastStatus
	this.DistributeHash = distributeHash
	return &this
}

// NewWithdrawalInfoWithDefaults instantiates a new WithdrawalInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawalInfoWithDefaults() *WithdrawalInfo {
	this := WithdrawalInfo{}
	return &this
}

// GetRecipient returns the Recipient field value
func (o *WithdrawalInfo) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *WithdrawalInfo) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *WithdrawalInfo) SetRecipient(v string) {
	o.Recipient = v
}

// GetFastStatus returns the FastStatus field value
func (o *WithdrawalInfo) GetFastStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FastStatus
}

// GetFastStatusOk returns a tuple with the FastStatus field value
// and a boolean to check if the value has been set.
func (o *WithdrawalInfo) GetFastStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FastStatus, true
}

// SetFastStatus sets field value
func (o *WithdrawalInfo) SetFastStatus(v string) {
	o.FastStatus = v
}

// GetDistributeHash returns the DistributeHash field value
func (o *WithdrawalInfo) GetDistributeHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DistributeHash
}

// GetDistributeHashOk returns a tuple with the DistributeHash field value
// and a boolean to check if the value has been set.
func (o *WithdrawalInfo) GetDistributeHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistributeHash, true
}

// SetDistributeHash sets field value
func (o *WithdrawalInfo) SetDistributeHash(v string) {
	o.DistributeHash = v
}

func (o WithdrawalInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recipient"] = o.Recipient
	}
	if true {
		toSerialize["fastStatus"] = o.FastStatus
	}
	if true {
		toSerialize["distributeHash"] = o.DistributeHash
	}
	return json.Marshal(toSerialize)
}

type NullableWithdrawalInfo struct {
	value *WithdrawalInfo
	isSet bool
}

func (v NullableWithdrawalInfo) Get() *WithdrawalInfo {
	return v.value
}

func (v *NullableWithdrawalInfo) Set(val *WithdrawalInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawalInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawalInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdrawalInfo(val *WithdrawalInfo) *NullableWithdrawalInfo {
	return &NullableWithdrawalInfo{value: val, isSet: true}
}

func (v NullableWithdrawalInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawalInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}