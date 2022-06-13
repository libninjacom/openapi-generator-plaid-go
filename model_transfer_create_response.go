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

// TransferCreateResponse Defines the response schema for `/transfer/create`
type TransferCreateResponse struct {
	Transfer Transfer `json:"transfer"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferCreateResponse TransferCreateResponse

// NewTransferCreateResponse instantiates a new TransferCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferCreateResponse(transfer Transfer, requestId string) *TransferCreateResponse {
	this := TransferCreateResponse{}
	this.Transfer = transfer
	this.RequestId = requestId
	return &this
}

// NewTransferCreateResponseWithDefaults instantiates a new TransferCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferCreateResponseWithDefaults() *TransferCreateResponse {
	this := TransferCreateResponse{}
	return &this
}

// GetTransfer returns the Transfer field value
func (o *TransferCreateResponse) GetTransfer() Transfer {
	if o == nil {
		var ret Transfer
		return ret
	}

	return o.Transfer
}

// GetTransferOk returns a tuple with the Transfer field value
// and a boolean to check if the value has been set.
func (o *TransferCreateResponse) GetTransferOk() (*Transfer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transfer, true
}

// SetTransfer sets field value
func (o *TransferCreateResponse) SetTransfer(v Transfer) {
	o.Transfer = v
}

// GetRequestId returns the RequestId field value
func (o *TransferCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transfer"] = o.Transfer
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferCreateResponse := _TransferCreateResponse{}

	if err = json.Unmarshal(bytes, &varTransferCreateResponse); err == nil {
		*o = TransferCreateResponse(varTransferCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transfer")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferCreateResponse struct {
	value *TransferCreateResponse
	isSet bool
}

func (v NullableTransferCreateResponse) Get() *TransferCreateResponse {
	return v.value
}

func (v *NullableTransferCreateResponse) Set(val *TransferCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferCreateResponse(val *TransferCreateResponse) *NullableTransferCreateResponse {
	return &NullableTransferCreateResponse{value: val, isSet: true}
}

func (v NullableTransferCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


