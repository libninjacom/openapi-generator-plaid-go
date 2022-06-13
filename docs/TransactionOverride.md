# TransactionOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateTransacted** | **string** | The date of the transaction, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format. Transactions in Sandbox will move from pending to posted once their transaction date has been reached. If a &#x60;date_transacted&#x60; is not provided by the institution, a transaction date may be available in the [&#x60;authorized_date&#x60;](https://plaid.com/docs/api/products/#transactions-get-response-transactions-authorized-date) field. | 
**DatePosted** | **string** | The date the transaction posted, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format. Posted dates in the past or present will result in posted transactions; posted dates in the future will result in pending transactions. | 
**Amount** | **float32** | The transaction amount. Can be negative. | 
**Description** | **string** | The transaction description. | 
**Currency** | Pointer to **string** | The ISO-4217 format currency code for the transaction. | [optional] 

## Methods

### NewTransactionOverride

`func NewTransactionOverride(dateTransacted string, datePosted string, amount float32, description string, ) *TransactionOverride`

NewTransactionOverride instantiates a new TransactionOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionOverrideWithDefaults

`func NewTransactionOverrideWithDefaults() *TransactionOverride`

NewTransactionOverrideWithDefaults instantiates a new TransactionOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateTransacted

`func (o *TransactionOverride) GetDateTransacted() string`

GetDateTransacted returns the DateTransacted field if non-nil, zero value otherwise.

### GetDateTransactedOk

`func (o *TransactionOverride) GetDateTransactedOk() (*string, bool)`

GetDateTransactedOk returns a tuple with the DateTransacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTransacted

`func (o *TransactionOverride) SetDateTransacted(v string)`

SetDateTransacted sets DateTransacted field to given value.


### GetDatePosted

`func (o *TransactionOverride) GetDatePosted() string`

GetDatePosted returns the DatePosted field if non-nil, zero value otherwise.

### GetDatePostedOk

`func (o *TransactionOverride) GetDatePostedOk() (*string, bool)`

GetDatePostedOk returns a tuple with the DatePosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePosted

`func (o *TransactionOverride) SetDatePosted(v string)`

SetDatePosted sets DatePosted field to given value.


### GetAmount

`func (o *TransactionOverride) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionOverride) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionOverride) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *TransactionOverride) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionOverride) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionOverride) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCurrency

`func (o *TransactionOverride) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionOverride) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionOverride) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransactionOverride) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


