# GetAmmPoolsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pools** | [**[]AmmPoolInfoV3**](AmmPoolInfoV3.md) | AMM pool list | 

## Methods

### NewGetAmmPoolsResponse

`func NewGetAmmPoolsResponse(pools []AmmPoolInfoV3, ) *GetAmmPoolsResponse`

NewGetAmmPoolsResponse instantiates a new GetAmmPoolsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAmmPoolsResponseWithDefaults

`func NewGetAmmPoolsResponseWithDefaults() *GetAmmPoolsResponse`

NewGetAmmPoolsResponseWithDefaults instantiates a new GetAmmPoolsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPools

`func (o *GetAmmPoolsResponse) GetPools() []AmmPoolInfoV3`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *GetAmmPoolsResponse) GetPoolsOk() (*[]AmmPoolInfoV3, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *GetAmmPoolsResponse) SetPools(v []AmmPoolInfoV3)`

SetPools sets Pools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


