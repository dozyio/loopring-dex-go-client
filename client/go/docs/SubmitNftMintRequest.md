# SubmitNftMintRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | field.SubmitNftMintRequest.exchange | 
**MinterId** | **int64** | minters accountId | 
**MinterAddress** | **string** | minters address | 
**ToAccountId** | **int64** | &#x3D;The account receive the minted NFT token, now should be minter himself. | 
**ToAddress** | Pointer to **string** | field.SubmitNftMintRequest.toAddress | [optional] 
**NftType** | **int32** | nftType: 0 for EIP1155, 1 for EIP712. EIP1155 by default. | 
**TokenAddress** | **string** | field.SubmitNftMintRequest.tokenAddress | 
**NftId** | **string** | NFT_ID | 
**Amount** | **string** | how many tokens to be mint. | 
**CreatorFeeBips** | Pointer to **map[string]interface{}** | fee to the creator of each NFT transaction. | [optional] 
**RoyaltyPercentage** | Pointer to **map[string]interface{}** | field.SubmitNftMintRequest.royaltyPercentage | [optional] 
**ValidUntil** | **int64** | field.SubmitNftMintRequest.validUntil | 
**StorageId** | **int32** | field.SubmitNftMintRequest.storageId | 
**MaxFee** | [**TokenAmountInfo**](TokenAmountInfo.md) |  | 
**EddsaSignature** | Pointer to **string** | field.SubmitNftMintRequest.eddsaSignature | [optional] 
**EcdsaSignature** | Pointer to **string** | field.SubmitNftMintRequest.ecdsaSignature | [optional] 
**HashApproved** | Pointer to **string** | field.SubmitNftMintRequest.hashApproved | [optional] 
**ForceToMint** | Pointer to **bool** | force to mint, regardless the previous mint record | [optional] 
**CounterFactualNftInfo** | Pointer to [**CounterFactualNftInfo**](CounterFactualNftInfo.md) |  | [optional] 
**RoyaltyAddress** | Pointer to **string** | field.CounterFactualNftInfo.royaltyAddress | [optional] 

## Methods

### NewSubmitNftMintRequest

`func NewSubmitNftMintRequest(exchange string, minterId int64, minterAddress string, toAccountId int64, nftType int32, tokenAddress string, nftId string, amount string, validUntil int64, storageId int32, maxFee TokenAmountInfo, ) *SubmitNftMintRequest`

NewSubmitNftMintRequest instantiates a new SubmitNftMintRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitNftMintRequestWithDefaults

`func NewSubmitNftMintRequestWithDefaults() *SubmitNftMintRequest`

NewSubmitNftMintRequestWithDefaults instantiates a new SubmitNftMintRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *SubmitNftMintRequest) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *SubmitNftMintRequest) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *SubmitNftMintRequest) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMinterId

`func (o *SubmitNftMintRequest) GetMinterId() int64`

GetMinterId returns the MinterId field if non-nil, zero value otherwise.

### GetMinterIdOk

`func (o *SubmitNftMintRequest) GetMinterIdOk() (*int64, bool)`

GetMinterIdOk returns a tuple with the MinterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinterId

`func (o *SubmitNftMintRequest) SetMinterId(v int64)`

SetMinterId sets MinterId field to given value.


### GetMinterAddress

`func (o *SubmitNftMintRequest) GetMinterAddress() string`

GetMinterAddress returns the MinterAddress field if non-nil, zero value otherwise.

### GetMinterAddressOk

`func (o *SubmitNftMintRequest) GetMinterAddressOk() (*string, bool)`

GetMinterAddressOk returns a tuple with the MinterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinterAddress

`func (o *SubmitNftMintRequest) SetMinterAddress(v string)`

SetMinterAddress sets MinterAddress field to given value.


### GetToAccountId

`func (o *SubmitNftMintRequest) GetToAccountId() int64`

GetToAccountId returns the ToAccountId field if non-nil, zero value otherwise.

### GetToAccountIdOk

`func (o *SubmitNftMintRequest) GetToAccountIdOk() (*int64, bool)`

GetToAccountIdOk returns a tuple with the ToAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccountId

