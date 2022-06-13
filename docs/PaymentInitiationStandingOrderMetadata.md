# PaymentInitiationStandingOrderMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsStandingOrderEndDate** | **bool** | Indicates whether the institution supports closed-ended standing orders by providing an end date. | 
**SupportsStandingOrderNegativeExecutionDays** | **bool** | This is only applicable to &#x60;MONTHLY&#x60; standing orders. Indicates whether the institution supports negative integers (-1 to -5) for setting up a &#x60;MONTHLY&#x60; standing order relative to the end of the month. | 
**ValidStandingOrderIntervals** | [**[]PaymentScheduleInterval**](PaymentScheduleInterval.md) | A list of the valid standing order intervals supported by the institution. | 

## Methods

### NewPaymentInitiationStandingOrderMetadata

`func NewPaymentInitiationStandingOrderMetadata(supportsStandingOrderEndDate bool, supportsStandingOrderNegativeExecutionDays bool, validStandingOrderIntervals []PaymentScheduleInterval, ) *PaymentInitiationStandingOrderMetadata`

NewPaymentInitiationStandingOrderMetadata instantiates a new PaymentInitiationStandingOrderMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationStandingOrderMetadataWithDefaults

`func NewPaymentInitiationStandingOrderMetadataWithDefaults() *PaymentInitiationStandingOrderMetadata`

NewPaymentInitiationStandingOrderMetadataWithDefaults instantiates a new PaymentInitiationStandingOrderMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportsStandingOrderEndDate

`func (o *PaymentInitiationStandingOrderMetadata) GetSupportsStandingOrderEndDate() bool`

GetSupportsStandingOrderEndDate returns the SupportsStandingOrderEndDate field if non-nil, zero value otherwise.

### GetSupportsStandingOrderEndDateOk

`func (o *PaymentInitiationStandingOrderMetadata) GetSupportsStandingOrderEndDateOk() (*bool, bool)`

GetSupportsStandingOrderEndDateOk returns a tuple with the SupportsStandingOrderEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsStandingOrderEndDate

`func (o *PaymentInitiationStandingOrderMetadata) SetSupportsStandingOrderEndDate(v bool)`

SetSupportsStandingOrderEndDate sets SupportsStandingOrderEndDate field to given value.


### GetSupportsStandingOrderNegativeExecutionDays

`func (o *PaymentInitiationStandingOrderMetadata) GetSupportsStandingOrderNegativeExecutionDays() bool`

GetSupportsStandingOrderNegativeExecutionDays returns the SupportsStandingOrderNegativeExecutionDays field if non-nil, zero value otherwise.

### GetSupportsStandingOrderNegativeExecutionDaysOk

`func (o *PaymentInitiationStandingOrderMetadata) GetSupportsStandingOrderNegativeExecutionDaysOk() (*bool, bool)`

GetSupportsStandingOrderNegativeExecutionDaysOk returns a tuple with the SupportsStandingOrderNegativeExecutionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsStandingOrderNegativeExecutionDays

`func (o *PaymentInitiationStandingOrderMetadata) SetSupportsStandingOrderNegativeExecutionDays(v bool)`

SetSupportsStandingOrderNegativeExecutionDays sets SupportsStandingOrderNegativeExecutionDays field to given value.


### GetValidStandingOrderIntervals

`func (o *PaymentInitiationStandingOrderMetadata) GetValidStandingOrderIntervals() []PaymentScheduleInterval`

GetValidStandingOrderIntervals returns the ValidStandingOrderIntervals field if non-nil, zero value otherwise.

### GetValidStandingOrderIntervalsOk

`func (o *PaymentInitiationStandingOrderMetadata) GetValidStandingOrderIntervalsOk() (*[]PaymentScheduleInterval, bool)`

GetValidStandingOrderIntervalsOk returns a tuple with the ValidStandingOrderIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidStandingOrderIntervals

`func (o *PaymentInitiationStandingOrderMetadata) SetValidStandingOrderIntervals(v []PaymentScheduleInterval)`

SetValidStandingOrderIntervals sets ValidStandingOrderIntervals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


