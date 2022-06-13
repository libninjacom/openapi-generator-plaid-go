# DepositSwitchAltCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**TargetAccount** | [**DepositSwitchTargetAccount**](DepositSwitchTargetAccount.md) |  | 
**TargetUser** | [**DepositSwitchTargetUser**](DepositSwitchTargetUser.md) |  | 
**Options** | Pointer to [**DepositSwitchCreateRequestOptions**](DepositSwitchCreateRequestOptions.md) |  | [optional] 
**CountryCode** | Pointer to **NullableString** | ISO-3166-1 alpha-2 country code standard. | [optional] 

## Methods

### NewDepositSwitchAltCreateRequest

`func NewDepositSwitchAltCreateRequest(targetAccount DepositSwitchTargetAccount, targetUser DepositSwitchTargetUser, ) *DepositSwitchAltCreateRequest`

NewDepositSwitchAltCreateRequest instantiates a new DepositSwitchAltCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchAltCreateRequestWithDefaults

`func NewDepositSwitchAltCreateRequestWithDefaults() *DepositSwitchAltCreateRequest`

NewDepositSwitchAltCreateRequestWithDefaults instantiates a new DepositSwitchAltCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *DepositSwitchAltCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DepositSwitchAltCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DepositSwitchAltCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *DepositSwitchAltCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *DepositSwitchAltCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DepositSwitchAltCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DepositSwitchAltCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DepositSwitchAltCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTargetAccount

`func (o *DepositSwitchAltCreateRequest) GetTargetAccount() DepositSwitchTargetAccount`

GetTargetAccount returns the TargetAccount field if non-nil, zero value otherwise.

### GetTargetAccountOk

`func (o *DepositSwitchAltCreateRequest) GetTargetAccountOk() (*DepositSwitchTargetAccount, bool)`

GetTargetAccountOk returns a tuple with the TargetAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccount

`func (o *DepositSwitchAltCreateRequest) SetTargetAccount(v DepositSwitchTargetAccount)`

SetTargetAccount sets TargetAccount field to given value.


### GetTargetUser

`func (o *DepositSwitchAltCreateRequest) GetTargetUser() DepositSwitchTargetUser`

GetTargetUser returns the TargetUser field if non-nil, zero value otherwise.

### GetTargetUserOk

`func (o *DepositSwitchAltCreateRequest) GetTargetUserOk() (*DepositSwitchTargetUser, bool)`

GetTargetUserOk returns a tuple with the TargetUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUser

`func (o *DepositSwitchAltCreateRequest) SetTargetUser(v DepositSwitchTargetUser)`

SetTargetUser sets TargetUser field to given value.


### GetOptions

`func (o *DepositSwitchAltCreateRequest) GetOptions() DepositSwitchCreateRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DepositSwitchAltCreateRequest) GetOptionsOk() (*DepositSwitchCreateRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DepositSwitchAltCreateRequest) SetOptions(v DepositSwitchCreateRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DepositSwitchAltCreateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetCountryCode

`func (o *DepositSwitchAltCreateRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *DepositSwitchAltCreateRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *DepositSwitchAltCreateRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *DepositSwitchAltCreateRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *DepositSwitchAltCreateRequest) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *DepositSwitchAltCreateRequest) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


