# ExchangeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | **int32** | Loopring&#39;s smart contract network ID. | 
**ExchangeAddress** | **string** | Contract address of exchange. | 
**DepositAddress** | **string** | field.ExchangeInfo.depositAddress | 
**OnchainFees** | [**[]FeeInfo**](FeeInfo.md) | Fees settings. | 
**OpenAccountFees** | [**[]OffFeeInfo**](OffFeeInfo.md) | field.ExchangeInfo.openAccountFee | 
**UpdateFees** | [**[]OffFeeInfo**](OffFeeInfo.md) | field.ExchangeInfo.updateFees | 
**TransferFees** | [**[]OffFeeInfo**](OffFeeInfo.md) | Transfer fee settings. | 
**WithdrawalFees** | [**[]OffFeeInfo**](OffFeeInfo.md) | Off-chain withdrawal fee settings. | 
**FastWithdrawalFees** | [**[]OffFeeInfo**](OffFeeInfo.md) | fast withdrawal fee settings. | 
**AmmExitFees** | [**[]OffFeeInfo**](OffFeeInfo.md) | AMM pool exit fee settings. | 

## Methods

### NewExchangeInfo

`func NewExchangeInfo(chainId int32, exchangeAddress string, depositAddress string, onchainFees []FeeInfo, openAccountFees []OffFeeInfo, updateFees []OffFeeInfo, transferFees []OffFeeInfo, withdrawalFees []OffFeeInfo, fastWithdrawalFees []OffFeeInfo, ammExitFees []OffFeeInfo, ) *ExchangeInfo`

NewExchangeInfo instantiates a new ExchangeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeInfoWithDefaults

`func NewExchangeInfoWithDefaults() *ExchangeInfo`

NewExchangeInfoWithDefaults instantiates a new ExchangeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *ExchangeInfo) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ExchangeInfo) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ExchangeInfo) SetChainId(v int32)`

SetChainId sets ChainId field to given value.


### GetExchangeAddress

`func (o *ExchangeInfo) GetExchangeAddress() string`

GetExchangeAddress returns the ExchangeAddress field if non-nil, zero value otherwise.

### GetExchangeAddressOk

`func (o *ExchangeInfo) GetExchangeAddressOk() (*string, bool)`

GetExchangeAddressOk returns a tuple with the ExchangeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeAddress

`func (o *ExchangeInfo) SetExchangeAddress(v string)`

SetExchangeAddress sets ExchangeAddress field to given value.


### GetDepositAddress

`func (o *ExchangeInfo) GetDepositAddress() string`

GetDepositAddress returns the DepositAddress field if non-nil, zero value otherwise.

### GetDepositAddressOk

`func (o *ExchangeInfo) GetDepositAddressOk() (*string, bool)`

GetDepositAddressOk returns a tuple with the DepositAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAddress

`func (o *ExchangeInfo) SetDepositAddress(v string)`

SetDepositAddress sets DepositAddress field to given value.


### GetOnchainFees

`func (o *ExchangeInfo) GetOnchainFees() []FeeInfo`

GetOnchainFees returns the OnchainFees field if non-nil, zero value otherwise.

### GetOnchainFeesOk

`func (o *ExchangeInfo) GetOnchainFeesOk() (*[]FeeInfo, bool)`

GetOnchainFeesOk returns a tuple with the OnchainFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnchainFees

`func (o *ExchangeInfo) SetOnchainFees(v []FeeInfo)`

SetOnchainFees sets OnchainFees field to given value.


### GetOpenAccountFees

`func (o *ExchangeInfo) GetOpenAccountFees() []OffFeeInfo`

GetOpenAccountFees returns the OpenAccountFees field if non-nil, zero value otherwise.

### GetOpenAccountFeesOk

`func (o *ExchangeInfo) GetOpenAccountFeesOk() (*[]OffFeeInfo, bool)`

GetOpenAccountFeesOk returns a tuple with the OpenAccountFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAccountFees

`func (o *ExchangeInfo) SetOpenAccountFees(v []OffFeeInfo)`

SetOpenAccountFees sets OpenAccountFees field to given value.


### GetUpdateFees

`func (o *ExchangeInfo) GetUpdateFees() []OffFeeInfo`

GetUpdateFees returns the UpdateFees field if non-nil, zero value otherwise.

### GetUpdateFeesOk

`func (o *ExchangeInfo) GetUpdateFeesOk() (*[]OffFeeInfo, bool)`

GetUpdateFeesOk returns a tuple with the UpdateFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFees

`func (o *ExchangeInfo) SetUpdateFees(v []OffFeeInfo)`

SetUpdateFees sets UpdateFees field to given value.


### GetTransferFees

`func (o *ExchangeInfo) GetTransferFees() []OffFeeInfo`

GetTransferFees returns the TransferFees field if non-nil, zero value otherwise.

### GetTransferFeesOk

`func (o *ExchangeInfo) GetTransferFeesOk() (*[]OffFeeInfo, bool)`

GetTransferFeesOk returns a tuple with the TransferFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFees

`func (o *ExchangeInfo) SetTransferFees(v []OffFeeInfo)`

SetTransferFees sets TransferFees field to given value.


### GetWithdrawalFees

`func (o *ExchangeInfo) GetWithdrawalFees() []OffFeeInfo`

GetWithdrawalFees returns the WithdrawalFees field if non-nil, zero value otherwise.

### GetWithdrawalFeesOk

`func (o *ExchangeInfo) GetWithdrawalFeesOk() (*[]OffFeeInfo, bool)`

GetWithdrawalFeesOk returns a tuple with the WithdrawalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalFees

`func (o *ExchangeInfo) SetWithdrawalFees(v []OffFeeInfo)`

SetWithdrawalFees sets WithdrawalFees field to given value.


### GetFastWithdrawalFees

`func (o *ExchangeInfo) GetFastWithdrawalFees() []OffFeeInfo`

GetFastWithdrawalFees returns the FastWithdrawalFees field if non-nil, zero value otherwise.

### GetFastWithdrawalFeesOk

`func (o *ExchangeInfo) GetFastWithdrawalFeesOk() (*[]OffFeeInfo, bool)`

GetFastWithdrawalFeesOk returns a tuple with the FastWithdrawalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastWithdrawalFees

`func (o *ExchangeInfo) SetFastWithdrawalFees(v []OffFeeInfo)`

SetFastWithdrawalFees sets FastWithdrawalFees field to given value.


### GetAmmExitFees

`func (o *ExchangeInfo) GetAmmExitFees() []OffFeeInfo`

GetAmmExitFees returns the AmmExitFees field if non-nil, zero value otherwise.

### GetAmmExitFeesOk

`func (o *ExchangeInfo) GetAmmExitFeesOk() (*[]OffFeeInfo, bool)`

GetAmmExitFeesOk returns a tuple with the AmmExitFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmmExitFees

`func (o *ExchangeInfo) SetAmmExitFees(v []OffFeeInfo)`

SetAmmExitFees sets AmmExitFees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


