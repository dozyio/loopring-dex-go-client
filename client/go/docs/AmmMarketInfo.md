# AmmMarketInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | field.AmmMarketInfo.name | 
**Market** | **string** | field.AmmMarketInfo.market | 
**Address** | **string** | field.AmmMarketInfo.address | 
**Version** | **string** | field.AmmMarketInfo.version | 
**InPoolTokens** | **[]map[string]interface{}** | field.AmmMarketInfo.inPoolTokens | 
**PoolTokenId** | **int32** | field.AmmMarketInfo.poolTokenId | 
**FeeBips** | **int32** | field.AmmMarketInfo.feeBips | 
**PricePrecision** | **int32** | field.AmmMarketInfo.pricePrecision | 
**AmountPrecision** | **int32** | field.AmmMarketInfo.amountPrecision | 
**CreatedAt** | **string** | field.AmmMarketInfo.createdAt | 
**Enabled** | **bool** | field.AmmMarketInfo.enabled | 
**Status** | **int32** | field.AmmMarketInfo.status | 

## Methods

### NewAmmMarketInfo

`func NewAmmMarketInfo(name string, market string, address string, version string, inPoolTokens []map[string]interface{}, poolTokenId int32, feeBips int32, pricePrecision int32, amountPrecision int32, createdAt string, enabled bool, status int32, ) *AmmMarketInfo`

NewAmmMarketInfo instantiates a new AmmMarketInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmMarketInfoWithDefaults

`func NewAmmMarketInfoWithDefaults() *AmmMarketInfo`

NewAmmMarketInfoWithDefaults instantiates a new AmmMarketInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AmmMarketInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmmMarketInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmmMarketInfo) SetName(v string)`

SetName sets Name field to given value.


### GetMarket

`func (o *AmmMarketInfo) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *AmmMarketInfo) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *AmmMarketInfo) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetAddress

`func (o *AmmMarketInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AmmMarketInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AmmMarketInfo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetVersion

`func (o *AmmMarketInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AmmMarketInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AmmMarketInfo) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetInPoolTokens

`func (o *AmmMarketInfo) GetInPoolTokens() []map[string]interface{}`

GetInPoolTokens returns the InPoolTokens field if non-nil, zero value otherwise.

### GetInPoolTokensOk

`func (o *AmmMarketInfo) GetInPoolTokensOk() (*[]map[string]interface{}, bool)`

GetInPoolTokensOk returns a tuple with the InPoolTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInPoolTokens

`func (o *AmmMarketInfo) SetInPoolTokens(v []map[string]interface{})`

SetInPoolTokens sets InPoolTokens field to given value.


### GetPoolTokenId

`func (o *AmmMarketInfo) GetPoolTokenId() int32`

GetPoolTokenId returns the PoolTokenId field if non-nil, zero value otherwise.

### GetPoolTokenIdOk

`func (o *AmmMarketInfo) GetPoolTokenIdOk() (*int32, bool)`

GetPoolTokenIdOk returns a tuple with the PoolTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolTokenId

`func (o *AmmMarketInfo) SetPoolTokenId(v int32)`

SetPoolTokenId sets PoolTokenId field to given value.


### GetFeeBips

`func (o *AmmMarketInfo) GetFeeBips() int32`

GetFeeBips returns the FeeBips field if non-nil, zero value otherwise.

### GetFeeBipsOk

`func (o *AmmMarketInfo) GetFeeBipsOk() (*int32, bool)`

GetFeeBipsOk returns a tuple with the FeeBips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBips

`func (o *AmmMarketInfo) SetFeeBips(v int32)`

SetFeeBips sets FeeBips field to given value.


### GetPricePrecision

`func (o *AmmMarketInfo) GetPricePrecision() int32`

GetPricePrecision returns the PricePrecision field if non-nil, zero value otherwise.

### GetPricePrecisionOk

`func (o *AmmMarketInfo) GetPricePrecisionOk() (*int32, bool)`

GetPricePrecisionOk returns a tuple with the PricePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePrecision

`func (o *AmmMarketInfo) SetPricePrecision(v int32)`

SetPricePrecision sets PricePrecision field to given value.


### GetAmountPrecision

`func (o *AmmMarketInfo) GetAmountPrecision() int32`

GetAmountPrecision returns the AmountPrecision field if non-nil, zero value otherwise.

### GetAmountPrecisionOk

`func (o *AmmMarketInfo) GetAmountPrecisionOk() (*int32, bool)`

GetAmountPrecisionOk returns a tuple with the AmountPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPrecision

`func (o *AmmMarketInfo) SetAmountPrecision(v int32)`

SetAmountPrecision sets AmountPrecision field to given value.


### GetCreatedAt

`func (o *AmmMarketInfo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AmmMarketInfo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AmmMarketInfo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnabled

`func (o *AmmMarketInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AmmMarketInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AmmMarketInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetStatus

`func (o *AmmMarketInfo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AmmMarketInfo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AmmMarketInfo) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


