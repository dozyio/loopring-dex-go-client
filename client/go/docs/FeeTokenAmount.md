# FeeTokenAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeTokenSymbol** | **string** | Fee token symbol | 
**MinAmount** | **string** | The minimum quote token amount to place a NFT order | 
**TradeCost** | **string** | The base cost of trade settlement | 
**MarketOrderInfo** | [**OrderAmountsV3**](OrderAmountsV3.md) |  | 

## Methods

### NewFeeTokenAmount

`func NewFeeTokenAmount(feeTokenSymbol string, minAmount string, tradeCost string, marketOrderInfo OrderAmountsV3, ) *FeeTokenAmount`

NewFeeTokenAmount instantiates a new FeeTokenAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeTokenAmountWithDefaults

`func NewFeeTokenAmountWithDefaults() *FeeTokenAmount`

NewFeeTokenAmountWithDefaults instantiates a new FeeTokenAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeTokenSymbol

`func (o *FeeTokenAmount) GetFeeTokenSymbol() string`

GetFeeTokenSymbol returns the FeeTokenSymbol field if non-nil, zero value otherwise.

### GetFeeTokenSymbolOk

`func (o *FeeTokenAmount) GetFeeTokenSymbolOk() (*string, bool)`

GetFeeTokenSymbolOk returns a tuple with the FeeTokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenSymbol

`func (o *FeeTokenAmount) SetFeeTokenSymbol(v string)`

SetFeeTokenSymbol sets FeeTokenSymbol field to given value.


### GetMinAmount

`func (o *FeeTokenAmount) GetMinAmount() string`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *FeeTokenAmount) GetMinAmountOk() (*string, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *FeeTokenAmount) SetMinAmount(v string)`

SetMinAmount sets MinAmount field to given value.


### GetTradeCost

`func (o *FeeTokenAmount) GetTradeCost() string`

GetTradeCost returns the TradeCost field if non-nil, zero value otherwise.

### GetTradeCostOk

`func (o *FeeTokenAmount) GetTradeCostOk() (*string, bool)`

GetTradeCostOk returns a tuple with the TradeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeCost

`func (o *FeeTokenAmount) SetTradeCost(v string)`

SetTradeCost sets TradeCost field to given value.


### GetMarketOrderInfo

`func (o *FeeTokenAmount) GetMarketOrderInfo() OrderAmountsV3`

GetMarketOrderInfo returns the MarketOrderInfo field if non-nil, zero value otherwise.

### GetMarketOrderInfoOk

`func (o *FeeTokenAmount) GetMarketOrderInfoOk() (*OrderAmountsV3, bool)`

GetMarketOrderInfoOk returns a tuple with the MarketOrderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketOrderInfo

`func (o *FeeTokenAmount) SetMarketOrderInfo(v OrderAmountsV3)`

SetMarketOrderInfo sets MarketOrderInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


