# NftDepositData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | field.NftDepositData.id | 
**RequestId** | **int64** | field.NftDepositData.requestId | 
**Hash** | **string** | field.NftDepositData.hash | 
**TxHash** | **string** | field.NftDepositData.txHash | 
**AccountId** | **int64** | field.NftDepositData.accountId | 
**Owner** | **string** | field.NftDepositData.owner | 
**Status** | **string** | field.NftDepositData.status | 
**NftData** | **string** | field.NftDepositData.nftData | 
**Amount** | **string** | field.NftDepositData.amount | 
**FeeTokenSymbol** | **string** | field.NftDepositData.feeTokenSymbol | 
**FeeAmount** | **string** | field.NftDepositData.feeAmount | 
**CreatedAt** | **int64** | field.NftDepositData.createdAt | 
**UpdatedAt** | **int64** | field.NftDepositData.updatedAt | 
**Memo** | Pointer to **string** | field.NftDepositData.memo | [optional] 
**DepositFrom** | **string** | field.NftDepositData.depositFrom | 
**DepositFromAccountId** | **int64** | field.NftDepositData.depositFromAccountId | 
**BlockIdInfo** | Pointer to [**BlockIdInfo**](BlockIdInfo.md) |  | [optional] 
**StorageInfo** | Pointer to [**StorageInfo**](StorageInfo.md) |  | [optional] 

## Methods

### NewNftDepositData

`func NewNftDepositData(id int64, requestId int64, hash string, txHash string, accountId int64, owner string, status string, nftData string, amount string, feeTokenSymbol string, feeAmount string, createdAt int64, updatedAt int64, depositFrom string, depositFromAccountId int64, ) *NftDepositData`

NewNftDepositData instantiates a new NftDepositData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftDepositDataWithDefaults

`func NewNftDepositDataWithDefaults() *NftDepositData`

NewNftDepositDataWithDefaults instantiates a new NftDepositData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NftDepositData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NftDepositData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NftDepositData) SetId(v int64)`

SetId sets Id field to given value.


### GetRequestId

`func (o *NftDepositData) GetRequestId() int64`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *NftDepositData) GetRequestIdOk() (*int64, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *NftDepositData) SetRequestId(v int64)`

SetRequestId sets RequestId field to given value.


### GetHash

`func (o *NftDepositData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *NftDepositData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *NftDepositData) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTxHash

`func (o *NftDepositData) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *NftDepositData) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *NftDepositData) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetAccountId

`func (o *NftDepositData) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NftDepositData) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NftDepositData) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetOwner

`func (o *NftDepositData) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NftDepositData) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NftDepositData) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetStatus

`func (o *NftDepositData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NftDepositData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NftDepositData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNftData

`func (o *NftDepositData) GetNftData() string`

GetNftData returns the NftData field if non-nil, zero value otherwise.

### GetNftDataOk

`func (o *NftDepositData) GetNftDataOk() (*string, bool)`

GetNftDataOk returns a tuple with the NftData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftData

`func (o *NftDepositData) SetNftData(v string)`

SetNftData sets NftData field to given value.


### GetAmount

`func (o *NftDepositData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NftDepositData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NftDepositData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFeeTokenSymbol

`func (o *NftDepositData) GetFeeTokenSymbol() string`

GetFeeTokenSymbol returns the FeeTokenSymbol field if non-nil, zero value otherwise.

### GetFeeTokenSymbolOk

`func (o *NftDepositData) GetFeeTokenSymbolOk() (*string, bool)`

GetFeeTokenSymbolOk returns a tuple with the FeeTokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenSymbol

`func (o *NftDepositData) SetFeeTokenSymbol(v string)`

SetFeeTokenSymbol sets FeeTokenSymbol field to given value.


### GetFeeAmount

`func (o *NftDepositData) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *NftDepositData) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *NftDepositData) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetCreatedAt

`func (o *NftDepositData) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NftDepositData) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NftDepositData) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NftDepositData) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NftDepositData) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NftDepositData) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMemo

`func (o *NftDepositData) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *NftDepositData) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *NftDepositData) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *NftDepositData) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetDepositFrom

`func (o *NftDepositData) GetDepositFrom() string`

GetDepositFrom returns the DepositFrom field if non-nil, zero value otherwise.

### GetDepositFromOk

`func (o *NftDepositData) GetDepositFromOk() (*string, bool)`

GetDepositFromOk returns a tuple with the DepositFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositFrom

`func (o *NftDepositData) SetDepositFrom(v string)`

SetDepositFrom sets DepositFrom field to given value.


### GetDepositFromAccountId

`func (o *NftDepositData) GetDepositFromAccountId() int64`

GetDepositFromAccountId returns the DepositFromAccountId field if non-nil, zero value otherwise.

### GetDepositFromAccountIdOk

`func (o *NftDepositData) GetDepositFromAccountIdOk() (*int64, bool)`

GetDepositFromAccountIdOk returns a tuple with the DepositFromAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositFromAccountId

`func (o *NftDepositData) SetDepositFromAccountId(v int64)`

SetDepositFromAccountId sets DepositFromAccountId field to given value.


### GetBlockIdInfo

`func (o *NftDepositData) GetBlockIdInfo() BlockIdInfo`

GetBlockIdInfo returns the BlockIdInfo field if non-nil, zero value otherwise.

### GetBlockIdInfoOk

`func (o *NftDepositData) GetBlockIdInfoOk() (*BlockIdInfo, bool)`

GetBlockIdInfoOk returns a tuple with the BlockIdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIdInfo

`func (o *NftDepositData) SetBlockIdInfo(v BlockIdInfo)`

SetBlockIdInfo sets BlockIdInfo field to given value.

### HasBlockIdInfo

`func (o *NftDepositData) HasBlockIdInfo() bool`

HasBlockIdInfo returns a boolean if a field has been set.

### GetStorageInfo

`func (o *NftDepositData) GetStorageInfo() StorageInfo`

GetStorageInfo returns the StorageInfo field if non-nil, zero value otherwise.

### GetStorageInfoOk

`func (o *NftDepositData) GetStorageInfoOk() (*StorageInfo, bool)`

GetStorageInfoOk returns a tuple with the StorageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInfo

`func (o *NftDepositData) SetStorageInfo(v StorageInfo)`

SetStorageInfo sets StorageInfo field to given value.

### HasStorageInfo

`func (o *NftDepositData) HasStorageInfo() bool`

HasStorageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


