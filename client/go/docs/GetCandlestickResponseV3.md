# GetCandlestickResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Candlesticks** | Pointer to **[][]string** | Candlestick data, each set of data includes start time, number of transactions, opening price, closing price, highest price, lowest price, total transaction volume of Base Token, total transaction volume of Quote Token | [optional] 

## Methods

### NewGetCandlestickResponseV3

`func NewGetCandlestickResponseV3() *GetCandlestickResponseV3`

NewGetCandlestickResponseV3 instantiates a new GetCandlestickResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCandlestickResponseV3WithDefaults

`func NewGetCandlestickResponseV3WithDefaults() *GetCandlestickResponseV3`

NewGetCandlestickResponseV3WithDefaults instantiates a new GetCandlestickResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandlesticks

`func (o *GetCandlestickResponseV3) GetCandlesticks() [][]string`

GetCandlesticks returns the Candlesticks field if non-nil, zero value otherwise.

### GetCandlesticksOk

`func (o *GetCandlestickResponseV3) GetCandlesticksOk() (*[][]string, bool)`

GetCandlesticksOk returns a tuple with the Candlesticks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandlesticks

`func (o *GetCandlestickResponseV3) SetCandlesticks(v [][]string)`

SetCandlesticks sets Candlesticks field to given value.

### HasCandlesticks

`func (o *GetCandlestickResponseV3) HasCandlesticks() bool`

HasCandlesticks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


