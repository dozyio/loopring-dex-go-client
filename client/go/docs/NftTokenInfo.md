# NftTokenInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NftData** | **string** | NFT token nftData, which is a hash of token address, NFT_ID, minter, etc. | 
**Minter** | **string** | The minter of the NFT token | 
**NftType** | **string** | field.NftTokenInfo.nftType | 
**TokenAddress** | **string** | field.NftTokenInfo.tokenAddress | 
**NftId** | **string** | field.NftTokenInfo.nftId | 
**CreatorFeeBips** | **int32** | field.NftTokenInfo.creatorFeeBips | 
**RoyaltyPercentage** | **int32** | field.NftTokenInfo.royaltyPercentage | 
**OriginalRoyaltyPercentage** | **int32** | field.NftTokenInfo.originalRoyaltyPercentage | 
**Status** | **bool** | field.NftTokenInfo.status | 
**NftFactory** | Pointer to **string** | field.NftTokenInfo.nftFactory | [optional] 
**NftOwner** | Pointer to **string** | field.NftTokenInfo.nftOwner | [optional] 
**NftBaseUri** | Pointer to **string** | field.NftTokenInfo.nftBaseUri | [optional] 
**RoyaltyAddress** | Pointer to **string** | field.NftTokenInfo.royaltyAddress | [optional] 
**OriginalMinter** | Pointer to **string** | field.NftTokenInfo.originalMinter | [optional] 
**CreatedAt** | Pointer to **int64** | field.NftTokenInfo.createdAt | [optional] 

## Methods

### NewNftTokenInfo

`func NewNftTokenInfo(nftData string, minter string, nftType string, tokenAddress string, nftId string, creatorFeeBips int32, royaltyPercentage int32, originalRoyaltyPercentage int32, status bool, ) *NftTokenInfo`

NewNftTokenInfo instantiates a new NftTokenInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftTokenInfoWithDefaults

`func NewNftTokenInfoWithDefaults() *NftTokenInfo`

NewNftTokenInfoWithDefaults instantiates a new NftTokenInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNftData

`func (o *NftTokenInfo) GetNftData() string`

GetNftData returns the NftData field if non-nil, zero value otherwise.

### GetNftDataOk

`func (o *NftTokenInfo) GetNftDataOk() (*string, bool)`

GetNftDataOk returns a tuple with the NftData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftData

`func (o *NftTokenInfo) SetNftData(v string)`

SetNftData sets NftData field to given value.


### GetMinter

`func (o *NftTokenInfo) GetMinter() string`

GetMinter returns the Minter field if non-nil, zero value otherwise.

### GetMinterOk

`func (o *NftTokenInfo) GetMinterOk() (*string, bool)`

GetMinterOk returns a tuple with the Minter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinter

`func (o *NftTokenInfo) SetMinter(v string)`

SetMinter sets Minter field to given value.


### GetNftType

`func (o *NftTokenInfo) GetNftType() string`

GetNftType returns the NftType field if non-nil, zero value otherwise.

### GetNftTypeOk

`func (o *NftTokenInfo) GetNftTypeOk() (*string, bool)`

GetNftTypeOk returns a tuple with the NftType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftType

`func (o *NftTokenInfo) SetNftType(v string)`

SetNftType sets NftType field to given value.


### GetTokenAddress

`func (o *NftTokenInfo) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *NftTokenInfo) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *NftTokenInfo) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetNftId

`func (o *NftTokenInfo) GetNftId() string`

GetNftId returns the NftId field if non-nil, zero value otherwise.

### GetNftIdOk

`func (o *NftTokenInfo) GetNftIdOk() (*string, bool)`

GetNftIdOk returns a tuple with the NftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftId

`func (o *NftTokenInfo) SetNftId(v string)`

SetNftId sets NftId field to given value.


### GetCreatorFeeBips

`func (o *NftTokenInfo) GetCreatorFeeBips() int32`

GetCreatorFeeBips returns the CreatorFeeBips field if non-nil, zero value otherwise.

### GetCreatorFeeBipsOk

`func (o *NftTokenInfo) GetCreatorFeeBipsOk() (*int32, bool)`

GetCreatorFeeBipsOk returns a tuple with the CreatorFeeBips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorFeeBips

`func (o *NftTokenInfo) SetCreatorFeeBips(v int32)`

SetCreatorFeeBips sets CreatorFeeBips field to given value.


### GetRoyaltyPercentage

`func (o *NftTokenInfo) GetRoyaltyPercentage() int32`

GetRoyaltyPercentage returns the RoyaltyPercentage field if non-nil, zero value otherwise.

### GetRoyaltyPercentageOk

`func (o *NftTokenInfo) GetRoyaltyPercentageOk() (*int32, bool)`

GetRoyaltyPercentageOk returns a tuple with the RoyaltyPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyaltyPercentage

`func (o *NftTokenInfo) SetRoyaltyPercentage(v int32)`

SetRoyaltyPercentage sets RoyaltyPercentage field to given value.


### GetOriginalRoyaltyPercentage

`func (o *NftTokenInfo) GetOriginalRoyaltyPercentage() int32`

GetOriginalRoyaltyPercentage returns the OriginalRoyaltyPercentage field if non-nil, zero value otherwise.

### GetOriginalRoyaltyPercentageOk

`func (o *NftTokenInfo) GetOriginalRoyaltyPercentageOk() (*int32, bool)`

GetOriginalRoyaltyPercentageOk returns a tuple with the OriginalRoyaltyPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalRoyaltyPercentage

`func (o *NftTokenInfo) SetOriginalRoyaltyPercentage(v int32)`

SetOriginalRoyaltyPercentage sets OriginalRoyaltyPercentage field to given value.


### GetStatus

`func (o *NftTokenInfo) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NftTokenInfo) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NftTokenInfo) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetNftFactory

