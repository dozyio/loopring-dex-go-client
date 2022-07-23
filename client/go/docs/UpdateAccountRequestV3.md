# UpdateAccountRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | exchange address | 
**Owner** | **string** | owner address | 
**AccountId** | **int64** | user account ID | 
**PublicKey** | [**PublicKey**](PublicKey.md) |  | 
**MaxFee** | [**TokenVolumeV3**](TokenVolumeV3.md) |  | 
**ValidUntil** | **int32** | Timestamp for order to become invalid | 
**Nonce** | **int32** | Nonce of users exchange account that used in off-chain requests. | 
**KeySeed** | Pointer to **string** | KeySeed of users L2 eddsaKey, the L2 key should be generated from this seed, i.e., L2_EDDSA_KEY&#x3D;eth.sign(keySeed). Otherwise, user may meet error in login loopring DEX | [optional] 
**CounterFactualInfo** | Pointer to [**CounterFactualInfo**](CounterFactualInfo.md) |  | [optional] 
**EddsaSignature** | Pointer to **string** | eddsa signature of this request | [optional] 
**EcdsaSignature** | Pointer to **string** | ecdsa signature of this request | [optional] 
**HashApproved** | Pointer to **string** | An approved hash string which was submitted on eth mainnet | [optional] 

## Methods

### NewUpdateAccountRequestV3

`func NewUpdateAccountRequestV3(exchange string, owner string, accountId int64, publicKey PublicKey, maxFee TokenVolumeV3, validUntil int32, nonce int32, ) *UpdateAccountRequestV3`

NewUpdateAccountRequestV3 instantiates a new UpdateAccountRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountRequestV3WithDefaults

`func NewUpdateAccountRequestV3WithDefaults() *UpdateAccountRequestV3`

NewUpdateAccountRequestV3WithDefaults instantiates a new UpdateAccountRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *UpdateAccountRequestV3) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *UpdateAccountRequestV3) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *UpdateAccountRequestV3) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetOwner

`func (o *UpdateAccountRequestV3) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateAccountRequestV3) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateAccountRequestV3) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetAccountId

`func (o *UpdateAccountRequestV3) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateAccountRequestV3) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateAccountRequestV3) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetPublicKey

`func (o *UpdateAccountRequestV3) GetPublicKey() PublicKey`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *UpdateAccountRequestV3) GetPublicKeyOk() (*PublicKey, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *UpdateAccountRequestV3) SetPublicKey(v PublicKey)`

SetPublicKey sets PublicKey field to given value.


### GetMaxFee

`func (o *UpdateAccountRequestV3) GetMaxFee() TokenVolumeV3`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *UpdateAccountRequestV3) GetMaxFeeOk() (*TokenVolumeV3, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *UpdateAccountRequestV3) SetMaxFee(v TokenVolumeV3)`

SetMaxFee sets MaxFee field to given value.


### GetValidUntil

`func (o *UpdateAccountRequestV3) GetValidUntil() int32`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *UpdateAccountRequestV3) GetValidUntilOk() (*int32, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *UpdateAccountRequestV3) SetValidUntil(v int32)`

SetValidUntil sets ValidUntil field to given value.


### GetNonce

`func (o *UpdateAccountRequestV3) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *UpdateAccountRequestV3) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *UpdateAccountRequestV3) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetKeySeed

`func (o *UpdateAccountRequestV3) GetKeySeed() string`

GetKeySeed returns the KeySeed field if non-nil, zero value otherwise.

### GetKeySeedOk

`func (o *UpdateAccountRequestV3) GetKeySeedOk() (*string, bool)`

GetKeySeedOk returns a tuple with the KeySeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySeed

`func (o *UpdateAccountRequestV3) SetKeySeed(v string)`

SetKeySeed sets KeySeed field to given value.

### HasKeySeed

`func (o *UpdateAccountRequestV3) HasKeySeed() bool`

HasKeySeed returns a boolean if a field has been set.

### GetCounterFactualInfo

`func (o *UpdateAccountRequestV3) GetCounterFactualInfo() CounterFactualInfo`

GetCounterFactualInfo returns the CounterFactualInfo field if non-nil, zero value otherwise.

### GetCounterFactualInfoOk

`func (o *UpdateAccountRequestV3) GetCounterFactualInfoOk() (*CounterFactualInfo, bool)`

GetCounterFactualInfoOk returns a tuple with the CounterFactualInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterFactualInfo

`func (o *UpdateAccountRequestV3) SetCounterFactualInfo(v CounterFactualInfo)`

SetCounterFactualInfo sets CounterFactualInfo field to given value.

### HasCounterFactualInfo

`func (o *UpdateAccountRequestV3) HasCounterFactualInfo() bool`

HasCounterFactualInfo returns a boolean if a field has been set.

### GetEddsaSignature

`func (o *UpdateAccountRequestV3) GetEddsaSignature() string`

GetEddsaSignature returns the EddsaSignature field if non-nil, zero value otherwise.

### GetEddsaSignatureOk

`func (o *UpdateAccountRequestV3) GetEddsaSignatureOk() (*string, bool)`

GetEddsaSignatureOk returns a tuple with the EddsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSignature

`func (o *UpdateAccountRequestV3) SetEddsaSignature(v string)`

SetEddsaSignature sets EddsaSignature field to given value.

### HasEddsaSignature

`func (o *UpdateAccountRequestV3) HasEddsaSignature() bool`

HasEddsaSignature returns a boolean if a field has been set.

### GetEcdsaSignature

`func (o *UpdateAccountRequestV3) GetEcdsaSignature() string`

GetEcdsaSignature returns the EcdsaSignature field if non-nil, zero value otherwise.

### GetEcdsaSignatureOk

`func (o *UpdateAccountRequestV3) GetEcdsaSignatureOk() (*string, bool)`

GetEcdsaSignatureOk returns a tuple with the EcdsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcdsaSignature

`func (o *UpdateAccountRequestV3) SetEcdsaSignature(v string)`

SetEcdsaSignature sets EcdsaSignature field to given value.

### HasEcdsaSignature

`func (o *UpdateAccountRequestV3) HasEcdsaSignature() bool`

HasEcdsaSignature returns a boolean if a field has been set.

### GetHashApproved

`func (o *UpdateAccountRequestV3) GetHashApproved() string`

GetHashApproved returns the HashApproved field if non-nil, zero value otherwise.

### GetHashApprovedOk

`func (o *UpdateAccountRequestV3) GetHashApprovedOk() (*string, bool)`

GetHashApprovedOk returns a tuple with the HashApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashApproved

`func (o *UpdateAccountRequestV3) SetHashApproved(v string)`

SetHashApproved sets HashApproved field to given value.

### HasHashApproved

`func (o *UpdateAccountRequestV3) HasHashApproved() bool`

HasHashApproved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


