# TransactionBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxType** | **string** | The txType of the transaction, one in [Noop,Deposit,Withdraw,Transfer,SpotTrade,AccountUpdate,AmmUpdate,JoinAmm,ExitAmm,SignatureVerification,NftMint,NftData] | 
**AccountId** | Pointer to **int64** | The accountId of the transaction | [optional] 
**Owner** | Pointer to **string** | The owner of the transaction | [optional] 
**Token** | Pointer to [**Token**](Token.md) |  | [optional] 
**ToToken** | Pointer to [**Token**](Token.md) |  | [optional] 
**Fee** | Pointer to [**Token**](Token.md) |  | [optional] 
**ValidUntil** | Pointer to **int64** | The validUntil of the transaction | [optional] 
**ToAccountId** | Pointer to **int64** | The toAccountId of the transaction if tx has a destination account | [optional] 
**ToAccountAddress** | Pointer to **string** | The toAccountAddress of the transaction if tx has a destination account | [optional] 
**StorageId** | Pointer to **int64** | The storageId of the transaction | [optional] 
**OrderA** | Pointer to [**Order**](Order.md) |  | [optional] 
**OrderB** | Pointer to [**Order**](Order.md) |  | [optional] 
**Valid** | Pointer to **bool** | The validness of the transaction | [optional] 
**Nonce** | Pointer to **int32** | The nonce of the transaction if it uses nonce | [optional] 
**MinterAccountId** | Pointer to **int64** | The minterAccountId of the transaction if its a mint tx | [optional] 
**Minter** | Pointer to **string** | The minter of the transaction if its a mint tx | [optional] 
**NftToken** | Pointer to [**Token**](Token.md) |  | [optional] 
**NftType** | Pointer to **string** | The nftType of the transaction if its a mint tx | [optional] 
**FromAddress** | Pointer to **string** | field.TransactionBlock.fromAddress | [optional] 
**ToAddress** | Pointer to **string** | field.TransactionBlock.toAddress | [optional] 
**SpotTradeInfo** | Pointer to [**SpotTradeInfo**](SpotTradeInfo.md) |  | [optional] 

## Methods

### NewTransactionBlock

`func NewTransactionBlock(txType string, ) *TransactionBlock`

NewTransactionBlock instantiates a new TransactionBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionBlockWithDefaults

`func NewTransactionBlockWithDefaults() *TransactionBlock`

NewTransactionBlockWithDefaults instantiates a new TransactionBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxType

`func (o *TransactionBlock) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *TransactionBlock) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *TransactionBlock) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetAccountId

`func (o *TransactionBlock) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransactionBlock) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransactionBlock) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TransactionBlock) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetOwner

`func (o *TransactionBlock) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TransactionBlock) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TransactionBlock) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *TransactionBlock) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetToken

`func (o *TransactionBlock) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TransactionBlock) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TransactionBlock) SetToken(v Token)`

SetToken sets Token field to given value.

### HasToken

`func (o *TransactionBlock) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetToToken

`func (o *TransactionBlock) GetToToken() Token`

GetToToken returns the ToToken field if non-nil, zero value otherwise.

### GetToTokenOk

`func (o *TransactionBlock) GetToTokenOk() (*Token, bool)`

GetToTokenOk returns a tuple with the ToToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToToken

`func (o *TransactionBlock) SetToToken(v Token)`

SetToToken sets ToToken field to given value.

### HasToToken

`func (o *TransactionBlock) HasToToken() bool`

HasToToken returns a boolean if a field has been set.

### GetFee

`func (o *TransactionBlock) GetFee() Token`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TransactionBlock) GetFeeOk() (*Token, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TransactionBlock) SetFee(v Token)`

SetFee sets Fee field to given value.

### HasFee

`func (o *TransactionBlock) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetValidUntil

`func (o *TransactionBlock) GetValidUntil() int64`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *TransactionBlock) GetValidUntilOk() (*int64, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *TransactionBlock) SetValidUntil(v int64)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *TransactionBlock) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetToAccountId

`func (o *TransactionBlock) GetToAccountId() int64`

GetToAccountId returns the ToAccountId field if non-nil, zero value otherwise.

### GetToAccountIdOk

`func (o *TransactionBlock) GetToAccountIdOk() (*int64, bool)`

GetToAccountIdOk returns a tuple with the ToAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccountId

`func (o *TransactionBlock) SetToAccountId(v int64)`

SetToAccountId sets ToAccountId field to given value.

### HasToAccountId

`func (o *TransactionBlock) HasToAccountId() bool`

HasToAccountId returns a boolean if a field has been set.

### GetToAccountAddress

`func (o *TransactionBlock) GetToAccountAddress() string`

GetToAccountAddress returns the ToAccountAddress field if non-nil, zero value otherwise.

### GetToAccountAddressOk

`func (o *TransactionBlock) GetToAccountAddressOk() (*string, bool)`

GetToAccountAddressOk returns a tuple with the ToAccountAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccountAddress

`func (o *TransactionBlock) SetToAccountAddress(v string)`

