# DexAccountV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | Account ID | 
**Owner** | **string** | Ethereum address | 
**Frozen** | **bool** | The frozen state of the account, true stands for frozen, if the account is frozen, the user cant submit order. | 
**PublicKey** | [**PublicKey**](PublicKey.md) |  | 
**Tags** | Pointer to **string** | Comma separated list of tags such as VIP levels, etc | [optional] 
**Nonce** | **int64** | field.DexAccountV3.nonce | 
**KeyNonce** | **int64** | Nonce of users key change request, for backward compatible | 
**KeySeed** | **string** | KeySeed of users L2 eddsaKey, the L2 key should be generated from this seed, i.e., L2_EDDSA_KEY&#x3D;eth.sign(keySeed). Otherwise, user may meet error in login loopring DEX | 

## Methods

### NewDexAccountV3

`func NewDexAccountV3(accountId int64, owner string, frozen bool, publicKey PublicKey, nonce int64, keyNonce int64, keySeed string, ) *DexAccountV3`

NewDexAccountV3 instantiates a new DexAccountV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexAccountV3WithDefaults

`func NewDexAccountV3WithDefaults() *DexAccountV3`

NewDexAccountV3WithDefaults instantiates a new DexAccountV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *DexAccountV3) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DexAccountV3) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DexAccountV3) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetOwner

`func (o *DexAccountV3) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DexAccountV3) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DexAccountV3) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetFrozen

`func (o *DexAccountV3) GetFrozen() bool`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *DexAccountV3) GetFrozenOk() (*bool, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *DexAccountV3) SetFrozen(v bool)`

SetFrozen sets Frozen field to given value.


### GetPublicKey

`func (o *DexAccountV3) GetPublicKey() PublicKey`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DexAccountV3) GetPublicKeyOk() (*PublicKey, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DexAccountV3) SetPublicKey(v PublicKey)`

SetPublicKey sets PublicKey field to given value.


### GetTags

`func (o *DexAccountV3) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DexAccountV3) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DexAccountV3) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DexAccountV3) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNonce

`func (o *DexAccountV3) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *DexAccountV3) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *DexAccountV3) SetNonce(v int64)`

SetNonce sets Nonce field to given value.


### GetKeyNonce

`func (o *DexAccountV3) GetKeyNonce() int64`

GetKeyNonce returns the KeyNonce field if non-nil, zero value otherwise.

### GetKeyNonceOk

`func (o *DexAccountV3) GetKeyNonceOk() (*int64, bool)`

GetKeyNonceOk returns a tuple with the KeyNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyNonce

`func (o *DexAccountV3) SetKeyNonce(v int64)`

SetKeyNonce sets KeyNonce field to given value.


### GetKeySeed

`func (o *DexAccountV3) GetKeySeed() string`

GetKeySeed returns the KeySeed field if non-nil, zero value otherwise.

### GetKeySeedOk

`func (o *DexAccountV3) GetKeySeedOk() (*string, bool)`

GetKeySeedOk returns a tuple with the KeySeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySeed

`func (o *DexAccountV3) SetKeySeed(v string)`

SetKeySeed sets KeySeed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


