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

// CancelOrdersResponseData struct for CancelOrdersResponseData
type CancelOrdersResponseData struct {
	ResultInfo ResultInfo `json:"resultInfo"`
	// field.CancelOrdersResponseData.data
	Data *bool `json:"data,omitempty"`
}

// NewCancelOrdersResponseData instantiates a new CancelOrdersResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelOrdersResponseData(resultInfo ResultInfo) *CancelOrdersResponseData {
	this := CancelOrdersResponseData{}
	this.ResultInfo = resultInfo
	return &this
}

// NewCancelOrdersResponseDataWithDefaults instantiates a new CancelOrdersResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelOrdersResponseDataWithDefaults() *CancelOrdersResponseData {
	this := CancelOrdersResponseData{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *CancelOrdersResponseData) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *CancelOrdersResponseData) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *CancelOrdersResponseData) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CancelOrdersResponseData) GetData() bool {
	if o == nil || o.Data == nil {
		var ret bool
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrdersResponseData) GetDataOk() (*bool, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CancelOrdersResponseData) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given bool and assigns it to the Data field.
func (o *CancelOrdersResponseData) SetData(v bool) {
	o.Data = &v
}

func (o CancelOrdersResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCancelOrdersResponseData struct {
	value *CancelOrdersResponseData
	isSet bool
}

func (v NullableCancelOrdersResponseData) Get() *CancelOrdersResponseData {
	return v.value
}

func (v *NullableCancelOrdersResponseData) Set(val *CancelOrdersResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelOrdersResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelOrdersResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelOrdersResponseData(val *CancelOrdersResponseData) *NullableCancelOrdersResponseData {
	return &NullableCancelOrdersResponseData{value: val, isSet: true}
}

func (v NullableCancelOrdersResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelOrdersResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
