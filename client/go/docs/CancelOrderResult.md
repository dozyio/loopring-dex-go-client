# CancelOrderResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Hash or clientOrderId of order cancelled. | 
**Result** | **bool** | Order cancellation result. | 
**Error** | Pointer to [**ResultInfo**](ResultInfo.md) |  | [optional] 

## Methods

### NewCancelOrderResult

`func NewCancelOrderResult(id string, result bool, ) *CancelOrderResult`

NewCancelOrderResult instantiates a new CancelOrderResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelOrderResultWithDefaults

`func NewCancelOrderResultWithDefaults() *CancelOrderResult`

NewCancelOrderResultWithDefaults instantiates a new CancelOrderResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CancelOrderResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CancelOrderResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CancelOrderResult) SetId(v string)`

SetId sets Id field to given value.


### GetResult

`func (o *CancelOrderResult) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CancelOrderResult) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CancelOrderResult) SetResult(v bool)`

SetResult sets Result field to given value.


### GetError

`func (o *CancelOrderResult) GetError() ResultInfo`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CancelOrderResult) GetErrorOk() (*ResultInfo, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CancelOrderResult) SetError(v ResultInfo)`

SetError sets Error field to given value.

### HasError

`func (o *CancelOrderResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


