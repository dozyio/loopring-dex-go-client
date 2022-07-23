# BlockIdInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockId** | **int64** |  | 
**IndexInBlock** | **int32** |  | 

## Methods

### NewBlockIdInfo

`func NewBlockIdInfo(blockId int64, indexInBlock int32, ) *BlockIdInfo`

NewBlockIdInfo instantiates a new BlockIdInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockIdInfoWithDefaults

`func NewBlockIdInfoWithDefaults() *BlockIdInfo`

NewBlockIdInfoWithDefaults instantiates a new BlockIdInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockId

`func (o *BlockIdInfo) GetBlockId() int64`

GetBlockId returns the BlockId field if non-nil, zero value otherwise.

### GetBlockIdOk

`func (o *BlockIdInfo) GetBlockIdOk() (*int64, bool)`

GetBlockIdOk returns a tuple with the BlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockId

`func (o *BlockIdInfo) SetBlockId(v int64)`

SetBlockId sets BlockId field to given value.


### GetIndexInBlock

`func (o *BlockIdInfo) GetIndexInBlock() int32`

GetIndexInBlock returns the IndexInBlock field if non-nil, zero value otherwise.

### GetIndexInBlockOk

`func (o *BlockIdInfo) GetIndexInBlockOk() (*int32, bool)`

GetIndexInBlockOk returns a tuple with the IndexInBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexInBlock

`func (o *BlockIdInfo) SetIndexInBlock(v int32)`

SetIndexInBlock sets IndexInBlock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


