# GetTickerResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tickers** | Pointer to **[][]string** | Each item in the list is an array that contains the following: trading pair ID, update timestamp, base token volume, quote token volume, open-price, highest price, lowest price, closing price, number of trades, highest bid price, lowest ask price, base fee amount, quote fee amount. All values are returned as strings. Fee amount is for AMM only. | [optional] 

## Methods

### NewGetTickerResponseV3

`func NewGetTickerResponseV3() *GetTickerResponseV3`

NewGetTickerResponseV3 instantiates a new GetTickerResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTickerResponseV3WithDefaults

`func NewGetTickerResponseV3WithDefaults() *GetTickerResponseV3`

NewGetTickerResponseV3WithDefaults instantiates a new GetTickerResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTickers

`func (o *GetTickerResponseV3) GetTickers() [][]string`

GetTickers returns the Tickers field if non-nil, zero value otherwise.

### GetTickersOk

`func (o *GetTickerResponseV3) GetTickersOk() (*[][]string, bool)`

GetTickersOk returns a tuple with the Tickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickers

`func (o *GetTickerResponseV3) SetTickers(v [][]string)`

SetTickers sets Tickers field to given value.

### HasTickers

`func (o *GetTickerResponseV3) HasTickers() bool`

HasTickers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


