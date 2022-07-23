# SubmitOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | field.SubmitOrderRequest.exchange | 
**StorageId** | **int32** | field.SubmitOrderRequest.storageId | 
**AccountId** | **int64** | field.SubmitOrderRequest.accountId | 
**TokenSId** | **int32** | field.SubmitOrderRequest.tokenSId | 
**TokenBId** | **int32** | field.SubmitOrderRequest.tokenBId | 
**AmountS** | **string** | field.SubmitOrderRequest.amountS | 
**AmountB** | **string** | field.SubmitOrderRequest.amountB | 
**AllOrNone** | Pointer to **string** | field.SubmitOrderRequest.allOrNone | [optional] 
**FillAmountBOrS** | **string** | field.SubmitOrderRequest.fillAmountBOrS | 
**ValidUntil** | **int64** | field.SubmitOrderRequest.validUntil | 
**MaxFeeBips** | **int32** | field.SubmitOrderRequest.maxFeeBips | 
**EddsaSig** | **string** | field.SubmitOrderRequest.eddsaSig | 
**ClientOrderId** | Pointer to **string** | field.SubmitOrderRequest.clientOrderId | [optional] 
**Affiliate** | Pointer to **string** | field.SubmitOrderRequest.affiliate | [optional] 
**OrderType** | Pointer to **string** | field.SubmitOrderRequest.orderType | [optional] 
**TradeChannel** | Pointer to **string** | field.SubmitOrderRequest.tradeChannel | [optional] 
**Taker** | Pointer to **string** | field.SubmitOrderRequest.taker | [optional] 
**PoolAddress** | Pointer to **string** | field.SubmitOrderRequest.taker | [optional] 

## Methods

### NewSubmitOrderRequest

`func NewSubmitOrderRequest(exchange string, storageId int32, accountId int64, tokenSId int32, tokenBId int32, amountS string, amountB string, fillAmountBOrS string, validUntil int64, maxFeeBips int32, eddsaSig string, ) *SubmitOrderRequest`

NewSubmitOrderRequest instantiates a new SubmitOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitOrderRequestWithDefaults

`func NewSubmitOrderRequestWithDefaults() *SubmitOrderRequest`

NewSubmitOrderRequestWithDefaults instantiates a new SubmitOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *SubmitOrderRequest) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *SubmitOrderRequest) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *SubmitOrderRequest) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetStorageId

`func (o *SubmitOrderRequest) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *SubmitOrderRequest) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *SubmitOrderRequest) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.


### GetAccountId

`func (o *SubmitOrderRequest) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SubmitOrderRequest) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SubmitOrderRequest) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetTokenSId

`func (o *SubmitOrderRequest) GetTokenSId() int32`

GetTokenSId returns the TokenSId field if non-nil, zero value otherwise.

### GetTokenSIdOk

`func (o *SubmitOrderRequest) GetTokenSIdOk() (*int32, bool)`

GetTokenSIdOk returns a tuple with the TokenSId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSId

`func (o *SubmitOrderRequest) SetTokenSId(v int32)`

SetTokenSId sets TokenSId field to given value.


### GetTokenBId

`func (o *SubmitOrderRequest) GetTokenBId() int32`

GetTokenBId returns the TokenBId field if non-nil, zero value otherwise.

### GetTokenBIdOk

`func (o *SubmitOrderRequest) GetTokenBIdOk() (*int32, bool)`

GetTokenBIdOk returns a tuple with the TokenBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBId

`func (o *SubmitOrderRequest) SetTokenBId(v int32)`

SetTokenBId sets TokenBId field to given value.


### GetAmountS

`func (o *SubmitOrderRequest) GetAmountS() string`

GetAmountS returns the AmountS field if non-nil, zero value otherwise.

### GetAmountSOk

`func (o *SubmitOrderRequest) GetAmountSOk() (*string, bool)`

GetAmountSOk returns a tuple with the AmountS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountS

`func (o *SubmitOrderRequest) SetAmountS(v string)`

SetAmountS sets AmountS field to given value.


### GetAmountB

`func (o *SubmitOrderRequest) GetAmountB() string`

GetAmountB returns the AmountB field if non-nil, zero value otherwise.

### GetAmountBOk

`func (o *SubmitOrderRequest) GetAmountBOk() (*string, bool)`

GetAmountBOk returns a tuple with the AmountB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountB

`func (o *SubmitOrderRequest) SetAmountB(v string)`

SetAmountB sets AmountB field to given value.


### GetAllOrNone

`func (o *SubmitOrderRequest) GetAllOrNone() string`

GetAllOrNone returns the AllOrNone field if non-nil, zero value otherwise.

### GetAllOrNoneOk

`func (o *SubmitOrderRequest) GetAllOrNoneOk() (*string, bool)`

