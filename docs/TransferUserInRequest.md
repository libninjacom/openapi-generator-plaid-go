# TransferUserInRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegalName** | **string** | The user&#39;s legal name. | 
**PhoneNumber** | Pointer to **string** | The user&#39;s phone number. | [optional] 
**EmailAddress** | Pointer to **string** | The user&#39;s email address. | [optional] 
**Address** | Pointer to [**TransferUserAddressInRequest**](TransferUserAddressInRequest.md) |  | [optional] 

## Methods

### NewTransferUserInRequest

`func NewTransferUserInRequest(legalName string, ) *TransferUserInRequest`

NewTransferUserInRequest instantiates a new TransferUserInRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferUserInRequestWithDefaults

`func NewTransferUserInRequestWithDefaults() *TransferUserInRequest`

NewTransferUserInRequestWithDefaults instantiates a new TransferUserInRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegalName

`func (o *TransferUserInRequest) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *TransferUserInRequest) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *TransferUserInRequest) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetPhoneNumber

`func (o *TransferUserInRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *TransferUserInRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *TransferUserInRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *TransferUserInRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmailAddress

`func (o *TransferUserInRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *TransferUserInRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *TransferUserInRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *TransferUserInRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetAddress

`func (o *TransferUserInRequest) GetAddress() TransferUserAddressInRequest`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TransferUserInRequest) GetAddressOk() (*TransferUserAddressInRequest, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TransferUserInRequest) SetAddress(v TransferUserAddressInRequest)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TransferUserInRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


