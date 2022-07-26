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

// GetMarketsResponseV3 model.GetMarketsResponseV3.description
type GetMarketsResponseV3 struct {
	// Markets list
	Markets []MarketInfo `json:"markets"`
}

// NewGetMarketsResponseV3 instantiates a new GetMarketsResponseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketsResponseV3(markets []MarketInfo) *GetMarketsResponseV3 {
	this := GetMarketsResponseV3{}
	this.Markets = markets
	return &this
}

// NewGetMarketsResponseV3WithDefaults instantiates a new GetMarketsResponseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketsResponseV3WithDefaults() *GetMarketsResponseV3 {
	this := GetMarketsResponseV3{}
	return &this
}

// GetMarkets returns the Markets field value
func (o *GetMarketsResponseV3) GetMarkets() []MarketInfo {
	if o == nil {
		var ret []MarketInfo
		return ret
	}

	return o.Markets
}

// GetMarketsOk returns a tuple with the Markets field value
// and a boolean to check if the value has been set.
func (o *GetMarketsResponseV3) GetMarketsOk() ([]MarketInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Markets, true
}

// SetMarkets sets field value
func (o *GetMarketsResponseV3) SetMarkets(v []MarketInfo) {
	o.Markets = v
}

func (o GetMarketsResponseV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["markets"] = o.Markets
	}
	return json.Marshal(toSerialize)
}

type NullableGetMarketsResponseV3 struct {
	value *GetMarketsResponseV3
	isSet bool
}

func (v NullableGetMarketsResponseV3) Get() *GetMarketsResponseV3 {
	return v.value
}

func (v *NullableGetMarketsResponseV3) Set(val *GetMarketsResponseV3) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketsResponseV3) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketsResponseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketsResponseV3(val *GetMarketsResponseV3) *NullableGetMarketsResponseV3 {
	return &NullableGetMarketsResponseV3{value: val, isSet: true}
}

func (v NullableGetMarketsResponseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketsResponseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
