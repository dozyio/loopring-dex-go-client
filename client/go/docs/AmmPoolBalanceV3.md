# AmmPoolBalanceV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolName** | **string** | field.AmmSnapshot.poolName | 
**PoolAddress** | **string** | field.AmmSnapshot.poolAddress | 
**Pooled** | [**[]TokenVolumeV3**](TokenVolumeV3.md) | AMM in pool tokens balances | 
**Lp** | [**TokenVolumeV3**](TokenVolumeV3.md) |  | 
**Risky** | **bool** | AMM pool risky flag, true if AMM pool TVL is low which means big slippage. | 

## Methods

### NewAmmPoolBalanceV3

`func NewAmmPoolBalanceV3(poolName string, poolAddress string, pooled []TokenVolumeV3, lp TokenVolumeV3, risky bool, ) *AmmPoolBalanceV3`

NewAmmPoolBalanceV3 instantiates a new AmmPoolBalanceV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmPoolBalanceV3WithDefaults

`func NewAmmPoolBalanceV3WithDefaults() *AmmPoolBalanceV3`

NewAmmPoolBalanceV3WithDefaults instantiates a new AmmPoolBalanceV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolName

`func (o *AmmPoolBalanceV3) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AmmPoolBalanceV3) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AmmPoolBalanceV3) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.


### GetPoolAddress

`func (o *AmmPoolBalanceV3) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *AmmPoolBalanceV3) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *AmmPoolBalanceV3) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetPooled

`func (o *AmmPoolBalanceV3) GetPooled() []TokenVolumeV3`

GetPooled returns the Pooled field if non-nil, zero value otherwise.

### GetPooledOk

`func (o *AmmPoolBalanceV3) GetPooledOk() (*[]TokenVolumeV3, bool)`

GetPooledOk returns a tuple with the Pooled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooled

`func (o *AmmPoolBalanceV3) SetPooled(v []TokenVolumeV3)`

SetPooled sets Pooled field to given value.


### GetLp

`func (o *AmmPoolBalanceV3) GetLp() TokenVolumeV3`

GetLp returns the Lp field if non-nil, zero value otherwise.

### GetLpOk

`func (o *AmmPoolBalanceV3) GetLpOk() (*TokenVolumeV3, bool)`

GetLpOk returns a tuple with the Lp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLp

`func (o *AmmPoolBalanceV3) SetLp(v TokenVolumeV3)`

SetLp sets Lp field to given value.


### GetRisky

`func (o *AmmPoolBalanceV3) GetRisky() bool`

GetRisky returns the Risky field if non-nil, zero value otherwise.

### GetRiskyOk

`func (o *AmmPoolBalanceV3) GetRiskyOk() (*bool, bool)`

GetRiskyOk returns a tuple with the Risky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisky

`func (o *AmmPoolBalanceV3) SetRisky(v bool)`

SetRisky sets Risky field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


