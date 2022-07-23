# OrderAmountsV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Minimum** | **string** | The minimum amount enforced when submitting orders for the token. | 
**Maximum** | **string** | The maximum amount enforced when submitting orders for the token. | 
**Dust** | **string** | The dust amount enforced when submitting orders for the token. | 

## Methods

### NewOrderAmountsV3

`func NewOrderAmountsV3(minimum string, maximum string, dust string, ) *OrderAmountsV3`

NewOrderAmountsV3 instantiates a new OrderAmountsV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmountsV3WithDefaults

`func NewOrderAmountsV3WithDefaults() *OrderAmountsV3`

NewOrderAmountsV3WithDefaults instantiates a new OrderAmountsV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimum

`func (o *OrderAmountsV3) GetMinimum() string`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *OrderAmountsV3) GetMinimumOk() (*string, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *OrderAmountsV3) SetMinimum(v string)`

SetMinimum sets Minimum field to given value.


### GetMaximum

`func (o *OrderAmountsV3) GetMaximum() string`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *OrderAmountsV3) GetMaximumOk() (*string, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *OrderAmountsV3) SetMaximum(v string)`

SetMaximum sets Maximum field to given value.


### GetDust

`func (o *OrderAmountsV3) GetDust() string`

GetDust returns the Dust field if non-nil, zero value otherwise.

### GetDustOk

`func (o *OrderAmountsV3) GetDustOk() (*string, bool)`

GetDustOk returns a tuple with the Dust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDust

`func (o *OrderAmountsV3) SetDust(v string)`

SetDust sets Dust field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


