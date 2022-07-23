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

// NftTokenAmountInfo Wrapper object used to describe a NFT token associated with a certain quantity.
type NftTokenAmountInfo struct {
	// The Loopring's NFT token identifier.
	TokenId int32 `json:"tokenId"`
	// The Loopring's NFT token data identifier which is a hash string of NFT token address and NFT_ID
	NftData *string `json:"nftData,omitempty"`
	// The amount of the NFT token
	Amount string `json:"amount"`
}

// NewNftTokenAmountInfo instantiates a new NftTokenAmountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftTokenAmountInfo(tokenId int32, amount string) *NftTokenAmountInfo {
	this := NftTokenAmountInfo{}
	this.TokenId = tokenId
	this.Amount = amount
	return &this
}

// NewNftTokenAmountInfoWithDefaults instantiates a new NftTokenAmountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftTokenAmountInfoWithDefaults() *NftTokenAmountInfo {
	this := NftTokenAmountInfo{}
	return &this
}

// GetTokenId returns the TokenId field value
func (o *NftTokenAmountInfo) GetTokenId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *NftTokenAmountInfo) GetTokenIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *NftTokenAmountInfo) SetTokenId(v int32) {
	o.TokenId = v
}

// GetNftData returns the NftData field value if set, zero value otherwise.
func (o *NftTokenAmountInfo) GetNftData() string {
	if o == nil || o.NftData == nil {
		var ret string
		return ret
	}
	return *o.NftData
}

// GetNftDataOk returns a tuple with the NftData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftTokenAmountInfo) GetNftDataOk() (*string, bool) {
	if o == nil || o.NftData == nil {
		return nil, false
	}
	return o.NftData, true
}

// HasNftData returns a boolean if a field has been set.
func (o *NftTokenAmountInfo) HasNftData() bool {
	if o != nil && o.NftData != nil {
		return true
	}

	return false
}

// SetNftData gets a reference to the given string and assigns it to the NftData field.
func (o *NftTokenAmountInfo) SetNftData(v string) {
	o.NftData = &v
}

// GetAmount returns the Amount field value
func (o *NftTokenAmountInfo) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *NftTokenAmountInfo) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *NftTokenAmountInfo) SetAmount(v string) {
	o.Amount = v
}

func (o NftTokenAmountInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tokenId"] = o.TokenId
	}
	if o.NftData != nil {
		toSerialize["nftData"] = o.NftData
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableNftTokenAmountInfo struct {
	value *NftTokenAmountInfo
	isSet bool
}

func (v NullableNftTokenAmountInfo) Get() *NftTokenAmountInfo {
	return v.value
}

func (v *NullableNftTokenAmountInfo) Set(val *NftTokenAmountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNftTokenAmountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNftTokenAmountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftTokenAmountInfo(val *NftTokenAmountInfo) *NullableNftTokenAmountInfo {
	return &NullableNftTokenAmountInfo{value: val, isSet: true}
}

func (v NullableNftTokenAmountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftTokenAmountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}