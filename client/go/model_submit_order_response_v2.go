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

// SubmitOrderResponseV2 model.SubmitOrderResponseV2.description
type SubmitOrderResponseV2 struct {
	ResultInfo ResultInfo                `json:"resultInfo"`
	Data       SubmitOrderResponseV2Item `json:"data"`
}

// NewSubmitOrderResponseV2 instantiates a new SubmitOrderResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitOrderResponseV2(resultInfo ResultInfo, data SubmitOrderResponseV2Item) *SubmitOrderResponseV2 {
	this := SubmitOrderResponseV2{}
	this.ResultInfo = resultInfo
	this.Data = data
	return &this
}

// NewSubmitOrderResponseV2WithDefaults instantiates a new SubmitOrderResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitOrderResponseV2WithDefaults() *SubmitOrderResponseV2 {
	this := SubmitOrderResponseV2{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *SubmitOrderResponseV2) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderResponseV2) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *SubmitOrderResponseV2) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value
func (o *SubmitOrderResponseV2) GetData() SubmitOrderResponseV2Item {
	if o == nil {
		var ret SubmitOrderResponseV2Item
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderResponseV2) GetDataOk() (*SubmitOrderResponseV2Item, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubmitOrderResponseV2) SetData(v SubmitOrderResponseV2Item) {
	o.Data = v
}

func (o SubmitOrderResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitOrderResponseV2 struct {
	value *SubmitOrderResponseV2
	isSet bool
}

func (v NullableSubmitOrderResponseV2) Get() *SubmitOrderResponseV2 {
	return v.value
}

func (v *NullableSubmitOrderResponseV2) Set(val *SubmitOrderResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitOrderResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitOrderResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitOrderResponseV2(val *SubmitOrderResponseV2) *NullableSubmitOrderResponseV2 {
	return &NullableSubmitOrderResponseV2{value: val, isSet: true}
}

func (v NullableSubmitOrderResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitOrderResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
