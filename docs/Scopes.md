# Scopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductAccess** | Pointer to [**ProductAccess**](ProductAccess.md) |  | [optional] 
**Accounts** | Pointer to [**[]AccountAccess**](AccountAccess.md) |  | [optional] 
**NewAccounts** | Pointer to **NullableBool** | Allow access to newly opened accounts as they are opened. If unset, defaults to &#x60;true&#x60;. | [optional] [default to true]

## Methods

### NewScopes

`func NewScopes() *Scopes`

NewScopes instantiates a new Scopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopesWithDefaults

`func NewScopesWithDefaults() *Scopes`

NewScopesWithDefaults instantiates a new Scopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductAccess

`func (o *Scopes) GetProductAccess() ProductAccess`

GetProductAccess returns the ProductAccess field if non-nil, zero value otherwise.

### GetProductAccessOk

`func (o *Scopes) GetProductAccessOk() (*ProductAccess, bool)`

GetProductAccessOk returns a tuple with the ProductAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAccess

`func (o *Scopes) SetProductAccess(v ProductAccess)`

SetProductAccess sets ProductAccess field to given value.

### HasProductAccess

`func (o *Scopes) HasProductAccess() bool`

HasProductAccess returns a boolean if a field has been set.

### GetAccounts

`func (o *Scopes) GetAccounts() []AccountAccess`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Scopes) GetAccountsOk() (*[]AccountAccess, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Scopes) SetAccounts(v []AccountAccess)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Scopes) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetNewAccounts

`func (o *Scopes) GetNewAccounts() bool`

GetNewAccounts returns the NewAccounts field if non-nil, zero value otherwise.

### GetNewAccountsOk

`func (o *Scopes) GetNewAccountsOk() (*bool, bool)`

GetNewAccountsOk returns a tuple with the NewAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAccounts

`func (o *Scopes) SetNewAccounts(v bool)`

SetNewAccounts sets NewAccounts field to given value.

### HasNewAccounts

`func (o *Scopes) HasNewAccounts() bool`

HasNewAccounts returns a boolean if a field has been set.

### SetNewAccountsNil

`func (o *Scopes) SetNewAccountsNil(b bool)`

 SetNewAccountsNil sets the value for NewAccounts to be an explicit nil

### UnsetNewAccounts
`func (o *Scopes) UnsetNewAccounts()`

UnsetNewAccounts ensures that no value is present for NewAccounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


