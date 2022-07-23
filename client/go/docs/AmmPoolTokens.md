# AmmPoolTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pooled** | **[]map[string]interface{}** | An array containing the unique identifiers of those tokens that are currently in the pool, seq matters as most requests ask for the same sequence. | 
**Lp** | **int32** | The unique identifier of the pool-specific LP token. This token is minted when supplying liquidity to the pool, and represents liquidity stakes in it | 

## Methods

### NewAmmPoolTokens

`func NewAmmPoolTokens(pooled []map[string]interface{}, lp int32, ) *AmmPoolTokens`

NewAmmPoolTokens instantiates a new AmmPoolTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmPoolTokensWithDefaults

`func NewAmmPoolTokensWithDefaults() *AmmPoolTokens`

NewAmmPoolTokensWithDefaults instantiates a new AmmPoolTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPooled

`func (o *AmmPoolTokens) GetPooled() []map[string]interface{}`

GetPooled returns the Pooled field if non-nil, zero value otherwise.

### GetPooledOk

`func (o *AmmPoolTokens) GetPooledOk() (*[]map[string]interface{}, bool)`

GetPooledOk returns a tuple with the Pooled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPooled

`func (o *AmmPoolTokens) SetPooled(v []map[string]interface{})`

SetPooled sets Pooled field to given value.


### GetLp

`func (o *AmmPoolTokens) GetLp() int32`

GetLp returns the Lp field if non-nil, zero value otherwise.

### GetLpOk

`func (o *AmmPoolTokens) GetLpOk() (*int32, bool)`

GetLpOk returns a tuple with the Lp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLp

`func (o *AmmPoolTokens) SetLp(v int32)`

SetLp sets Lp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


