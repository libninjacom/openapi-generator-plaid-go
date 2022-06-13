/*
The Plaid API

The Plaid REST API. Please see https://plaid.com/docs/api for more details.

API version: 2020-09-14_1.64.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TransferUserInResponse The legal name and other information for the account holder.
type TransferUserInResponse struct {
	// The user's legal name.
	LegalName string `json:"legal_name"`
	// The user's phone number.
	PhoneNumber NullableString `json:"phone_number"`
	// The user's email address.
	EmailAddress NullableString `json:"email_address"`
	Address NullableTransferUserAddressInResponse `json:"address"`
	AdditionalProperties map[string]interface{}
}

type _TransferUserInResponse TransferUserInResponse

// NewTransferUserInResponse instantiates a new TransferUserInResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferUserInResponse(legalName string, phoneNumber NullableString, emailAddress NullableString, address NullableTransferUserAddressInResponse) *TransferUserInResponse {
	this := TransferUserInResponse{}
	this.LegalName = legalName
	this.PhoneNumber = phoneNumber
	this.EmailAddress = emailAddress
	this.Address = address
	return &this
}

// NewTransferUserInResponseWithDefaults instantiates a new TransferUserInResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferUserInResponseWithDefaults() *TransferUserInResponse {
	this := TransferUserInResponse{}
	return &this
}

// GetLegalName returns the LegalName field value
func (o *TransferUserInResponse) GetLegalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value
// and a boolean to check if the value has been set.
func (o *TransferUserInResponse) GetLegalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegalName, true
}

// SetLegalName sets field value
func (o *TransferUserInResponse) SetLegalName(v string) {
	o.LegalName = v
}

// GetPhoneNumber returns the PhoneNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferUserInResponse) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferUserInResponse) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// SetPhoneNumber sets field value
func (o *TransferUserInResponse) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}

// GetEmailAddress returns the EmailAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferUserInResponse) GetEmailAddress() string {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferUserInResponse) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// SetEmailAddress sets field value
func (o *TransferUserInResponse) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for TransferUserAddressInResponse will be returned
func (o *TransferUserInResponse) GetAddress() TransferUserAddressInResponse {
	if o == nil || o.Address.Get() == nil {
		var ret TransferUserAddressInResponse
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferUserInResponse) GetAddressOk() (*TransferUserAddressInResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *TransferUserInResponse) SetAddress(v TransferUserAddressInResponse) {
	o.Address.Set(&v)
}

func (o TransferUserInResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["legal_name"] = o.LegalName
	}
	if true {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if true {
		toSerialize["email_address"] = o.EmailAddress.Get()
	}
	if true {
		toSerialize["address"] = o.Address.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferUserInResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferUserInResponse := _TransferUserInResponse{}

	if err = json.Unmarshal(bytes, &varTransferUserInResponse); err == nil {
		*o = TransferUserInResponse(varTransferUserInResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "legal_name")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "email_address")
		delete(additionalProperties, "address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferUserInResponse struct {
	value *TransferUserInResponse
	isSet bool
}

func (v NullableTransferUserInResponse) Get() *TransferUserInResponse {
	return v.value
}

func (v *NullableTransferUserInResponse) Set(val *TransferUserInResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferUserInResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferUserInResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferUserInResponse(val *TransferUserInResponse) *NullableTransferUserInResponse {
	return &NullableTransferUserInResponse{value: val, isSet: true}
}

func (v NullableTransferUserInResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferUserInResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

