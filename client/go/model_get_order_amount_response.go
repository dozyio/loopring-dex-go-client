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

// GetOrderAmountResponse The results returned by query place order amount
type GetOrderAmountResponse struct {
	ResultInfo ResultInfo                  `json:"resultInfo"`
	Data       *GetOrderAmountResponseData `json:"data,omitempty"`
}

// NewGetOrderAmountResponse instantiates a new GetOrderAmountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderAmountResponse(resultInfo ResultInfo) *GetOrderAmountResponse {
	this := GetOrderAmountResponse{}
	this.ResultInfo = resultInfo
	return &this
}

// NewGetOrderAmountResponseWithDefaults instantiates a new GetOrderAmountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderAmountResponseWithDefaults() *GetOrderAmountResponse {
	this := GetOrderAmountResponse{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *GetOrderAmountResponse) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *GetOrderAmountResponse) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *GetOrderAmountResponse) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetOrderAmountResponse) GetData() GetOrderAmountResponseData {
	if o == nil || o.Data == nil {
		var ret GetOrderAmountResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderAmountResponse) GetDataOk() (*GetOrderAmountResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetOrderAmountResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetOrderAmountResponseData and assigns it to the Data field.
func (o *GetOrderAmountResponse) SetData(v GetOrderAmountResponseData) {
	o.Data = &v
}

func (o GetOrderAmountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrderAmountResponse struct {
	value *GetOrderAmountResponse
	isSet bool
}

func (v NullableGetOrderAmountResponse) Get() *GetOrderAmountResponse {
	return v.value
}

func (v *NullableGetOrderAmountResponse) Set(val *GetOrderAmountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderAmountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderAmountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderAmountResponse(val *GetOrderAmountResponse) *NullableGetOrderAmountResponse {
	return &NullableGetOrderAmountResponse{value: val, isSet: true}
}

func (v NullableGetOrderAmountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderAmountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
