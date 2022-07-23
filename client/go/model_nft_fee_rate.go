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

// NftFeeRate The data structure of user query market fee rate
type NftFeeRate struct {
	// field.NftFeeRate.nftTokenAddress
	NftTokenAddress string `json:"nftTokenAddress"`
	// field.NftFeeRate.rate
	Rate int32 `json:"rate"`
}

// NewNftFeeRate instantiates a new NftFeeRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftFeeRate(nftTokenAddress string, rate int32) *NftFeeRate {
	this := NftFeeRate{}
	this.NftTokenAddress = nftTokenAddress
	this.Rate = rate
	return &this
}

// NewNftFeeRateWithDefaults instantiates a new NftFeeRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftFeeRateWithDefaults() *NftFeeRate {
	this := NftFeeRate{}
	return &this
}

// GetNftTokenAddress returns the NftTokenAddress field value
func (o *NftFeeRate) GetNftTokenAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NftTokenAddress
}

// GetNftTokenAddressOk returns a tuple with the NftTokenAddress field value
// and a boolean to check if the value has been set.
func (o *NftFeeRate) GetNftTokenAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NftTokenAddress, true
}

// SetNftTokenAddress sets field value
func (o *NftFeeRate) SetNftTokenAddress(v string) {
	o.NftTokenAddress = v
}

// GetRate returns the Rate field value
func (o *NftFeeRate) GetRate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *NftFeeRate) GetRateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *NftFeeRate) SetRate(v int32) {
	o.Rate = v
}

func (o NftFeeRate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nftTokenAddress"] = o.NftTokenAddress
	}
	if true {
		toSerialize["rate"] = o.Rate
	}
	return json.Marshal(toSerialize)
}

type NullableNftFeeRate struct {
	value *NftFeeRate
	isSet bool
}

func (v NullableNftFeeRate) Get() *NftFeeRate {
	return v.value
}

func (v *NullableNftFeeRate) Set(val *NftFeeRate) {
	v.value = val
	v.isSet = true
}

func (v NullableNftFeeRate) IsSet() bool {
	return v.isSet
}

func (v *NullableNftFeeRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftFeeRate(val *NftFeeRate) *NullableNftFeeRate {
	return &NullableNftFeeRate{value: val, isSet: true}
}

func (v NullableNftFeeRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftFeeRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}