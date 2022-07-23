# BalanceV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **int32** | Token ID | 
**Total** | **string** | Amount of the asset | 
**Locked** | **string** | The part of the total balance which is currently not liquid and not at the users disposal (because of pending withdrawals or orders for example) | 
**Pending** | [**PendingBalance**](PendingBalance.md) |  | 

## Methods

### NewBalanceV3

`func NewBalanceV3(tokenId int32, total string, locked string, pending PendingBalance, ) *BalanceV3`

NewBalanceV3 instantiates a new BalanceV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceV3WithDefaults

`func NewBalanceV3WithDefaults() *BalanceV3`

NewBalanceV3WithDefaults instantiates a new BalanceV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *BalanceV3) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *BalanceV3) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *BalanceV3) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetTotal

`func (o *BalanceV3) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BalanceV3) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BalanceV3) SetTotal(v string)`

SetTotal sets Total field to given value.


### GetLocked

`func (o *BalanceV3) GetLocked() string`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *BalanceV3) GetLockedOk() (*string, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *BalanceV3) SetLocked(v string)`

SetLocked sets Locked field to given value.


### GetPending

`func (o *BalanceV3) GetPending() PendingBalance`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *BalanceV3) GetPendingOk() (*PendingBalance, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *BalanceV3) SetPending(v PendingBalance)`

SetPending sets Pending field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


