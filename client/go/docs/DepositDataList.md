# DepositDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | Total number of records found | 
**Transactions** | [**[]DepositData**](DepositData.md) | List of DepositRecord | 

## Methods

### NewDepositDataList

`func NewDepositDataList(totalNum int64, transactions []DepositData, ) *DepositDataList`

NewDepositDataList instantiates a new DepositDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositDataListWithDefaults

`func NewDepositDataListWithDefaults() *DepositDataList`

NewDepositDataListWithDefaults instantiates a new DepositDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *DepositDataList) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *DepositDataList) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *DepositDataList) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetTransactions

`func (o *DepositDataList) GetTransactions() []DepositData`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *DepositDataList) GetTransactionsOk() (*[]DepositData, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *DepositDataList) SetTransactions(v []DepositData)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


