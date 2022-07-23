# SubmitOffChainRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**SubmitOffChainResponseItem**](SubmitOffChainResponseItem.md) |  | [optional] 

## Methods

### NewSubmitOffChainRequestResponse

`func NewSubmitOffChainRequestResponse(resultInfo ResultInfo, ) *SubmitOffChainRequestResponse`

NewSubmitOffChainRequestResponse instantiates a new SubmitOffChainRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitOffChainRequestResponseWithDefaults

`func NewSubmitOffChainRequestResponseWithDefaults() *SubmitOffChainRequestResponse`

NewSubmitOffChainRequestResponseWithDefaults instantiates a new SubmitOffChainRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *SubmitOffChainRequestResponse) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *SubmitOffChainRequestResponse) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *SubmitOffChainRequestResponse) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *SubmitOffChainRequestResponse) GetData() SubmitOffChainResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubmitOffChainRequestResponse) GetDataOk() (*SubmitOffChainResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubmitOffChainRequestResponse) SetData(v SubmitOffChainResponseItem)`

SetData sets Data field to given value.

### HasData

`func (o *SubmitOffChainRequestResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


