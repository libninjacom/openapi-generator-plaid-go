# DepositSwitchCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**TargetAccessToken** | **string** | Access token for the target Item, typically provided in the Import Item response.  | 
**TargetAccountId** | **string** | Plaid Account ID that specifies the target bank account. This account will become the recipient for a user&#39;s direct deposit. | 
**CountryCode** | Pointer to **NullableString** | ISO-3166-1 alpha-2 country code standard. | [optional] 
**Options** | Pointer to [**DepositSwitchCreateRequestOptions**](DepositSwitchCreateRequestOptions.md) |  | [optional] 

## Methods

### NewDepositSwitchCreateRequest

`func NewDepositSwitchCreateRequest(targetAccessToken string, targetAccountId string, ) *DepositSwitchCreateRequest`

NewDepositSwitchCreateRequest instantiates a new DepositSwitchCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchCreateRequestWithDefaults

`func NewDepositSwitchCreateRequestWithDefaults() *DepositSwitchCreateRequest`

NewDepositSwitchCreateRequestWithDefaults instantiates a new DepositSwitchCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *DepositSwitchCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DepositSwitchCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DepositSwitchCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *DepositSwitchCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *DepositSwitchCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DepositSwitchCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DepositSwitchCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DepositSwitchCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTargetAccessToken

`func (o *DepositSwitchCreateRequest) GetTargetAccessToken() string`

GetTargetAccessToken returns the TargetAccessToken field if non-nil, zero value otherwise.

### GetTargetAccessTokenOk

`func (o *DepositSwitchCreateRequest) GetTargetAccessTokenOk() (*string, bool)`

GetTargetAccessTokenOk returns a tuple with the TargetAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccessToken

`func (o *DepositSwitchCreateRequest) SetTargetAccessToken(v string)`

SetTargetAccessToken sets TargetAccessToken field to given value.


### GetTargetAccountId

`func (o *DepositSwitchCreateRequest) GetTargetAccountId() string`

GetTargetAccountId returns the TargetAccountId field if non-nil, zero value otherwise.

### GetTargetAccountIdOk

`func (o *DepositSwitchCreateRequest) GetTargetAccountIdOk() (*string, bool)`

GetTargetAccountIdOk returns a tuple with the TargetAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccountId

`func (o *DepositSwitchCreateRequest) SetTargetAccountId(v string)`

SetTargetAccountId sets TargetAccountId field to given value.


### GetCountryCode

`func (o *DepositSwitchCreateRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *DepositSwitchCreateRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *DepositSwitchCreateRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *DepositSwitchCreateRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *DepositSwitchCreateRequest) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *DepositSwitchCreateRequest) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetOptions

`func (o *DepositSwitchCreateRequest) GetOptions() DepositSwitchCreateRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DepositSwitchCreateRequest) GetOptionsOk() (*DepositSwitchCreateRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DepositSwitchCreateRequest) SetOptions(v DepositSwitchCreateRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DepositSwitchCreateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


