# OrderVolumesV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseAmount** | **string** | The amount of base tokens involved in the order. | 
**QuoteAmount** | **string** | The amount of quote tokens involved in the order. | 
**BaseFilled** | **string** | The amount of requested base tokens filled in the order. | 
**QuoteFilled** | **string** | The amount of requested quote tokens filled in the order. | 
**Fee** | **string** | The amount of quote or base token amount used to pay for the orders fee. Whether this data refers to the base or quote token, one can find out by looking at the orders side | 

## Methods

### NewOrderVolumesV3

`func NewOrderVolumesV3(baseAmount string, quoteAmount string, baseFilled string, quoteFilled string, fee string, ) *OrderVolumesV3`

NewOrderVolumesV3 instantiates a new OrderVolumesV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderVolumesV3WithDefaults

`func NewOrderVolumesV3WithDefaults() *OrderVolumesV3`

NewOrderVolumesV3WithDefaults instantiates a new OrderVolumesV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseAmount

`func (o *OrderVolumesV3) GetBaseAmount() string`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *OrderVolumesV3) GetBaseAmountOk() (*string, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *OrderVolumesV3) SetBaseAmount(v string)`

SetBaseAmount sets BaseAmount field to given value.


### GetQuoteAmount

`func (o *OrderVolumesV3) GetQuoteAmount() string`

GetQuoteAmount returns the QuoteAmount field if non-nil, zero value otherwise.

### GetQuoteAmountOk

`func (o *OrderVolumesV3) GetQuoteAmountOk() (*string, bool)`

GetQuoteAmountOk returns a tuple with the QuoteAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAmount

`func (o *OrderVolumesV3) SetQuoteAmount(v string)`

SetQuoteAmount sets QuoteAmount field to given value.


### GetBaseFilled

`func (o *OrderVolumesV3) GetBaseFilled() string`

GetBaseFilled returns the BaseFilled field if non-nil, zero value otherwise.

### GetBaseFilledOk

`func (o *OrderVolumesV3) GetBaseFilledOk() (*string, bool)`

GetBaseFilledOk returns a tuple with the BaseFilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFilled

`func (o *OrderVolumesV3) SetBaseFilled(v string)`

SetBaseFilled sets BaseFilled field to given value.


### GetQuoteFilled

`func (o *OrderVolumesV3) GetQuoteFilled() string`

GetQuoteFilled returns the QuoteFilled field if non-nil, zero value otherwise.

### GetQuoteFilledOk

`func (o *OrderVolumesV3) GetQuoteFilledOk() (*string, bool)`

GetQuoteFilledOk returns a tuple with the QuoteFilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteFilled

`func (o *OrderVolumesV3) SetQuoteFilled(v string)`

SetQuoteFilled sets QuoteFilled field to given value.


### GetFee

`func (o *OrderVolumesV3) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *OrderVolumesV3) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *OrderVolumesV3) SetFee(v string)`

SetFee sets Fee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


