# IncomeVerificationTaxformsGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**IncomeVerificationId** | Pointer to **NullableString** | The ID of the verification. | [optional] 
**AccessToken** | Pointer to **NullableString** | The access token associated with the Item data is being requested for. | [optional] 

## Methods

### NewIncomeVerificationTaxformsGetRequest

`func NewIncomeVerificationTaxformsGetRequest() *IncomeVerificationTaxformsGetRequest`

NewIncomeVerificationTaxformsGetRequest instantiates a new IncomeVerificationTaxformsGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationTaxformsGetRequestWithDefaults

`func NewIncomeVerificationTaxformsGetRequestWithDefaults() *IncomeVerificationTaxformsGetRequest`

NewIncomeVerificationTaxformsGetRequestWithDefaults instantiates a new IncomeVerificationTaxformsGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IncomeVerificationTaxformsGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IncomeVerificationTaxformsGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IncomeVerificationTaxformsGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IncomeVerificationTaxformsGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *IncomeVerificationTaxformsGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IncomeVerificationTaxformsGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IncomeVerificationTaxformsGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IncomeVerificationTaxformsGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetIncomeVerificationId

`func (o *IncomeVerificationTaxformsGetRequest) GetIncomeVerificationId() string`

GetIncomeVerificationId returns the IncomeVerificationId field if non-nil, zero value otherwise.

### GetIncomeVerificationIdOk

`func (o *IncomeVerificationTaxformsGetRequest) GetIncomeVerificationIdOk() (*string, bool)`

GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeVerificationId

`func (o *IncomeVerificationTaxformsGetRequest) SetIncomeVerificationId(v string)`

SetIncomeVerificationId sets IncomeVerificationId field to given value.

### HasIncomeVerificationId

`func (o *IncomeVerificationTaxformsGetRequest) HasIncomeVerificationId() bool`

HasIncomeVerificationId returns a boolean if a field has been set.

### SetIncomeVerificationIdNil

`func (o *IncomeVerificationTaxformsGetRequest) SetIncomeVerificationIdNil(b bool)`

 SetIncomeVerificationIdNil sets the value for IncomeVerificationId to be an explicit nil

### UnsetIncomeVerificationId
`func (o *IncomeVerificationTaxformsGetRequest) UnsetIncomeVerificationId()`

UnsetIncomeVerificationId ensures that no value is present for IncomeVerificationId, not even an explicit nil
### GetAccessToken

`func (o *IncomeVerificationTaxformsGetRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *IncomeVerificationTaxformsGetRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *IncomeVerificationTaxformsGetRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *IncomeVerificationTaxformsGetRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *IncomeVerificationTaxformsGetRequest) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *IncomeVerificationTaxformsGetRequest) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


