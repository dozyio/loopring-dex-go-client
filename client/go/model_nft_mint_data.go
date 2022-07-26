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

// NftMintData model.NftMintData.description
type NftMintData struct {
	// field.NftMintData.id
	Id int64 `json:"id"`
	// field.NftMintData.requestId
	RequestId int64 `json:"requestId"`
	// field.NftMintData.hash
	Hash string `json:"hash"`
	// field.NftMintData.txHash
	TxHash string `json:"txHash"`
	// field.NftMintData.accountId
	AccountId int64 `json:"accountId"`
	// field.NftMintData.owner
	Owner string `json:"owner"`
	// field.NftMintData.status
	Status string `json:"status"`
	// field.NftMintData.nftData
	NftData string `json:"nftData"`
	// field.NftMintData.amount
	Amount string `json:"amount"`
	// field.NftMintData.feeTokenSymbol
	FeeTokenSymbol string `json:"feeTokenSymbol"`
	// field.NftMintData.feeAmount
	FeeAmount string `json:"feeAmount"`
	// field.NftMintData.createdAt
	CreatedAt int64 `json:"createdAt"`
	// field.NftMintData.updatedAt
	UpdatedAt int64 `json:"updatedAt"`
	// field.NftMintData.memo
	Memo string `json:"memo"`
	// field.NftMintData.minterId
	MinterId int64 `json:"minterId"`
	// field.NftMintData.minterAddress
	MinterAddress string       `json:"minterAddress"`
	BlockIdInfo   *BlockIdInfo `json:"blockIdInfo,omitempty"`
	StorageInfo   *StorageInfo `json:"storageInfo,omitempty"`
}

// NewNftMintData instantiates a new NftMintData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftMintData(id int64, requestId int64, hash string, txHash string, accountId int64, owner string, status string, nftData string, amount string, feeTokenSymbol string, feeAmount string, createdAt int64, updatedAt int64, memo string, minterId int64, minterAddress string) *NftMintData {
	this := NftMintData{}
	this.Id = id
	this.RequestId = requestId
	this.Hash = hash
	this.TxHash = txHash
	this.AccountId = accountId
	this.Owner = owner
	this.Status = status
	this.NftData = nftData
	this.Amount = amount
	this.FeeTokenSymbol = feeTokenSymbol
	this.FeeAmount = feeAmount
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Memo = memo
	this.MinterId = minterId
	this.MinterAddress = minterAddress
	return &this
}

// NewNftMintDataWithDefaults instantiates a new NftMintData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftMintDataWithDefaults() *NftMintData {
	this := NftMintData{}
	return &this
}

// GetId returns the Id field value
func (o *NftMintData) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NftMintData) SetId(v int64) {
	o.Id = v
}

// GetRequestId returns the RequestId field value
func (o *NftMintData) GetRequestId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetRequestIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *NftMintData) SetRequestId(v int64) {
	o.RequestId = v
}

// GetHash returns the Hash field value
func (o *NftMintData) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *NftMintData) SetHash(v string) {
	o.Hash = v
}

// GetTxHash returns the TxHash field value
func (o *NftMintData) GetTxHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetTxHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxHash, true
}

// SetTxHash sets field value
func (o *NftMintData) SetTxHash(v string) {
	o.TxHash = v
}

// GetAccountId returns the AccountId field value
func (o *NftMintData) GetAccountId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetAccountIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NftMintData) SetAccountId(v int64) {
	o.AccountId = v
}

// GetOwner returns the Owner field value
func (o *NftMintData) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *NftMintData) SetOwner(v string) {
	o.Owner = v
}

// GetStatus returns the Status field value
func (o *NftMintData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NftMintData) SetStatus(v string) {
	o.Status = v
}

// GetNftData returns the NftData field value
func (o *NftMintData) GetNftData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NftData
}

// GetNftDataOk returns a tuple with the NftData field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetNftDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NftData, true
}

// SetNftData sets field value
func (o *NftMintData) SetNftData(v string) {
	o.NftData = v
}

// GetAmount returns the Amount field value
func (o *NftMintData) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *NftMintData) SetAmount(v string) {
	o.Amount = v
}

// GetFeeTokenSymbol returns the FeeTokenSymbol field value
func (o *NftMintData) GetFeeTokenSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeTokenSymbol
}

// GetFeeTokenSymbolOk returns a tuple with the FeeTokenSymbol field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetFeeTokenSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeTokenSymbol, true
}

// SetFeeTokenSymbol sets field value
func (o *NftMintData) SetFeeTokenSymbol(v string) {
	o.FeeTokenSymbol = v
}

