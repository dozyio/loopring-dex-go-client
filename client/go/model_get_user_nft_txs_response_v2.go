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

// GetUserNftTxsResponseV2 model.GetUserNftTxsResponseV2.description
type GetUserNftTxsResponseV2 struct {
	ResultInfo ResultInfo            `json:"resultInfo"`
	Data       GetUserNftTxsResponse `json:"data"`
}

// NewGetUserNftTxsResponseV2 instantiates a new GetUserNftTxsResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserNftTxsResponseV2(resultInfo ResultInfo, data GetUserNftTxsResponse) *GetUserNftTxsResponseV2 {
	this := GetUserNftTxsResponseV2{}
	this.ResultInfo = resultInfo
	this.Data = data
	return &this
}

// NewGetUserNftTxsResponseV2WithDefaults instantiates a new GetUserNftTxsResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserNftTxsResponseV2WithDefaults() *GetUserNftTxsResponseV2 {
	this := GetUserNftTxsResponseV2{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *GetUserNftTxsResponseV2) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *GetUserNftTxsResponseV2) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *GetUserNftTxsResponseV2) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value
func (o *GetUserNftTxsResponseV2) GetData() GetUserNftTxsResponse {
	if o == nil {
		var ret GetUserNftTxsResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetUserNftTxsResponseV2) GetDataOk() (*GetUserNftTxsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetUserNftTxsResponseV2) SetData(v GetUserNftTxsResponse) {
	o.Data = v
}

func (o GetUserNftTxsResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetUserNftTxsResponseV2 struct {
	value *GetUserNftTxsResponseV2
	isSet bool
}

func (v NullableGetUserNftTxsResponseV2) Get() *GetUserNftTxsResponseV2 {
	return v.value
}

func (v *NullableGetUserNftTxsResponseV2) Set(val *GetUserNftTxsResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserNftTxsResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserNftTxsResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserNftTxsResponseV2(val *GetUserNftTxsResponseV2) *NullableGetUserNftTxsResponseV2 {
	return &NullableGetUserNftTxsResponseV2{value: val, isSet: true}
}

func (v NullableGetUserNftTxsResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserNftTxsResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
