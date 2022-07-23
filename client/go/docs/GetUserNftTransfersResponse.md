# GetUserNftTransfersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | field.totalNum.description | 
**Transfers** | [**[]NftTransferData**](NftTransferData.md) | field.GetUserNftTransfersResponse.NftTransferDataList | 

## Methods

### NewGetUserNftTransfersResponse

`func NewGetUserNftTransfersResponse(totalNum int64, transfers []NftTransferData, ) *GetUserNftTransfersResponse`

NewGetUserNftTransfersResponse instantiates a new GetUserNftTransfersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserNftTransfersResponseWithDefaults

`func NewGetUserNftTransfersResponseWithDefaults() *GetUserNftTransfersResponse`

NewGetUserNftTransfersResponseWithDefaults instantiates a new GetUserNftTransfersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *GetUserNftTransfersResponse) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *GetUserNftTransfersResponse) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *GetUserNftTransfersResponse) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetTransfers

`func (o *GetUserNftTransfersResponse) GetTransfers() []NftTransferData`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *GetUserNftTransfersResponse) GetTransfersOk() (*[]NftTransferData, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *GetUserNftTransfersResponse) SetTransfers(v []NftTransferData)`

SetTransfers sets Transfers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


