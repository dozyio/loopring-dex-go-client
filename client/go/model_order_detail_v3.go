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

// OrderDetailV3 Response details of order
type OrderDetailV3 struct {
	// Order hash
	Hash string `json:"hash"`
	// Order's client-side ID
	ClientOrderId string `json:"clientOrderId"`
	// Order's side
	Side string `json:"side"`
	// Trading pair
	Market string `json:"market"`
	// Order price
	Price    string          `json:"price"`
	Volumes  OrderVolumesV3  `json:"volumes"`
	Validity OrderValidityV3 `json:"validity"`
	// Whether the order has to be treated as a limit, maker, or taker operation.
	OrderType string `json:"orderType"`
	// Order channel, can be ORDER_BOOK, AMM_POOL, MIXED
	TradeChannel string `json:"tradeChannel"`
	// Order status
	Status      string       `json:"status"`
	StorageInfo *StorageInfo `json:"storageInfo,omitempty"`
}

// NewOrderDetailV3 instantiates a new OrderDetailV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailV3(hash string, clientOrderId string, side string, market string, price string, volumes OrderVolumesV3, validity OrderValidityV3, orderType string, tradeChannel string, status string) *OrderDetailV3 {
	this := OrderDetailV3{}
	this.Hash = hash
	this.ClientOrderId = clientOrderId
	this.Side = side
	this.Market = market
	this.Price = price
	this.Volumes = volumes
	this.Validity = validity
	this.OrderType = orderType
	this.TradeChannel = tradeChannel
	this.Status = status
	return &this
}

// NewOrderDetailV3WithDefaults instantiates a new OrderDetailV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailV3WithDefaults() *OrderDetailV3 {
	this := OrderDetailV3{}
	return &this
}

// GetHash returns the Hash field value
func (o *OrderDetailV3) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *OrderDetailV3) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *OrderDetailV3) SetHash(v string) {
	o.Hash = v
}

// GetClientOrderId returns the ClientOrderId field value
func (o *OrderDetailV3) GetClientOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value
// and a boolean to check if the value has been set.
func (o *OrderDetailV3) GetClientOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientOrderId, true
}

// SetClientOrderId sets field value
func (o *OrderDetailV3) SetClientOrderId(v string) {
	o.ClientOrderId = v
}

// GetSide returns the Side field value
func (o *OrderDetailV3) GetSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *OrderDetailV3) GetSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *OrderDetailV3) SetSide(v string) {
	o.Side = v
}

// GetMarket returns the Market field value
func (o *OrderDetailV3) GetMarket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Market
}

// GetMarketOk returns a tuple with the Market field value
// and a boolean to check if the value has been set.
func (o *OrderDetailV3) GetMarketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Market, true
}

// SetMarket sets field value
func (o *OrderDetailV3) SetMarket(v string) {
	o.Market = v
}

// GetPrice returns the Price field value
func (o *OrderDetailV3) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *OrderDetailV3) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *OrderDetailV3) SetPrice(v string) {
	o.Price = v
}

// GetVolumes returns the Volumes field value
func (o *OrderDetailV3) GetVolumes() OrderVolumesV3 {
	if o == nil {
		var ret OrderVolumesV3
		return ret
	}

	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value
// and a boolean to check if the value has been set.
func (o *OrderDetailV3) GetVolumesOk() (*OrderVolumesV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// SetVolumes sets field value
func (o *OrderDetailV3) SetVolumes(v OrderVolumesV3) {
	o.Volumes = v
}

// GetValidity returns the Validity field value
func (o *OrderDetailV3) GetValidity() OrderValidityV3 {
	if o == nil {
		var ret OrderValidityV3
		return ret
	}

	return o.Validity
}

// GetValidityOk returns a tuple with the Validity field value
// and a boolean to check if the value has been set.
func (o *OrderDetailV3) GetValidityOk() (*OrderValidityV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Validity, true
}

// SetValidity sets field value
func (o *OrderDetailV3) SetValidity(v OrderValidityV3) {
	o.Validity = v
}

// GetOrderType returns the OrderType field value
func (o *OrderDetailV3) GetOrderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value
// and a boolean to check if the value has been set.
func (o *OrderDetailV3) GetOrderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderType, true
}

// SetOrderType sets field value
func (o *OrderDetailV3) SetOrderType(v string) {
	o.OrderType = v
}

// GetTradeChannel returns the TradeChannel field value
func (o *OrderDetailV3) GetTradeChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeChannel
}

// GetTradeChannelOk returns a tuple with the TradeChannel field value
// and a boolean to check if the value has been set.
func (o *OrderDetailV3) GetTradeChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeChannel, true
}

// SetTradeChannel sets field value
func (o *OrderDetailV3) SetTradeChannel(v string) {
	o.TradeChannel = v
}

// GetStatus returns the Status field value
func (o *OrderDetailV3) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrderDetailV3) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrderDetailV3) SetStatus(v string) {
	o.Status = v
}

// GetStorageInfo returns the StorageInfo field value if set, zero value otherwise.
func (o *OrderDetailV3) GetStorageInfo() StorageInfo {
	if o == nil || o.StorageInfo == nil {
		var ret StorageInfo
		return ret
	}
	return *o.StorageInfo
}

// GetStorageInfoOk returns a tuple with the StorageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailV3) GetStorageInfoOk() (*StorageInfo, bool) {
	if o == nil || o.StorageInfo == nil {
		return nil, false
	}
	return o.StorageInfo, true
}

// HasStorageInfo returns a boolean if a field has been set.
func (o *OrderDetailV3) HasStorageInfo() bool {
	if o != nil && o.StorageInfo != nil {
		return true
	}

	return false
}

// SetStorageInfo gets a reference to the given StorageInfo and assigns it to the StorageInfo field.
func (o *OrderDetailV3) SetStorageInfo(v StorageInfo) {
	o.StorageInfo = &v
}

func (o OrderDetailV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if true {
		toSerialize["side"] = o.Side
	}
	if true {
		toSerialize["market"] = o.Market
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if true {
		toSerialize["volumes"] = o.Volumes
	}
	if true {
		toSerialize["validity"] = o.Validity
	}
	if true {
		toSerialize["orderType"] = o.OrderType
	}
	if true {
		toSerialize["tradeChannel"] = o.TradeChannel
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.StorageInfo != nil {
		toSerialize["storageInfo"] = o.StorageInfo
	}
	return json.Marshal(toSerialize)
}

type NullableOrderDetailV3 struct {
	value *OrderDetailV3
	isSet bool
}

func (v NullableOrderDetailV3) Get() *OrderDetailV3 {
	return v.value
}

func (v *NullableOrderDetailV3) Set(val *OrderDetailV3) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailV3) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailV3(val *OrderDetailV3) *NullableOrderDetailV3 {
	return &NullableOrderDetailV3{value: val, isSet: true}
}

func (v NullableOrderDetailV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
