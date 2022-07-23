# SpotTradeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubmitterId** | Pointer to **int64** | field.spotTradeInfo.submitterId | [optional] 
**SmallOrder** | Pointer to **map[string]interface{}** | field.spotTradeInfo.smallOrder | [optional] 
**Affiliate** | Pointer to **map[string]interface{}** | field.spotTradeInfo.affiliate | [optional] 

## Methods

### NewSpotTradeInfo

`func NewSpotTradeInfo() *SpotTradeInfo`

NewSpotTradeInfo instantiates a new SpotTradeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotTradeInfoWithDefaults

`func NewSpotTradeInfoWithDefaults() *SpotTradeInfo`

NewSpotTradeInfoWithDefaults instantiates a new SpotTradeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubmitterId

`func (o *SpotTradeInfo) GetSubmitterId() int64`

GetSubmitterId returns the SubmitterId field if non-nil, zero value otherwise.

### GetSubmitterIdOk

`func (o *SpotTradeInfo) GetSubmitterIdOk() (*int64, bool)`

GetSubmitterIdOk returns a tuple with the SubmitterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterId

`func (o *SpotTradeInfo) SetSubmitterId(v int64)`

SetSubmitterId sets SubmitterId field to given value.

### HasSubmitterId

`func (o *SpotTradeInfo) HasSubmitterId() bool`

HasSubmitterId returns a boolean if a field has been set.

### GetSmallOrder

`func (o *SpotTradeInfo) GetSmallOrder() map[string]interface{}`

GetSmallOrder returns the SmallOrder field if non-nil, zero value otherwise.

### GetSmallOrderOk

`func (o *SpotTradeInfo) GetSmallOrderOk() (*map[string]interface{}, bool)`

GetSmallOrderOk returns a tuple with the SmallOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallOrder

`func (o *SpotTradeInfo) SetSmallOrder(v map[string]interface{})`

SetSmallOrder sets SmallOrder field to given value.

### HasSmallOrder

`func (o *SpotTradeInfo) HasSmallOrder() bool`

HasSmallOrder returns a boolean if a field has been set.

### GetAffiliate

`func (o *SpotTradeInfo) GetAffiliate() map[string]interface{}`

GetAffiliate returns the Affiliate field if non-nil, zero value otherwise.

### GetAffiliateOk

`func (o *SpotTradeInfo) GetAffiliateOk() (*map[string]interface{}, bool)`

GetAffiliateOk returns a tuple with the Affiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliate

`func (o *SpotTradeInfo) SetAffiliate(v map[string]interface{})`

SetAffiliate sets Affiliate field to given value.

### HasAffiliate

`func (o *SpotTradeInfo) HasAffiliate() bool`

HasAffiliate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


