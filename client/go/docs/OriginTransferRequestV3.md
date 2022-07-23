# OriginTransferRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | exchange address | 
**PayerId** | **int64** | payer account ID | 
**PayerAddr** | **string** | payer account address | 
**PayeeId** | **int64** | payee account ID | 
**PayeeAddr** | **string** | payer account address | 
**Token** | [**TokenVolumeV3**](TokenVolumeV3.md) |  | 
**MaxFee** | [**TokenVolumeV3**](TokenVolumeV3.md) |  | 
**StorageId** | **int64** | offchain Id | 
**ValidUntil** | **int32** | Timestamp for order to become invalid | 
**CounterFactualInfo** | Pointer to [**CounterFactualInfo**](CounterFactualInfo.md) |  | [optional] 
**EddsaSignature** | Pointer to **string** | eddsa signature | [optional] 
**EcdsaSignature** | Pointer to **string** | ecdsa signature | [optional] 
**HashApproved** | Pointer to **string** | An approved hash string which was already submitted on eth mainnet | [optional] 
**Memo** | Pointer to **string** | transfer memo | [optional] 
**ClientId** | Pointer to **string** | A user-defined id | [optional] 
**PayPayeeUpdateAccount** | Pointer to **bool** | field.OriginTransferRequestV3.payPayeeUpdateAccount | [optional] 

## Methods

### NewOriginTransferRequestV3

`func NewOriginTransferRequestV3(exchange string, payerId int64, payerAddr string, payeeId int64, payeeAddr string, token TokenVolumeV3, maxFee TokenVolumeV3, storageId int64, validUntil int32, ) *OriginTransferRequestV3`

NewOriginTransferRequestV3 instantiates a new OriginTransferRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginTransferRequestV3WithDefaults

`func NewOriginTransferRequestV3WithDefaults() *OriginTransferRequestV3`

NewOriginTransferRequestV3WithDefaults instantiates a new OriginTransferRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *OriginTransferRequestV3) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *OriginTransferRequestV3) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *OriginTransferRequestV3) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetPayerId

`func (o *OriginTransferRequestV3) GetPayerId() int64`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *OriginTransferRequestV3) GetPayerIdOk() (*int64, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *OriginTransferRequestV3) SetPayerId(v int64)`

SetPayerId sets PayerId field to given value.


### GetPayerAddr

`func (o *OriginTransferRequestV3) GetPayerAddr() string`

GetPayerAddr returns the PayerAddr field if non-nil, zero value otherwise.

### GetPayerAddrOk

`func (o *OriginTransferRequestV3) GetPayerAddrOk() (*string, bool)`

GetPayerAddrOk returns a tuple with the PayerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerAddr

`func (o *OriginTransferRequestV3) SetPayerAddr(v string)`

SetPayerAddr sets PayerAddr field to given value.


### GetPayeeId

`func (o *OriginTransferRequestV3) GetPayeeId() int64`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *OriginTransferRequestV3) GetPayeeIdOk() (*int64, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *OriginTransferRequestV3) SetPayeeId(v int64)`

SetPayeeId sets PayeeId field to given value.


### GetPayeeAddr

`func (o *OriginTransferRequestV3) GetPayeeAddr() string`

GetPayeeAddr returns the PayeeAddr field if non-nil, zero value otherwise.

### GetPayeeAddrOk

`func (o *OriginTransferRequestV3) GetPayeeAddrOk() (*string, bool)`

GetPayeeAddrOk returns a tuple with the PayeeAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeAddr

`func (o *OriginTransferRequestV3) SetPayeeAddr(v string)`

SetPayeeAddr sets PayeeAddr field to given value.


### GetToken

`func (o *OriginTransferRequestV3) GetToken() TokenVolumeV3`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OriginTransferRequestV3) GetTokenOk() (*TokenVolumeV3, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OriginTransferRequestV3) SetToken(v TokenVolumeV3)`

SetToken sets Token field to given value.


### GetMaxFee

`func (o *OriginTransferRequestV3) GetMaxFee() TokenVolumeV3`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *OriginTransferRequestV3) GetMaxFeeOk() (*TokenVolumeV3, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *OriginTransferRequestV3) SetMaxFee(v TokenVolumeV3)`

SetMaxFee sets MaxFee field to given value.


### GetStorageId

