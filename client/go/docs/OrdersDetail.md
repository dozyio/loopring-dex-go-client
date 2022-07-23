# OrdersDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | Total number of orders | 
**Orders** | Pointer to [**[]OrderDetail**](OrderDetail.md) | List of order | [optional] 

## Methods

### NewOrdersDetail

`func NewOrdersDetail(totalNum int64, ) *OrdersDetail`

NewOrdersDetail instantiates a new OrdersDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersDetailWithDefaults

`func NewOrdersDetailWithDefaults() *OrdersDetail`

NewOrdersDetailWithDefaults instantiates a new OrdersDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *OrdersDetail) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *OrdersDetail) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *OrdersDetail) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetOrders

`func (o *OrdersDetail) GetOrders() []OrderDetail`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OrdersDetail) GetOrdersOk() (*[]OrderDetail, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OrdersDetail) SetOrders(v []OrderDetail)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OrdersDetail) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


