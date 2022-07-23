# DepthV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int64** | An ever-increasing numeric version number that reflects its change history. | 
**Timestamp** | **int64** | Update timestamp | 
**Market** | **string** | Trading pair | 
**Bids** | **[][]string** | Array of bids, each item is an array that contains the price, size, volume and the number of orders aggregated at this price. | 
**Asks** | **[][]string** | Array of asks, each item is an array that contains the price, size, volume and the number of orders aggregated at this price. | 

## Methods

### NewDepthV3

`func NewDepthV3(version int64, timestamp int64, market string, bids [][]string, asks [][]string, ) *DepthV3`

NewDepthV3 instantiates a new DepthV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepthV3WithDefaults

`func NewDepthV3WithDefaults() *DepthV3`

NewDepthV3WithDefaults instantiates a new DepthV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DepthV3) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DepthV3) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DepthV3) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetTimestamp

`func (o *DepthV3) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DepthV3) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DepthV3) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetMarket

`func (o *DepthV3) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *DepthV3) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *DepthV3) SetMarket(v string)`

SetMarket sets Market field to given value.


### GetBids

`func (o *DepthV3) GetBids() [][]string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *DepthV3) GetBidsOk() (*[][]string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *DepthV3) SetBids(v [][]string)`

SetBids sets Bids field to given value.


### GetAsks

`func (o *DepthV3) GetAsks() [][]string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *DepthV3) GetAsksOk() (*[][]string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *DepthV3) SetAsks(v [][]string)`

SetAsks sets Asks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


