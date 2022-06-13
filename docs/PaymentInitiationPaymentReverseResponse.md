# PaymentInitiationPaymentReverseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefundId** | **string** | A unique ID identifying the refund | 
**Status** | **string** | The status of the refund.  &#x60;PROCESSING&#x60;: The refund is currently being processed. The refund will automatically exit this state when processing is complete.  &#x60;INITIATED&#x60;: The refund has been successfully initiated.  &#x60;EXECUTED&#x60;: Indicates that the refund has been successfully executed.  &#x60;FAILED&#x60;: The refund has failed to be executed. This error is retryable once the root cause is resolved. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewPaymentInitiationPaymentReverseResponse

`func NewPaymentInitiationPaymentReverseResponse(refundId string, status string, requestId string, ) *PaymentInitiationPaymentReverseResponse`

NewPaymentInitiationPaymentReverseResponse instantiates a new PaymentInitiationPaymentReverseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationPaymentReverseResponseWithDefaults

`func NewPaymentInitiationPaymentReverseResponseWithDefaults() *PaymentInitiationPaymentReverseResponse`

NewPaymentInitiationPaymentReverseResponseWithDefaults instantiates a new PaymentInitiationPaymentReverseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefundId

`func (o *PaymentInitiationPaymentReverseResponse) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *PaymentInitiationPaymentReverseResponse) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *PaymentInitiationPaymentReverseResponse) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.


### GetStatus

`func (o *PaymentInitiationPaymentReverseResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentInitiationPaymentReverseResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentInitiationPaymentReverseResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRequestId

`func (o *PaymentInitiationPaymentReverseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentInitiationPaymentReverseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentInitiationPaymentReverseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


