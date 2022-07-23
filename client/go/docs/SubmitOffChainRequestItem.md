# SubmitOffChainRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | The order hash identifier set by the user at the time of submission | 
**Status** | **string** | Whether the order was successfully submitted or not, please note, user may query after a while to get real process status, as most offchain requests are async processed | 
**IsIdempotent** | **bool** | Idempotent of submit order response, submit same order again when order was UNKNOWN or WAIT_FREEZE_BALANCE in relayer, idempotent will be true | 
**AccountId** | **int64** | field.SubmitOffChainResponseItem.accountId | 
**JoinTokensIds** | **[]map[string]interface{}** | field.SubmitOffChainResponseItem.tokenId | 
**JoinStorageIds** | **[]map[string]interface{}** | field.SubmitOffChainResponseItem.storageId | 

## Methods

### NewSubmitOffChainRequestItem

`func NewSubmitOffChainRequestItem(hash string, status string, isIdempotent bool, accountId int64, joinTokensIds []map[string]interface{}, joinStorageIds []map[string]interface{}, ) *SubmitOffChainRequestItem`

NewSubmitOffChainRequestItem instantiates a new SubmitOffChainRequestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitOffChainRequestItemWithDefaults

`func NewSubmitOffChainRequestItemWithDefaults() *SubmitOffChainRequestItem`

NewSubmitOffChainRequestItemWithDefaults instantiates a new SubmitOffChainRequestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *SubmitOffChainRequestItem) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SubmitOffChainRequestItem) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SubmitOffChainRequestItem) SetHash(v string)`

SetHash sets Hash field to given value.


### GetStatus

`func (o *SubmitOffChainRequestItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitOffChainRequestItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitOffChainRequestItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIsIdempotent

`func (o *SubmitOffChainRequestItem) GetIsIdempotent() bool`

GetIsIdempotent returns the IsIdempotent field if non-nil, zero value otherwise.

### GetIsIdempotentOk

`func (o *SubmitOffChainRequestItem) GetIsIdempotentOk() (*bool, bool)`

GetIsIdempotentOk returns a tuple with the IsIdempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIdempotent

`func (o *SubmitOffChainRequestItem) SetIsIdempotent(v bool)`

SetIsIdempotent sets IsIdempotent field to given value.


### GetAccountId

`func (o *SubmitOffChainRequestItem) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SubmitOffChainRequestItem) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SubmitOffChainRequestItem) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetJoinTokensIds

`func (o *SubmitOffChainRequestItem) GetJoinTokensIds() []map[string]interface{}`

GetJoinTokensIds returns the JoinTokensIds field if non-nil, zero value otherwise.

### GetJoinTokensIdsOk

`func (o *SubmitOffChainRequestItem) GetJoinTokensIdsOk() (*[]map[string]interface{}, bool)`

GetJoinTokensIdsOk returns a tuple with the JoinTokensIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTokensIds

`func (o *SubmitOffChainRequestItem) SetJoinTokensIds(v []map[string]interface{})`

SetJoinTokensIds sets JoinTokensIds field to given value.


### GetJoinStorageIds

`func (o *SubmitOffChainRequestItem) GetJoinStorageIds() []map[string]interface{}`

GetJoinStorageIds returns the JoinStorageIds field if non-nil, zero value otherwise.

### GetJoinStorageIdsOk

`func (o *SubmitOffChainRequestItem) GetJoinStorageIdsOk() (*[]map[string]interface{}, bool)`

GetJoinStorageIdsOk returns a tuple with the JoinStorageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinStorageIds

`func (o *SubmitOffChainRequestItem) SetJoinStorageIds(v []map[string]interface{})`

SetJoinStorageIds sets JoinStorageIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


