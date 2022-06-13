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

// Employee Data about the employee.
type Employee struct {
	Address PaystubAddress `json:"address"`
	// The name of the employee.
	Name NullableString `json:"name"`
	// Marital status of the employee - either `single` or `married`.
	MaritalStatus NullableString `json:"marital_status,omitempty"`
	TaxpayerId *TaxpayerID `json:"taxpayer_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Employee Employee

// NewEmployee instantiates a new Employee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployee(address PaystubAddress, name NullableString) *Employee {
	this := Employee{}
	this.Address = address
	this.Name = name
	return &this
}

// NewEmployeeWithDefaults instantiates a new Employee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeWithDefaults() *Employee {
	this := Employee{}
	return &this
}

// GetAddress returns the Address field value
func (o *Employee) GetAddress() PaystubAddress {
	if o == nil {
		var ret PaystubAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Employee) GetAddressOk() (*PaystubAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *Employee) SetAddress(v PaystubAddress) {
	o.Address = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Employee) SetName(v string) {
	o.Name.Set(&v)
}

// GetMaritalStatus returns the MaritalStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Employee) GetMaritalStatus() string {
	if o == nil || o.MaritalStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaritalStatus.Get()
}

// GetMaritalStatusOk returns a tuple with the MaritalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee) GetMaritalStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaritalStatus.Get(), o.MaritalStatus.IsSet()
}

// HasMaritalStatus returns a boolean if a field has been set.
func (o *Employee) HasMaritalStatus() bool {
	if o != nil && o.MaritalStatus.IsSet() {
		return true
	}

	return false
}

// SetMaritalStatus gets a reference to the given NullableString and assigns it to the MaritalStatus field.
func (o *Employee) SetMaritalStatus(v string) {
	o.MaritalStatus.Set(&v)
}
// SetMaritalStatusNil sets the value for MaritalStatus to be an explicit nil
func (o *Employee) SetMaritalStatusNil() {
	o.MaritalStatus.Set(nil)
}

// UnsetMaritalStatus ensures that no value is present for MaritalStatus, not even an explicit nil
func (o *Employee) UnsetMaritalStatus() {
	o.MaritalStatus.Unset()
}

// GetTaxpayerId returns the TaxpayerId field value if set, zero value otherwise.
func (o *Employee) GetTaxpayerId() TaxpayerID {
	if o == nil || o.TaxpayerId == nil {
		var ret TaxpayerID
		return ret
	}
	return *o.TaxpayerId
}

// GetTaxpayerIdOk returns a tuple with the TaxpayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employee) GetTaxpayerIdOk() (*TaxpayerID, bool) {
	if o == nil || o.TaxpayerId == nil {
		return nil, false
	}
	return o.TaxpayerId, true
}

// HasTaxpayerId returns a boolean if a field has been set.
func (o *Employee) HasTaxpayerId() bool {
	if o != nil && o.TaxpayerId != nil {
		return true
	}

	return false
}

// SetTaxpayerId gets a reference to the given TaxpayerID and assigns it to the TaxpayerId field.
func (o *Employee) SetTaxpayerId(v TaxpayerID) {
	o.TaxpayerId = &v
}

func (o Employee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if o.MaritalStatus.IsSet() {
		toSerialize["marital_status"] = o.MaritalStatus.Get()
	}
	if o.TaxpayerId != nil {
		toSerialize["taxpayer_id"] = o.TaxpayerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Employee) UnmarshalJSON(bytes []byte) (err error) {
	varEmployee := _Employee{}

	if err = json.Unmarshal(bytes, &varEmployee); err == nil {
		*o = Employee(varEmployee)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "name")
		delete(additionalProperties, "marital_status")
		delete(additionalProperties, "taxpayer_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmployee struct {
	value *Employee
	isSet bool
}

func (v NullableEmployee) Get() *Employee {
	return v.value
}

func (v *NullableEmployee) Set(val *Employee) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployee) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployee(val *Employee) *NullableEmployee {
	return &NullableEmployee{value: val, isSet: true}
}

func (v NullableEmployee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


