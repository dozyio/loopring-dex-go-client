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

// GetUserNftWithdrawalResponse model.GetUserNftWithdrawalResponse.description
type GetUserNftWithdrawalResponse struct {
	// field.totalNum.description
	TotalNum int64 `json:"totalNum"`
	// field.GetUserNftWithdrawalResponse.NftWithdrawalDataList
	Withdrawals []NftWithdrawalData `json:"withdrawals"`
}

// NewGetUserNftWithdrawalResponse instantiates a new GetUserNftWithdrawalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserNftWithdrawalResponse(totalNum int64, withdrawals []NftWithdrawalData) *GetUserNftWithdrawalResponse {
	this := GetUserNftWithdrawalResponse{}
	this.TotalNum = totalNum
	this.Withdrawals = withdrawals
	return &this
}

// NewGetUserNftWithdrawalResponseWithDefaults instantiates a new GetUserNftWithdrawalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserNftWithdrawalResponseWithDefaults() *GetUserNftWithdrawalResponse {
	this := GetUserNftWithdrawalResponse{}
	return &this
}

// GetTotalNum returns the TotalNum field value
func (o *GetUserNftWithdrawalResponse) GetTotalNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value
// and a boolean to check if the value has been set.
func (o *GetUserNftWithdrawalResponse) GetTotalNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalNum, true
}

// SetTotalNum sets field value
func (o *GetUserNftWithdrawalResponse) SetTotalNum(v int64) {
	o.TotalNum = v
}

// GetWithdrawals returns the Withdrawals field value
func (o *GetUserNftWithdrawalResponse) GetWithdrawals() []NftWithdrawalData {
	if o == nil {
		var ret []NftWithdrawalData
		return ret
	}

	return o.Withdrawals
}

// GetWithdrawalsOk returns a tuple with the Withdrawals field value
// and a boolean to check if the value has been set.
func (o *GetUserNftWithdrawalResponse) GetWithdrawalsOk() ([]NftWithdrawalData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Withdrawals, true
}

// SetWithdrawals sets field value
func (o *GetUserNftWithdrawalResponse) SetWithdrawals(v []NftWithdrawalData) {
	o.Withdrawals = v
}

func (o GetUserNftWithdrawalResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totalNum"] = o.TotalNum
	}
	if true {
		toSerialize["withdrawals"] = o.Withdrawals
	}
	return json.Marshal(toSerialize)
}

type NullableGetUserNftWithdrawalResponse struct {
	value *GetUserNftWithdrawalResponse
	isSet bool
}

func (v NullableGetUserNftWithdrawalResponse) Get() *GetUserNftWithdrawalResponse {
	return v.value
}

func (v *NullableGetUserNftWithdrawalResponse) Set(val *GetUserNftWithdrawalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserNftWithdrawalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserNftWithdrawalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserNftWithdrawalResponse(val *GetUserNftWithdrawalResponse) *NullableGetUserNftWithdrawalResponse {
	return &NullableGetUserNftWithdrawalResponse{value: val, isSet: true}
}

func (v NullableGetUserNftWithdrawalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserNftWithdrawalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}