# MarketInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | **string** | Trading pair ID | 
**BaseTokenId** | **int32** | The base token ID | 
**QuoteTokenId** | **int32** | The quote token ID | 
**PrecisionForPrice** | **int32** | The precision of price | 
**OrderbookAggLevels** | **int32** | The max level of orderbook price aggregation | 
**Enabled** | **bool** | True if trading is enabled for this trading pair | 

## Methods

### NewMarketInfo

`func NewMarketInfo(market string, baseTokenId int32, quoteTokenId int32, precisionForPrice int32, orderbookAggLevels int32, enabled bool, ) *MarketInfo`

NewMarketInfo instantiates a new MarketInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketInfoWithDefaults

`func NewMarketInfoWithDefaults() *MarketInfo`

NewMarketInfoWithDefaults instantiates a new MarketInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *MarketInfo) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *MarketInfo) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *MarketInfo) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetBaseTokenId

`func (o *MarketInfo) GetBaseTokenId() int32`

GetBaseTokenId returns the BaseTokenId field if non-nil, zero value otherwise.

### GetBaseTokenIdOk

`func (o *MarketInfo) GetBaseTokenIdOk() (*int32, bool)`

GetBaseTokenIdOk returns a tuple with the BaseTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseTokenId

`func (o *MarketInfo) SetBaseTokenId(v int32)`

SetBaseTokenId sets BaseTokenId field to given value.


### GetQuoteTokenId

`func (o *MarketInfo) GetQuoteTokenId() int32`

GetQuoteTokenId returns the QuoteTokenId field if non-nil, zero value otherwise.

### GetQuoteTokenIdOk

`func (o *MarketInfo) GetQuoteTokenIdOk() (*int32, bool)`

GetQuoteTokenIdOk returns a tuple with the QuoteTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTokenId

`func (o *MarketInfo) SetQuoteTokenId(v int32)`

SetQuoteTokenId sets QuoteTokenId field to given value.


### GetPrecisionForPrice

`func (o *MarketInfo) GetPrecisionForPrice() int32`

GetPrecisionForPrice returns the PrecisionForPrice field if non-nil, zero value otherwise.

### GetPrecisionForPriceOk

`func (o *MarketInfo) GetPrecisionForPriceOk() (*int32, bool)`

GetPrecisionForPriceOk returns a tuple with the PrecisionForPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecisionForPrice

`func (o *MarketInfo) SetPrecisionForPrice(v int32)`

SetPrecisionForPrice sets PrecisionForPrice field to given value.


### GetOrderbookAggLevels

`func (o *MarketInfo) GetOrderbookAggLevels() int32`

GetOrderbookAggLevels returns the OrderbookAggLevels field if non-nil, zero value otherwise.

### GetOrderbookAggLevelsOk

`func (o *MarketInfo) GetOrderbookAggLevelsOk() (*int32, bool)`

GetOrderbookAggLevelsOk returns a tuple with the OrderbookAggLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderbookAggLevels

`func (o *MarketInfo) SetOrderbookAggLevels(v int32)`

SetOrderbookAggLevels sets OrderbookAggLevels field to given value.


### GetEnabled

`func (o *MarketInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MarketInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MarketInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


