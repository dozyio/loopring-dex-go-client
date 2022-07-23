# TokenInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | field.tokenInfo.type | 
**TokenId** | **int32** | field.tokenInfo.tokenId | 
**Symbol** | **string** | field.tokenInfo.symbol | 
**Name** | **string** | field.tokenInfo.name | 
**Address** | **string** | field.tokenInfo.address | 
**Decimals** | **int32** | field.tokenInfo.decimals | 
**Precision** | **int32** | field.tokenInfo.precision | 
**PrecisionForOrder** | **int32** | field.tokenInfo.precision | 
**MinOrderAmount** | **string** | field.tokenInfo.minOrderAmount | 
**MaxOrderAmount** | **string** | field.tokenInfo.maxOrderAmount | 
**DustOrderAmount** | **string** | field.tokenInfo.dustOrderAmount | 
**MaxLuckyTokenAmount** | **string** | field.tokenInfo.maxLuckyTokenAmount | 
**MinLuckyTokenAmount** | **string** | field.tokenInfo.minLuckyTokenAmount | 
**FastWithdrawLimit** | **string** | field.tokenInfo.fastWithdrawLimit | 
**DistributionGas** | **string** | field.tokenInfo.distributionGas | 
**DepositGas** | **string** | field.tokenInfo.depositGas | 
**Enabled** | **bool** | field.tokenInfo.enabled | 

## Methods

### NewTokenInfo

`func NewTokenInfo(type_ string, tokenId int32, symbol string, name string, address string, decimals int32, precision int32, precisionForOrder int32, minOrderAmount string, maxOrderAmount string, dustOrderAmount string, maxLuckyTokenAmount string, minLuckyTokenAmount string, fastWithdrawLimit string, distributionGas string, depositGas string, enabled bool, ) *TokenInfo`

NewTokenInfo instantiates a new TokenInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenInfoWithDefaults

`func NewTokenInfoWithDefaults() *TokenInfo`

NewTokenInfoWithDefaults instantiates a new TokenInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TokenInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenInfo) SetType(v string)`

SetType sets Type field to given value.


### GetTokenId

`func (o *TokenInfo) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenInfo) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenInfo) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetSymbol

`func (o *TokenInfo) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenInfo) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenInfo) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *TokenInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenInfo) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *TokenInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenInfo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDecimals

`func (o *TokenInfo) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenInfo) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenInfo) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetPrecision

`func (o *TokenInfo) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *TokenInfo) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *TokenInfo) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.


### GetPrecisionForOrder

`func (o *TokenInfo) GetPrecisionForOrder() int32`

GetPrecisionForOrder returns the PrecisionForOrder field if non-nil, zero value otherwise.

### GetPrecisionForOrderOk

`func (o *TokenInfo) GetPrecisionForOrderOk() (*int32, bool)`

GetPrecisionForOrderOk returns a tuple with the PrecisionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecisionForOrder

`func (o *TokenInfo) SetPrecisionForOrder(v int32)`

SetPrecisionForOrder sets PrecisionForOrder field to given value.


### GetMinOrderAmount

`func (o *TokenInfo) GetMinOrderAmount() string`

GetMinOrderAmount returns the MinOrderAmount field if non-nil, zero value otherwise.

### GetMinOrderAmountOk

`func (o *TokenInfo) GetMinOrderAmountOk() (*string, bool)`

GetMinOrderAmountOk returns a tuple with the MinOrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrderAmount

`func (o *TokenInfo) SetMinOrderAmount(v string)`

SetMinOrderAmount sets MinOrderAmount field to given value.


### GetMaxOrderAmount

`func (o *TokenInfo) GetMaxOrderAmount() string`

GetMaxOrderAmount returns the MaxOrderAmount field if non-nil, zero value otherwise.

### GetMaxOrderAmountOk

`func (o *TokenInfo) GetMaxOrderAmountOk() (*string, bool)`

GetMaxOrderAmountOk returns a tuple with the MaxOrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrderAmount

`func (o *TokenInfo) SetMaxOrderAmount(v string)`

SetMaxOrderAmount sets MaxOrderAmount field to given value.


### GetDustOrderAmount

`func (o *TokenInfo) GetDustOrderAmount() string`

GetDustOrderAmount returns the DustOrderAmount field if non-nil, zero value otherwise.

### GetDustOrderAmountOk

`func (o *TokenInfo) GetDustOrderAmountOk() (*string, bool)`

GetDustOrderAmountOk returns a tuple with the DustOrderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDustOrderAmount

`func (o *TokenInfo) SetDustOrderAmount(v string)`

SetDustOrderAmount sets DustOrderAmount field to given value.


### GetMaxLuckyTokenAmount

`func (o *TokenInfo) GetMaxLuckyTokenAmount() string`

GetMaxLuckyTokenAmount returns the MaxLuckyTokenAmount field if non-nil, zero value otherwise.

### GetMaxLuckyTokenAmountOk

`func (o *TokenInfo) GetMaxLuckyTokenAmountOk() (*string, bool)`

GetMaxLuckyTokenAmountOk returns a tuple with the MaxLuckyTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLuckyTokenAmount

`func (o *TokenInfo) SetMaxLuckyTokenAmount(v string)`

SetMaxLuckyTokenAmount sets MaxLuckyTokenAmount field to given value.


### GetMinLuckyTokenAmount

`func (o *TokenInfo) GetMinLuckyTokenAmount() string`

GetMinLuckyTokenAmount returns the MinLuckyTokenAmount field if non-nil, zero value otherwise.

### GetMinLuckyTokenAmountOk

`func (o *TokenInfo) GetMinLuckyTokenAmountOk() (*string, bool)`

GetMinLuckyTokenAmountOk returns a tuple with the MinLuckyTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLuckyTokenAmount

`func (o *TokenInfo) SetMinLuckyTokenAmount(v string)`

SetMinLuckyTokenAmount sets MinLuckyTokenAmount field to given value.


### GetFastWithdrawLimit

`func (o *TokenInfo) GetFastWithdrawLimit() string`

GetFastWithdrawLimit returns the FastWithdrawLimit field if non-nil, zero value otherwise.

### GetFastWithdrawLimitOk

`func (o *TokenInfo) GetFastWithdrawLimitOk() (*string, bool)`

GetFastWithdrawLimitOk returns a tuple with the FastWithdrawLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastWithdrawLimit

`func (o *TokenInfo) SetFastWithdrawLimit(v string)`

SetFastWithdrawLimit sets FastWithdrawLimit field to given value.


### GetDistributionGas

`func (o *TokenInfo) GetDistributionGas() string`

GetDistributionGas returns the DistributionGas field if non-nil, zero value otherwise.

### GetDistributionGasOk

`func (o *TokenInfo) GetDistributionGasOk() (*string, bool)`

GetDistributionGasOk returns a tuple with the DistributionGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionGas

`func (o *TokenInfo) SetDistributionGas(v string)`

SetDistributionGas sets DistributionGas field to given value.


### GetDepositGas

`func (o *TokenInfo) GetDepositGas() string`

GetDepositGas returns the DepositGas field if non-nil, zero value otherwise.

### GetDepositGasOk

`func (o *TokenInfo) GetDepositGasOk() (*string, bool)`

GetDepositGasOk returns a tuple with the DepositGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositGas

`func (o *TokenInfo) SetDepositGas(v string)`

SetDepositGas sets DepositGas field to given value.


### GetEnabled

`func (o *TokenInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TokenInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TokenInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


