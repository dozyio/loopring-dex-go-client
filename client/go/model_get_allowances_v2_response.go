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

// GetAllowancesV2Response Query the allowance return of the Ethereum address on the exchange
type GetAllowancesV2Response struct {
	ResultInfo ResultInfo `json:"resultInfo"`
	// Allowance amount in wei
	Data []string `json:"data,omitempty"`
}

// NewGetAllowancesV2Response instantiates a new GetAllowancesV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllowancesV2Response(resultInfo ResultInfo) *GetAllowancesV2Response {
	this := GetAllowancesV2Response{}
	this.ResultInfo = resultInfo
	return &this
}

// NewGetAllowancesV2ResponseWithDefaults instantiates a new GetAllowancesV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllowancesV2ResponseWithDefaults() *GetAllowancesV2Response {
	this := GetAllowancesV2Response{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *GetAllowancesV2Response) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *GetAllowancesV2Response) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *GetAllowancesV2Response) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAllowancesV2Response) GetData() []string {
	if o == nil || o.Data == nil {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllowancesV2Response) GetDataOk() ([]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAllowancesV2Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *GetAllowancesV2Response) SetData(v []string) {
	o.Data = v
}

func (o GetAllowancesV2Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetAllowancesV2Response struct {
	value *GetAllowancesV2Response
	isSet bool
}

func (v NullableGetAllowancesV2Response) Get() *GetAllowancesV2Response {
	return v.value
}

func (v *NullableGetAllowancesV2Response) Set(val *GetAllowancesV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllowancesV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllowancesV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllowancesV2Response(val *GetAllowancesV2Response) *NullableGetAllowancesV2Response {
	return &NullableGetAllowancesV2Response{value: val, isSet: true}
}

func (v NullableGetAllowancesV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllowancesV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}