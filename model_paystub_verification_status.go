/*
The Plaid API

The Plaid REST API. Please see https://plaid.com/docs/api for more details.

API version: 2020-09-14_1.64.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PaystubVerificationStatus Derived verification status.
type PaystubVerificationStatus string

// List of PaystubVerificationStatus
const (
	PAYSTUB_VERIFICATION_STATUS_UNKNOWN PaystubVerificationStatus = "PAYSTUB_VERIFICATION_STATUS_UNKNOWN"
	PAYSTUB_VERIFICATION_STATUS_VERIFIED PaystubVerificationStatus = "PAYSTUB_VERIFICATION_STATUS_VERIFIED"
	PAYSTUB_VERIFICATION_STATUS_FRAUDULENT PaystubVerificationStatus = "PAYSTUB_VERIFICATION_STATUS_FRAUDULENT"
	NULL PaystubVerificationStatus = "null"
)

// All allowed values of PaystubVerificationStatus enum
var AllowedPaystubVerificationStatusEnumValues = []PaystubVerificationStatus{
	"PAYSTUB_VERIFICATION_STATUS_UNKNOWN",
	"PAYSTUB_VERIFICATION_STATUS_VERIFIED",
	"PAYSTUB_VERIFICATION_STATUS_FRAUDULENT",
	"null",
}

func (v *PaystubVerificationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaystubVerificationStatus(value)
	for _, existing := range AllowedPaystubVerificationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaystubVerificationStatus", value)
}

// NewPaystubVerificationStatusFromValue returns a pointer to a valid PaystubVerificationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaystubVerificationStatusFromValue(v string) (*PaystubVerificationStatus, error) {
	ev := PaystubVerificationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaystubVerificationStatus: valid values are %v", v, AllowedPaystubVerificationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaystubVerificationStatus) IsValid() bool {
	for _, existing := range AllowedPaystubVerificationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaystubVerificationStatus value
func (v PaystubVerificationStatus) Ptr() *PaystubVerificationStatus {
	return &v
}

type NullablePaystubVerificationStatus struct {
	value *PaystubVerificationStatus
	isSet bool
}

func (v NullablePaystubVerificationStatus) Get() *PaystubVerificationStatus {
	return v.value
}

func (v *NullablePaystubVerificationStatus) Set(val *PaystubVerificationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystubVerificationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystubVerificationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystubVerificationStatus(val *PaystubVerificationStatus) *NullablePaystubVerificationStatus {
	return &NullablePaystubVerificationStatus{value: val, isSet: true}
}

func (v NullablePaystubVerificationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystubVerificationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

