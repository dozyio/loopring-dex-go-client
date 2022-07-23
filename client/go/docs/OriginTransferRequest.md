# OriginTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | field.SubmitOriginTransferRequest.exchange | 
**PayerId** | **int64** | field.SubmitOriginTransferRequest.payerId | 
**PayerAddr** | **string** | field.SubmitOriginTransferRequest.payerAddr | 
**PayeeId** | **int64** | field.SubmitOriginTransferRequest.payeeId | 
**PayeeAddr** | **string** | field.SubmitOriginTransferRequest.payeeAddr | 
**Token** | **int32** | field.SubmitOriginTransferRequest.token | 
**Amount** | **string** | field.SubmitOriginTransferRequest.amount | 
**FeeToken** | **int32** | field.SubmitOriginTransferRequest.feeToken | 
**MaxFeeAmount** | **string** | field.SubmitOriginTransferRequest.maxFeeAmount | 
**StorageId** | **int64** | field.SubmitOriginTransferRequest.storageId | 
**ValidUntil** | **int32** | field.SubmitOriginTransferRequest.validUntil | 
**CounterFactualInfo** | Pointer to [**CounterFactualInfo**](CounterFactualInfo.md) |  | [optional] 
**EddsaSig** | Pointer to **string** | field.SubmitOriginTransferRequest.eddsaSig | [optional] 
**EcdsaSig** | Pointer to **string** | field.SubmitOriginTransferRequest.ecdsaSig | [optional] 
**Memo** | Pointer to **string** | field.SubmitOriginTransferRequest.memo | [optional] 
**ClientId** | Pointer to **string** | field.SubmitOriginTransferRequest.clientId | [optional] 

## Methods

### NewOriginTransferRequest

`func NewOriginTransferRequest(exchange string, payerId int64, payerAddr string, payeeId int64, payeeAddr string, token int32, amount string, feeToken int32, maxFeeAmount string, storageId int64, validUntil int32, ) *OriginTransferRequest`

NewOriginTransferRequest instantiates a new OriginTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginTransferRequestWithDefaults

`func NewOriginTransferRequestWithDefaults() *OriginTransferRequest`

NewOriginTransferRequestWithDefaults instantiates a new OriginTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *OriginTransferRequest) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *OriginTransferRequest) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *OriginTransferRequest) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetPayerId

`func (o *OriginTransferRequest) GetPayerId() int64`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *OriginTransferRequest) GetPayerIdOk() (*int64, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *OriginTransferRequest) SetPayerId(v int64)`

SetPayerId sets PayerId field to given value.


### GetPayerAddr

`func (o *OriginTransferRequest) GetPayerAddr() string`

GetPayerAddr returns the PayerAddr field if non-nil, zero value otherwise.

### GetPayerAddrOk

`func (o *OriginTransferRequest) GetPayerAddrOk() (*string, bool)`

GetPayerAddrOk returns a tuple with the PayerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerAddr

`func (o *OriginTransferRequest) SetPayerAddr(v string)`

SetPayerAddr sets PayerAddr field to given value.


### GetPayeeId

`func (o *OriginTransferRequest) GetPayeeId() int64`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *OriginTransferRequest) GetPayeeIdOk() (*int64, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *OriginTransferRequest) SetPayeeId(v int64)`

SetPayeeId sets PayeeId field to given value.


### GetPayeeAddr

`func (o *OriginTransferRequest) GetPayeeAddr() string`

GetPayeeAddr returns the PayeeAddr field if non-nil, zero value otherwise.

### GetPayeeAddrOk

`func (o *OriginTransferRequest) GetPayeeAddrOk() (*string, bool)`

GetPayeeAddrOk returns a tuple with the PayeeAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeAddr

`func (o *OriginTransferRequest) SetPayeeAddr(v string)`

SetPayeeAddr sets PayeeAddr field to given value.


### GetToken

`func (o *OriginTransferRequest) GetToken() int32`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OriginTransferRequest) GetTokenOk() (*int32, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OriginTransferRequest) SetToken(v int32)`

SetToken sets Token field to given value.


### GetAmount

`func (o *OriginTransferRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OriginTransferRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OriginTransferRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFeeToken

`func (o *OriginTransferRequest) GetFeeToken() int32`

GetFeeToken returns the FeeToken field if non-nil, zero value otherwise.

### GetFeeTokenOk

`func (o *OriginTransferRequest) GetFeeTokenOk() (*int32, bool)`

GetFeeTokenOk returns a tuple with the FeeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeToken

`func (o *OriginTransferRequest) SetFeeToken(v int32)`

SetFeeToken sets FeeToken field to given value.


### GetMaxFeeAmount

`func (o *OriginTransferRequest) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *OriginTransferRequest) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *OriginTransferRequest) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.


### GetStorageId

`func (o *OriginTransferRequest) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *OriginTransferRequest) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *OriginTransferRequest) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.


### GetValidUntil

`func (o *OriginTransferRequest) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *OriginTransferRequest) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *OriginTransferRequest) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetCounterFactualInfo

`func (o *OriginTransferRequest) GetCounterFactualInfo() CounterFactualInfo`

GetCounterFactualInfo returns the CounterFactualInfo field if non-nil, zero value otherwise.

### GetCounterFactualInfoOk

`func (o *OriginTransferRequest) GetCounterFactualInfoOk() (*CounterFactualInfo, bool)`

GetCounterFactualInfoOk returns a tuple with the CounterFactualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterFactualInfo

`func (o *OriginTransferRequest) SetCounterFactualInfo(v CounterFactualInfo)`

SetCounterFactualInfo sets CounterFactualInfo field to given value.

### HasCounterFactualInfo

`func (o *OriginTransferRequest) HasCounterFactualInfo() bool`

HasCounterFactualInfo returns a boolean if a field has been set.

### GetEddsaSig

`func (o *OriginTransferRequest) GetEddsaSig() string`

GetEddsaSig returns the EddsaSig field if non-nil, zero value otherwise.

### GetEddsaSigOk

`func (o *OriginTransferRequest) GetEddsaSigOk() (*string, bool)`

GetEddsaSigOk returns a tuple with the EddsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSig

`func (o *OriginTransferRequest) SetEddsaSig(v string)`

SetEddsaSig sets EddsaSig field to given value.

### HasEddsaSig

`func (o *OriginTransferRequest) HasEddsaSig() bool`

HasEddsaSig returns a boolean if a field has been set.

### GetEcdsaSig

`func (o *OriginTransferRequest) GetEcdsaSig() string`

GetEcdsaSig returns the EcdsaSig field if non-nil, zero value otherwise.

### GetEcdsaSigOk

`func (o *OriginTransferRequest) GetEcdsaSigOk() (*string, bool)`

GetEcdsaSigOk returns a tuple with the EcdsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSig

`func (o *OriginTransferRequest) SetEcdsaSig(v string)`

SetEcdsaSig sets EcdsaSig field to given value.

### HasEcdsaSig

`func (o *OriginTransferRequest) HasEcdsaSig() bool`

HasEcdsaSig returns a boolean if a field has been set.

### GetMemo

`func (o *OriginTransferRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *OriginTransferRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *OriginTransferRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *OriginTransferRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetClientId

`func (o *OriginTransferRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OriginTransferRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OriginTransferRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OriginTransferRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


