# UserAccountUpdateDataList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | Total number of records found | 
**Transactions** | [**[]UserAccountTxData**](UserAccountTxData.md) | List of UserPasswordChangeRecord | 

## Methods

### NewUserAccountUpdateDataList

`func NewUserAccountUpdateDataList(totalNum int64, transactions []UserAccountTxData, ) *UserAccountUpdateDataList`

NewUserAccountUpdateDataList instantiates a new UserAccountUpdateDataList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountUpdateDataListWithDefaults

`func NewUserAccountUpdateDataListWithDefaults() *UserAccountUpdateDataList`

NewUserAccountUpdateDataListWithDefaults instantiates a new UserAccountUpdateDataList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *UserAccountUpdateDataList) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *UserAccountUpdateDataList) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *UserAccountUpdateDataList) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetTransactions

`func (o *UserAccountUpdateDataList) GetTransactions() []UserAccountTxData`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *UserAccountUpdateDataList) GetTransactionsOk() (*[]UserAccountTxData, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *UserAccountUpdateDataList) SetTransactions(v []UserAccountTxData)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


