# PaymentStatusUpdateWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;PAYMENT_INITIATION&#x60; | 
**WebhookCode** | **string** | &#x60;PAYMENT_STATUS_UPDATE&#x60; | 
**PaymentId** | **string** | The &#x60;payment_id&#x60; for the payment being updated | 
**NewPaymentStatus** | [**PaymentInitiationPaymentStatus**](PaymentInitiationPaymentStatus.md) |  | 
**OldPaymentStatus** | [**PaymentInitiationPaymentStatus**](PaymentInitiationPaymentStatus.md) |  | 
**OriginalReference** | **NullableString** | The original value of the reference when creating the payment. | 
**AdjustedReference** | Pointer to **NullableString** | The value of the reference sent to the bank after adjustment to pass bank validation rules. | [optional] 
**OriginalStartDate** | **NullableString** | The original value of the &#x60;start_date&#x60; provided during the creation of a standing order. If the payment is not a standing order, this field will be &#x60;null&#x60;. | 
**AdjustedStartDate** | **NullableString** | The start date sent to the bank after adjusting for holidays or weekends.  Will be provided in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). If the start date did not require adjustment, or if the payment is not a standing order, this field will be &#x60;null&#x60;. | 
**Timestamp** | **time.Time** | The timestamp of the update, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. &#x60;\&quot;2017-09-14T14:42:19.350Z\&quot;&#x60; | 
**Error** | Pointer to [**PlaidError**](PlaidError.md) |  | [optional] 

## Methods

### NewPaymentStatusUpdateWebhook

`func NewPaymentStatusUpdateWebhook(webhookType string, webhookCode string, paymentId string, newPaymentStatus PaymentInitiationPaymentStatus, oldPaymentStatus PaymentInitiationPaymentStatus, originalReference NullableString, originalStartDate NullableString, adjustedStartDate NullableString, timestamp time.Time, ) *PaymentStatusUpdateWebhook`

NewPaymentStatusUpdateWebhook instantiates a new PaymentStatusUpdateWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentStatusUpdateWebhookWithDefaults

`func NewPaymentStatusUpdateWebhookWithDefaults() *PaymentStatusUpdateWebhook`

NewPaymentStatusUpdateWebhookWithDefaults instantiates a new PaymentStatusUpdateWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *PaymentStatusUpdateWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *PaymentStatusUpdateWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *PaymentStatusUpdateWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *PaymentStatusUpdateWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *PaymentStatusUpdateWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *PaymentStatusUpdateWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetPaymentId

