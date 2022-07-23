# GetOffchainFee2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**GetOffchainFee2ResponseData**](GetOffchainFee2ResponseData.md) |  | [optional] 

## Methods

### NewGetOffchainFee2Response

`func NewGetOffchainFee2Response(resultInfo ResultInfo, ) *GetOffchainFee2Response`

NewGetOffchainFee2Response instantiates a new GetOffchainFee2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOffchainFee2ResponseWithDefaults

`func NewGetOffchainFee2ResponseWithDefaults() *GetOffchainFee2Response`

NewGetOffchainFee2ResponseWithDefaults instantiates a new GetOffchainFee2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetOffchainFee2Response) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetOffchainFee2Response) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetOffchainFee2Response) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetOffchainFee2Response) GetData() GetOffchainFee2ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOffchainFee2Response) GetDataOk() (*GetOffchainFee2ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOffchainFee2Response) SetData(v GetOffchainFee2ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetOffchainFee2Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


