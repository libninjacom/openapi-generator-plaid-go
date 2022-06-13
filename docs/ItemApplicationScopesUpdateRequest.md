# ItemApplicationScopesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**ApplicationId** | **string** | This field will map to the application ID that is returned from /item/applications/list, or provided to the institution in an oauth redirect. | 
**Scopes** | [**Scopes**](Scopes.md) |  | 
**State** | Pointer to **string** | When scopes are updated during enrollment, this field must be populated with the state sent to the partner in the OAuth Login URI. This field is required when the context is &#x60;ENROLLMENT&#x60;. | [optional] 
**Context** | [**ScopesContext**](ScopesContext.md) |  | 

## Methods

### NewItemApplicationScopesUpdateRequest

`func NewItemApplicationScopesUpdateRequest(accessToken string, applicationId string, scopes Scopes, context ScopesContext, ) *ItemApplicationScopesUpdateRequest`

NewItemApplicationScopesUpdateRequest instantiates a new ItemApplicationScopesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemApplicationScopesUpdateRequestWithDefaults

`func NewItemApplicationScopesUpdateRequestWithDefaults() *ItemApplicationScopesUpdateRequest`

NewItemApplicationScopesUpdateRequestWithDefaults instantiates a new ItemApplicationScopesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ItemApplicationScopesUpdateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ItemApplicationScopesUpdateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ItemApplicationScopesUpdateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ItemApplicationScopesUpdateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *ItemApplicationScopesUpdateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ItemApplicationScopesUpdateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ItemApplicationScopesUpdateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ItemApplicationScopesUpdateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAccessToken

`func (o *ItemApplicationScopesUpdateRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ItemApplicationScopesUpdateRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ItemApplicationScopesUpdateRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetApplicationId

`func (o *ItemApplicationScopesUpdateRequest) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ItemApplicationScopesUpdateRequest) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ItemApplicationScopesUpdateRequest) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetScopes

`func (o *ItemApplicationScopesUpdateRequest) GetScopes() Scopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ItemApplicationScopesUpdateRequest) GetScopesOk() (*Scopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ItemApplicationScopesUpdateRequest) SetScopes(v Scopes)`

SetScopes sets Scopes field to given value.


### GetState

`func (o *ItemApplicationScopesUpdateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ItemApplicationScopesUpdateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ItemApplicationScopesUpdateRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ItemApplicationScopesUpdateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetContext

`func (o *ItemApplicationScopesUpdateRequest) GetContext() ScopesContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ItemApplicationScopesUpdateRequest) GetContextOk() (*ScopesContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ItemApplicationScopesUpdateRequest) SetContext(v ScopesContext)`

SetContext sets Context field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


