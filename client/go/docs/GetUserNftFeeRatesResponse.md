# GetUserNftFeeRatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeRate** | Pointer to [**NftFeeRate**](NftFeeRate.md) |  | [optional] 
**GasPrice** | Pointer to **string** | field.GetUserNftFeeRatesResponse.gasPrice | [optional] 

## Methods

### NewGetUserNftFeeRatesResponse

`func NewGetUserNftFeeRatesResponse() *GetUserNftFeeRatesResponse`

NewGetUserNftFeeRatesResponse instantiates a new GetUserNftFeeRatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserNftFeeRatesResponseWithDefaults

`func NewGetUserNftFeeRatesResponseWithDefaults() *GetUserNftFeeRatesResponse`

NewGetUserNftFeeRatesResponseWithDefaults instantiates a new GetUserNftFeeRatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeRate

`func (o *GetUserNftFeeRatesResponse) GetFeeRate() NftFeeRate`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *GetUserNftFeeRatesResponse) GetFeeRateOk() (*NftFeeRate, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *GetUserNftFeeRatesResponse) SetFeeRate(v NftFeeRate)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *GetUserNftFeeRatesResponse) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetGasPrice

`func (o *GetUserNftFeeRatesResponse) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetUserNftFeeRatesResponse) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetUserNftFeeRatesResponse) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *GetUserNftFeeRatesResponse) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


