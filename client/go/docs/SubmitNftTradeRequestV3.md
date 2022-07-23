# SubmitNftTradeRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Maker** | [**SubmitNftOrderRequestV3**](SubmitNftOrderRequestV3.md) |  | 
**MakerFeeBips** | **int32** | The maker feeBips, should &lt;&#x3D; maxFeeBips in makers order | 
**Taker** | [**SubmitNftOrderRequestV3**](SubmitNftOrderRequestV3.md) |  | 
**TakerFeeBips** | **int32** | The taker feeBips, should &lt;&#x3D; maxFeeBips in takers order | 
**IgnoreFees** | Pointer to **map[string]interface{}** | If this trade ignores the fees config of each order, only whitelisted user can set this flag. | [optional] 
**MatchByTaker** | Pointer to **map[string]interface{}** | field.SubmitNftTradeRequestV3.matchByTaker | [optional] 

## Methods

### NewSubmitNftTradeRequestV3

`func NewSubmitNftTradeRequestV3(maker SubmitNftOrderRequestV3, makerFeeBips int32, taker SubmitNftOrderRequestV3, takerFeeBips int32, ) *SubmitNftTradeRequestV3`

NewSubmitNftTradeRequestV3 instantiates a new SubmitNftTradeRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitNftTradeRequestV3WithDefaults

`func NewSubmitNftTradeRequestV3WithDefaults() *SubmitNftTradeRequestV3`

NewSubmitNftTradeRequestV3WithDefaults instantiates a new SubmitNftTradeRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaker

`func (o *SubmitNftTradeRequestV3) GetMaker() SubmitNftOrderRequestV3`

GetMaker returns the Maker field if non-nil, zero value otherwise.

### GetMakerOk

`func (o *SubmitNftTradeRequestV3) GetMakerOk() (*SubmitNftOrderRequestV3, bool)`

GetMakerOk returns a tuple with the Maker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaker

`func (o *SubmitNftTradeRequestV3) SetMaker(v SubmitNftOrderRequestV3)`

SetMaker sets Maker field to given value.


### GetMakerFeeBips

`func (o *SubmitNftTradeRequestV3) GetMakerFeeBips() int32`

GetMakerFeeBips returns the MakerFeeBips field if non-nil, zero value otherwise.

### GetMakerFeeBipsOk

`func (o *SubmitNftTradeRequestV3) GetMakerFeeBipsOk() (*int32, bool)`

GetMakerFeeBipsOk returns a tuple with the MakerFeeBips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerFeeBips

`func (o *SubmitNftTradeRequestV3) SetMakerFeeBips(v int32)`

SetMakerFeeBips sets MakerFeeBips field to given value.


### GetTaker

`func (o *SubmitNftTradeRequestV3) GetTaker() SubmitNftOrderRequestV3`

GetTaker returns the Taker field if non-nil, zero value otherwise.

### GetTakerOk

`func (o *SubmitNftTradeRequestV3) GetTakerOk() (*SubmitNftOrderRequestV3, bool)`

GetTakerOk returns a tuple with the Taker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaker

`func (o *SubmitNftTradeRequestV3) SetTaker(v SubmitNftOrderRequestV3)`

SetTaker sets Taker field to given value.


### GetTakerFeeBips

`func (o *SubmitNftTradeRequestV3) GetTakerFeeBips() int32`

GetTakerFeeBips returns the TakerFeeBips field if non-nil, zero value otherwise.

### GetTakerFeeBipsOk

`func (o *SubmitNftTradeRequestV3) GetTakerFeeBipsOk() (*int32, bool)`

GetTakerFeeBipsOk returns a tuple with the TakerFeeBips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerFeeBips

`func (o *SubmitNftTradeRequestV3) SetTakerFeeBips(v int32)`

SetTakerFeeBips sets TakerFeeBips field to given value.


### GetIgnoreFees

`func (o *SubmitNftTradeRequestV3) GetIgnoreFees() map[string]interface{}`

GetIgnoreFees returns the IgnoreFees field if non-nil, zero value otherwise.

### GetIgnoreFeesOk

`func (o *SubmitNftTradeRequestV3) GetIgnoreFeesOk() (*map[string]interface{}, bool)`

GetIgnoreFeesOk returns a tuple with the IgnoreFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreFees

`func (o *SubmitNftTradeRequestV3) SetIgnoreFees(v map[string]interface{})`

SetIgnoreFees sets IgnoreFees field to given value.

### HasIgnoreFees

`func (o *SubmitNftTradeRequestV3) HasIgnoreFees() bool`

HasIgnoreFees returns a boolean if a field has been set.

### GetMatchByTaker

`func (o *SubmitNftTradeRequestV3) GetMatchByTaker() map[string]interface{}`

GetMatchByTaker returns the MatchByTaker field if non-nil, zero value otherwise.

### GetMatchByTakerOk

`func (o *SubmitNftTradeRequestV3) GetMatchByTakerOk() (*map[string]interface{}, bool)`

GetMatchByTakerOk returns a tuple with the MatchByTaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchByTaker

`func (o *SubmitNftTradeRequestV3) SetMatchByTaker(v map[string]interface{})`

SetMatchByTaker sets MatchByTaker field to given value.

### HasMatchByTaker

`func (o *SubmitNftTradeRequestV3) HasMatchByTaker() bool`

HasMatchByTaker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