`func (o *SubmitNftMintRequest) SetToAccountId(v int64)`

SetToAccountId sets ToAccountId field to given value.


### GetToAddress

`func (o *SubmitNftMintRequest) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *SubmitNftMintRequest) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *SubmitNftMintRequest) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *SubmitNftMintRequest) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetNftType

`func (o *SubmitNftMintRequest) GetNftType() int32`

GetNftType returns the NftType field if non-nil, zero value otherwise.

### GetNftTypeOk

`func (o *SubmitNftMintRequest) GetNftTypeOk() (*int32, bool)`

GetNftTypeOk returns a tuple with the NftType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftType

`func (o *SubmitNftMintRequest) SetNftType(v int32)`

SetNftType sets NftType field to given value.


### GetTokenAddress

`func (o *SubmitNftMintRequest) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *SubmitNftMintRequest) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *SubmitNftMintRequest) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetNftId

`func (o *SubmitNftMintRequest) GetNftId() string`

GetNftId returns the NftId field if non-nil, zero value otherwise.

### GetNftIdOk

`func (o *SubmitNftMintRequest) GetNftIdOk() (*string, bool)`

GetNftIdOk returns a tuple with the NftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftId

`func (o *SubmitNftMintRequest) SetNftId(v string)`

SetNftId sets NftId field to given value.


### GetAmount

`func (o *SubmitNftMintRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SubmitNftMintRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SubmitNftMintRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCreatorFeeBips

`func (o *SubmitNftMintRequest) GetCreatorFeeBips() map[string]interface{}`

GetCreatorFeeBips returns the CreatorFeeBips field if non-nil, zero value otherwise.

### GetCreatorFeeBipsOk

`func (o *SubmitNftMintRequest) GetCreatorFeeBipsOk() (*map[string]interface{}, bool)`

GetCreatorFeeBipsOk returns a tuple with the CreatorFeeBips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorFeeBips

`func (o *SubmitNftMintRequest) SetCreatorFeeBips(v map[string]interface{})`

SetCreatorFeeBips sets CreatorFeeBips field to given value.

### HasCreatorFeeBips

`func (o *SubmitNftMintRequest) HasCreatorFeeBips() bool`

HasCreatorFeeBips returns a boolean if a field has been set.

### GetRoyaltyPercentage

`func (o *SubmitNftMintRequest) GetRoyaltyPercentage() map[string]interface{}`

GetRoyaltyPercentage returns the RoyaltyPercentage field if non-nil, zero value otherwise.

### GetRoyaltyPercentageOk

`func (o *SubmitNftMintRequest) GetRoyaltyPercentageOk() (*map[string]interface{}, bool)`

GetRoyaltyPercentageOk returns a tuple with the RoyaltyPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyaltyPercentage

`func (o *SubmitNftMintRequest) SetRoyaltyPercentage(v map[string]interface{})`

SetRoyaltyPercentage sets RoyaltyPercentage field to given value.

### HasRoyaltyPercentage

`func (o *SubmitNftMintRequest) HasRoyaltyPercentage() bool`

HasRoyaltyPercentage returns a boolean if a field has been set.

### GetValidUntil

`func (o *SubmitNftMintRequest) GetValidUntil() int64`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *SubmitNftMintRequest) GetValidUntilOk() (*int64, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *SubmitNftMintRequest) SetValidUntil(v int64)`

SetValidUntil sets ValidUntil field to given value.


### GetStorageId

`func (o *SubmitNftMintRequest) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *SubmitNftMintRequest) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *SubmitNftMintRequest) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.


### GetMaxFee

`func (o *SubmitNftMintRequest) GetMaxFee() TokenAmountInfo`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *SubmitNftMintRequest) GetMaxFeeOk() (*TokenAmountInfo, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *SubmitNftMintRequest) SetMaxFee(v TokenAmountInfo)`

SetMaxFee sets MaxFee field to given value.


### GetEddsaSignature

`func (o *SubmitNftMintRequest) GetEddsaSignature() string`

GetEddsaSignature returns the EddsaSignature field if non-nil, zero value otherwise.

### GetEddsaSignatureOk

