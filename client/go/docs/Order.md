# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageID** | **int64** | The storageId of the order | 
**AccountID** | **int64** | The accountID of the order | 
**AmountS** | **string** | The amountS of the order | 
**AmountB** | **string** | The amountB of the order | 
**TokenS** | **int32** | The tokenS of the order | 
**TokenB** | **int32** | The tokenB of the order | 
**ValidUntil** | **int64** | The validUntil of the order | 
**Taker** | **string** | The taker of the order | 
**FeeBips** | **int32** | The feeBips of the order | 
**IsAmm** | **bool** | If the order isAmm | 
**NftData** | **string** | The nftData of the order, if its NFT order | 
**FillS** | **int64** | The fillS of the order | 

## Methods

### NewOrder

`func NewOrder(storageID int64, accountID int64, amountS string, amountB string, tokenS int32, tokenB int32, validUntil int64, taker string, feeBips int32, isAmm bool, nftData string, fillS int64, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageID

`func (o *Order) GetStorageID() int64`

GetStorageID returns the StorageID field if non-nil, zero value otherwise.

### GetStorageIDOk

`func (o *Order) GetStorageIDOk() (*int64, bool)`

GetStorageIDOk returns a tuple with the StorageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageID

`func (o *Order) SetStorageID(v int64)`

SetStorageID sets StorageID field to given value.


### GetAccountID

`func (o *Order) GetAccountID() int64`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *Order) GetAccountIDOk() (*int64, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *Order) SetAccountID(v int64)`

SetAccountID sets AccountID field to given value.


### GetAmountS

`func (o *Order) GetAmountS() string`

GetAmountS returns the AmountS field if non-nil, zero value otherwise.

### GetAmountSOk

`func (o *Order) GetAmountSOk() (*string, bool)`

GetAmountSOk returns a tuple with the AmountS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountS

`func (o *Order) SetAmountS(v string)`

SetAmountS sets AmountS field to given value.


### GetAmountB

`func (o *Order) GetAmountB() string`

GetAmountB returns the AmountB field if non-nil, zero value otherwise.

### GetAmountBOk

`func (o *Order) GetAmountBOk() (*string, bool)`

GetAmountBOk returns a tuple with the AmountB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountB

`func (o *Order) SetAmountB(v string)`

SetAmountB sets AmountB field to given value.


### GetTokenS

`func (o *Order) GetTokenS() int32`

GetTokenS returns the TokenS field if non-nil, zero value otherwise.

### GetTokenSOk

`func (o *Order) GetTokenSOk() (*int32, bool)`

GetTokenSOk returns a tuple with the TokenS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenS

`func (o *Order) SetTokenS(v int32)`

SetTokenS sets TokenS field to given value.


### GetTokenB

`func (o *Order) GetTokenB() int32`

GetTokenB returns the TokenB field if non-nil, zero value otherwise.

### GetTokenBOk

`func (o *Order) GetTokenBOk() (*int32, bool)`

GetTokenBOk returns a tuple with the TokenB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenB

`func (o *Order) SetTokenB(v int32)`

SetTokenB sets TokenB field to given value.


### GetValidUntil

`func (o *Order) GetValidUntil() int64`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *Order) GetValidUntilOk() (*int64, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *Order) SetValidUntil(v int64)`

SetValidUntil sets ValidUntil field to given value.


### GetTaker

`func (o *Order) GetTaker() string`

GetTaker returns the Taker field if non-nil, zero value otherwise.

### GetTakerOk

`func (o *Order) GetTakerOk() (*string, bool)`

GetTakerOk returns a tuple with the Taker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaker

`func (o *Order) SetTaker(v string)`

SetTaker sets Taker field to given value.


### GetFeeBips

`func (o *Order) GetFeeBips() int32`

GetFeeBips returns the FeeBips field if non-nil, zero value otherwise.

### GetFeeBipsOk

`func (o *Order) GetFeeBipsOk() (*int32, bool)`

GetFeeBipsOk returns a tuple with the FeeBips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBips

`func (o *Order) SetFeeBips(v int32)`

SetFeeBips sets FeeBips field to given value.


### GetIsAmm

`func (o *Order) GetIsAmm() bool`

GetIsAmm returns the IsAmm field if non-nil, zero value otherwise.

### GetIsAmmOk

`func (o *Order) GetIsAmmOk() (*bool, bool)`

GetIsAmmOk returns a tuple with the IsAmm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAmm

`func (o *Order) SetIsAmm(v bool)`

SetIsAmm sets IsAmm field to given value.


### GetNftData

`func (o *Order) GetNftData() string`

GetNftData returns the NftData field if non-nil, zero value otherwise.

### GetNftDataOk

`func (o *Order) GetNftDataOk() (*string, bool)`

GetNftDataOk returns a tuple with the NftData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftData

`func (o *Order) SetNftData(v string)`

SetNftData sets NftData field to given value.


### GetFillS

`func (o *Order) GetFillS() int64`

GetFillS returns the FillS field if non-nil, zero value otherwise.

### GetFillSOk

`func (o *Order) GetFillSOk() (*int64, bool)`

GetFillSOk returns a tuple with the FillS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillS

`func (o *Order) SetFillS(v int64)`

SetFillS sets FillS field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


