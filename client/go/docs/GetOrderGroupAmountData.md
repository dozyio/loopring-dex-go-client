# GetOrderGroupAmountData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasPrice** | Pointer to **string** | The gas price use to calculate amount | [optional] 
**Amounts** | Pointer to [**[]TokenAmount**](TokenAmount.md) | Amounts | [optional] 
**CacheOverdueAt** | Pointer to **int64** | Cached price data overdue time | [optional] 

## Methods

### NewGetOrderGroupAmountData

`func NewGetOrderGroupAmountData() *GetOrderGroupAmountData`

NewGetOrderGroupAmountData instantiates a new GetOrderGroupAmountData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderGroupAmountDataWithDefaults

`func NewGetOrderGroupAmountDataWithDefaults() *GetOrderGroupAmountData`

NewGetOrderGroupAmountDataWithDefaults instantiates a new GetOrderGroupAmountData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasPrice

`func (o *GetOrderGroupAmountData) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetOrderGroupAmountData) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetOrderGroupAmountData) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *GetOrderGroupAmountData) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetAmounts

`func (o *GetOrderGroupAmountData) GetAmounts() []TokenAmount`

GetAmounts returns the Amounts field if non-nil, zero value otherwise.

### GetAmountsOk

`func (o *GetOrderGroupAmountData) GetAmountsOk() (*[]TokenAmount, bool)`

GetAmountsOk returns a tuple with the Amounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmounts

`func (o *GetOrderGroupAmountData) SetAmounts(v []TokenAmount)`

SetAmounts sets Amounts field to given value.

### HasAmounts

`func (o *GetOrderGroupAmountData) HasAmounts() bool`

HasAmounts returns a boolean if a field has been set.

### GetCacheOverdueAt

`func (o *GetOrderGroupAmountData) GetCacheOverdueAt() int64`

GetCacheOverdueAt returns the CacheOverdueAt field if non-nil, zero value otherwise.

### GetCacheOverdueAtOk

`func (o *GetOrderGroupAmountData) GetCacheOverdueAtOk() (*int64, bool)`

GetCacheOverdueAtOk returns a tuple with the CacheOverdueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheOverdueAt

`func (o *GetOrderGroupAmountData) SetCacheOverdueAt(v int64)`

SetCacheOverdueAt sets CacheOverdueAt field to given value.

### HasCacheOverdueAt

`func (o *GetOrderGroupAmountData) HasCacheOverdueAt() bool`

HasCacheOverdueAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


