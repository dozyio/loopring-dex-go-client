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

// SubmitOffChainRequestResponse model.submitOffChainRequest.response
type SubmitOffChainRequestResponse struct {
	ResultInfo ResultInfo                  `json:"resultInfo"`
	Data       *SubmitOffChainResponseItem `json:"data,omitempty"`
}

// NewSubmitOffChainRequestResponse instantiates a new SubmitOffChainRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitOffChainRequestResponse(resultInfo ResultInfo) *SubmitOffChainRequestResponse {
	this := SubmitOffChainRequestResponse{}
	this.ResultInfo = resultInfo
	return &this
}

// NewSubmitOffChainRequestResponseWithDefaults instantiates a new SubmitOffChainRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitOffChainRequestResponseWithDefaults() *SubmitOffChainRequestResponse {
	this := SubmitOffChainRequestResponse{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *SubmitOffChainRequestResponse) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *SubmitOffChainRequestResponse) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *SubmitOffChainRequestResponse) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SubmitOffChainRequestResponse) GetData() SubmitOffChainResponseItem {
	if o == nil || o.Data == nil {
		var ret SubmitOffChainResponseItem
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitOffChainRequestResponse) GetDataOk() (*SubmitOffChainResponseItem, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SubmitOffChainRequestResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given SubmitOffChainResponseItem and assigns it to the Data field.
func (o *SubmitOffChainRequestResponse) SetData(v SubmitOffChainResponseItem) {
	o.Data = &v
}

func (o SubmitOffChainRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitOffChainRequestResponse struct {
	value *SubmitOffChainRequestResponse
	isSet bool
}

func (v NullableSubmitOffChainRequestResponse) Get() *SubmitOffChainRequestResponse {
	return v.value
}

func (v *NullableSubmitOffChainRequestResponse) Set(val *SubmitOffChainRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitOffChainRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitOffChainRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitOffChainRequestResponse(val *SubmitOffChainRequestResponse) *NullableSubmitOffChainRequestResponse {
	return &NullableSubmitOffChainRequestResponse{value: val, isSet: true}
}

func (v NullableSubmitOffChainRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitOffChainRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
