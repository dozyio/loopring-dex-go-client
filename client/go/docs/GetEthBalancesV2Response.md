# GetEthBalancesV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to **string** | Balance in wei | [optional] 

## Methods

### NewGetEthBalancesV2Response

`func NewGetEthBalancesV2Response(resultInfo ResultInfo, ) *GetEthBalancesV2Response`

NewGetEthBalancesV2Response instantiates a new GetEthBalancesV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthBalancesV2ResponseWithDefaults

`func NewGetEthBalancesV2ResponseWithDefaults() *GetEthBalancesV2Response`

NewGetEthBalancesV2ResponseWithDefaults instantiates a new GetEthBalancesV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetEthBalancesV2Response) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetEthBalancesV2Response) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetEthBalancesV2Response) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetEthBalancesV2Response) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEthBalancesV2Response) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEthBalancesV2Response) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *GetEthBalancesV2Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


