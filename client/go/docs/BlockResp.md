# BlockResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockId** | **int64** | The num index of the block | 
**BlockSize** | **int32** | The block size | 
**Exchange** | **string** | The exchange address | 
**TxHash** | **string** | The txHash of the block | 
**Status** | **string** | The status of the block | 
**CreatedAt** | **int64** | The creation time of the block | 
**Transactions** | [**[]TransactionBlock**](TransactionBlock.md) | The txs list inside the block | 

## Methods

### NewBlockResp

`func NewBlockResp(blockId int64, blockSize int32, exchange string, txHash string, status string, createdAt int64, transactions []TransactionBlock, ) *BlockResp`

NewBlockResp instantiates a new BlockResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockRespWithDefaults

`func NewBlockRespWithDefaults() *BlockResp`

NewBlockRespWithDefaults instantiates a new BlockResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockId

`func (o *BlockResp) GetBlockId() int64`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *BlockResp) GetBlockIdOk() (*int64, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *BlockResp) SetBlockId(v int64)`

SetBlockId sets BlockId field to given value.


### GetBlockSize

`func (o *BlockResp) GetBlockSize() int32`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *BlockResp) GetBlockSizeOk() (*int32, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *BlockResp) SetBlockSize(v int32)`

SetBlockSize sets BlockSize field to given value.


### GetExchange

`func (o *BlockResp) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *BlockResp) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *BlockResp) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetTxHash

`func (o *BlockResp) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *BlockResp) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *BlockResp) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetStatus

`func (o *BlockResp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlockResp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlockResp) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *BlockResp) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BlockResp) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BlockResp) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetTransactions

`func (o *BlockResp) GetTransactions() []TransactionBlock`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *BlockResp) GetTransactionsOk() (*[]TransactionBlock, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *BlockResp) SetTransactions(v []TransactionBlock)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


