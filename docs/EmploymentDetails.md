# EmploymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnualSalary** | Pointer to [**Pay**](Pay.md) |  | [optional] 
**HireDate** | Pointer to **NullableString** | Date on which the employee was hired, in the YYYY-MM-DD format. | [optional] 

## Methods

### NewEmploymentDetails

`func NewEmploymentDetails() *EmploymentDetails`

NewEmploymentDetails instantiates a new EmploymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmploymentDetailsWithDefaults

`func NewEmploymentDetailsWithDefaults() *EmploymentDetails`

NewEmploymentDetailsWithDefaults instantiates a new EmploymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnualSalary

`func (o *EmploymentDetails) GetAnnualSalary() Pay`

GetAnnualSalary returns the AnnualSalary field if non-nil, zero value otherwise.

### GetAnnualSalaryOk

`func (o *EmploymentDetails) GetAnnualSalaryOk() (*Pay, bool)`

GetAnnualSalaryOk returns a tuple with the AnnualSalary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualSalary

`func (o *EmploymentDetails) SetAnnualSalary(v Pay)`

SetAnnualSalary sets AnnualSalary field to given value.

### HasAnnualSalary

`func (o *EmploymentDetails) HasAnnualSalary() bool`

HasAnnualSalary returns a boolean if a field has been set.

### GetHireDate

`func (o *EmploymentDetails) GetHireDate() string`

GetHireDate returns the HireDate field if non-nil, zero value otherwise.

### GetHireDateOk

`func (o *EmploymentDetails) GetHireDateOk() (*string, bool)`

GetHireDateOk returns a tuple with the HireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireDate

`func (o *EmploymentDetails) SetHireDate(v string)`

SetHireDate sets HireDate field to given value.

### HasHireDate

`func (o *EmploymentDetails) HasHireDate() bool`

HasHireDate returns a boolean if a field has been set.

### SetHireDateNil

`func (o *EmploymentDetails) SetHireDateNil(b bool)`

 SetHireDateNil sets the value for HireDate to be an explicit nil

### UnsetHireDate
`func (o *EmploymentDetails) UnsetHireDate()`

UnsetHireDate ensures that no value is present for HireDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


