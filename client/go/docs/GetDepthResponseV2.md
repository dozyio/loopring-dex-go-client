# GetDepthResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**Depth**](Depth.md) |  | [optional] 

## Methods

### NewGetDepthResponseV2

`func NewGetDepthResponseV2(resultInfo ResultInfo, ) *GetDepthResponseV2`

NewGetDepthResponseV2 instantiates a new GetDepthResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDepthResponseV2WithDefaults

`func NewGetDepthResponseV2WithDefaults() *GetDepthResponseV2`

NewGetDepthResponseV2WithDefaults instantiates a new GetDepthResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetDepthResponseV2) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetDepthResponseV2) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetDepthResponseV2) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetDepthResponseV2) GetData() Depth`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDepthResponseV2) GetDataOk() (*Depth, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDepthResponseV2) SetData(v Depth)`

SetData sets Data field to given value.

### HasData

`func (o *GetDepthResponseV2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


