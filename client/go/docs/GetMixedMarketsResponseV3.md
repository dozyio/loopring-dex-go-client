# GetMixedMarketsResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Markets** | [**[]CombineMarketInfo**](CombineMarketInfo.md) | Markets list | 

## Methods

### NewGetMixedMarketsResponseV3

`func NewGetMixedMarketsResponseV3(markets []CombineMarketInfo, ) *GetMixedMarketsResponseV3`

NewGetMixedMarketsResponseV3 instantiates a new GetMixedMarketsResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMixedMarketsResponseV3WithDefaults

`func NewGetMixedMarketsResponseV3WithDefaults() *GetMixedMarketsResponseV3`

NewGetMixedMarketsResponseV3WithDefaults instantiates a new GetMixedMarketsResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarkets

`func (o *GetMixedMarketsResponseV3) GetMarkets() []CombineMarketInfo`

GetMarkets returns the Markets field if non-nil, zero value otherwise.

### GetMarketsOk

`func (o *GetMixedMarketsResponseV3) GetMarketsOk() (*[]CombineMarketInfo, bool)`

GetMarketsOk returns a tuple with the Markets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkets

`func (o *GetMixedMarketsResponseV3) SetMarkets(v []CombineMarketInfo)`

SetMarkets sets Markets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


