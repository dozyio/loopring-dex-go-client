# GetOrdersResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**OrdersDetail**](OrdersDetail.md) |  | [optional] 

## Methods

### NewGetOrdersResponseV2

`func NewGetOrdersResponseV2(resultInfo ResultInfo, ) *GetOrdersResponseV2`

NewGetOrdersResponseV2 instantiates a new GetOrdersResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrdersResponseV2WithDefaults

`func NewGetOrdersResponseV2WithDefaults() *GetOrdersResponseV2`

NewGetOrdersResponseV2WithDefaults instantiates a new GetOrdersResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetOrdersResponseV2) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetOrdersResponseV2) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetOrdersResponseV2) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetOrdersResponseV2) GetData() OrdersDetail`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOrdersResponseV2) GetDataOk() (*OrdersDetail, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOrdersResponseV2) SetData(v OrdersDetail)`

SetData sets Data field to given value.

### HasData

`func (o *GetOrdersResponseV2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


