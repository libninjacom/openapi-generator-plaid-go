# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | This field will map to the application ID that is returned from /item/applications/list, or provided to the institution in an oauth redirect. | 
**Name** | **string** | The name of the application | 
**CreatedAt** | Pointer to **string** | The date this application was linked in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format in UTC. | [optional] 
**JoinDate** | **string** | The date this application was granted production access at Plaid in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format in UTC. | 
**LogoUrl** | **NullableString** | A URL that links to the application logo image. | 
**ApplicationUrl** | **NullableString** | The URL for the application&#39;s website | 
**ReasonForAccess** | **NullableString** | A string provided by the connected app stating why they use their respective enabled products. | 

## Methods

### NewApplication

`func NewApplication(applicationId string, name string, joinDate string, logoUrl NullableString, applicationUrl NullableString, reasonForAccess NullableString, ) *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *Application) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Application) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Application) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *Application) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Application) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Application) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Application) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetJoinDate

`func (o *Application) GetJoinDate() string`

GetJoinDate returns the JoinDate field if non-nil, zero value otherwise.

### GetJoinDateOk

`func (o *Application) GetJoinDateOk() (*string, bool)`

GetJoinDateOk returns a tuple with the JoinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDate

`func (o *Application) SetJoinDate(v string)`

SetJoinDate sets JoinDate field to given value.


### GetLogoUrl

`func (o *Application) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Application) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Application) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### SetLogoUrlNil

`func (o *Application) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *Application) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetApplicationUrl

`func (o *Application) GetApplicationUrl() string`

GetApplicationUrl returns the ApplicationUrl field if non-nil, zero value otherwise.

### GetApplicationUrlOk

`func (o *Application) GetApplicationUrlOk() (*string, bool)`

GetApplicationUrlOk returns a tuple with the ApplicationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUrl

`func (o *Application) SetApplicationUrl(v string)`

SetApplicationUrl sets ApplicationUrl field to given value.


### SetApplicationUrlNil

`func (o *Application) SetApplicationUrlNil(b bool)`

 SetApplicationUrlNil sets the value for ApplicationUrl to be an explicit nil

### UnsetApplicationUrl
`func (o *Application) UnsetApplicationUrl()`

UnsetApplicationUrl ensures that no value is present for ApplicationUrl, not even an explicit nil
### GetReasonForAccess

`func (o *Application) GetReasonForAccess() string`

GetReasonForAccess returns the ReasonForAccess field if non-nil, zero value otherwise.

### GetReasonForAccessOk

`func (o *Application) GetReasonForAccessOk() (*string, bool)`

GetReasonForAccessOk returns a tuple with the ReasonForAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForAccess

`func (o *Application) SetReasonForAccess(v string)`

SetReasonForAccess sets ReasonForAccess field to given value.


### SetReasonForAccessNil

`func (o *Application) SetReasonForAccessNil(b bool)`

 SetReasonForAccessNil sets the value for ReasonForAccess to be an explicit nil

### UnsetReasonForAccess
`func (o *Application) UnsetReasonForAccess()`

UnsetReasonForAccess ensures that no value is present for ReasonForAccess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


