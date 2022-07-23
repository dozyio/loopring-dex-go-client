# AmmTransactionDataV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | AMM transaction hash | 
**TxType** | **string** | AMM transaction type, i.e., join, exit, etc | 
**TxStatus** | **string** | AMM transaction processing status, i.e., processing, processed, failed, etc | 
**AmmPoolAddress** | **string** | AMM pool address of query | 
**AmmLayerType** | **string** | AMM transaction layer, 1 or 2 | 
**PoolTokens** | [**[]AmmTransferDataV3**](AmmTransferDataV3.md) | The in pool tokens transfers records of the AMM transaction | 
**LpToken** | [**AmmTransferDataV3**](AmmTransferDataV3.md) |  | 
**CreatedAt** | **int64** | Transaction creation time | 
**UpdatedAt** | **int64** | Transaction update time | 
**BlockId** | **int64** |  | 
**IndexInBlock** | **int32** |  | 
**StorageInfos** | Pointer to [**[]StorageInfo**](StorageInfo.md) |  | [optional] 

## Methods

### NewAmmTransactionDataV3

`func NewAmmTransactionDataV3(hash string, txType string, txStatus string, ammPoolAddress string, ammLayerType string, poolTokens []AmmTransferDataV3, lpToken AmmTransferDataV3, createdAt int64, updatedAt int64, blockId int64, indexInBlock int32, ) *AmmTransactionDataV3`

NewAmmTransactionDataV3 instantiates a new AmmTransactionDataV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmTransactionDataV3WithDefaults

`func NewAmmTransactionDataV3WithDefaults() *AmmTransactionDataV3`

NewAmmTransactionDataV3WithDefaults instantiates a new AmmTransactionDataV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *AmmTransactionDataV3) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AmmTransactionDataV3) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AmmTransactionDataV3) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTxType

`func (o *AmmTransactionDataV3) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *AmmTransactionDataV3) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *AmmTransactionDataV3) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetTxStatus

`func (o *AmmTransactionDataV3) GetTxStatus() string`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *AmmTransactionDataV3) GetTxStatusOk() (*string, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *AmmTransactionDataV3) SetTxStatus(v string)`

SetTxStatus sets TxStatus field to given value.


### GetAmmPoolAddress

`func (o *AmmTransactionDataV3) GetAmmPoolAddress() string`

GetAmmPoolAddress returns the AmmPoolAddress field if non-nil, zero value otherwise.

### GetAmmPoolAddressOk

`func (o *AmmTransactionDataV3) GetAmmPoolAddressOk() (*string, bool)`

GetAmmPoolAddressOk returns a tuple with the AmmPoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmmPoolAddress

`func (o *AmmTransactionDataV3) SetAmmPoolAddress(v string)`

SetAmmPoolAddress sets AmmPoolAddress field to given value.


### GetAmmLayerType

`func (o *AmmTransactionDataV3) GetAmmLayerType() string`

GetAmmLayerType returns the AmmLayerType field if non-nil, zero value otherwise.

### GetAmmLayerTypeOk

`func (o *AmmTransactionDataV3) GetAmmLayerTypeOk() (*string, bool)`

GetAmmLayerTypeOk returns a tuple with the AmmLayerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmmLayerType

`func (o *AmmTransactionDataV3) SetAmmLayerType(v string)`

SetAmmLayerType sets AmmLayerType field to given value.


### GetPoolTokens

`func (o *AmmTransactionDataV3) GetPoolTokens() []AmmTransferDataV3`

GetPoolTokens returns the PoolTokens field if non-nil, zero value otherwise.

### GetPoolTokensOk

`func (o *AmmTransactionDataV3) GetPoolTokensOk() (*[]AmmTransferDataV3, bool)`

GetPoolTokensOk returns a tuple with the PoolTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolTokens

`func (o *AmmTransactionDataV3) SetPoolTokens(v []AmmTransferDataV3)`

SetPoolTokens sets PoolTokens field to given value.


### GetLpToken

`func (o *AmmTransactionDataV3) GetLpToken() AmmTransferDataV3`

GetLpToken returns the LpToken field if non-nil, zero value otherwise.

### GetLpTokenOk

`func (o *AmmTransactionDataV3) GetLpTokenOk() (*AmmTransferDataV3, bool)`

GetLpTokenOk returns a tuple with the LpToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpToken

`func (o *AmmTransactionDataV3) SetLpToken(v AmmTransferDataV3)`

SetLpToken sets LpToken field to given value.


### GetCreatedAt

`func (o *AmmTransactionDataV3) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AmmTransactionDataV3) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AmmTransactionDataV3) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AmmTransactionDataV3) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AmmTransactionDataV3) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AmmTransactionDataV3) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBlockId

`func (o *AmmTransactionDataV3) GetBlockId() int64`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *AmmTransactionDataV3) GetBlockIdOk() (*int64, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *AmmTransactionDataV3) SetBlockId(v int64)`

SetBlockId sets BlockId field to given value.


### GetIndexInBlock

`func (o *AmmTransactionDataV3) GetIndexInBlock() int32`

GetIndexInBlock returns the IndexInBlock field if non-nil, zero value otherwise.

### GetIndexInBlockOk

`func (o *AmmTransactionDataV3) GetIndexInBlockOk() (*int32, bool)`

GetIndexInBlockOk returns a tuple with the IndexInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexInBlock

`func (o *AmmTransactionDataV3) SetIndexInBlock(v int32)`

SetIndexInBlock sets IndexInBlock field to given value.


### GetStorageInfos

`func (o *AmmTransactionDataV3) GetStorageInfos() []StorageInfo`

GetStorageInfos returns the StorageInfos field if non-nil, zero value otherwise.

### GetStorageInfosOk

`func (o *AmmTransactionDataV3) GetStorageInfosOk() (*[]StorageInfo, bool)`

GetStorageInfosOk returns a tuple with the StorageInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInfos

`func (o *AmmTransactionDataV3) SetStorageInfos(v []StorageInfo)`

SetStorageInfos sets StorageInfos field to given value.

### HasStorageInfos

`func (o *AmmTransactionDataV3) HasStorageInfos() bool`

HasStorageInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


