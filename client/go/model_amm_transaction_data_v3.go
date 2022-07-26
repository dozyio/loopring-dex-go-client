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

// AmmTransactionDataV3 AMM transaction info data
type AmmTransactionDataV3 struct {
	// AMM transaction hash
	Hash string `json:"hash"`
	// AMM transaction type, i.e., join, exit, etc
	TxType string `json:"txType"`
	// AMM transaction processing status, i.e., processing, processed, failed, etc
	TxStatus string `json:"txStatus"`
	// AMM pool address of query
	AmmPoolAddress string `json:"ammPoolAddress"`
	// AMM transaction layer, 1 or 2
	AmmLayerType string `json:"ammLayerType"`
	// The in pool tokens transfers records of the AMM transaction
	PoolTokens []AmmTransferDataV3 `json:"poolTokens"`
	LpToken    AmmTransferDataV3   `json:"lpToken"`
	// Transaction creation time
	CreatedAt int64 `json:"createdAt"`
	// Transaction update time
	UpdatedAt    int64         `json:"updatedAt"`
	BlockId      int64         `json:"blockId"`
	IndexInBlock int32         `json:"indexInBlock"`
	StorageInfos []StorageInfo `json:"storageInfos,omitempty"`
}

// NewAmmTransactionDataV3 instantiates a new AmmTransactionDataV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmmTransactionDataV3(hash string, txType string, txStatus string, ammPoolAddress string, ammLayerType string, poolTokens []AmmTransferDataV3, lpToken AmmTransferDataV3, createdAt int64, updatedAt int64, blockId int64, indexInBlock int32) *AmmTransactionDataV3 {
	this := AmmTransactionDataV3{}
	this.Hash = hash
	this.TxType = txType
	this.TxStatus = txStatus
	this.AmmPoolAddress = ammPoolAddress
	this.AmmLayerType = ammLayerType
	this.PoolTokens = poolTokens
	this.LpToken = lpToken
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.BlockId = blockId
	this.IndexInBlock = indexInBlock
	return &this
}

// NewAmmTransactionDataV3WithDefaults instantiates a new AmmTransactionDataV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmmTransactionDataV3WithDefaults() *AmmTransactionDataV3 {
	this := AmmTransactionDataV3{}
	return &this
}

// GetHash returns the Hash field value
func (o *AmmTransactionDataV3) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *AmmTransactionDataV3) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *AmmTransactionDataV3) SetHash(v string) {
	o.Hash = v
}

// GetTxType returns the TxType field value
func (o *AmmTransactionDataV3) GetTxType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxType
}

// GetTxTypeOk returns a tuple with the TxType field value
// and a boolean to check if the value has been set.
func (o *AmmTransactionDataV3) GetTxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxType, true
}

// SetTxType sets field value
func (o *AmmTransactionDataV3) SetTxType(v string) {
	o.TxType = v
}

// GetTxStatus returns the TxStatus field value
func (o *AmmTransactionDataV3) GetTxStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxStatus
}

// GetTxStatusOk returns a tuple with the TxStatus field value
// and a boolean to check if the value has been set.
func (o *AmmTransactionDataV3) GetTxStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxStatus, true
}

// SetTxStatus sets field value
func (o *AmmTransactionDataV3) SetTxStatus(v string) {
	o.TxStatus = v
}

// GetAmmPoolAddress returns the AmmPoolAddress field value
func (o *AmmTransactionDataV3) GetAmmPoolAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmmPoolAddress
}

// GetAmmPoolAddressOk returns a tuple with the AmmPoolAddress field value
// and a boolean to check if the value has been set.
func (o *AmmTransactionDataV3) GetAmmPoolAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmmPoolAddress, true
}

// SetAmmPoolAddress sets field value
func (o *AmmTransactionDataV3) SetAmmPoolAddress(v string) {
	o.AmmPoolAddress = v
}

// GetAmmLayerType returns the AmmLayerType field value
func (o *AmmTransactionDataV3) GetAmmLayerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmmLayerType
}

// GetAmmLayerTypeOk returns a tuple with the AmmLayerType field value
// and a boolean to check if the value has been set.
func (o *AmmTransactionDataV3) GetAmmLayerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmmLayerType, true
}

// SetAmmLayerType sets field value
func (o *AmmTransactionDataV3) SetAmmLayerType(v string) {
	o.AmmLayerType = v
}

// GetPoolTokens returns the PoolTokens field value
func (o *AmmTransactionDataV3) GetPoolTokens() []AmmTransferDataV3 {
	if o == nil {
		var ret []AmmTransferDataV3
		return ret
	}

	return o.PoolTokens
}

