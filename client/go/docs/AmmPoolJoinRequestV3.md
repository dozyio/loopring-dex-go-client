# AmmPoolJoinRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | The account owner adderss | 
**PoolAddress** | **string** | AMM pool address to be joined | 
**JoinTokens** | [**AmmPoolJoinTokens**](AmmPoolJoinTokens.md) |  | 
**StorageIds** | **string** | Offchain request storage Id | 
**Fee** | **string** | fee of join request | 
**ValidUntil** | **int32** | Timestamp for order to become invalid | 
**EddsaSignature** | Pointer to **string** | AMM join request eddsa signature | [optional] 
**EcdsaSignature** | Pointer to **string** | AMM join request ecdsa signature | [optional] 

## Methods

### NewAmmPoolJoinRequestV3

`func NewAmmPoolJoinRequestV3(owner string, poolAddress string, joinTokens AmmPoolJoinTokens, storageIds string, fee string, validUntil int32, ) *AmmPoolJoinRequestV3`

NewAmmPoolJoinRequestV3 instantiates a new AmmPoolJoinRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmPoolJoinRequestV3WithDefaults

`func NewAmmPoolJoinRequestV3WithDefaults() *AmmPoolJoinRequestV3`

NewAmmPoolJoinRequestV3WithDefaults instantiates a new AmmPoolJoinRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *AmmPoolJoinRequestV3) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AmmPoolJoinRequestV3) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AmmPoolJoinRequestV3) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPoolAddress

`func (o *AmmPoolJoinRequestV3) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *AmmPoolJoinRequestV3) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *AmmPoolJoinRequestV3) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetJoinTokens

`func (o *AmmPoolJoinRequestV3) GetJoinTokens() AmmPoolJoinTokens`

GetJoinTokens returns the JoinTokens field if non-nil, zero value otherwise.

### GetJoinTokensOk

`func (o *AmmPoolJoinRequestV3) GetJoinTokensOk() (*AmmPoolJoinTokens, bool)`

GetJoinTokensOk returns a tuple with the JoinTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTokens

`func (o *AmmPoolJoinRequestV3) SetJoinTokens(v AmmPoolJoinTokens)`

SetJoinTokens sets JoinTokens field to given value.


### GetStorageIds

`func (o *AmmPoolJoinRequestV3) GetStorageIds() string`

GetStorageIds returns the StorageIds field if non-nil, zero value otherwise.

### GetStorageIdsOk

`func (o *AmmPoolJoinRequestV3) GetStorageIdsOk() (*string, bool)`

GetStorageIdsOk returns a tuple with the StorageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageIds

`func (o *AmmPoolJoinRequestV3) SetStorageIds(v string)`

SetStorageIds sets StorageIds field to given value.


### GetFee

`func (o *AmmPoolJoinRequestV3) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *AmmPoolJoinRequestV3) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *AmmPoolJoinRequestV3) SetFee(v string)`

SetFee sets Fee field to given value.


### GetValidUntil

`func (o *AmmPoolJoinRequestV3) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *AmmPoolJoinRequestV3) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *AmmPoolJoinRequestV3) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetEddsaSignature

`func (o *AmmPoolJoinRequestV3) GetEddsaSignature() string`

GetEddsaSignature returns the EddsaSignature field if non-nil, zero value otherwise.

### GetEddsaSignatureOk

`func (o *AmmPoolJoinRequestV3) GetEddsaSignatureOk() (*string, bool)`

GetEddsaSignatureOk returns a tuple with the EddsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSignature

`func (o *AmmPoolJoinRequestV3) SetEddsaSignature(v string)`

SetEddsaSignature sets EddsaSignature field to given value.

### HasEddsaSignature

`func (o *AmmPoolJoinRequestV3) HasEddsaSignature() bool`

HasEddsaSignature returns a boolean if a field has been set.

### GetEcdsaSignature

`func (o *AmmPoolJoinRequestV3) GetEcdsaSignature() string`

GetEcdsaSignature returns the EcdsaSignature field if non-nil, zero value otherwise.

### GetEcdsaSignatureOk

`func (o *AmmPoolJoinRequestV3) GetEcdsaSignatureOk() (*string, bool)`

GetEcdsaSignatureOk returns a tuple with the EcdsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSignature

`func (o *AmmPoolJoinRequestV3) SetEcdsaSignature(v string)`

SetEcdsaSignature sets EcdsaSignature field to given value.

### HasEcdsaSignature

`func (o *AmmPoolJoinRequestV3) HasEcdsaSignature() bool`

HasEcdsaSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


