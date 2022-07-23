# GetUserAccountUpdateResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**UserAccountUpdateDataList**](UserAccountUpdateDataList.md) |  | [optional] 

## Methods

### NewGetUserAccountUpdateResponseV2

`func NewGetUserAccountUpdateResponseV2(resultInfo ResultInfo, ) *GetUserAccountUpdateResponseV2`

NewGetUserAccountUpdateResponseV2 instantiates a new GetUserAccountUpdateResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserAccountUpdateResponseV2WithDefaults

`func NewGetUserAccountUpdateResponseV2WithDefaults() *GetUserAccountUpdateResponseV2`

NewGetUserAccountUpdateResponseV2WithDefaults instantiates a new GetUserAccountUpdateResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetUserAccountUpdateResponseV2) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetUserAccountUpdateResponseV2) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetUserAccountUpdateResponseV2) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetUserAccountUpdateResponseV2) GetData() UserAccountUpdateDataList`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserAccountUpdateResponseV2) GetDataOk() (*UserAccountUpdateDataList, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserAccountUpdateResponseV2) SetData(v UserAccountUpdateDataList)`

SetData sets Data field to given value.

### HasData

`func (o *GetUserAccountUpdateResponseV2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


