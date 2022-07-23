# GetOffchainFee2ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GasPrice** | **string** |  | 
**Fees** | [**[]OffFeeInfo2**](OffFeeInfo2.md) |  | 

## Methods

### NewGetOffchainFee2ResponseData

`func NewGetOffchainFee2ResponseData(gasPrice string, fees []OffFeeInfo2, ) *GetOffchainFee2ResponseData`

NewGetOffchainFee2ResponseData instantiates a new GetOffchainFee2ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOffchainFee2ResponseDataWithDefaults

`func NewGetOffchainFee2ResponseDataWithDefaults() *GetOffchainFee2ResponseData`

NewGetOffchainFee2ResponseDataWithDefaults instantiates a new GetOffchainFee2ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGasPrice

`func (o *GetOffchainFee2ResponseData) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GetOffchainFee2ResponseData) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GetOffchainFee2ResponseData) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetFees

`func (o *GetOffchainFee2ResponseData) GetFees() []OffFeeInfo2`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *GetOffchainFee2ResponseData) GetFeesOk() (*[]OffFeeInfo2, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *GetOffchainFee2ResponseData) SetFees(v []OffFeeInfo2)`

SetFees sets Fees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


