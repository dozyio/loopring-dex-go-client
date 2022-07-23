# AmmTransferDataV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **int32** | AMM transfer token id | 
**Amount** | **string** | AMM transfer token amount | 
**ActualAmount** | **string** | Actual AMM transfer token amount due to precision processing | 
**FeeAmount** | **string** | Actual AMM transfer fee amount | 

## Methods

### NewAmmTransferDataV3

`func NewAmmTransferDataV3(tokenId int32, amount string, actualAmount string, feeAmount string, ) *AmmTransferDataV3`

NewAmmTransferDataV3 instantiates a new AmmTransferDataV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmTransferDataV3WithDefaults

`func NewAmmTransferDataV3WithDefaults() *AmmTransferDataV3`

NewAmmTransferDataV3WithDefaults instantiates a new AmmTransferDataV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *AmmTransferDataV3) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AmmTransferDataV3) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AmmTransferDataV3) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *AmmTransferDataV3) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AmmTransferDataV3) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AmmTransferDataV3) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetActualAmount

`func (o *AmmTransferDataV3) GetActualAmount() string`

GetActualAmount returns the ActualAmount field if non-nil, zero value otherwise.

### GetActualAmountOk

`func (o *AmmTransferDataV3) GetActualAmountOk() (*string, bool)`

GetActualAmountOk returns a tuple with the ActualAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualAmount

`func (o *AmmTransferDataV3) SetActualAmount(v string)`

SetActualAmount sets ActualAmount field to given value.


### GetFeeAmount

`func (o *AmmTransferDataV3) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *AmmTransferDataV3) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *AmmTransferDataV3) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


