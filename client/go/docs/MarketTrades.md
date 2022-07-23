# MarketTrades

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | The total number of trades in query result | 
**Trades** | **[][]string** | List of trades records, each trade record in the order of the array is the trade time, record id, buying and selling direction, volume, price, market, fees, blockId and indexInBlock | 

## Methods

### NewMarketTrades

`func NewMarketTrades(totalNum int64, trades [][]string, ) *MarketTrades`

NewMarketTrades instantiates a new MarketTrades object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketTradesWithDefaults

`func NewMarketTradesWithDefaults() *MarketTrades`

NewMarketTradesWithDefaults instantiates a new MarketTrades object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *MarketTrades) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *MarketTrades) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *MarketTrades) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetTrades

`func (o *MarketTrades) GetTrades() [][]string`

GetTrades returns the Trades field if non-nil, zero value otherwise.

### GetTradesOk

`func (o *MarketTrades) GetTradesOk() (*[][]string, bool)`

GetTradesOk returns a tuple with the Trades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrades

`func (o *MarketTrades) SetTrades(v [][]string)`

SetTrades sets Trades field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


