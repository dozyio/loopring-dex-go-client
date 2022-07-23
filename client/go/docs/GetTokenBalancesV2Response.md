# GetTokenBalancesV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to **[]string** | Balance in wei | [optional] 

## Methods

### NewGetTokenBalancesV2Response

`func NewGetTokenBalancesV2Response(resultInfo ResultInfo, ) *GetTokenBalancesV2Response`

NewGetTokenBalancesV2Response instantiates a new GetTokenBalancesV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTokenBalancesV2ResponseWithDefaults

`func NewGetTokenBalancesV2ResponseWithDefaults() *GetTokenBalancesV2Response`

NewGetTokenBalancesV2ResponseWithDefaults instantiates a new GetTokenBalancesV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetTokenBalancesV2Response) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetTokenBalancesV2Response) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetTokenBalancesV2Response) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetTokenBalancesV2Response) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTokenBalancesV2Response) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTokenBalancesV2Response) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *GetTokenBalancesV2Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