SetToAccountAddress sets ToAccountAddress field to given value.

### HasToAccountAddress

`func (o *TransactionBlock) HasToAccountAddress() bool`

HasToAccountAddress returns a boolean if a field has been set.

### GetStorageId

`func (o *TransactionBlock) GetStorageId() int64`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *TransactionBlock) GetStorageIdOk() (*int64, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *TransactionBlock) SetStorageId(v int64)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *TransactionBlock) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetOrderA

`func (o *TransactionBlock) GetOrderA() Order`

GetOrderA returns the OrderA field if non-nil, zero value otherwise.

### GetOrderAOk

`func (o *TransactionBlock) GetOrderAOk() (*Order, bool)`

GetOrderAOk returns a tuple with the OrderA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderA

`func (o *TransactionBlock) SetOrderA(v Order)`

SetOrderA sets OrderA field to given value.

### HasOrderA

`func (o *TransactionBlock) HasOrderA() bool`

HasOrderA returns a boolean if a field has been set.

### GetOrderB

`func (o *TransactionBlock) GetOrderB() Order`

GetOrderB returns the OrderB field if non-nil, zero value otherwise.

### GetOrderBOk

`func (o *TransactionBlock) GetOrderBOk() (*Order, bool)`

GetOrderBOk returns a tuple with the OrderB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderB

`func (o *TransactionBlock) SetOrderB(v Order)`

SetOrderB sets OrderB field to given value.

### HasOrderB

`func (o *TransactionBlock) HasOrderB() bool`

HasOrderB returns a boolean if a field has been set.

### GetValid

`func (o *TransactionBlock) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *TransactionBlock) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *TransactionBlock) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *TransactionBlock) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetNonce

`func (o *TransactionBlock) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TransactionBlock) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TransactionBlock) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *TransactionBlock) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetMinterAccountId

`func (o *TransactionBlock) GetMinterAccountId() int64`

GetMinterAccountId returns the MinterAccountId field if non-nil, zero value otherwise.

### GetMinterAccountIdOk

`func (o *TransactionBlock) GetMinterAccountIdOk() (*int64, bool)`

GetMinterAccountIdOk returns a tuple with the MinterAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinterAccountId

`func (o *TransactionBlock) SetMinterAccountId(v int64)`

SetMinterAccountId sets MinterAccountId field to given value.

### HasMinterAccountId

`func (o *TransactionBlock) HasMinterAccountId() bool`

HasMinterAccountId returns a boolean if a field has been set.

### GetMinter

`func (o *TransactionBlock) GetMinter() string`

GetMinter returns the Minter field if non-nil, zero value otherwise.

### GetMinterOk

`func (o *TransactionBlock) GetMinterOk() (*string, bool)`

GetMinterOk returns a tuple with the Minter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinter

`func (o *TransactionBlock) SetMinter(v string)`

SetMinter sets Minter field to given value.

### HasMinter

`func (o *TransactionBlock) HasMinter() bool`

HasMinter returns a boolean if a field has been set.

### GetNftToken

`func (o *TransactionBlock) GetNftToken() Token`

GetNftToken returns the NftToken field if non-nil, zero value otherwise.

### GetNftTokenOk

`func (o *TransactionBlock) GetNftTokenOk() (*Token, bool)`

GetNftTokenOk returns a tuple with the NftToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftToken

`func (o *TransactionBlock) SetNftToken(v Token)`

SetNftToken sets NftToken field to given value.

### HasNftToken

`func (o *TransactionBlock) HasNftToken() bool`

HasNftToken returns a boolean if a field has been set.

### GetNftType

`func (o *TransactionBlock) GetNftType() string`

GetNftType returns the NftType field if non-nil, zero value otherwise.

### GetNftTypeOk

`func (o *TransactionBlock) GetNftTypeOk() (*string, bool)`

GetNftTypeOk returns a tuple with the NftType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftType

`func (o *TransactionBlock) SetNftType(v string)`

SetNftType sets NftType field to given value.

### HasNftType

`func (o *TransactionBlock) HasNftType() bool`

HasNftType returns a boolean if a field has been set.

### GetFromAddress

`func (o *TransactionBlock) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *TransactionBlock) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *TransactionBlock) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *TransactionBlock) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetToAddress

`func (o *TransactionBlock) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *TransactionBlock) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *TransactionBlock) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *TransactionBlock) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetSpotTradeInfo

`func (o *TransactionBlock) GetSpotTradeInfo() SpotTradeInfo`

GetSpotTradeInfo returns the SpotTradeInfo field if non-nil, zero value otherwise.

### GetSpotTradeInfoOk

`func (o *TransactionBlock) GetSpotTradeInfoOk() (*SpotTradeInfo, bool)`

GetSpotTradeInfoOk returns a tuple with the SpotTradeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotTradeInfo

`func (o *TransactionBlock) SetSpotTradeInfo(v SpotTradeInfo)`

SetSpotTradeInfo sets SpotTradeInfo field to given value.

### HasSpotTradeInfo

`func (o *TransactionBlock) HasSpotTradeInfo() bool`

HasSpotTradeInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


