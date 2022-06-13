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

// PaymentInitiationPaymentCreateResponse PaymentInitiationPaymentCreateResponse defines the response schema for `/payment_initiation/payment/create`
type PaymentInitiationPaymentCreateResponse struct {
	// A unique ID identifying the payment
	PaymentId string `json:"payment_id"`
	// For a payment returned by this endpoint, there is only one possible value:  `PAYMENT_STATUS_INPUT_NEEDED`: The initial phase of the payment
	Status string `json:"status"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationPaymentCreateResponse PaymentInitiationPaymentCreateResponse

// NewPaymentInitiationPaymentCreateResponse instantiates a new PaymentInitiationPaymentCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationPaymentCreateResponse(paymentId string, status string, requestId string) *PaymentInitiationPaymentCreateResponse {
	this := PaymentInitiationPaymentCreateResponse{}
	this.PaymentId = paymentId
	this.Status = status
	this.RequestId = requestId
	return &this
}

// NewPaymentInitiationPaymentCreateResponseWithDefaults instantiates a new PaymentInitiationPaymentCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationPaymentCreateResponseWithDefaults() *PaymentInitiationPaymentCreateResponse {
	this := PaymentInitiationPaymentCreateResponse{}
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentInitiationPaymentCreateResponse) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentCreateResponse) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentInitiationPaymentCreateResponse) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetStatus returns the Status field value
func (o *PaymentInitiationPaymentCreateResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentCreateResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentInitiationPaymentCreateResponse) SetStatus(v string) {
	o.Status = v
}

// GetRequestId returns the RequestId field value
func (o *PaymentInitiationPaymentCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationPaymentCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentInitiationPaymentCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentInitiationPaymentCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_id"] = o.PaymentId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationPaymentCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationPaymentCreateResponse := _PaymentInitiationPaymentCreateResponse{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationPaymentCreateResponse); err == nil {
		*o = PaymentInitiationPaymentCreateResponse(varPaymentInitiationPaymentCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "payment_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationPaymentCreateResponse struct {
	value *PaymentInitiationPaymentCreateResponse
	isSet bool
}

func (v NullablePaymentInitiationPaymentCreateResponse) Get() *PaymentInitiationPaymentCreateResponse {
	return v.value
}

func (v *NullablePaymentInitiationPaymentCreateResponse) Set(val *PaymentInitiationPaymentCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentCreateResponse(val *PaymentInitiationPaymentCreateResponse) *NullablePaymentInitiationPaymentCreateResponse {
	return &NullablePaymentInitiationPaymentCreateResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


