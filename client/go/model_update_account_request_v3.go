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

// UpdateAccountRequestV3 Params to update account EDDSA key
type UpdateAccountRequestV3 struct {
	// exchange address
	Exchange string `json:"exchange"`
	// owner address
	Owner string `json:"owner"`
	// user account ID
	AccountId int64         `json:"accountId"`
	PublicKey PublicKey     `json:"publicKey"`
	MaxFee    TokenVolumeV3 `json:"maxFee"`
	// Timestamp for order to become invalid
	ValidUntil int32 `json:"validUntil"`
	// Nonce of users exchange account that used in off-chain requests.
	Nonce int32 `json:"nonce"`
	// KeySeed of users L2 eddsaKey, the L2 key should be generated from this seed, i.e., L2_EDDSA_KEY=eth.sign(keySeed). Otherwise, user may meet error in login loopring DEX
	KeySeed            *string             `json:"keySeed,omitempty"`
	CounterFactualInfo *CounterFactualInfo `json:"counterFactualInfo,omitempty"`
	// eddsa signature of this request
	EddsaSignature *string `json:"eddsaSignature,omitempty"`
	// ecdsa signature of this request
	EcdsaSignature *string `json:"ecdsaSignature,omitempty"`
	// An approved hash string which was submitted on eth mainnet
	HashApproved *string `json:"hashApproved,omitempty"`
}

// NewUpdateAccountRequestV3 instantiates a new UpdateAccountRequestV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountRequestV3(exchange string, owner string, accountId int64, publicKey PublicKey, maxFee TokenVolumeV3, validUntil int32, nonce int32) *UpdateAccountRequestV3 {
	this := UpdateAccountRequestV3{}
	this.Exchange = exchange
	this.Owner = owner
	this.AccountId = accountId
	this.PublicKey = publicKey
	this.MaxFee = maxFee
	this.ValidUntil = validUntil
	this.Nonce = nonce
	return &this
}

// NewUpdateAccountRequestV3WithDefaults instantiates a new UpdateAccountRequestV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountRequestV3WithDefaults() *UpdateAccountRequestV3 {
	this := UpdateAccountRequestV3{}
	return &this
}

// GetExchange returns the Exchange field value
func (o *UpdateAccountRequestV3) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequestV3) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *UpdateAccountRequestV3) SetExchange(v string) {
	o.Exchange = v
}

// GetOwner returns the Owner field value
func (o *UpdateAccountRequestV3) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequestV3) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *UpdateAccountRequestV3) SetOwner(v string) {
	o.Owner = v
}

// GetAccountId returns the AccountId field value
func (o *UpdateAccountRequestV3) GetAccountId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequestV3) GetAccountIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *UpdateAccountRequestV3) SetAccountId(v int64) {
	o.AccountId = v
}

// GetPublicKey returns the PublicKey field value
func (o *UpdateAccountRequestV3) GetPublicKey() PublicKey {
	if o == nil {
		var ret PublicKey
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequestV3) GetPublicKeyOk() (*PublicKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *UpdateAccountRequestV3) SetPublicKey(v PublicKey) {
	o.PublicKey = v
}

// GetMaxFee returns the MaxFee field value
func (o *UpdateAccountRequestV3) GetMaxFee() TokenVolumeV3 {
	if o == nil {
		var ret TokenVolumeV3
		return ret
	}

	return o.MaxFee
}

// GetMaxFeeOk returns a tuple with the MaxFee field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequestV3) GetMaxFeeOk() (*TokenVolumeV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxFee, true
}

// SetMaxFee sets field value
func (o *UpdateAccountRequestV3) SetMaxFee(v TokenVolumeV3) {
	o.MaxFee = v
}

// GetValidUntil returns the ValidUntil field value
func (o *UpdateAccountRequestV3) GetValidUntil() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequestV3) GetValidUntilOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidUntil, true
}

// SetValidUntil sets field value
func (o *UpdateAccountRequestV3) SetValidUntil(v int32) {
	o.ValidUntil = v
}

// GetNonce returns the Nonce field value
func (o *UpdateAccountRequestV3) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequestV3) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *UpdateAccountRequestV3) SetNonce(v int32) {
	o.Nonce = v
}

// GetKeySeed returns the KeySeed field value if set, zero value otherwise.
func (o *UpdateAccountRequestV3) GetKeySeed() string {
	if o == nil || o.KeySeed == nil {
		var ret string
		return ret
	}
	return *o.KeySeed
}

// GetKeySeedOk returns a tuple with the KeySeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequestV3) GetKeySeedOk() (*string, bool) {
	if o == nil || o.KeySeed == nil {
		return nil, false
	}
	return o.KeySeed, true
}

// HasKeySeed returns a boolean if a field has been set.
func (o *UpdateAccountRequestV3) HasKeySeed() bool {
	if o != nil && o.KeySeed != nil {
		return true
	}

	return false
}

// SetKeySeed gets a reference to the given string and assigns it to the KeySeed field.
func (o *UpdateAccountRequestV3) SetKeySeed(v string) {
	o.KeySeed = &v
}

// GetCounterFactualInfo returns the CounterFactualInfo field value if set, zero value otherwise.
func (o *UpdateAccountRequestV3) GetCounterFactualInfo() CounterFactualInfo {
	if o == nil || o.CounterFactualInfo == nil {
		var ret CounterFactualInfo
		return ret
	}
	return *o.CounterFactualInfo
}

