# IncomeVerificationCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**Webhook** | **string** | The URL endpoint to which Plaid should send webhooks related to the progress of the income verification process. | 
**PrecheckId** | Pointer to **string** | The ID of a precheck created with &#x60;/income/verification/precheck&#x60;. Will be used to improve conversion of the income verification flow. | [optional] 
**Options** | Pointer to [**IncomeVerificationCreateRequestOptions**](IncomeVerificationCreateRequestOptions.md) |  | [optional] 

## Methods

### NewIncomeVerificationCreateRequest

`func NewIncomeVerificationCreateRequest(webhook string, ) *IncomeVerificationCreateRequest`

NewIncomeVerificationCreateRequest instantiates a new IncomeVerificationCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationCreateRequestWithDefaults

`func NewIncomeVerificationCreateRequestWithDefaults() *IncomeVerificationCreateRequest`

NewIncomeVerificationCreateRequestWithDefaults instantiates a new IncomeVerificationCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IncomeVerificationCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IncomeVerificationCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IncomeVerificationCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IncomeVerificationCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *IncomeVerificationCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IncomeVerificationCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IncomeVerificationCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IncomeVerificationCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetWebhook

`func (o *IncomeVerificationCreateRequest) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *IncomeVerificationCreateRequest) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *IncomeVerificationCreateRequest) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.


### GetPrecheckId

`func (o *IncomeVerificationCreateRequest) GetPrecheckId() string`

GetPrecheckId returns the PrecheckId field if non-nil, zero value otherwise.

### GetPrecheckIdOk

`func (o *IncomeVerificationCreateRequest) GetPrecheckIdOk() (*string, bool)`

GetPrecheckIdOk returns a tuple with the PrecheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecheckId

`func (o *IncomeVerificationCreateRequest) SetPrecheckId(v string)`

SetPrecheckId sets PrecheckId field to given value.

### HasPrecheckId

`func (o *IncomeVerificationCreateRequest) HasPrecheckId() bool`

HasPrecheckId returns a boolean if a field has been set.

### GetOptions

`func (o *IncomeVerificationCreateRequest) GetOptions() IncomeVerificationCreateRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *IncomeVerificationCreateRequest) GetOptionsOk() (*IncomeVerificationCreateRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *IncomeVerificationCreateRequest) SetOptions(v IncomeVerificationCreateRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *IncomeVerificationCreateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


