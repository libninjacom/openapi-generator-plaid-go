# WalletTransactionsListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**WalletId** | **string** | The ID of the e-wallet to fetch transactions from | 
**Cursor** | Pointer to **string** | A base64 value representing the latest transaction that has already been requested. Set this to &#x60;next_cursor&#x60; received from the previous &#x60;/wallet/transactions/list&#x60; request. If provided, the response will only contain transactions created before that transaction. If omitted, the response will contain transactions starting from the most recent, and in descending order by the &#x60;created_at&#x60; time. | [optional] 
**Count** | Pointer to **int32** | The number of transactions to fetch | [optional] [default to 10]

## Methods

### NewWalletTransactionsListRequest

`func NewWalletTransactionsListRequest(walletId string, ) *WalletTransactionsListRequest`

NewWalletTransactionsListRequest instantiates a new WalletTransactionsListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionsListRequestWithDefaults

`func NewWalletTransactionsListRequestWithDefaults() *WalletTransactionsListRequest`

NewWalletTransactionsListRequestWithDefaults instantiates a new WalletTransactionsListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *WalletTransactionsListRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *WalletTransactionsListRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *WalletTransactionsListRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *WalletTransactionsListRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *WalletTransactionsListRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WalletTransactionsListRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WalletTransactionsListRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WalletTransactionsListRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetWalletId

`func (o *WalletTransactionsListRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *WalletTransactionsListRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *WalletTransactionsListRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetCursor

`func (o *WalletTransactionsListRequest) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *WalletTransactionsListRequest) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *WalletTransactionsListRequest) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *WalletTransactionsListRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetCount

`func (o *WalletTransactionsListRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WalletTransactionsListRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WalletTransactionsListRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *WalletTransactionsListRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


