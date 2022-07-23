# Balance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | field.Balance.accountId | 
**TokenId** | **int32** | field.Balance.tokenId | 
**TotalAmount** | **string** | field.Balance.totalAmount | 
**AmountLocked** | **string** | field.Balance.frozenAmount | 
**Pending** | [**PendingBalance**](PendingBalance.md) |  | 

## Methods

### NewBalance

`func NewBalance(accountId int64, tokenId int32, totalAmount string, amountLocked string, pending PendingBalance, ) *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Balance) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Balance) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Balance) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetTokenId

`func (o *Balance) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Balance) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Balance) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetTotalAmount

`func (o *Balance) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *Balance) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *Balance) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetAmountLocked

`func (o *Balance) GetAmountLocked() string`

GetAmountLocked returns the AmountLocked field if non-nil, zero value otherwise.

### GetAmountLockedOk

`func (o *Balance) GetAmountLockedOk() (*string, bool)`

GetAmountLockedOk returns a tuple with the AmountLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLocked

`func (o *Balance) SetAmountLocked(v string)`

SetAmountLocked sets AmountLocked field to given value.


### GetPending

`func (o *Balance) GetPending() PendingBalance`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *Balance) GetPendingOk() (*PendingBalance, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *Balance) SetPending(v PendingBalance)`

SetPending sets Pending field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


