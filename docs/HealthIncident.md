# HealthIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **time.Time** | The start date of the incident, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. &#x60;\&quot;2020-10-30T15:26:48Z\&quot;&#x60;. | 
**EndDate** | Pointer to **time.Time** | The end date of the incident, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. &#x60;\&quot;2020-10-30T15:26:48Z\&quot;&#x60;. | [optional] 
**Title** | **string** | The title of the incident | 
**IncidentUpdates** | [**[]IncidentUpdate**](IncidentUpdate.md) | Updates on the health incident. | 

## Methods

### NewHealthIncident

`func NewHealthIncident(startDate time.Time, title string, incidentUpdates []IncidentUpdate, ) *HealthIncident`

NewHealthIncident instantiates a new HealthIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthIncidentWithDefaults

`func NewHealthIncidentWithDefaults() *HealthIncident`

NewHealthIncidentWithDefaults instantiates a new HealthIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *HealthIncident) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *HealthIncident) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *HealthIncident) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *HealthIncident) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *HealthIncident) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *HealthIncident) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *HealthIncident) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetTitle

`func (o *HealthIncident) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HealthIncident) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HealthIncident) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIncidentUpdates

`func (o *HealthIncident) GetIncidentUpdates() []IncidentUpdate`

GetIncidentUpdates returns the IncidentUpdates field if non-nil, zero value otherwise.

### GetIncidentUpdatesOk

`func (o *HealthIncident) GetIncidentUpdatesOk() (*[]IncidentUpdate, bool)`

GetIncidentUpdatesOk returns a tuple with the IncidentUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentUpdates

`func (o *HealthIncident) SetIncidentUpdates(v []IncidentUpdate)`

SetIncidentUpdates sets IncidentUpdates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


