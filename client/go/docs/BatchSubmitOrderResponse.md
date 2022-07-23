# BatchSubmitOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | [**[]SubmitOrderResult**](SubmitOrderResult.md) | Result of batch submit orders. | 

## Methods

### NewBatchSubmitOrderResponse

`func NewBatchSubmitOrderResponse(resultInfo ResultInfo, data []SubmitOrderResult, ) *BatchSubmitOrderResponse`

NewBatchSubmitOrderResponse instantiates a new BatchSubmitOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSubmitOrderResponseWithDefaults

`func NewBatchSubmitOrderResponseWithDefaults() *BatchSubmitOrderResponse`

NewBatchSubmitOrderResponseWithDefaults instantiates a new BatchSubmitOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *BatchSubmitOrderResponse) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *BatchSubmitOrderResponse) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *BatchSubmitOrderResponse) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *BatchSubmitOrderResponse) GetData() []SubmitOrderResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BatchSubmitOrderResponse) GetDataOk() (*[]SubmitOrderResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BatchSubmitOrderResponse) SetData(v []SubmitOrderResult)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


