# OffChainWithdrawalRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | exchange address | 
**AccountId** | **int64** | account ID | 
**Owner** | **string** | account owner address | 
**Token** | [**TokenVolumeV3**](TokenVolumeV3.md) |  | 
**MaxFee** | [**TokenVolumeV3**](TokenVolumeV3.md) |  | 
**StorageId** | **int64** | offchain ID | 
**ValidUntil** | **int32** | Timestamp for order to become invalid | 
**MinGas** | Pointer to **int32** | min gas for on-chain withdraw, Loopring exchange allocates gas for each distribution, but people can also assign this min gas, so Loopring have to allocate higher gas value for this specific distribution, 0 means let loopring choose the reasonable gas | [optional] 
**To** | **string** | withdraw to address | 
**ExtraData** | Pointer to **string** | extra data for complex withdraw mode, normally none | [optional] 
**FastWithdrawalMode** | Pointer to **bool** | is fast withdraw mode | [optional] 
**CounterFactualInfo** | Pointer to [**CounterFactualInfo**](CounterFactualInfo.md) |  | [optional] 
**EddsaSignature** | Pointer to **string** | eddsa signature | [optional] 
**EcdsaSignature** | Pointer to **string** | ecdsa signature | [optional] 
**HashApproved** | Pointer to **string** | An approved hash string which was already submitted on eth mainnet | [optional] 

## Methods

### NewOffChainWithdrawalRequestV3

`func NewOffChainWithdrawalRequestV3(exchange string, accountId int64, owner string, token TokenVolumeV3, maxFee TokenVolumeV3, storageId int64, validUntil int32, to string, ) *OffChainWithdrawalRequestV3`

NewOffChainWithdrawalRequestV3 instantiates a new OffChainWithdrawalRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffChainWithdrawalRequestV3WithDefaults

`func NewOffChainWithdrawalRequestV3WithDefaults() *OffChainWithdrawalRequestV3`

NewOffChainWithdrawalRequestV3WithDefaults instantiates a new OffChainWithdrawalRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *OffChainWithdrawalRequestV3) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *OffChainWithdrawalRequestV3) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *OffChainWithdrawalRequestV3) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetAccountId

`func (o *OffChainWithdrawalRequestV3) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OffChainWithdrawalRequestV3) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *OffChainWithdrawalRequestV3) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetOwner

`func (o *OffChainWithdrawalRequestV3) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *OffChainWithdrawalRequestV3) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *OffChainWithdrawalRequestV3) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetToken

`func (o *OffChainWithdrawalRequestV3) GetToken() TokenVolumeV3`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OffChainWithdrawalRequestV3) GetTokenOk() (*TokenVolumeV3, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OffChainWithdrawalRequestV3) SetToken(v TokenVolumeV3)`

SetToken sets Token field to given value.


### GetMaxFee

`func (o *OffChainWithdrawalRequestV3) GetMaxFee() TokenVolumeV3`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *OffChainWithdrawalRequestV3) GetMaxFeeOk() (*TokenVolumeV3, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *OffChainWithdrawalRequestV3) SetMaxFee(v TokenVolumeV3)`

SetMaxFee sets MaxFee field to given value.


### GetStorageId

`func (o *OffChainWithdrawalRequestV3) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *OffChainWithdrawalRequestV3) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *OffChainWithdrawalRequestV3) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.


### GetValidUntil

`func (o *OffChainWithdrawalRequestV3) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *OffChainWithdrawalRequestV3) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *OffChainWithdrawalRequestV3) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetMinGas

`func (o *OffChainWithdrawalRequestV3) GetMinGas() int32`

GetMinGas returns the MinGas field if non-nil, zero value otherwise.

### GetMinGasOk

`func (o *OffChainWithdrawalRequestV3) GetMinGasOk() (*int32, bool)`

GetMinGasOk returns a tuple with the MinGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGas

`func (o *OffChainWithdrawalRequestV3) SetMinGas(v int32)`

SetMinGas sets MinGas field to given value.

### HasMinGas

`func (o *OffChainWithdrawalRequestV3) HasMinGas() bool`

HasMinGas returns a boolean if a field has been set.

### GetTo

