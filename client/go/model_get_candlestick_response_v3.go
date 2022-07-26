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

// GetCandlestickResponseV3 The response of query Candlestick data
type GetCandlestickResponseV3 struct {
	// Candlestick data, each set of data includes start time, number of transactions, opening price, closing price, highest price, lowest price, total transaction volume of Base Token, total transaction volume of Quote Token
	Candlesticks [][]string `json:"candlesticks,omitempty"`
}

// NewGetCandlestickResponseV3 instantiates a new GetCandlestickResponseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCandlestickResponseV3() *GetCandlestickResponseV3 {
	this := GetCandlestickResponseV3{}
	return &this
}

// NewGetCandlestickResponseV3WithDefaults instantiates a new GetCandlestickResponseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCandlestickResponseV3WithDefaults() *GetCandlestickResponseV3 {
	this := GetCandlestickResponseV3{}
	return &this
}

// GetCandlesticks returns the Candlesticks field value if set, zero value otherwise.
func (o *GetCandlestickResponseV3) GetCandlesticks() [][]string {
	if o == nil || o.Candlesticks == nil {
		var ret [][]string
		return ret
	}
	return o.Candlesticks
}

// GetCandlesticksOk returns a tuple with the Candlesticks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCandlestickResponseV3) GetCandlesticksOk() ([][]string, bool) {
	if o == nil || o.Candlesticks == nil {
		return nil, false
	}
	return o.Candlesticks, true
}

// HasCandlesticks returns a boolean if a field has been set.
func (o *GetCandlestickResponseV3) HasCandlesticks() bool {
	if o != nil && o.Candlesticks != nil {
		return true
	}

	return false
}

// SetCandlesticks gets a reference to the given [][]string and assigns it to the Candlesticks field.
func (o *GetCandlestickResponseV3) SetCandlesticks(v [][]string) {
	o.Candlesticks = v
}

func (o GetCandlestickResponseV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Candlesticks != nil {
		toSerialize["candlesticks"] = o.Candlesticks
	}
	return json.Marshal(toSerialize)
}

type NullableGetCandlestickResponseV3 struct {
	value *GetCandlestickResponseV3
	isSet bool
}

func (v NullableGetCandlestickResponseV3) Get() *GetCandlestickResponseV3 {
	return v.value
}

func (v *NullableGetCandlestickResponseV3) Set(val *GetCandlestickResponseV3) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCandlestickResponseV3) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCandlestickResponseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCandlestickResponseV3(val *GetCandlestickResponseV3) *NullableGetCandlestickResponseV3 {
	return &NullableGetCandlestickResponseV3{value: val, isSet: true}
}

func (v NullableGetCandlestickResponseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCandlestickResponseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
