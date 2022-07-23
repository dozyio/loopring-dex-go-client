# NftOffChainWithdrawalRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | exchange address | 
**AccountId** | **int64** | account ID | 
**Owner** | **string** | account owner address | 
**Token** | [**NftTokenAmountInfo**](NftTokenAmountInfo.md) |  | 
**MaxFee** | [**TokenAmountInfo**](TokenAmountInfo.md) |  | 
**StorageId** | **int64** | offchain ID | 
**ValidUntil** | **int32** | Timestamp for order to become invalid | 
**MinGas** | Pointer to **int32** | min gas for on-chain withdraw, Loopring exchange allocates gas for each distribution, but people can also assign this min gas, so Loopring have to allocate higher gas value for this specific distribution, 0 means let loopring choose the reasonable gas | [optional] 
**To** | **string** | withdraw to address | 
**ExtraData** | Pointer to **string** | extra data for complex withdraw mode, normally none | [optional] 
**CounterFactualInfo** | Pointer to [**CounterFactualInfo**](CounterFactualInfo.md) |  | [optional] 
**EddsaSignature** | Pointer to **string** | eddsa signature | [optional] 
**EcdsaSignature** | Pointer to **string** | ecdsa signature | [optional] 
**HashApproved** | Pointer to **string** | An approved hash string which was already submitted on eth mainnet | [optional] 

## Methods

### NewNftOffChainWithdrawalRequestV3

`func NewNftOffChainWithdrawalRequestV3(exchange string, accountId int64, owner string, token NftTokenAmountInfo, maxFee TokenAmountInfo, storageId int64, validUntil int32, to string, ) *NftOffChainWithdrawalRequestV3`

NewNftOffChainWithdrawalRequestV3 instantiates a new NftOffChainWithdrawalRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftOffChainWithdrawalRequestV3WithDefaults

`func NewNftOffChainWithdrawalRequestV3WithDefaults() *NftOffChainWithdrawalRequestV3`

NewNftOffChainWithdrawalRequestV3WithDefaults instantiates a new NftOffChainWithdrawalRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *NftOffChainWithdrawalRequestV3) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *NftOffChainWithdrawalRequestV3) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *NftOffChainWithdrawalRequestV3) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetAccountId

`func (o *NftOffChainWithdrawalRequestV3) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NftOffChainWithdrawalRequestV3) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NftOffChainWithdrawalRequestV3) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetOwner

`func (o *NftOffChainWithdrawalRequestV3) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NftOffChainWithdrawalRequestV3) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NftOffChainWithdrawalRequestV3) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetToken

`func (o *NftOffChainWithdrawalRequestV3) GetToken() NftTokenAmountInfo`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *NftOffChainWithdrawalRequestV3) GetTokenOk() (*NftTokenAmountInfo, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *NftOffChainWithdrawalRequestV3) SetToken(v NftTokenAmountInfo)`

SetToken sets Token field to given value.


### GetMaxFee

`func (o *NftOffChainWithdrawalRequestV3) GetMaxFee() TokenAmountInfo`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *NftOffChainWithdrawalRequestV3) GetMaxFeeOk() (*TokenAmountInfo, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *NftOffChainWithdrawalRequestV3) SetMaxFee(v TokenAmountInfo)`

SetMaxFee sets MaxFee field to given value.


### GetStorageId

`func (o *NftOffChainWithdrawalRequestV3) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *NftOffChainWithdrawalRequestV3) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *NftOffChainWithdrawalRequestV3) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.


### GetValidUntil

`func (o *NftOffChainWithdrawalRequestV3) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *NftOffChainWithdrawalRequestV3) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *NftOffChainWithdrawalRequestV3) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetMinGas

`func (o *NftOffChainWithdrawalRequestV3) GetMinGas() int32`

GetMinGas returns the MinGas field if non-nil, zero value otherwise.

### GetMinGasOk

`func (o *NftOffChainWithdrawalRequestV3) GetMinGasOk() (*int32, bool)`

