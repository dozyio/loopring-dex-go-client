# GetAmmSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | [**AmmSnapshot**](AmmSnapshot.md) |  | 

## Methods

### NewGetAmmSnapshotResponse

`func NewGetAmmSnapshotResponse(resultInfo ResultInfo, data AmmSnapshot, ) *GetAmmSnapshotResponse`

NewGetAmmSnapshotResponse instantiates a new GetAmmSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAmmSnapshotResponseWithDefaults

`func NewGetAmmSnapshotResponseWithDefaults() *GetAmmSnapshotResponse`

NewGetAmmSnapshotResponseWithDefaults instantiates a new GetAmmSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetAmmSnapshotResponse) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetAmmSnapshotResponse) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetAmmSnapshotResponse) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetAmmSnapshotResponse) GetData() AmmSnapshot`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAmmSnapshotResponse) GetDataOk() (*AmmSnapshot, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAmmSnapshotResponse) SetData(v AmmSnapshot)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


