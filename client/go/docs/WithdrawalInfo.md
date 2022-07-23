# WithdrawalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | **string** |  | 
**FastStatus** | **string** |  | 
**DistributeHash** | **string** |  | 

## Methods

### NewWithdrawalInfo

`func NewWithdrawalInfo(recipient string, fastStatus string, distributeHash string, ) *WithdrawalInfo`

NewWithdrawalInfo instantiates a new WithdrawalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawalInfoWithDefaults

`func NewWithdrawalInfoWithDefaults() *WithdrawalInfo`

NewWithdrawalInfoWithDefaults instantiates a new WithdrawalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *WithdrawalInfo) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *WithdrawalInfo) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *WithdrawalInfo) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetFastStatus

`func (o *WithdrawalInfo) GetFastStatus() string`

GetFastStatus returns the FastStatus field if non-nil, zero value otherwise.

### GetFastStatusOk

`func (o *WithdrawalInfo) GetFastStatusOk() (*string, bool)`

GetFastStatusOk returns a tuple with the FastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastStatus

`func (o *WithdrawalInfo) SetFastStatus(v string)`

SetFastStatus sets FastStatus field to given value.


### GetDistributeHash

`func (o *WithdrawalInfo) GetDistributeHash() string`

GetDistributeHash returns the DistributeHash field if non-nil, zero value otherwise.

### GetDistributeHashOk

`func (o *WithdrawalInfo) GetDistributeHashOk() (*string, bool)`

GetDistributeHashOk returns a tuple with the DistributeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeHash

`func (o *WithdrawalInfo) SetDistributeHash(v string)`

SetDistributeHash sets DistributeHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


