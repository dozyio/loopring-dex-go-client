# AmmPoolInfoV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | AMM pool name, used to calculate domainSeparator of EIP712 hash if use chooses ECDSA sign path. | 
**Market** | **string** | AMM pool market name, can be used to query market related info like ticker, etc | 
**Address** | **string** | AMM pool address | 
**Version** | **string** | AMM pool contract version | 
**Tokens** | [**AmmPoolTokens**](AmmPoolTokens.md) |  | 
**FeeBips** | **int32** | AMM fee bips | 
**Precisions** | [**AmmPoolPrecisions**](AmmPoolPrecisions.md) |  | 
**CreatedAt** | **string** | AMM pool online date | 
**Status** | **int32** | AMM market status | 

## Methods

### NewAmmPoolInfoV3

`func NewAmmPoolInfoV3(name string, market string, address string, version string, tokens AmmPoolTokens, feeBips int32, precisions AmmPoolPrecisions, createdAt string, status int32, ) *AmmPoolInfoV3`

NewAmmPoolInfoV3 instantiates a new AmmPoolInfoV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmPoolInfoV3WithDefaults

`func NewAmmPoolInfoV3WithDefaults() *AmmPoolInfoV3`

NewAmmPoolInfoV3WithDefaults instantiates a new AmmPoolInfoV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AmmPoolInfoV3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmmPoolInfoV3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmmPoolInfoV3) SetName(v string)`

SetName sets Name field to given value.


### GetMarket

`func (o *AmmPoolInfoV3) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *AmmPoolInfoV3) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *AmmPoolInfoV3) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetAddress

`func (o *AmmPoolInfoV3) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AmmPoolInfoV3) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AmmPoolInfoV3) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetVersion

`func (o *AmmPoolInfoV3) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AmmPoolInfoV3) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AmmPoolInfoV3) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetTokens

`func (o *AmmPoolInfoV3) GetTokens() AmmPoolTokens`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *AmmPoolInfoV3) GetTokensOk() (*AmmPoolTokens, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *AmmPoolInfoV3) SetTokens(v AmmPoolTokens)`

SetTokens sets Tokens field to given value.


### GetFeeBips

`func (o *AmmPoolInfoV3) GetFeeBips() int32`

GetFeeBips returns the FeeBips field if non-nil, zero value otherwise.

### GetFeeBipsOk

`func (o *AmmPoolInfoV3) GetFeeBipsOk() (*int32, bool)`

GetFeeBipsOk returns a tuple with the FeeBips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBips

`func (o *AmmPoolInfoV3) SetFeeBips(v int32)`

SetFeeBips sets FeeBips field to given value.


### GetPrecisions

`func (o *AmmPoolInfoV3) GetPrecisions() AmmPoolPrecisions`

GetPrecisions returns the Precisions field if non-nil, zero value otherwise.

### GetPrecisionsOk

`func (o *AmmPoolInfoV3) GetPrecisionsOk() (*AmmPoolPrecisions, bool)`

GetPrecisionsOk returns a tuple with the Precisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecisions

`func (o *AmmPoolInfoV3) SetPrecisions(v AmmPoolPrecisions)`

SetPrecisions sets Precisions field to given value.


### GetCreatedAt

`func (o *AmmPoolInfoV3) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AmmPoolInfoV3) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AmmPoolInfoV3) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetStatus

`func (o *AmmPoolInfoV3) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AmmPoolInfoV3) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AmmPoolInfoV3) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


