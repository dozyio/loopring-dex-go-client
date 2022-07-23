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

// GetTickerResponseV3 model.getTicker.description
type GetTickerResponseV3 struct {
	// Each item in the list is an array that contains the following: trading pair ID, update timestamp, base token volume, quote token volume, open-price, highest price, lowest price, closing price, number of trades, highest bid price, lowest ask price, base fee amount, quote fee amount. All values are returned as strings. Fee amount is for AMM only.
	Tickers [][]string `json:"tickers,omitempty"`
}

// NewGetTickerResponseV3 instantiates a new GetTickerResponseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTickerResponseV3() *GetTickerResponseV3 {
	this := GetTickerResponseV3{}
	return &this
}

// NewGetTickerResponseV3WithDefaults instantiates a new GetTickerResponseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTickerResponseV3WithDefaults() *GetTickerResponseV3 {
	this := GetTickerResponseV3{}
	return &this
}

// GetTickers returns the Tickers field value if set, zero value otherwise.
func (o *GetTickerResponseV3) GetTickers() [][]string {
	if o == nil || o.Tickers == nil {
		var ret [][]string
		return ret
	}
	return o.Tickers
}

// GetTickersOk returns a tuple with the Tickers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerResponseV3) GetTickersOk() ([][]string, bool) {
	if o == nil || o.Tickers == nil {
		return nil, false
	}
	return o.Tickers, true
}

// HasTickers returns a boolean if a field has been set.
func (o *GetTickerResponseV3) HasTickers() bool {
	if o != nil && o.Tickers != nil {
		return true
	}

	return false
}

// SetTickers gets a reference to the given [][]string and assigns it to the Tickers field.
func (o *GetTickerResponseV3) SetTickers(v [][]string) {
	o.Tickers = v
}

func (o GetTickerResponseV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tickers != nil {
		toSerialize["tickers"] = o.Tickers
	}
	return json.Marshal(toSerialize)
}

type NullableGetTickerResponseV3 struct {
	value *GetTickerResponseV3
	isSet bool
}

func (v NullableGetTickerResponseV3) Get() *GetTickerResponseV3 {
	return v.value
}

func (v *NullableGetTickerResponseV3) Set(val *GetTickerResponseV3) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTickerResponseV3) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTickerResponseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTickerResponseV3(val *GetTickerResponseV3) *NullableGetTickerResponseV3 {
	return &NullableGetTickerResponseV3{value: val, isSet: true}
}

func (v NullableGetTickerResponseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTickerResponseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
