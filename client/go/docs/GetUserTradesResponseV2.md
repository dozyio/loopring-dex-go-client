# GetUserTradesResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**TradeList**](TradeList.md) |  | [optional] 

## Methods

### NewGetUserTradesResponseV2

`func NewGetUserTradesResponseV2(resultInfo ResultInfo, ) *GetUserTradesResponseV2`

NewGetUserTradesResponseV2 instantiates a new GetUserTradesResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserTradesResponseV2WithDefaults

`func NewGetUserTradesResponseV2WithDefaults() *GetUserTradesResponseV2`

NewGetUserTradesResponseV2WithDefaults instantiates a new GetUserTradesResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetUserTradesResponseV2) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetUserTradesResponseV2) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetUserTradesResponseV2) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetUserTradesResponseV2) GetData() TradeList`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserTradesResponseV2) GetDataOk() (*TradeList, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserTradesResponseV2) SetData(v TradeList)`

SetData sets Data field to given value.

### HasData

`func (o *GetUserTradesResponseV2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


