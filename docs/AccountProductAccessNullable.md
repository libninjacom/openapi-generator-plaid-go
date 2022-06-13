# AccountProductAccessNullable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountData** | Pointer to **NullableBool** | Allow the application to access account data. Only used by certain partners. If relevant to the partner and unset, defaults to &#x60;true&#x60;. | [optional] [default to true]
**Statements** | Pointer to **NullableBool** | Allow the application to access bank statements. Only used by certain partners. If relevant to the partner and unset, defaults to &#x60;true&#x60;. | [optional] [default to true]
**TaxDocuments** | Pointer to **NullableBool** | Allow the application to access tax documents. Only used by certain partners. If relevant to the partner and unset, defaults to &#x60;true&#x60;. | [optional] [default to true]

## Methods

### NewAccountProductAccessNullable

`func NewAccountProductAccessNullable() *AccountProductAccessNullable`

NewAccountProductAccessNullable instantiates a new AccountProductAccessNullable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountProductAccessNullableWithDefaults

`func NewAccountProductAccessNullableWithDefaults() *AccountProductAccessNullable`

NewAccountProductAccessNullableWithDefaults instantiates a new AccountProductAccessNullable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountData

`func (o *AccountProductAccessNullable) GetAccountData() bool`

GetAccountData returns the AccountData field if non-nil, zero value otherwise.

### GetAccountDataOk

`func (o *AccountProductAccessNullable) GetAccountDataOk() (*bool, bool)`

GetAccountDataOk returns a tuple with the AccountData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountData

`func (o *AccountProductAccessNullable) SetAccountData(v bool)`

SetAccountData sets AccountData field to given value.

### HasAccountData

`func (o *AccountProductAccessNullable) HasAccountData() bool`

HasAccountData returns a boolean if a field has been set.

### SetAccountDataNil

`func (o *AccountProductAccessNullable) SetAccountDataNil(b bool)`

 SetAccountDataNil sets the value for AccountData to be an explicit nil

### UnsetAccountData
`func (o *AccountProductAccessNullable) UnsetAccountData()`

UnsetAccountData ensures that no value is present for AccountData, not even an explicit nil
### GetStatements

`func (o *AccountProductAccessNullable) GetStatements() bool`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AccountProductAccessNullable) GetStatementsOk() (*bool, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AccountProductAccessNullable) SetStatements(v bool)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *AccountProductAccessNullable) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### SetStatementsNil

`func (o *AccountProductAccessNullable) SetStatementsNil(b bool)`

 SetStatementsNil sets the value for Statements to be an explicit nil

### UnsetStatements
`func (o *AccountProductAccessNullable) UnsetStatements()`

UnsetStatements ensures that no value is present for Statements, not even an explicit nil
### GetTaxDocuments

`func (o *AccountProductAccessNullable) GetTaxDocuments() bool`

GetTaxDocuments returns the TaxDocuments field if non-nil, zero value otherwise.

### GetTaxDocumentsOk

`func (o *AccountProductAccessNullable) GetTaxDocumentsOk() (*bool, bool)`

GetTaxDocumentsOk returns a tuple with the TaxDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDocuments

`func (o *AccountProductAccessNullable) SetTaxDocuments(v bool)`

SetTaxDocuments sets TaxDocuments field to given value.

### HasTaxDocuments

`func (o *AccountProductAccessNullable) HasTaxDocuments() bool`

HasTaxDocuments returns a boolean if a field has been set.

### SetTaxDocumentsNil

`func (o *AccountProductAccessNullable) SetTaxDocumentsNil(b bool)`

 SetTaxDocumentsNil sets the value for TaxDocuments to be an explicit nil

### UnsetTaxDocuments
`func (o *AccountProductAccessNullable) UnsetTaxDocuments()`

UnsetTaxDocuments ensures that no value is present for TaxDocuments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


