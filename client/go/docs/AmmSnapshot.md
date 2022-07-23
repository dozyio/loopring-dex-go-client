# AmmSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolName** | **string** | field.AmmSnapshot.poolName | 
**PoolAddress** | **string** | field.AmmSnapshot.poolAddress | 
**PoolTokenId** | **int32** | field.AmmSnapshot.poolTokenId | 
**PoolTokenAmount** | **string** | field.AmmSnapshot.PoolTokenAmount | 
**TokenIds** | **[]map[string]interface{}** | field.AmmSnapshot.tokenIds | 
**TokenAmounts** | **[]string** | field.AmmSnapshot.tokenAmounts | 
**Risky** | **bool** | field.AmmSnapshot.risky | 

## Methods

### NewAmmSnapshot

`func NewAmmSnapshot(poolName string, poolAddress string, poolTokenId int32, poolTokenAmount string, tokenIds []map[string]interface{}, tokenAmounts []string, risky bool, ) *AmmSnapshot`

NewAmmSnapshot instantiates a new AmmSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmmSnapshotWithDefaults

`func NewAmmSnapshotWithDefaults() *AmmSnapshot`

NewAmmSnapshotWithDefaults instantiates a new AmmSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolName

`func (o *AmmSnapshot) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AmmSnapshot) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AmmSnapshot) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.


### GetPoolAddress

`func (o *AmmSnapshot) GetPoolAddress() string`

GetPoolAddress returns the PoolAddress field if non-nil, zero value otherwise.

### GetPoolAddressOk

`func (o *AmmSnapshot) GetPoolAddressOk() (*string, bool)`

GetPoolAddressOk returns a tuple with the PoolAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAddress

`func (o *AmmSnapshot) SetPoolAddress(v string)`

SetPoolAddress sets PoolAddress field to given value.


### GetPoolTokenId

`func (o *AmmSnapshot) GetPoolTokenId() int32`

GetPoolTokenId returns the PoolTokenId field if non-nil, zero value otherwise.

### GetPoolTokenIdOk

`func (o *AmmSnapshot) GetPoolTokenIdOk() (*int32, bool)`

GetPoolTokenIdOk returns a tuple with the PoolTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolTokenId

`func (o *AmmSnapshot) SetPoolTokenId(v int32)`

SetPoolTokenId sets PoolTokenId field to given value.


### GetPoolTokenAmount

`func (o *AmmSnapshot) GetPoolTokenAmount() string`

GetPoolTokenAmount returns the PoolTokenAmount field if non-nil, zero value otherwise.

### GetPoolTokenAmountOk

`func (o *AmmSnapshot) GetPoolTokenAmountOk() (*string, bool)`

GetPoolTokenAmountOk returns a tuple with the PoolTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolTokenAmount

`func (o *AmmSnapshot) SetPoolTokenAmount(v string)`

SetPoolTokenAmount sets PoolTokenAmount field to given value.


### GetTokenIds

`func (o *AmmSnapshot) GetTokenIds() []map[string]interface{}`

GetTokenIds returns the TokenIds field if non-nil, zero value otherwise.

### GetTokenIdsOk

`func (o *AmmSnapshot) GetTokenIdsOk() (*[]map[string]interface{}, bool)`

GetTokenIdsOk returns a tuple with the TokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIds

`func (o *AmmSnapshot) SetTokenIds(v []map[string]interface{})`

SetTokenIds sets TokenIds field to given value.


### GetTokenAmounts

`func (o *AmmSnapshot) GetTokenAmounts() []string`

GetTokenAmounts returns the TokenAmounts field if non-nil, zero value otherwise.

### GetTokenAmountsOk

`func (o *AmmSnapshot) GetTokenAmountsOk() (*[]string, bool)`

GetTokenAmountsOk returns a tuple with the TokenAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAmounts

`func (o *AmmSnapshot) SetTokenAmounts(v []string)`

SetTokenAmounts sets TokenAmounts field to given value.


### GetRisky

`func (o *AmmSnapshot) GetRisky() bool`

GetRisky returns the Risky field if non-nil, zero value otherwise.

### GetRiskyOk

`func (o *AmmSnapshot) GetRiskyOk() (*bool, bool)`

GetRiskyOk returns a tuple with the Risky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisky

`func (o *AmmSnapshot) SetRisky(v bool)`

SetRisky sets Risky field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


