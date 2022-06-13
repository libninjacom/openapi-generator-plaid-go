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

// ProcessorBalanceGetResponse ProcessorBalanceGetResponse defines the response schema for `/processor/balance/get`
type ProcessorBalanceGetResponse struct {
	Account AccountBase `json:"account"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorBalanceGetResponse ProcessorBalanceGetResponse

// NewProcessorBalanceGetResponse instantiates a new ProcessorBalanceGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorBalanceGetResponse(account AccountBase, requestId string) *ProcessorBalanceGetResponse {
	this := ProcessorBalanceGetResponse{}
	this.Account = account
	this.RequestId = requestId
	return &this
}

// NewProcessorBalanceGetResponseWithDefaults instantiates a new ProcessorBalanceGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorBalanceGetResponseWithDefaults() *ProcessorBalanceGetResponse {
	this := ProcessorBalanceGetResponse{}
	return &this
}

// GetAccount returns the Account field value
func (o *ProcessorBalanceGetResponse) GetAccount() AccountBase {
	if o == nil {
		var ret AccountBase
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *ProcessorBalanceGetResponse) GetAccountOk() (*AccountBase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *ProcessorBalanceGetResponse) SetAccount(v AccountBase) {
	o.Account = v
}

// GetRequestId returns the RequestId field value
func (o *ProcessorBalanceGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ProcessorBalanceGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ProcessorBalanceGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ProcessorBalanceGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorBalanceGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorBalanceGetResponse := _ProcessorBalanceGetResponse{}

	if err = json.Unmarshal(bytes, &varProcessorBalanceGetResponse); err == nil {
		*o = ProcessorBalanceGetResponse(varProcessorBalanceGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessorBalanceGetResponse struct {
	value *ProcessorBalanceGetResponse
	isSet bool
}

func (v NullableProcessorBalanceGetResponse) Get() *ProcessorBalanceGetResponse {
	return v.value
}

func (v *NullableProcessorBalanceGetResponse) Set(val *ProcessorBalanceGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorBalanceGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorBalanceGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorBalanceGetResponse(val *ProcessorBalanceGetResponse) *NullableProcessorBalanceGetResponse {
	return &NullableProcessorBalanceGetResponse{value: val, isSet: true}
}

func (v NullableProcessorBalanceGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorBalanceGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

