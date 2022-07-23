# GetUserTransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**TransferDataList**](TransferDataList.md) |  | [optional] 

## Methods

### NewGetUserTransferResponse

`func NewGetUserTransferResponse(resultInfo ResultInfo, ) *GetUserTransferResponse`

NewGetUserTransferResponse instantiates a new GetUserTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserTransferResponseWithDefaults

`func NewGetUserTransferResponseWithDefaults() *GetUserTransferResponse`

NewGetUserTransferResponseWithDefaults instantiates a new GetUserTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetUserTransferResponse) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetUserTransferResponse) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetUserTransferResponse) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetUserTransferResponse) GetData() TransferDataList`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserTransferResponse) GetDataOk() (*TransferDataList, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserTransferResponse) SetData(v TransferDataList)`

SetData sets Data field to given value.

### HasData

`func (o *GetUserTransferResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


