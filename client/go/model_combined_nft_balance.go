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

// CombinedNftBalance model.nftBalance.description
type CombinedNftBalance struct {
	// field.Balance.accountId
	Id int64 `json:"id"`
	// field.Balance.accountId
	AccountId int64 `json:"accountId"`
	// field.Balance.tokenId
	TokenId int32 `json:"tokenId"`
	// Users NFT token nftData
	NftData *string `json:"nftData,omitempty"`
	// field.NftBalance.tokenAddress
	TokenAddress *string `json:"tokenAddress,omitempty"`
	// field.NftBalance.nftId
	NftId *string `json:"nftId,omitempty"`
	// field.NftBalance.nftType
	NftType *string `json:"nftType,omitempty"`
	// field.Balance.totalAmount
	Total string `json:"total"`
	// field.Balance.frozenAmount
	Locked  string         `json:"locked"`
	Pending PendingBalance `json:"pending"`
	// field.Balance.deploymentStatus
	DeploymentStatus *string `json:"deploymentStatus,omitempty"`
	// field.Balance.isCounterFactualNFT
	IsCounterFactualNFT *bool              `json:"isCounterFactualNFT,omitempty"`
	Metadata            *NftMetadataL2Info `json:"metadata,omitempty"`
}

// NewCombinedNftBalance instantiates a new CombinedNftBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCombinedNftBalance(id int64, accountId int64, tokenId int32, total string, locked string, pending PendingBalance) *CombinedNftBalance {
	this := CombinedNftBalance{}
	this.Id = id
	this.AccountId = accountId
	this.TokenId = tokenId
	this.Total = total
	this.Locked = locked
	this.Pending = pending
	return &this
}

// NewCombinedNftBalanceWithDefaults instantiates a new CombinedNftBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCombinedNftBalanceWithDefaults() *CombinedNftBalance {
	this := CombinedNftBalance{}
	return &this
}

// GetId returns the Id field value
func (o *CombinedNftBalance) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CombinedNftBalance) SetId(v int64) {
	o.Id = v
}

// GetAccountId returns the AccountId field value
func (o *CombinedNftBalance) GetAccountId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetAccountIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CombinedNftBalance) SetAccountId(v int64) {
	o.AccountId = v
}

// GetTokenId returns the TokenId field value
func (o *CombinedNftBalance) GetTokenId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetTokenIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *CombinedNftBalance) SetTokenId(v int32) {
	o.TokenId = v
}

// GetNftData returns the NftData field value if set, zero value otherwise.
func (o *CombinedNftBalance) GetNftData() string {
	if o == nil || o.NftData == nil {
		var ret string
		return ret
	}
	return *o.NftData
}

// GetNftDataOk returns a tuple with the NftData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetNftDataOk() (*string, bool) {
	if o == nil || o.NftData == nil {
		return nil, false
	}
	return o.NftData, true
}

// HasNftData returns a boolean if a field has been set.
func (o *CombinedNftBalance) HasNftData() bool {
	if o != nil && o.NftData != nil {
		return true
	}

	return false
}

// SetNftData gets a reference to the given string and assigns it to the NftData field.
func (o *CombinedNftBalance) SetNftData(v string) {
	o.NftData = &v
}

// GetTokenAddress returns the TokenAddress field value if set, zero value otherwise.
func (o *CombinedNftBalance) GetTokenAddress() string {
	if o == nil || o.TokenAddress == nil {
		var ret string
		return ret
	}
	return *o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetTokenAddressOk() (*string, bool) {
	if o == nil || o.TokenAddress == nil {
		return nil, false
	}
	return o.TokenAddress, true
}

// HasTokenAddress returns a boolean if a field has been set.
func (o *CombinedNftBalance) HasTokenAddress() bool {
	if o != nil && o.TokenAddress != nil {
		return true
	}

	return false
}

// SetTokenAddress gets a reference to the given string and assigns it to the TokenAddress field.
func (o *CombinedNftBalance) SetTokenAddress(v string) {
	o.TokenAddress = &v
}

// GetNftId returns the NftId field value if set, zero value otherwise.
func (o *CombinedNftBalance) GetNftId() string {
	if o == nil || o.NftId == nil {
		var ret string
		return ret
	}
	return *o.NftId
}

// GetNftIdOk returns a tuple with the NftId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetNftIdOk() (*string, bool) {
	if o == nil || o.NftId == nil {
		return nil, false
	}
	return o.NftId, true
}

// HasNftId returns a boolean if a field has been set.
func (o *CombinedNftBalance) HasNftId() bool {
	if o != nil && o.NftId != nil {
		return true
	}

	return false
}

// SetNftId gets a reference to the given string and assigns it to the NftId field.
func (o *CombinedNftBalance) SetNftId(v string) {
	o.NftId = &v
}

// GetNftType returns the NftType field value if set, zero value otherwise.
func (o *CombinedNftBalance) GetNftType() string {
	if o == nil || o.NftType == nil {
		var ret string
		return ret
	}
	return *o.NftType
}