`func (o *OffChainWithdrawalRequestV3) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *OffChainWithdrawalRequestV3) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *OffChainWithdrawalRequestV3) SetTo(v string)`

SetTo sets To field to given value.


### GetExtraData

`func (o *OffChainWithdrawalRequestV3) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OffChainWithdrawalRequestV3) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OffChainWithdrawalRequestV3) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *OffChainWithdrawalRequestV3) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetFastWithdrawalMode

`func (o *OffChainWithdrawalRequestV3) GetFastWithdrawalMode() bool`

GetFastWithdrawalMode returns the FastWithdrawalMode field if non-nil, zero value otherwise.

### GetFastWithdrawalModeOk

`func (o *OffChainWithdrawalRequestV3) GetFastWithdrawalModeOk() (*bool, bool)`

GetFastWithdrawalModeOk returns a tuple with the FastWithdrawalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastWithdrawalMode

`func (o *OffChainWithdrawalRequestV3) SetFastWithdrawalMode(v bool)`

SetFastWithdrawalMode sets FastWithdrawalMode field to given value.

### HasFastWithdrawalMode

`func (o *OffChainWithdrawalRequestV3) HasFastWithdrawalMode() bool`

HasFastWithdrawalMode returns a boolean if a field has been set.

### GetCounterFactualInfo

`func (o *OffChainWithdrawalRequestV3) GetCounterFactualInfo() CounterFactualInfo`

GetCounterFactualInfo returns the CounterFactualInfo field if non-nil, zero value otherwise.

### GetCounterFactualInfoOk

`func (o *OffChainWithdrawalRequestV3) GetCounterFactualInfoOk() (*CounterFactualInfo, bool)`

GetCounterFactualInfoOk returns a tuple with the CounterFactualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterFactualInfo

`func (o *OffChainWithdrawalRequestV3) SetCounterFactualInfo(v CounterFactualInfo)`

SetCounterFactualInfo sets CounterFactualInfo field to given value.

### HasCounterFactualInfo

`func (o *OffChainWithdrawalRequestV3) HasCounterFactualInfo() bool`

HasCounterFactualInfo returns a boolean if a field has been set.

### GetEddsaSignature

`func (o *OffChainWithdrawalRequestV3) GetEddsaSignature() string`

GetEddsaSignature returns the EddsaSignature field if non-nil, zero value otherwise.

### GetEddsaSignatureOk

`func (o *OffChainWithdrawalRequestV3) GetEddsaSignatureOk() (*string, bool)`

GetEddsaSignatureOk returns a tuple with the EddsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSignature

`func (o *OffChainWithdrawalRequestV3) SetEddsaSignature(v string)`

SetEddsaSignature sets EddsaSignature field to given value.

### HasEddsaSignature

`func (o *OffChainWithdrawalRequestV3) HasEddsaSignature() bool`

HasEddsaSignature returns a boolean if a field has been set.

### GetEcdsaSignature

`func (o *OffChainWithdrawalRequestV3) GetEcdsaSignature() string`

GetEcdsaSignature returns the EcdsaSignature field if non-nil, zero value otherwise.

### GetEcdsaSignatureOk

`func (o *OffChainWithdrawalRequestV3) GetEcdsaSignatureOk() (*string, bool)`

GetEcdsaSignatureOk returns a tuple with the EcdsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSignature

`func (o *OffChainWithdrawalRequestV3) SetEcdsaSignature(v string)`

SetEcdsaSignature sets EcdsaSignature field to given value.

### HasEcdsaSignature

`func (o *OffChainWithdrawalRequestV3) HasEcdsaSignature() bool`

HasEcdsaSignature returns a boolean if a field has been set.

### GetHashApproved

`func (o *OffChainWithdrawalRequestV3) GetHashApproved() string`

GetHashApproved returns the HashApproved field if non-nil, zero value otherwise.

### GetHashApprovedOk

`func (o *OffChainWithdrawalRequestV3) GetHashApprovedOk() (*string, bool)`

GetHashApprovedOk returns a tuple with the HashApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashApproved

`func (o *OffChainWithdrawalRequestV3) SetHashApproved(v string)`

SetHashApproved sets HashApproved field to given value.

### HasHashApproved

`func (o *OffChainWithdrawalRequestV3) HasHashApproved() bool`

HasHashApproved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


