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

// AmmPoolPrecisions The precision requirement of a AMM pool
type AmmPoolPrecisions struct {
	// The price precision requirement of a AMM pool
	Price int32 `json:"price"`
	// The amount precision requirement of a AMM pool
	Amount int32 `json:"amount"`
}

// NewAmmPoolPrecisions instantiates a new AmmPoolPrecisions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmmPoolPrecisions(price int32, amount int32) *AmmPoolPrecisions {
	this := AmmPoolPrecisions{}
	this.Price = price
	this.Amount = amount
	return &this
}

// NewAmmPoolPrecisionsWithDefaults instantiates a new AmmPoolPrecisions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmmPoolPrecisionsWithDefaults() *AmmPoolPrecisions {
	this := AmmPoolPrecisions{}
	return &this
}

// GetPrice returns the Price field value
func (o *AmmPoolPrecisions) GetPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *AmmPoolPrecisions) GetPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *AmmPoolPrecisions) SetPrice(v int32) {
	o.Price = v
}

// GetAmount returns the Amount field value
func (o *AmmPoolPrecisions) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AmmPoolPrecisions) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AmmPoolPrecisions) SetAmount(v int32) {
	o.Amount = v
}

func (o AmmPoolPrecisions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableAmmPoolPrecisions struct {
	value *AmmPoolPrecisions
	isSet bool
}

func (v NullableAmmPoolPrecisions) Get() *AmmPoolPrecisions {
	return v.value
}

func (v *NullableAmmPoolPrecisions) Set(val *AmmPoolPrecisions) {
	v.value = val
	v.isSet = true
}

func (v NullableAmmPoolPrecisions) IsSet() bool {
	return v.isSet
}

func (v *NullableAmmPoolPrecisions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmmPoolPrecisions(val *AmmPoolPrecisions) *NullableAmmPoolPrecisions {
	return &NullableAmmPoolPrecisions{value: val, isSet: true}
}

func (v NullableAmmPoolPrecisions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmmPoolPrecisions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
