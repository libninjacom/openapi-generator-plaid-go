# TransactionStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account to which the stream belongs | 
**StreamId** | **string** | A unique id for the stream | 
**CategoryId** | **string** | The ID of the category to which this transaction belongs. See [Categories](https://plaid.com/docs/#category-overview). | 
**Category** | **[]string** | A hierarchical array of the categories to which this transaction belongs. See [Categories](https://plaid.com/docs/#category-overview). | 
**Description** | **string** | A description of the transaction stream. | 
**FirstDate** | **string** | The posted date of the earliest transaction in the stream. | 
**LastDate** | **string** | The posted date of the latest transaction in the stream. | 
**Frequency** | [**RecurringTransactionFrequency**](RecurringTransactionFrequency.md) |  | 
**TransactionIds** | **[]string** | An array of Plaid transaction IDs belonging to the stream, sorted by posted date. | 
**AverageAmount** | [**TransactionStreamAmount**](TransactionStreamAmount.md) |  | 
**IsActive** | **bool** | indicates whether the transaction stream is still live. | 

## Methods

### NewTransactionStream

`func NewTransactionStream(accountId string, streamId string, categoryId string, category []string, description string, firstDate string, lastDate string, frequency RecurringTransactionFrequency, transactionIds []string, averageAmount TransactionStreamAmount, isActive bool, ) *TransactionStream`

NewTransactionStream instantiates a new TransactionStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionStreamWithDefaults

`func NewTransactionStreamWithDefaults() *TransactionStream`

NewTransactionStreamWithDefaults instantiates a new TransactionStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TransactionStream) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransactionStream) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransactionStream) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetStreamId

`func (o *TransactionStream) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *TransactionStream) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *TransactionStream) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.


### GetCategoryId

`func (o *TransactionStream) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *TransactionStream) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *TransactionStream) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetCategory

`func (o *TransactionStream) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TransactionStream) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TransactionStream) SetCategory(v []string)`

SetCategory sets Category field to given value.


### GetDescription

`func (o *TransactionStream) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionStream) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionStream) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFirstDate

`func (o *TransactionStream) GetFirstDate() string`

GetFirstDate returns the FirstDate field if non-nil, zero value otherwise.

### GetFirstDateOk

`func (o *TransactionStream) GetFirstDateOk() (*string, bool)`

GetFirstDateOk returns a tuple with the FirstDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDate

`func (o *TransactionStream) SetFirstDate(v string)`

SetFirstDate sets FirstDate field to given value.


### GetLastDate

`func (o *TransactionStream) GetLastDate() string`

GetLastDate returns the LastDate field if non-nil, zero value otherwise.

### GetLastDateOk

`func (o *TransactionStream) GetLastDateOk() (*string, bool)`

GetLastDateOk returns a tuple with the LastDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDate

`func (o *TransactionStream) SetLastDate(v string)`

SetLastDate sets LastDate field to given value.


### GetFrequency

`func (o *TransactionStream) GetFrequency() RecurringTransactionFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *TransactionStream) GetFrequencyOk() (*RecurringTransactionFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *TransactionStream) SetFrequency(v RecurringTransactionFrequency)`

SetFrequency sets Frequency field to given value.


### GetTransactionIds

`func (o *TransactionStream) GetTransactionIds() []string`

GetTransactionIds returns the TransactionIds field if non-nil, zero value otherwise.

### GetTransactionIdsOk

`func (o *TransactionStream) GetTransactionIdsOk() (*[]string, bool)`

GetTransactionIdsOk returns a tuple with the TransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIds

`func (o *TransactionStream) SetTransactionIds(v []string)`

SetTransactionIds sets TransactionIds field to given value.


### GetAverageAmount

`func (o *TransactionStream) GetAverageAmount() TransactionStreamAmount`

GetAverageAmount returns the AverageAmount field if non-nil, zero value otherwise.

### GetAverageAmountOk

`func (o *TransactionStream) GetAverageAmountOk() (*TransactionStreamAmount, bool)`

GetAverageAmountOk returns a tuple with the AverageAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageAmount

`func (o *TransactionStream) SetAverageAmount(v TransactionStreamAmount)`

SetAverageAmount sets AverageAmount field to given value.


### GetIsActive

`func (o *TransactionStream) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TransactionStream) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TransactionStream) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


