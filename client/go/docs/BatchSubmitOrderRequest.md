# BatchSubmitOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | [**[]SubmitOrderRequest**](SubmitOrderRequest.md) | Orders to submit. | 

## Methods

### NewBatchSubmitOrderRequest

`func NewBatchSubmitOrderRequest(orders []SubmitOrderRequest, ) *BatchSubmitOrderRequest`

NewBatchSubmitOrderRequest instantiates a new BatchSubmitOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSubmitOrderRequestWithDefaults

`func NewBatchSubmitOrderRequestWithDefaults() *BatchSubmitOrderRequest`

NewBatchSubmitOrderRequestWithDefaults instantiates a new BatchSubmitOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *BatchSubmitOrderRequest) GetOrders() []SubmitOrderRequest`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *BatchSubmitOrderRequest) GetOrdersOk() (*[]SubmitOrderRequest, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *BatchSubmitOrderRequest) SetOrders(v []SubmitOrderRequest)`

SetOrders sets Orders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


