# EmploymentVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NullableEmploymentVerificationStatus**](EmploymentVerificationStatus.md) |  | [optional] 
**StartDate** | Pointer to **NullableString** | Start of employment in ISO 8601 format (YYYY-MM-DD). | [optional] 
**EndDate** | Pointer to **NullableString** | End of employment, if applicable. Provided in ISO 8601 format (YYY-MM-DD). | [optional] 
**Employer** | Pointer to [**EmployerVerification**](EmployerVerification.md) |  | [optional] 
**Title** | Pointer to **NullableString** | Current title of employee. | [optional] 
**PlatformIds** | Pointer to [**PlatformIds**](PlatformIds.md) |  | [optional] 

## Methods

### NewEmploymentVerification

`func NewEmploymentVerification() *EmploymentVerification`

NewEmploymentVerification instantiates a new EmploymentVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmploymentVerificationWithDefaults

`func NewEmploymentVerificationWithDefaults() *EmploymentVerification`

NewEmploymentVerificationWithDefaults instantiates a new EmploymentVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EmploymentVerification) GetStatus() EmploymentVerificationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmploymentVerification) GetStatusOk() (*EmploymentVerificationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmploymentVerification) SetStatus(v EmploymentVerificationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmploymentVerification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *EmploymentVerification) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *EmploymentVerification) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStartDate

`func (o *EmploymentVerification) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EmploymentVerification) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EmploymentVerification) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EmploymentVerification) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *EmploymentVerification) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *EmploymentVerification) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *EmploymentVerification) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EmploymentVerification) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EmploymentVerification) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EmploymentVerification) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *EmploymentVerification) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *EmploymentVerification) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetEmployer

`func (o *EmploymentVerification) GetEmployer() EmployerVerification`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *EmploymentVerification) GetEmployerOk() (*EmployerVerification, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *EmploymentVerification) SetEmployer(v EmployerVerification)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *EmploymentVerification) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### GetTitle

`func (o *EmploymentVerification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EmploymentVerification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EmploymentVerification) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EmploymentVerification) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *EmploymentVerification) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *EmploymentVerification) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPlatformIds

`func (o *EmploymentVerification) GetPlatformIds() PlatformIds`

GetPlatformIds returns the PlatformIds field if non-nil, zero value otherwise.

### GetPlatformIdsOk

`func (o *EmploymentVerification) GetPlatformIdsOk() (*PlatformIds, bool)`

GetPlatformIdsOk returns a tuple with the PlatformIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIds

`func (o *EmploymentVerification) SetPlatformIds(v PlatformIds)`

SetPlatformIds sets PlatformIds field to given value.

### HasPlatformIds

`func (o *EmploymentVerification) HasPlatformIds() bool`

HasPlatformIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


