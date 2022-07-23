# GetUserNftDepositResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNum** | **int64** | field.totalNum.description | 
**Deposits** | [**[]NftDepositData**](NftDepositData.md) | field.GetUserNftDepositResponse.NftDepositData | 

## Methods

### NewGetUserNftDepositResponse

`func NewGetUserNftDepositResponse(totalNum int64, deposits []NftDepositData, ) *GetUserNftDepositResponse`

NewGetUserNftDepositResponse instantiates a new GetUserNftDepositResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserNftDepositResponseWithDefaults

`func NewGetUserNftDepositResponseWithDefaults() *GetUserNftDepositResponse`

NewGetUserNftDepositResponseWithDefaults instantiates a new GetUserNftDepositResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNum

`func (o *GetUserNftDepositResponse) GetTotalNum() int64`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *GetUserNftDepositResponse) GetTotalNumOk() (*int64, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *GetUserNftDepositResponse) SetTotalNum(v int64)`

SetTotalNum sets TotalNum field to given value.


### GetDeposits

`func (o *GetUserNftDepositResponse) GetDeposits() []NftDepositData`

GetDeposits returns the Deposits field if non-nil, zero value otherwise.

### GetDepositsOk

`func (o *GetUserNftDepositResponse) GetDepositsOk() (*[]NftDepositData, bool)`

GetDepositsOk returns a tuple with the Deposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposits

`func (o *GetUserNftDepositResponse) SetDeposits(v []NftDepositData)`

SetDeposits sets Deposits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


