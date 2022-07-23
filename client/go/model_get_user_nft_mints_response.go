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

// GetUserNftMintsResponse model.GetUserNftMintsResponse.description
type GetUserNftMintsResponse struct {
	// field.totalNum.description
	TotalNum int64 `json:"totalNum"`
	// field.GetUserNftMintsResponse.NftMintDataList
	Mints []NftMintData `json:"mints"`
}

// NewGetUserNftMintsResponse instantiates a new GetUserNftMintsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserNftMintsResponse(totalNum int64, mints []NftMintData) *GetUserNftMintsResponse {
	this := GetUserNftMintsResponse{}
	this.TotalNum = totalNum
	this.Mints = mints
	return &this
}

// NewGetUserNftMintsResponseWithDefaults instantiates a new GetUserNftMintsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserNftMintsResponseWithDefaults() *GetUserNftMintsResponse {
	this := GetUserNftMintsResponse{}
	return &this
}

// GetTotalNum returns the TotalNum field value
func (o *GetUserNftMintsResponse) GetTotalNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value
// and a boolean to check if the value has been set.
func (o *GetUserNftMintsResponse) GetTotalNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalNum, true
}

// SetTotalNum sets field value
func (o *GetUserNftMintsResponse) SetTotalNum(v int64) {
	o.TotalNum = v
}

// GetMints returns the Mints field value
func (o *GetUserNftMintsResponse) GetMints() []NftMintData {
	if o == nil {
		var ret []NftMintData
		return ret
	}

	return o.Mints
}

// GetMintsOk returns a tuple with the Mints field value
// and a boolean to check if the value has been set.
func (o *GetUserNftMintsResponse) GetMintsOk() ([]NftMintData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mints, true
}

// SetMints sets field value
func (o *GetUserNftMintsResponse) SetMints(v []NftMintData) {
	o.Mints = v
}

func (o GetUserNftMintsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totalNum"] = o.TotalNum
	}
	if true {
		toSerialize["mints"] = o.Mints
	}
	return json.Marshal(toSerialize)
}

type NullableGetUserNftMintsResponse struct {
	value *GetUserNftMintsResponse
	isSet bool
}

func (v NullableGetUserNftMintsResponse) Get() *GetUserNftMintsResponse {
	return v.value
}

func (v *NullableGetUserNftMintsResponse) Set(val *GetUserNftMintsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserNftMintsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserNftMintsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserNftMintsResponse(val *GetUserNftMintsResponse) *NullableGetUserNftMintsResponse {
	return &NullableGetUserNftMintsResponse{value: val, isSet: true}
}

func (v NullableGetUserNftMintsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserNftMintsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
