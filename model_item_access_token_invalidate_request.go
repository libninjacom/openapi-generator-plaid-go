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

// ItemAccessTokenInvalidateRequest ItemAccessTokenInvalidateRequest defines the request schema for `/item/access_token/invalidate`
type ItemAccessTokenInvalidateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
}

// NewItemAccessTokenInvalidateRequest instantiates a new ItemAccessTokenInvalidateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemAccessTokenInvalidateRequest(accessToken string) *ItemAccessTokenInvalidateRequest {
	this := ItemAccessTokenInvalidateRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewItemAccessTokenInvalidateRequestWithDefaults instantiates a new ItemAccessTokenInvalidateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemAccessTokenInvalidateRequestWithDefaults() *ItemAccessTokenInvalidateRequest {
	this := ItemAccessTokenInvalidateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ItemAccessTokenInvalidateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemAccessTokenInvalidateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ItemAccessTokenInvalidateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ItemAccessTokenInvalidateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ItemAccessTokenInvalidateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemAccessTokenInvalidateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ItemAccessTokenInvalidateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ItemAccessTokenInvalidateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *ItemAccessTokenInvalidateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ItemAccessTokenInvalidateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ItemAccessTokenInvalidateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

func (o ItemAccessTokenInvalidateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	return json.Marshal(toSerialize)
}

type NullableItemAccessTokenInvalidateRequest struct {
	value *ItemAccessTokenInvalidateRequest
	isSet bool
}

func (v NullableItemAccessTokenInvalidateRequest) Get() *ItemAccessTokenInvalidateRequest {
	return v.value
}

func (v *NullableItemAccessTokenInvalidateRequest) Set(val *ItemAccessTokenInvalidateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableItemAccessTokenInvalidateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableItemAccessTokenInvalidateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemAccessTokenInvalidateRequest(val *ItemAccessTokenInvalidateRequest) *NullableItemAccessTokenInvalidateRequest {
	return &NullableItemAccessTokenInvalidateRequest{value: val, isSet: true}
}

func (v NullableItemAccessTokenInvalidateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemAccessTokenInvalidateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

