# SubmitNftOrderRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exchange** | **string** | The adderss of the exchange which has to process this order | 
**AccountId** | **int64** | Loopring&#39;s account ID | 
**StorageId** | **int32** | The unique identifier of the L2 Merkle tree storage slot where the burn made in order to exit the pool will or has been stored. | 
**SellToken** | [**NftTokenAmountInfo**](NftTokenAmountInfo.md) |  | 
**BuyToken** | [**NftTokenAmountInfo**](NftTokenAmountInfo.md) |  | 
**AllOrNone** | Pointer to **string** | Whether the order supports partial fills or not.Currently only supports false as a valid value | [optional] 
**FillAmountBOrS** | Pointer to **string** | Fill size by buy token or by sell token | [optional] 
**ValidUntil** | **int64** | Order expiration time, accuracy is in seconds | 
**MaxFeeBips** | **int32** | Maximum order fee that the user can accept, value range (in ten thousandths) 1 ~ 63 | 
**EddsaSignature** | **string** | The orders EdDSA signature. The signature is a hexadecimal string obtained by signing the order itself and concatenating the resulting signature parts (Rx, Ry, and S). Used to authenticate and authorize the operation. | 
**ClientOrderId** | Pointer to **string** | An arbitrary, client-set unique order identifier, max length is 120 bytes | [optional] 
**OrderType** | Pointer to **string** | Order types, can be AMM, LIMIT_ORDER, MAKER_ONLY, TAKER_ONLY | [optional] 
**TradeChannel** | Pointer to **string** | Order channel, can be ORDER_BOOK, AMM_POOL, MIXED | [optional] 
**Taker** | Pointer to **string** | Used by the P2P order which user specify the taker, so far its 0x0000000000000000000000000000000000000000 | [optional] 
**Affiliate** | Pointer to **string** | An accountID who will recieve a share of the fee of this order | [optional] 

## Methods

### NewSubmitNftOrderRequestV3

`func NewSubmitNftOrderRequestV3(exchange string, accountId int64, storageId int32, sellToken NftTokenAmountInfo, buyToken NftTokenAmountInfo, validUntil int64, maxFeeBips int32, eddsaSignature string, ) *SubmitNftOrderRequestV3`

NewSubmitNftOrderRequestV3 instantiates a new SubmitNftOrderRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitNftOrderRequestV3WithDefaults

`func NewSubmitNftOrderRequestV3WithDefaults() *SubmitNftOrderRequestV3`

NewSubmitNftOrderRequestV3WithDefaults instantiates a new SubmitNftOrderRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchange

`func (o *SubmitNftOrderRequestV3) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *SubmitNftOrderRequestV3) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *SubmitNftOrderRequestV3) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetAccountId

`func (o *SubmitNftOrderRequestV3) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SubmitNftOrderRequestV3) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SubmitNftOrderRequestV3) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetStorageId

`func (o *SubmitNftOrderRequestV3) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *SubmitNftOrderRequestV3) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *SubmitNftOrderRequestV3) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.


### GetSellToken

`func (o *SubmitNftOrderRequestV3) GetSellToken() NftTokenAmountInfo`

GetSellToken returns the SellToken field if non-nil, zero value otherwise.

### GetSellTokenOk

`func (o *SubmitNftOrderRequestV3) GetSellTokenOk() (*NftTokenAmountInfo, bool)`

GetSellTokenOk returns a tuple with the SellToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellToken

`func (o *SubmitNftOrderRequestV3) SetSellToken(v NftTokenAmountInfo)`

SetSellToken sets SellToken field to given value.


### GetBuyToken

`func (o *SubmitNftOrderRequestV3) GetBuyToken() NftTokenAmountInfo`

GetBuyToken returns the BuyToken field if non-nil, zero value otherwise.

### GetBuyTokenOk

`func (o *SubmitNftOrderRequestV3) GetBuyTokenOk() (*NftTokenAmountInfo, bool)`

GetBuyTokenOk returns a tuple with the BuyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyToken

`func (o *SubmitNftOrderRequestV3) SetBuyToken(v NftTokenAmountInfo)`

SetBuyToken sets BuyToken field to given value.


### GetAllOrNone

`func (o *SubmitNftOrderRequestV3) GetAllOrNone() string`

GetAllOrNone returns the AllOrNone field if non-nil, zero value otherwise.

### GetAllOrNoneOk

`func (o *SubmitNftOrderRequestV3) GetAllOrNoneOk() (*string, bool)`

GetAllOrNoneOk returns a tuple with the AllOrNone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllOrNone

`func (o *SubmitNftOrderRequestV3) SetAllOrNone(v string)`

SetAllOrNone sets AllOrNone field to given value.

### HasAllOrNone

`func (o *SubmitNftOrderRequestV3) HasAllOrNone() bool`

HasAllOrNone returns a boolean if a field has been set.

### GetFillAmountBOrS

`func (o *SubmitNftOrderRequestV3) GetFillAmountBOrS() string`

GetFillAmountBOrS returns the FillAmountBOrS field if non-nil, zero value otherwise.

### GetFillAmountBOrSOk

`func (o *SubmitNftOrderRequestV3) GetFillAmountBOrSOk() (*string, bool)`

GetFillAmountBOrSOk returns a tuple with the FillAmountBOrS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillAmountBOrS

`func (o *SubmitNftOrderRequestV3) SetFillAmountBOrS(v string)`

SetFillAmountBOrS sets FillAmountBOrS field to given value.

### HasFillAmountBOrS

`func (o *SubmitNftOrderRequestV3) HasFillAmountBOrS() bool`

HasFillAmountBOrS returns a boolean if a field has been set.

### GetValidUntil

`func (o *SubmitNftOrderRequestV3) GetValidUntil() int64`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *SubmitNftOrderRequestV3) GetValidUntilOk() (*int64, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *SubmitNftOrderRequestV3) SetValidUntil(v int64)`

