# GetTickerResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | Pointer to **[][]string** | Each item in the list is an array that contains the following: trading pair ID, update timestamp, base token volume, quote token volume, open-price, highest price, lowest price, closing price, number of trades, highest bid price, lowest ask price, base fee amount, quote fee amount. All values are returned as strings. Fee amount is for AMM only. | [optional] 

## Methods

### NewGetTickerResponseV2

`func NewGetTickerResponseV2(resultInfo ResultInfo, ) *GetTickerResponseV2`

NewGetTickerResponseV2 instantiates a new GetTickerResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTickerResponseV2WithDefaults

`func NewGetTickerResponseV2WithDefaults() *GetTickerResponseV2`

NewGetTickerResponseV2WithDefaults instantiates a new GetTickerResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetTickerResponseV2) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetTickerResponseV2) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetTickerResponseV2) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetTickerResponseV2) GetData() [][]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTickerResponseV2) GetDataOk() (*[][]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTickerResponseV2) SetData(v [][]string)`

SetData sets Data field to given value.

### HasData

`func (o *GetTickerResponseV2) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


