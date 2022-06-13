# TransferUserInResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegalName** | **string** | The user&#39;s legal name. | 
**PhoneNumber** | **NullableString** | The user&#39;s phone number. | 
**EmailAddress** | **NullableString** | The user&#39;s email address. | 
**Address** | [**NullableTransferUserAddressInResponse**](TransferUserAddressInResponse.md) |  | 

## Methods

### NewTransferUserInResponse

`func NewTransferUserInResponse(legalName string, phoneNumber NullableString, emailAddress NullableString, address NullableTransferUserAddressInResponse, ) *TransferUserInResponse`

NewTransferUserInResponse instantiates a new TransferUserInResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferUserInResponseWithDefaults

`func NewTransferUserInResponseWithDefaults() *TransferUserInResponse`

NewTransferUserInResponseWithDefaults instantiates a new TransferUserInResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegalName

`func (o *TransferUserInResponse) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *TransferUserInResponse) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *TransferUserInResponse) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetPhoneNumber

`func (o *TransferUserInResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *TransferUserInResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *TransferUserInResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### SetPhoneNumberNil

`func (o *TransferUserInResponse) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *TransferUserInResponse) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetEmailAddress

`func (o *TransferUserInResponse) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *TransferUserInResponse) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *TransferUserInResponse) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### SetEmailAddressNil

`func (o *TransferUserInResponse) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *TransferUserInResponse) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetAddress

`func (o *TransferUserInResponse) GetAddress() TransferUserAddressInResponse`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransferUserInResponse) GetAddressOk() (*TransferUserAddressInResponse, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransferUserInResponse) SetAddress(v TransferUserAddressInResponse)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *TransferUserInResponse) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *TransferUserInResponse) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


