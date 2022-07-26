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

// GetOrderAmountResponseData The result of query minimum amount
type GetOrderAmountResponseData struct {
	// Amount
	Amount string `json:"amount"`
	// The gas price use to calculate amount
	GasPrice string `json:"gasPrice"`
}

// NewGetOrderAmountResponseData instantiates a new GetOrderAmountResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderAmountResponseData(amount string, gasPrice string) *GetOrderAmountResponseData {
	this := GetOrderAmountResponseData{}
	this.Amount = amount
	this.GasPrice = gasPrice
	return &this
}

// NewGetOrderAmountResponseDataWithDefaults instantiates a new GetOrderAmountResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderAmountResponseDataWithDefaults() *GetOrderAmountResponseData {
	this := GetOrderAmountResponseData{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetOrderAmountResponseData) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetOrderAmountResponseData) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetOrderAmountResponseData) SetAmount(v string) {
	o.Amount = v
}

// GetGasPrice returns the GasPrice field value
func (o *GetOrderAmountResponseData) GetGasPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *GetOrderAmountResponseData) GetGasPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *GetOrderAmountResponseData) SetGasPrice(v string) {
	o.GasPrice = v
}

func (o GetOrderAmountResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["gasPrice"] = o.GasPrice
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrderAmountResponseData struct {
	value *GetOrderAmountResponseData
	isSet bool
}

func (v NullableGetOrderAmountResponseData) Get() *GetOrderAmountResponseData {
	return v.value
}

func (v *NullableGetOrderAmountResponseData) Set(val *GetOrderAmountResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderAmountResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderAmountResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderAmountResponseData(val *GetOrderAmountResponseData) *NullableGetOrderAmountResponseData {
	return &NullableGetOrderAmountResponseData{value: val, isSet: true}
}

func (v NullableGetOrderAmountResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderAmountResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
