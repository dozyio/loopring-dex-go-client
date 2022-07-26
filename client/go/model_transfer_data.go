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

// TransferData struct for TransferData
type TransferData struct {
	// Unique ID
	Id int64 `json:"id"`
	// hash
	Hash string `json:"hash"`
	// User transaction type
	TxType string `json:"txType"`
	// Token symbol
	Symbol string `json:"symbol"`
	// Amount requested by the user
	Amount string `json:"amount"`
	// field.TxData.senderAddress
	SenderAddress *string `json:"senderAddress,omitempty"`
	// Receiver ID
	Receiver *int64 `json:"receiver,omitempty"`
	// The transfer receiver's address
	ReceiverAddress *string `json:"receiverAddress,omitempty"`
	// Fee token symbol
	FeeTokenSymbol string `json:"feeTokenSymbol"`
	// Fee amount in wei
	FeeAmount string `json:"feeAmount"`
	// Current status
	Status string `json:"status"`
	// Progress
	Progress string `json:"progress"`
	// Create time
	Timestamp int64 `json:"timestamp"`
	// Update time
	UpdatedAt int64 `json:"updatedAt"`
	// field.TxData.memo
	Memo         *string      `json:"memo,omitempty"`
	BlockId      int64        `json:"blockId"`
	IndexInBlock int32        `json:"indexInBlock"`
	StorageInfo  *StorageInfo `json:"storageInfo,omitempty"`
}

// NewTransferData instantiates a new TransferData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferData(id int64, hash string, txType string, symbol string, amount string, feeTokenSymbol string, feeAmount string, status string, progress string, timestamp int64, updatedAt int64, blockId int64, indexInBlock int32) *TransferData {
	this := TransferData{}
	this.Id = id
	this.Hash = hash
	this.TxType = txType
	this.Symbol = symbol
	this.Amount = amount
	this.FeeTokenSymbol = feeTokenSymbol
	this.FeeAmount = feeAmount
	this.Status = status
	this.Progress = progress
	this.Timestamp = timestamp
	this.UpdatedAt = updatedAt
	this.BlockId = blockId
	this.IndexInBlock = indexInBlock
	return &this
}

// NewTransferDataWithDefaults instantiates a new TransferData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferDataWithDefaults() *TransferData {
	this := TransferData{}
	return &this
}

// GetId returns the Id field value
func (o *TransferData) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransferData) SetId(v int64) {
	o.Id = v
}

// GetHash returns the Hash field value
func (o *TransferData) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *TransferData) SetHash(v string) {
	o.Hash = v
}

// GetTxType returns the TxType field value
func (o *TransferData) GetTxType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxType
}

// GetTxTypeOk returns a tuple with the TxType field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetTxTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxType, true
}

// SetTxType sets field value
func (o *TransferData) SetTxType(v string) {
	o.TxType = v
}

// GetSymbol returns the Symbol field value
func (o *TransferData) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *TransferData) SetSymbol(v string) {
	o.Symbol = v
}

// GetAmount returns the Amount field value
func (o *TransferData) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransferData) SetAmount(v string) {
	o.Amount = v
}

// GetSenderAddress returns the SenderAddress field value if set, zero value otherwise.
func (o *TransferData) GetSenderAddress() string {
	if o == nil || o.SenderAddress == nil {
		var ret string
		return ret
	}
	return *o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetSenderAddressOk() (*string, bool) {
	if o == nil || o.SenderAddress == nil {
		return nil, false
	}
	return o.SenderAddress, true
}

// HasSenderAddress returns a boolean if a field has been set.
func (o *TransferData) HasSenderAddress() bool {
	if o != nil && o.SenderAddress != nil {
		return true
	}

	return false
}

// SetSenderAddress gets a reference to the given string and assigns it to the SenderAddress field.
func (o *TransferData) SetSenderAddress(v string) {
	o.SenderAddress = &v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise.
func (o *TransferData) GetReceiver() int64 {
	if o == nil || o.Receiver == nil {
		var ret int64
		return ret
	}
	return *o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetReceiverOk() (*int64, bool) {
	if o == nil || o.Receiver == nil {
		return nil, false
	}
	return o.Receiver, true
}

// HasReceiver returns a boolean if a field has been set.
func (o *TransferData) HasReceiver() bool {
	if o != nil && o.Receiver != nil {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given int64 and assigns it to the Receiver field.
func (o *TransferData) SetReceiver(v int64) {
	o.Receiver = &v
}

// GetReceiverAddress returns the ReceiverAddress field value if set, zero value otherwise.
func (o *TransferData) GetReceiverAddress() string {
	if o == nil || o.ReceiverAddress == nil {
		var ret string
		return ret
	}
	return *o.ReceiverAddress
}

// GetReceiverAddressOk returns a tuple with the ReceiverAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetReceiverAddressOk() (*string, bool) {
	if o == nil || o.ReceiverAddress == nil {
		return nil, false
	}
	return o.ReceiverAddress, true
}

// HasReceiverAddress returns a boolean if a field has been set.
func (o *TransferData) HasReceiverAddress() bool {
	if o != nil && o.ReceiverAddress != nil {
		return true
	}

	return false
}

// SetReceiverAddress gets a reference to the given string and assigns it to the ReceiverAddress field.
func (o *TransferData) SetReceiverAddress(v string) {
	o.ReceiverAddress = &v
}

// GetFeeTokenSymbol returns the FeeTokenSymbol field value
func (o *TransferData) GetFeeTokenSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeTokenSymbol
}

// GetFeeTokenSymbolOk returns a tuple with the FeeTokenSymbol field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetFeeTokenSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeTokenSymbol, true
}

// SetFeeTokenSymbol sets field value
func (o *TransferData) SetFeeTokenSymbol(v string) {
	o.FeeTokenSymbol = v
}

// GetFeeAmount returns the FeeAmount field value
func (o *TransferData) GetFeeAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetFeeAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeAmount, true
}

// SetFeeAmount sets field value
func (o *TransferData) SetFeeAmount(v string) {
	o.FeeAmount = v
}

// GetStatus returns the Status field value
func (o *TransferData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransferData) SetStatus(v string) {
	o.Status = v
}

// GetProgress returns the Progress field value
func (o *TransferData) GetProgress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetProgressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *TransferData) SetProgress(v string) {
	o.Progress = v
}

