# NftTxData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique ID | 
**RequestId** | Pointer to **int64** | Request Id | [optional] 
**Hash** | **string** | hash | 
**TxHash** | **string** | Transaction hash | 
**NftTxType** | **string** | User transaction type | 
**NftData** | **string** | The NFT tokens nftData in this transaction | 
**Amount** | **string** | Amount requested by the user | 
**FeeTokenSymbol** | **string** | fee Token Id | 
**FeeAmount** | **string** | Fee amount in wei | 
**Status** | **string** | Current transaction status | 
**Timestamp** | **int64** | Create time | 
**UpdatedAt** | **int64** | Update time | 
**Memo** | **string** | User memo | 
**ReceiverAddress** | Pointer to **string** | The transfer receiver&#39;s address | [optional] 
**SenderAddress** | Pointer to **string** | field.NftTxData.senderAddress | [optional] 
**Receiver** | Pointer to **int64** | Receiver ID | [optional] 
**StorageInfo** | Pointer to [**StorageInfo**](StorageInfo.md) |  | [optional] 
**WithdrawalInfo** | Pointer to [**WithdrawalInfo**](WithdrawalInfo.md) |  | [optional] 
**MinterInfo** | Pointer to [**MinterInfo**](MinterInfo.md) |  | [optional] 
**NftStatusInfo** | Pointer to [**NftStatusInfo**](NftStatusInfo.md) |  | [optional] 
**Metadata** | Pointer to [**NftMetadataL2Info**](NftMetadataL2Info.md) |  | [optional] 

## Methods

### NewNftTxData

`func NewNftTxData(id int64, hash string, txHash string, nftTxType string, nftData string, amount string, feeTokenSymbol string, feeAmount string, status string, timestamp int64, updatedAt int64, memo string, ) *NftTxData`

NewNftTxData instantiates a new NftTxData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftTxDataWithDefaults

`func NewNftTxDataWithDefaults() *NftTxData`

NewNftTxDataWithDefaults instantiates a new NftTxData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NftTxData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NftTxData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NftTxData) SetId(v int64)`

SetId sets Id field to given value.


### GetRequestId

`func (o *NftTxData) GetRequestId() int64`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *NftTxData) GetRequestIdOk() (*int64, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *NftTxData) SetRequestId(v int64)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *NftTxData) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetHash

`func (o *NftTxData) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *NftTxData) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *NftTxData) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTxHash

`func (o *NftTxData) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *NftTxData) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *NftTxData) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetNftTxType

`func (o *NftTxData) GetNftTxType() string`

GetNftTxType returns the NftTxType field if non-nil, zero value otherwise.

### GetNftTxTypeOk

`func (o *NftTxData) GetNftTxTypeOk() (*string, bool)`

GetNftTxTypeOk returns a tuple with the NftTxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftTxType

`func (o *NftTxData) SetNftTxType(v string)`

SetNftTxType sets NftTxType field to given value.


### GetNftData

`func (o *NftTxData) GetNftData() string`

GetNftData returns the NftData field if non-nil, zero value otherwise.

### GetNftDataOk

`func (o *NftTxData) GetNftDataOk() (*string, bool)`

GetNftDataOk returns a tuple with the NftData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftData

`func (o *NftTxData) SetNftData(v string)`

SetNftData sets NftData field to given value.


### GetAmount

`func (o *NftTxData) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NftTxData) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NftTxData) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFeeTokenSymbol

`func (o *NftTxData) GetFeeTokenSymbol() string`

GetFeeTokenSymbol returns the FeeTokenSymbol field if non-nil, zero value otherwise.

### GetFeeTokenSymbolOk

`func (o *NftTxData) GetFeeTokenSymbolOk() (*string, bool)`

GetFeeTokenSymbolOk returns a tuple with the FeeTokenSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenSymbol

`func (o *NftTxData) SetFeeTokenSymbol(v string)`

SetFeeTokenSymbol sets FeeTokenSymbol field to given value.


### GetFeeAmount

`func (o *NftTxData) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *NftTxData) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *NftTxData) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.


### GetStatus

`func (o *NftTxData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NftTxData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NftTxData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *NftTxData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NftTxData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NftTxData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetUpdatedAt

`func (o *NftTxData) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NftTxData) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NftTxData) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMemo

`func (o *NftTxData) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *NftTxData) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *NftTxData) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetReceiverAddress

