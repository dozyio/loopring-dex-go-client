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

// GetOrderGroupAmountResponse struct for GetOrderGroupAmountResponse
type GetOrderGroupAmountResponse struct {
	ResultInfo ResultInfo               `json:"resultInfo"`
	Data       *GetOrderGroupAmountData `json:"data,omitempty"`
}

// NewGetOrderGroupAmountResponse instantiates a new GetOrderGroupAmountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderGroupAmountResponse(resultInfo ResultInfo) *GetOrderGroupAmountResponse {
	this := GetOrderGroupAmountResponse{}
	this.ResultInfo = resultInfo
	return &this
}

// NewGetOrderGroupAmountResponseWithDefaults instantiates a new GetOrderGroupAmountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderGroupAmountResponseWithDefaults() *GetOrderGroupAmountResponse {
	this := GetOrderGroupAmountResponse{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *GetOrderGroupAmountResponse) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *GetOrderGroupAmountResponse) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *GetOrderGroupAmountResponse) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetOrderGroupAmountResponse) GetData() GetOrderGroupAmountData {
	if o == nil || o.Data == nil {
		var ret GetOrderGroupAmountData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderGroupAmountResponse) GetDataOk() (*GetOrderGroupAmountData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetOrderGroupAmountResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetOrderGroupAmountData and assigns it to the Data field.
func (o *GetOrderGroupAmountResponse) SetData(v GetOrderGroupAmountData) {
	o.Data = &v
}

func (o GetOrderGroupAmountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrderGroupAmountResponse struct {
	value *GetOrderGroupAmountResponse
	isSet bool
}

func (v NullableGetOrderGroupAmountResponse) Get() *GetOrderGroupAmountResponse {
	return v.value
}

func (v *NullableGetOrderGroupAmountResponse) Set(val *GetOrderGroupAmountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderGroupAmountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderGroupAmountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderGroupAmountResponse(val *GetOrderGroupAmountResponse) *NullableGetOrderGroupAmountResponse {
	return &NullableGetOrderGroupAmountResponse{value: val, isSet: true}
}

func (v NullableGetOrderGroupAmountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderGroupAmountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}