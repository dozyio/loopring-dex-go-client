# GetEthNonceV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to **int64** | The result of nonce | [optional] 

## Methods

### NewGetEthNonceV2Response

`func NewGetEthNonceV2Response(resultInfo ResultInfo, ) *GetEthNonceV2Response`

NewGetEthNonceV2Response instantiates a new GetEthNonceV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthNonceV2ResponseWithDefaults

`func NewGetEthNonceV2ResponseWithDefaults() *GetEthNonceV2Response`

NewGetEthNonceV2ResponseWithDefaults instantiates a new GetEthNonceV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetEthNonceV2Response) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetEthNonceV2Response) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetEthNonceV2Response) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetEthNonceV2Response) GetData() int64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEthNonceV2Response) GetDataOk() (*int64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEthNonceV2Response) SetData(v int64)`

SetData sets Data field to given value.

### HasData

`func (o *GetEthNonceV2Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


