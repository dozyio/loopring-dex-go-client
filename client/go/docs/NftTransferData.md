# NftTransferData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | field.NftTransferData.id | 
**RequestId** | **int64** | field.NftTransferData.requestId | 
**Hash** | **string** | field.NftTransferData.hash | 
**TxHash** | **string** | field.NftTransferData.txHash | 
**AccountId** | **int64** | field.NftTransferData.accountId | 
**Owner** | **string** | field.NftTransferData.owner | 
**Status** | **string** | field.NftTransferData.status | 
**NftData** | **string** | field.NftTransferData.nftData | 
**Amount** | **string** | field.NftTransferData.amount | 
**FeeTokenSymbol** | **string** | field.NftTransferData.feeTokenSymbol | 
**FeeAmount** | **string** | field.NftTransferData.feeAmount | 
**CreatedAt** | **int64** | field.NftTransferData.createdAt | 
**UpdatedAt** | **int64** | field.NftTransferData.updatedAt | 
**Memo** | **string** | field.NftTransferData.memo | 
**PayeeId** | **int64** | field.NftTransferData.payeeId | 
**PayeeAddress** | **string** | field.NftTransferData.payeeAddress | 
**BlockIdInfo** | Pointer to [**BlockIdInfo**](BlockIdInfo.md) |  | [optional] 
**StorageInfo** | Pointer to [**StorageInfo**](StorageInfo.md) |  | [optional] 

## Methods

### NewNftTransferData

`func NewNftTransferData(id string, requestId int64, hash string, txHash string, accountId int64, owner string, status string, nftData string, amount string, feeTokenSymbol string, feeAmount string, createdAt int64, updatedAt int64, memo string, payeeId int64, payeeAddress string, ) *NftTransferData`

NewNftTransferData instantiates a new NftTransferData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftTransferDataWithDefaults

`func NewNftTransferDataWithDefaults() *NftTransferData`

NewNftTransferDataWithDefaults instantiates a new NftTransferData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NftTransferData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NftTransferData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NftTransferData) SetId(v string)`

SetId sets Id field to given value.


### GetRequestId

`func (o *NftTransferData) GetRequestId() int64`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *NftTransferData) GetRequestIdOk() (*int64, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *NftTransferData) SetRequestId(v int64)`

SetRequestId sets RequestId field to given value.


### GetHash

`func (o *NftTransferData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *NftTransferData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *NftTransferData) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTxHash

`func (o *NftTransferData) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *NftTransferData) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *NftTransferData) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetAccountId

`func (o *NftTransferData) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NftTransferData) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NftTransferData) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetOwner

`func (o *NftTransferData) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NftTransferData) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NftTransferData) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetStatus

`func (o *NftTransferData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NftTransferData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NftTransferData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNftData

`func (o *NftTransferData) GetNftData() string`

GetNftData returns the NftData field if non-nil, zero value otherwise.

### GetNftDataOk

`func (o *NftTransferData) GetNftDataOk() (*string, bool)`

GetNftDataOk returns a tuple with the NftData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftData

`func (o *NftTransferData) SetNftData(v string)`

SetNftData sets NftData field to given value.


### GetAmount

`func (o *NftTransferData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NftTransferData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NftTransferData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFeeTokenSymbol

`func (o *NftTransferData) GetFeeTokenSymbol() string`

GetFeeTokenSymbol returns the FeeTokenSymbol field if non-nil, zero value otherwise.

### GetFeeTokenSymbolOk

`func (o *NftTransferData) GetFeeTokenSymbolOk() (*string, bool)`

GetFeeTokenSymbolOk returns a tuple with the FeeTokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenSymbol

`func (o *NftTransferData) SetFeeTokenSymbol(v string)`

SetFeeTokenSymbol sets FeeTokenSymbol field to given value.


### GetFeeAmount

`func (o *NftTransferData) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *NftTransferData) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *NftTransferData) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetCreatedAt

`func (o *NftTransferData) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NftTransferData) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NftTransferData) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NftTransferData) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NftTransferData) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NftTransferData) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMemo

`func (o *NftTransferData) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *NftTransferData) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *NftTransferData) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetPayeeId

`func (o *NftTransferData) GetPayeeId() int64`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *NftTransferData) GetPayeeIdOk() (*int64, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *NftTransferData) SetPayeeId(v int64)`

SetPayeeId sets PayeeId field to given value.


### GetPayeeAddress

`func (o *NftTransferData) GetPayeeAddress() string`

GetPayeeAddress returns the PayeeAddress field if non-nil, zero value otherwise.

### GetPayeeAddressOk

`func (o *NftTransferData) GetPayeeAddressOk() (*string, bool)`

GetPayeeAddressOk returns a tuple with the PayeeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeAddress

`func (o *NftTransferData) SetPayeeAddress(v string)`

SetPayeeAddress sets PayeeAddress field to given value.


### GetBlockIdInfo

`func (o *NftTransferData) GetBlockIdInfo() BlockIdInfo`

GetBlockIdInfo returns the BlockIdInfo field if non-nil, zero value otherwise.

### GetBlockIdInfoOk

`func (o *NftTransferData) GetBlockIdInfoOk() (*BlockIdInfo, bool)`

GetBlockIdInfoOk returns a tuple with the BlockIdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIdInfo

`func (o *NftTransferData) SetBlockIdInfo(v BlockIdInfo)`

SetBlockIdInfo sets BlockIdInfo field to given value.

### HasBlockIdInfo

`func (o *NftTransferData) HasBlockIdInfo() bool`

HasBlockIdInfo returns a boolean if a field has been set.

### GetStorageInfo

`func (o *NftTransferData) GetStorageInfo() StorageInfo`

GetStorageInfo returns the StorageInfo field if non-nil, zero value otherwise.

### GetStorageInfoOk

`func (o *NftTransferData) GetStorageInfoOk() (*StorageInfo, bool)`

GetStorageInfoOk returns a tuple with the StorageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInfo

`func (o *NftTransferData) SetStorageInfo(v StorageInfo)`

SetStorageInfo sets StorageInfo field to given value.

### HasStorageInfo

`func (o *NftTransferData) HasStorageInfo() bool`

HasStorageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


