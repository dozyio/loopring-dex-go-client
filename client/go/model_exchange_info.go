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

// ExchangeInfo field.exchangeInfo.description
type ExchangeInfo struct {
	// Loopring's smart contract network ID.
	ChainId int32 `json:"chainId"`
	// Contract address of exchange.
	ExchangeAddress string `json:"exchangeAddress"`
	// field.ExchangeInfo.depositAddress
	DepositAddress string `json:"depositAddress"`
	// Fees settings.
	OnchainFees []FeeInfo `json:"onchainFees"`
	// field.ExchangeInfo.openAccountFee
	OpenAccountFees []OffFeeInfo `json:"openAccountFees"`
	// field.ExchangeInfo.updateFees
	UpdateFees []OffFeeInfo `json:"updateFees"`
	// Transfer fee settings.
	TransferFees []OffFeeInfo `json:"transferFees"`
	// Off-chain withdrawal fee settings.
	WithdrawalFees []OffFeeInfo `json:"withdrawalFees"`
	// fast withdrawal fee settings.
	FastWithdrawalFees []OffFeeInfo `json:"fastWithdrawalFees"`
	// AMM pool exit fee settings.
	AmmExitFees []OffFeeInfo `json:"ammExitFees"`
}

// NewExchangeInfo instantiates a new ExchangeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInfo(chainId int32, exchangeAddress string, depositAddress string, onchainFees []FeeInfo, openAccountFees []OffFeeInfo, updateFees []OffFeeInfo, transferFees []OffFeeInfo, withdrawalFees []OffFeeInfo, fastWithdrawalFees []OffFeeInfo, ammExitFees []OffFeeInfo) *ExchangeInfo {
	this := ExchangeInfo{}
	this.ChainId = chainId
	this.ExchangeAddress = exchangeAddress
	this.DepositAddress = depositAddress
	this.OnchainFees = onchainFees
	this.OpenAccountFees = openAccountFees
	this.UpdateFees = updateFees
	this.TransferFees = transferFees
	this.WithdrawalFees = withdrawalFees
	this.FastWithdrawalFees = fastWithdrawalFees
	this.AmmExitFees = ammExitFees
	return &this
}

// NewExchangeInfoWithDefaults instantiates a new ExchangeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInfoWithDefaults() *ExchangeInfo {
	this := ExchangeInfo{}
	return &this
}

// GetChainId returns the ChainId field value
func (o *ExchangeInfo) GetChainId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *ExchangeInfo) GetChainIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *ExchangeInfo) SetChainId(v int32) {
	o.ChainId = v
}

// GetExchangeAddress returns the ExchangeAddress field value
func (o *ExchangeInfo) GetExchangeAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeAddress
}

// GetExchangeAddressOk returns a tuple with the ExchangeAddress field value
// and a boolean to check if the value has been set.
func (o *ExchangeInfo) GetExchangeAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeAddress, true
}

// SetExchangeAddress sets field value
func (o *ExchangeInfo) SetExchangeAddress(v string) {
	o.ExchangeAddress = v
}

// GetDepositAddress returns the DepositAddress field value
func (o *ExchangeInfo) GetDepositAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositAddress
}

// GetDepositAddressOk returns a tuple with the DepositAddress field value
// and a boolean to check if the value has been set.
func (o *ExchangeInfo) GetDepositAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositAddress, true
}

// SetDepositAddress sets field value
func (o *ExchangeInfo) SetDepositAddress(v string) {
	o.DepositAddress = v
}

// GetOnchainFees returns the OnchainFees field value
func (o *ExchangeInfo) GetOnchainFees() []FeeInfo {
	if o == nil {
		var ret []FeeInfo
		return ret
	}

	return o.OnchainFees
}

// GetOnchainFeesOk returns a tuple with the OnchainFees field value
// and a boolean to check if the value has been set.
func (o *ExchangeInfo) GetOnchainFeesOk() ([]FeeInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnchainFees, true
}

// SetOnchainFees sets field value
func (o *ExchangeInfo) SetOnchainFees(v []FeeInfo) {
	o.OnchainFees = v
}

// GetOpenAccountFees returns the OpenAccountFees field value
func (o *ExchangeInfo) GetOpenAccountFees() []OffFeeInfo {
	if o == nil {
		var ret []OffFeeInfo
		return ret
	}

	return o.OpenAccountFees
}

// GetOpenAccountFeesOk returns a tuple with the OpenAccountFees field value
// and a boolean to check if the value has been set.
func (o *ExchangeInfo) GetOpenAccountFeesOk() ([]OffFeeInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenAccountFees, true
}

// SetOpenAccountFees sets field value
func (o *ExchangeInfo) SetOpenAccountFees(v []OffFeeInfo) {
	o.OpenAccountFees = v
}

