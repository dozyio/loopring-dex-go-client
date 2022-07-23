# NftTokenAmountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **int32** | The Loopring&#39;s NFT token identifier. | 
**NftData** | Pointer to **string** | The Loopring&#39;s NFT token data identifier which is a hash string of NFT token address and NFT_ID | [optional] 
**Amount** | **string** | The amount of the NFT token | 

## Methods

### NewNftTokenAmountInfo

`func NewNftTokenAmountInfo(tokenId int32, amount string, ) *NftTokenAmountInfo`

NewNftTokenAmountInfo instantiates a new NftTokenAmountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftTokenAmountInfoWithDefaults

`func NewNftTokenAmountInfoWithDefaults() *NftTokenAmountInfo`

NewNftTokenAmountInfoWithDefaults instantiates a new NftTokenAmountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *NftTokenAmountInfo) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *NftTokenAmountInfo) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *NftTokenAmountInfo) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetNftData

`func (o *NftTokenAmountInfo) GetNftData() string`

GetNftData returns the NftData field if non-nil, zero value otherwise.

### GetNftDataOk

`func (o *NftTokenAmountInfo) GetNftDataOk() (*string, bool)`

GetNftDataOk returns a tuple with the NftData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftData

`func (o *NftTokenAmountInfo) SetNftData(v string)`

SetNftData sets NftData field to given value.

### HasNftData

`func (o *NftTokenAmountInfo) HasNftData() bool`

HasNftData returns a boolean if a field has been set.

### GetAmount

`func (o *NftTokenAmountInfo) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NftTokenAmountInfo) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NftTokenAmountInfo) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