// GetNftTypeOk returns a tuple with the NftType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetNftTypeOk() (*string, bool) {
	if o == nil || o.NftType == nil {
		return nil, false
	}
	return o.NftType, true
}

// HasNftType returns a boolean if a field has been set.
func (o *CombinedNftBalance) HasNftType() bool {
	if o != nil && o.NftType != nil {
		return true
	}

	return false
}

// SetNftType gets a reference to the given string and assigns it to the NftType field.
func (o *CombinedNftBalance) SetNftType(v string) {
	o.NftType = &v
}

// GetTotal returns the Total field value
func (o *CombinedNftBalance) GetTotal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetTotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CombinedNftBalance) SetTotal(v string) {
	o.Total = v
}

// GetLocked returns the Locked field value
func (o *CombinedNftBalance) GetLocked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetLockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *CombinedNftBalance) SetLocked(v string) {
	o.Locked = v
}

// GetPending returns the Pending field value
func (o *CombinedNftBalance) GetPending() PendingBalance {
	if o == nil {
		var ret PendingBalance
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetPendingOk() (*PendingBalance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *CombinedNftBalance) SetPending(v PendingBalance) {
	o.Pending = v
}

// GetDeploymentStatus returns the DeploymentStatus field value if set, zero value otherwise.
func (o *CombinedNftBalance) GetDeploymentStatus() string {
	if o == nil || o.DeploymentStatus == nil {
		var ret string
		return ret
	}
	return *o.DeploymentStatus
}

// GetDeploymentStatusOk returns a tuple with the DeploymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetDeploymentStatusOk() (*string, bool) {
	if o == nil || o.DeploymentStatus == nil {
		return nil, false
	}
	return o.DeploymentStatus, true
}

// HasDeploymentStatus returns a boolean if a field has been set.
func (o *CombinedNftBalance) HasDeploymentStatus() bool {
	if o != nil && o.DeploymentStatus != nil {
		return true
	}

	return false
}

// SetDeploymentStatus gets a reference to the given string and assigns it to the DeploymentStatus field.
func (o *CombinedNftBalance) SetDeploymentStatus(v string) {
	o.DeploymentStatus = &v
}

// GetIsCounterFactualNFT returns the IsCounterFactualNFT field value if set, zero value otherwise.
func (o *CombinedNftBalance) GetIsCounterFactualNFT() bool {
	if o == nil || o.IsCounterFactualNFT == nil {
		var ret bool
		return ret
	}
	return *o.IsCounterFactualNFT
}

// GetIsCounterFactualNFTOk returns a tuple with the IsCounterFactualNFT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetIsCounterFactualNFTOk() (*bool, bool) {
	if o == nil || o.IsCounterFactualNFT == nil {
		return nil, false
	}
	return o.IsCounterFactualNFT, true
}

// HasIsCounterFactualNFT returns a boolean if a field has been set.
func (o *CombinedNftBalance) HasIsCounterFactualNFT() bool {
	if o != nil && o.IsCounterFactualNFT != nil {
		return true
	}

	return false
}

// SetIsCounterFactualNFT gets a reference to the given bool and assigns it to the IsCounterFactualNFT field.
func (o *CombinedNftBalance) SetIsCounterFactualNFT(v bool) {
	o.IsCounterFactualNFT = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CombinedNftBalance) GetMetadata() NftMetadataL2Info {
	if o == nil || o.Metadata == nil {
		var ret NftMetadataL2Info
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombinedNftBalance) GetMetadataOk() (*NftMetadataL2Info, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CombinedNftBalance) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given NftMetadataL2Info and assigns it to the Metadata field.
func (o *CombinedNftBalance) SetMetadata(v NftMetadataL2Info) {
	o.Metadata = &v
}

func (o CombinedNftBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["tokenId"] = o.TokenId
	}
	if o.NftData != nil {
		toSerialize["nftData"] = o.NftData
	}
	if o.TokenAddress != nil {
		toSerialize["tokenAddress"] = o.TokenAddress
	}
	if o.NftId != nil {
		toSerialize["nftId"] = o.NftId
	}
	if o.NftType != nil {
		toSerialize["nftType"] = o.NftType
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["locked"] = o.Locked
	}
	if true {
		toSerialize["pending"] = o.Pending
	}
	if o.DeploymentStatus != nil {
		toSerialize["deploymentStatus"] = o.DeploymentStatus
	}
	if o.IsCounterFactualNFT != nil {
		toSerialize["isCounterFactualNFT"] = o.IsCounterFactualNFT
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableCombinedNftBalance struct {
	value *CombinedNftBalance
	isSet bool
}

func (v NullableCombinedNftBalance) Get() *CombinedNftBalance {
	return v.value
}

func (v *NullableCombinedNftBalance) Set(val *CombinedNftBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableCombinedNftBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableCombinedNftBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCombinedNftBalance(val *CombinedNftBalance) *NullableCombinedNftBalance {
	return &NullableCombinedNftBalance{value: val, isSet: true}
}

func (v NullableCombinedNftBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCombinedNftBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