SetValidUntil sets ValidUntil field to given value.


### GetMaxFeeBips

`func (o *SubmitNftOrderRequestV3) GetMaxFeeBips() int32`

GetMaxFeeBips returns the MaxFeeBips field if non-nil, zero value otherwise.

### GetMaxFeeBipsOk

`func (o *SubmitNftOrderRequestV3) GetMaxFeeBipsOk() (*int32, bool)`

GetMaxFeeBipsOk returns a tuple with the MaxFeeBips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeeBips

`func (o *SubmitNftOrderRequestV3) SetMaxFeeBips(v int32)`

SetMaxFeeBips sets MaxFeeBips field to given value.


### GetEddsaSignature

`func (o *SubmitNftOrderRequestV3) GetEddsaSignature() string`

GetEddsaSignature returns the EddsaSignature field if non-nil, zero value otherwise.

### GetEddsaSignatureOk

`func (o *SubmitNftOrderRequestV3) GetEddsaSignatureOk() (*string, bool)`

GetEddsaSignatureOk returns a tuple with the EddsaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddsaSignature

`func (o *SubmitNftOrderRequestV3) SetEddsaSignature(v string)`

SetEddsaSignature sets EddsaSignature field to given value.


### GetClientOrderId

`func (o *SubmitNftOrderRequestV3) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *SubmitNftOrderRequestV3) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *SubmitNftOrderRequestV3) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *SubmitNftOrderRequestV3) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetOrderType

`func (o *SubmitNftOrderRequestV3) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *SubmitNftOrderRequestV3) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *SubmitNftOrderRequestV3) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *SubmitNftOrderRequestV3) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetTradeChannel

`func (o *SubmitNftOrderRequestV3) GetTradeChannel() string`

GetTradeChannel returns the TradeChannel field if non-nil, zero value otherwise.

### GetTradeChannelOk

`func (o *SubmitNftOrderRequestV3) GetTradeChannelOk() (*string, bool)`

GetTradeChannelOk returns a tuple with the TradeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeChannel

`func (o *SubmitNftOrderRequestV3) SetTradeChannel(v string)`

SetTradeChannel sets TradeChannel field to given value.

### HasTradeChannel

`func (o *SubmitNftOrderRequestV3) HasTradeChannel() bool`

HasTradeChannel returns a boolean if a field has been set.

### GetTaker

`func (o *SubmitNftOrderRequestV3) GetTaker() string`

GetTaker returns the Taker field if non-nil, zero value otherwise.

### GetTakerOk

`func (o *SubmitNftOrderRequestV3) GetTakerOk() (*string, bool)`

GetTakerOk returns a tuple with the Taker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaker

`func (o *SubmitNftOrderRequestV3) SetTaker(v string)`

SetTaker sets Taker field to given value.

### HasTaker

`func (o *SubmitNftOrderRequestV3) HasTaker() bool`

HasTaker returns a boolean if a field has been set.

### GetAffiliate

`func (o *SubmitNftOrderRequestV3) GetAffiliate() string`

GetAffiliate returns the Affiliate field if non-nil, zero value otherwise.

### GetAffiliateOk

`func (o *SubmitNftOrderRequestV3) GetAffiliateOk() (*string, bool)`

GetAffiliateOk returns a tuple with the Affiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliate

`func (o *SubmitNftOrderRequestV3) SetAffiliate(v string)`

SetAffiliate sets Affiliate field to given value.

### HasAffiliate

`func (o *SubmitNftOrderRequestV3) HasAffiliate() bool`

HasAffiliate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


