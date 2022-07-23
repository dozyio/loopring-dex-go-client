# UserAccountTxData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique ID | 
**Hash** | **string** | hash | 
**Owner** | **string** | Owner address of the Tx | 
**TxHash** | **string** | Transaction hash | 
**FeeTokenSymbol** | **string** | Fee token symbol | 
**FeeAmount** | **string** | Fee amount in wei | 
**Status** | **string** | Current status | 
**Progress** | **string** | Progress | 
**Timestamp** | **int64** | Create time | 
**BlockNum** | **int64** | Block height | 
**UpdatedAt** | **int64** | Update time | 
**BlockId** | **int64** |  | 
**IndexInBlock** | **int32** |  | 

## Methods

### NewUserAccountTxData

`func NewUserAccountTxData(id int64, hash string, owner string, txHash string, feeTokenSymbol string, feeAmount string, status string, progress string, timestamp int64, blockNum int64, updatedAt int64, blockId int64, indexInBlock int32, ) *UserAccountTxData`

NewUserAccountTxData instantiates a new UserAccountTxData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountTxDataWithDefaults

`func NewUserAccountTxDataWithDefaults() *UserAccountTxData`

NewUserAccountTxDataWithDefaults instantiates a new UserAccountTxData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserAccountTxData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAccountTxData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAccountTxData) SetId(v int64)`

SetId sets Id field to given value.


### GetHash

`func (o *UserAccountTxData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *UserAccountTxData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *UserAccountTxData) SetHash(v string)`

SetHash sets Hash field to given value.


### GetOwner

`func (o *UserAccountTxData) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UserAccountTxData) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UserAccountTxData) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTxHash

`func (o *UserAccountTxData) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *UserAccountTxData) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *UserAccountTxData) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetFeeTokenSymbol

`func (o *UserAccountTxData) GetFeeTokenSymbol() string`

GetFeeTokenSymbol returns the FeeTokenSymbol field if non-nil, zero value otherwise.

### GetFeeTokenSymbolOk

`func (o *UserAccountTxData) GetFeeTokenSymbolOk() (*string, bool)`

GetFeeTokenSymbolOk returns a tuple with the FeeTokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenSymbol

`func (o *UserAccountTxData) SetFeeTokenSymbol(v string)`

SetFeeTokenSymbol sets FeeTokenSymbol field to given value.


### GetFeeAmount

`func (o *UserAccountTxData) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *UserAccountTxData) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *UserAccountTxData) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetStatus

`func (o *UserAccountTxData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserAccountTxData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserAccountTxData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *UserAccountTxData) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *UserAccountTxData) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *UserAccountTxData) SetProgress(v string)`

SetProgress sets Progress field to given value.


### GetTimestamp

`func (o *UserAccountTxData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserAccountTxData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserAccountTxData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetBlockNum

`func (o *UserAccountTxData) GetBlockNum() int64`

GetBlockNum returns the BlockNum field if non-nil, zero value otherwise.

### GetBlockNumOk

`func (o *UserAccountTxData) GetBlockNumOk() (*int64, bool)`

GetBlockNumOk returns a tuple with the BlockNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNum

`func (o *UserAccountTxData) SetBlockNum(v int64)`

SetBlockNum sets BlockNum field to given value.


### GetUpdatedAt

`func (o *UserAccountTxData) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserAccountTxData) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserAccountTxData) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBlockId

`func (o *UserAccountTxData) GetBlockId() int64`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *UserAccountTxData) GetBlockIdOk() (*int64, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *UserAccountTxData) SetBlockId(v int64)`

SetBlockId sets BlockId field to given value.


### GetIndexInBlock

`func (o *UserAccountTxData) GetIndexInBlock() int32`

GetIndexInBlock returns the IndexInBlock field if non-nil, zero value otherwise.

### GetIndexInBlockOk

`func (o *UserAccountTxData) GetIndexInBlockOk() (*int32, bool)`

GetIndexInBlockOk returns a tuple with the IndexInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexInBlock

`func (o *UserAccountTxData) SetIndexInBlock(v int32)`

SetIndexInBlock sets IndexInBlock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


