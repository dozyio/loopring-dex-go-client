# DualAuthTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | field.SubmitDualAuthTransferRequest.exchange | 
**PayerId** | **int32** | field.SubmitDualAuthTransferRequest.payerId | 
**PayerAddr** | **string** | field.SubmitDualAuthTransferRequest.payerAddr | 
**Token** | **int32** | field.SubmitDualAuthTransferRequest.token | 
**Amount** | **string** | field.SubmitDualAuthTransferRequest.amount | 
**FeeToken** | **int32** | field.SubmitDualAuthTransferRequest.feeToken | 
**MaxFeeAmount** | **string** | field.SubmitDualAuthTransferRequest.maxFeeAmount | 
**StorageId** | **int64** | field.SubmitDualAuthTransferRequest.storageId | 
**ValidUntil** | **int32** | field.SubmitDualAuthTransferRequest.validUntil | 
**EddsaSig** | **string** | field.SubmitDualAuthTransferRequest.eddsaSig | 
**DualAuthKeyX** | **string** | field.SubmitDualAuthTransferRequest.dualAuthKeyX | 
**DualAuthKeyY** | **string** | field.SubmitDualAuthTransferRequest.dualAuthKeyY | 
**DualEddsaSig** | **string** | field.SubmitDualAuthTransferRequest.dualEddsaSig | 
**DualPayeeId** | **int64** | field.SubmitDualAuthTransferRequest.dualPayeeId | 
**DualPayeeAddr** | **string** | field.SubmitDualAuthTransferRequest.dualPayeeAddr | 
**Memo** | Pointer to **string** | field.SubmitDualAuthTransferRequest.memo | [optional] 

## Methods

### NewDualAuthTransferRequest

`func NewDualAuthTransferRequest(exchange string, payerId int32, payerAddr string, token int32, amount string, feeToken int32, maxFeeAmount string, storageId int64, validUntil int32, eddsaSig string, dualAuthKeyX string, dualAuthKeyY string, dualEddsaSig string, dualPayeeId int64, dualPayeeAddr string, ) *DualAuthTransferRequest`

NewDualAuthTransferRequest instantiates a new DualAuthTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDualAuthTransferRequestWithDefaults

`func NewDualAuthTransferRequestWithDefaults() *DualAuthTransferRequest`

NewDualAuthTransferRequestWithDefaults instantiates a new DualAuthTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *DualAuthTransferRequest) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *DualAuthTransferRequest) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *DualAuthTransferRequest) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetPayerId

`func (o *DualAuthTransferRequest) GetPayerId() int32`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *DualAuthTransferRequest) GetPayerIdOk() (*int32, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *DualAuthTransferRequest) SetPayerId(v int32)`

SetPayerId sets PayerId field to given value.


### GetPayerAddr

`func (o *DualAuthTransferRequest) GetPayerAddr() string`

GetPayerAddr returns the PayerAddr field if non-nil, zero value otherwise.

### GetPayerAddrOk

`func (o *DualAuthTransferRequest) GetPayerAddrOk() (*string, bool)`

GetPayerAddrOk returns a tuple with the PayerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerAddr

`func (o *DualAuthTransferRequest) SetPayerAddr(v string)`

SetPayerAddr sets PayerAddr field to given value.


### GetToken

`func (o *DualAuthTransferRequest) GetToken() int32`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DualAuthTransferRequest) GetTokenOk() (*int32, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DualAuthTransferRequest) SetToken(v int32)`

SetToken sets Token field to given value.


### GetAmount

`func (o *DualAuthTransferRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DualAuthTransferRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DualAuthTransferRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFeeToken

`func (o *DualAuthTransferRequest) GetFeeToken() int32`

GetFeeToken returns the FeeToken field if non-nil, zero value otherwise.

### GetFeeTokenOk

`func (o *DualAuthTransferRequest) GetFeeTokenOk() (*int32, bool)`

GetFeeTokenOk returns a tuple with the FeeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeToken

`func (o *DualAuthTransferRequest) SetFeeToken(v int32)`

SetFeeToken sets FeeToken field to given value.


### GetMaxFeeAmount

`func (o *DualAuthTransferRequest) GetMaxFeeAmount() string`

GetMaxFeeAmount returns the MaxFeeAmount field if non-nil, zero value otherwise.

### GetMaxFeeAmountOk

`func (o *DualAuthTransferRequest) GetMaxFeeAmountOk() (*string, bool)`

GetMaxFeeAmountOk returns a tuple with the MaxFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeAmount

`func (o *DualAuthTransferRequest) SetMaxFeeAmount(v string)`

SetMaxFeeAmount sets MaxFeeAmount field to given value.


### GetStorageId

`func (o *DualAuthTransferRequest) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *DualAuthTransferRequest) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *DualAuthTransferRequest) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.


