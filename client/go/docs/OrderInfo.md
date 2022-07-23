# OrderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinAmount** | Pointer to **string** | The minimum amount | [optional] 
**MakerRate** | Pointer to **int32** | Maker rate | [optional] 
**TakerRate** | Pointer to **int32** | Taker rate | [optional] 

## Methods

### NewOrderInfo

`func NewOrderInfo() *OrderInfo`

NewOrderInfo instantiates a new OrderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderInfoWithDefaults

`func NewOrderInfoWithDefaults() *OrderInfo`

NewOrderInfoWithDefaults instantiates a new OrderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinAmount

`func (o *OrderInfo) GetMinAmount() string`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *OrderInfo) GetMinAmountOk() (*string, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *OrderInfo) SetMinAmount(v string)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *OrderInfo) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### GetMakerRate

`func (o *OrderInfo) GetMakerRate() int32`

GetMakerRate returns the MakerRate field if non-nil, zero value otherwise.

### GetMakerRateOk

`func (o *OrderInfo) GetMakerRateOk() (*int32, bool)`

GetMakerRateOk returns a tuple with the MakerRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerRate

`func (o *OrderInfo) SetMakerRate(v int32)`

SetMakerRate sets MakerRate field to given value.

### HasMakerRate

`func (o *OrderInfo) HasMakerRate() bool`

HasMakerRate returns a boolean if a field has been set.

### GetTakerRate

`func (o *OrderInfo) GetTakerRate() int32`

GetTakerRate returns the TakerRate field if non-nil, zero value otherwise.

### GetTakerRateOk

`func (o *OrderInfo) GetTakerRateOk() (*int32, bool)`

GetTakerRateOk returns a tuple with the TakerRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerRate

`func (o *OrderInfo) SetTakerRate(v int32)`

SetTakerRate sets TakerRate field to given value.

### HasTakerRate

`func (o *OrderInfo) HasTakerRate() bool`

HasTakerRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


