# GetAllowancesV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to **[]string** | Allowance amount in wei | [optional] 

## Methods

### NewGetAllowancesV2Response

`func NewGetAllowancesV2Response(resultInfo ResultInfo, ) *GetAllowancesV2Response`

NewGetAllowancesV2Response instantiates a new GetAllowancesV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllowancesV2ResponseWithDefaults

`func NewGetAllowancesV2ResponseWithDefaults() *GetAllowancesV2Response`

NewGetAllowancesV2ResponseWithDefaults instantiates a new GetAllowancesV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetAllowancesV2Response) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetAllowancesV2Response) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetAllowancesV2Response) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetAllowancesV2Response) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllowancesV2Response) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllowancesV2Response) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllowancesV2Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


