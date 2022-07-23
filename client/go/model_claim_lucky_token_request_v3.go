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

// ClaimLuckyTokenRequestV3 struct for ClaimLuckyTokenRequestV3
type ClaimLuckyTokenRequestV3 struct {
	// field.ClaimLuckyTokenRequestV3.hash
	Hash string `json:"hash"`
	// field.SendLuckyTokenRequestV3.claimer
	Claimer string `json:"claimer"`
	// field.SendLuckyTokenRequestV3.referrer
	Referrer *string `json:"referrer,omitempty"`
}

// NewClaimLuckyTokenRequestV3 instantiates a new ClaimLuckyTokenRequestV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimLuckyTokenRequestV3(hash string, claimer string) *ClaimLuckyTokenRequestV3 {
	this := ClaimLuckyTokenRequestV3{}
	this.Hash = hash
	this.Claimer = claimer
	return &this
}

// NewClaimLuckyTokenRequestV3WithDefaults instantiates a new ClaimLuckyTokenRequestV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimLuckyTokenRequestV3WithDefaults() *ClaimLuckyTokenRequestV3 {
	this := ClaimLuckyTokenRequestV3{}
	return &this
}

// GetHash returns the Hash field value
func (o *ClaimLuckyTokenRequestV3) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *ClaimLuckyTokenRequestV3) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *ClaimLuckyTokenRequestV3) SetHash(v string) {
	o.Hash = v
}

// GetClaimer returns the Claimer field value
func (o *ClaimLuckyTokenRequestV3) GetClaimer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Claimer
}

// GetClaimerOk returns a tuple with the Claimer field value
// and a boolean to check if the value has been set.
func (o *ClaimLuckyTokenRequestV3) GetClaimerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Claimer, true
}

// SetClaimer sets field value
func (o *ClaimLuckyTokenRequestV3) SetClaimer(v string) {
	o.Claimer = v
}

// GetReferrer returns the Referrer field value if set, zero value otherwise.
func (o *ClaimLuckyTokenRequestV3) GetReferrer() string {
	if o == nil || o.Referrer == nil {
		var ret string
		return ret
	}
	return *o.Referrer
}

// GetReferrerOk returns a tuple with the Referrer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimLuckyTokenRequestV3) GetReferrerOk() (*string, bool) {
	if o == nil || o.Referrer == nil {
		return nil, false
	}
	return o.Referrer, true
}

// HasReferrer returns a boolean if a field has been set.
func (o *ClaimLuckyTokenRequestV3) HasReferrer() bool {
	if o != nil && o.Referrer != nil {
		return true
	}

	return false
}

// SetReferrer gets a reference to the given string and assigns it to the Referrer field.
func (o *ClaimLuckyTokenRequestV3) SetReferrer(v string) {
	o.Referrer = &v
}

func (o ClaimLuckyTokenRequestV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["claimer"] = o.Claimer
	}
	if o.Referrer != nil {
		toSerialize["referrer"] = o.Referrer
	}
	return json.Marshal(toSerialize)
}

type NullableClaimLuckyTokenRequestV3 struct {
	value *ClaimLuckyTokenRequestV3
	isSet bool
}

func (v NullableClaimLuckyTokenRequestV3) Get() *ClaimLuckyTokenRequestV3 {
	return v.value
}

func (v *NullableClaimLuckyTokenRequestV3) Set(val *ClaimLuckyTokenRequestV3) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimLuckyTokenRequestV3) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimLuckyTokenRequestV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimLuckyTokenRequestV3(val *ClaimLuckyTokenRequestV3) *NullableClaimLuckyTokenRequestV3 {
	return &NullableClaimLuckyTokenRequestV3{value: val, isSet: true}
}

func (v NullableClaimLuckyTokenRequestV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimLuckyTokenRequestV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
