# ExternalPaymentScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | [**PaymentScheduleInterval**](PaymentScheduleInterval.md) |  | 
**IntervalExecutionDay** | **int32** | The day of the interval on which to schedule the payment.  If the payment interval is weekly, &#x60;interval_execution_day&#x60; should be an integer from 1 (Monday) to 7 (Sunday).  If the payment interval is monthly, &#x60;interval_execution_day&#x60; should be an integer indicating which day of the month to make the payment on. Integers from 1 to 28 can be used to make a payment on that day of the month. Negative integers from -1 to -5 can be used to make a payment relative to the end of the month. To make a payment on the last day of the month, use -1; to make the payment on the second-to-last day, use -2, and so on. | 
**StartDate** | **string** | A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will begin on the first &#x60;interval_execution_day&#x60; on or after the &#x60;start_date&#x60;.  If the first &#x60;interval_execution_day&#x60; on or after the start date is also the same day that &#x60;/payment_initiation/payment/create&#x60; was called, the bank *may* make the first payment on that day, but it is not guaranteed to do so. | 
**EndDate** | Pointer to **NullableString** | A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will end on the last &#x60;interval_execution_day&#x60; on or before the &#x60;end_date&#x60;. If the only &#x60;interval_execution_day&#x60; between the start date and the end date (inclusive) is also the same day that &#x60;/payment_initiation/payment/create&#x60; was called, the bank *may* make a payment on that day, but it is not guaranteed to do so. | [optional] 
**AdjustedStartDate** | Pointer to **NullableString** | The start date sent to the bank after adjusting for holidays or weekends.  Will be provided in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). If the start date did not require adjustment, this field will be &#x60;null&#x60;. | [optional] 

## Methods

### NewExternalPaymentScheduleRequest

`func NewExternalPaymentScheduleRequest(interval PaymentScheduleInterval, intervalExecutionDay int32, startDate string, ) *ExternalPaymentScheduleRequest`

NewExternalPaymentScheduleRequest instantiates a new ExternalPaymentScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPaymentScheduleRequestWithDefaults

`func NewExternalPaymentScheduleRequestWithDefaults() *ExternalPaymentScheduleRequest`

NewExternalPaymentScheduleRequestWithDefaults instantiates a new ExternalPaymentScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *ExternalPaymentScheduleRequest) GetInterval() PaymentScheduleInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ExternalPaymentScheduleRequest) GetIntervalOk() (*PaymentScheduleInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ExternalPaymentScheduleRequest) SetInterval(v PaymentScheduleInterval)`

SetInterval sets Interval field to given value.


### GetIntervalExecutionDay

`func (o *ExternalPaymentScheduleRequest) GetIntervalExecutionDay() int32`

GetIntervalExecutionDay returns the IntervalExecutionDay field if non-nil, zero value otherwise.

### GetIntervalExecutionDayOk

`func (o *ExternalPaymentScheduleRequest) GetIntervalExecutionDayOk() (*int32, bool)`

GetIntervalExecutionDayOk returns a tuple with the IntervalExecutionDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalExecutionDay

`func (o *ExternalPaymentScheduleRequest) SetIntervalExecutionDay(v int32)`

SetIntervalExecutionDay sets IntervalExecutionDay field to given value.


### GetStartDate

`func (o *ExternalPaymentScheduleRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ExternalPaymentScheduleRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ExternalPaymentScheduleRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *ExternalPaymentScheduleRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ExternalPaymentScheduleRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ExternalPaymentScheduleRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ExternalPaymentScheduleRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ExternalPaymentScheduleRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ExternalPaymentScheduleRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetAdjustedStartDate

`func (o *ExternalPaymentScheduleRequest) GetAdjustedStartDate() string`

GetAdjustedStartDate returns the AdjustedStartDate field if non-nil, zero value otherwise.

### GetAdjustedStartDateOk

`func (o *ExternalPaymentScheduleRequest) GetAdjustedStartDateOk() (*string, bool)`

GetAdjustedStartDateOk returns a tuple with the AdjustedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedStartDate

`func (o *ExternalPaymentScheduleRequest) SetAdjustedStartDate(v string)`

SetAdjustedStartDate sets AdjustedStartDate field to given value.

### HasAdjustedStartDate

`func (o *ExternalPaymentScheduleRequest) HasAdjustedStartDate() bool`

HasAdjustedStartDate returns a boolean if a field has been set.

### SetAdjustedStartDateNil

`func (o *ExternalPaymentScheduleRequest) SetAdjustedStartDateNil(b bool)`

 SetAdjustedStartDateNil sets the value for AdjustedStartDate to be an explicit nil

### UnsetAdjustedStartDate
`func (o *ExternalPaymentScheduleRequest) UnsetAdjustedStartDate()`

UnsetAdjustedStartDate ensures that no value is present for AdjustedStartDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


