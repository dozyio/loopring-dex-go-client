# SubmitOrderResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Hash of order submitted | 
**Error** | Pointer to [**ResultInfo**](ResultInfo.md) |  | [optional] 

## Methods

### NewSubmitOrderResult

`func NewSubmitOrderResult(hash string, ) *SubmitOrderResult`

NewSubmitOrderResult instantiates a new SubmitOrderResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitOrderResultWithDefaults

`func NewSubmitOrderResultWithDefaults() *SubmitOrderResult`

NewSubmitOrderResultWithDefaults instantiates a new SubmitOrderResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *SubmitOrderResult) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SubmitOrderResult) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SubmitOrderResult) SetHash(v string)`

SetHash sets Hash field to given value.


### GetError

`func (o *SubmitOrderResult) GetError() ResultInfo`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SubmitOrderResult) GetErrorOk() (*ResultInfo, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SubmitOrderResult) SetError(v ResultInfo)`

SetError sets Error field to given value.

### HasError

`func (o *SubmitOrderResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


