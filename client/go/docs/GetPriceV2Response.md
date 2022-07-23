# GetPriceV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to [**[]Price**](Price.md) | field.getPriceResponse.data | [optional] 

## Methods

### NewGetPriceV2Response

`func NewGetPriceV2Response(resultInfo ResultInfo, ) *GetPriceV2Response`

NewGetPriceV2Response instantiates a new GetPriceV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPriceV2ResponseWithDefaults

`func NewGetPriceV2ResponseWithDefaults() *GetPriceV2Response`

NewGetPriceV2ResponseWithDefaults instantiates a new GetPriceV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetPriceV2Response) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetPriceV2Response) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetPriceV2Response) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetPriceV2Response) GetData() []Price`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPriceV2Response) GetDataOk() (*[]Price, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPriceV2Response) SetData(v []Price)`

SetData sets Data field to given value.

### HasData

`func (o *GetPriceV2Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


