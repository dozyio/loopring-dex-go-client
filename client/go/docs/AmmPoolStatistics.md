# AmmPoolStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | **string** | field.AmmPoolStatistics.market | 
**Liquidity** | **[]string** | field.AmmPoolStatistics.liquidity | 
**LpLiquidity** | **string** | field.AmmPoolStatistics.lpLiquidity | 
**LiquidityUSD** | **string** | field.AmmPoolStatistics.liquidityUSD | 
**Ohlc** | **[]string** | field.AmmPoolStatistics.ohlc | 
**Volume** | **[]string** | field.AmmPoolStatistics.volume | 
**Fees** | **[]string** | field.AmmPoolStatistics.fees | 
**ApyBips** | **string** | field.AmmPoolStatistics.apyBips | 
**IsRecommended** | **bool** | field.AmmPoolStatistics.isRecommended | 
**Rewards** | [**[]TokenVolumeV3**](TokenVolumeV3.md) | field.AmmPoolStatistics.rewards | 

## Methods

### NewAmmPoolStatistics

`func NewAmmPoolStatistics(market string, liquidity []string, lpLiquidity string, liquidityUSD string, ohlc []string, volume []string, fees []string, apyBips string, isRecommended bool, rewards []TokenVolumeV3, ) *AmmPoolStatistics`

NewAmmPoolStatistics instantiates a new AmmPoolStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmPoolStatisticsWithDefaults

`func NewAmmPoolStatisticsWithDefaults() *AmmPoolStatistics`

NewAmmPoolStatisticsWithDefaults instantiates a new AmmPoolStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *AmmPoolStatistics) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *AmmPoolStatistics) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *AmmPoolStatistics) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetLiquidity

`func (o *AmmPoolStatistics) GetLiquidity() []string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *AmmPoolStatistics) GetLiquidityOk() (*[]string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *AmmPoolStatistics) SetLiquidity(v []string)`

SetLiquidity sets Liquidity field to given value.


### GetLpLiquidity

`func (o *AmmPoolStatistics) GetLpLiquidity() string`

GetLpLiquidity returns the LpLiquidity field if non-nil, zero value otherwise.

### GetLpLiquidityOk

`func (o *AmmPoolStatistics) GetLpLiquidityOk() (*string, bool)`

GetLpLiquidityOk returns a tuple with the LpLiquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpLiquidity

`func (o *AmmPoolStatistics) SetLpLiquidity(v string)`

SetLpLiquidity sets LpLiquidity field to given value.


### GetLiquidityUSD

`func (o *AmmPoolStatistics) GetLiquidityUSD() string`

GetLiquidityUSD returns the LiquidityUSD field if non-nil, zero value otherwise.

### GetLiquidityUSDOk

`func (o *AmmPoolStatistics) GetLiquidityUSDOk() (*string, bool)`

GetLiquidityUSDOk returns a tuple with the LiquidityUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidityUSD

`func (o *AmmPoolStatistics) SetLiquidityUSD(v string)`

SetLiquidityUSD sets LiquidityUSD field to given value.


### GetOhlc

`func (o *AmmPoolStatistics) GetOhlc() []string`

GetOhlc returns the Ohlc field if non-nil, zero value otherwise.

### GetOhlcOk

`func (o *AmmPoolStatistics) GetOhlcOk() (*[]string, bool)`

GetOhlcOk returns a tuple with the Ohlc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOhlc

`func (o *AmmPoolStatistics) SetOhlc(v []string)`

SetOhlc sets Ohlc field to given value.


### GetVolume

`func (o *AmmPoolStatistics) GetVolume() []string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *AmmPoolStatistics) GetVolumeOk() (*[]string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *AmmPoolStatistics) SetVolume(v []string)`

SetVolume sets Volume field to given value.


### GetFees

`func (o *AmmPoolStatistics) GetFees() []string`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *AmmPoolStatistics) GetFeesOk() (*[]string, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *AmmPoolStatistics) SetFees(v []string)`

SetFees sets Fees field to given value.


### GetApyBips

`func (o *AmmPoolStatistics) GetApyBips() string`

GetApyBips returns the ApyBips field if non-nil, zero value otherwise.

### GetApyBipsOk

`func (o *AmmPoolStatistics) GetApyBipsOk() (*string, bool)`

GetApyBipsOk returns a tuple with the ApyBips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApyBips

`func (o *AmmPoolStatistics) SetApyBips(v string)`

SetApyBips sets ApyBips field to given value.


### GetIsRecommended

`func (o *AmmPoolStatistics) GetIsRecommended() bool`

GetIsRecommended returns the IsRecommended field if non-nil, zero value otherwise.

### GetIsRecommendedOk

`func (o *AmmPoolStatistics) GetIsRecommendedOk() (*bool, bool)`

GetIsRecommendedOk returns a tuple with the IsRecommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecommended

`func (o *AmmPoolStatistics) SetIsRecommended(v bool)`

SetIsRecommended sets IsRecommended field to given value.


### GetRewards

`func (o *AmmPoolStatistics) GetRewards() []TokenVolumeV3`

GetRewards returns the Rewards field if non-nil, zero value otherwise.

### GetRewardsOk

`func (o *AmmPoolStatistics) GetRewardsOk() (*[]TokenVolumeV3, bool)`

GetRewardsOk returns a tuple with the Rewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewards

`func (o *AmmPoolStatistics) SetRewards(v []TokenVolumeV3)`

SetRewards sets Rewards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


