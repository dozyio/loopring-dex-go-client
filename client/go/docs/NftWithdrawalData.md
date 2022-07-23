# NftWithdrawalData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | field.NftWithdrawalData.id | 
**RequestId** | **int64** | field.NftWithdrawalData.requestId | 
**Hash** | **string** | field.NftWithdrawalData.hash | 
**TxHash** | **string** | field.NftWithdrawalData.txHash | 
**AccountId** | **int64** | field.NftWithdrawalData.accountId | 
**Owner** | **string** | field.NftWithdrawalData.owner | 
**Status** | **string** | field.NftWithdrawalData.status | 
**NftData** | Pointer to **string** | field.NftWithdrawalData.nftData | [optional] 
**Amount** | Pointer to **string** | field.NftWithdrawalData.amount | [optional] 
**FeeTokenSymbol** | Pointer to **string** | field.NftWithdrawalData.feeTokenSymbol | [optional] 
**FeeAmount** | **string** | field.NftWithdrawalData.feeAmount | 
**CreatedAt** | **int64** | field.NftWithdrawalData.createdAt | 
**UpdatedAt** | **int64** | field.NftWithdrawalData.updatedAt | 
**Memo** | Pointer to **string** | field.NftWithdrawalData.memo | [optional] 
**Recipient** | **string** | field.NftWithdrawalData.recipient | 
**DistributeHash** | **string** | field.NftWithdrawalData.distributeHash | 
**FastWithdrawStatus** | **string** | field.NftWithdrawalData.fastWithdrawStatus | 
**IsFast** | **bool** | field.NftWithdrawalData.isFast | 
**BlockIdInfo** | Pointer to [**BlockIdInfo**](BlockIdInfo.md) |  | [optional] 
**StorageInfo** | Pointer to [**StorageInfo**](StorageInfo.md) |  | [optional] 

## Methods

### NewNftWithdrawalData

`func NewNftWithdrawalData(id int64, requestId int64, hash string, txHash string, accountId int64, owner string, status string, feeAmount string, createdAt int64, updatedAt int64, recipient string, distributeHash string, fastWithdrawStatus string, isFast bool, ) *NftWithdrawalData`

NewNftWithdrawalData instantiates a new NftWithdrawalData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftWithdrawalDataWithDefaults

`func NewNftWithdrawalDataWithDefaults() *NftWithdrawalData`

NewNftWithdrawalDataWithDefaults instantiates a new NftWithdrawalData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NftWithdrawalData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NftWithdrawalData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NftWithdrawalData) SetId(v int64)`

SetId sets Id field to given value.


### GetRequestId

`func (o *NftWithdrawalData) GetRequestId() int64`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *NftWithdrawalData) GetRequestIdOk() (*int64, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *NftWithdrawalData) SetRequestId(v int64)`

SetRequestId sets RequestId field to given value.


### GetHash

`func (o *NftWithdrawalData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *NftWithdrawalData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *NftWithdrawalData) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTxHash

`func (o *NftWithdrawalData) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *NftWithdrawalData) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *NftWithdrawalData) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetAccountId

`func (o *NftWithdrawalData) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NftWithdrawalData) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NftWithdrawalData) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetOwner

`func (o *NftWithdrawalData) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NftWithdrawalData) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NftWithdrawalData) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetStatus

`func (o *NftWithdrawalData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NftWithdrawalData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NftWithdrawalData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetNftData

`func (o *NftWithdrawalData) GetNftData() string`

GetNftData returns the NftData field if non-nil, zero value otherwise.

### GetNftDataOk

`func (o *NftWithdrawalData) GetNftDataOk() (*string, bool)`

GetNftDataOk returns a tuple with the NftData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftData

`func (o *NftWithdrawalData) SetNftData(v string)`

SetNftData sets NftData field to given value.

### HasNftData

`func (o *NftWithdrawalData) HasNftData() bool`

HasNftData returns a boolean if a field has been set.

### GetAmount

`func (o *NftWithdrawalData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NftWithdrawalData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NftWithdrawalData) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *NftWithdrawalData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFeeTokenSymbol

`func (o *NftWithdrawalData) GetFeeTokenSymbol() string`

