/*
The Plaid API

The Plaid REST API. Please see https://plaid.com/docs/api for more details.

API version: 2020-09-14_1.64.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// PaymentInitiationPaymentListRequest PaymentInitiationPaymentListRequest defines the request schema for `/payment_initiation/payment/list`
type PaymentInitiationPaymentListRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The maximum number of payments to return. If `count` is not specified, a maximum of 10 payments will be returned, beginning with the most recent payment before the cursor (if specified).
	Count NullableInt32 `json:"count,omitempty"`
	// A string in RFC 3339 format (i.e. \"2019-12-06T22:35:49Z\"). Only payments created before the cursor will be returned.
	Cursor NullableTime `json:"cursor,omitempty"`
}

// NewPaymentInitiationPaymentListRequest instantiates a new PaymentInitiationPaymentListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationPaymentListRequest() *PaymentInitiationPaymentListRequest {
	this := PaymentInitiationPaymentListRequest{}
	var count int32 = 10
	this.Count = *NewNullableInt32(&count)
	return &this
}

// NewPaymentInitiationPaymentListRequestWithDefaults instantiates a new PaymentInitiationPaymentListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationPaymentListRequestWithDefaults() *PaymentInitiationPaymentListRequest {
	this := PaymentInitiationPaymentListRequest{}
	var count int32 = 10
	this.Count = *NewNullableInt32(&count)
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentInitiationPaymentListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentInitiationPaymentListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentInitiationPaymentListRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentListRequest) GetCount() int32 {
	if o == nil || o.Count.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentListRequest) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentListRequest) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt32 and assigns it to the Count field.
func (o *PaymentInitiationPaymentListRequest) SetCount(v int32) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *PaymentInitiationPaymentListRequest) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *PaymentInitiationPaymentListRequest) UnsetCount() {
	o.Count.Unset()
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationPaymentListRequest) GetCursor() time.Time {
	if o == nil || o.Cursor.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationPaymentListRequest) GetCursorOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *PaymentInitiationPaymentListRequest) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableTime and assigns it to the Cursor field.
func (o *PaymentInitiationPaymentListRequest) SetCursor(v time.Time) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *PaymentInitiationPaymentListRequest) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *PaymentInitiationPaymentListRequest) UnsetCursor() {
	o.Cursor.Unset()
}

func (o PaymentInitiationPaymentListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	if o.Cursor.IsSet() {
		toSerialize["cursor"] = o.Cursor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInitiationPaymentListRequest struct {
	value *PaymentInitiationPaymentListRequest
	isSet bool
}

func (v NullablePaymentInitiationPaymentListRequest) Get() *PaymentInitiationPaymentListRequest {
	return v.value
}

func (v *NullablePaymentInitiationPaymentListRequest) Set(val *PaymentInitiationPaymentListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentListRequest(val *PaymentInitiationPaymentListRequest) *NullablePaymentInitiationPaymentListRequest {
	return &NullablePaymentInitiationPaymentListRequest{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


