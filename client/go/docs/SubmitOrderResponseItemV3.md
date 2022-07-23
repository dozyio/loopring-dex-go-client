# SubmitOrderResponseItemV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Order hash of submit order response | 
**ClientOrderId** | **string** | The clientOrderId of the submitted order | 
**Status** | **string** | Order status of submit order response | 
**IsIdempotent** | **bool** | Idempotent of submit order response, submit same order again when order was UNKNOWN or WAIT_FREEZE_BALANCE in relayer, idempotent will be true | 
**AccountId** | **int64** | field.SubmitOffChainResponseItem.accountId | 
**Tokens** | **[]map[string]interface{}** | field.SubmitOffChainResponseItem.tokenId | 
**StorageId** | **int64** | field.SubmitOffChainResponseItem.storageId | 

## Methods

### NewSubmitOrderResponseItemV3

`func NewSubmitOrderResponseItemV3(hash string, clientOrderId string, status string, isIdempotent bool, accountId int64, tokens []map[string]interface{}, storageId int64, ) *SubmitOrderResponseItemV3`

NewSubmitOrderResponseItemV3 instantiates a new SubmitOrderResponseItemV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitOrderResponseItemV3WithDefaults

`func NewSubmitOrderResponseItemV3WithDefaults() *SubmitOrderResponseItemV3`

NewSubmitOrderResponseItemV3WithDefaults instantiates a new SubmitOrderResponseItemV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *SubmitOrderResponseItemV3) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SubmitOrderResponseItemV3) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SubmitOrderResponseItemV3) SetHash(v string)`

SetHash sets Hash field to given value.


### GetClientOrderId

`func (o *SubmitOrderResponseItemV3) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *SubmitOrderResponseItemV3) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *SubmitOrderResponseItemV3) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.


### GetStatus

`func (o *SubmitOrderResponseItemV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmitOrderResponseItemV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmitOrderResponseItemV3) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIsIdempotent

`func (o *SubmitOrderResponseItemV3) GetIsIdempotent() bool`

GetIsIdempotent returns the IsIdempotent field if non-nil, zero value otherwise.

### GetIsIdempotentOk

`func (o *SubmitOrderResponseItemV3) GetIsIdempotentOk() (*bool, bool)`

GetIsIdempotentOk returns a tuple with the IsIdempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIdempotent

`func (o *SubmitOrderResponseItemV3) SetIsIdempotent(v bool)`

SetIsIdempotent sets IsIdempotent field to given value.


### GetAccountId

`func (o *SubmitOrderResponseItemV3) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SubmitOrderResponseItemV3) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SubmitOrderResponseItemV3) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetTokens

`func (o *SubmitOrderResponseItemV3) GetTokens() []map[string]interface{}`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *SubmitOrderResponseItemV3) GetTokensOk() (*[]map[string]interface{}, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *SubmitOrderResponseItemV3) SetTokens(v []map[string]interface{})`

SetTokens sets Tokens field to given value.


### GetStorageId

`func (o *SubmitOrderResponseItemV3) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *SubmitOrderResponseItemV3) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *SubmitOrderResponseItemV3) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


