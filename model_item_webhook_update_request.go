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

// ItemWebhookUpdateRequest ItemWebhookUpdateRequest defines the request schema for `/item/webhook/update`
type ItemWebhookUpdateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// The new webhook URL to associate with the Item.
	Webhook NullableString `json:"webhook,omitempty"`
}

// NewItemWebhookUpdateRequest instantiates a new ItemWebhookUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemWebhookUpdateRequest(accessToken string) *ItemWebhookUpdateRequest {
	this := ItemWebhookUpdateRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewItemWebhookUpdateRequestWithDefaults instantiates a new ItemWebhookUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWebhookUpdateRequestWithDefaults() *ItemWebhookUpdateRequest {
	this := ItemWebhookUpdateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ItemWebhookUpdateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemWebhookUpdateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ItemWebhookUpdateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ItemWebhookUpdateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ItemWebhookUpdateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemWebhookUpdateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ItemWebhookUpdateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ItemWebhookUpdateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *ItemWebhookUpdateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ItemWebhookUpdateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ItemWebhookUpdateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemWebhookUpdateRequest) GetWebhook() string {
	if o == nil || o.Webhook.Get() == nil {
		var ret string
		return ret
	}
	return *o.Webhook.Get()
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemWebhookUpdateRequest) GetWebhookOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Webhook.Get(), o.Webhook.IsSet()
}

// HasWebhook returns a boolean if a field has been set.
func (o *ItemWebhookUpdateRequest) HasWebhook() bool {
	if o != nil && o.Webhook.IsSet() {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given NullableString and assigns it to the Webhook field.
func (o *ItemWebhookUpdateRequest) SetWebhook(v string) {
	o.Webhook.Set(&v)
}
// SetWebhookNil sets the value for Webhook to be an explicit nil
func (o *ItemWebhookUpdateRequest) SetWebhookNil() {
	o.Webhook.Set(nil)
}

// UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
func (o *ItemWebhookUpdateRequest) UnsetWebhook() {
	o.Webhook.Unset()
}

func (o ItemWebhookUpdateRequest) MarshalJSON() ([]byte, error) {
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
	if o.Webhook.IsSet() {
		toSerialize["webhook"] = o.Webhook.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableItemWebhookUpdateRequest struct {
	value *ItemWebhookUpdateRequest
	isSet bool
}

func (v NullableItemWebhookUpdateRequest) Get() *ItemWebhookUpdateRequest {
	return v.value
}

func (v *NullableItemWebhookUpdateRequest) Set(val *ItemWebhookUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableItemWebhookUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableItemWebhookUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemWebhookUpdateRequest(val *ItemWebhookUpdateRequest) *NullableItemWebhookUpdateRequest {
	return &NullableItemWebhookUpdateRequest{value: val, isSet: true}
}

func (v NullableItemWebhookUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemWebhookUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


