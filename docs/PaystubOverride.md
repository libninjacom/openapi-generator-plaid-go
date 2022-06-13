# PaystubOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Employer** | Pointer to [**PaystubOverrideEmployer**](PaystubOverrideEmployer.md) |  | [optional] 
**Employee** | Pointer to [**PaystubOverrideEmployee**](PaystubOverrideEmployee.md) |  | [optional] 
**IncomeBreakdown** | Pointer to [**[]IncomeBreakdown**](IncomeBreakdown.md) |  | [optional] 
**PayPeriodDetails** | Pointer to [**PayPeriodDetails**](PayPeriodDetails.md) |  | [optional] 

## Methods

### NewPaystubOverride

`func NewPaystubOverride() *PaystubOverride`

NewPaystubOverride instantiates a new PaystubOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaystubOverrideWithDefaults

`func NewPaystubOverrideWithDefaults() *PaystubOverride`

NewPaystubOverrideWithDefaults instantiates a new PaystubOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployer

`func (o *PaystubOverride) GetEmployer() PaystubOverrideEmployer`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *PaystubOverride) GetEmployerOk() (*PaystubOverrideEmployer, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *PaystubOverride) SetEmployer(v PaystubOverrideEmployer)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *PaystubOverride) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### GetEmployee

`func (o *PaystubOverride) GetEmployee() PaystubOverrideEmployee`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *PaystubOverride) GetEmployeeOk() (*PaystubOverrideEmployee, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *PaystubOverride) SetEmployee(v PaystubOverrideEmployee)`

SetEmployee sets Employee field to given value.

### HasEmployee

`func (o *PaystubOverride) HasEmployee() bool`

HasEmployee returns a boolean if a field has been set.

### GetIncomeBreakdown

`func (o *PaystubOverride) GetIncomeBreakdown() []IncomeBreakdown`

GetIncomeBreakdown returns the IncomeBreakdown field if non-nil, zero value otherwise.

### GetIncomeBreakdownOk

`func (o *PaystubOverride) GetIncomeBreakdownOk() (*[]IncomeBreakdown, bool)`

GetIncomeBreakdownOk returns a tuple with the IncomeBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeBreakdown

`func (o *PaystubOverride) SetIncomeBreakdown(v []IncomeBreakdown)`

SetIncomeBreakdown sets IncomeBreakdown field to given value.

### HasIncomeBreakdown

`func (o *PaystubOverride) HasIncomeBreakdown() bool`

HasIncomeBreakdown returns a boolean if a field has been set.

### GetPayPeriodDetails

`func (o *PaystubOverride) GetPayPeriodDetails() PayPeriodDetails`

GetPayPeriodDetails returns the PayPeriodDetails field if non-nil, zero value otherwise.

### GetPayPeriodDetailsOk

`func (o *PaystubOverride) GetPayPeriodDetailsOk() (*PayPeriodDetails, bool)`

GetPayPeriodDetailsOk returns a tuple with the PayPeriodDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriodDetails

`func (o *PaystubOverride) SetPayPeriodDetails(v PayPeriodDetails)`

SetPayPeriodDetails sets PayPeriodDetails field to given value.

### HasPayPeriodDetails

`func (o *PaystubOverride) HasPayPeriodDetails() bool`

HasPayPeriodDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


