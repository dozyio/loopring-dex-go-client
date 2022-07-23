# GetUserFeeRatesResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeRate** | Pointer to [**FeeRate**](FeeRate.md) |  | [optional] 
**GasPrice** | Pointer to **string** | The gas price use to calculate fee rate | [optional] 

## Methods

### NewGetUserFeeRatesResponseData

`func NewGetUserFeeRatesResponseData() *GetUserFeeRatesResponseData`

NewGetUserFeeRatesResponseData instantiates a new GetUserFeeRatesResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserFeeRatesResponseDataWithDefaults

`func NewGetUserFeeRatesResponseDataWithDefaults() *GetUserFeeRatesResponseData`

NewGetUserFeeRatesResponseDataWithDefaults instantiates a new GetUserFeeRatesResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeRate

`func (o *GetUserFeeRatesResponseData) GetFeeRate() FeeRate`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *GetUserFeeRatesResponseData) GetFeeRateOk() (*FeeRate, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *GetUserFeeRatesResponseData) SetFeeRate(v FeeRate)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *GetUserFeeRatesResponseData) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetGasPrice

`func (o *GetUserFeeRatesResponseData) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetUserFeeRatesResponseData) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetUserFeeRatesResponseData) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *GetUserFeeRatesResponseData) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


