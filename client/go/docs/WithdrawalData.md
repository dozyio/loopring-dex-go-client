# WithdrawalData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique ID | 
**TxType** | **string** | User transaction type | 
**Hash** | **string** | hash | 
**Symbol** | **string** | Token symbol | 
**Amount** | **string** | Amount requested by the user | 
**TxHash** | **string** | Transaction hash | 
**FeeTokenSymbol** | **string** | Fee token symbol | 
**FeeAmount** | **string** | Fee amount in wei | 
**Status** | **string** | Current status | 
**Progress** | **string** | Progress | 
**Timestamp** | **int64** | Create time | 
**BlockNum** | **int64** | Block height | 
**UpdatedAt** | **int64** | Update time | 
**DistributeHash** | **string** | Distribute Hash | 
**RequestId** | Pointer to **int64** | Request Id | [optional] 
**FastStatus** | Pointer to **string** | Fast withdrawal status | [optional] 
**BlockId** | **int64** |  | 
**IndexInBlock** | **int32** |  | 

## Methods

### NewWithdrawalData

`func NewWithdrawalData(id int64, txType string, hash string, symbol string, amount string, txHash string, feeTokenSymbol string, feeAmount string, status string, progress string, timestamp int64, blockNum int64, updatedAt int64, distributeHash string, blockId int64, indexInBlock int32, ) *WithdrawalData`

NewWithdrawalData instantiates a new WithdrawalData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawalDataWithDefaults

`func NewWithdrawalDataWithDefaults() *WithdrawalData`

NewWithdrawalDataWithDefaults instantiates a new WithdrawalData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WithdrawalData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WithdrawalData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WithdrawalData) SetId(v int64)`

SetId sets Id field to given value.


### GetTxType

`func (o *WithdrawalData) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *WithdrawalData) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *WithdrawalData) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetHash

`func (o *WithdrawalData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *WithdrawalData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *WithdrawalData) SetHash(v string)`

SetHash sets Hash field to given value.


### GetSymbol

`func (o *WithdrawalData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *WithdrawalData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *WithdrawalData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetAmount

`func (o *WithdrawalData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WithdrawalData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WithdrawalData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTxHash

`func (o *WithdrawalData) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *WithdrawalData) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *WithdrawalData) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetFeeTokenSymbol

`func (o *WithdrawalData) GetFeeTokenSymbol() string`

GetFeeTokenSymbol returns the FeeTokenSymbol field if non-nil, zero value otherwise.

### GetFeeTokenSymbolOk

`func (o *WithdrawalData) GetFeeTokenSymbolOk() (*string, bool)`

GetFeeTokenSymbolOk returns a tuple with the FeeTokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenSymbol

`func (o *WithdrawalData) SetFeeTokenSymbol(v string)`

SetFeeTokenSymbol sets FeeTokenSymbol field to given value.


### GetFeeAmount

`func (o *WithdrawalData) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *WithdrawalData) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *WithdrawalData) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetStatus

`func (o *WithdrawalData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WithdrawalData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WithdrawalData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *WithdrawalData) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *WithdrawalData) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *WithdrawalData) SetProgress(v string)`

SetProgress sets Progress field to given value.


### GetTimestamp

`func (o *WithdrawalData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WithdrawalData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WithdrawalData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetBlockNum

`func (o *WithdrawalData) GetBlockNum() int64`

GetBlockNum returns the BlockNum field if non-nil, zero value otherwise.

### GetBlockNumOk

`func (o *WithdrawalData) GetBlockNumOk() (*int64, bool)`

GetBlockNumOk returns a tuple with the BlockNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNum

`func (o *WithdrawalData) SetBlockNum(v int64)`

SetBlockNum sets BlockNum field to given value.


### GetUpdatedAt

`func (o *WithdrawalData) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WithdrawalData) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WithdrawalData) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDistributeHash

`func (o *WithdrawalData) GetDistributeHash() string`

GetDistributeHash returns the DistributeHash field if non-nil, zero value otherwise.

### GetDistributeHashOk

`func (o *WithdrawalData) GetDistributeHashOk() (*string, bool)`

GetDistributeHashOk returns a tuple with the DistributeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeHash

`func (o *WithdrawalData) SetDistributeHash(v string)`

SetDistributeHash sets DistributeHash field to given value.


### GetRequestId

`func (o *WithdrawalData) GetRequestId() int64`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WithdrawalData) GetRequestIdOk() (*int64, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WithdrawalData) SetRequestId(v int64)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *WithdrawalData) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetFastStatus

`func (o *WithdrawalData) GetFastStatus() string`

GetFastStatus returns the FastStatus field if non-nil, zero value otherwise.

### GetFastStatusOk

`func (o *WithdrawalData) GetFastStatusOk() (*string, bool)`

GetFastStatusOk returns a tuple with the FastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastStatus

`func (o *WithdrawalData) SetFastStatus(v string)`

SetFastStatus sets FastStatus field to given value.

### HasFastStatus

`func (o *WithdrawalData) HasFastStatus() bool`

HasFastStatus returns a boolean if a field has been set.

### GetBlockId

`func (o *WithdrawalData) GetBlockId() int64`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *WithdrawalData) GetBlockIdOk() (*int64, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *WithdrawalData) SetBlockId(v int64)`

SetBlockId sets BlockId field to given value.


### GetIndexInBlock

`func (o *WithdrawalData) GetIndexInBlock() int32`

GetIndexInBlock returns the IndexInBlock field if non-nil, zero value otherwise.

### GetIndexInBlockOk

`func (o *WithdrawalData) GetIndexInBlockOk() (*int32, bool)`

GetIndexInBlockOk returns a tuple with the IndexInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexInBlock

`func (o *WithdrawalData) SetIndexInBlock(v int32)`

SetIndexInBlock sets IndexInBlock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


