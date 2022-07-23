# MinterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** |  | 
**Minter** | **string** |  | 
**OriginalMinter** | Pointer to **string** |  | [optional] 

## Methods

### NewMinterInfo

`func NewMinterInfo(accountId int64, minter string, ) *MinterInfo`

NewMinterInfo instantiates a new MinterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinterInfoWithDefaults

`func NewMinterInfoWithDefaults() *MinterInfo`

NewMinterInfoWithDefaults instantiates a new MinterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *MinterInfo) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MinterInfo) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MinterInfo) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetMinter

`func (o *MinterInfo) GetMinter() string`

GetMinter returns the Minter field if non-nil, zero value otherwise.

### GetMinterOk

`func (o *MinterInfo) GetMinterOk() (*string, bool)`

GetMinterOk returns a tuple with the Minter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinter

`func (o *MinterInfo) SetMinter(v string)`

SetMinter sets Minter field to given value.


### GetOriginalMinter

`func (o *MinterInfo) GetOriginalMinter() string`

GetOriginalMinter returns the OriginalMinter field if non-nil, zero value otherwise.

### GetOriginalMinterOk

`func (o *MinterInfo) GetOriginalMinterOk() (*string, bool)`

GetOriginalMinterOk returns a tuple with the OriginalMinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMinter

`func (o *MinterInfo) SetOriginalMinter(v string)`

SetOriginalMinter sets OriginalMinter field to given value.

### HasOriginalMinter

`func (o *MinterInfo) HasOriginalMinter() bool`

HasOriginalMinter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


