# AmmPoolPrecisions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **int32** | The price precision requirement of a AMM pool | 
**Amount** | **int32** | The amount precision requirement of a AMM pool | 

## Methods

### NewAmmPoolPrecisions

`func NewAmmPoolPrecisions(price int32, amount int32, ) *AmmPoolPrecisions`

NewAmmPoolPrecisions instantiates a new AmmPoolPrecisions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmPoolPrecisionsWithDefaults

`func NewAmmPoolPrecisionsWithDefaults() *AmmPoolPrecisions`

NewAmmPoolPrecisionsWithDefaults instantiates a new AmmPoolPrecisions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *AmmPoolPrecisions) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AmmPoolPrecisions) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AmmPoolPrecisions) SetPrice(v int32)`

SetPrice sets Price field to given value.


### GetAmount

`func (o *AmmPoolPrecisions) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AmmPoolPrecisions) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AmmPoolPrecisions) SetAmount(v int32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


