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

// GetAmmTradesResponseV2 struct for GetAmmTradesResponseV2
type GetAmmTradesResponseV2 struct {
	ResultInfo ResultInfo       `json:"resultInfo"`
	Data       AmmTradeDataList `json:"data"`
}

// NewGetAmmTradesResponseV2 instantiates a new GetAmmTradesResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAmmTradesResponseV2(resultInfo ResultInfo, data AmmTradeDataList) *GetAmmTradesResponseV2 {
	this := GetAmmTradesResponseV2{}
	this.ResultInfo = resultInfo
	this.Data = data
	return &this
}

// NewGetAmmTradesResponseV2WithDefaults instantiates a new GetAmmTradesResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAmmTradesResponseV2WithDefaults() *GetAmmTradesResponseV2 {
	this := GetAmmTradesResponseV2{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *GetAmmTradesResponseV2) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *GetAmmTradesResponseV2) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *GetAmmTradesResponseV2) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value
func (o *GetAmmTradesResponseV2) GetData() AmmTradeDataList {
	if o == nil {
		var ret AmmTradeDataList
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetAmmTradesResponseV2) GetDataOk() (*AmmTradeDataList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetAmmTradesResponseV2) SetData(v AmmTradeDataList) {
	o.Data = v
}

func (o GetAmmTradesResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetAmmTradesResponseV2 struct {
	value *GetAmmTradesResponseV2
	isSet bool
}

func (v NullableGetAmmTradesResponseV2) Get() *GetAmmTradesResponseV2 {
	return v.value
}

func (v *NullableGetAmmTradesResponseV2) Set(val *GetAmmTradesResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAmmTradesResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAmmTradesResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAmmTradesResponseV2(val *GetAmmTradesResponseV2) *NullableGetAmmTradesResponseV2 {
	return &NullableGetAmmTradesResponseV2{value: val, isSet: true}
}

func (v NullableGetAmmTradesResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAmmTradesResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}