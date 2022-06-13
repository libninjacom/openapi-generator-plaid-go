# ProductAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statements** | Pointer to **NullableBool** | Allow access to statements. Only used by certain partners. If relevant to the partner and unset, defaults to &#x60;true&#x60;. | [optional] [default to true]
**Identity** | Pointer to **NullableBool** | Allow access to the Identity product (name, email, phone, address). Only used by certain partners. If relevant to the partner and unset, defaults to &#x60;true&#x60;. | [optional] [default to true]
**Auth** | Pointer to **NullableBool** | Allow access to account number details. Only used by certain partners. If relevant to the partner and unset, defaults to &#x60;true&#x60;. | [optional] [default to true]
**Transactions** | Pointer to **NullableBool** | Allow access to transaction details. Only used by certain partners. If relevant to the partner and unset, defaults to &#x60;true&#x60;. | [optional] [default to true]

## Methods

### NewProductAccess

`func NewProductAccess() *ProductAccess`

NewProductAccess instantiates a new ProductAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAccessWithDefaults

`func NewProductAccessWithDefaults() *ProductAccess`

NewProductAccessWithDefaults instantiates a new ProductAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatements

`func (o *ProductAccess) GetStatements() bool`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *ProductAccess) GetStatementsOk() (*bool, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *ProductAccess) SetStatements(v bool)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *ProductAccess) HasStatements() bool`

HasStatements returns a boolean if a field has been set.

### SetStatementsNil

`func (o *ProductAccess) SetStatementsNil(b bool)`

 SetStatementsNil sets the value for Statements to be an explicit nil

### UnsetStatements
`func (o *ProductAccess) UnsetStatements()`

UnsetStatements ensures that no value is present for Statements, not even an explicit nil
### GetIdentity

`func (o *ProductAccess) GetIdentity() bool`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ProductAccess) GetIdentityOk() (*bool, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ProductAccess) SetIdentity(v bool)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ProductAccess) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *ProductAccess) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *ProductAccess) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetAuth

`func (o *ProductAccess) GetAuth() bool`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ProductAccess) GetAuthOk() (*bool, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ProductAccess) SetAuth(v bool)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ProductAccess) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### SetAuthNil

`func (o *ProductAccess) SetAuthNil(b bool)`

 SetAuthNil sets the value for Auth to be an explicit nil

### UnsetAuth
`func (o *ProductAccess) UnsetAuth()`

UnsetAuth ensures that no value is present for Auth, not even an explicit nil
### GetTransactions

`func (o *ProductAccess) GetTransactions() bool`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *ProductAccess) GetTransactionsOk() (*bool, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *ProductAccess) SetTransactions(v bool)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *ProductAccess) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### SetTransactionsNil

`func (o *ProductAccess) SetTransactionsNil(b bool)`

 SetTransactionsNil sets the value for Transactions to be an explicit nil

### UnsetTransactions
`func (o *ProductAccess) UnsetTransactions()`

UnsetTransactions ensures that no value is present for Transactions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


