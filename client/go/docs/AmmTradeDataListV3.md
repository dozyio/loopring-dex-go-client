# AmmTradeDataListV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | Number of trades | 
**Trades** | [**[]AmmTradeDataV3**](AmmTradeDataV3.md) | AMM trade list | 

## Methods

### NewAmmTradeDataListV3

`func NewAmmTradeDataListV3(totalNum int64, trades []AmmTradeDataV3, ) *AmmTradeDataListV3`

NewAmmTradeDataListV3 instantiates a new AmmTradeDataListV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmTradeDataListV3WithDefaults

`func NewAmmTradeDataListV3WithDefaults() *AmmTradeDataListV3`

NewAmmTradeDataListV3WithDefaults instantiates a new AmmTradeDataListV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *AmmTradeDataListV3) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *AmmTradeDataListV3) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *AmmTradeDataListV3) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetTrades

`func (o *AmmTradeDataListV3) GetTrades() []AmmTradeDataV3`

GetTrades returns the Trades field if non-nil, zero value otherwise.

### GetTradesOk

`func (o *AmmTradeDataListV3) GetTradesOk() (*[]AmmTradeDataV3, bool)`

GetTradesOk returns a tuple with the Trades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrades

`func (o *AmmTradeDataListV3) SetTrades(v []AmmTradeDataV3)`

SetTrades sets Trades field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


