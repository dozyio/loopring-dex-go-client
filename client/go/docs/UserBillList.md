# UserBillList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** |  | 
**Bills** | [**[]UserBill**](UserBill.md) |  | 

## Methods

### NewUserBillList

`func NewUserBillList(totalNum int64, bills []UserBill, ) *UserBillList`

NewUserBillList instantiates a new UserBillList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBillListWithDefaults

`func NewUserBillListWithDefaults() *UserBillList`

NewUserBillListWithDefaults instantiates a new UserBillList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *UserBillList) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *UserBillList) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *UserBillList) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetBills

`func (o *UserBillList) GetBills() []UserBill`

GetBills returns the Bills field if non-nil, zero value otherwise.

### GetBillsOk

`func (o *UserBillList) GetBillsOk() (*[]UserBill, bool)`

GetBillsOk returns a tuple with the Bills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBills

`func (o *UserBillList) SetBills(v []UserBill)`

SetBills sets Bills field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


