# NftTradeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MakerFills** | [**NftTradeFill**](NftTradeFill.md) |  | 
**TakerFills** | [**NftTradeFill**](NftTradeFill.md) |  | 
**TradeHash** | **string** | The trade hash which can be queried in loopring scan web. | 

## Methods

### NewNftTradeResponse

`func NewNftTradeResponse(makerFills NftTradeFill, takerFills NftTradeFill, tradeHash string, ) *NftTradeResponse`

NewNftTradeResponse instantiates a new NftTradeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftTradeResponseWithDefaults

`func NewNftTradeResponseWithDefaults() *NftTradeResponse`

NewNftTradeResponseWithDefaults instantiates a new NftTradeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMakerFills

`func (o *NftTradeResponse) GetMakerFills() NftTradeFill`

GetMakerFills returns the MakerFills field if non-nil, zero value otherwise.

### GetMakerFillsOk

`func (o *NftTradeResponse) GetMakerFillsOk() (*NftTradeFill, bool)`

GetMakerFillsOk returns a tuple with the MakerFills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerFills

`func (o *NftTradeResponse) SetMakerFills(v NftTradeFill)`

SetMakerFills sets MakerFills field to given value.


### GetTakerFills

`func (o *NftTradeResponse) GetTakerFills() NftTradeFill`

GetTakerFills returns the TakerFills field if non-nil, zero value otherwise.

### GetTakerFillsOk

`func (o *NftTradeResponse) GetTakerFillsOk() (*NftTradeFill, bool)`

GetTakerFillsOk returns a tuple with the TakerFills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerFills

`func (o *NftTradeResponse) SetTakerFills(v NftTradeFill)`

SetTakerFills sets TakerFills field to given value.


### GetTradeHash

`func (o *NftTradeResponse) GetTradeHash() string`

GetTradeHash returns the TradeHash field if non-nil, zero value otherwise.

### GetTradeHashOk

`func (o *NftTradeResponse) GetTradeHashOk() (*string, bool)`

GetTradeHashOk returns a tuple with the TradeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeHash

`func (o *NftTradeResponse) SetTradeHash(v string)`

SetTradeHash sets TradeHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


