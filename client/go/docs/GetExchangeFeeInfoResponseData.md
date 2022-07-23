# GetExchangeFeeInfoResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ORDERBOOK_TRADING_FEES_STABLECOIN** | [**map[string]FeeRate**](FeeRate.md) | Stable coin orderbook fee rate | 
**ORDERBOOK_TRADING_FEES** | [**map[string]FeeRate**](FeeRate.md) | Common orderbook fee rate | 
**AMM_TRADING_FEES** | [**map[string]FeeRate**](FeeRate.md) | AMM fee rate | 
**OTHER_FEES** | **map[string]string** | Other fee | 

## Methods

### NewGetExchangeFeeInfoResponseData

`func NewGetExchangeFeeInfoResponseData(oRDERBOOKTRADINGFEESSTABLECOIN map[string]FeeRate, oRDERBOOKTRADINGFEES map[string]FeeRate, aMMTRADINGFEES map[string]FeeRate, oTHERFEES map[string]string, ) *GetExchangeFeeInfoResponseData`

NewGetExchangeFeeInfoResponseData instantiates a new GetExchangeFeeInfoResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeFeeInfoResponseDataWithDefaults

`func NewGetExchangeFeeInfoResponseDataWithDefaults() *GetExchangeFeeInfoResponseData`

NewGetExchangeFeeInfoResponseDataWithDefaults instantiates a new GetExchangeFeeInfoResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetORDERBOOK_TRADING_FEES_STABLECOIN

`func (o *GetExchangeFeeInfoResponseData) GetORDERBOOK_TRADING_FEES_STABLECOIN() map[string]FeeRate`

GetORDERBOOK_TRADING_FEES_STABLECOIN returns the ORDERBOOK_TRADING_FEES_STABLECOIN field if non-nil, zero value otherwise.

### GetORDERBOOK_TRADING_FEES_STABLECOINOk

`func (o *GetExchangeFeeInfoResponseData) GetORDERBOOK_TRADING_FEES_STABLECOINOk() (*map[string]FeeRate, bool)`

GetORDERBOOK_TRADING_FEES_STABLECOINOk returns a tuple with the ORDERBOOK_TRADING_FEES_STABLECOIN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORDERBOOK_TRADING_FEES_STABLECOIN

`func (o *GetExchangeFeeInfoResponseData) SetORDERBOOK_TRADING_FEES_STABLECOIN(v map[string]FeeRate)`

SetORDERBOOK_TRADING_FEES_STABLECOIN sets ORDERBOOK_TRADING_FEES_STABLECOIN field to given value.


### GetORDERBOOK_TRADING_FEES

`func (o *GetExchangeFeeInfoResponseData) GetORDERBOOK_TRADING_FEES() map[string]FeeRate`

GetORDERBOOK_TRADING_FEES returns the ORDERBOOK_TRADING_FEES field if non-nil, zero value otherwise.

### GetORDERBOOK_TRADING_FEESOk

`func (o *GetExchangeFeeInfoResponseData) GetORDERBOOK_TRADING_FEESOk() (*map[string]FeeRate, bool)`

GetORDERBOOK_TRADING_FEESOk returns a tuple with the ORDERBOOK_TRADING_FEES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORDERBOOK_TRADING_FEES

`func (o *GetExchangeFeeInfoResponseData) SetORDERBOOK_TRADING_FEES(v map[string]FeeRate)`

SetORDERBOOK_TRADING_FEES sets ORDERBOOK_TRADING_FEES field to given value.


### GetAMM_TRADING_FEES

`func (o *GetExchangeFeeInfoResponseData) GetAMM_TRADING_FEES() map[string]FeeRate`

GetAMM_TRADING_FEES returns the AMM_TRADING_FEES field if non-nil, zero value otherwise.

### GetAMM_TRADING_FEESOk

`func (o *GetExchangeFeeInfoResponseData) GetAMM_TRADING_FEESOk() (*map[string]FeeRate, bool)`

GetAMM_TRADING_FEESOk returns a tuple with the AMM_TRADING_FEES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAMM_TRADING_FEES

`func (o *GetExchangeFeeInfoResponseData) SetAMM_TRADING_FEES(v map[string]FeeRate)`

SetAMM_TRADING_FEES sets AMM_TRADING_FEES field to given value.


### GetOTHER_FEES

`func (o *GetExchangeFeeInfoResponseData) GetOTHER_FEES() map[string]string`

GetOTHER_FEES returns the OTHER_FEES field if non-nil, zero value otherwise.

### GetOTHER_FEESOk

`func (o *GetExchangeFeeInfoResponseData) GetOTHER_FEESOk() (*map[string]string, bool)`

GetOTHER_FEESOk returns a tuple with the OTHER_FEES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOTHER_FEES

`func (o *GetExchangeFeeInfoResponseData) SetOTHER_FEES(v map[string]string)`

SetOTHER_FEES sets OTHER_FEES field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


