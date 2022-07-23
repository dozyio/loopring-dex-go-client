# GetUserNftWithdrawalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | field.totalNum.description | 
**Withdrawals** | [**[]NftWithdrawalData**](NftWithdrawalData.md) | field.GetUserNftWithdrawalResponse.NftWithdrawalDataList | 

## Methods

### NewGetUserNftWithdrawalResponse

`func NewGetUserNftWithdrawalResponse(totalNum int64, withdrawals []NftWithdrawalData, ) *GetUserNftWithdrawalResponse`

NewGetUserNftWithdrawalResponse instantiates a new GetUserNftWithdrawalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserNftWithdrawalResponseWithDefaults

`func NewGetUserNftWithdrawalResponseWithDefaults() *GetUserNftWithdrawalResponse`

NewGetUserNftWithdrawalResponseWithDefaults instantiates a new GetUserNftWithdrawalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *GetUserNftWithdrawalResponse) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *GetUserNftWithdrawalResponse) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *GetUserNftWithdrawalResponse) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetWithdrawals

`func (o *GetUserNftWithdrawalResponse) GetWithdrawals() []NftWithdrawalData`

GetWithdrawals returns the Withdrawals field if non-nil, zero value otherwise.

### GetWithdrawalsOk

`func (o *GetUserNftWithdrawalResponse) GetWithdrawalsOk() (*[]NftWithdrawalData, bool)`

GetWithdrawalsOk returns a tuple with the Withdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawals

`func (o *GetUserNftWithdrawalResponse) SetWithdrawals(v []NftWithdrawalData)`

SetWithdrawals sets Withdrawals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


