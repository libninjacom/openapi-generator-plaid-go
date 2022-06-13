# WalletTransactionsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | [**[]WalletTransaction**](WalletTransaction.md) | An array of transactions of an e-wallet, associated with the given &#x60;wallet_id&#x60; | 
**NextCursor** | Pointer to **string** | Cursor used for fetching transactions created before the latest transaction provided in this response | [optional] 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewWalletTransactionsListResponse

`func NewWalletTransactionsListResponse(transactions []WalletTransaction, requestId string, ) *WalletTransactionsListResponse`

NewWalletTransactionsListResponse instantiates a new WalletTransactionsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionsListResponseWithDefaults

`func NewWalletTransactionsListResponseWithDefaults() *WalletTransactionsListResponse`

NewWalletTransactionsListResponseWithDefaults instantiates a new WalletTransactionsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *WalletTransactionsListResponse) GetTransactions() []WalletTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *WalletTransactionsListResponse) GetTransactionsOk() (*[]WalletTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *WalletTransactionsListResponse) SetTransactions(v []WalletTransaction)`

SetTransactions sets Transactions field to given value.


### GetNextCursor

`func (o *WalletTransactionsListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *WalletTransactionsListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *WalletTransactionsListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *WalletTransactionsListResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetRequestId

`func (o *WalletTransactionsListResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WalletTransactionsListResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WalletTransactionsListResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


