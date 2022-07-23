# GetUserFeeRates2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | Account ID | 
**Market** | **string** | Market symbol | 
**TokenB** | **int32** | Token ID | 
**AmountB** | **string** | Amount to buy | 

## Methods

### NewGetUserFeeRates2Request

`func NewGetUserFeeRates2Request(accountId int64, market string, tokenB int32, amountB string, ) *GetUserFeeRates2Request`

NewGetUserFeeRates2Request instantiates a new GetUserFeeRates2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserFeeRates2RequestWithDefaults

`func NewGetUserFeeRates2RequestWithDefaults() *GetUserFeeRates2Request`

NewGetUserFeeRates2RequestWithDefaults instantiates a new GetUserFeeRates2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GetUserFeeRates2Request) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetUserFeeRates2Request) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetUserFeeRates2Request) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetMarket

`func (o *GetUserFeeRates2Request) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *GetUserFeeRates2Request) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *GetUserFeeRates2Request) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetTokenB

`func (o *GetUserFeeRates2Request) GetTokenB() int32`

GetTokenB returns the TokenB field if non-nil, zero value otherwise.

### GetTokenBOk

`func (o *GetUserFeeRates2Request) GetTokenBOk() (*int32, bool)`

GetTokenBOk returns a tuple with the TokenB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenB

`func (o *GetUserFeeRates2Request) SetTokenB(v int32)`

SetTokenB sets TokenB field to given value.


### GetAmountB

`func (o *GetUserFeeRates2Request) GetAmountB() string`

GetAmountB returns the AmountB field if non-nil, zero value otherwise.

### GetAmountBOk

`func (o *GetUserFeeRates2Request) GetAmountBOk() (*string, bool)`

GetAmountBOk returns a tuple with the AmountB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountB

`func (o *GetUserFeeRates2Request) SetAmountB(v string)`

SetAmountB sets AmountB field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


