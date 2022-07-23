# CancelOrdersResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to **bool** | field.CancelOrdersResponseData.data | [optional] 

## Methods

### NewCancelOrdersResponseData

`func NewCancelOrdersResponseData(resultInfo ResultInfo, ) *CancelOrdersResponseData`

NewCancelOrdersResponseData instantiates a new CancelOrdersResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelOrdersResponseDataWithDefaults

`func NewCancelOrdersResponseDataWithDefaults() *CancelOrdersResponseData`

NewCancelOrdersResponseDataWithDefaults instantiates a new CancelOrdersResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *CancelOrdersResponseData) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *CancelOrdersResponseData) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *CancelOrdersResponseData) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *CancelOrdersResponseData) GetData() bool`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CancelOrdersResponseData) GetDataOk() (*bool, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CancelOrdersResponseData) SetData(v bool)`

SetData sets Data field to given value.

### HasData

`func (o *CancelOrdersResponseData) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


