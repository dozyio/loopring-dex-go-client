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

// SubmitOrderRequestV3 Describes an orders structure.
type SubmitOrderRequestV3 struct {
	// The adderss of the exchange which has to process this order
	Exchange string `json:"exchange"`
	// Loopring's account ID
	AccountId int64 `json:"accountId"`
	// The unique identifier of the L2 Merkle tree storage slot where the burn made in order to exit the pool will or has been stored.
	StorageId int32         `json:"storageId"`
	SellToken TokenVolumeV3 `json:"sellToken"`
	BuyToken  TokenVolumeV3 `json:"buyToken"`
	// Whether the order supports partial fills or not.Currently only supports false as a valid value
	AllOrNone string `json:"allOrNone"`
	// Fill size by buy token or by sell token
	FillAmountBOrS string `json:"fillAmountBOrS"`
	// Order expiration time, accuracy is in seconds
	ValidUntil int64 `json:"validUntil"`
	// Maximum order fee that the user can accept, value range (in ten thousandths) 1 ~ 63
	MaxFeeBips int32 `json:"maxFeeBips"`
	// The orders EdDSA signature. The signature is a hexadecimal string obtained by signing the order itself and concatenating the resulting signature parts (Rx, Ry, and S). Used to authenticate and authorize the operation.
	EddsaSignature string `json:"eddsaSignature"`
	// An arbitrary, client-set unique order identifier, max length is 120 bytes
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	// Order types, can be AMM, LIMIT_ORDER, MAKER_ONLY, TAKER_ONLY
	OrderType *string `json:"orderType,omitempty"`
	// Order channel, can be ORDER_BOOK, AMM_POOL, MIXED
	TradeChannel *string `json:"tradeChannel,omitempty"`
	// Used by the P2P order which user specify the taker, so far its 0x0000000000000000000000000000000000000000
	Taker *string `json:"taker,omitempty"`
	// The AMM pool address if order type is AMM
	PoolAddress *string `json:"poolAddress,omitempty"`
	// An accountID who will recieve a share of the fee of this order
	Affiliate *string `json:"affiliate,omitempty"`
}

// NewSubmitOrderRequestV3 instantiates a new SubmitOrderRequestV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitOrderRequestV3(exchange string, accountId int64, storageId int32, sellToken TokenVolumeV3, buyToken TokenVolumeV3, allOrNone string, fillAmountBOrS string, validUntil int64, maxFeeBips int32, eddsaSignature string) *SubmitOrderRequestV3 {
	this := SubmitOrderRequestV3{}
	this.Exchange = exchange
	this.AccountId = accountId
	this.StorageId = storageId
	this.SellToken = sellToken
	this.BuyToken = buyToken
	this.AllOrNone = allOrNone
	this.FillAmountBOrS = fillAmountBOrS
	this.ValidUntil = validUntil
	this.MaxFeeBips = maxFeeBips
	this.EddsaSignature = eddsaSignature
	return &this
}

// NewSubmitOrderRequestV3WithDefaults instantiates a new SubmitOrderRequestV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitOrderRequestV3WithDefaults() *SubmitOrderRequestV3 {
	this := SubmitOrderRequestV3{}
	return &this
}

// GetExchange returns the Exchange field value
func (o *SubmitOrderRequestV3) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *SubmitOrderRequestV3) SetExchange(v string) {
	o.Exchange = v
}

// GetAccountId returns the AccountId field value
func (o *SubmitOrderRequestV3) GetAccountId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetAccountIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *SubmitOrderRequestV3) SetAccountId(v int64) {
	o.AccountId = v
}

// GetStorageId returns the StorageId field value
func (o *SubmitOrderRequestV3) GetStorageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetStorageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageId, true
}

// SetStorageId sets field value
func (o *SubmitOrderRequestV3) SetStorageId(v int32) {
	o.StorageId = v
}

// GetSellToken returns the SellToken field value
func (o *SubmitOrderRequestV3) GetSellToken() TokenVolumeV3 {
	if o == nil {
		var ret TokenVolumeV3
		return ret
	}

	return o.SellToken
}

// GetSellTokenOk returns a tuple with the SellToken field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetSellTokenOk() (*TokenVolumeV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellToken, true
}

// SetSellToken sets field value
func (o *SubmitOrderRequestV3) SetSellToken(v TokenVolumeV3) {
	o.SellToken = v
}