// GetTimestamp returns the Timestamp field value
func (o *TransferData) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *TransferData) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TransferData) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TransferData) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *TransferData) GetMemo() string {
	if o == nil || o.Memo == nil {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetMemoOk() (*string, bool) {
	if o == nil || o.Memo == nil {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *TransferData) HasMemo() bool {
	if o != nil && o.Memo != nil {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *TransferData) SetMemo(v string) {
	o.Memo = &v
}

// GetBlockId returns the BlockId field value
func (o *TransferData) GetBlockId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BlockId
}

// GetBlockIdOk returns a tuple with the BlockId field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetBlockIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockId, true
}

// SetBlockId sets field value
func (o *TransferData) SetBlockId(v int64) {
	o.BlockId = v
}

// GetIndexInBlock returns the IndexInBlock field value
func (o *TransferData) GetIndexInBlock() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IndexInBlock
}

// GetIndexInBlockOk returns a tuple with the IndexInBlock field value
// and a boolean to check if the value has been set.
func (o *TransferData) GetIndexInBlockOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexInBlock, true
}

// SetIndexInBlock sets field value
func (o *TransferData) SetIndexInBlock(v int32) {
	o.IndexInBlock = v
}

// GetStorageInfo returns the StorageInfo field value if set, zero value otherwise.
func (o *TransferData) GetStorageInfo() StorageInfo {
	if o == nil || o.StorageInfo == nil {
		var ret StorageInfo
		return ret
	}
	return *o.StorageInfo
}

// GetStorageInfoOk returns a tuple with the StorageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferData) GetStorageInfoOk() (*StorageInfo, bool) {
	if o == nil || o.StorageInfo == nil {
		return nil, false
	}
	return o.StorageInfo, true
}

// HasStorageInfo returns a boolean if a field has been set.
func (o *TransferData) HasStorageInfo() bool {
	if o != nil && o.StorageInfo != nil {
		return true
	}

	return false
}

// SetStorageInfo gets a reference to the given StorageInfo and assigns it to the StorageInfo field.
func (o *TransferData) SetStorageInfo(v StorageInfo) {
	o.StorageInfo = &v
}

func (o TransferData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["txType"] = o.TxType
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.SenderAddress != nil {
		toSerialize["senderAddress"] = o.SenderAddress
	}
	if o.Receiver != nil {
		toSerialize["receiver"] = o.Receiver
	}
	if o.ReceiverAddress != nil {
		toSerialize["receiverAddress"] = o.ReceiverAddress
	}
	if true {
		toSerialize["feeTokenSymbol"] = o.FeeTokenSymbol
	}
	if true {
		toSerialize["feeAmount"] = o.FeeAmount
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["progress"] = o.Progress
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Memo != nil {
		toSerialize["memo"] = o.Memo
	}
	if true {
		toSerialize["blockId"] = o.BlockId
	}
	if true {
		toSerialize["indexInBlock"] = o.IndexInBlock
	}
	if o.StorageInfo != nil {
		toSerialize["storageInfo"] = o.StorageInfo
	}
	return json.Marshal(toSerialize)
}

type NullableTransferData struct {
	value *TransferData
	isSet bool
}

func (v NullableTransferData) Get() *TransferData {
	return v.value
}

func (v *NullableTransferData) Set(val *TransferData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferData(val *TransferData) *NullableTransferData {
	return &NullableTransferData{value: val, isSet: true}
}

func (v NullableTransferData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
