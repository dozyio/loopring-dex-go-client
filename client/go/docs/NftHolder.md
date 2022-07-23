# NftHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | Sender accountId of the Tx | 
**TokenId** | **int32** | The Loopring&#39;s NFT token identifier. | 
**Amount** | **string** | The amount of the NFT token | 

## Methods

### NewNftHolder

`func NewNftHolder(accountId int64, tokenId int32, amount string, ) *NftHolder`

NewNftHolder instantiates a new NftHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftHolderWithDefaults

`func NewNftHolderWithDefaults() *NftHolder`

NewNftHolderWithDefaults instantiates a new NftHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NftHolder) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NftHolder) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NftHolder) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetTokenId

`func (o *NftHolder) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *NftHolder) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *NftHolder) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *NftHolder) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NftHolder) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NftHolder) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


