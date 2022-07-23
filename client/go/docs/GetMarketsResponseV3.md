# GetMarketsResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Markets** | [**[]MarketInfo**](MarketInfo.md) | Markets list | 

## Methods

### NewGetMarketsResponseV3

`func NewGetMarketsResponseV3(markets []MarketInfo, ) *GetMarketsResponseV3`

NewGetMarketsResponseV3 instantiates a new GetMarketsResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketsResponseV3WithDefaults

`func NewGetMarketsResponseV3WithDefaults() *GetMarketsResponseV3`

NewGetMarketsResponseV3WithDefaults instantiates a new GetMarketsResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarkets

`func (o *GetMarketsResponseV3) GetMarkets() []MarketInfo`

GetMarkets returns the Markets field if non-nil, zero value otherwise.

### GetMarketsOk

`func (o *GetMarketsResponseV3) GetMarketsOk() (*[]MarketInfo, bool)`

GetMarketsOk returns a tuple with the Markets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkets

`func (o *GetMarketsResponseV3) SetMarkets(v []MarketInfo)`

SetMarkets sets Markets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


