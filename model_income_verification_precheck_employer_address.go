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

// IncomeVerificationPrecheckEmployerAddress The address of the employer
type IncomeVerificationPrecheckEmployerAddress struct {
	// The full city name
	City *string `json:"city,omitempty"`
	// The region or state. In API versions 2018-05-22 and earlier, this field is called `state`. Example: `\"NC\"`
	Region *string `json:"region,omitempty"`
	// The full street address Example: `\"564 Main Street, APT 15\"`
	Street *string `json:"street,omitempty"`
	// The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.
	PostalCode *string `json:"postal_code,omitempty"`
	// The ISO 3166-1 alpha-2 country code
	Country *string `json:"country,omitempty"`
}

// NewIncomeVerificationPrecheckEmployerAddress instantiates a new IncomeVerificationPrecheckEmployerAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationPrecheckEmployerAddress() *IncomeVerificationPrecheckEmployerAddress {
	this := IncomeVerificationPrecheckEmployerAddress{}
	return &this
}

// NewIncomeVerificationPrecheckEmployerAddressWithDefaults instantiates a new IncomeVerificationPrecheckEmployerAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationPrecheckEmployerAddressWithDefaults() *IncomeVerificationPrecheckEmployerAddress {
	this := IncomeVerificationPrecheckEmployerAddress{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckEmployerAddress) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckEmployerAddress) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployerAddress) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *IncomeVerificationPrecheckEmployerAddress) SetCity(v string) {
	o.City = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckEmployerAddress) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckEmployerAddress) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployerAddress) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *IncomeVerificationPrecheckEmployerAddress) SetRegion(v string) {
	o.Region = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckEmployerAddress) GetStreet() string {
	if o == nil || o.Street == nil {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckEmployerAddress) GetStreetOk() (*string, bool) {
	if o == nil || o.Street == nil {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployerAddress) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *IncomeVerificationPrecheckEmployerAddress) SetStreet(v string) {
	o.Street = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckEmployerAddress) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckEmployerAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployerAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *IncomeVerificationPrecheckEmployerAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *IncomeVerificationPrecheckEmployerAddress) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationPrecheckEmployerAddress) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckEmployerAddress) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *IncomeVerificationPrecheckEmployerAddress) SetCountry(v string) {
	o.Country = &v
}

func (o IncomeVerificationPrecheckEmployerAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Street != nil {
		toSerialize["street"] = o.Street
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationPrecheckEmployerAddress struct {
	value *IncomeVerificationPrecheckEmployerAddress
	isSet bool
}

func (v NullableIncomeVerificationPrecheckEmployerAddress) Get() *IncomeVerificationPrecheckEmployerAddress {
	return v.value
}

func (v *NullableIncomeVerificationPrecheckEmployerAddress) Set(val *IncomeVerificationPrecheckEmployerAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPrecheckEmployerAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPrecheckEmployerAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPrecheckEmployerAddress(val *IncomeVerificationPrecheckEmployerAddress) *NullableIncomeVerificationPrecheckEmployerAddress {
	return &NullableIncomeVerificationPrecheckEmployerAddress{value: val, isSet: true}
}

func (v NullableIncomeVerificationPrecheckEmployerAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPrecheckEmployerAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


