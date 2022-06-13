# ConnectedApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | This field will map to the application ID that is returned from /item/applications/list, or provided to the institution in an oauth redirect. | 
**Name** | **string** | The name of the application | 
**Logo** | **NullableString** | A URL that links to the application logo image (will be deprecated in the future, please use logo_url). | 
**LogoUrl** | **NullableString** | A URL that links to the application logo image. | 
**ApplicationUrl** | **NullableString** | The URL for the application&#39;s website | 
**ReasonForAccess** | **NullableString** | A string provided by the connected app stating why they use their respective enabled products. | 
**CreatedAt** | **string** | The date this application was linked in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format in UTC. | 
**JoinDate** | **string** | The date this application was granted production access at Plaid in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format in UTC. | 
**ProductDataTypes** | **[]string** | (Deprecated) A list of enums representing the data collected and products enabled for this connected application. | 
**Scopes** | Pointer to [**NullableScopesNullable**](ScopesNullable.md) |  | [optional] 
**RequestedScopes** | Pointer to [**RequestedScopes**](RequestedScopes.md) |  | [optional] 

## Methods

### NewConnectedApplication

`func NewConnectedApplication(applicationId string, name string, logo NullableString, logoUrl NullableString, applicationUrl NullableString, reasonForAccess NullableString, createdAt string, joinDate string, productDataTypes []string, ) *ConnectedApplication`

NewConnectedApplication instantiates a new ConnectedApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedApplicationWithDefaults

`func NewConnectedApplicationWithDefaults() *ConnectedApplication`

NewConnectedApplicationWithDefaults instantiates a new ConnectedApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ConnectedApplication) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ConnectedApplication) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ConnectedApplication) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetName

`func (o *ConnectedApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectedApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectedApplication) SetName(v string)`

SetName sets Name field to given value.


### GetLogo

`func (o *ConnectedApplication) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ConnectedApplication) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ConnectedApplication) SetLogo(v string)`

SetLogo sets Logo field to given value.


### SetLogoNil

`func (o *ConnectedApplication) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *ConnectedApplication) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetLogoUrl

`func (o *ConnectedApplication) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *ConnectedApplication) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *ConnectedApplication) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### SetLogoUrlNil

`func (o *ConnectedApplication) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *ConnectedApplication) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetApplicationUrl

`func (o *ConnectedApplication) GetApplicationUrl() string`

GetApplicationUrl returns the ApplicationUrl field if non-nil, zero value otherwise.

### GetApplicationUrlOk

`func (o *ConnectedApplication) GetApplicationUrlOk() (*string, bool)`

GetApplicationUrlOk returns a tuple with the ApplicationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUrl

`func (o *ConnectedApplication) SetApplicationUrl(v string)`

SetApplicationUrl sets ApplicationUrl field to given value.


### SetApplicationUrlNil

`func (o *ConnectedApplication) SetApplicationUrlNil(b bool)`

 SetApplicationUrlNil sets the value for ApplicationUrl to be an explicit nil

### UnsetApplicationUrl
`func (o *ConnectedApplication) UnsetApplicationUrl()`

UnsetApplicationUrl ensures that no value is present for ApplicationUrl, not even an explicit nil
### GetReasonForAccess

`func (o *ConnectedApplication) GetReasonForAccess() string`

GetReasonForAccess returns the ReasonForAccess field if non-nil, zero value otherwise.

### GetReasonForAccessOk

`func (o *ConnectedApplication) GetReasonForAccessOk() (*string, bool)`

GetReasonForAccessOk returns a tuple with the ReasonForAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForAccess

`func (o *ConnectedApplication) SetReasonForAccess(v string)`

SetReasonForAccess sets ReasonForAccess field to given value.


### SetReasonForAccessNil

`func (o *ConnectedApplication) SetReasonForAccessNil(b bool)`

 SetReasonForAccessNil sets the value for ReasonForAccess to be an explicit nil

### UnsetReasonForAccess
`func (o *ConnectedApplication) UnsetReasonForAccess()`

UnsetReasonForAccess ensures that no value is present for ReasonForAccess, not even an explicit nil
### GetCreatedAt

`func (o *ConnectedApplication) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectedApplication) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectedApplication) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetJoinDate

`func (o *ConnectedApplication) GetJoinDate() string`

GetJoinDate returns the JoinDate field if non-nil, zero value otherwise.

### GetJoinDateOk

`func (o *ConnectedApplication) GetJoinDateOk() (*string, bool)`

GetJoinDateOk returns a tuple with the JoinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDate

`func (o *ConnectedApplication) SetJoinDate(v string)`

SetJoinDate sets JoinDate field to given value.


### GetProductDataTypes

`func (o *ConnectedApplication) GetProductDataTypes() []string`

GetProductDataTypes returns the ProductDataTypes field if non-nil, zero value otherwise.

### GetProductDataTypesOk

`func (o *ConnectedApplication) GetProductDataTypesOk() (*[]string, bool)`

GetProductDataTypesOk returns a tuple with the ProductDataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDataTypes

`func (o *ConnectedApplication) SetProductDataTypes(v []string)`

SetProductDataTypes sets ProductDataTypes field to given value.


### GetScopes

`func (o *ConnectedApplication) GetScopes() ScopesNullable`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ConnectedApplication) GetScopesOk() (*ScopesNullable, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ConnectedApplication) SetScopes(v ScopesNullable)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ConnectedApplication) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *ConnectedApplication) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *ConnectedApplication) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetRequestedScopes

`func (o *ConnectedApplication) GetRequestedScopes() RequestedScopes`

GetRequestedScopes returns the RequestedScopes field if non-nil, zero value otherwise.

### GetRequestedScopesOk

`func (o *ConnectedApplication) GetRequestedScopesOk() (*RequestedScopes, bool)`

GetRequestedScopesOk returns a tuple with the RequestedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedScopes

`func (o *ConnectedApplication) SetRequestedScopes(v RequestedScopes)`

SetRequestedScopes sets RequestedScopes field to given value.

### HasRequestedScopes

`func (o *ConnectedApplication) HasRequestedScopes() bool`

HasRequestedScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


