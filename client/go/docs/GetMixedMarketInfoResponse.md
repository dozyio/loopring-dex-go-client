# GetMixedMarketInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**[]CombineMarketInfo**](CombineMarketInfo.md) | field.getMarketInfoResponse.data | [optional] 

## Methods

### NewGetMixedMarketInfoResponse

`func NewGetMixedMarketInfoResponse(resultInfo ResultInfo, ) *GetMixedMarketInfoResponse`

NewGetMixedMarketInfoResponse instantiates a new GetMixedMarketInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMixedMarketInfoResponseWithDefaults

`func NewGetMixedMarketInfoResponseWithDefaults() *GetMixedMarketInfoResponse`

NewGetMixedMarketInfoResponseWithDefaults instantiates a new GetMixedMarketInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetMixedMarketInfoResponse) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetMixedMarketInfoResponse) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetMixedMarketInfoResponse) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetMixedMarketInfoResponse) GetData() []CombineMarketInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMixedMarketInfoResponse) GetDataOk() (*[]CombineMarketInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMixedMarketInfoResponse) SetData(v []CombineMarketInfo)`

SetData sets Data field to given value.

### HasData

`func (o *GetMixedMarketInfoResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