GetFeeTokenSymbol returns the FeeTokenSymbol field if non-nil, zero value otherwise.

### GetFeeTokenSymbolOk

`func (o *NftWithdrawalData) GetFeeTokenSymbolOk() (*string, bool)`

GetFeeTokenSymbolOk returns a tuple with the FeeTokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenSymbol

`func (o *NftWithdrawalData) SetFeeTokenSymbol(v string)`

SetFeeTokenSymbol sets FeeTokenSymbol field to given value.

### HasFeeTokenSymbol

`func (o *NftWithdrawalData) HasFeeTokenSymbol() bool`

HasFeeTokenSymbol returns a boolean if a field has been set.

### GetFeeAmount

`func (o *NftWithdrawalData) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *NftWithdrawalData) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *NftWithdrawalData) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetCreatedAt

`func (o *NftWithdrawalData) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NftWithdrawalData) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NftWithdrawalData) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *NftWithdrawalData) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NftWithdrawalData) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NftWithdrawalData) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMemo

`func (o *NftWithdrawalData) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *NftWithdrawalData) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *NftWithdrawalData) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *NftWithdrawalData) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetRecipient

`func (o *NftWithdrawalData) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NftWithdrawalData) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NftWithdrawalData) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetDistributeHash

`func (o *NftWithdrawalData) GetDistributeHash() string`

GetDistributeHash returns the DistributeHash field if non-nil, zero value otherwise.

### GetDistributeHashOk

`func (o *NftWithdrawalData) GetDistributeHashOk() (*string, bool)`

GetDistributeHashOk returns a tuple with the DistributeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeHash

`func (o *NftWithdrawalData) SetDistributeHash(v string)`

SetDistributeHash sets DistributeHash field to given value.


### GetFastWithdrawStatus

`func (o *NftWithdrawalData) GetFastWithdrawStatus() string`

GetFastWithdrawStatus returns the FastWithdrawStatus field if non-nil, zero value otherwise.

### GetFastWithdrawStatusOk

`func (o *NftWithdrawalData) GetFastWithdrawStatusOk() (*string, bool)`

GetFastWithdrawStatusOk returns a tuple with the FastWithdrawStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastWithdrawStatus

`func (o *NftWithdrawalData) SetFastWithdrawStatus(v string)`

SetFastWithdrawStatus sets FastWithdrawStatus field to given value.


### GetIsFast

`func (o *NftWithdrawalData) GetIsFast() bool`

GetIsFast returns the IsFast field if non-nil, zero value otherwise.

### GetIsFastOk

`func (o *NftWithdrawalData) GetIsFastOk() (*bool, bool)`

GetIsFastOk returns a tuple with the IsFast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFast

`func (o *NftWithdrawalData) SetIsFast(v bool)`

SetIsFast sets IsFast field to given value.


### GetBlockIdInfo

`func (o *NftWithdrawalData) GetBlockIdInfo() BlockIdInfo`

GetBlockIdInfo returns the BlockIdInfo field if non-nil, zero value otherwise.

### GetBlockIdInfoOk

`func (o *NftWithdrawalData) GetBlockIdInfoOk() (*BlockIdInfo, bool)`

GetBlockIdInfoOk returns a tuple with the BlockIdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIdInfo

`func (o *NftWithdrawalData) SetBlockIdInfo(v BlockIdInfo)`

SetBlockIdInfo sets BlockIdInfo field to given value.

### HasBlockIdInfo

`func (o *NftWithdrawalData) HasBlockIdInfo() bool`

HasBlockIdInfo returns a boolean if a field has been set.

### GetStorageInfo

`func (o *NftWithdrawalData) GetStorageInfo() StorageInfo`

GetStorageInfo returns the StorageInfo field if non-nil, zero value otherwise.

### GetStorageInfoOk

`func (o *NftWithdrawalData) GetStorageInfoOk() (*StorageInfo, bool)`

GetStorageInfoOk returns a tuple with the StorageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInfo

`func (o *NftWithdrawalData) SetStorageInfo(v StorageInfo)`

SetStorageInfo sets StorageInfo field to given value.

### HasStorageInfo

`func (o *NftWithdrawalData) HasStorageInfo() bool`

HasStorageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


