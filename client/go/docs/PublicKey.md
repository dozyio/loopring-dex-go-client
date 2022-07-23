# PublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | **string** | The public keys x part. | 
**Y** | **string** | The public keys y part. | 

## Methods

### NewPublicKey

`func NewPublicKey(x string, y string, ) *PublicKey`

NewPublicKey instantiates a new PublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicKeyWithDefaults

`func NewPublicKeyWithDefaults() *PublicKey`

NewPublicKeyWithDefaults instantiates a new PublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *PublicKey) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *PublicKey) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *PublicKey) SetX(v string)`

SetX sets X field to given value.


### GetY

`func (o *PublicKey) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *PublicKey) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *PublicKey) SetY(v string)`

SetY sets Y field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


