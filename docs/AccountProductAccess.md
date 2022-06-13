# AccountProductAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountData** | Pointer to **NullableBool** | Allow the application to access account data. Only used by certain partners. If relevant to the partner and unset, defaults to &#x60;true&#x60;. | [optional] [default to true]
**Statements** | Pointer to **NullableBool** | Allow the application to access bank statements. Only used by certain partners. If relevant to the partner and unset, defaults to &#x60;true&#x60;. | [optional] [default to true]
**TaxDocuments** | Pointer to **NullableBool** | Allow the application to access tax documents. Only used by certain partners. If relevant to the partner and unset, defaults to &#x60;true&#x60;. | [optional] [default to true]

## Methods

### NewAccountProductAccess

`func NewAccountProductAccess() *AccountProductAccess`

NewAccountProductAccess instantiates a new AccountProductAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountProductAccessWithDefaults

`func NewAccountProductAccessWithDefaults() *AccountProductAccess`

NewAccountProductAccessWithDefaults instantiates a new AccountProductAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountData

`func (o *AccountProductAccess) GetAccountData() bool`

GetAccountData returns the AccountData field if non-nil, zero value otherwise.

### GetAccountDataOk

`func (o *AccountProductAccess) GetAccountDataOk() (*bool, bool)`

GetAccountDataOk returns a tuple with the AccountData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountData

`func (o *AccountProductAccess) SetAccountData(v bool)`

SetAccountData sets AccountData field to given value.

### HasAccountData

`func (o *AccountProductAccess) HasAccountData() bool`

HasAccountData returns a boolean if a field has been set.

### SetAccountDataNil

`func (o *AccountProductAccess) SetAccountDataNil(b bool)`

 SetAccountDataNil sets the value for AccountData to be an explicit nil

### UnsetAccountData
`func (o *AccountProductAccess) UnsetAccountData()`

UnsetAccountData ensures that no value is present for AccountData, not even an explicit nil
### GetStatements

`func (o *AccountProductAccess) GetStatements() bool`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AccountProductAccess) GetStatementsOk() (*bool, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AccountProductAccess) SetStatements(v bool)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *AccountProductAccess) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### SetStatementsNil

`func (o *AccountProductAccess) SetStatementsNil(b bool)`

 SetStatementsNil sets the value for Statements to be an explicit nil

### UnsetStatements
`func (o *AccountProductAccess) UnsetStatements()`

UnsetStatements ensures that no value is present for Statements, not even an explicit nil
### GetTaxDocuments

`func (o *AccountProductAccess) GetTaxDocuments() bool`

GetTaxDocuments returns the TaxDocuments field if non-nil, zero value otherwise.

### GetTaxDocumentsOk

`func (o *AccountProductAccess) GetTaxDocumentsOk() (*bool, bool)`

GetTaxDocumentsOk returns a tuple with the TaxDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxDocuments

`func (o *AccountProductAccess) SetTaxDocuments(v bool)`

SetTaxDocuments sets TaxDocuments field to given value.

### HasTaxDocuments

`func (o *AccountProductAccess) HasTaxDocuments() bool`

HasTaxDocuments returns a boolean if a field has been set.

### SetTaxDocumentsNil

`func (o *AccountProductAccess) SetTaxDocumentsNil(b bool)`

 SetTaxDocumentsNil sets the value for TaxDocuments to be an explicit nil

### UnsetTaxDocuments
`func (o *AccountProductAccess) UnsetTaxDocuments()`

UnsetTaxDocuments ensures that no value is present for TaxDocuments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


