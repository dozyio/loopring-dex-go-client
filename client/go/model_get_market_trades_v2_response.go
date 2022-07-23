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

// GetMarketTradesV2Response Query market trades returns
type GetMarketTradesV2Response struct {
	ResultInfo ResultInfo    `json:"resultInfo"`
	Data       *MarketTrades `json:"data,omitempty"`
}

// NewGetMarketTradesV2Response instantiates a new GetMarketTradesV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketTradesV2Response(resultInfo ResultInfo) *GetMarketTradesV2Response {
	this := GetMarketTradesV2Response{}
	this.ResultInfo = resultInfo
	return &this
}

// NewGetMarketTradesV2ResponseWithDefaults instantiates a new GetMarketTradesV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketTradesV2ResponseWithDefaults() *GetMarketTradesV2Response {
	this := GetMarketTradesV2Response{}
	return &this
}

// GetResultInfo returns the ResultInfo field value
func (o *GetMarketTradesV2Response) GetResultInfo() ResultInfo {
	if o == nil {
		var ret ResultInfo
		return ret
	}

	return o.ResultInfo
}

// GetResultInfoOk returns a tuple with the ResultInfo field value
// and a boolean to check if the value has been set.
func (o *GetMarketTradesV2Response) GetResultInfoOk() (*ResultInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultInfo, true
}

// SetResultInfo sets field value
func (o *GetMarketTradesV2Response) SetResultInfo(v ResultInfo) {
	o.ResultInfo = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetMarketTradesV2Response) GetData() MarketTrades {
	if o == nil || o.Data == nil {
		var ret MarketTrades
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketTradesV2Response) GetDataOk() (*MarketTrades, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetMarketTradesV2Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given MarketTrades and assigns it to the Data field.
func (o *GetMarketTradesV2Response) SetData(v MarketTrades) {
	o.Data = &v
}

func (o GetMarketTradesV2Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resultInfo"] = o.ResultInfo
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetMarketTradesV2Response struct {
	value *GetMarketTradesV2Response
	isSet bool
}

func (v NullableGetMarketTradesV2Response) Get() *GetMarketTradesV2Response {
	return v.value
}

func (v *NullableGetMarketTradesV2Response) Set(val *GetMarketTradesV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketTradesV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketTradesV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketTradesV2Response(val *GetMarketTradesV2Response) *NullableGetMarketTradesV2Response {
	return &NullableGetMarketTradesV2Response{value: val, isSet: true}
}

func (v NullableGetMarketTradesV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketTradesV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}