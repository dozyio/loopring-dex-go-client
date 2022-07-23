# GetAmmMarketInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**[]AmmMarketInfo**](AmmMarketInfo.md) | field.getAmmMarketInfoResponse.data | [optional] 

## Methods

### NewGetAmmMarketInfoResponse

`func NewGetAmmMarketInfoResponse(resultInfo ResultInfo, ) *GetAmmMarketInfoResponse`

NewGetAmmMarketInfoResponse instantiates a new GetAmmMarketInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAmmMarketInfoResponseWithDefaults

`func NewGetAmmMarketInfoResponseWithDefaults() *GetAmmMarketInfoResponse`

NewGetAmmMarketInfoResponseWithDefaults instantiates a new GetAmmMarketInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetAmmMarketInfoResponse) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetAmmMarketInfoResponse) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetAmmMarketInfoResponse) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetAmmMarketInfoResponse) GetData() []AmmMarketInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAmmMarketInfoResponse) GetDataOk() (*[]AmmMarketInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAmmMarketInfoResponse) SetData(v []AmmMarketInfo)`

SetData sets Data field to given value.

### HasData

`func (o *GetAmmMarketInfoResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


