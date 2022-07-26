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

// GetTimestampV2Response model.getTimestampV2Response
type GetTimestampV2Response struct {
	ResultInfo ResultInfo `json:"resultInfo"`
	// field.getTimestampV2Response.data
	Data *int64 `json:"data,omitempty"`
}

// NewGetTimestampV2Response instantiates a new GetTimestampV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimestampV2Response(resultInfo ResultInfo) *GetTimestampV2Response {
	this := GetTimestampV2Response{}
	this.ResultInfo = resultInfo
	return &this
}

// NewGetTimestampV2ResponseWithDefaults instantiates a new GetTimestampV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimestampV2ResponseWithDefaults() *GetTimestampV2Response {
	this := GetTimestampV2Response{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *GetTimestampV2Response) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *GetTimestampV2Response) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *GetTimestampV2Response) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetTimestampV2Response) GetData() int64 {
	if o == nil || o.Data == nil {
		var ret int64
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimestampV2Response) GetDataOk() (*int64, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetTimestampV2Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given int64 and assigns it to the Data field.
func (o *GetTimestampV2Response) SetData(v int64) {
	o.Data = &v
}

func (o GetTimestampV2Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetTimestampV2Response struct {
	value *GetTimestampV2Response
	isSet bool
}

func (v NullableGetTimestampV2Response) Get() *GetTimestampV2Response {
	return v.value
}

func (v *NullableGetTimestampV2Response) Set(val *GetTimestampV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimestampV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimestampV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimestampV2Response(val *GetTimestampV2Response) *NullableGetTimestampV2Response {
	return &NullableGetTimestampV2Response{value: val, isSet: true}
}

func (v NullableGetTimestampV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimestampV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
