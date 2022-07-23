# UserBill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**From** | **string** |  | 
**To** | **string** |  | 
**Token** | **string** |  | 
**Amount** | **string** |  | 
**TokenF** | **string** |  | 
**AmountF** | **string** |  | 
**Status** | **string** |  | 
**TxHash** | **string** |  | 
**BillType** | **string** |  | 
**Income** | **bool** |  | 
**Timestamp** | **int64** |  | 
**Memo** | **string** |  | 
**Price** | **string** |  | 
**TransferType** | **string** |  | 
**Hash** | **string** |  | 
**AccountId** | **int64** |  | 
**TokenId** | **int32** |  | 
**StorageId** | **int64** |  | 
**Label** | **string** |  | 

## Methods

### NewUserBill

`func NewUserBill(id int64, from string, to string, token string, amount string, tokenF string, amountF string, status string, txHash string, billType string, income bool, timestamp int64, memo string, price string, transferType string, hash string, accountId int64, tokenId int32, storageId int64, label string, ) *UserBill`

NewUserBill instantiates a new UserBill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBillWithDefaults

`func NewUserBillWithDefaults() *UserBill`

NewUserBillWithDefaults instantiates a new UserBill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserBill) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserBill) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserBill) SetId(v int64)`

SetId sets Id field to given value.


### GetFrom

`func (o *UserBill) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *UserBill) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *UserBill) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *UserBill) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *UserBill) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *UserBill) SetTo(v string)`

SetTo sets To field to given value.


### GetToken

`func (o *UserBill) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserBill) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserBill) SetToken(v string)`

SetToken sets Token field to given value.


### GetAmount

`func (o *UserBill) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UserBill) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UserBill) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTokenF

`func (o *UserBill) GetTokenF() string`

GetTokenF returns the TokenF field if non-nil, zero value otherwise.

### GetTokenFOk

`func (o *UserBill) GetTokenFOk() (*string, bool)`

GetTokenFOk returns a tuple with the TokenF field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenF

`func (o *UserBill) SetTokenF(v string)`

SetTokenF sets TokenF field to given value.


### GetAmountF

`func (o *UserBill) GetAmountF() string`

GetAmountF returns the AmountF field if non-nil, zero value otherwise.

### GetAmountFOk

`func (o *UserBill) GetAmountFOk() (*string, bool)`

GetAmountFOk returns a tuple with the AmountF field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountF

`func (o *UserBill) SetAmountF(v string)`

SetAmountF sets AmountF field to given value.


### GetStatus

`func (o *UserBill) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserBill) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserBill) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTxHash

`func (o *UserBill) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *UserBill) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *UserBill) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetBillType

`func (o *UserBill) GetBillType() string`

GetBillType returns the BillType field if non-nil, zero value otherwise.

### GetBillTypeOk

`func (o *UserBill) GetBillTypeOk() (*string, bool)`

GetBillTypeOk returns a tuple with the BillType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillType

`func (o *UserBill) SetBillType(v string)`

SetBillType sets BillType field to given value.


### GetIncome

`func (o *UserBill) GetIncome() bool`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *UserBill) GetIncomeOk() (*bool, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *UserBill) SetIncome(v bool)`

SetIncome sets Income field to given value.


### GetTimestamp

`func (o *UserBill) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserBill) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserBill) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetMemo

`func (o *UserBill) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *UserBill) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *UserBill) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetPrice

`func (o *UserBill) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UserBill) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UserBill) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetTransferType

`func (o *UserBill) GetTransferType() string`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *UserBill) GetTransferTypeOk() (*string, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *UserBill) SetTransferType(v string)`

SetTransferType sets TransferType field to given value.


### GetHash

`func (o *UserBill) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *UserBill) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *UserBill) SetHash(v string)`

SetHash sets Hash field to given value.


### GetAccountId

`func (o *UserBill) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserBill) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserBill) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetTokenId

`func (o *UserBill) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *UserBill) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *UserBill) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetStorageId

`func (o *UserBill) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *UserBill) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *UserBill) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.


### GetLabel

`func (o *UserBill) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UserBill) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UserBill) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