`func (o *PaymentStatusUpdateWebhook) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentStatusUpdateWebhook) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentStatusUpdateWebhook) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetNewPaymentStatus

`func (o *PaymentStatusUpdateWebhook) GetNewPaymentStatus() PaymentInitiationPaymentStatus`

GetNewPaymentStatus returns the NewPaymentStatus field if non-nil, zero value otherwise.

### GetNewPaymentStatusOk

`func (o *PaymentStatusUpdateWebhook) GetNewPaymentStatusOk() (*PaymentInitiationPaymentStatus, bool)`

GetNewPaymentStatusOk returns a tuple with the NewPaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPaymentStatus

`func (o *PaymentStatusUpdateWebhook) SetNewPaymentStatus(v PaymentInitiationPaymentStatus)`

SetNewPaymentStatus sets NewPaymentStatus field to given value.


### GetOldPaymentStatus

`func (o *PaymentStatusUpdateWebhook) GetOldPaymentStatus() PaymentInitiationPaymentStatus`

GetOldPaymentStatus returns the OldPaymentStatus field if non-nil, zero value otherwise.

### GetOldPaymentStatusOk

`func (o *PaymentStatusUpdateWebhook) GetOldPaymentStatusOk() (*PaymentInitiationPaymentStatus, bool)`

GetOldPaymentStatusOk returns a tuple with the OldPaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPaymentStatus

`func (o *PaymentStatusUpdateWebhook) SetOldPaymentStatus(v PaymentInitiationPaymentStatus)`

SetOldPaymentStatus sets OldPaymentStatus field to given value.


### GetOriginalReference

`func (o *PaymentStatusUpdateWebhook) GetOriginalReference() string`

GetOriginalReference returns the OriginalReference field if non-nil, zero value otherwise.

### GetOriginalReferenceOk

`func (o *PaymentStatusUpdateWebhook) GetOriginalReferenceOk() (*string, bool)`

GetOriginalReferenceOk returns a tuple with the OriginalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReference

`func (o *PaymentStatusUpdateWebhook) SetOriginalReference(v string)`

SetOriginalReference sets OriginalReference field to given value.


### SetOriginalReferenceNil

`func (o *PaymentStatusUpdateWebhook) SetOriginalReferenceNil(b bool)`

 SetOriginalReferenceNil sets the value for OriginalReference to be an explicit nil

### UnsetOriginalReference
`func (o *PaymentStatusUpdateWebhook) UnsetOriginalReference()`

UnsetOriginalReference ensures that no value is present for OriginalReference, not even an explicit nil
### GetAdjustedReference

`func (o *PaymentStatusUpdateWebhook) GetAdjustedReference() string`

GetAdjustedReference returns the AdjustedReference field if non-nil, zero value otherwise.

### GetAdjustedReferenceOk

`func (o *PaymentStatusUpdateWebhook) GetAdjustedReferenceOk() (*string, bool)`

GetAdjustedReferenceOk returns a tuple with the AdjustedReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedReference

`func (o *PaymentStatusUpdateWebhook) SetAdjustedReference(v string)`

SetAdjustedReference sets AdjustedReference field to given value.

### HasAdjustedReference

`func (o *PaymentStatusUpdateWebhook) HasAdjustedReference() bool`

HasAdjustedReference returns a boolean if a field has been set.

### SetAdjustedReferenceNil

`func (o *PaymentStatusUpdateWebhook) SetAdjustedReferenceNil(b bool)`

 SetAdjustedReferenceNil sets the value for AdjustedReference to be an explicit nil

### UnsetAdjustedReference
`func (o *PaymentStatusUpdateWebhook) UnsetAdjustedReference()`

UnsetAdjustedReference ensures that no value is present for AdjustedReference, not even an explicit nil
### GetOriginalStartDate

`func (o *PaymentStatusUpdateWebhook) GetOriginalStartDate() string`

GetOriginalStartDate returns the OriginalStartDate field if non-nil, zero value otherwise.

### GetOriginalStartDateOk

`func (o *PaymentStatusUpdateWebhook) GetOriginalStartDateOk() (*string, bool)`

GetOriginalStartDateOk returns a tuple with the OriginalStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStartDate

`func (o *PaymentStatusUpdateWebhook) SetOriginalStartDate(v string)`

SetOriginalStartDate sets OriginalStartDate field to given value.


### SetOriginalStartDateNil

`func (o *PaymentStatusUpdateWebhook) SetOriginalStartDateNil(b bool)`

 SetOriginalStartDateNil sets the value for OriginalStartDate to be an explicit nil

### UnsetOriginalStartDate
`func (o *PaymentStatusUpdateWebhook) UnsetOriginalStartDate()`

UnsetOriginalStartDate ensures that no value is present for OriginalStartDate, not even an explicit nil
### GetAdjustedStartDate

`func (o *PaymentStatusUpdateWebhook) GetAdjustedStartDate() string`

GetAdjustedStartDate returns the AdjustedStartDate field if non-nil, zero value otherwise.

### GetAdjustedStartDateOk

`func (o *PaymentStatusUpdateWebhook) GetAdjustedStartDateOk() (*string, bool)`

GetAdjustedStartDateOk returns a tuple with the AdjustedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedStartDate

`func (o *PaymentStatusUpdateWebhook) SetAdjustedStartDate(v string)`

SetAdjustedStartDate sets AdjustedStartDate field to given value.


### SetAdjustedStartDateNil

`func (o *PaymentStatusUpdateWebhook) SetAdjustedStartDateNil(b bool)`

 SetAdjustedStartDateNil sets the value for AdjustedStartDate to be an explicit nil

### UnsetAdjustedStartDate
`func (o *PaymentStatusUpdateWebhook) UnsetAdjustedStartDate()`

UnsetAdjustedStartDate ensures that no value is present for AdjustedStartDate, not even an explicit nil
### GetTimestamp

`func (o *PaymentStatusUpdateWebhook) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaymentStatusUpdateWebhook) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaymentStatusUpdateWebhook) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetError

`func (o *PaymentStatusUpdateWebhook) GetError() PlaidError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PaymentStatusUpdateWebhook) GetErrorOk() (*PlaidError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PaymentStatusUpdateWebhook) SetError(v PlaidError)`

SetError sets Error field to given value.

### HasError

`func (o *PaymentStatusUpdateWebhook) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


