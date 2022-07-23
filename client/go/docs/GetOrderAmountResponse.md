# GetOrderAmountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**GetOrderAmountResponseData**](GetOrderAmountResponseData.md) |  | [optional] 

## Methods

### NewGetOrderAmountResponse

`func NewGetOrderAmountResponse(resultInfo ResultInfo, ) *GetOrderAmountResponse`

NewGetOrderAmountResponse instantiates a new GetOrderAmountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderAmountResponseWithDefaults

`func NewGetOrderAmountResponseWithDefaults() *GetOrderAmountResponse`

NewGetOrderAmountResponseWithDefaults instantiates a new GetOrderAmountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetOrderAmountResponse) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetOrderAmountResponse) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetOrderAmountResponse) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetOrderAmountResponse) GetData() GetOrderAmountResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOrderAmountResponse) GetDataOk() (*GetOrderAmountResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOrderAmountResponse) SetData(v GetOrderAmountResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetOrderAmountResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


