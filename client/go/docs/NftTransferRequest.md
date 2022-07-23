# NftTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | exchange address | 
**FromAccountId** | **int64** | payer account ID | 
**FromAddress** | **string** | payer account address | 
**ToAccountId** | **int64** | payee account ID | 
**ToAddress** | **string** | payer account address | 
**Token** | [**NftTokenAmountInfo**](NftTokenAmountInfo.md) |  | 
**MaxFee** | [**TokenAmountInfo**](TokenAmountInfo.md) |  | 
**StorageId** | **int64** | offchain Id | 
**ValidUntil** | **int32** | Timestamp for order to become invalid | 
**CounterFactualInfo** | Pointer to [**CounterFactualInfo**](CounterFactualInfo.md) |  | [optional] 
**EddsaSignature** | Pointer to **string** | eddsa signature | [optional] 
**EcdsaSignature** | Pointer to **string** | ecdsa signature | [optional] 
**HashApproved** | Pointer to **string** | An approved hash string which was already submitted on eth mainnet | [optional] 
**Memo** | Pointer to **string** | transfer memo | [optional] 
**PayPayeeUpdateAccount** | Pointer to **bool** | field.OriginTransferRequestV3.payPayeeUpdateAccount | [optional] 

## Methods

### NewNftTransferRequest

`func NewNftTransferRequest(exchange string, fromAccountId int64, fromAddress string, toAccountId int64, toAddress string, token NftTokenAmountInfo, maxFee TokenAmountInfo, storageId int64, validUntil int32, ) *NftTransferRequest`

NewNftTransferRequest instantiates a new NftTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftTransferRequestWithDefaults

`func NewNftTransferRequestWithDefaults() *NftTransferRequest`

NewNftTransferRequestWithDefaults instantiates a new NftTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *NftTransferRequest) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *NftTransferRequest) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *NftTransferRequest) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetFromAccountId

`func (o *NftTransferRequest) GetFromAccountId() int64`

GetFromAccountId returns the FromAccountId field if non-nil, zero value otherwise.

### GetFromAccountIdOk

`func (o *NftTransferRequest) GetFromAccountIdOk() (*int64, bool)`

GetFromAccountIdOk returns a tuple with the FromAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccountId

`func (o *NftTransferRequest) SetFromAccountId(v int64)`

SetFromAccountId sets FromAccountId field to given value.


### GetFromAddress

`func (o *NftTransferRequest) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *NftTransferRequest) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *NftTransferRequest) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetToAccountId

`func (o *NftTransferRequest) GetToAccountId() int64`

GetToAccountId returns the ToAccountId field if non-nil, zero value otherwise.

### GetToAccountIdOk

`func (o *NftTransferRequest) GetToAccountIdOk() (*int64, bool)`

GetToAccountIdOk returns a tuple with the ToAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccountId

`func (o *NftTransferRequest) SetToAccountId(v int64)`

SetToAccountId sets ToAccountId field to given value.


### GetToAddress

`func (o *NftTransferRequest) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *NftTransferRequest) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *NftTransferRequest) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetToken

`func (o *NftTransferRequest) GetToken() NftTokenAmountInfo`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *NftTransferRequest) GetTokenOk() (*NftTokenAmountInfo, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *NftTransferRequest) SetToken(v NftTokenAmountInfo)`

SetToken sets Token field to given value.


### GetMaxFee

`func (o *NftTransferRequest) GetMaxFee() TokenAmountInfo`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *NftTransferRequest) GetMaxFeeOk() (*TokenAmountInfo, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *NftTransferRequest) SetMaxFee(v TokenAmountInfo)`

SetMaxFee sets MaxFee field to given value.


### GetStorageId

`func (o *NftTransferRequest) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *NftTransferRequest) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *NftTransferRequest) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.


### GetValidUntil

`func (o *NftTransferRequest) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *NftTransferRequest) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *NftTransferRequest) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetCounterFactualInfo

`func (o *NftTransferRequest) GetCounterFactualInfo() CounterFactualInfo`

GetCounterFactualInfo returns the CounterFactualInfo field if non-nil, zero value otherwise.

### GetCounterFactualInfoOk

`func (o *NftTransferRequest) GetCounterFactualInfoOk() (*CounterFactualInfo, bool)`

GetCounterFactualInfoOk returns a tuple with the CounterFactualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterFactualInfo

`func (o *NftTransferRequest) SetCounterFactualInfo(v CounterFactualInfo)`

SetCounterFactualInfo sets CounterFactualInfo field to given value.

### HasCounterFactualInfo

`func (o *NftTransferRequest) HasCounterFactualInfo() bool`

HasCounterFactualInfo returns a boolean if a field has been set.

### GetEddsaSignature

`func (o *NftTransferRequest) GetEddsaSignature() string`

GetEddsaSignature returns the EddsaSignature field if non-nil, zero value otherwise.

### GetEddsaSignatureOk

`func (o *NftTransferRequest) GetEddsaSignatureOk() (*string, bool)`

GetEddsaSignatureOk returns a tuple with the EddsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSignature

`func (o *NftTransferRequest) SetEddsaSignature(v string)`

SetEddsaSignature sets EddsaSignature field to given value.

### HasEddsaSignature

`func (o *NftTransferRequest) HasEddsaSignature() bool`

HasEddsaSignature returns a boolean if a field has been set.

### GetEcdsaSignature

`func (o *NftTransferRequest) GetEcdsaSignature() string`

GetEcdsaSignature returns the EcdsaSignature field if non-nil, zero value otherwise.

### GetEcdsaSignatureOk

`func (o *NftTransferRequest) GetEcdsaSignatureOk() (*string, bool)`

GetEcdsaSignatureOk returns a tuple with the EcdsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSignature

`func (o *NftTransferRequest) SetEcdsaSignature(v string)`

SetEcdsaSignature sets EcdsaSignature field to given value.

### HasEcdsaSignature

`func (o *NftTransferRequest) HasEcdsaSignature() bool`

HasEcdsaSignature returns a boolean if a field has been set.

### GetHashApproved

`func (o *NftTransferRequest) GetHashApproved() string`

GetHashApproved returns the HashApproved field if non-nil, zero value otherwise.

### GetHashApprovedOk

`func (o *NftTransferRequest) GetHashApprovedOk() (*string, bool)`

GetHashApprovedOk returns a tuple with the HashApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashApproved

`func (o *NftTransferRequest) SetHashApproved(v string)`

SetHashApproved sets HashApproved field to given value.

### HasHashApproved

`func (o *NftTransferRequest) HasHashApproved() bool`

HasHashApproved returns a boolean if a field has been set.

### GetMemo

`func (o *NftTransferRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *NftTransferRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *NftTransferRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *NftTransferRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPayPayeeUpdateAccount

`func (o *NftTransferRequest) GetPayPayeeUpdateAccount() bool`

GetPayPayeeUpdateAccount returns the PayPayeeUpdateAccount field if non-nil, zero value otherwise.

### GetPayPayeeUpdateAccountOk

`func (o *NftTransferRequest) GetPayPayeeUpdateAccountOk() (*bool, bool)`

GetPayPayeeUpdateAccountOk returns a tuple with the PayPayeeUpdateAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPayeeUpdateAccount

`func (o *NftTransferRequest) SetPayPayeeUpdateAccount(v bool)`

SetPayPayeeUpdateAccount sets PayPayeeUpdateAccount field to given value.

### HasPayPayeeUpdateAccount

`func (o *NftTransferRequest) HasPayPayeeUpdateAccount() bool`

HasPayPayeeUpdateAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


