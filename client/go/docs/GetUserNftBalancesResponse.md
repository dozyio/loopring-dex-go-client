# GetUserNftBalancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | field.totalNum.description | 
**Data** | [**[]CombinedNftBalance**](CombinedNftBalance.md) | NFT tokens balance info | 

## Methods

### NewGetUserNftBalancesResponse

`func NewGetUserNftBalancesResponse(totalNum int64, data []CombinedNftBalance, ) *GetUserNftBalancesResponse`

NewGetUserNftBalancesResponse instantiates a new GetUserNftBalancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserNftBalancesResponseWithDefaults

`func NewGetUserNftBalancesResponseWithDefaults() *GetUserNftBalancesResponse`

NewGetUserNftBalancesResponseWithDefaults instantiates a new GetUserNftBalancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *GetUserNftBalancesResponse) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *GetUserNftBalancesResponse) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *GetUserNftBalancesResponse) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetData

`func (o *GetUserNftBalancesResponse) GetData() []CombinedNftBalance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserNftBalancesResponse) GetDataOk() (*[]CombinedNftBalance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserNftBalancesResponse) SetData(v []CombinedNftBalance)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


