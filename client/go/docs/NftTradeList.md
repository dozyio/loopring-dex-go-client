# NftTradeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | field.totalNum.description | 
**Trades** | **[][]string** | Trade info | 

## Methods

### NewNftTradeList

`func NewNftTradeList(totalNum int64, trades [][]string, ) *NftTradeList`

NewNftTradeList instantiates a new NftTradeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftTradeListWithDefaults

`func NewNftTradeListWithDefaults() *NftTradeList`

NewNftTradeListWithDefaults instantiates a new NftTradeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *NftTradeList) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *NftTradeList) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *NftTradeList) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetTrades

`func (o *NftTradeList) GetTrades() [][]string`

GetTrades returns the Trades field if non-nil, zero value otherwise.

### GetTradesOk

`func (o *NftTradeList) GetTradesOk() (*[][]string, bool)`

GetTradesOk returns a tuple with the Trades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrades

`func (o *NftTradeList) SetTrades(v [][]string)`

SetTrades sets Trades field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