`func (o *OriginTransferRequestV3) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *OriginTransferRequestV3) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *OriginTransferRequestV3) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.


### GetValidUntil

`func (o *OriginTransferRequestV3) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *OriginTransferRequestV3) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *OriginTransferRequestV3) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetCounterFactualInfo

`func (o *OriginTransferRequestV3) GetCounterFactualInfo() CounterFactualInfo`

GetCounterFactualInfo returns the CounterFactualInfo field if non-nil, zero value otherwise.

### GetCounterFactualInfoOk

`func (o *OriginTransferRequestV3) GetCounterFactualInfoOk() (*CounterFactualInfo, bool)`

GetCounterFactualInfoOk returns a tuple with the CounterFactualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterFactualInfo

`func (o *OriginTransferRequestV3) SetCounterFactualInfo(v CounterFactualInfo)`

SetCounterFactualInfo sets CounterFactualInfo field to given value.

### HasCounterFactualInfo

`func (o *OriginTransferRequestV3) HasCounterFactualInfo() bool`

HasCounterFactualInfo returns a boolean if a field has been set.

### GetEddsaSignature

`func (o *OriginTransferRequestV3) GetEddsaSignature() string`

GetEddsaSignature returns the EddsaSignature field if non-nil, zero value otherwise.

### GetEddsaSignatureOk

`func (o *OriginTransferRequestV3) GetEddsaSignatureOk() (*string, bool)`

GetEddsaSignatureOk returns a tuple with the EddsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSignature

`func (o *OriginTransferRequestV3) SetEddsaSignature(v string)`

SetEddsaSignature sets EddsaSignature field to given value.

### HasEddsaSignature

`func (o *OriginTransferRequestV3) HasEddsaSignature() bool`

HasEddsaSignature returns a boolean if a field has been set.

### GetEcdsaSignature

`func (o *OriginTransferRequestV3) GetEcdsaSignature() string`

GetEcdsaSignature returns the EcdsaSignature field if non-nil, zero value otherwise.

### GetEcdsaSignatureOk

`func (o *OriginTransferRequestV3) GetEcdsaSignatureOk() (*string, bool)`

GetEcdsaSignatureOk returns a tuple with the EcdsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSignature

`func (o *OriginTransferRequestV3) SetEcdsaSignature(v string)`

SetEcdsaSignature sets EcdsaSignature field to given value.

### HasEcdsaSignature

`func (o *OriginTransferRequestV3) HasEcdsaSignature() bool`

HasEcdsaSignature returns a boolean if a field has been set.

### GetHashApproved

`func (o *OriginTransferRequestV3) GetHashApproved() string`

GetHashApproved returns the HashApproved field if non-nil, zero value otherwise.

### GetHashApprovedOk

`func (o *OriginTransferRequestV3) GetHashApprovedOk() (*string, bool)`

GetHashApprovedOk returns a tuple with the HashApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashApproved

`func (o *OriginTransferRequestV3) SetHashApproved(v string)`

SetHashApproved sets HashApproved field to given value.

### HasHashApproved

`func (o *OriginTransferRequestV3) HasHashApproved() bool`

HasHashApproved returns a boolean if a field has been set.

### GetMemo

`func (o *OriginTransferRequestV3) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *OriginTransferRequestV3) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *OriginTransferRequestV3) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *OriginTransferRequestV3) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetClientId

`func (o *OriginTransferRequestV3) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OriginTransferRequestV3) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OriginTransferRequestV3) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OriginTransferRequestV3) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetPayPayeeUpdateAccount

`func (o *OriginTransferRequestV3) GetPayPayeeUpdateAccount() bool`

GetPayPayeeUpdateAccount returns the PayPayeeUpdateAccount field if non-nil, zero value otherwise.

### GetPayPayeeUpdateAccountOk

`func (o *OriginTransferRequestV3) GetPayPayeeUpdateAccountOk() (*bool, bool)`

GetPayPayeeUpdateAccountOk returns a tuple with the PayPayeeUpdateAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPayeeUpdateAccount

`func (o *OriginTransferRequestV3) SetPayPayeeUpdateAccount(v bool)`

SetPayPayeeUpdateAccount sets PayPayeeUpdateAccount field to given value.

### HasPayPayeeUpdateAccount

`func (o *OriginTransferRequestV3) HasPayPayeeUpdateAccount() bool`

HasPayPayeeUpdateAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


