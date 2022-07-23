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

// NftHolders model.NftHolders.description
type NftHolders struct {
	// field.totalNum.description
	TotalNum int32 `json:"totalNum"`
	// model.nftHolders.description
	NftHolders []NftHolder `json:"nftHolders"`
}

// NewNftHolders instantiates a new NftHolders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftHolders(totalNum int32, nftHolders []NftHolder) *NftHolders {
	this := NftHolders{}
	this.TotalNum = totalNum
	this.NftHolders = nftHolders
	return &this
}

// NewNftHoldersWithDefaults instantiates a new NftHolders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftHoldersWithDefaults() *NftHolders {
	this := NftHolders{}
	return &this
}

// GetTotalNum returns the TotalNum field value
func (o *NftHolders) GetTotalNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value
// and a boolean to check if the value has been set.
func (o *NftHolders) GetTotalNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalNum, true
}

// SetTotalNum sets field value
func (o *NftHolders) SetTotalNum(v int32) {
	o.TotalNum = v
}

// GetNftHolders returns the NftHolders field value
func (o *NftHolders) GetNftHolders() []NftHolder {
	if o == nil {
		var ret []NftHolder
		return ret
	}

	return o.NftHolders
}

// GetNftHoldersOk returns a tuple with the NftHolders field value
// and a boolean to check if the value has been set.
func (o *NftHolders) GetNftHoldersOk() ([]NftHolder, bool) {
	if o == nil {
		return nil, false
	}
	return o.NftHolders, true
}

// SetNftHolders sets field value
func (o *NftHolders) SetNftHolders(v []NftHolder) {
	o.NftHolders = v
}

func (o NftHolders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totalNum"] = o.TotalNum
	}
	if true {
		toSerialize["nftHolders"] = o.NftHolders
	}
	return json.Marshal(toSerialize)
}

type NullableNftHolders struct {
	value *NftHolders
	isSet bool
}

func (v NullableNftHolders) Get() *NftHolders {
	return v.value
}

func (v *NullableNftHolders) Set(val *NftHolders) {
	v.value = val
	v.isSet = true
}

func (v NullableNftHolders) IsSet() bool {
	return v.isSet
}

func (v *NullableNftHolders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftHolders(val *NftHolders) *NullableNftHolders {
	return &NullableNftHolders{value: val, isSet: true}
}

func (v NullableNftHolders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftHolders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}