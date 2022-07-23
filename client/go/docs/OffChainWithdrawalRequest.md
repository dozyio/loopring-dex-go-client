# OffChainWithdrawalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | exchange address | 
**AccountId** | **int64** | account id in exchange | 
**Owner** | **string** | owner of accountId in exchange | 
**Token** | **int32** | withdraw token id | 
**Amount** | **string** | withdraw amount of token, decimal string in WEI | 
**FeeToken** | **int32** | fee token id in exchange | 
**MaxFeeAmount** | **string** | fee amount of token, decimal string in WEI | 
**StorageId** | **int64** | dex offchain request storageID of the account | 
**ValidUntil** | **int32** | valid until | 
**MinGas** | **int32** | min gas | 
**To** | **string** | to address | 
**ExtraData** | Pointer to **string** | extra data | [optional] 
**FastWithdrawalMode** | Pointer to **bool** | fastWithdrawal mode | [optional] 
**CounterFactualInfo** | Pointer to [**CounterFactualInfo**](CounterFactualInfo.md) |  | [optional] 
**EddsaSig** | Pointer to **string** | eddsaSig of this request, hex string | [optional] 
**EcdsaSig** | Pointer to **string** | trading ecdsa_sig signature of this order, decimal string | [optional] 

## Methods

### NewOffChainWithdrawalRequest

`func NewOffChainWithdrawalRequest(exchange string, accountId int64, owner string, token int32, amount string, feeToken int32, maxFeeAmount string, storageId int64, validUntil int32, minGas int32, to string, ) *OffChainWithdrawalRequest`

NewOffChainWithdrawalRequest instantiates a new OffChainWithdrawalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffChainWithdrawalRequestWithDefaults

`func NewOffChainWithdrawalRequestWithDefaults() *OffChainWithdrawalRequest`

NewOffChainWithdrawalRequestWithDefaults instantiates a new OffChainWithdrawalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *OffChainWithdrawalRequest) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *OffChainWithdrawalRequest) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *OffChainWithdrawalRequest) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetAccountId

`func (o *OffChainWithdrawalRequest) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OffChainWithdrawalRequest) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *OffChainWithdrawalRequest) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetOwner

`func (o *OffChainWithdrawalRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *OffChainWithdrawalRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *OffChainWithdrawalRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetToken

`func (o *OffChainWithdrawalRequest) GetToken() int32`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OffChainWithdrawalRequest) GetTokenOk() (*int32, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OffChainWithdrawalRequest) SetToken(v int32)`

SetToken sets Token field to given value.


### GetAmount

`func (o *OffChainWithdrawalRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OffChainWithdrawalRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OffChainWithdrawalRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFeeToken

`func (o *OffChainWithdrawalRequest) GetFeeToken() int32`

GetFeeToken returns the FeeToken field if non-nil, zero value otherwise.

### GetFeeTokenOk

`func (o *OffChainWithdrawalRequest) GetFeeTokenOk() (*int32, bool)`

GetFeeTokenOk returns a tuple with the FeeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeToken

`func (o *OffChainWithdrawalRequest) SetFeeToken(v int32)`

SetFeeToken sets FeeToken field to given value.


### GetMaxFeeAmount

`func (o *OffChainWithdrawalRequest) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *OffChainWithdrawalRequest) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *OffChainWithdrawalRequest) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.


### GetStorageId

`func (o *OffChainWithdrawalRequest) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *OffChainWithdrawalRequest) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *OffChainWithdrawalRequest) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.


### GetValidUntil

`func (o *OffChainWithdrawalRequest) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *OffChainWithdrawalRequest) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *OffChainWithdrawalRequest) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetMinGas

`func (o *OffChainWithdrawalRequest) GetMinGas() int32`

GetMinGas returns the MinGas field if non-nil, zero value otherwise.

### GetMinGasOk

`func (o *OffChainWithdrawalRequest) GetMinGasOk() (*int32, bool)`

GetMinGasOk returns a tuple with the MinGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGas

`func (o *OffChainWithdrawalRequest) SetMinGas(v int32)`

SetMinGas sets MinGas field to given value.


### GetTo

`func (o *OffChainWithdrawalRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *OffChainWithdrawalRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *OffChainWithdrawalRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetExtraData

`func (o *OffChainWithdrawalRequest) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *OffChainWithdrawalRequest) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *OffChainWithdrawalRequest) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *OffChainWithdrawalRequest) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetFastWithdrawalMode

`func (o *OffChainWithdrawalRequest) GetFastWithdrawalMode() bool`

GetFastWithdrawalMode returns the FastWithdrawalMode field if non-nil, zero value otherwise.

### GetFastWithdrawalModeOk

`func (o *OffChainWithdrawalRequest) GetFastWithdrawalModeOk() (*bool, bool)`

GetFastWithdrawalModeOk returns a tuple with the FastWithdrawalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastWithdrawalMode

`func (o *OffChainWithdrawalRequest) SetFastWithdrawalMode(v bool)`

SetFastWithdrawalMode sets FastWithdrawalMode field to given value.

### HasFastWithdrawalMode

`func (o *OffChainWithdrawalRequest) HasFastWithdrawalMode() bool`

HasFastWithdrawalMode returns a boolean if a field has been set.

### GetCounterFactualInfo

`func (o *OffChainWithdrawalRequest) GetCounterFactualInfo() CounterFactualInfo`

GetCounterFactualInfo returns the CounterFactualInfo field if non-nil, zero value otherwise.

### GetCounterFactualInfoOk

`func (o *OffChainWithdrawalRequest) GetCounterFactualInfoOk() (*CounterFactualInfo, bool)`

GetCounterFactualInfoOk returns a tuple with the CounterFactualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterFactualInfo

`func (o *OffChainWithdrawalRequest) SetCounterFactualInfo(v CounterFactualInfo)`

SetCounterFactualInfo sets CounterFactualInfo field to given value.

### HasCounterFactualInfo

`func (o *OffChainWithdrawalRequest) HasCounterFactualInfo() bool`

HasCounterFactualInfo returns a boolean if a field has been set.

### GetEddsaSig

`func (o *OffChainWithdrawalRequest) GetEddsaSig() string`

GetEddsaSig returns the EddsaSig field if non-nil, zero value otherwise.

### GetEddsaSigOk

`func (o *OffChainWithdrawalRequest) GetEddsaSigOk() (*string, bool)`

GetEddsaSigOk returns a tuple with the EddsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSig

`func (o *OffChainWithdrawalRequest) SetEddsaSig(v string)`

SetEddsaSig sets EddsaSig field to given value.

### HasEddsaSig

`func (o *OffChainWithdrawalRequest) HasEddsaSig() bool`

HasEddsaSig returns a boolean if a field has been set.

### GetEcdsaSig

`func (o *OffChainWithdrawalRequest) GetEcdsaSig() string`

GetEcdsaSig returns the EcdsaSig field if non-nil, zero value otherwise.

### GetEcdsaSigOk

`func (o *OffChainWithdrawalRequest) GetEcdsaSigOk() (*string, bool)`

GetEcdsaSigOk returns a tuple with the EcdsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSig

`func (o *OffChainWithdrawalRequest) SetEcdsaSig(v string)`

SetEcdsaSig sets EcdsaSig field to given value.

### HasEcdsaSig

`func (o *OffChainWithdrawalRequest) HasEcdsaSig() bool`

HasEcdsaSig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


