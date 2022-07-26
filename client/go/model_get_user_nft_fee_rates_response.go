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

// GetUserNftFeeRatesResponse The user fee rates result in NFT market
type GetUserNftFeeRatesResponse struct {
	FeeRate *NftFeeRate `json:"feeRate,omitempty"`
	// field.GetUserNftFeeRatesResponse.gasPrice
	GasPrice *string `json:"gasPrice,omitempty"`
}

// NewGetUserNftFeeRatesResponse instantiates a new GetUserNftFeeRatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserNftFeeRatesResponse() *GetUserNftFeeRatesResponse {
	this := GetUserNftFeeRatesResponse{}
	return &this
}

// NewGetUserNftFeeRatesResponseWithDefaults instantiates a new GetUserNftFeeRatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserNftFeeRatesResponseWithDefaults() *GetUserNftFeeRatesResponse {
	this := GetUserNftFeeRatesResponse{}
	return &this
}

// GetFeeRate returns the FeeRate field value if set, zero value otherwise.
func (o *GetUserNftFeeRatesResponse) GetFeeRate() NftFeeRate {
	if o == nil || o.FeeRate == nil {
		var ret NftFeeRate
		return ret
	}
	return *o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserNftFeeRatesResponse) GetFeeRateOk() (*NftFeeRate, bool) {
	if o == nil || o.FeeRate == nil {
		return nil, false
	}
	return o.FeeRate, true
}

// HasFeeRate returns a boolean if a field has been set.
func (o *GetUserNftFeeRatesResponse) HasFeeRate() bool {
	if o != nil && o.FeeRate != nil {
		return true
	}

	return false
}

// SetFeeRate gets a reference to the given NftFeeRate and assigns it to the FeeRate field.
func (o *GetUserNftFeeRatesResponse) SetFeeRate(v NftFeeRate) {
	o.FeeRate = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *GetUserNftFeeRatesResponse) GetGasPrice() string {
	if o == nil || o.GasPrice == nil {
		var ret string
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserNftFeeRatesResponse) GetGasPriceOk() (*string, bool) {
	if o == nil || o.GasPrice == nil {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *GetUserNftFeeRatesResponse) HasGasPrice() bool {
	if o != nil && o.GasPrice != nil {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given string and assigns it to the GasPrice field.
func (o *GetUserNftFeeRatesResponse) SetGasPrice(v string) {
	o.GasPrice = &v
}

func (o GetUserNftFeeRatesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FeeRate != nil {
		toSerialize["feeRate"] = o.FeeRate
	}
	if o.GasPrice != nil {
		toSerialize["gasPrice"] = o.GasPrice
	}
	return json.Marshal(toSerialize)
}

type NullableGetUserNftFeeRatesResponse struct {
	value *GetUserNftFeeRatesResponse
	isSet bool
}

func (v NullableGetUserNftFeeRatesResponse) Get() *GetUserNftFeeRatesResponse {
	return v.value
}

func (v *NullableGetUserNftFeeRatesResponse) Set(val *GetUserNftFeeRatesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserNftFeeRatesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserNftFeeRatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserNftFeeRatesResponse(val *GetUserNftFeeRatesResponse) *NullableGetUserNftFeeRatesResponse {
	return &NullableGetUserNftFeeRatesResponse{value: val, isSet: true}
}

func (v NullableGetUserNftFeeRatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserNftFeeRatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
