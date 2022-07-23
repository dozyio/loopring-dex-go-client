# RequestInBlockResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**BlockId** | **int64** | The block Id. | 
**IndexInBlock** | **int32** | The tx index/postion in block | 

## Methods

### NewRequestInBlockResponse

`func NewRequestInBlockResponse(hash string, blockId int64, indexInBlock int32, ) *RequestInBlockResponse`

NewRequestInBlockResponse instantiates a new RequestInBlockResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestInBlockResponseWithDefaults

`func NewRequestInBlockResponseWithDefaults() *RequestInBlockResponse`

NewRequestInBlockResponseWithDefaults instantiates a new RequestInBlockResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *RequestInBlockResponse) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *RequestInBlockResponse) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *RequestInBlockResponse) SetHash(v string)`

SetHash sets Hash field to given value.


### GetBlockId

`func (o *RequestInBlockResponse) GetBlockId() int64`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *RequestInBlockResponse) GetBlockIdOk() (*int64, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *RequestInBlockResponse) SetBlockId(v int64)`

SetBlockId sets BlockId field to given value.


### GetIndexInBlock

`func (o *RequestInBlockResponse) GetIndexInBlock() int32`

GetIndexInBlock returns the IndexInBlock field if non-nil, zero value otherwise.

### GetIndexInBlockOk

`func (o *RequestInBlockResponse) GetIndexInBlockOk() (*int32, bool)`

GetIndexInBlockOk returns a tuple with the IndexInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexInBlock

`func (o *RequestInBlockResponse) SetIndexInBlock(v int32)`

SetIndexInBlock sets IndexInBlock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


