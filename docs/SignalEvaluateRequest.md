# SignalEvaluateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**AccountId** | **string** | The &#x60;account_id&#x60; of the account whose verification status is to be modified | 
**ClientTransactionId** | **string** | The unique ID that you would like to use to refer to this transaction. For your convenience mapping your internal data, you could use your internal ID/identifier for this transaction. The max length for this field is 36 characters. | 
**Amount** | **float32** | The transaction amount, in USD (e.g. &#x60;102.05&#x60;) | 
**UserPresent** | Pointer to **NullableBool** | &#x60;true&#x60; if the end user is present while initiating the ACH transfer and the endpoint is being called; &#x60;false&#x60; otherwise (for example, when the ACH transfer is scheduled and the end user is not present, or you call this endpoint after the ACH transfer but before submitting the Nacha file for ACH processing). | [optional] 
**ClientUserId** | Pointer to **string** | A unique ID that identifies the end user in your system. This ID is used to correlate requests by a user with multiple Items. The max length for this field is 36 characters. | [optional] 
**User** | Pointer to [**SignalUser**](SignalUser.md) |  | [optional] 
**Device** | Pointer to [**SignalDevice**](SignalDevice.md) |  | [optional] 

## Methods

### NewSignalEvaluateRequest

`func NewSignalEvaluateRequest(accessToken string, accountId string, clientTransactionId string, amount float32, ) *SignalEvaluateRequest`

NewSignalEvaluateRequest instantiates a new SignalEvaluateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalEvaluateRequestWithDefaults

`func NewSignalEvaluateRequestWithDefaults() *SignalEvaluateRequest`

NewSignalEvaluateRequestWithDefaults instantiates a new SignalEvaluateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SignalEvaluateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SignalEvaluateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SignalEvaluateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SignalEvaluateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *SignalEvaluateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SignalEvaluateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SignalEvaluateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SignalEvaluateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAccessToken

`func (o *SignalEvaluateRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *SignalEvaluateRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *SignalEvaluateRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAccountId

`func (o *SignalEvaluateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SignalEvaluateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SignalEvaluateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetClientTransactionId

`func (o *SignalEvaluateRequest) GetClientTransactionId() string`

GetClientTransactionId returns the ClientTransactionId field if non-nil, zero value otherwise.

### GetClientTransactionIdOk

`func (o *SignalEvaluateRequest) GetClientTransactionIdOk() (*string, bool)`

GetClientTransactionIdOk returns a tuple with the ClientTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTransactionId

`func (o *SignalEvaluateRequest) SetClientTransactionId(v string)`

SetClientTransactionId sets ClientTransactionId field to given value.


### GetAmount

`func (o *SignalEvaluateRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SignalEvaluateRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SignalEvaluateRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetUserPresent

`func (o *SignalEvaluateRequest) GetUserPresent() bool`

GetUserPresent returns the UserPresent field if non-nil, zero value otherwise.

### GetUserPresentOk

`func (o *SignalEvaluateRequest) GetUserPresentOk() (*bool, bool)`

GetUserPresentOk returns a tuple with the UserPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPresent

`func (o *SignalEvaluateRequest) SetUserPresent(v bool)`

SetUserPresent sets UserPresent field to given value.

### HasUserPresent

`func (o *SignalEvaluateRequest) HasUserPresent() bool`

HasUserPresent returns a boolean if a field has been set.

### SetUserPresentNil

`func (o *SignalEvaluateRequest) SetUserPresentNil(b bool)`

 SetUserPresentNil sets the value for UserPresent to be an explicit nil

### UnsetUserPresent
`func (o *SignalEvaluateRequest) UnsetUserPresent()`

UnsetUserPresent ensures that no value is present for UserPresent, not even an explicit nil
### GetClientUserId

`func (o *SignalEvaluateRequest) GetClientUserId() string`

GetClientUserId returns the ClientUserId field if non-nil, zero value otherwise.

### GetClientUserIdOk

`func (o *SignalEvaluateRequest) GetClientUserIdOk() (*string, bool)`

GetClientUserIdOk returns a tuple with the ClientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUserId

`func (o *SignalEvaluateRequest) SetClientUserId(v string)`

SetClientUserId sets ClientUserId field to given value.

### HasClientUserId

`func (o *SignalEvaluateRequest) HasClientUserId() bool`

HasClientUserId returns a boolean if a field has been set.

### GetUser

`func (o *SignalEvaluateRequest) GetUser() SignalUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SignalEvaluateRequest) GetUserOk() (*SignalUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SignalEvaluateRequest) SetUser(v SignalUser)`

SetUser sets User field to given value.

### HasUser

`func (o *SignalEvaluateRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDevice

`func (o *SignalEvaluateRequest) GetDevice() SignalDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SignalEvaluateRequest) GetDeviceOk() (*SignalDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SignalEvaluateRequest) SetDevice(v SignalDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SignalEvaluateRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