### GetValidUntil

`func (o *DualAuthTransferRequest) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *DualAuthTransferRequest) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *DualAuthTransferRequest) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetEddsaSig

`func (o *DualAuthTransferRequest) GetEddsaSig() string`

GetEddsaSig returns the EddsaSig field if non-nil, zero value otherwise.

### GetEddsaSigOk

`func (o *DualAuthTransferRequest) GetEddsaSigOk() (*string, bool)`

GetEddsaSigOk returns a tuple with the EddsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSig

`func (o *DualAuthTransferRequest) SetEddsaSig(v string)`

SetEddsaSig sets EddsaSig field to given value.


### GetDualAuthKeyX

`func (o *DualAuthTransferRequest) GetDualAuthKeyX() string`

GetDualAuthKeyX returns the DualAuthKeyX field if non-nil, zero value otherwise.

### GetDualAuthKeyXOk

`func (o *DualAuthTransferRequest) GetDualAuthKeyXOk() (*string, bool)`

GetDualAuthKeyXOk returns a tuple with the DualAuthKeyX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualAuthKeyX

`func (o *DualAuthTransferRequest) SetDualAuthKeyX(v string)`

SetDualAuthKeyX sets DualAuthKeyX field to given value.


### GetDualAuthKeyY

`func (o *DualAuthTransferRequest) GetDualAuthKeyY() string`

GetDualAuthKeyY returns the DualAuthKeyY field if non-nil, zero value otherwise.

### GetDualAuthKeyYOk

`func (o *DualAuthTransferRequest) GetDualAuthKeyYOk() (*string, bool)`

GetDualAuthKeyYOk returns a tuple with the DualAuthKeyY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualAuthKeyY

`func (o *DualAuthTransferRequest) SetDualAuthKeyY(v string)`

SetDualAuthKeyY sets DualAuthKeyY field to given value.


### GetDualEddsaSig

`func (o *DualAuthTransferRequest) GetDualEddsaSig() string`

GetDualEddsaSig returns the DualEddsaSig field if non-nil, zero value otherwise.

### GetDualEddsaSigOk

`func (o *DualAuthTransferRequest) GetDualEddsaSigOk() (*string, bool)`

GetDualEddsaSigOk returns a tuple with the DualEddsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualEddsaSig

`func (o *DualAuthTransferRequest) SetDualEddsaSig(v string)`

SetDualEddsaSig sets DualEddsaSig field to given value.


### GetDualPayeeId

`func (o *DualAuthTransferRequest) GetDualPayeeId() int64`

GetDualPayeeId returns the DualPayeeId field if non-nil, zero value otherwise.

### GetDualPayeeIdOk

`func (o *DualAuthTransferRequest) GetDualPayeeIdOk() (*int64, bool)`

GetDualPayeeIdOk returns a tuple with the DualPayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualPayeeId

`func (o *DualAuthTransferRequest) SetDualPayeeId(v int64)`

SetDualPayeeId sets DualPayeeId field to given value.


### GetDualPayeeAddr

`func (o *DualAuthTransferRequest) GetDualPayeeAddr() string`

GetDualPayeeAddr returns the DualPayeeAddr field if non-nil, zero value otherwise.

### GetDualPayeeAddrOk

`func (o *DualAuthTransferRequest) GetDualPayeeAddrOk() (*string, bool)`

GetDualPayeeAddrOk returns a tuple with the DualPayeeAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualPayeeAddr

`func (o *DualAuthTransferRequest) SetDualPayeeAddr(v string)`

SetDualPayeeAddr sets DualPayeeAddr field to given value.


### GetMemo

`func (o *DualAuthTransferRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *DualAuthTransferRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *DualAuthTransferRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *DualAuthTransferRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


