# GetAccountResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**Data** | [**DexAccount**](DexAccount.md) |  | 

## Methods

### NewGetAccountResponseV2

`func NewGetAccountResponseV2(resultInfo ResultInfo, data DexAccount, ) *GetAccountResponseV2`

NewGetAccountResponseV2 instantiates a new GetAccountResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountResponseV2WithDefaults

`func NewGetAccountResponseV2WithDefaults() *GetAccountResponseV2`

NewGetAccountResponseV2WithDefaults instantiates a new GetAccountResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *GetAccountResponseV2) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *GetAccountResponseV2) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *GetAccountResponseV2) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetData

`func (o *GetAccountResponseV2) GetData() DexAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAccountResponseV2) GetDataOk() (*DexAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAccountResponseV2) SetData(v DexAccount)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


