# CombineMarketInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | **string** | Trading pair ID | 
**BaseTokenId** | **int32** | The base token ID | 
**QuoteTokenId** | **int32** | The quote token ID | 
**PrecisionForPrice** | **int32** | The precision of price | 
**OrderbookAggLevels** | **int32** | The max level of orderbook price aggregation | 
**Enabled** | **bool** | True if trading is enabled for this trading pair | 
**Status** | **int32** | field.CombineMarketInfo.status | 
**CreatedAt** | **string** | field.AmmMarketInfo.createdAt | 

## Methods

### NewCombineMarketInfo

`func NewCombineMarketInfo(market string, baseTokenId int32, quoteTokenId int32, precisionForPrice int32, orderbookAggLevels int32, enabled bool, status int32, createdAt string, ) *CombineMarketInfo`

NewCombineMarketInfo instantiates a new CombineMarketInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombineMarketInfoWithDefaults

`func NewCombineMarketInfoWithDefaults() *CombineMarketInfo`

NewCombineMarketInfoWithDefaults instantiates a new CombineMarketInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *CombineMarketInfo) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *CombineMarketInfo) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *CombineMarketInfo) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetBaseTokenId

`func (o *CombineMarketInfo) GetBaseTokenId() int32`

GetBaseTokenId returns the BaseTokenId field if non-nil, zero value otherwise.

### GetBaseTokenIdOk

`func (o *CombineMarketInfo) GetBaseTokenIdOk() (*int32, bool)`

GetBaseTokenIdOk returns a tuple with the BaseTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseTokenId

`func (o *CombineMarketInfo) SetBaseTokenId(v int32)`

SetBaseTokenId sets BaseTokenId field to given value.


### GetQuoteTokenId

`func (o *CombineMarketInfo) GetQuoteTokenId() int32`

GetQuoteTokenId returns the QuoteTokenId field if non-nil, zero value otherwise.

### GetQuoteTokenIdOk

`func (o *CombineMarketInfo) GetQuoteTokenIdOk() (*int32, bool)`

GetQuoteTokenIdOk returns a tuple with the QuoteTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTokenId

`func (o *CombineMarketInfo) SetQuoteTokenId(v int32)`

SetQuoteTokenId sets QuoteTokenId field to given value.


### GetPrecisionForPrice

`func (o *CombineMarketInfo) GetPrecisionForPrice() int32`

GetPrecisionForPrice returns the PrecisionForPrice field if non-nil, zero value otherwise.

### GetPrecisionForPriceOk

`func (o *CombineMarketInfo) GetPrecisionForPriceOk() (*int32, bool)`

GetPrecisionForPriceOk returns a tuple with the PrecisionForPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecisionForPrice

`func (o *CombineMarketInfo) SetPrecisionForPrice(v int32)`

SetPrecisionForPrice sets PrecisionForPrice field to given value.


### GetOrderbookAggLevels

`func (o *CombineMarketInfo) GetOrderbookAggLevels() int32`

GetOrderbookAggLevels returns the OrderbookAggLevels field if non-nil, zero value otherwise.

### GetOrderbookAggLevelsOk

`func (o *CombineMarketInfo) GetOrderbookAggLevelsOk() (*int32, bool)`

GetOrderbookAggLevelsOk returns a tuple with the OrderbookAggLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderbookAggLevels

`func (o *CombineMarketInfo) SetOrderbookAggLevels(v int32)`

SetOrderbookAggLevels sets OrderbookAggLevels field to given value.


### GetEnabled

`func (o *CombineMarketInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CombineMarketInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CombineMarketInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetStatus

`func (o *CombineMarketInfo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CombineMarketInfo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CombineMarketInfo) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *CombineMarketInfo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CombineMarketInfo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CombineMarketInfo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


