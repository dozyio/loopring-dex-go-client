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

// GetNextStorageIdResponseData model.GetNextStorageIdResponseData.description
type GetNextStorageIdResponseData struct {
	ResultInfo ResultInfo            `json:"resultInfo"`
	Data       NextStorageIdResponse `json:"data"`
}

// NewGetNextStorageIdResponseData instantiates a new GetNextStorageIdResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNextStorageIdResponseData(resultInfo ResultInfo, data NextStorageIdResponse) *GetNextStorageIdResponseData {
	this := GetNextStorageIdResponseData{}
	this.ResultInfo = resultInfo
	this.Data = data
	return &this
}

// NewGetNextStorageIdResponseDataWithDefaults instantiates a new GetNextStorageIdResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNextStorageIdResponseDataWithDefaults() *GetNextStorageIdResponseData {
	this := GetNextStorageIdResponseData{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *GetNextStorageIdResponseData) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *GetNextStorageIdResponseData) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *GetNextStorageIdResponseData) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value
func (o *GetNextStorageIdResponseData) GetData() NextStorageIdResponse {
	if o == nil {
		var ret NextStorageIdResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetNextStorageIdResponseData) GetDataOk() (*NextStorageIdResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetNextStorageIdResponseData) SetData(v NextStorageIdResponse) {
	o.Data = v
}

func (o GetNextStorageIdResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetNextStorageIdResponseData struct {
	value *GetNextStorageIdResponseData
	isSet bool
}

func (v NullableGetNextStorageIdResponseData) Get() *GetNextStorageIdResponseData {
	return v.value
}

func (v *NullableGetNextStorageIdResponseData) Set(val *GetNextStorageIdResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNextStorageIdResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNextStorageIdResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNextStorageIdResponseData(val *GetNextStorageIdResponseData) *NullableGetNextStorageIdResponseData {
	return &NullableGetNextStorageIdResponseData{value: val, isSet: true}
}

func (v NullableGetNextStorageIdResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNextStorageIdResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}