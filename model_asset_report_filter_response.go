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

// AssetReportFilterResponse AssetReportFilterResponse defines the response schema for `/asset_report/filter`
type AssetReportFilterResponse struct {
	// A token that can be provided to endpoints such as `/asset_report/get` or `/asset_report/pdf/get` to fetch or update an Asset Report.
	AssetReportToken string `json:"asset_report_token"`
	// A unique ID identifying an Asset Report. Like all Plaid identifiers, this ID is case sensitive.
	AssetReportId string `json:"asset_report_id"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportFilterResponse AssetReportFilterResponse

// NewAssetReportFilterResponse instantiates a new AssetReportFilterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportFilterResponse(assetReportToken string, assetReportId string, requestId string) *AssetReportFilterResponse {
	this := AssetReportFilterResponse{}
	this.AssetReportToken = assetReportToken
	this.AssetReportId = assetReportId
	this.RequestId = requestId
	return &this
}

// NewAssetReportFilterResponseWithDefaults instantiates a new AssetReportFilterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportFilterResponseWithDefaults() *AssetReportFilterResponse {
	this := AssetReportFilterResponse{}
	return &this
}

// GetAssetReportToken returns the AssetReportToken field value
func (o *AssetReportFilterResponse) GetAssetReportToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportToken
}

// GetAssetReportTokenOk returns a tuple with the AssetReportToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportFilterResponse) GetAssetReportTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetReportToken, true
}

// SetAssetReportToken sets field value
func (o *AssetReportFilterResponse) SetAssetReportToken(v string) {
	o.AssetReportToken = v
}

// GetAssetReportId returns the AssetReportId field value
func (o *AssetReportFilterResponse) GetAssetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportId
}

// GetAssetReportIdOk returns a tuple with the AssetReportId field value
// and a boolean to check if the value has been set.
func (o *AssetReportFilterResponse) GetAssetReportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetReportId, true
}

// SetAssetReportId sets field value
func (o *AssetReportFilterResponse) SetAssetReportId(v string) {
	o.AssetReportId = v
}

// GetRequestId returns the RequestId field value
func (o *AssetReportFilterResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AssetReportFilterResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AssetReportFilterResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o AssetReportFilterResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset_report_token"] = o.AssetReportToken
	}
	if true {
		toSerialize["asset_report_id"] = o.AssetReportId
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportFilterResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportFilterResponse := _AssetReportFilterResponse{}

	if err = json.Unmarshal(bytes, &varAssetReportFilterResponse); err == nil {
		*o = AssetReportFilterResponse(varAssetReportFilterResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "asset_report_token")
		delete(additionalProperties, "asset_report_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportFilterResponse struct {
	value *AssetReportFilterResponse
	isSet bool
}

func (v NullableAssetReportFilterResponse) Get() *AssetReportFilterResponse {
	return v.value
}

func (v *NullableAssetReportFilterResponse) Set(val *AssetReportFilterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportFilterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportFilterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportFilterResponse(val *AssetReportFilterResponse) *NullableAssetReportFilterResponse {
	return &NullableAssetReportFilterResponse{value: val, isSet: true}
}

func (v NullableAssetReportFilterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportFilterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


