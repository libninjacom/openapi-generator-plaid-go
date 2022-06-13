# IncomeVerificationPaystubsGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**IncomeVerificationId** | Pointer to **NullableString** | The ID of the verification for which to get paystub information. | [optional] 
**AccessToken** | Pointer to **NullableString** | The access token associated with the Item data is being requested for. | [optional] 

## Methods

### NewIncomeVerificationPaystubsGetRequest

`func NewIncomeVerificationPaystubsGetRequest() *IncomeVerificationPaystubsGetRequest`

NewIncomeVerificationPaystubsGetRequest instantiates a new IncomeVerificationPaystubsGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationPaystubsGetRequestWithDefaults

`func NewIncomeVerificationPaystubsGetRequestWithDefaults() *IncomeVerificationPaystubsGetRequest`

NewIncomeVerificationPaystubsGetRequestWithDefaults instantiates a new IncomeVerificationPaystubsGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IncomeVerificationPaystubsGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IncomeVerificationPaystubsGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IncomeVerificationPaystubsGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IncomeVerificationPaystubsGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *IncomeVerificationPaystubsGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IncomeVerificationPaystubsGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IncomeVerificationPaystubsGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IncomeVerificationPaystubsGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetIncomeVerificationId

`func (o *IncomeVerificationPaystubsGetRequest) GetIncomeVerificationId() string`

GetIncomeVerificationId returns the IncomeVerificationId field if non-nil, zero value otherwise.

### GetIncomeVerificationIdOk

`func (o *IncomeVerificationPaystubsGetRequest) GetIncomeVerificationIdOk() (*string, bool)`

GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeVerificationId

`func (o *IncomeVerificationPaystubsGetRequest) SetIncomeVerificationId(v string)`

SetIncomeVerificationId sets IncomeVerificationId field to given value.

### HasIncomeVerificationId

`func (o *IncomeVerificationPaystubsGetRequest) HasIncomeVerificationId() bool`

HasIncomeVerificationId returns a boolean if a field has been set.

### SetIncomeVerificationIdNil

`func (o *IncomeVerificationPaystubsGetRequest) SetIncomeVerificationIdNil(b bool)`

 SetIncomeVerificationIdNil sets the value for IncomeVerificationId to be an explicit nil

### UnsetIncomeVerificationId
`func (o *IncomeVerificationPaystubsGetRequest) UnsetIncomeVerificationId()`

UnsetIncomeVerificationId ensures that no value is present for IncomeVerificationId, not even an explicit nil
### GetAccessToken

`func (o *IncomeVerificationPaystubsGetRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *IncomeVerificationPaystubsGetRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *IncomeVerificationPaystubsGetRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *IncomeVerificationPaystubsGetRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *IncomeVerificationPaystubsGetRequest) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *IncomeVerificationPaystubsGetRequest) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