`func (o *NftTxData) GetReceiverAddress() string`

GetReceiverAddress returns the ReceiverAddress field if non-nil, zero value otherwise.

### GetReceiverAddressOk

`func (o *NftTxData) GetReceiverAddressOk() (*string, bool)`

GetReceiverAddressOk returns a tuple with the ReceiverAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverAddress

`func (o *NftTxData) SetReceiverAddress(v string)`

SetReceiverAddress sets ReceiverAddress field to given value.

### HasReceiverAddress

`func (o *NftTxData) HasReceiverAddress() bool`

HasReceiverAddress returns a boolean if a field has been set.

### GetSenderAddress

`func (o *NftTxData) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *NftTxData) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *NftTxData) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.

### HasSenderAddress

`func (o *NftTxData) HasSenderAddress() bool`

HasSenderAddress returns a boolean if a field has been set.

### GetReceiver

`func (o *NftTxData) GetReceiver() int64`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *NftTxData) GetReceiverOk() (*int64, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *NftTxData) SetReceiver(v int64)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *NftTxData) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetStorageInfo

`func (o *NftTxData) GetStorageInfo() StorageInfo`

GetStorageInfo returns the StorageInfo field if non-nil, zero value otherwise.

### GetStorageInfoOk

`func (o *NftTxData) GetStorageInfoOk() (*StorageInfo, bool)`

GetStorageInfoOk returns a tuple with the StorageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInfo

`func (o *NftTxData) SetStorageInfo(v StorageInfo)`

SetStorageInfo sets StorageInfo field to given value.

### HasStorageInfo

`func (o *NftTxData) HasStorageInfo() bool`

HasStorageInfo returns a boolean if a field has been set.

### GetWithdrawalInfo

`func (o *NftTxData) GetWithdrawalInfo() WithdrawalInfo`

GetWithdrawalInfo returns the WithdrawalInfo field if non-nil, zero value otherwise.

### GetWithdrawalInfoOk

`func (o *NftTxData) GetWithdrawalInfoOk() (*WithdrawalInfo, bool)`

GetWithdrawalInfoOk returns a tuple with the WithdrawalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalInfo

`func (o *NftTxData) SetWithdrawalInfo(v WithdrawalInfo)`

SetWithdrawalInfo sets WithdrawalInfo field to given value.

### HasWithdrawalInfo

`func (o *NftTxData) HasWithdrawalInfo() bool`

HasWithdrawalInfo returns a boolean if a field has been set.

### GetMinterInfo

`func (o *NftTxData) GetMinterInfo() MinterInfo`

GetMinterInfo returns the MinterInfo field if non-nil, zero value otherwise.

### GetMinterInfoOk

`func (o *NftTxData) GetMinterInfoOk() (*MinterInfo, bool)`

GetMinterInfoOk returns a tuple with the MinterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinterInfo

`func (o *NftTxData) SetMinterInfo(v MinterInfo)`

SetMinterInfo sets MinterInfo field to given value.

### HasMinterInfo

`func (o *NftTxData) HasMinterInfo() bool`

HasMinterInfo returns a boolean if a field has been set.

### GetNftStatusInfo

`func (o *NftTxData) GetNftStatusInfo() NftStatusInfo`

GetNftStatusInfo returns the NftStatusInfo field if non-nil, zero value otherwise.

### GetNftStatusInfoOk

`func (o *NftTxData) GetNftStatusInfoOk() (*NftStatusInfo, bool)`

GetNftStatusInfoOk returns a tuple with the NftStatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftStatusInfo

`func (o *NftTxData) SetNftStatusInfo(v NftStatusInfo)`

SetNftStatusInfo sets NftStatusInfo field to given value.

### HasNftStatusInfo

`func (o *NftTxData) HasNftStatusInfo() bool`

HasNftStatusInfo returns a boolean if a field has been set.

### GetMetadata

`func (o *NftTxData) GetMetadata() NftMetadataL2Info`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NftTxData) GetMetadataOk() (*NftMetadataL2Info, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NftTxData) SetMetadata(v NftMetadataL2Info)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NftTxData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


