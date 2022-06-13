# IncomeVerificationPrecheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**User** | Pointer to [**NullableIncomeVerificationPrecheckUser**](IncomeVerificationPrecheckUser.md) |  | [optional] 
**Employer** | Pointer to [**NullableIncomeVerificationPrecheckEmployer**](IncomeVerificationPrecheckEmployer.md) |  | [optional] 
**TransactionsAccessToken** | Pointer to **NullableString** |  | [optional] 
**TransactionsAccessTokens** | Pointer to **[]string** | An array of access tokens corresponding to Items belonging to the user whose eligibility is being checked. Note that if the Items specified here are not already initialized with &#x60;transactions&#x60;, providing them in this field will cause these Items to be initialized with (and billed for) the Transactions product. | [optional] 
**UsMilitaryInfo** | Pointer to [**NullableIncomeVerificationPrecheckMilitaryInfo**](IncomeVerificationPrecheckMilitaryInfo.md) |  | [optional] 

## Methods

### NewIncomeVerificationPrecheckRequest

`func NewIncomeVerificationPrecheckRequest() *IncomeVerificationPrecheckRequest`

NewIncomeVerificationPrecheckRequest instantiates a new IncomeVerificationPrecheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationPrecheckRequestWithDefaults

`func NewIncomeVerificationPrecheckRequestWithDefaults() *IncomeVerificationPrecheckRequest`

NewIncomeVerificationPrecheckRequestWithDefaults instantiates a new IncomeVerificationPrecheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IncomeVerificationPrecheckRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IncomeVerificationPrecheckRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IncomeVerificationPrecheckRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IncomeVerificationPrecheckRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *IncomeVerificationPrecheckRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IncomeVerificationPrecheckRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IncomeVerificationPrecheckRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IncomeVerificationPrecheckRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUser

`func (o *IncomeVerificationPrecheckRequest) GetUser() IncomeVerificationPrecheckUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IncomeVerificationPrecheckRequest) GetUserOk() (*IncomeVerificationPrecheckUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IncomeVerificationPrecheckRequest) SetUser(v IncomeVerificationPrecheckUser)`

SetUser sets User field to given value.

### HasUser

`func (o *IncomeVerificationPrecheckRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *IncomeVerificationPrecheckRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *IncomeVerificationPrecheckRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetEmployer

`func (o *IncomeVerificationPrecheckRequest) GetEmployer() IncomeVerificationPrecheckEmployer`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *IncomeVerificationPrecheckRequest) GetEmployerOk() (*IncomeVerificationPrecheckEmployer, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *IncomeVerificationPrecheckRequest) SetEmployer(v IncomeVerificationPrecheckEmployer)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *IncomeVerificationPrecheckRequest) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### SetEmployerNil

`func (o *IncomeVerificationPrecheckRequest) SetEmployerNil(b bool)`

 SetEmployerNil sets the value for Employer to be an explicit nil

### UnsetEmployer
`func (o *IncomeVerificationPrecheckRequest) UnsetEmployer()`

UnsetEmployer ensures that no value is present for Employer, not even an explicit nil
### GetTransactionsAccessToken

`func (o *IncomeVerificationPrecheckRequest) GetTransactionsAccessToken() string`

GetTransactionsAccessToken returns the TransactionsAccessToken field if non-nil, zero value otherwise.

### GetTransactionsAccessTokenOk

`func (o *IncomeVerificationPrecheckRequest) GetTransactionsAccessTokenOk() (*string, bool)`

GetTransactionsAccessTokenOk returns a tuple with the TransactionsAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsAccessToken

`func (o *IncomeVerificationPrecheckRequest) SetTransactionsAccessToken(v string)`

SetTransactionsAccessToken sets TransactionsAccessToken field to given value.

### HasTransactionsAccessToken

`func (o *IncomeVerificationPrecheckRequest) HasTransactionsAccessToken() bool`

HasTransactionsAccessToken returns a boolean if a field has been set.

### SetTransactionsAccessTokenNil

`func (o *IncomeVerificationPrecheckRequest) SetTransactionsAccessTokenNil(b bool)`

 SetTransactionsAccessTokenNil sets the value for TransactionsAccessToken to be an explicit nil

### UnsetTransactionsAccessToken
`func (o *IncomeVerificationPrecheckRequest) UnsetTransactionsAccessToken()`

UnsetTransactionsAccessToken ensures that no value is present for TransactionsAccessToken, not even an explicit nil
### GetTransactionsAccessTokens

`func (o *IncomeVerificationPrecheckRequest) GetTransactionsAccessTokens() []string`

GetTransactionsAccessTokens returns the TransactionsAccessTokens field if non-nil, zero value otherwise.

### GetTransactionsAccessTokensOk

`func (o *IncomeVerificationPrecheckRequest) GetTransactionsAccessTokensOk() (*[]string, bool)`

GetTransactionsAccessTokensOk returns a tuple with the TransactionsAccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsAccessTokens

`func (o *IncomeVerificationPrecheckRequest) SetTransactionsAccessTokens(v []string)`

SetTransactionsAccessTokens sets TransactionsAccessTokens field to given value.

### HasTransactionsAccessTokens

`func (o *IncomeVerificationPrecheckRequest) HasTransactionsAccessTokens() bool`

HasTransactionsAccessTokens returns a boolean if a field has been set.

### GetUsMilitaryInfo

`func (o *IncomeVerificationPrecheckRequest) GetUsMilitaryInfo() IncomeVerificationPrecheckMilitaryInfo`

GetUsMilitaryInfo returns the UsMilitaryInfo field if non-nil, zero value otherwise.

### GetUsMilitaryInfoOk

`func (o *IncomeVerificationPrecheckRequest) GetUsMilitaryInfoOk() (*IncomeVerificationPrecheckMilitaryInfo, bool)`

GetUsMilitaryInfoOk returns a tuple with the UsMilitaryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsMilitaryInfo

`func (o *IncomeVerificationPrecheckRequest) SetUsMilitaryInfo(v IncomeVerificationPrecheckMilitaryInfo)`

SetUsMilitaryInfo sets UsMilitaryInfo field to given value.

### HasUsMilitaryInfo

`func (o *IncomeVerificationPrecheckRequest) HasUsMilitaryInfo() bool`

HasUsMilitaryInfo returns a boolean if a field has been set.

### SetUsMilitaryInfoNil

`func (o *IncomeVerificationPrecheckRequest) SetUsMilitaryInfoNil(b bool)`

 SetUsMilitaryInfoNil sets the value for UsMilitaryInfo to be an explicit nil

### UnsetUsMilitaryInfo
`func (o *IncomeVerificationPrecheckRequest) UnsetUsMilitaryInfo()`

UnsetUsMilitaryInfo ensures that no value is present for UsMilitaryInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


