# DepositData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique ID | 
**Hash** | **string** | hash | 
**Symbol** | **string** | Token symbol | 
**Amount** | **string** | Amount requested by the user | 
**TxHash** | **string** | Transaction hash | 
**Status** | **string** | Current status | 
**Progress** | **string** | Progress | 
**Timestamp** | **int64** | Create time | 
**BlockNum** | **int64** | Block height | 
**UpdatedAt** | **int64** | Update time | 
**BlockId** | **int64** |  | 
**IndexInBlock** | **int32** |  | 

## Methods

### NewDepositData

`func NewDepositData(id int64, hash string, symbol string, amount string, txHash string, status string, progress string, timestamp int64, blockNum int64, updatedAt int64, blockId int64, indexInBlock int32, ) *DepositData`

NewDepositData instantiates a new DepositData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositDataWithDefaults

`func NewDepositDataWithDefaults() *DepositData`

NewDepositDataWithDefaults instantiates a new DepositData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DepositData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DepositData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DepositData) SetId(v int64)`

SetId sets Id field to given value.


### GetHash

`func (o *DepositData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *DepositData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *DepositData) SetHash(v string)`

SetHash sets Hash field to given value.


### GetSymbol

`func (o *DepositData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *DepositData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *DepositData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetAmount

`func (o *DepositData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DepositData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DepositData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTxHash

`func (o *DepositData) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *DepositData) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *DepositData) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetStatus

`func (o *DepositData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DepositData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DepositData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *DepositData) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *DepositData) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *DepositData) SetProgress(v string)`

SetProgress sets Progress field to given value.


### GetTimestamp

`func (o *DepositData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DepositData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DepositData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetBlockNum

`func (o *DepositData) GetBlockNum() int64`

GetBlockNum returns the BlockNum field if non-nil, zero value otherwise.

### GetBlockNumOk

`func (o *DepositData) GetBlockNumOk() (*int64, bool)`

GetBlockNumOk returns a tuple with the BlockNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNum

`func (o *DepositData) SetBlockNum(v int64)`

SetBlockNum sets BlockNum field to given value.


### GetUpdatedAt

`func (o *DepositData) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DepositData) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DepositData) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBlockId

`func (o *DepositData) GetBlockId() int64`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *DepositData) GetBlockIdOk() (*int64, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *DepositData) SetBlockId(v int64)`

SetBlockId sets BlockId field to given value.


### GetIndexInBlock

`func (o *DepositData) GetIndexInBlock() int32`

GetIndexInBlock returns the IndexInBlock field if non-nil, zero value otherwise.

### GetIndexInBlockOk

`func (o *DepositData) GetIndexInBlockOk() (*int32, bool)`

GetIndexInBlockOk returns a tuple with the IndexInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexInBlock

`func (o *DepositData) SetIndexInBlock(v int32)`

SetIndexInBlock sets IndexInBlock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


