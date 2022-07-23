# NextStorageIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** | Next storage ID for order request, must be even | [optional] 
**OffchainId** | Pointer to **int64** | Next storage ID for offchain requests, i.e., transfer/withdraw/updateAccount, must be odd | [optional] 

## Methods

### NewNextStorageIdResponse

`func NewNextStorageIdResponse() *NextStorageIdResponse`

NewNextStorageIdResponse instantiates a new NextStorageIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextStorageIdResponseWithDefaults

`func NewNextStorageIdResponseWithDefaults() *NextStorageIdResponse`

NewNextStorageIdResponseWithDefaults instantiates a new NextStorageIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *NextStorageIdResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *NextStorageIdResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *NextStorageIdResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *NextStorageIdResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOffchainId

`func (o *NextStorageIdResponse) GetOffchainId() int64`

GetOffchainId returns the OffchainId field if non-nil, zero value otherwise.

### GetOffchainIdOk

`func (o *NextStorageIdResponse) GetOffchainIdOk() (*int64, bool)`

GetOffchainIdOk returns a tuple with the OffchainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffchainId

`func (o *NextStorageIdResponse) SetOffchainId(v int64)`

SetOffchainId sets OffchainId field to given value.

### HasOffchainId

`func (o *NextStorageIdResponse) HasOffchainId() bool`

HasOffchainId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


