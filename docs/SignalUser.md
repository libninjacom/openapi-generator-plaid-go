# SignalUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**NullableSignalPersonName**](SignalPersonName.md) |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The user&#39;s phone number, in E.164 format: +{countrycode}{number}. For example: \&quot;+14151234567\&quot; | [optional] 
**EmailAddress** | Pointer to **NullableString** | The user&#39;s email address. | [optional] 
**Address** | Pointer to [**NullableSignalAddressData**](SignalAddressData.md) |  | [optional] 

## Methods

### NewSignalUser

`func NewSignalUser() *SignalUser`

NewSignalUser instantiates a new SignalUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalUserWithDefaults

`func NewSignalUserWithDefaults() *SignalUser`

NewSignalUserWithDefaults instantiates a new SignalUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SignalUser) GetName() SignalPersonName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignalUser) GetNameOk() (*SignalPersonName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignalUser) SetName(v SignalPersonName)`

SetName sets Name field to given value.

### HasName

`func (o *SignalUser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SignalUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SignalUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhoneNumber

`func (o *SignalUser) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *SignalUser) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *SignalUser) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *SignalUser) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *SignalUser) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *SignalUser) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetEmailAddress

`func (o *SignalUser) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SignalUser) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SignalUser) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *SignalUser) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *SignalUser) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *SignalUser) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetAddress

`func (o *SignalUser) GetAddress() SignalAddressData`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SignalUser) GetAddressOk() (*SignalAddressData, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SignalUser) SetAddress(v SignalAddressData)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SignalUser) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *SignalUser) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *SignalUser) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


