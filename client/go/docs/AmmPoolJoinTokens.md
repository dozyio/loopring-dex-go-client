# AmmPoolJoinTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pooled** | [**[]TokenVolumeV3**](TokenVolumeV3.md) | Describes an amount of a specific token to be supplied to the pool, the seq matters, and must be the same as the pool token list returned by pool info | 
**MinimumLp** | [**TokenVolumeV3**](TokenVolumeV3.md) |  | 

## Methods

### NewAmmPoolJoinTokens

`func NewAmmPoolJoinTokens(pooled []TokenVolumeV3, minimumLp TokenVolumeV3, ) *AmmPoolJoinTokens`

NewAmmPoolJoinTokens instantiates a new AmmPoolJoinTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmPoolJoinTokensWithDefaults

`func NewAmmPoolJoinTokensWithDefaults() *AmmPoolJoinTokens`

NewAmmPoolJoinTokensWithDefaults instantiates a new AmmPoolJoinTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPooled

`func (o *AmmPoolJoinTokens) GetPooled() []TokenVolumeV3`

GetPooled returns the Pooled field if non-nil, zero value otherwise.

### GetPooledOk

`func (o *AmmPoolJoinTokens) GetPooledOk() (*[]TokenVolumeV3, bool)`

GetPooledOk returns a tuple with the Pooled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooled

`func (o *AmmPoolJoinTokens) SetPooled(v []TokenVolumeV3)`

SetPooled sets Pooled field to given value.


### GetMinimumLp

`func (o *AmmPoolJoinTokens) GetMinimumLp() TokenVolumeV3`

GetMinimumLp returns the MinimumLp field if non-nil, zero value otherwise.

### GetMinimumLpOk

`func (o *AmmPoolJoinTokens) GetMinimumLpOk() (*TokenVolumeV3, bool)`

GetMinimumLpOk returns a tuple with the MinimumLp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumLp

`func (o *AmmPoolJoinTokens) SetMinimumLp(v TokenVolumeV3)`

SetMinimumLp sets MinimumLp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