GetAllOrNoneOk returns a tuple with the AllOrNone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllOrNone

`func (o *SubmitOrderRequest) SetAllOrNone(v string)`

SetAllOrNone sets AllOrNone field to given value.

### HasAllOrNone

`func (o *SubmitOrderRequest) HasAllOrNone() bool`

HasAllOrNone returns a boolean if a field has been set.

### GetFillAmountBOrS

`func (o *SubmitOrderRequest) GetFillAmountBOrS() string`

GetFillAmountBOrS returns the FillAmountBOrS field if non-nil, zero value otherwise.

### GetFillAmountBOrSOk

`func (o *SubmitOrderRequest) GetFillAmountBOrSOk() (*string, bool)`

GetFillAmountBOrSOk returns a tuple with the FillAmountBOrS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillAmountBOrS

`func (o *SubmitOrderRequest) SetFillAmountBOrS(v string)`

SetFillAmountBOrS sets FillAmountBOrS field to given value.


### GetValidUntil

`func (o *SubmitOrderRequest) GetValidUntil() int64`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *SubmitOrderRequest) GetValidUntilOk() (*int64, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *SubmitOrderRequest) SetValidUntil(v int64)`

SetValidUntil sets ValidUntil field to given value.


### GetMaxFeeBips

`func (o *SubmitOrderRequest) GetMaxFeeBips() int32`

GetMaxFeeBips returns the MaxFeeBips field if non-nil, zero value otherwise.

### GetMaxFeeBipsOk

`func (o *SubmitOrderRequest) GetMaxFeeBipsOk() (*int32, bool)`

GetMaxFeeBipsOk returns a tuple with the MaxFeeBips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeBips

`func (o *SubmitOrderRequest) SetMaxFeeBips(v int32)`

SetMaxFeeBips sets MaxFeeBips field to given value.


### GetEddsaSig

`func (o *SubmitOrderRequest) GetEddsaSig() string`

GetEddsaSig returns the EddsaSig field if non-nil, zero value otherwise.

### GetEddsaSigOk

`func (o *SubmitOrderRequest) GetEddsaSigOk() (*string, bool)`

GetEddsaSigOk returns a tuple with the EddsaSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSig

`func (o *SubmitOrderRequest) SetEddsaSig(v string)`

SetEddsaSig sets EddsaSig field to given value.


### GetClientOrderId

`func (o *SubmitOrderRequest) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *SubmitOrderRequest) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *SubmitOrderRequest) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *SubmitOrderRequest) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetAffiliate

`func (o *SubmitOrderRequest) GetAffiliate() string`

GetAffiliate returns the Affiliate field if non-nil, zero value otherwise.

### GetAffiliateOk

`func (o *SubmitOrderRequest) GetAffiliateOk() (*string, bool)`

GetAffiliateOk returns a tuple with the Affiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliate

`func (o *SubmitOrderRequest) SetAffiliate(v string)`

SetAffiliate sets Affiliate field to given value.

### HasAffiliate

`func (o *SubmitOrderRequest) HasAffiliate() bool`

HasAffiliate returns a boolean if a field has been set.

### GetOrderType

`func (o *SubmitOrderRequest) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *SubmitOrderRequest) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *SubmitOrderRequest) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *SubmitOrderRequest) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetTradeChannel

`func (o *SubmitOrderRequest) GetTradeChannel() string`

GetTradeChannel returns the TradeChannel field if non-nil, zero value otherwise.

### GetTradeChannelOk

`func (o *SubmitOrderRequest) GetTradeChannelOk() (*string, bool)`

GetTradeChannelOk returns a tuple with the TradeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeChannel

`func (o *SubmitOrderRequest) SetTradeChannel(v string)`

SetTradeChannel sets TradeChannel field to given value.

### HasTradeChannel

`func (o *SubmitOrderRequest) HasTradeChannel() bool`

HasTradeChannel returns a boolean if a field has been set.

### GetTaker

`func (o *SubmitOrderRequest) GetTaker() string`

GetTaker returns the Taker field if non-nil, zero value otherwise.

### GetTakerOk

`func (o *SubmitOrderRequest) GetTakerOk() (*string, bool)`

GetTakerOk returns a tuple with the Taker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaker

`func (o *SubmitOrderRequest) SetTaker(v string)`

SetTaker sets Taker field to given value.

### HasTaker

`func (o *SubmitOrderRequest) HasTaker() bool`

HasTaker returns a boolean if a field has been set.

### GetPoolAddress

`func (o *SubmitOrderRequest) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *SubmitOrderRequest) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *SubmitOrderRequest) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.

### HasPoolAddress

`func (o *SubmitOrderRequest) HasPoolAddress() bool`

HasPoolAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


