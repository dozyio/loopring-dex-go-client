# DexAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | field.dexAccount.accountId | 
**Owner** | **string** | field.dexAccount.owner | 
**Frozen** | **bool** | field.dexAccount.isFreeze | 
**PublicKeyX** | **string** | field.dexAccount.publicKeyX | 
**PublicKeyY** | **string** | field.dexAccount.publicKeyY | 
**Tags** | Pointer to **string** | field.dexAccount.tags | [optional] 
**KeyNonce** | **int64** | field.dexAccount.keyNonce | 
**AccountNonce** | **int64** | field.dexAccount.accountNonce | 

## Methods

### NewDexAccount

`func NewDexAccount(accountId int64, owner string, frozen bool, publicKeyX string, publicKeyY string, keyNonce int64, accountNonce int64, ) *DexAccount`

NewDexAccount instantiates a new DexAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexAccountWithDefaults

`func NewDexAccountWithDefaults() *DexAccount`

NewDexAccountWithDefaults instantiates a new DexAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *DexAccount) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DexAccount) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DexAccount) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetOwner

`func (o *DexAccount) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DexAccount) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DexAccount) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetFrozen

`func (o *DexAccount) GetFrozen() bool`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *DexAccount) GetFrozenOk() (*bool, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *DexAccount) SetFrozen(v bool)`

SetFrozen sets Frozen field to given value.


### GetPublicKeyX

`func (o *DexAccount) GetPublicKeyX() string`

GetPublicKeyX returns the PublicKeyX field if non-nil, zero value otherwise.

### GetPublicKeyXOk

`func (o *DexAccount) GetPublicKeyXOk() (*string, bool)`

GetPublicKeyXOk returns a tuple with the PublicKeyX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyX

`func (o *DexAccount) SetPublicKeyX(v string)`

SetPublicKeyX sets PublicKeyX field to given value.


### GetPublicKeyY

`func (o *DexAccount) GetPublicKeyY() string`

GetPublicKeyY returns the PublicKeyY field if non-nil, zero value otherwise.

### GetPublicKeyYOk

`func (o *DexAccount) GetPublicKeyYOk() (*string, bool)`

GetPublicKeyYOk returns a tuple with the PublicKeyY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyY

`func (o *DexAccount) SetPublicKeyY(v string)`

SetPublicKeyY sets PublicKeyY field to given value.


### GetTags

`func (o *DexAccount) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DexAccount) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DexAccount) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DexAccount) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetKeyNonce

`func (o *DexAccount) GetKeyNonce() int64`

GetKeyNonce returns the KeyNonce field if non-nil, zero value otherwise.

### GetKeyNonceOk

`func (o *DexAccount) GetKeyNonceOk() (*int64, bool)`

GetKeyNonceOk returns a tuple with the KeyNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyNonce

`func (o *DexAccount) SetKeyNonce(v int64)`

SetKeyNonce sets KeyNonce field to given value.


### GetAccountNonce

`func (o *DexAccount) GetAccountNonce() int64`

GetAccountNonce returns the AccountNonce field if non-nil, zero value otherwise.

### GetAccountNonceOk

`func (o *DexAccount) GetAccountNonceOk() (*int64, bool)`

GetAccountNonceOk returns a tuple with the AccountNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNonce

`func (o *DexAccount) SetAccountNonce(v int64)`

SetAccountNonce sets AccountNonce field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


