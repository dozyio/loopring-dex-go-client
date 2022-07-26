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

// GetMixedMarketsResponseV3 model.GetMarketsResponseV3.description
type GetMixedMarketsResponseV3 struct {
	// Markets list
	Markets []CombineMarketInfo `json:"markets"`
}

// NewGetMixedMarketsResponseV3 instantiates a new GetMixedMarketsResponseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMixedMarketsResponseV3(markets []CombineMarketInfo) *GetMixedMarketsResponseV3 {
	this := GetMixedMarketsResponseV3{}
	this.Markets = markets
	return &this
}

// NewGetMixedMarketsResponseV3WithDefaults instantiates a new GetMixedMarketsResponseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMixedMarketsResponseV3WithDefaults() *GetMixedMarketsResponseV3 {
	this := GetMixedMarketsResponseV3{}
	return &this
}

// GetMarkets returns the Markets field value
func (o *GetMixedMarketsResponseV3) GetMarkets() []CombineMarketInfo {
	if o == nil {
		var ret []CombineMarketInfo
		return ret
	}

	return o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value
// and a boolean to check if the value has been set.
func (o *GetMixedMarketsResponseV3) GetMarketsOk() ([]CombineMarketInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Markets, true
}

// SetMarkets sets field value
func (o *GetMixedMarketsResponseV3) SetMarkets(v []CombineMarketInfo) {
	o.Markets = v
}

func (o GetMixedMarketsResponseV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["markets"] = o.Markets
	}
	return json.Marshal(toSerialize)
}

type NullableGetMixedMarketsResponseV3 struct {
	value *GetMixedMarketsResponseV3
	isSet bool
}

func (v NullableGetMixedMarketsResponseV3) Get() *GetMixedMarketsResponseV3 {
	return v.value
}

func (v *NullableGetMixedMarketsResponseV3) Set(val *GetMixedMarketsResponseV3) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMixedMarketsResponseV3) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMixedMarketsResponseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMixedMarketsResponseV3(val *GetMixedMarketsResponseV3) *NullableGetMixedMarketsResponseV3 {
	return &NullableGetMixedMarketsResponseV3{value: val, isSet: true}
}

func (v NullableGetMixedMarketsResponseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMixedMarketsResponseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
