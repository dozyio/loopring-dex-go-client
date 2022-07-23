# FeeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Market | 
**MakerRate** | **int32** | Maker rate | 
**TakerRate** | **int32** | Taker rate | 

## Methods

### NewFeeRate

`func NewFeeRate(symbol string, makerRate int32, takerRate int32, ) *FeeRate`

NewFeeRate instantiates a new FeeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeRateWithDefaults

`func NewFeeRateWithDefaults() *FeeRate`

NewFeeRateWithDefaults instantiates a new FeeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *FeeRate) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FeeRate) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FeeRate) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetMakerRate

`func (o *FeeRate) GetMakerRate() int32`

GetMakerRate returns the MakerRate field if non-nil, zero value otherwise.

### GetMakerRateOk

`func (o *FeeRate) GetMakerRateOk() (*int32, bool)`

GetMakerRateOk returns a tuple with the MakerRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerRate

`func (o *FeeRate) SetMakerRate(v int32)`

SetMakerRate sets MakerRate field to given value.


### GetTakerRate

`func (o *FeeRate) GetTakerRate() int32`

GetTakerRate returns the TakerRate field if non-nil, zero value otherwise.

### GetTakerRateOk

`func (o *FeeRate) GetTakerRateOk() (*int32, bool)`

GetTakerRateOk returns a tuple with the TakerRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerRate

`func (o *FeeRate) SetTakerRate(v int32)`

SetTakerRate sets TakerRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


