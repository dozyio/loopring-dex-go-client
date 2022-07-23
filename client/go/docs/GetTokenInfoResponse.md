# GetTokenInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**[]TokenInfo**](TokenInfo.md) | Token list | [optional] 

## Methods

### NewGetTokenInfoResponse

`func NewGetTokenInfoResponse(resultInfo ResultInfo, ) *GetTokenInfoResponse`

NewGetTokenInfoResponse instantiates a new GetTokenInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTokenInfoResponseWithDefaults

`func NewGetTokenInfoResponseWithDefaults() *GetTokenInfoResponse`

NewGetTokenInfoResponseWithDefaults instantiates a new GetTokenInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetTokenInfoResponse) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetTokenInfoResponse) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetTokenInfoResponse) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetTokenInfoResponse) GetData() []TokenInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTokenInfoResponse) GetDataOk() (*[]TokenInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTokenInfoResponse) SetData(v []TokenInfo)`

SetData sets Data field to given value.

### HasData

`func (o *GetTokenInfoResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


