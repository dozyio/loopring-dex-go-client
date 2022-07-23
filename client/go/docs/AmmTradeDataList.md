# AmmTradeDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** |  | 
**Trades** | [**[]AmmTradeData**](AmmTradeData.md) |  | 

## Methods

### NewAmmTradeDataList

`func NewAmmTradeDataList(totalNum int64, trades []AmmTradeData, ) *AmmTradeDataList`

NewAmmTradeDataList instantiates a new AmmTradeDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmTradeDataListWithDefaults

`func NewAmmTradeDataListWithDefaults() *AmmTradeDataList`

NewAmmTradeDataListWithDefaults instantiates a new AmmTradeDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *AmmTradeDataList) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *AmmTradeDataList) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *AmmTradeDataList) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetTrades

`func (o *AmmTradeDataList) GetTrades() []AmmTradeData`

GetTrades returns the Trades field if non-nil, zero value otherwise.

### GetTradesOk

`func (o *AmmTradeDataList) GetTradesOk() (*[]AmmTradeData, bool)`

GetTradesOk returns a tuple with the Trades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrades

`func (o *AmmTradeDataList) SetTrades(v []AmmTradeData)`

SetTrades sets Trades field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


