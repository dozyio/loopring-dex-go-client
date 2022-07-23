# AmmPoolExitTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnPooled** | [**[]TokenVolumeV3**](TokenVolumeV3.md) | Describes list of the amount of a specific token to be removed from the pool, tokens seq should be same as AMM pool info | 
**Burned** | **string** | The minimum amoun of LP token to burn | 

## Methods

### NewAmmPoolExitTokens

`func NewAmmPoolExitTokens(unPooled []TokenVolumeV3, burned string, ) *AmmPoolExitTokens`

NewAmmPoolExitTokens instantiates a new AmmPoolExitTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmPoolExitTokensWithDefaults

`func NewAmmPoolExitTokensWithDefaults() *AmmPoolExitTokens`

NewAmmPoolExitTokensWithDefaults instantiates a new AmmPoolExitTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnPooled

`func (o *AmmPoolExitTokens) GetUnPooled() []TokenVolumeV3`

GetUnPooled returns the UnPooled field if non-nil, zero value otherwise.

### GetUnPooledOk

`func (o *AmmPoolExitTokens) GetUnPooledOk() (*[]TokenVolumeV3, bool)`

GetUnPooledOk returns a tuple with the UnPooled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnPooled

`func (o *AmmPoolExitTokens) SetUnPooled(v []TokenVolumeV3)`

SetUnPooled sets UnPooled field to given value.


### GetBurned

`func (o *AmmPoolExitTokens) GetBurned() string`

GetBurned returns the Burned field if non-nil, zero value otherwise.

### GetBurnedOk

`func (o *AmmPoolExitTokens) GetBurnedOk() (*string, bool)`

GetBurnedOk returns a tuple with the Burned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurned

`func (o *AmmPoolExitTokens) SetBurned(v string)`

SetBurned sets Burned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


