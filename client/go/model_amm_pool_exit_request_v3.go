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

// AmmPoolExitRequestV3 struct for AmmPoolExitRequestV3
type AmmPoolExitRequestV3 struct {
	// The account owner adderss
	Owner string `json:"owner"`
	// AMM pool address to be joined
	PoolAddress string            `json:"poolAddress"`
	ExitTokens  AmmPoolExitTokens `json:"exitTokens"`
	// Offchain request storage Id
	StorageId int64 `json:"storageId"`
	// Maximum fee of exit request, use the last in pool token by default
	MaxFee string `json:"maxFee"`
	// Timestamp for order to become invalid
	ValidUntil int32 `json:"validUntil"`
	// AMM exit request eddsa signature
	EddsaSignature *string `json:"eddsaSignature,omitempty"`
	// AMM exit request ecdsa signature
	EcdsaSignature *string `json:"ecdsaSignature,omitempty"`
}

// NewAmmPoolExitRequestV3 instantiates a new AmmPoolExitRequestV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmmPoolExitRequestV3(owner string, poolAddress string, exitTokens AmmPoolExitTokens, storageId int64, maxFee string, validUntil int32) *AmmPoolExitRequestV3 {
	this := AmmPoolExitRequestV3{}
	this.Owner = owner
	this.PoolAddress = poolAddress
	this.ExitTokens = exitTokens
	this.StorageId = storageId
	this.MaxFee = maxFee
	this.ValidUntil = validUntil
	return &this
}

// NewAmmPoolExitRequestV3WithDefaults instantiates a new AmmPoolExitRequestV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmmPoolExitRequestV3WithDefaults() *AmmPoolExitRequestV3 {
	this := AmmPoolExitRequestV3{}
	return &this
}

// GetOwner returns the Owner field value
func (o *AmmPoolExitRequestV3) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *AmmPoolExitRequestV3) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *AmmPoolExitRequestV3) SetOwner(v string) {
	o.Owner = v
}

// GetPoolAddress returns the PoolAddress field value
func (o *AmmPoolExitRequestV3) GetPoolAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoolAddress
}

// GetPoolAddressOk returns a tuple with the PoolAddress field value
// and a boolean to check if the value has been set.
func (o *AmmPoolExitRequestV3) GetPoolAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolAddress, true
}

// SetPoolAddress sets field value
func (o *AmmPoolExitRequestV3) SetPoolAddress(v string) {
	o.PoolAddress = v
}

// GetExitTokens returns the ExitTokens field value
func (o *AmmPoolExitRequestV3) GetExitTokens() AmmPoolExitTokens {
	if o == nil {
		var ret AmmPoolExitTokens
		return ret
	}

	return o.ExitTokens
}

// GetExitTokensOk returns a tuple with the ExitTokens field value
// and a boolean to check if the value has been set.
func (o *AmmPoolExitRequestV3) GetExitTokensOk() (*AmmPoolExitTokens, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExitTokens, true
}

// SetExitTokens sets field value
func (o *AmmPoolExitRequestV3) SetExitTokens(v AmmPoolExitTokens) {
	o.ExitTokens = v
}

// GetStorageId returns the StorageId field value
func (o *AmmPoolExitRequestV3) GetStorageId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value
// and a boolean to check if the value has been set.
func (o *AmmPoolExitRequestV3) GetStorageIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageId, true
}

// SetStorageId sets field value
func (o *AmmPoolExitRequestV3) SetStorageId(v int64) {
	o.StorageId = v
}

// GetMaxFee returns the MaxFee field value
func (o *AmmPoolExitRequestV3) GetMaxFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxFee
}

// GetMaxFeeOk returns a tuple with the MaxFee field value
// and a boolean to check if the value has been set.
func (o *AmmPoolExitRequestV3) GetMaxFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxFee, true
}

// SetMaxFee sets field value
func (o *AmmPoolExitRequestV3) SetMaxFee(v string) {
	o.MaxFee = v
}

