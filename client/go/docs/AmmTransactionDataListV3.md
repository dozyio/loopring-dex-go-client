# AmmTransactionDataListV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | Number of AMM pool transactions | 
**Transactions** | [**[]AmmTransactionDataV3**](AmmTransactionDataV3.md) | List of AMM pool transactions | 

## Methods

### NewAmmTransactionDataListV3

`func NewAmmTransactionDataListV3(totalNum int64, transactions []AmmTransactionDataV3, ) *AmmTransactionDataListV3`

NewAmmTransactionDataListV3 instantiates a new AmmTransactionDataListV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmTransactionDataListV3WithDefaults

`func NewAmmTransactionDataListV3WithDefaults() *AmmTransactionDataListV3`

NewAmmTransactionDataListV3WithDefaults instantiates a new AmmTransactionDataListV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *AmmTransactionDataListV3) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *AmmTransactionDataListV3) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *AmmTransactionDataListV3) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetTransactions

`func (o *AmmTransactionDataListV3) GetTransactions() []AmmTransactionDataV3`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *AmmTransactionDataListV3) GetTransactionsOk() (*[]AmmTransactionDataV3, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *AmmTransactionDataListV3) SetTransactions(v []AmmTransactionDataV3)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


