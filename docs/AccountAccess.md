# AccountAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueId** | **string** | The unique account identifier for this account. This value must match that returned by the data access API for this account. | 
**Authorized** | Pointer to **NullableBool** | Allow the application to see this account (and associated details, including balance) in the list of accounts  If unset, defaults to &#x60;true&#x60;. | [optional] [default to true]
**AccountProductAccess** | Pointer to [**NullableAccountProductAccessNullable**](AccountProductAccessNullable.md) |  | [optional] 

## Methods

### NewAccountAccess

`func NewAccountAccess(uniqueId string, ) *AccountAccess`

NewAccountAccess instantiates a new AccountAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAccessWithDefaults

`func NewAccountAccessWithDefaults() *AccountAccess`

NewAccountAccessWithDefaults instantiates a new AccountAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueId

`func (o *AccountAccess) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *AccountAccess) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *AccountAccess) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.


### GetAuthorized

`func (o *AccountAccess) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *AccountAccess) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *AccountAccess) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *AccountAccess) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### SetAuthorizedNil

`func (o *AccountAccess) SetAuthorizedNil(b bool)`

 SetAuthorizedNil sets the value for Authorized to be an explicit nil

### UnsetAuthorized
`func (o *AccountAccess) UnsetAuthorized()`

UnsetAuthorized ensures that no value is present for Authorized, not even an explicit nil
### GetAccountProductAccess

`func (o *AccountAccess) GetAccountProductAccess() AccountProductAccessNullable`

GetAccountProductAccess returns the AccountProductAccess field if non-nil, zero value otherwise.

### GetAccountProductAccessOk

`func (o *AccountAccess) GetAccountProductAccessOk() (*AccountProductAccessNullable, bool)`

GetAccountProductAccessOk returns a tuple with the AccountProductAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountProductAccess

`func (o *AccountAccess) SetAccountProductAccess(v AccountProductAccessNullable)`

SetAccountProductAccess sets AccountProductAccess field to given value.

### HasAccountProductAccess

`func (o *AccountAccess) HasAccountProductAccess() bool`

HasAccountProductAccess returns a boolean if a field has been set.

### SetAccountProductAccessNil

`func (o *AccountAccess) SetAccountProductAccessNil(b bool)`

 SetAccountProductAccessNil sets the value for AccountProductAccess to be an explicit nil

### UnsetAccountProductAccess
`func (o *AccountAccess) UnsetAccountProductAccess()`

UnsetAccountProductAccess ensures that no value is present for AccountProductAccess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


