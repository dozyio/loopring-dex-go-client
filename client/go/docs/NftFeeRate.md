# NftFeeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NftTokenAddress** | **string** | field.NftFeeRate.nftTokenAddress | 
**Rate** | **int32** | field.NftFeeRate.rate | 

## Methods

### NewNftFeeRate

`func NewNftFeeRate(nftTokenAddress string, rate int32, ) *NftFeeRate`

NewNftFeeRate instantiates a new NftFeeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftFeeRateWithDefaults

`func NewNftFeeRateWithDefaults() *NftFeeRate`

NewNftFeeRateWithDefaults instantiates a new NftFeeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNftTokenAddress

`func (o *NftFeeRate) GetNftTokenAddress() string`

GetNftTokenAddress returns the NftTokenAddress field if non-nil, zero value otherwise.

### GetNftTokenAddressOk

`func (o *NftFeeRate) GetNftTokenAddressOk() (*string, bool)`

GetNftTokenAddressOk returns a tuple with the NftTokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftTokenAddress

`func (o *NftFeeRate) SetNftTokenAddress(v string)`

SetNftTokenAddress sets NftTokenAddress field to given value.


### GetRate

`func (o *NftFeeRate) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *NftFeeRate) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *NftFeeRate) SetRate(v int32)`

SetRate sets Rate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


