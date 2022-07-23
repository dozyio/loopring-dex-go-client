# GetRecommendedGasPriceResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to **string** | Response value of current recommended gas price in Gwei | [optional] 

## Methods

### NewGetRecommendedGasPriceResponseV2

`func NewGetRecommendedGasPriceResponseV2(resultInfo ResultInfo, ) *GetRecommendedGasPriceResponseV2`

NewGetRecommendedGasPriceResponseV2 instantiates a new GetRecommendedGasPriceResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecommendedGasPriceResponseV2WithDefaults

`func NewGetRecommendedGasPriceResponseV2WithDefaults() *GetRecommendedGasPriceResponseV2`

NewGetRecommendedGasPriceResponseV2WithDefaults instantiates a new GetRecommendedGasPriceResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetRecommendedGasPriceResponseV2) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetRecommendedGasPriceResponseV2) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetRecommendedGasPriceResponseV2) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetRecommendedGasPriceResponseV2) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRecommendedGasPriceResponseV2) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRecommendedGasPriceResponseV2) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *GetRecommendedGasPriceResponseV2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


