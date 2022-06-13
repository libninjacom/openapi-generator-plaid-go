# WalletTransactionExecuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**IdempotencyKey** | **string** | A random key provided by the client, per unique wallet transaction. Maximum of 128 characters.  The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. If a request to execute a wallet transaction fails due to a network connection error, then after a minimum delay of one minute, you can retry the request with the same idempotency key to guarantee that only a single wallet transaction is created. If the request was successfully processed, it will prevent any transaction that uses the same idempotency key, and was received within 24 hours of the first request, from being processed. | 
**WalletId** | **string** | The ID of the e-wallet to debit from | 
**Counterparty** | [**WalletTransactionCounterparty**](WalletTransactionCounterparty.md) |  | 
**Amount** | [**WalletTransactionAmount**](WalletTransactionAmount.md) |  | 
**Reference** | **string** | A reference for the transaction. This must be an alphanumeric string with at most 18 characters and must not contain any special characters or spaces. | 

## Methods

### NewWalletTransactionExecuteRequest

`func NewWalletTransactionExecuteRequest(idempotencyKey string, walletId string, counterparty WalletTransactionCounterparty, amount WalletTransactionAmount, reference string, ) *WalletTransactionExecuteRequest`

NewWalletTransactionExecuteRequest instantiates a new WalletTransactionExecuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionExecuteRequestWithDefaults

`func NewWalletTransactionExecuteRequestWithDefaults() *WalletTransactionExecuteRequest`

NewWalletTransactionExecuteRequestWithDefaults instantiates a new WalletTransactionExecuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *WalletTransactionExecuteRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *WalletTransactionExecuteRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *WalletTransactionExecuteRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *WalletTransactionExecuteRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *WalletTransactionExecuteRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WalletTransactionExecuteRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WalletTransactionExecuteRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WalletTransactionExecuteRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *WalletTransactionExecuteRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *WalletTransactionExecuteRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *WalletTransactionExecuteRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetWalletId

`func (o *WalletTransactionExecuteRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *WalletTransactionExecuteRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *WalletTransactionExecuteRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetCounterparty

`func (o *WalletTransactionExecuteRequest) GetCounterparty() WalletTransactionCounterparty`

GetCounterparty returns the Counterparty field if non-nil, zero value otherwise.

### GetCounterpartyOk

`func (o *WalletTransactionExecuteRequest) GetCounterpartyOk() (*WalletTransactionCounterparty, bool)`

GetCounterpartyOk returns a tuple with the Counterparty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparty

`func (o *WalletTransactionExecuteRequest) SetCounterparty(v WalletTransactionCounterparty)`

SetCounterparty sets Counterparty field to given value.


### GetAmount

`func (o *WalletTransactionExecuteRequest) GetAmount() WalletTransactionAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WalletTransactionExecuteRequest) GetAmountOk() (*WalletTransactionAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WalletTransactionExecuteRequest) SetAmount(v WalletTransactionAmount)`

SetAmount sets Amount field to given value.


### GetReference

`func (o *WalletTransactionExecuteRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *WalletTransactionExecuteRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *WalletTransactionExecuteRequest) SetReference(v string)`

SetReference sets Reference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


