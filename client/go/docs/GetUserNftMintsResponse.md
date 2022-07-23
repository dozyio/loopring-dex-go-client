# GetUserNftMintsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | field.totalNum.description | 
**Mints** | [**[]NftMintData**](NftMintData.md) | field.GetUserNftMintsResponse.NftMintDataList | 

## Methods

### NewGetUserNftMintsResponse

`func NewGetUserNftMintsResponse(totalNum int64, mints []NftMintData, ) *GetUserNftMintsResponse`

NewGetUserNftMintsResponse instantiates a new GetUserNftMintsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserNftMintsResponseWithDefaults

`func NewGetUserNftMintsResponseWithDefaults() *GetUserNftMintsResponse`

NewGetUserNftMintsResponseWithDefaults instantiates a new GetUserNftMintsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *GetUserNftMintsResponse) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *GetUserNftMintsResponse) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *GetUserNftMintsResponse) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetMints

`func (o *GetUserNftMintsResponse) GetMints() []NftMintData`

GetMints returns the Mints field if non-nil, zero value otherwise.

### GetMintsOk

`func (o *GetUserNftMintsResponse) GetMintsOk() (*[]NftMintData, bool)`

GetMintsOk returns a tuple with the Mints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMints

`func (o *GetUserNftMintsResponse) SetMints(v []NftMintData)`

SetMints sets Mints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


