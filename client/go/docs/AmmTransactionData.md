# AmmTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Hash** | **string** |  | 
**TxHash** | **string** |  | 
**Owner** | **string** |  | 
**TxType** | **string** |  | 
**TxStatus** | **string** |  | 
**AmmLayerType** | **string** |  | 
**AmmPoolId** | **int64** |  | 
**AmmPoolAddress** | **string** |  | 
**LpTokenId** | **int64** |  | 
**LpTokenSymbol** | **string** |  | 
**LpTokenAmount** | **string** |  | 
**Transfers** | [**[]AmmTransferData**](AmmTransferData.md) |  | 
**BlockHeight** | **int64** |  | 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 
**StorageInfos** | Pointer to [**[]StorageInfo**](StorageInfo.md) |  | [optional] 

## Methods

### NewAmmTransactionData

`func NewAmmTransactionData(id int64, hash string, txHash string, owner string, txType string, txStatus string, ammLayerType string, ammPoolId int64, ammPoolAddress string, lpTokenId int64, lpTokenSymbol string, lpTokenAmount string, transfers []AmmTransferData, blockHeight int64, createdAt int64, updatedAt int64, ) *AmmTransactionData`

NewAmmTransactionData instantiates a new AmmTransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmTransactionDataWithDefaults

`func NewAmmTransactionDataWithDefaults() *AmmTransactionData`

NewAmmTransactionDataWithDefaults instantiates a new AmmTransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AmmTransactionData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AmmTransactionData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AmmTransactionData) SetId(v int64)`

SetId sets Id field to given value.


### GetHash

`func (o *AmmTransactionData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AmmTransactionData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AmmTransactionData) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTxHash

`func (o *AmmTransactionData) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *AmmTransactionData) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *AmmTransactionData) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetOwner

`func (o *AmmTransactionData) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AmmTransactionData) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AmmTransactionData) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetTxType

`func (o *AmmTransactionData) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *AmmTransactionData) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *AmmTransactionData) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetTxStatus

`func (o *AmmTransactionData) GetTxStatus() string`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *AmmTransactionData) GetTxStatusOk() (*string, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *AmmTransactionData) SetTxStatus(v string)`

SetTxStatus sets TxStatus field to given value.


### GetAmmLayerType

`func (o *AmmTransactionData) GetAmmLayerType() string`

GetAmmLayerType returns the AmmLayerType field if non-nil, zero value otherwise.

### GetAmmLayerTypeOk

`func (o *AmmTransactionData) GetAmmLayerTypeOk() (*string, bool)`

GetAmmLayerTypeOk returns a tuple with the AmmLayerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmmLayerType

`func (o *AmmTransactionData) SetAmmLayerType(v string)`

SetAmmLayerType sets AmmLayerType field to given value.


### GetAmmPoolId

`func (o *AmmTransactionData) GetAmmPoolId() int64`

GetAmmPoolId returns the AmmPoolId field if non-nil, zero value otherwise.

### GetAmmPoolIdOk

`func (o *AmmTransactionData) GetAmmPoolIdOk() (*int64, bool)`

GetAmmPoolIdOk returns a tuple with the AmmPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmmPoolId

`func (o *AmmTransactionData) SetAmmPoolId(v int64)`

SetAmmPoolId sets AmmPoolId field to given value.


### GetAmmPoolAddress

`func (o *AmmTransactionData) GetAmmPoolAddress() string`

GetAmmPoolAddress returns the AmmPoolAddress field if non-nil, zero value otherwise.

### GetAmmPoolAddressOk

`func (o *AmmTransactionData) GetAmmPoolAddressOk() (*string, bool)`

GetAmmPoolAddressOk returns a tuple with the AmmPoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmmPoolAddress

`func (o *AmmTransactionData) SetAmmPoolAddress(v string)`

SetAmmPoolAddress sets AmmPoolAddress field to given value.


### GetLpTokenId

`func (o *AmmTransactionData) GetLpTokenId() int64`

GetLpTokenId returns the LpTokenId field if non-nil, zero value otherwise.

### GetLpTokenIdOk

`func (o *AmmTransactionData) GetLpTokenIdOk() (*int64, bool)`

GetLpTokenIdOk returns a tuple with the LpTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpTokenId

`func (o *AmmTransactionData) SetLpTokenId(v int64)`

SetLpTokenId sets LpTokenId field to given value.


### GetLpTokenSymbol

`func (o *AmmTransactionData) GetLpTokenSymbol() string`

GetLpTokenSymbol returns the LpTokenSymbol field if non-nil, zero value otherwise.

### GetLpTokenSymbolOk

`func (o *AmmTransactionData) GetLpTokenSymbolOk() (*string, bool)`

GetLpTokenSymbolOk returns a tuple with the LpTokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpTokenSymbol

`func (o *AmmTransactionData) SetLpTokenSymbol(v string)`

SetLpTokenSymbol sets LpTokenSymbol field to given value.


### GetLpTokenAmount

`func (o *AmmTransactionData) GetLpTokenAmount() string`

GetLpTokenAmount returns the LpTokenAmount field if non-nil, zero value otherwise.

### GetLpTokenAmountOk

`func (o *AmmTransactionData) GetLpTokenAmountOk() (*string, bool)`

GetLpTokenAmountOk returns a tuple with the LpTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpTokenAmount

`func (o *AmmTransactionData) SetLpTokenAmount(v string)`

SetLpTokenAmount sets LpTokenAmount field to given value.


### GetTransfers

`func (o *AmmTransactionData) GetTransfers() []AmmTransferData`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *AmmTransactionData) GetTransfersOk() (*[]AmmTransferData, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *AmmTransactionData) SetTransfers(v []AmmTransferData)`

SetTransfers sets Transfers field to given value.


### GetBlockHeight

`func (o *AmmTransactionData) GetBlockHeight() int64`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *AmmTransactionData) GetBlockHeightOk() (*int64, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *AmmTransactionData) SetBlockHeight(v int64)`

SetBlockHeight sets BlockHeight field to given value.


### GetCreatedAt

`func (o *AmmTransactionData) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AmmTransactionData) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AmmTransactionData) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AmmTransactionData) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AmmTransactionData) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AmmTransactionData) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStorageInfos

`func (o *AmmTransactionData) GetStorageInfos() []StorageInfo`

GetStorageInfos returns the StorageInfos field if non-nil, zero value otherwise.

### GetStorageInfosOk

`func (o *AmmTransactionData) GetStorageInfosOk() (*[]StorageInfo, bool)`

GetStorageInfosOk returns a tuple with the StorageInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInfos

`func (o *AmmTransactionData) SetStorageInfos(v []StorageInfo)`

SetStorageInfos sets StorageInfos field to given value.

### HasStorageInfos

`func (o *AmmTransactionData) HasStorageInfos() bool`

HasStorageInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


