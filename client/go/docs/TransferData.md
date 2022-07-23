# TransferData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique ID | 
**Hash** | **string** | hash | 
**TxType** | **string** | User transaction type | 
**Symbol** | **string** | Token symbol | 
**Amount** | **string** | Amount requested by the user | 
**SenderAddress** | Pointer to **string** | field.TxData.senderAddress | [optional] 
**Receiver** | Pointer to **int64** | Receiver ID | [optional] 
**ReceiverAddress** | Pointer to **string** | The transfer receiver&#39;s address | [optional] 
**FeeTokenSymbol** | **string** | Fee token symbol | 
**FeeAmount** | **string** | Fee amount in wei | 
**Status** | **string** | Current status | 
**Progress** | **string** | Progress | 
**Timestamp** | **int64** | Create time | 
**UpdatedAt** | **int64** | Update time | 
**Memo** | Pointer to **string** | field.TxData.memo | [optional] 
**BlockId** | **int64** |  | 
**IndexInBlock** | **int32** |  | 
**StorageInfo** | Pointer to [**StorageInfo**](StorageInfo.md) |  | [optional] 

## Methods

### NewTransferData

`func NewTransferData(id int64, hash string, txType string, symbol string, amount string, feeTokenSymbol string, feeAmount string, status string, progress string, timestamp int64, updatedAt int64, blockId int64, indexInBlock int32, ) *TransferData`

NewTransferData instantiates a new TransferData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferDataWithDefaults

`func NewTransferDataWithDefaults() *TransferData`

NewTransferDataWithDefaults instantiates a new TransferData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransferData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferData) SetId(v int64)`

SetId sets Id field to given value.


### GetHash

`func (o *TransferData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *TransferData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *TransferData) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTxType

`func (o *TransferData) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *TransferData) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *TransferData) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetSymbol

`func (o *TransferData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TransferData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TransferData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetAmount

`func (o *TransferData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetSenderAddress

`func (o *TransferData) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *TransferData) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *TransferData) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.

### HasSenderAddress

`func (o *TransferData) HasSenderAddress() bool`

HasSenderAddress returns a boolean if a field has been set.

### GetReceiver

`func (o *TransferData) GetReceiver() int64`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *TransferData) GetReceiverOk() (*int64, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *TransferData) SetReceiver(v int64)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *TransferData) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetReceiverAddress

`func (o *TransferData) GetReceiverAddress() string`

GetReceiverAddress returns the ReceiverAddress field if non-nil, zero value otherwise.

### GetReceiverAddressOk

`func (o *TransferData) GetReceiverAddressOk() (*string, bool)`

GetReceiverAddressOk returns a tuple with the ReceiverAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverAddress

`func (o *TransferData) SetReceiverAddress(v string)`

SetReceiverAddress sets ReceiverAddress field to given value.

### HasReceiverAddress

`func (o *TransferData) HasReceiverAddress() bool`

HasReceiverAddress returns a boolean if a field has been set.

### GetFeeTokenSymbol

`func (o *TransferData) GetFeeTokenSymbol() string`

GetFeeTokenSymbol returns the FeeTokenSymbol field if non-nil, zero value otherwise.

### GetFeeTokenSymbolOk

`func (o *TransferData) GetFeeTokenSymbolOk() (*string, bool)`

GetFeeTokenSymbolOk returns a tuple with the FeeTokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenSymbol

`func (o *TransferData) SetFeeTokenSymbol(v string)`

SetFeeTokenSymbol sets FeeTokenSymbol field to given value.


### GetFeeAmount

`func (o *TransferData) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *TransferData) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *TransferData) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetStatus

`func (o *TransferData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *TransferData) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *TransferData) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *TransferData) SetProgress(v string)`

SetProgress sets Progress field to given value.


### GetTimestamp

`func (o *TransferData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TransferData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TransferData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetUpdatedAt

`func (o *TransferData) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransferData) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransferData) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMemo

`func (o *TransferData) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransferData) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransferData) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransferData) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetBlockId

`func (o *TransferData) GetBlockId() int64`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *TransferData) GetBlockIdOk() (*int64, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *TransferData) SetBlockId(v int64)`

SetBlockId sets BlockId field to given value.


### GetIndexInBlock

`func (o *TransferData) GetIndexInBlock() int32`

GetIndexInBlock returns the IndexInBlock field if non-nil, zero value otherwise.

### GetIndexInBlockOk

`func (o *TransferData) GetIndexInBlockOk() (*int32, bool)`

GetIndexInBlockOk returns a tuple with the IndexInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexInBlock

`func (o *TransferData) SetIndexInBlock(v int32)`

SetIndexInBlock sets IndexInBlock field to given value.


### GetStorageInfo

`func (o *TransferData) GetStorageInfo() StorageInfo`

GetStorageInfo returns the StorageInfo field if non-nil, zero value otherwise.

### GetStorageInfoOk

`func (o *TransferData) GetStorageInfoOk() (*StorageInfo, bool)`

GetStorageInfoOk returns a tuple with the StorageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInfo

`func (o *TransferData) SetStorageInfo(v StorageInfo)`

SetStorageInfo sets StorageInfo field to given value.

### HasStorageInfo

`func (o *TransferData) HasStorageInfo() bool`

HasStorageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