// GetBuyToken returns the BuyToken field value
func (o *SubmitOrderRequestV3) GetBuyToken() TokenVolumeV3 {
	if o == nil {
		var ret TokenVolumeV3
		return ret
	}

	return o.BuyToken
}

// GetBuyTokenOk returns a tuple with the BuyToken field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetBuyTokenOk() (*TokenVolumeV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuyToken, true
}

// SetBuyToken sets field value
func (o *SubmitOrderRequestV3) SetBuyToken(v TokenVolumeV3) {
	o.BuyToken = v
}

// GetAllOrNone returns the AllOrNone field value
func (o *SubmitOrderRequestV3) GetAllOrNone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AllOrNone
}

// GetAllOrNoneOk returns a tuple with the AllOrNone field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetAllOrNoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllOrNone, true
}

// SetAllOrNone sets field value
func (o *SubmitOrderRequestV3) SetAllOrNone(v string) {
	o.AllOrNone = v
}

// GetFillAmountBOrS returns the FillAmountBOrS field value
func (o *SubmitOrderRequestV3) GetFillAmountBOrS() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FillAmountBOrS
}

// GetFillAmountBOrSOk returns a tuple with the FillAmountBOrS field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetFillAmountBOrSOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FillAmountBOrS, true
}

// SetFillAmountBOrS sets field value
func (o *SubmitOrderRequestV3) SetFillAmountBOrS(v string) {
	o.FillAmountBOrS = v
}

// GetValidUntil returns the ValidUntil field value
func (o *SubmitOrderRequestV3) GetValidUntil() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetValidUntilOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidUntil, true
}

// SetValidUntil sets field value
func (o *SubmitOrderRequestV3) SetValidUntil(v int64) {
	o.ValidUntil = v
}

// GetMaxFeeBips returns the MaxFeeBips field value
func (o *SubmitOrderRequestV3) GetMaxFeeBips() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxFeeBips
}

// GetMaxFeeBipsOk returns a tuple with the MaxFeeBips field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetMaxFeeBipsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxFeeBips, true
}

// SetMaxFeeBips sets field value
func (o *SubmitOrderRequestV3) SetMaxFeeBips(v int32) {
	o.MaxFeeBips = v
}

// GetEddsaSignature returns the EddsaSignature field value
func (o *SubmitOrderRequestV3) GetEddsaSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EddsaSignature
}

// GetEddsaSignatureOk returns a tuple with the EddsaSignature field value
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetEddsaSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EddsaSignature, true
}

