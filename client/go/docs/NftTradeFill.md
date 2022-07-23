# NftTradeFill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderHash** | **string** | The order hash | 
**SellFilled** | **string** | The orders sell fill | 
**BuyFilled** | **string** | The orders buy fill | 
**Fee** | **string** | The orders fee | 

## Methods

### NewNftTradeFill

`func NewNftTradeFill(orderHash string, sellFilled string, buyFilled string, fee string, ) *NftTradeFill`

NewNftTradeFill instantiates a new NftTradeFill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftTradeFillWithDefaults

`func NewNftTradeFillWithDefaults() *NftTradeFill`

NewNftTradeFillWithDefaults instantiates a new NftTradeFill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderHash

`func (o *NftTradeFill) GetOrderHash() string`

GetOrderHash returns the OrderHash field if non-nil, zero value otherwise.

### GetOrderHashOk

`func (o *NftTradeFill) GetOrderHashOk() (*string, bool)`

GetOrderHashOk returns a tuple with the OrderHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderHash

`func (o *NftTradeFill) SetOrderHash(v string)`

SetOrderHash sets OrderHash field to given value.


### GetSellFilled

`func (o *NftTradeFill) GetSellFilled() string`

GetSellFilled returns the SellFilled field if non-nil, zero value otherwise.

### GetSellFilledOk

`func (o *NftTradeFill) GetSellFilledOk() (*string, bool)`

GetSellFilledOk returns a tuple with the SellFilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellFilled

`func (o *NftTradeFill) SetSellFilled(v string)`

SetSellFilled sets SellFilled field to given value.


### GetBuyFilled

`func (o *NftTradeFill) GetBuyFilled() string`

GetBuyFilled returns the BuyFilled field if non-nil, zero value otherwise.

### GetBuyFilledOk

`func (o *NftTradeFill) GetBuyFilledOk() (*string, bool)`

GetBuyFilledOk returns a tuple with the BuyFilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyFilled

`func (o *NftTradeFill) SetBuyFilled(v string)`

SetBuyFilled sets BuyFilled field to given value.


### GetFee

`func (o *NftTradeFill) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *NftTradeFill) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *NftTradeFill) SetFee(v string)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


