# GetUserNftOrderRatesAmountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NftTokenAddress** | **string** | field.NftFeeRate.nftTokenAddress | 
**FeeRate** | **int32** | field.NftFeeRate.rate | 
**Amounts** | [**[]FeeTokenAmount**](FeeTokenAmount.md) | Amounts | 
**GasPrice** | **string** | field.GetUserNftFeeRatesResponse.gasPrice | 
**CacheOverdueAt** | **int64** | Cached price data overdue time | 

## Methods

### NewGetUserNftOrderRatesAmountsResponse

`func NewGetUserNftOrderRatesAmountsResponse(nftTokenAddress string, feeRate int32, amounts []FeeTokenAmount, gasPrice string, cacheOverdueAt int64, ) *GetUserNftOrderRatesAmountsResponse`

NewGetUserNftOrderRatesAmountsResponse instantiates a new GetUserNftOrderRatesAmountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserNftOrderRatesAmountsResponseWithDefaults

`func NewGetUserNftOrderRatesAmountsResponseWithDefaults() *GetUserNftOrderRatesAmountsResponse`

NewGetUserNftOrderRatesAmountsResponseWithDefaults instantiates a new GetUserNftOrderRatesAmountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNftTokenAddress

`func (o *GetUserNftOrderRatesAmountsResponse) GetNftTokenAddress() string`

GetNftTokenAddress returns the NftTokenAddress field if non-nil, zero value otherwise.

### GetNftTokenAddressOk

`func (o *GetUserNftOrderRatesAmountsResponse) GetNftTokenAddressOk() (*string, bool)`

GetNftTokenAddressOk returns a tuple with the NftTokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftTokenAddress

`func (o *GetUserNftOrderRatesAmountsResponse) SetNftTokenAddress(v string)`

SetNftTokenAddress sets NftTokenAddress field to given value.


### GetFeeRate

`func (o *GetUserNftOrderRatesAmountsResponse) GetFeeRate() int32`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *GetUserNftOrderRatesAmountsResponse) GetFeeRateOk() (*int32, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *GetUserNftOrderRatesAmountsResponse) SetFeeRate(v int32)`

SetFeeRate sets FeeRate field to given value.


### GetAmounts

`func (o *GetUserNftOrderRatesAmountsResponse) GetAmounts() []FeeTokenAmount`

GetAmounts returns the Amounts field if non-nil, zero value otherwise.

### GetAmountsOk

`func (o *GetUserNftOrderRatesAmountsResponse) GetAmountsOk() (*[]FeeTokenAmount, bool)`

GetAmountsOk returns a tuple with the Amounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmounts

`func (o *GetUserNftOrderRatesAmountsResponse) SetAmounts(v []FeeTokenAmount)`

SetAmounts sets Amounts field to given value.


### GetGasPrice

`func (o *GetUserNftOrderRatesAmountsResponse) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetUserNftOrderRatesAmountsResponse) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetUserNftOrderRatesAmountsResponse) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetCacheOverdueAt

`func (o *GetUserNftOrderRatesAmountsResponse) GetCacheOverdueAt() int64`

GetCacheOverdueAt returns the CacheOverdueAt field if non-nil, zero value otherwise.

### GetCacheOverdueAtOk

`func (o *GetUserNftOrderRatesAmountsResponse) GetCacheOverdueAtOk() (*int64, bool)`

GetCacheOverdueAtOk returns a tuple with the CacheOverdueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheOverdueAt

`func (o *GetUserNftOrderRatesAmountsResponse) SetCacheOverdueAt(v int64)`

SetCacheOverdueAt sets CacheOverdueAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