GetMinGasOk returns a tuple with the MinGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGas

`func (o *NftOffChainWithdrawalRequestV3) SetMinGas(v int32)`

SetMinGas sets MinGas field to given value.

### HasMinGas

`func (o *NftOffChainWithdrawalRequestV3) HasMinGas() bool`

HasMinGas returns a boolean if a field has been set.

### GetTo

`func (o *NftOffChainWithdrawalRequestV3) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *NftOffChainWithdrawalRequestV3) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *NftOffChainWithdrawalRequestV3) SetTo(v string)`

SetTo sets To field to given value.


### GetExtraData

`func (o *NftOffChainWithdrawalRequestV3) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *NftOffChainWithdrawalRequestV3) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *NftOffChainWithdrawalRequestV3) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *NftOffChainWithdrawalRequestV3) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetCounterFactualInfo

`func (o *NftOffChainWithdrawalRequestV3) GetCounterFactualInfo() CounterFactualInfo`

GetCounterFactualInfo returns the CounterFactualInfo field if non-nil, zero value otherwise.

### GetCounterFactualInfoOk

`func (o *NftOffChainWithdrawalRequestV3) GetCounterFactualInfoOk() (*CounterFactualInfo, bool)`

GetCounterFactualInfoOk returns a tuple with the CounterFactualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterFactualInfo

`func (o *NftOffChainWithdrawalRequestV3) SetCounterFactualInfo(v CounterFactualInfo)`

SetCounterFactualInfo sets CounterFactualInfo field to given value.

### HasCounterFactualInfo

`func (o *NftOffChainWithdrawalRequestV3) HasCounterFactualInfo() bool`

HasCounterFactualInfo returns a boolean if a field has been set.

### GetEddsaSignature

`func (o *NftOffChainWithdrawalRequestV3) GetEddsaSignature() string`

GetEddsaSignature returns the EddsaSignature field if non-nil, zero value otherwise.

### GetEddsaSignatureOk

`func (o *NftOffChainWithdrawalRequestV3) GetEddsaSignatureOk() (*string, bool)`

GetEddsaSignatureOk returns a tuple with the EddsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSignature

`func (o *NftOffChainWithdrawalRequestV3) SetEddsaSignature(v string)`

SetEddsaSignature sets EddsaSignature field to given value.

### HasEddsaSignature

`func (o *NftOffChainWithdrawalRequestV3) HasEddsaSignature() bool`

HasEddsaSignature returns a boolean if a field has been set.

### GetEcdsaSignature

`func (o *NftOffChainWithdrawalRequestV3) GetEcdsaSignature() string`

GetEcdsaSignature returns the EcdsaSignature field if non-nil, zero value otherwise.

### GetEcdsaSignatureOk

`func (o *NftOffChainWithdrawalRequestV3) GetEcdsaSignatureOk() (*string, bool)`

GetEcdsaSignatureOk returns a tuple with the EcdsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSignature

`func (o *NftOffChainWithdrawalRequestV3) SetEcdsaSignature(v string)`

SetEcdsaSignature sets EcdsaSignature field to given value.

### HasEcdsaSignature

`func (o *NftOffChainWithdrawalRequestV3) HasEcdsaSignature() bool`

HasEcdsaSignature returns a boolean if a field has been set.

### GetHashApproved

`func (o *NftOffChainWithdrawalRequestV3) GetHashApproved() string`

GetHashApproved returns the HashApproved field if non-nil, zero value otherwise.

### GetHashApprovedOk

`func (o *NftOffChainWithdrawalRequestV3) GetHashApprovedOk() (*string, bool)`

GetHashApprovedOk returns a tuple with the HashApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashApproved

`func (o *NftOffChainWithdrawalRequestV3) SetHashApproved(v string)`

SetHashApproved sets HashApproved field to given value.

### HasHashApproved

`func (o *NftOffChainWithdrawalRequestV3) HasHashApproved() bool`

HasHashApproved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


