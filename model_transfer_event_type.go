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

// TransferEventType The type of event that this transfer represents.  `pending`: A new transfer was created; it is in the pending state.  `cancelled`: The transfer was cancelled by the client.  `failed`: The transfer failed, no funds were moved.  `posted`: The transfer has been successfully submitted to the payment network.  `reversed`: A posted transfer was reversed.  `swept`: The transfer was swept to / from the sweep account.  `reverse_swept`: Due to the transfer reversing, funds were pulled from or pushed back to the sweep account.
type TransferEventType string

// List of TransferEventType
const (
	PENDING TransferEventType = "pending"
	CANCELLED TransferEventType = "cancelled"
	FAILED TransferEventType = "failed"
	POSTED TransferEventType = "posted"
	REVERSED TransferEventType = "reversed"
	SWEPT TransferEventType = "swept"
	REVERSE_SWEPT TransferEventType = "reverse_swept"
)

// All allowed values of TransferEventType enum
var AllowedTransferEventTypeEnumValues = []TransferEventType{
	"pending",
	"cancelled",
	"failed",
	"posted",
	"reversed",
	"swept",
	"reverse_swept",
}

func (v *TransferEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferEventType(value)
	for _, existing := range AllowedTransferEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferEventType", value)
}

// NewTransferEventTypeFromValue returns a pointer to a valid TransferEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferEventTypeFromValue(v string) (*TransferEventType, error) {
	ev := TransferEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferEventType: valid values are %v", v, AllowedTransferEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferEventType) IsValid() bool {
	for _, existing := range AllowedTransferEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferEventType value
func (v TransferEventType) Ptr() *TransferEventType {
	return &v
}

type NullableTransferEventType struct {
	value *TransferEventType
	isSet bool
}

func (v NullableTransferEventType) Get() *TransferEventType {
	return v.value
}

func (v *NullableTransferEventType) Set(val *TransferEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEventType(val *TransferEventType) *NullableTransferEventType {
	return &NullableTransferEventType{value: val, isSet: true}
}

func (v NullableTransferEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
