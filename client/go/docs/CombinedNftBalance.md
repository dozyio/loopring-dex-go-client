# CombinedNftBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | field.Balance.accountId | 
**AccountId** | **int64** | field.Balance.accountId | 
**TokenId** | **int32** | field.Balance.tokenId | 
**NftData** | Pointer to **string** | Users NFT token nftData | [optional] 
**TokenAddress** | Pointer to **string** | field.NftBalance.tokenAddress | [optional] 
**NftId** | Pointer to **string** | field.NftBalance.nftId | [optional] 
**NftType** | Pointer to **string** | field.NftBalance.nftType | [optional] 
**Total** | **string** | field.Balance.totalAmount | 
**Locked** | **string** | field.Balance.frozenAmount | 
**Pending** | [**PendingBalance**](PendingBalance.md) |  | 
**DeploymentStatus** | Pointer to **string** | field.Balance.deploymentStatus | [optional] 
**IsCounterFactualNFT** | Pointer to **bool** | field.Balance.isCounterFactualNFT | [optional] 
**Metadata** | Pointer to [**NftMetadataL2Info**](NftMetadataL2Info.md) |  | [optional] 

## Methods

### NewCombinedNftBalance

`func NewCombinedNftBalance(id int64, accountId int64, tokenId int32, total string, locked string, pending PendingBalance, ) *CombinedNftBalance`

NewCombinedNftBalance instantiates a new CombinedNftBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombinedNftBalanceWithDefaults

`func NewCombinedNftBalanceWithDefaults() *CombinedNftBalance`

NewCombinedNftBalanceWithDefaults instantiates a new CombinedNftBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CombinedNftBalance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CombinedNftBalance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CombinedNftBalance) SetId(v int64)`

SetId sets Id field to given value.


### GetAccountId

`func (o *CombinedNftBalance) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CombinedNftBalance) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CombinedNftBalance) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetTokenId

`func (o *CombinedNftBalance) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CombinedNftBalance) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CombinedNftBalance) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetNftData

`func (o *CombinedNftBalance) GetNftData() string`

GetNftData returns the NftData field if non-nil, zero value otherwise.

### GetNftDataOk

`func (o *CombinedNftBalance) GetNftDataOk() (*string, bool)`

GetNftDataOk returns a tuple with the NftData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftData

`func (o *CombinedNftBalance) SetNftData(v string)`

SetNftData sets NftData field to given value.

### HasNftData

`func (o *CombinedNftBalance) HasNftData() bool`

HasNftData returns a boolean if a field has been set.

### GetTokenAddress

`func (o *CombinedNftBalance) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *CombinedNftBalance) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *CombinedNftBalance) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *CombinedNftBalance) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetNftId

`func (o *CombinedNftBalance) GetNftId() string`

GetNftId returns the NftId field if non-nil, zero value otherwise.

### GetNftIdOk

`func (o *CombinedNftBalance) GetNftIdOk() (*string, bool)`

GetNftIdOk returns a tuple with the NftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftId

`func (o *CombinedNftBalance) SetNftId(v string)`

SetNftId sets NftId field to given value.

### HasNftId

`func (o *CombinedNftBalance) HasNftId() bool`

HasNftId returns a boolean if a field has been set.

### GetNftType

`func (o *CombinedNftBalance) GetNftType() string`

GetNftType returns the NftType field if non-nil, zero value otherwise.

### GetNftTypeOk

`func (o *CombinedNftBalance) GetNftTypeOk() (*string, bool)`

GetNftTypeOk returns a tuple with the NftType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftType

`func (o *CombinedNftBalance) SetNftType(v string)`

SetNftType sets NftType field to given value.

### HasNftType

`func (o *CombinedNftBalance) HasNftType() bool`

HasNftType returns a boolean if a field has been set.

### GetTotal

`func (o *CombinedNftBalance) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CombinedNftBalance) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CombinedNftBalance) SetTotal(v string)`

SetTotal sets Total field to given value.


### GetLocked

`func (o *CombinedNftBalance) GetLocked() string`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CombinedNftBalance) GetLockedOk() (*string, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CombinedNftBalance) SetLocked(v string)`

SetLocked sets Locked field to given value.


### GetPending

`func (o *CombinedNftBalance) GetPending() PendingBalance`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *CombinedNftBalance) GetPendingOk() (*PendingBalance, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *CombinedNftBalance) SetPending(v PendingBalance)`

SetPending sets Pending field to given value.


### GetDeploymentStatus

`func (o *CombinedNftBalance) GetDeploymentStatus() string`

GetDeploymentStatus returns the DeploymentStatus field if non-nil, zero value otherwise.

### GetDeploymentStatusOk

`func (o *CombinedNftBalance) GetDeploymentStatusOk() (*string, bool)`

GetDeploymentStatusOk returns a tuple with the DeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentStatus

`func (o *CombinedNftBalance) SetDeploymentStatus(v string)`

SetDeploymentStatus sets DeploymentStatus field to given value.

### HasDeploymentStatus

`func (o *CombinedNftBalance) HasDeploymentStatus() bool`

HasDeploymentStatus returns a boolean if a field has been set.

### GetIsCounterFactualNFT

`func (o *CombinedNftBalance) GetIsCounterFactualNFT() bool`

GetIsCounterFactualNFT returns the IsCounterFactualNFT field if non-nil, zero value otherwise.

### GetIsCounterFactualNFTOk

`func (o *CombinedNftBalance) GetIsCounterFactualNFTOk() (*bool, bool)`

GetIsCounterFactualNFTOk returns a tuple with the IsCounterFactualNFT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCounterFactualNFT

`func (o *CombinedNftBalance) SetIsCounterFactualNFT(v bool)`

SetIsCounterFactualNFT sets IsCounterFactualNFT field to given value.

### HasIsCounterFactualNFT

`func (o *CombinedNftBalance) HasIsCounterFactualNFT() bool`

HasIsCounterFactualNFT returns a boolean if a field has been set.

### GetMetadata

`func (o *CombinedNftBalance) GetMetadata() NftMetadataL2Info`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CombinedNftBalance) GetMetadataOk() (*NftMetadataL2Info, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CombinedNftBalance) SetMetadata(v NftMetadataL2Info)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CombinedNftBalance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


