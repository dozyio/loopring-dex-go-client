# CounterFactualInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletFactory** | **string** | Counter factual wallet factory contract address | 
**WalletOwner** | **string** | Counter factual wallet owner address, NOT the wallet address | 
**WalletSalt** | **string** | Salt to generate address from owner &amp; other related info | 

## Methods

### NewCounterFactualInfo

`func NewCounterFactualInfo(walletFactory string, walletOwner string, walletSalt string, ) *CounterFactualInfo`

NewCounterFactualInfo instantiates a new CounterFactualInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCounterFactualInfoWithDefaults

`func NewCounterFactualInfoWithDefaults() *CounterFactualInfo`

NewCounterFactualInfoWithDefaults instantiates a new CounterFactualInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletFactory

`func (o *CounterFactualInfo) GetWalletFactory() string`

GetWalletFactory returns the WalletFactory field if non-nil, zero value otherwise.

### GetWalletFactoryOk

`func (o *CounterFactualInfo) GetWalletFactoryOk() (*string, bool)`

GetWalletFactoryOk returns a tuple with the WalletFactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletFactory

`func (o *CounterFactualInfo) SetWalletFactory(v string)`

SetWalletFactory sets WalletFactory field to given value.


### GetWalletOwner

`func (o *CounterFactualInfo) GetWalletOwner() string`

GetWalletOwner returns the WalletOwner field if non-nil, zero value otherwise.

### GetWalletOwnerOk

`func (o *CounterFactualInfo) GetWalletOwnerOk() (*string, bool)`

GetWalletOwnerOk returns a tuple with the WalletOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletOwner

`func (o *CounterFactualInfo) SetWalletOwner(v string)`

SetWalletOwner sets WalletOwner field to given value.


### GetWalletSalt

`func (o *CounterFactualInfo) GetWalletSalt() string`

GetWalletSalt returns the WalletSalt field if non-nil, zero value otherwise.

### GetWalletSaltOk

`func (o *CounterFactualInfo) GetWalletSaltOk() (*string, bool)`

GetWalletSaltOk returns a tuple with the WalletSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletSalt

`func (o *CounterFactualInfo) SetWalletSalt(v string)`

SetWalletSalt sets WalletSalt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


