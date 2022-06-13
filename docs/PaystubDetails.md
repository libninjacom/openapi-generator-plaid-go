# PaystubDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayPeriodStartDate** | Pointer to **NullableString** | Beginning date of the pay period on the paystub in the &#39;YYYY-MM-DD&#39; format. | [optional] 
**PayPeriodEndDate** | Pointer to **NullableString** | Ending date of the pay period on the paystub in the &#39;YYYY-MM-DD&#39; format. | [optional] 
**PayDate** | Pointer to **NullableString** | Pay date on the paystub in the &#39;YYYY-MM-DD&#39; format. | [optional] 
**PaystubProvider** | Pointer to **NullableString** | The name of the payroll provider that generated the paystub, e.g. ADP | [optional] 
**PayFrequency** | Pointer to [**NullablePaystubPayFrequency**](PaystubPayFrequency.md) |  | [optional] 

## Methods

### NewPaystubDetails

`func NewPaystubDetails() *PaystubDetails`

NewPaystubDetails instantiates a new PaystubDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaystubDetailsWithDefaults

`func NewPaystubDetailsWithDefaults() *PaystubDetails`

NewPaystubDetailsWithDefaults instantiates a new PaystubDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayPeriodStartDate

`func (o *PaystubDetails) GetPayPeriodStartDate() string`

GetPayPeriodStartDate returns the PayPeriodStartDate field if non-nil, zero value otherwise.

### GetPayPeriodStartDateOk

`func (o *PaystubDetails) GetPayPeriodStartDateOk() (*string, bool)`

GetPayPeriodStartDateOk returns a tuple with the PayPeriodStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriodStartDate

`func (o *PaystubDetails) SetPayPeriodStartDate(v string)`

SetPayPeriodStartDate sets PayPeriodStartDate field to given value.

### HasPayPeriodStartDate

`func (o *PaystubDetails) HasPayPeriodStartDate() bool`

HasPayPeriodStartDate returns a boolean if a field has been set.

### SetPayPeriodStartDateNil

`func (o *PaystubDetails) SetPayPeriodStartDateNil(b bool)`

 SetPayPeriodStartDateNil sets the value for PayPeriodStartDate to be an explicit nil

### UnsetPayPeriodStartDate
`func (o *PaystubDetails) UnsetPayPeriodStartDate()`

UnsetPayPeriodStartDate ensures that no value is present for PayPeriodStartDate, not even an explicit nil
### GetPayPeriodEndDate

`func (o *PaystubDetails) GetPayPeriodEndDate() string`

GetPayPeriodEndDate returns the PayPeriodEndDate field if non-nil, zero value otherwise.

### GetPayPeriodEndDateOk

`func (o *PaystubDetails) GetPayPeriodEndDateOk() (*string, bool)`

GetPayPeriodEndDateOk returns a tuple with the PayPeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriodEndDate

`func (o *PaystubDetails) SetPayPeriodEndDate(v string)`

SetPayPeriodEndDate sets PayPeriodEndDate field to given value.

### HasPayPeriodEndDate

`func (o *PaystubDetails) HasPayPeriodEndDate() bool`

HasPayPeriodEndDate returns a boolean if a field has been set.

### SetPayPeriodEndDateNil

`func (o *PaystubDetails) SetPayPeriodEndDateNil(b bool)`

 SetPayPeriodEndDateNil sets the value for PayPeriodEndDate to be an explicit nil

### UnsetPayPeriodEndDate
`func (o *PaystubDetails) UnsetPayPeriodEndDate()`

UnsetPayPeriodEndDate ensures that no value is present for PayPeriodEndDate, not even an explicit nil
### GetPayDate

`func (o *PaystubDetails) GetPayDate() string`

GetPayDate returns the PayDate field if non-nil, zero value otherwise.

### GetPayDateOk

`func (o *PaystubDetails) GetPayDateOk() (*string, bool)`

GetPayDateOk returns a tuple with the PayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayDate

`func (o *PaystubDetails) SetPayDate(v string)`

SetPayDate sets PayDate field to given value.

### HasPayDate

`func (o *PaystubDetails) HasPayDate() bool`

HasPayDate returns a boolean if a field has been set.

### SetPayDateNil

`func (o *PaystubDetails) SetPayDateNil(b bool)`

 SetPayDateNil sets the value for PayDate to be an explicit nil

### UnsetPayDate
`func (o *PaystubDetails) UnsetPayDate()`

UnsetPayDate ensures that no value is present for PayDate, not even an explicit nil
### GetPaystubProvider

`func (o *PaystubDetails) GetPaystubProvider() string`

GetPaystubProvider returns the PaystubProvider field if non-nil, zero value otherwise.

### GetPaystubProviderOk

`func (o *PaystubDetails) GetPaystubProviderOk() (*string, bool)`

GetPaystubProviderOk returns a tuple with the PaystubProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaystubProvider

`func (o *PaystubDetails) SetPaystubProvider(v string)`

SetPaystubProvider sets PaystubProvider field to given value.

### HasPaystubProvider

`func (o *PaystubDetails) HasPaystubProvider() bool`

HasPaystubProvider returns a boolean if a field has been set.

### SetPaystubProviderNil

`func (o *PaystubDetails) SetPaystubProviderNil(b bool)`

 SetPaystubProviderNil sets the value for PaystubProvider to be an explicit nil

### UnsetPaystubProvider
`func (o *PaystubDetails) UnsetPaystubProvider()`

UnsetPaystubProvider ensures that no value is present for PaystubProvider, not even an explicit nil
### GetPayFrequency

`func (o *PaystubDetails) GetPayFrequency() PaystubPayFrequency`

GetPayFrequency returns the PayFrequency field if non-nil, zero value otherwise.

### GetPayFrequencyOk

`func (o *PaystubDetails) GetPayFrequencyOk() (*PaystubPayFrequency, bool)`

GetPayFrequencyOk returns a tuple with the PayFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayFrequency

`func (o *PaystubDetails) SetPayFrequency(v PaystubPayFrequency)`

SetPayFrequency sets PayFrequency field to given value.

### HasPayFrequency

`func (o *PaystubDetails) HasPayFrequency() bool`

HasPayFrequency returns a boolean if a field has been set.

### SetPayFrequencyNil

`func (o *PaystubDetails) SetPayFrequencyNil(b bool)`

 SetPayFrequencyNil sets the value for PayFrequency to be an explicit nil

### UnsetPayFrequency
`func (o *PaystubDetails) UnsetPayFrequency()`

UnsetPayFrequency ensures that no value is present for PayFrequency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


