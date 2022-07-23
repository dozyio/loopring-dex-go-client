# TradeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | Total number of tradings | 
**Trades** | **[][]string** | field.marketTrades.trades | 

## Methods

### NewTradeList

`func NewTradeList(totalNum int64, trades [][]string, ) *TradeList`

NewTradeList instantiates a new TradeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeListWithDefaults

`func NewTradeListWithDefaults() *TradeList`

NewTradeListWithDefaults instantiates a new TradeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *TradeList) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *TradeList) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *TradeList) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetTrades

`func (o *TradeList) GetTrades() [][]string`

GetTrades returns the Trades field if non-nil, zero value otherwise.

### GetTradesOk

`func (o *TradeList) GetTradesOk() (*[][]string, bool)`

GetTradesOk returns a tuple with the Trades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrades

`func (o *TradeList) SetTrades(v [][]string)`

SetTrades sets Trades field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


