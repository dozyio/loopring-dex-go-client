# TokenAmountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **int32** | The Loopring&#39;s NFT token identifier. | 
**Amount** | **string** | The amount of the NFT token | 

## Methods

### NewTokenAmountInfo

`func NewTokenAmountInfo(tokenId int32, amount string, ) *TokenAmountInfo`

NewTokenAmountInfo instantiates a new TokenAmountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenAmountInfoWithDefaults

`func NewTokenAmountInfoWithDefaults() *TokenAmountInfo`

NewTokenAmountInfoWithDefaults instantiates a new TokenAmountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TokenAmountInfo) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenAmountInfo) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenAmountInfo) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *TokenAmountInfo) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenAmountInfo) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenAmountInfo) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