// GetCounterFactualInfoOk returns a tuple with the CounterFactualInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequestV3) GetCounterFactualInfoOk() (*CounterFactualInfo, bool) {
	if o == nil || o.CounterFactualInfo == nil {
		return nil, false
	}
	return o.CounterFactualInfo, true
}

// HasCounterFactualInfo returns a boolean if a field has been set.
func (o *UpdateAccountRequestV3) HasCounterFactualInfo() bool {
	if o != nil && o.CounterFactualInfo != nil {
		return true
	}

	return false
}

// SetCounterFactualInfo gets a reference to the given CounterFactualInfo and assigns it to the CounterFactualInfo field.
func (o *UpdateAccountRequestV3) SetCounterFactualInfo(v CounterFactualInfo) {
	o.CounterFactualInfo = &v
}

// GetEddsaSignature returns the EddsaSignature field value if set, zero value otherwise.
func (o *UpdateAccountRequestV3) GetEddsaSignature() string {
	if o == nil || o.EddsaSignature == nil {
		var ret string
		return ret
	}
	return *o.EddsaSignature
}

// GetEddsaSignatureOk returns a tuple with the EddsaSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequestV3) GetEddsaSignatureOk() (*string, bool) {
	if o == nil || o.EddsaSignature == nil {
		return nil, false
	}
	return o.EddsaSignature, true
}

// HasEddsaSignature returns a boolean if a field has been set.
func (o *UpdateAccountRequestV3) HasEddsaSignature() bool {
	if o != nil && o.EddsaSignature != nil {
		return true
	}

	return false
}

// SetEddsaSignature gets a reference to the given string and assigns it to the EddsaSignature field.
func (o *UpdateAccountRequestV3) SetEddsaSignature(v string) {
	o.EddsaSignature = &v
}

// GetEcdsaSignature returns the EcdsaSignature field value if set, zero value otherwise.
func (o *UpdateAccountRequestV3) GetEcdsaSignature() string {
	if o == nil || o.EcdsaSignature == nil {
		var ret string
		return ret
	}
	return *o.EcdsaSignature
}

// GetEcdsaSignatureOk returns a tuple with the EcdsaSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequestV3) GetEcdsaSignatureOk() (*string, bool) {
	if o == nil || o.EcdsaSignature == nil {
		return nil, false
	}
	return o.EcdsaSignature, true
}

// HasEcdsaSignature returns a boolean if a field has been set.
func (o *UpdateAccountRequestV3) HasEcdsaSignature() bool {
	if o != nil && o.EcdsaSignature != nil {
		return true
	}

	return false
}

// SetEcdsaSignature gets a reference to the given string and assigns it to the EcdsaSignature field.
func (o *UpdateAccountRequestV3) SetEcdsaSignature(v string) {
	o.EcdsaSignature = &v
}

// GetHashApproved returns the HashApproved field value if set, zero value otherwise.
func (o *UpdateAccountRequestV3) GetHashApproved() string {
	if o == nil || o.HashApproved == nil {
		var ret string
		return ret
	}
	return *o.HashApproved
}

// GetHashApprovedOk returns a tuple with the HashApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequestV3) GetHashApprovedOk() (*string, bool) {
	if o == nil || o.HashApproved == nil {
		return nil, false
	}
	return o.HashApproved, true
}

// HasHashApproved returns a boolean if a field has been set.
func (o *UpdateAccountRequestV3) HasHashApproved() bool {
	if o != nil && o.HashApproved != nil {
		return true
	}

	return false
}

// SetHashApproved gets a reference to the given string and assigns it to the HashApproved field.
func (o *UpdateAccountRequestV3) SetHashApproved(v string) {
	o.HashApproved = &v
}

func (o UpdateAccountRequestV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["exchange"] = o.Exchange
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["publicKey"] = o.PublicKey
	}
	if true {
		toSerialize["maxFee"] = o.MaxFee
	}
	if true {
		toSerialize["validUntil"] = o.ValidUntil
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if o.KeySeed != nil {
		toSerialize["keySeed"] = o.KeySeed
	}
	if o.CounterFactualInfo != nil {
		toSerialize["counterFactualInfo"] = o.CounterFactualInfo
	}
	if o.EddsaSignature != nil {
		toSerialize["eddsaSignature"] = o.EddsaSignature
	}
	if o.EcdsaSignature != nil {
		toSerialize["ecdsaSignature"] = o.EcdsaSignature
	}
	if o.HashApproved != nil {
		toSerialize["hashApproved"] = o.HashApproved
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAccountRequestV3 struct {
	value *UpdateAccountRequestV3
	isSet bool
}

func (v NullableUpdateAccountRequestV3) Get() *UpdateAccountRequestV3 {
	return v.value
}

func (v *NullableUpdateAccountRequestV3) Set(val *UpdateAccountRequestV3) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountRequestV3) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountRequestV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountRequestV3(val *UpdateAccountRequestV3) *NullableUpdateAccountRequestV3 {
	return &NullableUpdateAccountRequestV3{value: val, isSet: true}
}

func (v NullableUpdateAccountRequestV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountRequestV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
