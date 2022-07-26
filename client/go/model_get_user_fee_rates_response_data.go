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

// GetUserFeeRatesResponseData struct for GetUserFeeRatesResponseData
type GetUserFeeRatesResponseData struct {
	FeeRate *FeeRate `json:"feeRate,omitempty"`
	// The gas price use to calculate fee rate
	GasPrice *string `json:"gasPrice,omitempty"`
}

// NewGetUserFeeRatesResponseData instantiates a new GetUserFeeRatesResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserFeeRatesResponseData() *GetUserFeeRatesResponseData {
	this := GetUserFeeRatesResponseData{}
	return &this
}

// NewGetUserFeeRatesResponseDataWithDefaults instantiates a new GetUserFeeRatesResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserFeeRatesResponseDataWithDefaults() *GetUserFeeRatesResponseData {
	this := GetUserFeeRatesResponseData{}
	return &this
}

// GetFeeRate returns the FeeRate field value if set, zero value otherwise.
func (o *GetUserFeeRatesResponseData) GetFeeRate() FeeRate {
	if o == nil || o.FeeRate == nil {
		var ret FeeRate
		return ret
	}
	return *o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserFeeRatesResponseData) GetFeeRateOk() (*FeeRate, bool) {
	if o == nil || o.FeeRate == nil {
		return nil, false
	}
	return o.FeeRate, true
}

// HasFeeRate returns a boolean if a field has been set.
func (o *GetUserFeeRatesResponseData) HasFeeRate() bool {
	if o != nil && o.FeeRate != nil {
		return true
	}

	return false
}

// SetFeeRate gets a reference to the given FeeRate and assigns it to the FeeRate field.
func (o *GetUserFeeRatesResponseData) SetFeeRate(v FeeRate) {
	o.FeeRate = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *GetUserFeeRatesResponseData) GetGasPrice() string {
	if o == nil || o.GasPrice == nil {
		var ret string
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserFeeRatesResponseData) GetGasPriceOk() (*string, bool) {
	if o == nil || o.GasPrice == nil {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *GetUserFeeRatesResponseData) HasGasPrice() bool {
	if o != nil && o.GasPrice != nil {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given string and assigns it to the GasPrice field.
func (o *GetUserFeeRatesResponseData) SetGasPrice(v string) {
	o.GasPrice = &v
}

func (o GetUserFeeRatesResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeeRate != nil {
		toSerialize["feeRate"] = o.FeeRate
	}
	if o.GasPrice != nil {
		toSerialize["gasPrice"] = o.GasPrice
	}
	return json.Marshal(toSerialize)
}

type NullableGetUserFeeRatesResponseData struct {
	value *GetUserFeeRatesResponseData
	isSet bool
}

func (v NullableGetUserFeeRatesResponseData) Get() *GetUserFeeRatesResponseData {
	return v.value
}

func (v *NullableGetUserFeeRatesResponseData) Set(val *GetUserFeeRatesResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserFeeRatesResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserFeeRatesResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserFeeRatesResponseData(val *GetUserFeeRatesResponseData) *NullableGetUserFeeRatesResponseData {
	return &NullableGetUserFeeRatesResponseData{value: val, isSet: true}
}

func (v NullableGetUserFeeRatesResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserFeeRatesResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
