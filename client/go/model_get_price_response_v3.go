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

// GetPriceResponseV3 model.GetPriceResponseV3
type GetPriceResponseV3 struct {
	// field.GetPriceResponseV3.prices
	Prices []PriceV3 `json:"prices,omitempty"`
}

// NewGetPriceResponseV3 instantiates a new GetPriceResponseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPriceResponseV3() *GetPriceResponseV3 {
	this := GetPriceResponseV3{}
	return &this
}

// NewGetPriceResponseV3WithDefaults instantiates a new GetPriceResponseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPriceResponseV3WithDefaults() *GetPriceResponseV3 {
	this := GetPriceResponseV3{}
	return &this
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *GetPriceResponseV3) GetPrices() []PriceV3 {
	if o == nil || o.Prices == nil {
		var ret []PriceV3
		return ret
	}
	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceResponseV3) GetPricesOk() ([]PriceV3, bool) {
	if o == nil || o.Prices == nil {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *GetPriceResponseV3) HasPrices() bool {
	if o != nil && o.Prices != nil {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []PriceV3 and assigns it to the Prices field.
func (o *GetPriceResponseV3) SetPrices(v []PriceV3) {
	o.Prices = v
}

func (o GetPriceResponseV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Prices != nil {
		toSerialize["prices"] = o.Prices
	}
	return json.Marshal(toSerialize)
}

type NullableGetPriceResponseV3 struct {
	value *GetPriceResponseV3
	isSet bool
}

func (v NullableGetPriceResponseV3) Get() *GetPriceResponseV3 {
	return v.value
}

func (v *NullableGetPriceResponseV3) Set(val *GetPriceResponseV3) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPriceResponseV3) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPriceResponseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPriceResponseV3(val *GetPriceResponseV3) *NullableGetPriceResponseV3 {
	return &NullableGetPriceResponseV3{value: val, isSet: true}
}

func (v NullableGetPriceResponseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPriceResponseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}