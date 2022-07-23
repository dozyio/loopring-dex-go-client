# MarketTradesV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | The total number of trades in query result | 
**Trades** | **[][]string** | List of trades records, each trade record in the order of the array is the trade time, record id, buying and selling direction, volume, price, market, fees, blockId and indexInBlock | 

## Methods

### NewMarketTradesV3

`func NewMarketTradesV3(totalNum int64, trades [][]string, ) *MarketTradesV3`

NewMarketTradesV3 instantiates a new MarketTradesV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketTradesV3WithDefaults

`func NewMarketTradesV3WithDefaults() *MarketTradesV3`

NewMarketTradesV3WithDefaults instantiates a new MarketTradesV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *MarketTradesV3) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *MarketTradesV3) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *MarketTradesV3) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetTrades

`func (o *MarketTradesV3) GetTrades() [][]string`

GetTrades returns the Trades field if non-nil, zero value otherwise.

### GetTradesOk

`func (o *MarketTradesV3) GetTradesOk() (*[][]string, bool)`

GetTradesOk returns a tuple with the Trades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrades

`func (o *MarketTradesV3) SetTrades(v [][]string)`

SetTrades sets Trades field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


