# AmmPoolExitRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | The account owner adderss | 
**PoolAddress** | **string** | AMM pool address to be joined | 
**ExitTokens** | [**AmmPoolExitTokens**](AmmPoolExitTokens.md) |  | 
**StorageId** | **int64** | Offchain request storage Id | 
**MaxFee** | **string** | Maximum fee of exit request, use the last in pool token by default | 
**ValidUntil** | **int32** | Timestamp for order to become invalid | 
**EddsaSignature** | Pointer to **string** | AMM exit request eddsa signature | [optional] 
**EcdsaSignature** | Pointer to **string** | AMM exit request ecdsa signature | [optional] 

## Methods

### NewAmmPoolExitRequestV3

`func NewAmmPoolExitRequestV3(owner string, poolAddress string, exitTokens AmmPoolExitTokens, storageId int64, maxFee string, validUntil int32, ) *AmmPoolExitRequestV3`

NewAmmPoolExitRequestV3 instantiates a new AmmPoolExitRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmPoolExitRequestV3WithDefaults

`func NewAmmPoolExitRequestV3WithDefaults() *AmmPoolExitRequestV3`

NewAmmPoolExitRequestV3WithDefaults instantiates a new AmmPoolExitRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *AmmPoolExitRequestV3) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AmmPoolExitRequestV3) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AmmPoolExitRequestV3) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPoolAddress

`func (o *AmmPoolExitRequestV3) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *AmmPoolExitRequestV3) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *AmmPoolExitRequestV3) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetExitTokens

`func (o *AmmPoolExitRequestV3) GetExitTokens() AmmPoolExitTokens`

GetExitTokens returns the ExitTokens field if non-nil, zero value otherwise.

### GetExitTokensOk

`func (o *AmmPoolExitRequestV3) GetExitTokensOk() (*AmmPoolExitTokens, bool)`

GetExitTokensOk returns a tuple with the ExitTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitTokens

`func (o *AmmPoolExitRequestV3) SetExitTokens(v AmmPoolExitTokens)`

SetExitTokens sets ExitTokens field to given value.


### GetStorageId

`func (o *AmmPoolExitRequestV3) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *AmmPoolExitRequestV3) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *AmmPoolExitRequestV3) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.


### GetMaxFee

`func (o *AmmPoolExitRequestV3) GetMaxFee() string`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *AmmPoolExitRequestV3) GetMaxFeeOk() (*string, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *AmmPoolExitRequestV3) SetMaxFee(v string)`

SetMaxFee sets MaxFee field to given value.


### GetValidUntil

`func (o *AmmPoolExitRequestV3) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *AmmPoolExitRequestV3) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *AmmPoolExitRequestV3) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetEddsaSignature

`func (o *AmmPoolExitRequestV3) GetEddsaSignature() string`

GetEddsaSignature returns the EddsaSignature field if non-nil, zero value otherwise.

### GetEddsaSignatureOk

`func (o *AmmPoolExitRequestV3) GetEddsaSignatureOk() (*string, bool)`

GetEddsaSignatureOk returns a tuple with the EddsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSignature

`func (o *AmmPoolExitRequestV3) SetEddsaSignature(v string)`

SetEddsaSignature sets EddsaSignature field to given value.

### HasEddsaSignature

`func (o *AmmPoolExitRequestV3) HasEddsaSignature() bool`

HasEddsaSignature returns a boolean if a field has been set.

### GetEcdsaSignature

`func (o *AmmPoolExitRequestV3) GetEcdsaSignature() string`

GetEcdsaSignature returns the EcdsaSignature field if non-nil, zero value otherwise.

### GetEcdsaSignatureOk

`func (o *AmmPoolExitRequestV3) GetEcdsaSignatureOk() (*string, bool)`

GetEcdsaSignatureOk returns a tuple with the EcdsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSignature

`func (o *AmmPoolExitRequestV3) SetEcdsaSignature(v string)`

SetEcdsaSignature sets EcdsaSignature field to given value.

### HasEcdsaSignature

`func (o *AmmPoolExitRequestV3) HasEcdsaSignature() bool`

HasEcdsaSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


