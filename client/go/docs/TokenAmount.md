# TokenAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenSymbol** | Pointer to **string** | Token | [optional] 
**BaseOrderInfo** | Pointer to [**OrderInfo**](OrderInfo.md) |  | [optional] 
**UserOrderInfo** | Pointer to [**OrderInfo**](OrderInfo.md) |  | [optional] 
**TradeCost** | **string** | The base cost of trade settlement | 

## Methods

### NewTokenAmount

`func NewTokenAmount(tradeCost string, ) *TokenAmount`

NewTokenAmount instantiates a new TokenAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenAmountWithDefaults

`func NewTokenAmountWithDefaults() *TokenAmount`

NewTokenAmountWithDefaults instantiates a new TokenAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenSymbol

`func (o *TokenAmount) GetTokenSymbol() string`

GetTokenSymbol returns the TokenSymbol field if non-nil, zero value otherwise.

### GetTokenSymbolOk

`func (o *TokenAmount) GetTokenSymbolOk() (*string, bool)`

GetTokenSymbolOk returns a tuple with the TokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSymbol

`func (o *TokenAmount) SetTokenSymbol(v string)`

SetTokenSymbol sets TokenSymbol field to given value.

### HasTokenSymbol

`func (o *TokenAmount) HasTokenSymbol() bool`

HasTokenSymbol returns a boolean if a field has been set.

### GetBaseOrderInfo

`func (o *TokenAmount) GetBaseOrderInfo() OrderInfo`

GetBaseOrderInfo returns the BaseOrderInfo field if non-nil, zero value otherwise.

### GetBaseOrderInfoOk

`func (o *TokenAmount) GetBaseOrderInfoOk() (*OrderInfo, bool)`

GetBaseOrderInfoOk returns a tuple with the BaseOrderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseOrderInfo

`func (o *TokenAmount) SetBaseOrderInfo(v OrderInfo)`

SetBaseOrderInfo sets BaseOrderInfo field to given value.

### HasBaseOrderInfo

`func (o *TokenAmount) HasBaseOrderInfo() bool`

HasBaseOrderInfo returns a boolean if a field has been set.

### GetUserOrderInfo

`func (o *TokenAmount) GetUserOrderInfo() OrderInfo`

GetUserOrderInfo returns the UserOrderInfo field if non-nil, zero value otherwise.

### GetUserOrderInfoOk

`func (o *TokenAmount) GetUserOrderInfoOk() (*OrderInfo, bool)`

GetUserOrderInfoOk returns a tuple with the UserOrderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOrderInfo

`func (o *TokenAmount) SetUserOrderInfo(v OrderInfo)`

SetUserOrderInfo sets UserOrderInfo field to given value.

### HasUserOrderInfo

`func (o *TokenAmount) HasUserOrderInfo() bool`

HasUserOrderInfo returns a boolean if a field has been set.

### GetTradeCost

`func (o *TokenAmount) GetTradeCost() string`

GetTradeCost returns the TradeCost field if non-nil, zero value otherwise.

### GetTradeCostOk

`func (o *TokenAmount) GetTradeCostOk() (*string, bool)`

GetTradeCostOk returns a tuple with the TradeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeCost

`func (o *TokenAmount) SetTradeCost(v string)`

SetTradeCost sets TradeCost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


