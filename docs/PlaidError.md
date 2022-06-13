# PlaidError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorType** | **string** | A broad categorization of the error. Safe for programmatic use. | 
**ErrorCode** | **string** | The particular error code. Safe for programmatic use. | 
**ErrorMessage** | **string** | A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use. | 
**DisplayMessage** | **NullableString** | A user-friendly representation of the error code. &#x60;null&#x60; if the error is not related to user action.  This may change over time and is not safe for programmatic use. | 
**RequestId** | Pointer to **string** | A unique ID identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks. | [optional] 
**Causes** | Pointer to **[]interface{}** | In the Assets product, a request can pertain to more than one Item. If an error is returned for such a request, &#x60;causes&#x60; will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  &#x60;causes&#x60; will only be provided for the &#x60;error_type&#x60; &#x60;ASSET_REPORT_ERROR&#x60;. &#x60;causes&#x60; will also not be populated inside an error nested within a &#x60;warning&#x60; object. | [optional] 
**Status** | Pointer to **NullableFloat32** | The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook. | [optional] 
**DocumentationUrl** | Pointer to **string** | The URL of a Plaid documentation page with more information about the error | [optional] 
**SuggestedAction** | Pointer to **string** | Suggested steps for resolving the error | [optional] 

## Methods

### NewPlaidError

`func NewPlaidError(errorType string, errorCode string, errorMessage string, displayMessage NullableString, ) *PlaidError`

NewPlaidError instantiates a new PlaidError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaidErrorWithDefaults

`func NewPlaidErrorWithDefaults() *PlaidError`

NewPlaidErrorWithDefaults instantiates a new PlaidError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorType

`func (o *PlaidError) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *PlaidError) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *PlaidError) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.


### GetErrorCode

`func (o *PlaidError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PlaidError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *PlaidError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *PlaidError) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PlaidError) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PlaidError) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetDisplayMessage

`func (o *PlaidError) GetDisplayMessage() string`

GetDisplayMessage returns the DisplayMessage field if non-nil, zero value otherwise.

### GetDisplayMessageOk

`func (o *PlaidError) GetDisplayMessageOk() (*string, bool)`

GetDisplayMessageOk returns a tuple with the DisplayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMessage

`func (o *PlaidError) SetDisplayMessage(v string)`

SetDisplayMessage sets DisplayMessage field to given value.


### SetDisplayMessageNil

`func (o *PlaidError) SetDisplayMessageNil(b bool)`

 SetDisplayMessageNil sets the value for DisplayMessage to be an explicit nil

### UnsetDisplayMessage
`func (o *PlaidError) UnsetDisplayMessage()`

UnsetDisplayMessage ensures that no value is present for DisplayMessage, not even an explicit nil
### GetRequestId

`func (o *PlaidError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PlaidError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PlaidError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *PlaidError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCauses

`func (o *PlaidError) GetCauses() []interface{}`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *PlaidError) GetCausesOk() (*[]interface{}, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *PlaidError) SetCauses(v []interface{})`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *PlaidError) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetStatus

`func (o *PlaidError) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlaidError) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlaidError) SetStatus(v float32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlaidError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PlaidError) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PlaidError) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDocumentationUrl

`func (o *PlaidError) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *PlaidError) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *PlaidError) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *PlaidError) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetSuggestedAction

`func (o *PlaidError) GetSuggestedAction() string`

GetSuggestedAction returns the SuggestedAction field if non-nil, zero value otherwise.

### GetSuggestedActionOk

`func (o *PlaidError) GetSuggestedActionOk() (*string, bool)`

GetSuggestedActionOk returns a tuple with the SuggestedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedAction

`func (o *PlaidError) SetSuggestedAction(v string)`

SetSuggestedAction sets SuggestedAction field to given value.

### HasSuggestedAction

`func (o *PlaidError) HasSuggestedAction() bool`

HasSuggestedAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