// GetUpdateFees returns the UpdateFees field value
func (o *ExchangeInfo) GetUpdateFees() []OffFeeInfo {
	if o == nil {
		var ret []OffFeeInfo
		return ret
	}

	return o.UpdateFees
}

// GetUpdateFeesOk returns a tuple with the UpdateFees field value
// and a boolean to check if the value has been set.
func (o *ExchangeInfo) GetUpdateFeesOk() ([]OffFeeInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdateFees, true
}

// SetUpdateFees sets field value
func (o *ExchangeInfo) SetUpdateFees(v []OffFeeInfo) {
	o.UpdateFees = v
}

// GetTransferFees returns the TransferFees field value
func (o *ExchangeInfo) GetTransferFees() []OffFeeInfo {
	if o == nil {
		var ret []OffFeeInfo
		return ret
	}

	return o.TransferFees
}

// GetTransferFeesOk returns a tuple with the TransferFees field value
// and a boolean to check if the value has been set.
func (o *ExchangeInfo) GetTransferFeesOk() ([]OffFeeInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransferFees, true
}

// SetTransferFees sets field value
func (o *ExchangeInfo) SetTransferFees(v []OffFeeInfo) {
	o.TransferFees = v
}

// GetWithdrawalFees returns the WithdrawalFees field value
func (o *ExchangeInfo) GetWithdrawalFees() []OffFeeInfo {
	if o == nil {
		var ret []OffFeeInfo
		return ret
	}

	return o.WithdrawalFees
}

// GetWithdrawalFeesOk returns a tuple with the WithdrawalFees field value
// and a boolean to check if the value has been set.
func (o *ExchangeInfo) GetWithdrawalFeesOk() ([]OffFeeInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithdrawalFees, true
}

// SetWithdrawalFees sets field value
func (o *ExchangeInfo) SetWithdrawalFees(v []OffFeeInfo) {
	o.WithdrawalFees = v
}

// GetFastWithdrawalFees returns the FastWithdrawalFees field value
func (o *ExchangeInfo) GetFastWithdrawalFees() []OffFeeInfo {
	if o == nil {
		var ret []OffFeeInfo
		return ret
	}

	return o.FastWithdrawalFees
}

// GetFastWithdrawalFeesOk returns a tuple with the FastWithdrawalFees field value
// and a boolean to check if the value has been set.
func (o *ExchangeInfo) GetFastWithdrawalFeesOk() ([]OffFeeInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.FastWithdrawalFees, true
}

// SetFastWithdrawalFees sets field value
func (o *ExchangeInfo) SetFastWithdrawalFees(v []OffFeeInfo) {
	o.FastWithdrawalFees = v
}

// GetAmmExitFees returns the AmmExitFees field value
func (o *ExchangeInfo) GetAmmExitFees() []OffFeeInfo {
	if o == nil {
		var ret []OffFeeInfo
		return ret
	}

	return o.AmmExitFees
}

// GetAmmExitFeesOk returns a tuple with the AmmExitFees field value
// and a boolean to check if the value has been set.
func (o *ExchangeInfo) GetAmmExitFeesOk() ([]OffFeeInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmmExitFees, true
}

// SetAmmExitFees sets field value
func (o *ExchangeInfo) SetAmmExitFees(v []OffFeeInfo) {
	o.AmmExitFees = v
}

func (o ExchangeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["chainId"] = o.ChainId
	}
	if true {
		toSerialize["exchangeAddress"] = o.ExchangeAddress
	}
	if true {
		toSerialize["depositAddress"] = o.DepositAddress
	}
	if true {
		toSerialize["onchainFees"] = o.OnchainFees
	}
	if true {
		toSerialize["openAccountFees"] = o.OpenAccountFees
	}
	if true {
		toSerialize["updateFees"] = o.UpdateFees
	}
	if true {
		toSerialize["transferFees"] = o.TransferFees
	}
	if true {
		toSerialize["withdrawalFees"] = o.WithdrawalFees
	}
	if true {
		toSerialize["fastWithdrawalFees"] = o.FastWithdrawalFees
	}
	if true {
		toSerialize["ammExitFees"] = o.AmmExitFees
	}
	return json.Marshal(toSerialize)
}

type NullableExchangeInfo struct {
	value *ExchangeInfo
	isSet bool
}

func (v NullableExchangeInfo) Get() *ExchangeInfo {
	return v.value
}

func (v *NullableExchangeInfo) Set(val *ExchangeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInfo(val *ExchangeInfo) *NullableExchangeInfo {
	return &NullableExchangeInfo{value: val, isSet: true}
}

func (v NullableExchangeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
