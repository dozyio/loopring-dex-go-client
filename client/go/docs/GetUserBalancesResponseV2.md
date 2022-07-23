# GetUserBalancesResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | [**[]Balance**](Balance.md) | field.Balance.balances | 

## Methods

### NewGetUserBalancesResponseV2

`func NewGetUserBalancesResponseV2(resultInfo ResultInfo, data []Balance, ) *GetUserBalancesResponseV2`

NewGetUserBalancesResponseV2 instantiates a new GetUserBalancesResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserBalancesResponseV2WithDefaults

`func NewGetUserBalancesResponseV2WithDefaults() *GetUserBalancesResponseV2`

NewGetUserBalancesResponseV2WithDefaults instantiates a new GetUserBalancesResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetUserBalancesResponseV2) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetUserBalancesResponseV2) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetUserBalancesResponseV2) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetUserBalancesResponseV2) GetData() []Balance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserBalancesResponseV2) GetDataOk() (*[]Balance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserBalancesResponseV2) SetData(v []Balance)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


