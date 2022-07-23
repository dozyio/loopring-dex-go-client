/*
LightCone 2.0 API Documentation

LightCone DEX function interpretation

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loopring

import (
	"encoding/json"
)

// OrderDetail Response details of order
type OrderDetail struct {
	// Order hash
	Hash string `json:"hash"`
	// Order's client-side ID
	ClientOrderId string `json:"clientOrderId"`
	// field.OrderDetail.size
	Size string `json:"size"`
	// field.OrderDetail.volume
	Volume string `json:"volume"`
	// Order price
	Price string `json:"price"`
	// field.OrderDetail.filledSize
	FilledSize string `json:"filledSize"`
	// field.OrderDetail.filledVolume
	FilledVolume string `json:"filledVolume"`
	// field.OrderDetail.filledFee
	FilledFee string `json:"filledFee"`
	// Order status
	Status string `json:"status"`
	// field.OrderDetail.validUntil
	ValidUntil int64 `json:"validUntil"`
	// field.OrderDetail.createdAt
	CreatedAt int64 `json:"createdAt"`
	// Order's side
	Side string `json:"side"`
	// Trading pair
	Market string `json:"market"`
	// Whether the order has to be treated as a limit, maker, or taker operation.
	OrderType string `json:"orderType"`
	// field.SubmitOrderRequest.tradeChannel
	TradeChannel string `json:"tradeChannel"`
}

// NewOrderDetail instantiates a new OrderDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetail(hash string, clientOrderId string, size string, volume string, price string, filledSize string, filledVolume string, filledFee string, status string, validUntil int64, createdAt int64, side string, market string, orderType string, tradeChannel string) *OrderDetail {
	this := OrderDetail{}
	this.Hash = hash
	this.ClientOrderId = clientOrderId
	this.Size = size
	this.Volume = volume
	this.Price = price
	this.FilledSize = filledSize
	this.FilledVolume = filledVolume
	this.FilledFee = filledFee
	this.Status = status
	this.ValidUntil = validUntil
	this.CreatedAt = createdAt
	this.Side = side
	this.Market = market
	this.OrderType = orderType
	this.TradeChannel = tradeChannel
	return &this
}

// NewOrderDetailWithDefaults instantiates a new OrderDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailWithDefaults() *OrderDetail {
	this := OrderDetail{}
	return &this
}

// GetHash returns the Hash field value
func (o *OrderDetail) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *OrderDetail) SetHash(v string) {
	o.Hash = v
}

// GetClientOrderId returns the ClientOrderId field value
func (o *OrderDetail) GetClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetClientOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientOrderId, true
}

// SetClientOrderId sets field value
func (o *OrderDetail) SetClientOrderId(v string) {
	o.ClientOrderId = v
}

// GetSize returns the Size field value
func (o *OrderDetail) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *OrderDetail) SetSize(v string) {
	o.Size = v
}

// GetVolume returns the Volume field value
func (o *OrderDetail) GetVolume() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetVolumeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volume, true
}

// SetVolume sets field value
func (o *OrderDetail) SetVolume(v string) {
	o.Volume = v
}

// GetPrice returns the Price field value
func (o *OrderDetail) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *OrderDetail) SetPrice(v string) {
	o.Price = v
}

// GetFilledSize returns the FilledSize field value
func (o *OrderDetail) GetFilledSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilledSize
}

// GetFilledSizeOk returns a tuple with the FilledSize field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetFilledSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilledSize, true
}

// SetFilledSize sets field value
func (o *OrderDetail) SetFilledSize(v string) {
	o.FilledSize = v
}

// GetFilledVolume returns the FilledVolume field value
func (o *OrderDetail) GetFilledVolume() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilledVolume
}

// GetFilledVolumeOk returns a tuple with the FilledVolume field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetFilledVolumeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilledVolume, true
}

// SetFilledVolume sets field value
func (o *OrderDetail) SetFilledVolume(v string) {
	o.FilledVolume = v
}

// GetFilledFee returns the FilledFee field value
func (o *OrderDetail) GetFilledFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilledFee
}

// GetFilledFeeOk returns a tuple with the FilledFee field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetFilledFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilledFee, true
}

// SetFilledFee sets field value
func (o *OrderDetail) SetFilledFee(v string) {
	o.FilledFee = v
}

// GetStatus returns the Status field value
func (o *OrderDetail) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrderDetail) SetStatus(v string) {
	o.Status = v
}

// GetValidUntil returns the ValidUntil field value
func (o *OrderDetail) GetValidUntil() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetValidUntilOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidUntil, true
}

// SetValidUntil sets field value
func (o *OrderDetail) SetValidUntil(v int64) {
	o.ValidUntil = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrderDetail) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrderDetail) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetSide returns the Side field value
func (o *OrderDetail) GetSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *OrderDetail) SetSide(v string) {
	o.Side = v
}

// GetMarket returns the Market field value
func (o *OrderDetail) GetMarket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetMarketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *OrderDetail) SetMarket(v string) {
	o.Market = v
}

// GetOrderType returns the OrderType field value
func (o *OrderDetail) GetOrderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetOrderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderType, true
}

// SetOrderType sets field value
func (o *OrderDetail) SetOrderType(v string) {
	o.OrderType = v
}

// GetTradeChannel returns the TradeChannel field value
func (o *OrderDetail) GetTradeChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeChannel
}

// GetTradeChannelOk returns a tuple with the TradeChannel field value
// and a boolean to check if the value has been set.
func (o *OrderDetail) GetTradeChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeChannel, true
}

// SetTradeChannel sets field value
func (o *OrderDetail) SetTradeChannel(v string) {
	o.TradeChannel = v
}

func (o OrderDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["volume"] = o.Volume
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["filledSize"] = o.FilledSize
	}
	if true {
		toSerialize["filledVolume"] = o.FilledVolume
	}
	if true {
		toSerialize["filledFee"] = o.FilledFee
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["validUntil"] = o.ValidUntil
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["side"] = o.Side
	}
	if true {
		toSerialize["market"] = o.Market
	}
	if true {
		toSerialize["orderType"] = o.OrderType
	}
	if true {
		toSerialize["tradeChannel"] = o.TradeChannel
	}
	return json.Marshal(toSerialize)
}

type NullableOrderDetail struct {
	value *OrderDetail
	isSet bool
}

func (v NullableOrderDetail) Get() *OrderDetail {
	return v.value
}

func (v *NullableOrderDetail) Set(val *OrderDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetail(val *OrderDetail) *NullableOrderDetail {
	return &NullableOrderDetail{value: val, isSet: true}
}

func (v NullableOrderDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}