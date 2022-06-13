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

// InstitutionsSearchPaymentInitiationOptions Additional options that will be used to filter institutions by various Payment Initiation configurations.
type InstitutionsSearchPaymentInitiationOptions struct {
	// A unique ID identifying the payment
	PaymentId NullableString `json:"payment_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InstitutionsSearchPaymentInitiationOptions InstitutionsSearchPaymentInitiationOptions

// NewInstitutionsSearchPaymentInitiationOptions instantiates a new InstitutionsSearchPaymentInitiationOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsSearchPaymentInitiationOptions() *InstitutionsSearchPaymentInitiationOptions {
	this := InstitutionsSearchPaymentInitiationOptions{}
	return &this
}

// NewInstitutionsSearchPaymentInitiationOptionsWithDefaults instantiates a new InstitutionsSearchPaymentInitiationOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsSearchPaymentInitiationOptionsWithDefaults() *InstitutionsSearchPaymentInitiationOptions {
	this := InstitutionsSearchPaymentInitiationOptions{}
	return &this
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionsSearchPaymentInitiationOptions) GetPaymentId() string {
	if o == nil || o.PaymentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentId.Get()
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionsSearchPaymentInitiationOptions) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentId.Get(), o.PaymentId.IsSet()
}

// HasPaymentId returns a boolean if a field has been set.
func (o *InstitutionsSearchPaymentInitiationOptions) HasPaymentId() bool {
	if o != nil && o.PaymentId.IsSet() {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given NullableString and assigns it to the PaymentId field.
func (o *InstitutionsSearchPaymentInitiationOptions) SetPaymentId(v string) {
	o.PaymentId.Set(&v)
}
// SetPaymentIdNil sets the value for PaymentId to be an explicit nil
func (o *InstitutionsSearchPaymentInitiationOptions) SetPaymentIdNil() {
	o.PaymentId.Set(nil)
}

// UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
func (o *InstitutionsSearchPaymentInitiationOptions) UnsetPaymentId() {
	o.PaymentId.Unset()
}

func (o InstitutionsSearchPaymentInitiationOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentId.IsSet() {
		toSerialize["payment_id"] = o.PaymentId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InstitutionsSearchPaymentInitiationOptions) UnmarshalJSON(bytes []byte) (err error) {
	varInstitutionsSearchPaymentInitiationOptions := _InstitutionsSearchPaymentInitiationOptions{}

	if err = json.Unmarshal(bytes, &varInstitutionsSearchPaymentInitiationOptions); err == nil {
		*o = InstitutionsSearchPaymentInitiationOptions(varInstitutionsSearchPaymentInitiationOptions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "payment_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstitutionsSearchPaymentInitiationOptions struct {
	value *InstitutionsSearchPaymentInitiationOptions
	isSet bool
}

func (v NullableInstitutionsSearchPaymentInitiationOptions) Get() *InstitutionsSearchPaymentInitiationOptions {
	return v.value
}

func (v *NullableInstitutionsSearchPaymentInitiationOptions) Set(val *InstitutionsSearchPaymentInitiationOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsSearchPaymentInitiationOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsSearchPaymentInitiationOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsSearchPaymentInitiationOptions(val *InstitutionsSearchPaymentInitiationOptions) *NullableInstitutionsSearchPaymentInitiationOptions {
	return &NullableInstitutionsSearchPaymentInitiationOptions{value: val, isSet: true}
}

func (v NullableInstitutionsSearchPaymentInitiationOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsSearchPaymentInitiationOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


