# AmmTransactionDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** |  | 
**Transactions** | [**[]AmmTransactionData**](AmmTransactionData.md) |  | 

## Methods

### NewAmmTransactionDataList

`func NewAmmTransactionDataList(totalNum int64, transactions []AmmTransactionData, ) *AmmTransactionDataList`

NewAmmTransactionDataList instantiates a new AmmTransactionDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmTransactionDataListWithDefaults

`func NewAmmTransactionDataListWithDefaults() *AmmTransactionDataList`

NewAmmTransactionDataListWithDefaults instantiates a new AmmTransactionDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *AmmTransactionDataList) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *AmmTransactionDataList) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *AmmTransactionDataList) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetTransactions

`func (o *AmmTransactionDataList) GetTransactions() []AmmTransactionData`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *AmmTransactionDataList) GetTransactionsOk() (*[]AmmTransactionData, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *AmmTransactionDataList) SetTransactions(v []AmmTransactionData)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