// GetPoolTokensOk returns a tuple with the PoolTokens field value
// and a boolean to check if the value has been set.
func (o *AmmTransactionDataV3) GetPoolTokensOk() ([]AmmTransferDataV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoolTokens, true
}

// SetPoolTokens sets field value
func (o *AmmTransactionDataV3) SetPoolTokens(v []AmmTransferDataV3) {
	o.PoolTokens = v
}

// GetLpToken returns the LpToken field value
func (o *AmmTransactionDataV3) GetLpToken() AmmTransferDataV3 {
	if o == nil {
		var ret AmmTransferDataV3
		return ret
	}

	return o.LpToken
}

// GetLpTokenOk returns a tuple with the LpToken field value
// and a boolean to check if the value has been set.
func (o *AmmTransactionDataV3) GetLpTokenOk() (*AmmTransferDataV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LpToken, true
}

// SetLpToken sets field value
func (o *AmmTransactionDataV3) SetLpToken(v AmmTransferDataV3) {
	o.LpToken = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AmmTransactionDataV3) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AmmTransactionDataV3) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AmmTransactionDataV3) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AmmTransactionDataV3) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AmmTransactionDataV3) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AmmTransactionDataV3) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetBlockId returns the BlockId field value
func (o *AmmTransactionDataV3) GetBlockId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BlockId
}

// GetBlockIdOk returns a tuple with the BlockId field value
// and a boolean to check if the value has been set.
func (o *AmmTransactionDataV3) GetBlockIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockId, true
}

// SetBlockId sets field value
func (o *AmmTransactionDataV3) SetBlockId(v int64) {
	o.BlockId = v
}

// GetIndexInBlock returns the IndexInBlock field value
func (o *AmmTransactionDataV3) GetIndexInBlock() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IndexInBlock
}

// GetIndexInBlockOk returns a tuple with the IndexInBlock field value
// and a boolean to check if the value has been set.
func (o *AmmTransactionDataV3) GetIndexInBlockOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexInBlock, true
}

// SetIndexInBlock sets field value
func (o *AmmTransactionDataV3) SetIndexInBlock(v int32) {
	o.IndexInBlock = v
}

// GetStorageInfos returns the StorageInfos field value if set, zero value otherwise.
func (o *AmmTransactionDataV3) GetStorageInfos() []StorageInfo {
	if o == nil || o.StorageInfos == nil {
		var ret []StorageInfo
		return ret
	}
	return o.StorageInfos
}

// GetStorageInfosOk returns a tuple with the StorageInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmmTransactionDataV3) GetStorageInfosOk() ([]StorageInfo, bool) {
	if o == nil || o.StorageInfos == nil {
		return nil, false
	}
	return o.StorageInfos, true
}

// HasStorageInfos returns a boolean if a field has been set.
func (o *AmmTransactionDataV3) HasStorageInfos() bool {
	if o != nil && o.StorageInfos != nil {
		return true
	}

	return false
}

// SetStorageInfos gets a reference to the given []StorageInfo and assigns it to the StorageInfos field.
func (o *AmmTransactionDataV3) SetStorageInfos(v []StorageInfo) {
	o.StorageInfos = v
}

func (o AmmTransactionDataV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["txType"] = o.TxType
	}
	if true {
		toSerialize["txStatus"] = o.TxStatus
	}
	if true {
		toSerialize["ammPoolAddress"] = o.AmmPoolAddress
	}
	if true {
		toSerialize["ammLayerType"] = o.AmmLayerType
	}
	if true {
		toSerialize["poolTokens"] = o.PoolTokens
	}
	if true {
		toSerialize["lpToken"] = o.LpToken
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["blockId"] = o.BlockId
	}
	if true {
		toSerialize["indexInBlock"] = o.IndexInBlock
	}
	if o.StorageInfos != nil {
		toSerialize["storageInfos"] = o.StorageInfos
	}
	return json.Marshal(toSerialize)
}

type NullableAmmTransactionDataV3 struct {
	value *AmmTransactionDataV3
	isSet bool
}

func (v NullableAmmTransactionDataV3) Get() *AmmTransactionDataV3 {
	return v.value
}

func (v *NullableAmmTransactionDataV3) Set(val *AmmTransactionDataV3) {
	v.value = val
	v.isSet = true
}

func (v NullableAmmTransactionDataV3) IsSet() bool {
	return v.isSet
}

func (v *NullableAmmTransactionDataV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmmTransactionDataV3(val *AmmTransactionDataV3) *NullableAmmTransactionDataV3 {
	return &NullableAmmTransactionDataV3{value: val, isSet: true}
}

func (v NullableAmmTransactionDataV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmmTransactionDataV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
