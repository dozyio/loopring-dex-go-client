# SubmitMintNftResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Mint request hash. | 
**NftTokenId** | **int32** | field.SubmitMintNftResponseItem.nftTokenId | 
**NftData** | **string** | Mint tokens nftData. | 
**Status** | **string** | Mint status. | 
**IsIdempotent** | **bool** | Idempotent of submit order response. True if the same request is sent more than once. | 
**AccountId** | **int64** | field.SubmitOffChainResponseItem.accountId | 
**StorageId** | **int64** | field.SubmitOffChainResponseItem.storageId | 

## Methods

### NewSubmitMintNftResponseItem

`func NewSubmitMintNftResponseItem(hash string, nftTokenId int32, nftData string, status string, isIdempotent bool, accountId int64, storageId int64, ) *SubmitMintNftResponseItem`

NewSubmitMintNftResponseItem instantiates a new SubmitMintNftResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitMintNftResponseItemWithDefaults

`func NewSubmitMintNftResponseItemWithDefaults() *SubmitMintNftResponseItem`

NewSubmitMintNftResponseItemWithDefaults instantiates a new SubmitMintNftResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *SubmitMintNftResponseItem) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SubmitMintNftResponseItem) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SubmitMintNftResponseItem) SetHash(v string)`

SetHash sets Hash field to given value.


### GetNftTokenId

`func (o *SubmitMintNftResponseItem) GetNftTokenId() int32`

GetNftTokenId returns the NftTokenId field if non-nil, zero value otherwise.

### GetNftTokenIdOk

`func (o *SubmitMintNftResponseItem) GetNftTokenIdOk() (*int32, bool)`

GetNftTokenIdOk returns a tuple with the NftTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftTokenId

`func (o *SubmitMintNftResponseItem) SetNftTokenId(v int32)`

SetNftTokenId sets NftTokenId field to given value.


### GetNftData

`func (o *SubmitMintNftResponseItem) GetNftData() string`

GetNftData returns the NftData field if non-nil, zero value otherwise.

### GetNftDataOk

`func (o *SubmitMintNftResponseItem) GetNftDataOk() (*string, bool)`

GetNftDataOk returns a tuple with the NftData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftData

`func (o *SubmitMintNftResponseItem) SetNftData(v string)`

SetNftData sets NftData field to given value.


### GetStatus

`func (o *SubmitMintNftResponseItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitMintNftResponseItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitMintNftResponseItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIsIdempotent

`func (o *SubmitMintNftResponseItem) GetIsIdempotent() bool`

GetIsIdempotent returns the IsIdempotent field if non-nil, zero value otherwise.

### GetIsIdempotentOk

`func (o *SubmitMintNftResponseItem) GetIsIdempotentOk() (*bool, bool)`

GetIsIdempotentOk returns a tuple with the IsIdempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIdempotent

`func (o *SubmitMintNftResponseItem) SetIsIdempotent(v bool)`

SetIsIdempotent sets IsIdempotent field to given value.


### GetAccountId

`func (o *SubmitMintNftResponseItem) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SubmitMintNftResponseItem) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SubmitMintNftResponseItem) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetStorageId

`func (o *SubmitMintNftResponseItem) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *SubmitMintNftResponseItem) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *SubmitMintNftResponseItem) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


