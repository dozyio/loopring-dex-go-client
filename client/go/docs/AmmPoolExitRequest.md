# AmmPoolExitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** | field.AmmPoolJoinRequest.owner | 
**PoolAddress** | **string** | field.AmmPoolJoinRequest.poolAddress | 
**BurnAmount** | **string** | field.AmmPoolExitRequest.burnAmount | 
**BurnStorageID** | **int64** | field.AmmPoolExitRequest.burnStorageID | 
**ExitMinAmounts** | **string** | field.AmmPoolExitRequest.exitMinAmounts | 
**Fee** | **string** | field.AmmPoolExitRequest.fee | 
**ValidUntil** | **int32** | field.AmmPoolExitRequest.validUntil | 
**EcdsaSig** | **string** | field.AmmPoolExitRequest.ecdsaSig | 
**EddsaSig** | **string** | field.AmmPoolExitRequest.eddsaSig | 

## Methods

### NewAmmPoolExitRequest

`func NewAmmPoolExitRequest(owner string, poolAddress string, burnAmount string, burnStorageID int64, exitMinAmounts string, fee string, validUntil int32, ecdsaSig string, eddsaSig string, ) *AmmPoolExitRequest`

NewAmmPoolExitRequest instantiates a new AmmPoolExitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmPoolExitRequestWithDefaults

`func NewAmmPoolExitRequestWithDefaults() *AmmPoolExitRequest`

NewAmmPoolExitRequestWithDefaults instantiates a new AmmPoolExitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *AmmPoolExitRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AmmPoolExitRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AmmPoolExitRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPoolAddress

`func (o *AmmPoolExitRequest) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *AmmPoolExitRequest) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *AmmPoolExitRequest) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetBurnAmount

`func (o *AmmPoolExitRequest) GetBurnAmount() string`

GetBurnAmount returns the BurnAmount field if non-nil, zero value otherwise.

### GetBurnAmountOk

`func (o *AmmPoolExitRequest) GetBurnAmountOk() (*string, bool)`

GetBurnAmountOk returns a tuple with the BurnAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnAmount

`func (o *AmmPoolExitRequest) SetBurnAmount(v string)`

SetBurnAmount sets BurnAmount field to given value.


### GetBurnStorageID

`func (o *AmmPoolExitRequest) GetBurnStorageID() int64`

GetBurnStorageID returns the BurnStorageID field if non-nil, zero value otherwise.

### GetBurnStorageIDOk

`func (o *AmmPoolExitRequest) GetBurnStorageIDOk() (*int64, bool)`

GetBurnStorageIDOk returns a tuple with the BurnStorageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnStorageID

`func (o *AmmPoolExitRequest) SetBurnStorageID(v int64)`

SetBurnStorageID sets BurnStorageID field to given value.


### GetExitMinAmounts

`func (o *AmmPoolExitRequest) GetExitMinAmounts() string`

GetExitMinAmounts returns the ExitMinAmounts field if non-nil, zero value otherwise.

### GetExitMinAmountsOk

`func (o *AmmPoolExitRequest) GetExitMinAmountsOk() (*string, bool)`

GetExitMinAmountsOk returns a tuple with the ExitMinAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitMinAmounts

`func (o *AmmPoolExitRequest) SetExitMinAmounts(v string)`

SetExitMinAmounts sets ExitMinAmounts field to given value.


### GetFee

`func (o *AmmPoolExitRequest) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *AmmPoolExitRequest) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *AmmPoolExitRequest) SetFee(v string)`

SetFee sets Fee field to given value.


### GetValidUntil

`func (o *AmmPoolExitRequest) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *AmmPoolExitRequest) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *AmmPoolExitRequest) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetEcdsaSig

`func (o *AmmPoolExitRequest) GetEcdsaSig() string`

GetEcdsaSig returns the EcdsaSig field if non-nil, zero value otherwise.

### GetEcdsaSigOk

`func (o *AmmPoolExitRequest) GetEcdsaSigOk() (*string, bool)`

GetEcdsaSigOk returns a tuple with the EcdsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSig

`func (o *AmmPoolExitRequest) SetEcdsaSig(v string)`

SetEcdsaSig sets EcdsaSig field to given value.


### GetEddsaSig

`func (o *AmmPoolExitRequest) GetEddsaSig() string`

GetEddsaSig returns the EddsaSig field if non-nil, zero value otherwise.

### GetEddsaSigOk

`func (o *AmmPoolExitRequest) GetEddsaSigOk() (*string, bool)`

GetEddsaSigOk returns a tuple with the EddsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSig

`func (o *AmmPoolExitRequest) SetEddsaSig(v string)`

SetEddsaSig sets EddsaSig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


