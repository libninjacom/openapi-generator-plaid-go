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

// TransferRepaymentListResponse Defines the response schema for `/transfer/repayments/list`
type TransferRepaymentListResponse struct {
	Repayments []TransferRepayment `json:"repayments"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferRepaymentListResponse TransferRepaymentListResponse

// NewTransferRepaymentListResponse instantiates a new TransferRepaymentListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRepaymentListResponse(repayments []TransferRepayment, requestId string) *TransferRepaymentListResponse {
	this := TransferRepaymentListResponse{}
	this.Repayments = repayments
	this.RequestId = requestId
	return &this
}

// NewTransferRepaymentListResponseWithDefaults instantiates a new TransferRepaymentListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRepaymentListResponseWithDefaults() *TransferRepaymentListResponse {
	this := TransferRepaymentListResponse{}
	return &this
}

// GetRepayments returns the Repayments field value
func (o *TransferRepaymentListResponse) GetRepayments() []TransferRepayment {
	if o == nil {
		var ret []TransferRepayment
		return ret
	}

	return o.Repayments
}

// GetRepaymentsOk returns a tuple with the Repayments field value
// and a boolean to check if the value has been set.
func (o *TransferRepaymentListResponse) GetRepaymentsOk() ([]TransferRepayment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repayments, true
}

// SetRepayments sets field value
func (o *TransferRepaymentListResponse) SetRepayments(v []TransferRepayment) {
	o.Repayments = v
}

// GetRequestId returns the RequestId field value
func (o *TransferRepaymentListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferRepaymentListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferRepaymentListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferRepaymentListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["repayments"] = o.Repayments
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferRepaymentListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferRepaymentListResponse := _TransferRepaymentListResponse{}

	if err = json.Unmarshal(bytes, &varTransferRepaymentListResponse); err == nil {
		*o = TransferRepaymentListResponse(varTransferRepaymentListResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "repayments")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferRepaymentListResponse struct {
	value *TransferRepaymentListResponse
	isSet bool
}

func (v NullableTransferRepaymentListResponse) Get() *TransferRepaymentListResponse {
	return v.value
}

func (v *NullableTransferRepaymentListResponse) Set(val *TransferRepaymentListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRepaymentListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRepaymentListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRepaymentListResponse(val *TransferRepaymentListResponse) *NullableTransferRepaymentListResponse {
	return &NullableTransferRepaymentListResponse{value: val, isSet: true}
}

func (v NullableTransferRepaymentListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRepaymentListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