`func (o *SubmitNftMintRequest) GetEddsaSignatureOk() (*string, bool)`

GetEddsaSignatureOk returns a tuple with the EddsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSignature

`func (o *SubmitNftMintRequest) SetEddsaSignature(v string)`

SetEddsaSignature sets EddsaSignature field to given value.

### HasEddsaSignature

`func (o *SubmitNftMintRequest) HasEddsaSignature() bool`

HasEddsaSignature returns a boolean if a field has been set.

### GetEcdsaSignature

`func (o *SubmitNftMintRequest) GetEcdsaSignature() string`

GetEcdsaSignature returns the EcdsaSignature field if non-nil, zero value otherwise.

### GetEcdsaSignatureOk

`func (o *SubmitNftMintRequest) GetEcdsaSignatureOk() (*string, bool)`

GetEcdsaSignatureOk returns a tuple with the EcdsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSignature

`func (o *SubmitNftMintRequest) SetEcdsaSignature(v string)`

SetEcdsaSignature sets EcdsaSignature field to given value.

### HasEcdsaSignature

`func (o *SubmitNftMintRequest) HasEcdsaSignature() bool`

HasEcdsaSignature returns a boolean if a field has been set.

### GetHashApproved

`func (o *SubmitNftMintRequest) GetHashApproved() string`

GetHashApproved returns the HashApproved field if non-nil, zero value otherwise.

### GetHashApprovedOk

`func (o *SubmitNftMintRequest) GetHashApprovedOk() (*string, bool)`

GetHashApprovedOk returns a tuple with the HashApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashApproved

`func (o *SubmitNftMintRequest) SetHashApproved(v string)`

SetHashApproved sets HashApproved field to given value.

### HasHashApproved

`func (o *SubmitNftMintRequest) HasHashApproved() bool`

HasHashApproved returns a boolean if a field has been set.

### GetForceToMint

`func (o *SubmitNftMintRequest) GetForceToMint() bool`

GetForceToMint returns the ForceToMint field if non-nil, zero value otherwise.

### GetForceToMintOk

`func (o *SubmitNftMintRequest) GetForceToMintOk() (*bool, bool)`

GetForceToMintOk returns a tuple with the ForceToMint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceToMint

`func (o *SubmitNftMintRequest) SetForceToMint(v bool)`

SetForceToMint sets ForceToMint field to given value.

### HasForceToMint

`func (o *SubmitNftMintRequest) HasForceToMint() bool`

HasForceToMint returns a boolean if a field has been set.

### GetCounterFactualNftInfo

`func (o *SubmitNftMintRequest) GetCounterFactualNftInfo() CounterFactualNftInfo`

GetCounterFactualNftInfo returns the CounterFactualNftInfo field if non-nil, zero value otherwise.

### GetCounterFactualNftInfoOk

`func (o *SubmitNftMintRequest) GetCounterFactualNftInfoOk() (*CounterFactualNftInfo, bool)`

GetCounterFactualNftInfoOk returns a tuple with the CounterFactualNftInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterFactualNftInfo

`func (o *SubmitNftMintRequest) SetCounterFactualNftInfo(v CounterFactualNftInfo)`

SetCounterFactualNftInfo sets CounterFactualNftInfo field to given value.

### HasCounterFactualNftInfo

`func (o *SubmitNftMintRequest) HasCounterFactualNftInfo() bool`

HasCounterFactualNftInfo returns a boolean if a field has been set.

### GetRoyaltyAddress

`func (o *SubmitNftMintRequest) GetRoyaltyAddress() string`

GetRoyaltyAddress returns the RoyaltyAddress field if non-nil, zero value otherwise.

### GetRoyaltyAddressOk

`func (o *SubmitNftMintRequest) GetRoyaltyAddressOk() (*string, bool)`

GetRoyaltyAddressOk returns a tuple with the RoyaltyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyaltyAddress

`func (o *SubmitNftMintRequest) SetRoyaltyAddress(v string)`

SetRoyaltyAddress sets RoyaltyAddress field to given value.

### HasRoyaltyAddress

`func (o *SubmitNftMintRequest) HasRoyaltyAddress() bool`

HasRoyaltyAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


