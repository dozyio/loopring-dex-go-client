# GetUserCreateResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**UserCreateDataList**](UserCreateDataList.md) |  | [optional] 

## Methods

### NewGetUserCreateResponseV2

`func NewGetUserCreateResponseV2(resultInfo ResultInfo, ) *GetUserCreateResponseV2`

NewGetUserCreateResponseV2 instantiates a new GetUserCreateResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserCreateResponseV2WithDefaults

`func NewGetUserCreateResponseV2WithDefaults() *GetUserCreateResponseV2`

NewGetUserCreateResponseV2WithDefaults instantiates a new GetUserCreateResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetUserCreateResponseV2) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetUserCreateResponseV2) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetUserCreateResponseV2) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetUserCreateResponseV2) GetData() UserCreateDataList`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserCreateResponseV2) GetDataOk() (*UserCreateDataList, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserCreateResponseV2) SetData(v UserCreateDataList)`

SetData sets Data field to given value.

### HasData

`func (o *GetUserCreateResponseV2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


