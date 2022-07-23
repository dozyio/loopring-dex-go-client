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

// CombineMarketInfo struct for CombineMarketInfo
type CombineMarketInfo struct {
	// Trading pair ID
	Market string `json:"market"`
	// The base token ID
	BaseTokenId int32 `json:"baseTokenId"`
	// The quote token ID
	QuoteTokenId int32 `json:"quoteTokenId"`
	// The precision of price
	PrecisionForPrice int32 `json:"precisionForPrice"`
	// The max level of orderbook price aggregation
	OrderbookAggLevels int32 `json:"orderbookAggLevels"`
	// True if trading is enabled for this trading pair
	Enabled bool `json:"enabled"`
	// field.CombineMarketInfo.status
	Status int32 `json:"status"`
	// field.AmmMarketInfo.createdAt
	CreatedAt string `json:"createdAt"`
}

// NewCombineMarketInfo instantiates a new CombineMarketInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCombineMarketInfo(market string, baseTokenId int32, quoteTokenId int32, precisionForPrice int32, orderbookAggLevels int32, enabled bool, status int32, createdAt string) *CombineMarketInfo {
	this := CombineMarketInfo{}
	this.Market = market
	this.BaseTokenId = baseTokenId
	this.QuoteTokenId = quoteTokenId
	this.PrecisionForPrice = precisionForPrice
	this.OrderbookAggLevels = orderbookAggLevels
	this.Enabled = enabled
	this.Status = status
	this.CreatedAt = createdAt
	return &this
}

// NewCombineMarketInfoWithDefaults instantiates a new CombineMarketInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCombineMarketInfoWithDefaults() *CombineMarketInfo {
	this := CombineMarketInfo{}
	return &this
}

// GetMarket returns the Market field value
func (o *CombineMarketInfo) GetMarket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *CombineMarketInfo) GetMarketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *CombineMarketInfo) SetMarket(v string) {
	o.Market = v
}

// GetBaseTokenId returns the BaseTokenId field value
func (o *CombineMarketInfo) GetBaseTokenId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BaseTokenId
}

// GetBaseTokenIdOk returns a tuple with the BaseTokenId field value
// and a boolean to check if the value has been set.
func (o *CombineMarketInfo) GetBaseTokenIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseTokenId, true
}

// SetBaseTokenId sets field value
func (o *CombineMarketInfo) SetBaseTokenId(v int32) {
	o.BaseTokenId = v
}

// GetQuoteTokenId returns the QuoteTokenId field value
func (o *CombineMarketInfo) GetQuoteTokenId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QuoteTokenId
}

// GetQuoteTokenIdOk returns a tuple with the QuoteTokenId field value
// and a boolean to check if the value has been set.
func (o *CombineMarketInfo) GetQuoteTokenIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteTokenId, true
}

// SetQuoteTokenId sets field value
func (o *CombineMarketInfo) SetQuoteTokenId(v int32) {
	o.QuoteTokenId = v
}

// GetPrecisionForPrice returns the PrecisionForPrice field value
func (o *CombineMarketInfo) GetPrecisionForPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrecisionForPrice
}

// GetPrecisionForPriceOk returns a tuple with the PrecisionForPrice field value
// and a boolean to check if the value has been set.
func (o *CombineMarketInfo) GetPrecisionForPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrecisionForPrice, true
}

// SetPrecisionForPrice sets field value
func (o *CombineMarketInfo) SetPrecisionForPrice(v int32) {
	o.PrecisionForPrice = v
}

// GetOrderbookAggLevels returns the OrderbookAggLevels field value
func (o *CombineMarketInfo) GetOrderbookAggLevels() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderbookAggLevels
}

// GetOrderbookAggLevelsOk returns a tuple with the OrderbookAggLevels field value
// and a boolean to check if the value has been set.
func (o *CombineMarketInfo) GetOrderbookAggLevelsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderbookAggLevels, true
}

// SetOrderbookAggLevels sets field value
func (o *CombineMarketInfo) SetOrderbookAggLevels(v int32) {
	o.OrderbookAggLevels = v
}

// GetEnabled returns the Enabled field value
func (o *CombineMarketInfo) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CombineMarketInfo) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CombineMarketInfo) SetEnabled(v bool) {
	o.Enabled = v
}

// GetStatus returns the Status field value
func (o *CombineMarketInfo) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CombineMarketInfo) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CombineMarketInfo) SetStatus(v int32) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CombineMarketInfo) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CombineMarketInfo) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CombineMarketInfo) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o CombineMarketInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["market"] = o.Market
	}
	if true {
		toSerialize["baseTokenId"] = o.BaseTokenId
	}
	if true {
		toSerialize["quoteTokenId"] = o.QuoteTokenId
	}
	if true {
		toSerialize["precisionForPrice"] = o.PrecisionForPrice
	}
	if true {
		toSerialize["orderbookAggLevels"] = o.OrderbookAggLevels
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableCombineMarketInfo struct {
	value *CombineMarketInfo
	isSet bool
}

func (v NullableCombineMarketInfo) Get() *CombineMarketInfo {
	return v.value
}

func (v *NullableCombineMarketInfo) Set(val *CombineMarketInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCombineMarketInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCombineMarketInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCombineMarketInfo(val *CombineMarketInfo) *NullableCombineMarketInfo {
	return &NullableCombineMarketInfo{value: val, isSet: true}
}

func (v NullableCombineMarketInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCombineMarketInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
