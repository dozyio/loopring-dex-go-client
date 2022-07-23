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

// NftMetadataL2Info struct for NftMetadataL2Info
type NftMetadataL2Info struct {
	Uri          string               `json:"uri"`
	Base         *NftMetadataBaseInfo `json:"base,omitempty"`
	ImageSize    map[string]string    `json:"imageSize"`
	Extra        *NftExtra            `json:"extra,omitempty"`
	Status       int32                `json:"status"`
	NftType      int32                `json:"nftType"`
	Network      int32                `json:"network"`
	TokenAddress string               `json:"tokenAddress"`
	NftId        string               `json:"nftId"`
}

// NewNftMetadataL2Info instantiates a new NftMetadataL2Info object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftMetadataL2Info(uri string, imageSize map[string]string, status int32, nftType int32, network int32, tokenAddress string, nftId string) *NftMetadataL2Info {
	this := NftMetadataL2Info{}
	this.Uri = uri
	this.ImageSize = imageSize
	this.Status = status
	this.NftType = nftType
	this.Network = network
	this.TokenAddress = tokenAddress
	this.NftId = nftId
	return &this
}

// NewNftMetadataL2InfoWithDefaults instantiates a new NftMetadataL2Info object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftMetadataL2InfoWithDefaults() *NftMetadataL2Info {
	this := NftMetadataL2Info{}
	return &this
}

// GetUri returns the Uri field value
func (o *NftMetadataL2Info) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *NftMetadataL2Info) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *NftMetadataL2Info) SetUri(v string) {
	o.Uri = v
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *NftMetadataL2Info) GetBase() NftMetadataBaseInfo {
	if o == nil || o.Base == nil {
		var ret NftMetadataBaseInfo
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftMetadataL2Info) GetBaseOk() (*NftMetadataBaseInfo, bool) {
	if o == nil || o.Base == nil {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *NftMetadataL2Info) HasBase() bool {
	if o != nil && o.Base != nil {
		return true
	}

	return false
}

// SetBase gets a reference to the given NftMetadataBaseInfo and assigns it to the Base field.
func (o *NftMetadataL2Info) SetBase(v NftMetadataBaseInfo) {
	o.Base = &v
}

// GetImageSize returns the ImageSize field value
func (o *NftMetadataL2Info) GetImageSize() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.ImageSize
}

// GetImageSizeOk returns a tuple with the ImageSize field value
// and a boolean to check if the value has been set.
func (o *NftMetadataL2Info) GetImageSizeOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageSize, true
}

// SetImageSize sets field value
func (o *NftMetadataL2Info) SetImageSize(v map[string]string) {
	o.ImageSize = v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *NftMetadataL2Info) GetExtra() NftExtra {
	if o == nil || o.Extra == nil {
		var ret NftExtra
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftMetadataL2Info) GetExtraOk() (*NftExtra, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *NftMetadataL2Info) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given NftExtra and assigns it to the Extra field.
func (o *NftMetadataL2Info) SetExtra(v NftExtra) {
	o.Extra = &v
}

// GetStatus returns the Status field value
func (o *NftMetadataL2Info) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NftMetadataL2Info) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NftMetadataL2Info) SetStatus(v int32) {
	o.Status = v
}

// GetNftType returns the NftType field value
func (o *NftMetadataL2Info) GetNftType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NftType
}

// GetNftTypeOk returns a tuple with the NftType field value
// and a boolean to check if the value has been set.
func (o *NftMetadataL2Info) GetNftTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NftType, true
}

// SetNftType sets field value
func (o *NftMetadataL2Info) SetNftType(v int32) {
	o.NftType = v
}

// GetNetwork returns the Network field value
func (o *NftMetadataL2Info) GetNetwork() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *NftMetadataL2Info) GetNetworkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *NftMetadataL2Info) SetNetwork(v int32) {
	o.Network = v
}

// GetTokenAddress returns the TokenAddress field value
func (o *NftMetadataL2Info) GetTokenAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value
// and a boolean to check if the value has been set.
func (o *NftMetadataL2Info) GetTokenAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenAddress, true
}

// SetTokenAddress sets field value
func (o *NftMetadataL2Info) SetTokenAddress(v string) {
	o.TokenAddress = v
}

// GetNftId returns the NftId field value
func (o *NftMetadataL2Info) GetNftId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NftId
}

// GetNftIdOk returns a tuple with the NftId field value
// and a boolean to check if the value has been set.
func (o *NftMetadataL2Info) GetNftIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NftId, true
}

// SetNftId sets field value
func (o *NftMetadataL2Info) SetNftId(v string) {
	o.NftId = v
}

func (o NftMetadataL2Info) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uri"] = o.Uri
	}
	if o.Base != nil {
		toSerialize["base"] = o.Base
	}
	if true {
		toSerialize["imageSize"] = o.ImageSize
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["nftType"] = o.NftType
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["tokenAddress"] = o.TokenAddress
	}
	if true {
		toSerialize["nftId"] = o.NftId
	}
	return json.Marshal(toSerialize)
}

type NullableNftMetadataL2Info struct {
	value *NftMetadataL2Info
	isSet bool
}

func (v NullableNftMetadataL2Info) Get() *NftMetadataL2Info {
	return v.value
}

func (v *NullableNftMetadataL2Info) Set(val *NftMetadataL2Info) {
	v.value = val
	v.isSet = true
}

func (v NullableNftMetadataL2Info) IsSet() bool {
	return v.isSet
}

func (v *NullableNftMetadataL2Info) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftMetadataL2Info(val *NftMetadataL2Info) *NullableNftMetadataL2Info {
	return &NullableNftMetadataL2Info{value: val, isSet: true}
}

func (v NullableNftMetadataL2Info) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftMetadataL2Info) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}