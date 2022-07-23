# SubmitOrderResponseV2Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderHash** | **string** | field.SubmitOrderResponseV2Item.orderHash | 
**Status** | **string** | field.SubmitOrderResponseV2Item.status | 
**IsIdempotent** | **bool** | field.SubmitOrderResponseV2Item.isIdempotent | 

## Methods

### NewSubmitOrderResponseV2Item

`func NewSubmitOrderResponseV2Item(orderHash string, status string, isIdempotent bool, ) *SubmitOrderResponseV2Item`

NewSubmitOrderResponseV2Item instantiates a new SubmitOrderResponseV2Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitOrderResponseV2ItemWithDefaults

`func NewSubmitOrderResponseV2ItemWithDefaults() *SubmitOrderResponseV2Item`

NewSubmitOrderResponseV2ItemWithDefaults instantiates a new SubmitOrderResponseV2Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderHash

`func (o *SubmitOrderResponseV2Item) GetOrderHash() string`

GetOrderHash returns the OrderHash field if non-nil, zero value otherwise.

### GetOrderHashOk

`func (o *SubmitOrderResponseV2Item) GetOrderHashOk() (*string, bool)`

GetOrderHashOk returns a tuple with the OrderHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderHash

`func (o *SubmitOrderResponseV2Item) SetOrderHash(v string)`

SetOrderHash sets OrderHash field to given value.


### GetStatus

`func (o *SubmitOrderResponseV2Item) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitOrderResponseV2Item) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitOrderResponseV2Item) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIsIdempotent

`func (o *SubmitOrderResponseV2Item) GetIsIdempotent() bool`

GetIsIdempotent returns the IsIdempotent field if non-nil, zero value otherwise.

### GetIsIdempotentOk

`func (o *SubmitOrderResponseV2Item) GetIsIdempotentOk() (*bool, bool)`

GetIsIdempotentOk returns a tuple with the IsIdempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIdempotent

`func (o *SubmitOrderResponseV2Item) SetIsIdempotent(v bool)`

SetIsIdempotent sets IsIdempotent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


