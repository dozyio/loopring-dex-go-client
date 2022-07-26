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

// GetMarketInfoResponse model.getMarketInfoResponse
type GetMarketInfoResponse struct {
	ResultInfo ResultInfo `json:"resultInfo"`
	// field.getMarketInfoResponse.data
	Data []MarketInfo `json:"data,omitempty"`
}

// NewGetMarketInfoResponse instantiates a new GetMarketInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketInfoResponse(resultInfo ResultInfo) *GetMarketInfoResponse {
	this := GetMarketInfoResponse{}
	this.ResultInfo = resultInfo
	return &this
}

// NewGetMarketInfoResponseWithDefaults instantiates a new GetMarketInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketInfoResponseWithDefaults() *GetMarketInfoResponse {
	this := GetMarketInfoResponse{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *GetMarketInfoResponse) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *GetMarketInfoResponse) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *GetMarketInfoResponse) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetMarketInfoResponse) GetData() []MarketInfo {
	if o == nil || o.Data == nil {
		var ret []MarketInfo
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketInfoResponse) GetDataOk() ([]MarketInfo, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetMarketInfoResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []MarketInfo and assigns it to the Data field.
func (o *GetMarketInfoResponse) SetData(v []MarketInfo) {
	o.Data = v
}

func (o GetMarketInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetMarketInfoResponse struct {
	value *GetMarketInfoResponse
	isSet bool
}

func (v NullableGetMarketInfoResponse) Get() *GetMarketInfoResponse {
	return v.value
}

func (v *NullableGetMarketInfoResponse) Set(val *GetMarketInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketInfoResponse(val *GetMarketInfoResponse) *NullableGetMarketInfoResponse {
	return &NullableGetMarketInfoResponse{value: val, isSet: true}
}

func (v NullableGetMarketInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
