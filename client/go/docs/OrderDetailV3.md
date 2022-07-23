# OrderDetailV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Order hash | 
**ClientOrderId** | **string** | Order&#39;s client-side ID | 
**Side** | **string** | Order&#39;s side | 
**Market** | **string** | Trading pair | 
**Price** | **string** | Order price | 
**Volumes** | [**OrderVolumesV3**](OrderVolumesV3.md) |  | 
**Validity** | [**OrderValidityV3**](OrderValidityV3.md) |  | 
**OrderType** | **string** | Whether the order has to be treated as a limit, maker, or taker operation. | 
**TradeChannel** | **string** | Order channel, can be ORDER_BOOK, AMM_POOL, MIXED | 
**Status** | **string** | Order status | 
**StorageInfo** | Pointer to [**StorageInfo**](StorageInfo.md) |  | [optional] 

## Methods

### NewOrderDetailV3

`func NewOrderDetailV3(hash string, clientOrderId string, side string, market string, price string, volumes OrderVolumesV3, validity OrderValidityV3, orderType string, tradeChannel string, status string, ) *OrderDetailV3`

NewOrderDetailV3 instantiates a new OrderDetailV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailV3WithDefaults

`func NewOrderDetailV3WithDefaults() *OrderDetailV3`

NewOrderDetailV3WithDefaults instantiates a new OrderDetailV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *OrderDetailV3) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *OrderDetailV3) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *OrderDetailV3) SetHash(v string)`

SetHash sets Hash field to given value.


### GetClientOrderId

`func (o *OrderDetailV3) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OrderDetailV3) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OrderDetailV3) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.


### GetSide

`func (o *OrderDetailV3) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderDetailV3) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderDetailV3) SetSide(v string)`

SetSide sets Side field to given value.


### GetMarket

`func (o *OrderDetailV3) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *OrderDetailV3) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *OrderDetailV3) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetPrice

`func (o *OrderDetailV3) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderDetailV3) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderDetailV3) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetVolumes

`func (o *OrderDetailV3) GetVolumes() OrderVolumesV3`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *OrderDetailV3) GetVolumesOk() (*OrderVolumesV3, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *OrderDetailV3) SetVolumes(v OrderVolumesV3)`

SetVolumes sets Volumes field to given value.


### GetValidity

`func (o *OrderDetailV3) GetValidity() OrderValidityV3`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *OrderDetailV3) GetValidityOk() (*OrderValidityV3, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *OrderDetailV3) SetValidity(v OrderValidityV3)`

SetValidity sets Validity field to given value.


### GetOrderType

`func (o *OrderDetailV3) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderDetailV3) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderDetailV3) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.


### GetTradeChannel

`func (o *OrderDetailV3) GetTradeChannel() string`

GetTradeChannel returns the TradeChannel field if non-nil, zero value otherwise.

### GetTradeChannelOk

`func (o *OrderDetailV3) GetTradeChannelOk() (*string, bool)`

GetTradeChannelOk returns a tuple with the TradeChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeChannel

`func (o *OrderDetailV3) SetTradeChannel(v string)`

SetTradeChannel sets TradeChannel field to given value.


### GetStatus

`func (o *OrderDetailV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderDetailV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderDetailV3) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStorageInfo

`func (o *OrderDetailV3) GetStorageInfo() StorageInfo`

GetStorageInfo returns the StorageInfo field if non-nil, zero value otherwise.

### GetStorageInfoOk

`func (o *OrderDetailV3) GetStorageInfoOk() (*StorageInfo, bool)`

GetStorageInfoOk returns a tuple with the StorageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInfo

`func (o *OrderDetailV3) SetStorageInfo(v StorageInfo)`

SetStorageInfo sets StorageInfo field to given value.

### HasStorageInfo

`func (o *OrderDetailV3) HasStorageInfo() bool`

HasStorageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


