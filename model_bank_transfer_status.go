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

// BankTransferStatus The status of the transfer.
type BankTransferStatus string

// List of BankTransferStatus
const (
	PENDING BankTransferStatus = "pending"
	POSTED BankTransferStatus = "posted"
	CANCELLED BankTransferStatus = "cancelled"
	FAILED BankTransferStatus = "failed"
	REVERSED BankTransferStatus = "reversed"
)

// All allowed values of BankTransferStatus enum
var AllowedBankTransferStatusEnumValues = []BankTransferStatus{
	"pending",
	"posted",
	"cancelled",
	"failed",
	"reversed",
}

func (v *BankTransferStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BankTransferStatus(value)
	for _, existing := range AllowedBankTransferStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BankTransferStatus", value)
}

// NewBankTransferStatusFromValue returns a pointer to a valid BankTransferStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBankTransferStatusFromValue(v string) (*BankTransferStatus, error) {
	ev := BankTransferStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BankTransferStatus: valid values are %v", v, AllowedBankTransferStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankTransferStatus) IsValid() bool {
	for _, existing := range AllowedBankTransferStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BankTransferStatus value
func (v BankTransferStatus) Ptr() *BankTransferStatus {
	return &v
}

type NullableBankTransferStatus struct {
	value *BankTransferStatus
	isSet bool
}

func (v NullableBankTransferStatus) Get() *BankTransferStatus {
	return v.value
}

func (v *NullableBankTransferStatus) Set(val *BankTransferStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferStatus(val *BankTransferStatus) *NullableBankTransferStatus {
	return &NullableBankTransferStatus{value: val, isSet: true}
}

func (v NullableBankTransferStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

