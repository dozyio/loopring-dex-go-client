# AmmPoolJoinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | field.AmmPoolJoinRequest.owner | 
**PoolAddress** | **string** | field.AmmPoolJoinRequest.poolAddress | 
**JoinAmounts** | **string** | field.AmmPoolJoinRequest.joinAmounts | 
**JoinStorageIDs** | **string** | field.AmmPoolJoinRequest.joinStorageIDs | 
**MintMinAmount** | **string** | field.AmmPoolJoinRequest.mintMinAmount | 
**Fee** | **string** | field.AmmPoolJoinRequest.fee | 
**ValidUntil** | **int32** | field.AmmPoolJoinRequest.validUntil | 
**EcdsaSig** | **string** | field.AmmPoolJoinRequest.ecdsaSig | 
**EddsaSig** | **string** | field.AmmPoolJoinRequest.eddsaSig | 

## Methods

### NewAmmPoolJoinRequest

`func NewAmmPoolJoinRequest(owner string, poolAddress string, joinAmounts string, joinStorageIDs string, mintMinAmount string, fee string, validUntil int32, ecdsaSig string, eddsaSig string, ) *AmmPoolJoinRequest`

NewAmmPoolJoinRequest instantiates a new AmmPoolJoinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmPoolJoinRequestWithDefaults

`func NewAmmPoolJoinRequestWithDefaults() *AmmPoolJoinRequest`

NewAmmPoolJoinRequestWithDefaults instantiates a new AmmPoolJoinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *AmmPoolJoinRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AmmPoolJoinRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AmmPoolJoinRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPoolAddress

`func (o *AmmPoolJoinRequest) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *AmmPoolJoinRequest) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *AmmPoolJoinRequest) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetJoinAmounts

`func (o *AmmPoolJoinRequest) GetJoinAmounts() string`

GetJoinAmounts returns the JoinAmounts field if non-nil, zero value otherwise.

### GetJoinAmountsOk

`func (o *AmmPoolJoinRequest) GetJoinAmountsOk() (*string, bool)`

GetJoinAmountsOk returns a tuple with the JoinAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinAmounts

`func (o *AmmPoolJoinRequest) SetJoinAmounts(v string)`

SetJoinAmounts sets JoinAmounts field to given value.


### GetJoinStorageIDs

`func (o *AmmPoolJoinRequest) GetJoinStorageIDs() string`

GetJoinStorageIDs returns the JoinStorageIDs field if non-nil, zero value otherwise.

### GetJoinStorageIDsOk

`func (o *AmmPoolJoinRequest) GetJoinStorageIDsOk() (*string, bool)`

GetJoinStorageIDsOk returns a tuple with the JoinStorageIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinStorageIDs

`func (o *AmmPoolJoinRequest) SetJoinStorageIDs(v string)`

SetJoinStorageIDs sets JoinStorageIDs field to given value.


### GetMintMinAmount

`func (o *AmmPoolJoinRequest) GetMintMinAmount() string`

GetMintMinAmount returns the MintMinAmount field if non-nil, zero value otherwise.

### GetMintMinAmountOk

`func (o *AmmPoolJoinRequest) GetMintMinAmountOk() (*string, bool)`

GetMintMinAmountOk returns a tuple with the MintMinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintMinAmount

`func (o *AmmPoolJoinRequest) SetMintMinAmount(v string)`

SetMintMinAmount sets MintMinAmount field to given value.


### GetFee

`func (o *AmmPoolJoinRequest) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *AmmPoolJoinRequest) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *AmmPoolJoinRequest) SetFee(v string)`

SetFee sets Fee field to given value.


### GetValidUntil

`func (o *AmmPoolJoinRequest) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *AmmPoolJoinRequest) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *AmmPoolJoinRequest) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetEcdsaSig

`func (o *AmmPoolJoinRequest) GetEcdsaSig() string`

GetEcdsaSig returns the EcdsaSig field if non-nil, zero value otherwise.

### GetEcdsaSigOk

`func (o *AmmPoolJoinRequest) GetEcdsaSigOk() (*string, bool)`

GetEcdsaSigOk returns a tuple with the EcdsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSig

`func (o *AmmPoolJoinRequest) SetEcdsaSig(v string)`

SetEcdsaSig sets EcdsaSig field to given value.


### GetEddsaSig

`func (o *AmmPoolJoinRequest) GetEddsaSig() string`

GetEddsaSig returns the EddsaSig field if non-nil, zero value otherwise.

### GetEddsaSigOk

`func (o *AmmPoolJoinRequest) GetEddsaSigOk() (*string, bool)`

GetEddsaSigOk returns a tuple with the EddsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSig

`func (o *AmmPoolJoinRequest) SetEddsaSig(v string)`

SetEddsaSig sets EddsaSig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


