# PendingBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Withdraw** | **string** | Withdrawal pending balance which means the token is in withdrawal state but not arrived L1 | 
**Deposit** | **string** | Deposit pending balance which means the token is in deposit state but not arrived L2 | 

## Methods

### NewPendingBalance

`func NewPendingBalance(withdraw string, deposit string, ) *PendingBalance`

NewPendingBalance instantiates a new PendingBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingBalanceWithDefaults

`func NewPendingBalanceWithDefaults() *PendingBalance`

NewPendingBalanceWithDefaults instantiates a new PendingBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWithdraw

`func (o *PendingBalance) GetWithdraw() string`

GetWithdraw returns the Withdraw field if non-nil, zero value otherwise.

### GetWithdrawOk

`func (o *PendingBalance) GetWithdrawOk() (*string, bool)`

GetWithdrawOk returns a tuple with the Withdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdraw

`func (o *PendingBalance) SetWithdraw(v string)`

SetWithdraw sets Withdraw field to given value.


### GetDeposit

`func (o *PendingBalance) GetDeposit() string`

GetDeposit returns the Deposit field if non-nil, zero value otherwise.

### GetDepositOk

`func (o *PendingBalance) GetDepositOk() (*string, bool)`

GetDepositOk returns a tuple with the Deposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposit

`func (o *PendingBalance) SetDeposit(v string)`

SetDeposit sets Deposit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


