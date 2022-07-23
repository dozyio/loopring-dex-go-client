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

// GetTickerResponseV2 model.getTicker.description
type GetTickerResponseV2 struct {
	ResultInfo ResultInfo `json:"resultInfo"`
	// Each item in the list is an array that contains the following: trading pair ID, update timestamp, base token volume, quote token volume, open-price, highest price, lowest price, closing price, number of trades, highest bid price, lowest ask price, base fee amount, quote fee amount. All values are returned as strings. Fee amount is for AMM only.
	Data [][]string `json:"data,omitempty"`
}

// NewGetTickerResponseV2 instantiates a new GetTickerResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTickerResponseV2(resultInfo ResultInfo) *GetTickerResponseV2 {
	this := GetTickerResponseV2{}
	this.ResultInfo = resultInfo
	return &this
}

// NewGetTickerResponseV2WithDefaults instantiates a new GetTickerResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTickerResponseV2WithDefaults() *GetTickerResponseV2 {
	this := GetTickerResponseV2{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *GetTickerResponseV2) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *GetTickerResponseV2) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *GetTickerResponseV2) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetTickerResponseV2) GetData() [][]string {
	if o == nil || o.Data == nil {
		var ret [][]string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerResponseV2) GetDataOk() ([][]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetTickerResponseV2) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given [][]string and assigns it to the Data field.
func (o *GetTickerResponseV2) SetData(v [][]string) {
	o.Data = v
}

func (o GetTickerResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetTickerResponseV2 struct {
	value *GetTickerResponseV2
	isSet bool
}

func (v NullableGetTickerResponseV2) Get() *GetTickerResponseV2 {
	return v.value
}

func (v *NullableGetTickerResponseV2) Set(val *GetTickerResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTickerResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTickerResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTickerResponseV2(val *GetTickerResponseV2) *NullableGetTickerResponseV2 {
	return &NullableGetTickerResponseV2{value: val, isSet: true}
}

func (v NullableGetTickerResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTickerResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}