// GetFeeAmount returns the FeeAmount field value
func (o *NftMintData) GetFeeAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetFeeAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeAmount, true
}

// SetFeeAmount sets field value
func (o *NftMintData) SetFeeAmount(v string) {
	o.FeeAmount = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NftMintData) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NftMintData) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *NftMintData) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *NftMintData) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetMemo returns the Memo field value
func (o *NftMintData) GetMemo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Memo
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memo, true
}

// SetMemo sets field value
func (o *NftMintData) SetMemo(v string) {
	o.Memo = v
}

// GetMinterId returns the MinterId field value
func (o *NftMintData) GetMinterId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinterId
}

// GetMinterIdOk returns a tuple with the MinterId field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetMinterIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinterId, true
}

// SetMinterId sets field value
func (o *NftMintData) SetMinterId(v int64) {
	o.MinterId = v
}

// GetMinterAddress returns the MinterAddress field value
func (o *NftMintData) GetMinterAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinterAddress
}

// GetMinterAddressOk returns a tuple with the MinterAddress field value
// and a boolean to check if the value has been set.
func (o *NftMintData) GetMinterAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinterAddress, true
}

// SetMinterAddress sets field value
func (o *NftMintData) SetMinterAddress(v string) {
	o.MinterAddress = v
}

// GetBlockIdInfo returns the BlockIdInfo field value if set, zero value otherwise.
func (o *NftMintData) GetBlockIdInfo() BlockIdInfo {
	if o == nil || o.BlockIdInfo == nil {
		var ret BlockIdInfo
		return ret
	}
	return *o.BlockIdInfo
}

// GetBlockIdInfoOk returns a tuple with the BlockIdInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftMintData) GetBlockIdInfoOk() (*BlockIdInfo, bool) {
	if o == nil || o.BlockIdInfo == nil {
		return nil, false
	}
	return o.BlockIdInfo, true
}

// HasBlockIdInfo returns a boolean if a field has been set.
func (o *NftMintData) HasBlockIdInfo() bool {
	if o != nil && o.BlockIdInfo != nil {
		return true
	}

	return false
}

// SetBlockIdInfo gets a reference to the given BlockIdInfo and assigns it to the BlockIdInfo field.
func (o *NftMintData) SetBlockIdInfo(v BlockIdInfo) {
	o.BlockIdInfo = &v
}

// GetStorageInfo returns the StorageInfo field value if set, zero value otherwise.
func (o *NftMintData) GetStorageInfo() StorageInfo {
	if o == nil || o.StorageInfo == nil {
		var ret StorageInfo
		return ret
	}
	return *o.StorageInfo
}

// GetStorageInfoOk returns a tuple with the StorageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftMintData) GetStorageInfoOk() (*StorageInfo, bool) {
	if o == nil || o.StorageInfo == nil {
		return nil, false
	}
	return o.StorageInfo, true
}

// HasStorageInfo returns a boolean if a field has been set.
func (o *NftMintData) HasStorageInfo() bool {
	if o != nil && o.StorageInfo != nil {
		return true
	}

	return false
}

// SetStorageInfo gets a reference to the given StorageInfo and assigns it to the StorageInfo field.
func (o *NftMintData) SetStorageInfo(v StorageInfo) {
	o.StorageInfo = &v
}

func (o NftMintData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["requestId"] = o.RequestId
	}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["txHash"] = o.TxHash
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["nftData"] = o.NftData
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["feeTokenSymbol"] = o.FeeTokenSymbol
	}
	if true {
		toSerialize["feeAmount"] = o.FeeAmount
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["memo"] = o.Memo
	}
	if true {
		toSerialize["minterId"] = o.MinterId
	}
	if true {
		toSerialize["minterAddress"] = o.MinterAddress
	}
	if o.BlockIdInfo != nil {
		toSerialize["blockIdInfo"] = o.BlockIdInfo
	}
	if o.StorageInfo != nil {
		toSerialize["storageInfo"] = o.StorageInfo
	}
	return json.Marshal(toSerialize)
}

type NullableNftMintData struct {
	value *NftMintData
	isSet bool
}

func (v NullableNftMintData) Get() *NftMintData {
	return v.value
}

func (v *NullableNftMintData) Set(val *NftMintData) {
	v.value = val
	v.isSet = true
}

func (v NullableNftMintData) IsSet() bool {
	return v.isSet
}

func (v *NullableNftMintData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftMintData(val *NftMintData) *NullableNftMintData {
	return &NullableNftMintData{value: val, isSet: true}
}

func (v NullableNftMintData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftMintData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
