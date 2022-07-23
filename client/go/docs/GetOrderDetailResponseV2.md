# GetOrderDetailResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**OrderDetail**](OrderDetail.md) |  | [optional] 

## Methods

### NewGetOrderDetailResponseV2

`func NewGetOrderDetailResponseV2(resultInfo ResultInfo, ) *GetOrderDetailResponseV2`

NewGetOrderDetailResponseV2 instantiates a new GetOrderDetailResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderDetailResponseV2WithDefaults

`func NewGetOrderDetailResponseV2WithDefaults() *GetOrderDetailResponseV2`

NewGetOrderDetailResponseV2WithDefaults instantiates a new GetOrderDetailResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetOrderDetailResponseV2) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetOrderDetailResponseV2) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetOrderDetailResponseV2) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetOrderDetailResponseV2) GetData() OrderDetail`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOrderDetailResponseV2) GetDataOk() (*OrderDetail, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOrderDetailResponseV2) SetData(v OrderDetail)`

SetData sets Data field to given value.

### HasData

`func (o *GetOrderDetailResponseV2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


