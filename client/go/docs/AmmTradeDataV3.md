# AmmTradeDataV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | trader&#39;s accountId | 
**OrderHash** | **string** | AMM trade&#39;s order hash | 
**Market** | **string** | The trade market | 
**Side** | **string** | The trade direction, buy or sell | 
**Size** | **string** | The trade volume | 
**Price** | **float64** | The trade price | 
**FeeAmount** | **string** | The trade fee | 
**CreatedAt** | **int64** | The trade&#39;s creation time | 
**TokenId** | **int32** | field.SubmitOffChainResponseItem.tokenId | 
**StorageId** | **int64** | field.SubmitOffChainResponseItem.storageId | 

## Methods

### NewAmmTradeDataV3

`func NewAmmTradeDataV3(accountId int64, orderHash string, market string, side string, size string, price float64, feeAmount string, createdAt int64, tokenId int32, storageId int64, ) *AmmTradeDataV3`

NewAmmTradeDataV3 instantiates a new AmmTradeDataV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmTradeDataV3WithDefaults

`func NewAmmTradeDataV3WithDefaults() *AmmTradeDataV3`

NewAmmTradeDataV3WithDefaults instantiates a new AmmTradeDataV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AmmTradeDataV3) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AmmTradeDataV3) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AmmTradeDataV3) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetOrderHash

`func (o *AmmTradeDataV3) GetOrderHash() string`

GetOrderHash returns the OrderHash field if non-nil, zero value otherwise.

### GetOrderHashOk

`func (o *AmmTradeDataV3) GetOrderHashOk() (*string, bool)`

GetOrderHashOk returns a tuple with the OrderHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderHash

`func (o *AmmTradeDataV3) SetOrderHash(v string)`

SetOrderHash sets OrderHash field to given value.


### GetMarket

`func (o *AmmTradeDataV3) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *AmmTradeDataV3) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *AmmTradeDataV3) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetSide

`func (o *AmmTradeDataV3) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *AmmTradeDataV3) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *AmmTradeDataV3) SetSide(v string)`

SetSide sets Side field to given value.


### GetSize

`func (o *AmmTradeDataV3) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AmmTradeDataV3) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AmmTradeDataV3) SetSize(v string)`

SetSize sets Size field to given value.


### GetPrice

`func (o *AmmTradeDataV3) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AmmTradeDataV3) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AmmTradeDataV3) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetFeeAmount

`func (o *AmmTradeDataV3) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *AmmTradeDataV3) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *AmmTradeDataV3) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetCreatedAt

`func (o *AmmTradeDataV3) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AmmTradeDataV3) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AmmTradeDataV3) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetTokenId

`func (o *AmmTradeDataV3) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AmmTradeDataV3) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AmmTradeDataV3) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetStorageId

`func (o *AmmTradeDataV3) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *AmmTradeDataV3) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *AmmTradeDataV3) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


