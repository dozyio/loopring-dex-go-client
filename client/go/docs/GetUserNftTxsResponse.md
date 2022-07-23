# GetUserNftTxsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | field.totalNum.description | 
**Transactions** | [**[]NftTxData**](NftTxData.md) | field.GetUserNftTxsResponse.transactions | 

## Methods

### NewGetUserNftTxsResponse

`func NewGetUserNftTxsResponse(totalNum int64, transactions []NftTxData, ) *GetUserNftTxsResponse`

NewGetUserNftTxsResponse instantiates a new GetUserNftTxsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserNftTxsResponseWithDefaults

`func NewGetUserNftTxsResponseWithDefaults() *GetUserNftTxsResponse`

NewGetUserNftTxsResponseWithDefaults instantiates a new GetUserNftTxsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *GetUserNftTxsResponse) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *GetUserNftTxsResponse) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *GetUserNftTxsResponse) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetTransactions

`func (o *GetUserNftTxsResponse) GetTransactions() []NftTxData`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *GetUserNftTxsResponse) GetTransactionsOk() (*[]NftTxData, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *GetUserNftTxsResponse) SetTransactions(v []NftTxData)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


