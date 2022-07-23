# Depth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int64** | An ever-increasing numeric version number that reflects its change history. | 
**Timestamp** | **int64** | Update timestamp | 
**Market** | **string** | field.trade.market | 
**Bids** | **[][]string** | Array of bids, each item is an array that contains the price, size, volume and the number of orders aggregated at this price. | 
**Asks** | **[][]string** | Array of asks, each item is an array that contains the price, size, volume and the number of orders aggregated at this price. | 

## Methods

### NewDepth

`func NewDepth(version int64, timestamp int64, market string, bids [][]string, asks [][]string, ) *Depth`

NewDepth instantiates a new Depth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepthWithDefaults

`func NewDepthWithDefaults() *Depth`

NewDepthWithDefaults instantiates a new Depth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *Depth) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Depth) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Depth) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetTimestamp

`func (o *Depth) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Depth) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Depth) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetMarket

`func (o *Depth) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *Depth) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *Depth) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetBids

`func (o *Depth) GetBids() [][]string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *Depth) GetBidsOk() (*[][]string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *Depth) SetBids(v [][]string)`

SetBids sets Bids field to given value.


### GetAsks

`func (o *Depth) GetAsks() [][]string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *Depth) GetAsksOk() (*[][]string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *Depth) SetAsks(v [][]string)`

SetAsks sets Asks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


