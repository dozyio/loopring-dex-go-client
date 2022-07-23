# NftMintData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | field.NftMintData.id | 
**RequestId** | **int64** | field.NftMintData.requestId | 
**Hash** | **string** | field.NftMintData.hash | 
**TxHash** | **string** | field.NftMintData.txHash | 
**AccountId** | **int64** | field.NftMintData.accountId | 
**Owner** | **string** | field.NftMintData.owner | 
**Status** | **string** | field.NftMintData.status | 
**NftData** | **string** | field.NftMintData.nftData | 
**Amount** | **string** | field.NftMintData.amount | 
**FeeTokenSymbol** | **string** | field.NftMintData.feeTokenSymbol | 
**FeeAmount** | **string** | field.NftMintData.feeAmount | 
**CreatedAt** | **int64** | field.NftMintData.createdAt | 
**UpdatedAt** | **int64** | field.NftMintData.updatedAt | 
**Memo** | **string** | field.NftMintData.memo | 
**MinterId** | **int64** | field.NftMintData.minterId | 
**MinterAddress** | **string** | field.NftMintData.minterAddress | 
**BlockIdInfo** | Pointer to [**BlockIdInfo**](BlockIdInfo.md) |  | [optional] 
**StorageInfo** | Pointer to [**StorageInfo**](StorageInfo.md) |  | [optional] 

## Methods

### NewNftMintData

`func NewNftMintData(id int64, requestId int64, hash string, txHash string, accountId int64, owner string, status string, nftData string, amount string, feeTokenSymbol string, feeAmount string, createdAt int64, updatedAt int64, memo string, minterId int64, minterAddress string, ) *NftMintData`

NewNftMintData instantiates a new NftMintData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftMintDataWithDefaults

`func NewNftMintDataWithDefaults() *NftMintData`

NewNftMintDataWithDefaults instantiates a new NftMintData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NftMintData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NftMintData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NftMintData) SetId(v int64)`

SetId sets Id field to given value.


### GetRequestId

`func (o *NftMintData) GetRequestId() int64`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *NftMintData) GetRequestIdOk() (*int64, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *NftMintData) SetRequestId(v int64)`

SetRequestId sets RequestId field to given value.


### GetHash

`func (o *NftMintData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *NftMintData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *NftMintData) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTxHash

`func (o *NftMintData) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *NftMintData) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *NftMintData) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetAccountId

`func (o *NftMintData) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NftMintData) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NftMintData) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetOwner

`func (o *NftMintData) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NftMintData) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NftMintData) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetStatus

`func (o *NftMintData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NftMintData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NftMintData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNftData

`func (o *NftMintData) GetNftData() string`

GetNftData returns the NftData field if non-nil, zero value otherwise.

### GetNftDataOk

`func (o *NftMintData) GetNftDataOk() (*string, bool)`

GetNftDataOk returns a tuple with the NftData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftData

`func (o *NftMintData) SetNftData(v string)`

SetNftData sets NftData field to given value.


### GetAmount

`func (o *NftMintData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NftMintData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NftMintData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFeeTokenSymbol

`func (o *NftMintData) GetFeeTokenSymbol() string`

GetFeeTokenSymbol returns the FeeTokenSymbol field if non-nil, zero value otherwise.

### GetFeeTokenSymbolOk

`func (o *NftMintData) GetFeeTokenSymbolOk() (*string, bool)`

GetFeeTokenSymbolOk returns a tuple with the FeeTokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenSymbol

`func (o *NftMintData) SetFeeTokenSymbol(v string)`

SetFeeTokenSymbol sets FeeTokenSymbol field to given value.


### GetFeeAmount

`func (o *NftMintData) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *NftMintData) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *NftMintData) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetCreatedAt

`func (o *NftMintData) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NftMintData) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NftMintData) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NftMintData) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NftMintData) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NftMintData) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMemo

`func (o *NftMintData) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *NftMintData) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *NftMintData) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetMinterId

`func (o *NftMintData) GetMinterId() int64`

GetMinterId returns the MinterId field if non-nil, zero value otherwise.

### GetMinterIdOk

`func (o *NftMintData) GetMinterIdOk() (*int64, bool)`

GetMinterIdOk returns a tuple with the MinterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinterId

`func (o *NftMintData) SetMinterId(v int64)`

SetMinterId sets MinterId field to given value.


### GetMinterAddress

`func (o *NftMintData) GetMinterAddress() string`

GetMinterAddress returns the MinterAddress field if non-nil, zero value otherwise.

### GetMinterAddressOk

`func (o *NftMintData) GetMinterAddressOk() (*string, bool)`

GetMinterAddressOk returns a tuple with the MinterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinterAddress

`func (o *NftMintData) SetMinterAddress(v string)`

SetMinterAddress sets MinterAddress field to given value.


### GetBlockIdInfo

`func (o *NftMintData) GetBlockIdInfo() BlockIdInfo`

GetBlockIdInfo returns the BlockIdInfo field if non-nil, zero value otherwise.

### GetBlockIdInfoOk

`func (o *NftMintData) GetBlockIdInfoOk() (*BlockIdInfo, bool)`

GetBlockIdInfoOk returns a tuple with the BlockIdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIdInfo

`func (o *NftMintData) SetBlockIdInfo(v BlockIdInfo)`

SetBlockIdInfo sets BlockIdInfo field to given value.

### HasBlockIdInfo

`func (o *NftMintData) HasBlockIdInfo() bool`

HasBlockIdInfo returns a boolean if a field has been set.

### GetStorageInfo

`func (o *NftMintData) GetStorageInfo() StorageInfo`

GetStorageInfo returns the StorageInfo field if non-nil, zero value otherwise.

### GetStorageInfoOk

`func (o *NftMintData) GetStorageInfoOk() (*StorageInfo, bool)`

GetStorageInfoOk returns a tuple with the StorageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInfo

`func (o *NftMintData) SetStorageInfo(v StorageInfo)`

SetStorageInfo sets StorageInfo field to given value.

### HasStorageInfo

`func (o *NftMintData) HasStorageInfo() bool`

HasStorageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