// GetValidUntil returns the ValidUntil field value
func (o *AmmPoolExitRequestV3) GetValidUntil() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
func (o *AmmPoolExitRequestV3) GetValidUntilOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidUntil, true
}

// SetValidUntil sets field value
func (o *AmmPoolExitRequestV3) SetValidUntil(v int32) {
	o.ValidUntil = v
}

// GetEddsaSignature returns the EddsaSignature field value if set, zero value otherwise.
func (o *AmmPoolExitRequestV3) GetEddsaSignature() string {
	if o == nil || o.EddsaSignature == nil {
		var ret string
		return ret
	}
	return *o.EddsaSignature
}

// GetEddsaSignatureOk returns a tuple with the EddsaSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmmPoolExitRequestV3) GetEddsaSignatureOk() (*string, bool) {
	if o == nil || o.EddsaSignature == nil {
		return nil, false
	}
	return o.EddsaSignature, true
}

// HasEddsaSignature returns a boolean if a field has been set.
func (o *AmmPoolExitRequestV3) HasEddsaSignature() bool {
	if o != nil && o.EddsaSignature != nil {
		return true
	}

	return false
}

// SetEddsaSignature gets a reference to the given string and assigns it to the EddsaSignature field.
func (o *AmmPoolExitRequestV3) SetEddsaSignature(v string) {
	o.EddsaSignature = &v
}

// GetEcdsaSignature returns the EcdsaSignature field value if set, zero value otherwise.
func (o *AmmPoolExitRequestV3) GetEcdsaSignature() string {
	if o == nil || o.EcdsaSignature == nil {
		var ret string
		return ret
	}
	return *o.EcdsaSignature
}

// GetEcdsaSignatureOk returns a tuple with the EcdsaSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmmPoolExitRequestV3) GetEcdsaSignatureOk() (*string, bool) {
	if o == nil || o.EcdsaSignature == nil {
		return nil, false
	}
	return o.EcdsaSignature, true
}

// HasEcdsaSignature returns a boolean if a field has been set.
func (o *AmmPoolExitRequestV3) HasEcdsaSignature() bool {
	if o != nil && o.EcdsaSignature != nil {
		return true
	}

	return false
}

// SetEcdsaSignature gets a reference to the given string and assigns it to the EcdsaSignature field.
func (o *AmmPoolExitRequestV3) SetEcdsaSignature(v string) {
	o.EcdsaSignature = &v
}

func (o AmmPoolExitRequestV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["poolAddress"] = o.PoolAddress
	}
	if true {
		toSerialize["exitTokens"] = o.ExitTokens
	}
	if true {
		toSerialize["storageId"] = o.StorageId
	}
	if true {
		toSerialize["maxFee"] = o.MaxFee
	}
	if true {
		toSerialize["validUntil"] = o.ValidUntil
	}
	if o.EddsaSignature != nil {
		toSerialize["eddsaSignature"] = o.EddsaSignature
	}
	if o.EcdsaSignature != nil {
		toSerialize["ecdsaSignature"] = o.EcdsaSignature
	}
	return json.Marshal(toSerialize)
}

type NullableAmmPoolExitRequestV3 struct {
	value *AmmPoolExitRequestV3
	isSet bool
}

func (v NullableAmmPoolExitRequestV3) Get() *AmmPoolExitRequestV3 {
	return v.value
}

func (v *NullableAmmPoolExitRequestV3) Set(val *AmmPoolExitRequestV3) {
	v.value = val
	v.isSet = true
}

func (v NullableAmmPoolExitRequestV3) IsSet() bool {
	return v.isSet
}

func (v *NullableAmmPoolExitRequestV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmmPoolExitRequestV3(val *AmmPoolExitRequestV3) *NullableAmmPoolExitRequestV3 {
	return &NullableAmmPoolExitRequestV3{value: val, isSet: true}
}

func (v NullableAmmPoolExitRequestV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmmPoolExitRequestV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
