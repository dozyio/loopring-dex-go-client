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

// GasAmountLimitV3 Contains information about the gas amounts required by ETH L1 requests.
type GasAmountLimitV3 struct {
	// The gas amount for withdrawal
	Distribution string `json:"distribution"`
	// The gas amount for deposit
	Deposit string `json:"deposit"`
}

// NewGasAmountLimitV3 instantiates a new GasAmountLimitV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGasAmountLimitV3(distribution string, deposit string) *GasAmountLimitV3 {
	this := GasAmountLimitV3{}
	this.Distribution = distribution
	this.Deposit = deposit
	return &this
}

// NewGasAmountLimitV3WithDefaults instantiates a new GasAmountLimitV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGasAmountLimitV3WithDefaults() *GasAmountLimitV3 {
	this := GasAmountLimitV3{}
	return &this
}

// GetDistribution returns the Distribution field value
func (o *GasAmountLimitV3) GetDistribution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value
// and a boolean to check if the value has been set.
func (o *GasAmountLimitV3) GetDistributionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distribution, true
}

// SetDistribution sets field value
func (o *GasAmountLimitV3) SetDistribution(v string) {
	o.Distribution = v
}

// GetDeposit returns the Deposit field value
func (o *GasAmountLimitV3) GetDeposit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Deposit
}

// GetDepositOk returns a tuple with the Deposit field value
// and a boolean to check if the value has been set.
func (o *GasAmountLimitV3) GetDepositOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deposit, true
}

// SetDeposit sets field value
func (o *GasAmountLimitV3) SetDeposit(v string) {
	o.Deposit = v
}

func (o GasAmountLimitV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["distribution"] = o.Distribution
	}
	if true {
		toSerialize["deposit"] = o.Deposit
	}
	return json.Marshal(toSerialize)
}

type NullableGasAmountLimitV3 struct {
	value *GasAmountLimitV3
	isSet bool
}

func (v NullableGasAmountLimitV3) Get() *GasAmountLimitV3 {
	return v.value
}

func (v *NullableGasAmountLimitV3) Set(val *GasAmountLimitV3) {
	v.value = val
	v.isSet = true
}

func (v NullableGasAmountLimitV3) IsSet() bool {
	return v.isSet
}

func (v *NullableGasAmountLimitV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGasAmountLimitV3(val *GasAmountLimitV3) *NullableGasAmountLimitV3 {
	return &NullableGasAmountLimitV3{value: val, isSet: true}
}

func (v NullableGasAmountLimitV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGasAmountLimitV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}