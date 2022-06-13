# PaymentInitiationRefund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefundId** | **string** | The ID of the refund. Like all Plaid identifiers, the &#x60;refund_id&#x60; is case sensitive. | 
**Amount** | [**PaymentAmount**](PaymentAmount.md) |  | 
**Status** | **string** | The status of the refund.  &#x60;PROCESSING&#x60;: The refund is currently being processed. The refund will automatically exit this state when processing is complete.  &#x60;INITIATED&#x60;: The refund has been successfully initiated.  &#x60;EXECUTED&#x60;: Indicates that the refund has been successfully executed.  &#x60;FAILED&#x60;: The refund has failed to be executed. This error is retryable once the root cause is resolved. | 
**LastStatusUpdate** | **time.Time** | The date and time of the last time the &#x60;status&#x60; was updated, in IS0 8601 format | 

## Methods

### NewPaymentInitiationRefund

`func NewPaymentInitiationRefund(refundId string, amount PaymentAmount, status string, lastStatusUpdate time.Time, ) *PaymentInitiationRefund`

NewPaymentInitiationRefund instantiates a new PaymentInitiationRefund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationRefundWithDefaults

`func NewPaymentInitiationRefundWithDefaults() *PaymentInitiationRefund`

NewPaymentInitiationRefundWithDefaults instantiates a new PaymentInitiationRefund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefundId

`func (o *PaymentInitiationRefund) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *PaymentInitiationRefund) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *PaymentInitiationRefund) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.


### GetAmount

`func (o *PaymentInitiationRefund) GetAmount() PaymentAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentInitiationRefund) GetAmountOk() (*PaymentAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentInitiationRefund) SetAmount(v PaymentAmount)`

SetAmount sets Amount field to given value.


### GetStatus

`func (o *PaymentInitiationRefund) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentInitiationRefund) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentInitiationRefund) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastStatusUpdate

`func (o *PaymentInitiationRefund) GetLastStatusUpdate() time.Time`

GetLastStatusUpdate returns the LastStatusUpdate field if non-nil, zero value otherwise.

### GetLastStatusUpdateOk

`func (o *PaymentInitiationRefund) GetLastStatusUpdateOk() (*time.Time, bool)`

GetLastStatusUpdateOk returns a tuple with the LastStatusUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusUpdate

`func (o *PaymentInitiationRefund) SetLastStatusUpdate(v time.Time)`

SetLastStatusUpdate sets LastStatusUpdate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


