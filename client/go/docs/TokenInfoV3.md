# TokenInfoV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Token Type | 
**TokenId** | **int32** | Token&#39;s numeric ID | 
**Symbol** | **string** | Token symbol | 
**Name** | **string** | Token name | 
**Address** | **string** | Token ERC20 contract address | 
**Decimals** | **int32** | Token decimals | 
**Precision** | **int32** | Max decimals that relayer uses for the token, smaller amount will be treated as zero. | 
**PrecisionForOrder** | **int32** | Max decimals that relayer uses for the token, smaller amount will be treated as zero. | 
**OrderAmounts** | [**OrderAmountsV3**](OrderAmountsV3.md) |  | 
**LuckyTokenAmounts** | [**OrderAmountsV3**](OrderAmountsV3.md) |  | 
**FastWithdrawLimit** | **string** | The maximum amount for single fast withdrawal | 
**GasAmounts** | [**GasAmountLimitV3**](GasAmountLimitV3.md) |  | 
**Enabled** | **bool** | Whether the token is currently enabled for deposits and withdrawals. | 

## Methods

### NewTokenInfoV3

`func NewTokenInfoV3(type_ string, tokenId int32, symbol string, name string, address string, decimals int32, precision int32, precisionForOrder int32, orderAmounts OrderAmountsV3, luckyTokenAmounts OrderAmountsV3, fastWithdrawLimit string, gasAmounts GasAmountLimitV3, enabled bool, ) *TokenInfoV3`

NewTokenInfoV3 instantiates a new TokenInfoV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenInfoV3WithDefaults

`func NewTokenInfoV3WithDefaults() *TokenInfoV3`

NewTokenInfoV3WithDefaults instantiates a new TokenInfoV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TokenInfoV3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenInfoV3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenInfoV3) SetType(v string)`

SetType sets Type field to given value.


### GetTokenId

`func (o *TokenInfoV3) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenInfoV3) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenInfoV3) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetSymbol

`func (o *TokenInfoV3) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenInfoV3) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenInfoV3) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *TokenInfoV3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenInfoV3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenInfoV3) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *TokenInfoV3) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenInfoV3) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenInfoV3) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDecimals

`func (o *TokenInfoV3) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenInfoV3) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenInfoV3) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetPrecision

`func (o *TokenInfoV3) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *TokenInfoV3) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *TokenInfoV3) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.


### GetPrecisionForOrder

`func (o *TokenInfoV3) GetPrecisionForOrder() int32`

GetPrecisionForOrder returns the PrecisionForOrder field if non-nil, zero value otherwise.

### GetPrecisionForOrderOk

`func (o *TokenInfoV3) GetPrecisionForOrderOk() (*int32, bool)`

GetPrecisionForOrderOk returns a tuple with the PrecisionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecisionForOrder

`func (o *TokenInfoV3) SetPrecisionForOrder(v int32)`

SetPrecisionForOrder sets PrecisionForOrder field to given value.


### GetOrderAmounts

`func (o *TokenInfoV3) GetOrderAmounts() OrderAmountsV3`

GetOrderAmounts returns the OrderAmounts field if non-nil, zero value otherwise.

### GetOrderAmountsOk

`func (o *TokenInfoV3) GetOrderAmountsOk() (*OrderAmountsV3, bool)`

GetOrderAmountsOk returns a tuple with the OrderAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmounts

`func (o *TokenInfoV3) SetOrderAmounts(v OrderAmountsV3)`

SetOrderAmounts sets OrderAmounts field to given value.


### GetLuckyTokenAmounts

`func (o *TokenInfoV3) GetLuckyTokenAmounts() OrderAmountsV3`

GetLuckyTokenAmounts returns the LuckyTokenAmounts field if non-nil, zero value otherwise.

### GetLuckyTokenAmountsOk

`func (o *TokenInfoV3) GetLuckyTokenAmountsOk() (*OrderAmountsV3, bool)`

GetLuckyTokenAmountsOk returns a tuple with the LuckyTokenAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLuckyTokenAmounts

`func (o *TokenInfoV3) SetLuckyTokenAmounts(v OrderAmountsV3)`

SetLuckyTokenAmounts sets LuckyTokenAmounts field to given value.


### GetFastWithdrawLimit

`func (o *TokenInfoV3) GetFastWithdrawLimit() string`

GetFastWithdrawLimit returns the FastWithdrawLimit field if non-nil, zero value otherwise.

### GetFastWithdrawLimitOk

`func (o *TokenInfoV3) GetFastWithdrawLimitOk() (*string, bool)`

GetFastWithdrawLimitOk returns a tuple with the FastWithdrawLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastWithdrawLimit

`func (o *TokenInfoV3) SetFastWithdrawLimit(v string)`

SetFastWithdrawLimit sets FastWithdrawLimit field to given value.


### GetGasAmounts

`func (o *TokenInfoV3) GetGasAmounts() GasAmountLimitV3`

GetGasAmounts returns the GasAmounts field if non-nil, zero value otherwise.

### GetGasAmountsOk

`func (o *TokenInfoV3) GetGasAmountsOk() (*GasAmountLimitV3, bool)`

GetGasAmountsOk returns a tuple with the GasAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmounts

`func (o *TokenInfoV3) SetGasAmounts(v GasAmountLimitV3)`

SetGasAmounts sets GasAmounts field to given value.


### GetEnabled

`func (o *TokenInfoV3) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TokenInfoV3) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TokenInfoV3) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


