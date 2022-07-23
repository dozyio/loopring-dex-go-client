# WithdrawLuckyTokenRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **int32** | field.WithdrawLuckyTokenRequestV3.withdrawTokenId | 
**FeeTokenId** | **int32** | field.WithdrawLuckyTokenRequestV3.feeTokenId | 
**Amount** | **string** | field.WithdrawLuckyTokenRequestV3.amount | 
**Claimer** | **string** | field.WithdrawLuckyTokenRequestV3.claimer | 
**Layer2Transfer** | Pointer to **string** | field.WithdrawLuckyTokenRequestV3.layer2Transfer | [optional] 

## Methods

### NewWithdrawLuckyTokenRequestV3

`func NewWithdrawLuckyTokenRequestV3(tokenId int32, feeTokenId int32, amount string, claimer string, ) *WithdrawLuckyTokenRequestV3`

NewWithdrawLuckyTokenRequestV3 instantiates a new WithdrawLuckyTokenRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawLuckyTokenRequestV3WithDefaults

`func NewWithdrawLuckyTokenRequestV3WithDefaults() *WithdrawLuckyTokenRequestV3`

NewWithdrawLuckyTokenRequestV3WithDefaults instantiates a new WithdrawLuckyTokenRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *WithdrawLuckyTokenRequestV3) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *WithdrawLuckyTokenRequestV3) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *WithdrawLuckyTokenRequestV3) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetFeeTokenId

`func (o *WithdrawLuckyTokenRequestV3) GetFeeTokenId() int32`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *WithdrawLuckyTokenRequestV3) GetFeeTokenIdOk() (*int32, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *WithdrawLuckyTokenRequestV3) SetFeeTokenId(v int32)`

SetFeeTokenId sets FeeTokenId field to given value.


### GetAmount

`func (o *WithdrawLuckyTokenRequestV3) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WithdrawLuckyTokenRequestV3) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WithdrawLuckyTokenRequestV3) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetClaimer

`func (o *WithdrawLuckyTokenRequestV3) GetClaimer() string`

GetClaimer returns the Claimer field if non-nil, zero value otherwise.

### GetClaimerOk

`func (o *WithdrawLuckyTokenRequestV3) GetClaimerOk() (*string, bool)`

GetClaimerOk returns a tuple with the Claimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimer

`func (o *WithdrawLuckyTokenRequestV3) SetClaimer(v string)`

SetClaimer sets Claimer field to given value.


### GetLayer2Transfer

`func (o *WithdrawLuckyTokenRequestV3) GetLayer2Transfer() string`

GetLayer2Transfer returns the Layer2Transfer field if non-nil, zero value otherwise.

### GetLayer2TransferOk

`func (o *WithdrawLuckyTokenRequestV3) GetLayer2TransferOk() (*string, bool)`

GetLayer2TransferOk returns a tuple with the Layer2Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer2Transfer

`func (o *WithdrawLuckyTokenRequestV3) SetLayer2Transfer(v string)`

SetLayer2Transfer sets Layer2Transfer field to given value.

### HasLayer2Transfer

`func (o *WithdrawLuckyTokenRequestV3) HasLayer2Transfer() bool`

HasLayer2Transfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


