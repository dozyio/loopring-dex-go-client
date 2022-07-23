# NftMetadataL2Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** |  | 
**Base** | Pointer to [**NftMetadataBaseInfo**](NftMetadataBaseInfo.md) |  | [optional] 
**ImageSize** | **map[string]string** |  | 
**Extra** | Pointer to [**NftExtra**](NftExtra.md) |  | [optional] 
**Status** | **int32** |  | 
**NftType** | **int32** |  | 
**Network** | **int32** |  | 
**TokenAddress** | **string** |  | 
**NftId** | **string** |  | 

## Methods

### NewNftMetadataL2Info

`func NewNftMetadataL2Info(uri string, imageSize map[string]string, status int32, nftType int32, network int32, tokenAddress string, nftId string, ) *NftMetadataL2Info`

NewNftMetadataL2Info instantiates a new NftMetadataL2Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftMetadataL2InfoWithDefaults

`func NewNftMetadataL2InfoWithDefaults() *NftMetadataL2Info`

NewNftMetadataL2InfoWithDefaults instantiates a new NftMetadataL2Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *NftMetadataL2Info) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *NftMetadataL2Info) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *NftMetadataL2Info) SetUri(v string)`

SetUri sets Uri field to given value.


### GetBase

`func (o *NftMetadataL2Info) GetBase() NftMetadataBaseInfo`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *NftMetadataL2Info) GetBaseOk() (*NftMetadataBaseInfo, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *NftMetadataL2Info) SetBase(v NftMetadataBaseInfo)`

SetBase sets Base field to given value.

### HasBase

`func (o *NftMetadataL2Info) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetImageSize

`func (o *NftMetadataL2Info) GetImageSize() map[string]string`

GetImageSize returns the ImageSize field if non-nil, zero value otherwise.

### GetImageSizeOk

`func (o *NftMetadataL2Info) GetImageSizeOk() (*map[string]string, bool)`

GetImageSizeOk returns a tuple with the ImageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSize

`func (o *NftMetadataL2Info) SetImageSize(v map[string]string)`

SetImageSize sets ImageSize field to given value.


### GetExtra

`func (o *NftMetadataL2Info) GetExtra() NftExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *NftMetadataL2Info) GetExtraOk() (*NftExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *NftMetadataL2Info) SetExtra(v NftExtra)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *NftMetadataL2Info) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetStatus

`func (o *NftMetadataL2Info) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NftMetadataL2Info) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NftMetadataL2Info) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetNftType

`func (o *NftMetadataL2Info) GetNftType() int32`

GetNftType returns the NftType field if non-nil, zero value otherwise.

### GetNftTypeOk

`func (o *NftMetadataL2Info) GetNftTypeOk() (*int32, bool)`

GetNftTypeOk returns a tuple with the NftType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftType

`func (o *NftMetadataL2Info) SetNftType(v int32)`

SetNftType sets NftType field to given value.


### GetNetwork

`func (o *NftMetadataL2Info) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NftMetadataL2Info) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NftMetadataL2Info) SetNetwork(v int32)`

SetNetwork sets Network field to given value.


### GetTokenAddress

`func (o *NftMetadataL2Info) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *NftMetadataL2Info) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *NftMetadataL2Info) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetNftId

`func (o *NftMetadataL2Info) GetNftId() string`

GetNftId returns the NftId field if non-nil, zero value otherwise.

### GetNftIdOk

`func (o *NftMetadataL2Info) GetNftIdOk() (*string, bool)`

GetNftIdOk returns a tuple with the NftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftId

`func (o *NftMetadataL2Info) SetNftId(v string)`

SetNftId sets NftId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


