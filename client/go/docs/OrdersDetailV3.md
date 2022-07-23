# OrdersDetailV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | Total number of orders | 
**Orders** | Pointer to [**[]OrderDetailV3**](OrderDetailV3.md) | List of order | [optional] 

## Methods

### NewOrdersDetailV3

`func NewOrdersDetailV3(totalNum int64, ) *OrdersDetailV3`

NewOrdersDetailV3 instantiates a new OrdersDetailV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersDetailV3WithDefaults

`func NewOrdersDetailV3WithDefaults() *OrdersDetailV3`

NewOrdersDetailV3WithDefaults instantiates a new OrdersDetailV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *OrdersDetailV3) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *OrdersDetailV3) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *OrdersDetailV3) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetOrders

`func (o *OrdersDetailV3) GetOrders() []OrderDetailV3`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *OrdersDetailV3) GetOrdersOk() (*[]OrderDetailV3, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *OrdersDetailV3) SetOrders(v []OrderDetailV3)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *OrdersDetailV3) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


