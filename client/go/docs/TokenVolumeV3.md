# TokenVolumeV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **int32** | The Loopring&#39;s token identifier. | 
**Volume** | **string** | The volume of the token | 

## Methods

### NewTokenVolumeV3

`func NewTokenVolumeV3(tokenId int32, volume string, ) *TokenVolumeV3`

NewTokenVolumeV3 instantiates a new TokenVolumeV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenVolumeV3WithDefaults

`func NewTokenVolumeV3WithDefaults() *TokenVolumeV3`

NewTokenVolumeV3WithDefaults instantiates a new TokenVolumeV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TokenVolumeV3) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenVolumeV3) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenVolumeV3) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetVolume

`func (o *TokenVolumeV3) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *TokenVolumeV3) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *TokenVolumeV3) SetVolume(v string)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