`func (o *NftTokenInfo) GetNftFactory() string`

GetNftFactory returns the NftFactory field if non-nil, zero value otherwise.

### GetNftFactoryOk

`func (o *NftTokenInfo) GetNftFactoryOk() (*string, bool)`

GetNftFactoryOk returns a tuple with the NftFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftFactory

`func (o *NftTokenInfo) SetNftFactory(v string)`

SetNftFactory sets NftFactory field to given value.

### HasNftFactory

`func (o *NftTokenInfo) HasNftFactory() bool`

HasNftFactory returns a boolean if a field has been set.

### GetNftOwner

`func (o *NftTokenInfo) GetNftOwner() string`

GetNftOwner returns the NftOwner field if non-nil, zero value otherwise.

### GetNftOwnerOk

`func (o *NftTokenInfo) GetNftOwnerOk() (*string, bool)`

GetNftOwnerOk returns a tuple with the NftOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftOwner

`func (o *NftTokenInfo) SetNftOwner(v string)`

SetNftOwner sets NftOwner field to given value.

### HasNftOwner

`func (o *NftTokenInfo) HasNftOwner() bool`

HasNftOwner returns a boolean if a field has been set.

### GetNftBaseUri

`func (o *NftTokenInfo) GetNftBaseUri() string`

GetNftBaseUri returns the NftBaseUri field if non-nil, zero value otherwise.

### GetNftBaseUriOk

`func (o *NftTokenInfo) GetNftBaseUriOk() (*string, bool)`

GetNftBaseUriOk returns a tuple with the NftBaseUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftBaseUri

`func (o *NftTokenInfo) SetNftBaseUri(v string)`

SetNftBaseUri sets NftBaseUri field to given value.

### HasNftBaseUri

`func (o *NftTokenInfo) HasNftBaseUri() bool`

HasNftBaseUri returns a boolean if a field has been set.

### GetRoyaltyAddress

`func (o *NftTokenInfo) GetRoyaltyAddress() string`

GetRoyaltyAddress returns the RoyaltyAddress field if non-nil, zero value otherwise.

### GetRoyaltyAddressOk

`func (o *NftTokenInfo) GetRoyaltyAddressOk() (*string, bool)`

GetRoyaltyAddressOk returns a tuple with the RoyaltyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyaltyAddress

`func (o *NftTokenInfo) SetRoyaltyAddress(v string)`

SetRoyaltyAddress sets RoyaltyAddress field to given value.

### HasRoyaltyAddress

`func (o *NftTokenInfo) HasRoyaltyAddress() bool`

HasRoyaltyAddress returns a boolean if a field has been set.

### GetOriginalMinter

`func (o *NftTokenInfo) GetOriginalMinter() string`

GetOriginalMinter returns the OriginalMinter field if non-nil, zero value otherwise.

### GetOriginalMinterOk

`func (o *NftTokenInfo) GetOriginalMinterOk() (*string, bool)`

GetOriginalMinterOk returns a tuple with the OriginalMinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMinter

`func (o *NftTokenInfo) SetOriginalMinter(v string)`

SetOriginalMinter sets OriginalMinter field to given value.

### HasOriginalMinter

`func (o *NftTokenInfo) HasOriginalMinter() bool`

HasOriginalMinter returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NftTokenInfo) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NftTokenInfo) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NftTokenInfo) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NftTokenInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


