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

// AmmTradeData struct for AmmTradeData
type AmmTradeData struct {
	SequenceId          int64   `json:"sequenceId"`
	AccountId           int64   `json:"accountId"`
	SettlementRequestId int64   `json:"settlementRequestId"`
	OrderHash           string  `json:"orderHash"`
	Market              string  `json:"market"`
	Side                string  `json:"side"`
	Size                string  `json:"size"`
	Price               float64 `json:"price"`
	FeeAmount           string  `json:"feeAmount"`
	CreatedAt           int64   `json:"createdAt"`
}

// NewAmmTradeData instantiates a new AmmTradeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmmTradeData(sequenceId int64, accountId int64, settlementRequestId int64, orderHash string, market string, side string, size string, price float64, feeAmount string, createdAt int64) *AmmTradeData {
	this := AmmTradeData{}
	this.SequenceId = sequenceId
	this.AccountId = accountId
	this.SettlementRequestId = settlementRequestId
	this.OrderHash = orderHash
	this.Market = market
	this.Side = side
	this.Size = size
	this.Price = price
	this.FeeAmount = feeAmount
	this.CreatedAt = createdAt
	return &this
}

// NewAmmTradeDataWithDefaults instantiates a new AmmTradeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmmTradeDataWithDefaults() *AmmTradeData {
	this := AmmTradeData{}
	return &this
}

// GetSequenceId returns the SequenceId field value
func (o *AmmTradeData) GetSequenceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SequenceId
}

// GetSequenceIdOk returns a tuple with the SequenceId field value
// and a boolean to check if the value has been set.
func (o *AmmTradeData) GetSequenceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceId, true
}

// SetSequenceId sets field value
func (o *AmmTradeData) SetSequenceId(v int64) {
	o.SequenceId = v
}

// GetAccountId returns the AccountId field value
func (o *AmmTradeData) GetAccountId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AmmTradeData) GetAccountIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AmmTradeData) SetAccountId(v int64) {
	o.AccountId = v
}

// GetSettlementRequestId returns the SettlementRequestId field value
func (o *AmmTradeData) GetSettlementRequestId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SettlementRequestId
}

// GetSettlementRequestIdOk returns a tuple with the SettlementRequestId field value
// and a boolean to check if the value has been set.
func (o *AmmTradeData) GetSettlementRequestIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettlementRequestId, true
}

// SetSettlementRequestId sets field value
func (o *AmmTradeData) SetSettlementRequestId(v int64) {
	o.SettlementRequestId = v
}

// GetOrderHash returns the OrderHash field value
func (o *AmmTradeData) GetOrderHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderHash
}

// GetOrderHashOk returns a tuple with the OrderHash field value
// and a boolean to check if the value has been set.
func (o *AmmTradeData) GetOrderHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderHash, true
}

// SetOrderHash sets field value
func (o *AmmTradeData) SetOrderHash(v string) {
	o.OrderHash = v
}

// GetMarket returns the Market field value
func (o *AmmTradeData) GetMarket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *AmmTradeData) GetMarketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *AmmTradeData) SetMarket(v string) {
	o.Market = v
}

// GetSide returns the Side field value
func (o *AmmTradeData) GetSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *AmmTradeData) GetSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *AmmTradeData) SetSide(v string) {
	o.Side = v
}

// GetSize returns the Size field value
func (o *AmmTradeData) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *AmmTradeData) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *AmmTradeData) SetSize(v string) {
	o.Size = v
}

// GetPrice returns the Price field value
func (o *AmmTradeData) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *AmmTradeData) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *AmmTradeData) SetPrice(v float64) {
	o.Price = v
}

// GetFeeAmount returns the FeeAmount field value
func (o *AmmTradeData) GetFeeAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value
// and a boolean to check if the value has been set.
func (o *AmmTradeData) GetFeeAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeAmount, true
}

// SetFeeAmount sets field value
func (o *AmmTradeData) SetFeeAmount(v string) {
	o.FeeAmount = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AmmTradeData) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AmmTradeData) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AmmTradeData) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

func (o AmmTradeData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sequenceId"] = o.SequenceId
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["settlementRequestId"] = o.SettlementRequestId
	}
	if true {
		toSerialize["orderHash"] = o.OrderHash
	}
	if true {
		toSerialize["market"] = o.Market
	}
	if true {
		toSerialize["side"] = o.Side
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["feeAmount"] = o.FeeAmount
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableAmmTradeData struct {
	value *AmmTradeData
	isSet bool
}

func (v NullableAmmTradeData) Get() *AmmTradeData {
	return v.value
}

func (v *NullableAmmTradeData) Set(val *AmmTradeData) {
	v.value = val
	v.isSet = true
}

func (v NullableAmmTradeData) IsSet() bool {
	return v.isSet
}

func (v *NullableAmmTradeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmmTradeData(val *AmmTradeData) *NullableAmmTradeData {
	return &NullableAmmTradeData{value: val, isSet: true}
}

func (v NullableAmmTradeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmmTradeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}