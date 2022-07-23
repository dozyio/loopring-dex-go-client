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

// GetPriceV2Response model.getPriceResponse
type GetPriceV2Response struct {
	ResultInfo ResultInfo `json:"resultInfo"`
	// field.getPriceResponse.data
	Data []Price `json:"data,omitempty"`
}

// NewGetPriceV2Response instantiates a new GetPriceV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPriceV2Response(resultInfo ResultInfo) *GetPriceV2Response {
	this := GetPriceV2Response{}
	this.ResultInfo = resultInfo
	return &this
}

// NewGetPriceV2ResponseWithDefaults instantiates a new GetPriceV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPriceV2ResponseWithDefaults() *GetPriceV2Response {
	this := GetPriceV2Response{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *GetPriceV2Response) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *GetPriceV2Response) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *GetPriceV2Response) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetPriceV2Response) GetData() []Price {
	if o == nil || o.Data == nil {
		var ret []Price
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceV2Response) GetDataOk() ([]Price, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetPriceV2Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Price and assigns it to the Data field.
func (o *GetPriceV2Response) SetData(v []Price) {
	o.Data = v
}

func (o GetPriceV2Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetPriceV2Response struct {
	value *GetPriceV2Response
	isSet bool
}

func (v NullableGetPriceV2Response) Get() *GetPriceV2Response {
	return v.value
}

func (v *NullableGetPriceV2Response) Set(val *GetPriceV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPriceV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPriceV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPriceV2Response(val *GetPriceV2Response) *NullableGetPriceV2Response {
	return &NullableGetPriceV2Response{value: val, isSet: true}
}

func (v NullableGetPriceV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPriceV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}