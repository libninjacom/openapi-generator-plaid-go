# IncomeVerificationPrecheckUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The user&#39;s first name | [optional] 
**LastName** | Pointer to **NullableString** | The user&#39;s last name | [optional] 
**EmailAddress** | Pointer to **NullableString** | The user&#39;s email address | [optional] 
**HomeAddress** | Pointer to [**NullableSignalAddressData**](SignalAddressData.md) |  | [optional] 

## Methods

### NewIncomeVerificationPrecheckUser

`func NewIncomeVerificationPrecheckUser() *IncomeVerificationPrecheckUser`

NewIncomeVerificationPrecheckUser instantiates a new IncomeVerificationPrecheckUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationPrecheckUserWithDefaults

`func NewIncomeVerificationPrecheckUserWithDefaults() *IncomeVerificationPrecheckUser`

NewIncomeVerificationPrecheckUserWithDefaults instantiates a new IncomeVerificationPrecheckUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *IncomeVerificationPrecheckUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IncomeVerificationPrecheckUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IncomeVerificationPrecheckUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *IncomeVerificationPrecheckUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *IncomeVerificationPrecheckUser) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *IncomeVerificationPrecheckUser) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *IncomeVerificationPrecheckUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IncomeVerificationPrecheckUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IncomeVerificationPrecheckUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *IncomeVerificationPrecheckUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *IncomeVerificationPrecheckUser) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *IncomeVerificationPrecheckUser) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetEmailAddress

`func (o *IncomeVerificationPrecheckUser) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *IncomeVerificationPrecheckUser) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *IncomeVerificationPrecheckUser) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *IncomeVerificationPrecheckUser) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *IncomeVerificationPrecheckUser) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *IncomeVerificationPrecheckUser) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetHomeAddress

`func (o *IncomeVerificationPrecheckUser) GetHomeAddress() SignalAddressData`

GetHomeAddress returns the HomeAddress field if non-nil, zero value otherwise.

### GetHomeAddressOk

`func (o *IncomeVerificationPrecheckUser) GetHomeAddressOk() (*SignalAddressData, bool)`

GetHomeAddressOk returns a tuple with the HomeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeAddress

`func (o *IncomeVerificationPrecheckUser) SetHomeAddress(v SignalAddressData)`

SetHomeAddress sets HomeAddress field to given value.

### HasHomeAddress

`func (o *IncomeVerificationPrecheckUser) HasHomeAddress() bool`

HasHomeAddress returns a boolean if a field has been set.

### SetHomeAddressNil

`func (o *IncomeVerificationPrecheckUser) SetHomeAddressNil(b bool)`

 SetHomeAddressNil sets the value for HomeAddress to be an explicit nil

### UnsetHomeAddress
`func (o *IncomeVerificationPrecheckUser) UnsetHomeAddress()`

UnsetHomeAddress ensures that no value is present for HomeAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


