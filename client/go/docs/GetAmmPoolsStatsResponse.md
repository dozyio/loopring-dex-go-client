# GetAmmPoolsStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | [**[]AmmPoolStatistics**](AmmPoolStatistics.md) | model.AmmPoolStatistics | 

## Methods

### NewGetAmmPoolsStatsResponse

`func NewGetAmmPoolsStatsResponse(resultInfo ResultInfo, data []AmmPoolStatistics, ) *GetAmmPoolsStatsResponse`

NewGetAmmPoolsStatsResponse instantiates a new GetAmmPoolsStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAmmPoolsStatsResponseWithDefaults

`func NewGetAmmPoolsStatsResponseWithDefaults() *GetAmmPoolsStatsResponse`

NewGetAmmPoolsStatsResponseWithDefaults instantiates a new GetAmmPoolsStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetAmmPoolsStatsResponse) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetAmmPoolsStatsResponse) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetAmmPoolsStatsResponse) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetAmmPoolsStatsResponse) GetData() []AmmPoolStatistics`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAmmPoolsStatsResponse) GetDataOk() (*[]AmmPoolStatistics, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAmmPoolsStatsResponse) SetData(v []AmmPoolStatistics)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