// SetEddsaSignature sets field value
func (o *SubmitOrderRequestV3) SetEddsaSignature(v string) {
	o.EddsaSignature = v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *SubmitOrderRequestV3) GetClientOrderId() string {
	if o == nil || o.ClientOrderId == nil {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetClientOrderIdOk() (*string, bool) {
	if o == nil || o.ClientOrderId == nil {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *SubmitOrderRequestV3) HasClientOrderId() bool {
	if o != nil && o.ClientOrderId != nil {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *SubmitOrderRequestV3) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *SubmitOrderRequestV3) GetOrderType() string {
	if o == nil || o.OrderType == nil {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetOrderTypeOk() (*string, bool) {
	if o == nil || o.OrderType == nil {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *SubmitOrderRequestV3) HasOrderType() bool {
	if o != nil && o.OrderType != nil {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *SubmitOrderRequestV3) SetOrderType(v string) {
	o.OrderType = &v
}

// GetTradeChannel returns the TradeChannel field value if set, zero value otherwise.
func (o *SubmitOrderRequestV3) GetTradeChannel() string {
	if o == nil || o.TradeChannel == nil {
		var ret string
		return ret
	}
	return *o.TradeChannel
}

// GetTradeChannelOk returns a tuple with the TradeChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetTradeChannelOk() (*string, bool) {
	if o == nil || o.TradeChannel == nil {
		return nil, false
	}
	return o.TradeChannel, true
}

// HasTradeChannel returns a boolean if a field has been set.
func (o *SubmitOrderRequestV3) HasTradeChannel() bool {
	if o != nil && o.TradeChannel != nil {
		return true
	}

	return false
}

// SetTradeChannel gets a reference to the given string and assigns it to the TradeChannel field.
func (o *SubmitOrderRequestV3) SetTradeChannel(v string) {
	o.TradeChannel = &v
}

// GetTaker returns the Taker field value if set, zero value otherwise.
func (o *SubmitOrderRequestV3) GetTaker() string {
	if o == nil || o.Taker == nil {
		var ret string
		return ret
	}
	return *o.Taker
}

// GetTakerOk returns a tuple with the Taker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetTakerOk() (*string, bool) {
	if o == nil || o.Taker == nil {
		return nil, false
	}
	return o.Taker, true
}

// HasTaker returns a boolean if a field has been set.
func (o *SubmitOrderRequestV3) HasTaker() bool {
	if o != nil && o.Taker != nil {
		return true
	}

	return false
}

// SetTaker gets a reference to the given string and assigns it to the Taker field.
func (o *SubmitOrderRequestV3) SetTaker(v string) {
	o.Taker = &v
}

// GetPoolAddress returns the PoolAddress field value if set, zero value otherwise.
func (o *SubmitOrderRequestV3) GetPoolAddress() string {
	if o == nil || o.PoolAddress == nil {
		var ret string
		return ret
	}
	return *o.PoolAddress
}

// GetPoolAddressOk returns a tuple with the PoolAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetPoolAddressOk() (*string, bool) {
	if o == nil || o.PoolAddress == nil {
		return nil, false
	}
	return o.PoolAddress, true
}

// HasPoolAddress returns a boolean if a field has been set.
func (o *SubmitOrderRequestV3) HasPoolAddress() bool {
	if o != nil && o.PoolAddress != nil {
		return true
	}

	return false
}

// SetPoolAddress gets a reference to the given string and assigns it to the PoolAddress field.
func (o *SubmitOrderRequestV3) SetPoolAddress(v string) {
	o.PoolAddress = &v
}

// GetAffiliate returns the Affiliate field value if set, zero value otherwise.
func (o *SubmitOrderRequestV3) GetAffiliate() string {
	if o == nil || o.Affiliate == nil {
		var ret string
		return ret
	}
	return *o.Affiliate
}

// GetAffiliateOk returns a tuple with the Affiliate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitOrderRequestV3) GetAffiliateOk() (*string, bool) {
	if o == nil || o.Affiliate == nil {
		return nil, false
	}
	return o.Affiliate, true
}

// HasAffiliate returns a boolean if a field has been set.
func (o *SubmitOrderRequestV3) HasAffiliate() bool {
	if o != nil && o.Affiliate != nil {
		return true
	}

	return false
}

// SetAffiliate gets a reference to the given string and assigns it to the Affiliate field.
func (o *SubmitOrderRequestV3) SetAffiliate(v string) {
	o.Affiliate = &v
}

func (o SubmitOrderRequestV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["exchange"] = o.Exchange
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["storageId"] = o.StorageId
	}
	if true {
		toSerialize["sellToken"] = o.SellToken
	}
	if true {
		toSerialize["buyToken"] = o.BuyToken
	}
	if true {
		toSerialize["allOrNone"] = o.AllOrNone
	}
	if true {
		toSerialize["fillAmountBOrS"] = o.FillAmountBOrS
	}
	if true {
		toSerialize["validUntil"] = o.ValidUntil
	}
	if true {
		toSerialize["maxFeeBips"] = o.MaxFeeBips
	}
	if true {
		toSerialize["eddsaSignature"] = o.EddsaSignature
	}
	if o.ClientOrderId != nil {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if o.OrderType != nil {
		toSerialize["orderType"] = o.OrderType
	}
	if o.TradeChannel != nil {
		toSerialize["tradeChannel"] = o.TradeChannel
	}
	if o.Taker != nil {
		toSerialize["taker"] = o.Taker
	}
	if o.PoolAddress != nil {
		toSerialize["poolAddress"] = o.PoolAddress
	}
	if o.Affiliate != nil {
		toSerialize["affiliate"] = o.Affiliate
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitOrderRequestV3 struct {
	value *SubmitOrderRequestV3
	isSet bool
}

func (v NullableSubmitOrderRequestV3) Get() *SubmitOrderRequestV3 {
	return v.value
}

func (v *NullableSubmitOrderRequestV3) Set(val *SubmitOrderRequestV3) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitOrderRequestV3) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitOrderRequestV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitOrderRequestV3(val *SubmitOrderRequestV3) *NullableSubmitOrderRequestV3 {
	return &NullableSubmitOrderRequestV3{value: val, isSet: true}
}

func (v NullableSubmitOrderRequestV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitOrderRequestV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}