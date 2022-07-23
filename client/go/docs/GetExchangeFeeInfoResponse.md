# GetExchangeFeeInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**GetExchangeFeeInfoResponseData**](GetExchangeFeeInfoResponseData.md) |  | [optional] 

## Methods

### NewGetExchangeFeeInfoResponse

`func NewGetExchangeFeeInfoResponse(resultInfo ResultInfo, ) *GetExchangeFeeInfoResponse`

NewGetExchangeFeeInfoResponse instantiates a new GetExchangeFeeInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeFeeInfoResponseWithDefaults

`func NewGetExchangeFeeInfoResponseWithDefaults() *GetExchangeFeeInfoResponse`

NewGetExchangeFeeInfoResponseWithDefaults instantiates a new GetExchangeFeeInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetExchangeFeeInfoResponse) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetExchangeFeeInfoResponse) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetExchangeFeeInfoResponse) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetExchangeFeeInfoResponse) GetData() GetExchangeFeeInfoResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetExchangeFeeInfoResponse) GetDataOk() (*GetExchangeFeeInfoResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetExchangeFeeInfoResponse) SetData(v GetExchangeFeeInfoResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetExchangeFeeInfoResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


