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

// NftTradeList NFT trade list
type NftTradeList struct {
	// field.totalNum.description
	TotalNum int64 `json:"totalNum"`
	// Trade info
	Trades [][]string `json:"trades"`
}

// NewNftTradeList instantiates a new NftTradeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftTradeList(totalNum int64, trades [][]string) *NftTradeList {
	this := NftTradeList{}
	this.TotalNum = totalNum
	this.Trades = trades
	return &this
}

// NewNftTradeListWithDefaults instantiates a new NftTradeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftTradeListWithDefaults() *NftTradeList {
	this := NftTradeList{}
	return &this
}

// GetTotalNum returns the TotalNum field value
func (o *NftTradeList) GetTotalNum() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalNum
}

// GetTotalNumOk returns a tuple with the TotalNum field value
// and a boolean to check if the value has been set.
func (o *NftTradeList) GetTotalNumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalNum, true
}

// SetTotalNum sets field value
func (o *NftTradeList) SetTotalNum(v int64) {
	o.TotalNum = v
}

// GetTrades returns the Trades field value
func (o *NftTradeList) GetTrades() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.Trades
}

// GetTradesOk returns a tuple with the Trades field value
// and a boolean to check if the value has been set.
func (o *NftTradeList) GetTradesOk() ([][]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Trades, true
}

// SetTrades sets field value
func (o *NftTradeList) SetTrades(v [][]string) {
	o.Trades = v
}

func (o NftTradeList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totalNum"] = o.TotalNum
	}
	if true {
		toSerialize["trades"] = o.Trades
	}
	return json.Marshal(toSerialize)
}

type NullableNftTradeList struct {
	value *NftTradeList
	isSet bool
}

func (v NullableNftTradeList) Get() *NftTradeList {
	return v.value
}

func (v *NullableNftTradeList) Set(val *NftTradeList) {
	v.value = val
	v.isSet = true
}

func (v NullableNftTradeList) IsSet() bool {
	return v.isSet
}

func (v *NullableNftTradeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftTradeList(val *NftTradeList) *NullableNftTradeList {
	return &NullableNftTradeList{value: val, isSet: true}
}

func (v NullableNftTradeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftTradeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
