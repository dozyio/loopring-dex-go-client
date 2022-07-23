# UpdateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | field.UpdateAccountRequest.exchange | 
**Owner** | **string** | field.UpdateAccountRequest.owner | 
**AccountId** | **int64** | field.UpdateAccountRequest.accountId | 
**ValidUntil** | **int32** | field.UpdateAccountRequest.validUntil | 
**Nonce** | **int32** | field.UpdateAccountRequest.nonce | 
**PublicKeyX** | **string** | field.UpdateAccountRequest.publicKeyX | 
**PublicKeyY** | **string** | field.UpdateAccountRequest.publicKeyY | 
**FeeToken** | **int32** | field.UpdateAccountRequest.feeTokenId | 
**MaxFeeAmount** | **string** | field.UpdateAccountRequest.maxFee | 
**CounterFactualInfo** | Pointer to [**CounterFactualInfo**](CounterFactualInfo.md) |  | [optional] 
**EddsaSig** | Pointer to **string** | field.UpdateAccountRequest.eddsaSig | [optional] 
**EcdsaSig** | Pointer to **string** | field.UpdateAccountRequest.ecdsaSig | [optional] 

## Methods

### NewUpdateAccountRequest

`func NewUpdateAccountRequest(exchange string, owner string, accountId int64, validUntil int32, nonce int32, publicKeyX string, publicKeyY string, feeToken int32, maxFeeAmount string, ) *UpdateAccountRequest`

NewUpdateAccountRequest instantiates a new UpdateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountRequestWithDefaults

`func NewUpdateAccountRequestWithDefaults() *UpdateAccountRequest`

NewUpdateAccountRequestWithDefaults instantiates a new UpdateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *UpdateAccountRequest) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *UpdateAccountRequest) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *UpdateAccountRequest) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetOwner

`func (o *UpdateAccountRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateAccountRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateAccountRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetAccountId

`func (o *UpdateAccountRequest) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateAccountRequest) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateAccountRequest) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetValidUntil

`func (o *UpdateAccountRequest) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *UpdateAccountRequest) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *UpdateAccountRequest) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetNonce

`func (o *UpdateAccountRequest) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *UpdateAccountRequest) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *UpdateAccountRequest) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetPublicKeyX

`func (o *UpdateAccountRequest) GetPublicKeyX() string`

GetPublicKeyX returns the PublicKeyX field if non-nil, zero value otherwise.

### GetPublicKeyXOk

`func (o *UpdateAccountRequest) GetPublicKeyXOk() (*string, bool)`

GetPublicKeyXOk returns a tuple with the PublicKeyX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyX

`func (o *UpdateAccountRequest) SetPublicKeyX(v string)`

SetPublicKeyX sets PublicKeyX field to given value.


### GetPublicKeyY

`func (o *UpdateAccountRequest) GetPublicKeyY() string`

GetPublicKeyY returns the PublicKeyY field if non-nil, zero value otherwise.

### GetPublicKeyYOk

`func (o *UpdateAccountRequest) GetPublicKeyYOk() (*string, bool)`

GetPublicKeyYOk returns a tuple with the PublicKeyY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyY

`func (o *UpdateAccountRequest) SetPublicKeyY(v string)`

SetPublicKeyY sets PublicKeyY field to given value.


### GetFeeToken

`func (o *UpdateAccountRequest) GetFeeToken() int32`

GetFeeToken returns the FeeToken field if non-nil, zero value otherwise.

### GetFeeTokenOk

`func (o *UpdateAccountRequest) GetFeeTokenOk() (*int32, bool)`

GetFeeTokenOk returns a tuple with the FeeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeToken

`func (o *UpdateAccountRequest) SetFeeToken(v int32)`

SetFeeToken sets FeeToken field to given value.


### GetMaxFeeAmount

`func (o *UpdateAccountRequest) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *UpdateAccountRequest) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *UpdateAccountRequest) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.


### GetCounterFactualInfo

`func (o *UpdateAccountRequest) GetCounterFactualInfo() CounterFactualInfo`

GetCounterFactualInfo returns the CounterFactualInfo field if non-nil, zero value otherwise.

### GetCounterFactualInfoOk

`func (o *UpdateAccountRequest) GetCounterFactualInfoOk() (*CounterFactualInfo, bool)`

GetCounterFactualInfoOk returns a tuple with the CounterFactualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterFactualInfo

`func (o *UpdateAccountRequest) SetCounterFactualInfo(v CounterFactualInfo)`

SetCounterFactualInfo sets CounterFactualInfo field to given value.

### HasCounterFactualInfo

`func (o *UpdateAccountRequest) HasCounterFactualInfo() bool`

HasCounterFactualInfo returns a boolean if a field has been set.

### GetEddsaSig

`func (o *UpdateAccountRequest) GetEddsaSig() string`

GetEddsaSig returns the EddsaSig field if non-nil, zero value otherwise.

### GetEddsaSigOk

`func (o *UpdateAccountRequest) GetEddsaSigOk() (*string, bool)`

GetEddsaSigOk returns a tuple with the EddsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSig

`func (o *UpdateAccountRequest) SetEddsaSig(v string)`

SetEddsaSig sets EddsaSig field to given value.

### HasEddsaSig

`func (o *UpdateAccountRequest) HasEddsaSig() bool`

HasEddsaSig returns a boolean if a field has been set.

### GetEcdsaSig

`func (o *UpdateAccountRequest) GetEcdsaSig() string`

GetEcdsaSig returns the EcdsaSig field if non-nil, zero value otherwise.

### GetEcdsaSigOk

`func (o *UpdateAccountRequest) GetEcdsaSigOk() (*string, bool)`

GetEcdsaSigOk returns a tuple with the EcdsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSig

`func (o *UpdateAccountRequest) SetEcdsaSig(v string)`

SetEcdsaSig sets EcdsaSig field to given value.

### HasEcdsaSig

`func (o *UpdateAccountRequest) HasEcdsaSig() bool`

HasEcdsaSig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


