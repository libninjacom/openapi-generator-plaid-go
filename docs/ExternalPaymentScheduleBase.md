# ExternalPaymentScheduleBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to [**PaymentScheduleInterval**](PaymentScheduleInterval.md) |  | [optional] 
**IntervalExecutionDay** | Pointer to **int32** | The day of the interval on which to schedule the payment.  If the payment interval is weekly, &#x60;interval_execution_day&#x60; should be an integer from 1 (Monday) to 7 (Sunday).  If the payment interval is monthly, &#x60;interval_execution_day&#x60; should be an integer indicating which day of the month to make the payment on. Integers from 1 to 28 can be used to make a payment on that day of the month. Negative integers from -1 to -5 can be used to make a payment relative to the end of the month. To make a payment on the last day of the month, use -1; to make the payment on the second-to-last day, use -2, and so on. | [optional] 
**StartDate** | Pointer to **string** | A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will begin on the first &#x60;interval_execution_day&#x60; on or after the &#x60;start_date&#x60;.  If the first &#x60;interval_execution_day&#x60; on or after the start date is also the same day that &#x60;/payment_initiation/payment/create&#x60; was called, the bank *may* make the first payment on that day, but it is not guaranteed to do so. | [optional] 
**EndDate** | Pointer to **NullableString** | A date in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Standing order payments will end on the last &#x60;interval_execution_day&#x60; on or before the &#x60;end_date&#x60;. If the only &#x60;interval_execution_day&#x60; between the start date and the end date (inclusive) is also the same day that &#x60;/payment_initiation/payment/create&#x60; was called, the bank *may* make a payment on that day, but it is not guaranteed to do so. | [optional] 
**AdjustedStartDate** | Pointer to **NullableString** | The start date sent to the bank after adjusting for holidays or weekends.  Will be provided in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). If the start date did not require adjustment, this field will be &#x60;null&#x60;. | [optional] 

## Methods

### NewExternalPaymentScheduleBase

`func NewExternalPaymentScheduleBase() *ExternalPaymentScheduleBase`

NewExternalPaymentScheduleBase instantiates a new ExternalPaymentScheduleBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPaymentScheduleBaseWithDefaults

`func NewExternalPaymentScheduleBaseWithDefaults() *ExternalPaymentScheduleBase`

NewExternalPaymentScheduleBaseWithDefaults instantiates a new ExternalPaymentScheduleBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *ExternalPaymentScheduleBase) GetInterval() PaymentScheduleInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ExternalPaymentScheduleBase) GetIntervalOk() (*PaymentScheduleInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ExternalPaymentScheduleBase) SetInterval(v PaymentScheduleInterval)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ExternalPaymentScheduleBase) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetIntervalExecutionDay

`func (o *ExternalPaymentScheduleBase) GetIntervalExecutionDay() int32`

GetIntervalExecutionDay returns the IntervalExecutionDay field if non-nil, zero value otherwise.

### GetIntervalExecutionDayOk

`func (o *ExternalPaymentScheduleBase) GetIntervalExecutionDayOk() (*int32, bool)`

GetIntervalExecutionDayOk returns a tuple with the IntervalExecutionDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalExecutionDay

`func (o *ExternalPaymentScheduleBase) SetIntervalExecutionDay(v int32)`

SetIntervalExecutionDay sets IntervalExecutionDay field to given value.

### HasIntervalExecutionDay

`func (o *ExternalPaymentScheduleBase) HasIntervalExecutionDay() bool`

HasIntervalExecutionDay returns a boolean if a field has been set.

### GetStartDate

`func (o *ExternalPaymentScheduleBase) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ExternalPaymentScheduleBase) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ExternalPaymentScheduleBase) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ExternalPaymentScheduleBase) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ExternalPaymentScheduleBase) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ExternalPaymentScheduleBase) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ExternalPaymentScheduleBase) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ExternalPaymentScheduleBase) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ExternalPaymentScheduleBase) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ExternalPaymentScheduleBase) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetAdjustedStartDate

`func (o *ExternalPaymentScheduleBase) GetAdjustedStartDate() string`

GetAdjustedStartDate returns the AdjustedStartDate field if non-nil, zero value otherwise.

### GetAdjustedStartDateOk

`func (o *ExternalPaymentScheduleBase) GetAdjustedStartDateOk() (*string, bool)`

GetAdjustedStartDateOk returns a tuple with the AdjustedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedStartDate

`func (o *ExternalPaymentScheduleBase) SetAdjustedStartDate(v string)`

SetAdjustedStartDate sets AdjustedStartDate field to given value.

### HasAdjustedStartDate

`func (o *ExternalPaymentScheduleBase) HasAdjustedStartDate() bool`

HasAdjustedStartDate returns a boolean if a field has been set.

### SetAdjustedStartDateNil

`func (o *ExternalPaymentScheduleBase) SetAdjustedStartDateNil(b bool)`

 SetAdjustedStartDateNil sets the value for AdjustedStartDate to be an explicit nil

### UnsetAdjustedStartDate
`func (o *ExternalPaymentScheduleBase) UnsetAdjustedStartDate()`

UnsetAdjustedStartDate ensures that no value is present for AdjustedStartDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


