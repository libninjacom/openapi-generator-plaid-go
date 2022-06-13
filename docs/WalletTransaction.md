# WalletTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | A unique ID identifying the transaction | 
**Reference** | **string** | A reference for the transaction | 
**Type** | **string** | The type of of the transaction. Currently, only &#x60;\&quot;PAYOUT\&quot;&#x60; is supported. | 
**Amount** | [**WalletTransactionAmount**](WalletTransactionAmount.md) |  | 
**Counterparty** | [**WalletTransactionCounterparty**](WalletTransactionCounterparty.md) |  | 
**Status** | [**WalletTransactionStatus**](WalletTransactionStatus.md) |  | 
**CreatedAt** | **time.Time** | Timestamp when the transaction was created, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format. | 

## Methods

### NewWalletTransaction

`func NewWalletTransaction(transactionId string, reference string, type_ string, amount WalletTransactionAmount, counterparty WalletTransactionCounterparty, status WalletTransactionStatus, createdAt time.Time, ) *WalletTransaction`

NewWalletTransaction instantiates a new WalletTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionWithDefaults

`func NewWalletTransactionWithDefaults() *WalletTransaction`

NewWalletTransactionWithDefaults instantiates a new WalletTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *WalletTransaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *WalletTransaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *WalletTransaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetReference

`func (o *WalletTransaction) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *WalletTransaction) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *WalletTransaction) SetReference(v string)`

SetReference sets Reference field to given value.


### GetType

`func (o *WalletTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WalletTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WalletTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetAmount

`func (o *WalletTransaction) GetAmount() WalletTransactionAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WalletTransaction) GetAmountOk() (*WalletTransactionAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WalletTransaction) SetAmount(v WalletTransactionAmount)`

SetAmount sets Amount field to given value.


### GetCounterparty

`func (o *WalletTransaction) GetCounterparty() WalletTransactionCounterparty`

GetCounterparty returns the Counterparty field if non-nil, zero value otherwise.

### GetCounterpartyOk

`func (o *WalletTransaction) GetCounterpartyOk() (*WalletTransactionCounterparty, bool)`

GetCounterpartyOk returns a tuple with the Counterparty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparty

`func (o *WalletTransaction) SetCounterparty(v WalletTransactionCounterparty)`

SetCounterparty sets Counterparty field to given value.


### GetStatus

`func (o *WalletTransaction) GetStatus() WalletTransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WalletTransaction) GetStatusOk() (*WalletTransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WalletTransaction) SetStatus(v WalletTransactionStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *WalletTransaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WalletTransaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WalletTransaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


