# OrderDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Order hash | 
**ClientOrderId** | **string** | Order&#39;s client-side ID | 
**Size** | **string** | field.OrderDetail.size | 
**Volume** | **string** | field.OrderDetail.volume | 
**Price** | **string** | Order price | 
**FilledSize** | **string** | field.OrderDetail.filledSize | 
**FilledVolume** | **string** | field.OrderDetail.filledVolume | 
**FilledFee** | **string** | field.OrderDetail.filledFee | 
**Status** | **string** | Order status | 
**ValidUntil** | **int64** | field.OrderDetail.validUntil | 
**CreatedAt** | **int64** | field.OrderDetail.createdAt | 
**Side** | **string** | Order&#39;s side | 
**Market** | **string** | Trading pair | 
**OrderType** | **string** | Whether the order has to be treated as a limit, maker, or taker operation. | 
**TradeChannel** | **string** | field.SubmitOrderRequest.tradeChannel | 

## Methods

### NewOrderDetail

`func NewOrderDetail(hash string, clientOrderId string, size string, volume string, price string, filledSize string, filledVolume string, filledFee string, status string, validUntil int64, createdAt int64, side string, market string, orderType string, tradeChannel string, ) *OrderDetail`

NewOrderDetail instantiates a new OrderDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailWithDefaults

`func NewOrderDetailWithDefaults() *OrderDetail`

NewOrderDetailWithDefaults instantiates a new OrderDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *OrderDetail) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *OrderDetail) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *OrderDetail) SetHash(v string)`

SetHash sets Hash field to given value.


### GetClientOrderId

`func (o *OrderDetail) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OrderDetail) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OrderDetail) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.


### GetSize

`func (o *OrderDetail) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OrderDetail) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OrderDetail) SetSize(v string)`

SetSize sets Size field to given value.


### GetVolume

`func (o *OrderDetail) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *OrderDetail) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *OrderDetail) SetVolume(v string)`

SetVolume sets Volume field to given value.


### GetPrice

`func (o *OrderDetail) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderDetail) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderDetail) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetFilledSize

`func (o *OrderDetail) GetFilledSize() string`

GetFilledSize returns the FilledSize field if non-nil, zero value otherwise.

### GetFilledSizeOk

`func (o *OrderDetail) GetFilledSizeOk() (*string, bool)`

GetFilledSizeOk returns a tuple with the FilledSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledSize

`func (o *OrderDetail) SetFilledSize(v string)`

SetFilledSize sets FilledSize field to given value.


### GetFilledVolume

`func (o *OrderDetail) GetFilledVolume() string`

GetFilledVolume returns the FilledVolume field if non-nil, zero value otherwise.

### GetFilledVolumeOk

`func (o *OrderDetail) GetFilledVolumeOk() (*string, bool)`

GetFilledVolumeOk returns a tuple with the FilledVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledVolume

`func (o *OrderDetail) SetFilledVolume(v string)`

SetFilledVolume sets FilledVolume field to given value.


### GetFilledFee

`func (o *OrderDetail) GetFilledFee() string`

GetFilledFee returns the FilledFee field if non-nil, zero value otherwise.

### GetFilledFeeOk

`func (o *OrderDetail) GetFilledFeeOk() (*string, bool)`

GetFilledFeeOk returns a tuple with the FilledFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledFee

`func (o *OrderDetail) SetFilledFee(v string)`

SetFilledFee sets FilledFee field to given value.


### GetStatus

`func (o *OrderDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderDetail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetValidUntil

`func (o *OrderDetail) GetValidUntil() int64`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *OrderDetail) GetValidUntilOk() (*int64, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *OrderDetail) SetValidUntil(v int64)`

SetValidUntil sets ValidUntil field to given value.


### GetCreatedAt

`func (o *OrderDetail) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderDetail) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderDetail) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetSide

`func (o *OrderDetail) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderDetail) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderDetail) SetSide(v string)`

SetSide sets Side field to given value.


### GetMarket

`func (o *OrderDetail) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *OrderDetail) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *OrderDetail) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetOrderType

`func (o *OrderDetail) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderDetail) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderDetail) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.


### GetTradeChannel

`func (o *OrderDetail) GetTradeChannel() string`

GetTradeChannel returns the TradeChannel field if non-nil, zero value otherwise.

### GetTradeChannelOk

`func (o *OrderDetail) GetTradeChannelOk() (*string, bool)`

GetTradeChannelOk returns a tuple with the TradeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeChannel

`func (o *OrderDetail) SetTradeChannel(v string)`

SetTradeChannel sets TradeChannel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


