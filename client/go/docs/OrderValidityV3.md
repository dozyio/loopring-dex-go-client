# OrderValidityV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **int64** | Timestamp from when the order officially becomes valid and fillable | 
**End** | **int64** | Timestamp from when the order ceases to be valid and fillable | 

## Methods

### NewOrderValidityV3

`func NewOrderValidityV3(start int64, end int64, ) *OrderValidityV3`

NewOrderValidityV3 instantiates a new OrderValidityV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderValidityV3WithDefaults

`func NewOrderValidityV3WithDefaults() *OrderValidityV3`

NewOrderValidityV3WithDefaults instantiates a new OrderValidityV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *OrderValidityV3) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *OrderValidityV3) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *OrderValidityV3) SetStart(v int64)`

SetStart sets Start field to given value.


### GetEnd

`func (o *OrderValidityV3) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *OrderValidityV3) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *OrderValidityV3) SetEnd(v int64)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


