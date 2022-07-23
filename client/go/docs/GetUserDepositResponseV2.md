# GetUserDepositResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**DepositDataList**](DepositDataList.md) |  | [optional] 

## Methods

### NewGetUserDepositResponseV2

`func NewGetUserDepositResponseV2(resultInfo ResultInfo, ) *GetUserDepositResponseV2`

NewGetUserDepositResponseV2 instantiates a new GetUserDepositResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserDepositResponseV2WithDefaults

`func NewGetUserDepositResponseV2WithDefaults() *GetUserDepositResponseV2`

NewGetUserDepositResponseV2WithDefaults instantiates a new GetUserDepositResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetUserDepositResponseV2) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetUserDepositResponseV2) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetUserDepositResponseV2) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetUserDepositResponseV2) GetData() DepositDataList`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserDepositResponseV2) GetDataOk() (*DepositDataList, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserDepositResponseV2) SetData(v DepositDataList)`

SetData sets Data field to given value.

### HasData

`func (o *GetUserDepositResponseV2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


