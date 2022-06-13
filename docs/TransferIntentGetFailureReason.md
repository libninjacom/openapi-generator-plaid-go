# TransferIntentGetFailureReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorType** | Pointer to **string** | A broad categorization of the error. | [optional] 
**ErrorCode** | Pointer to **string** | A code representing the reason for a failed transfer intent (i.e., an API error or the authorization being declined).  For a full listing of bank transfer errors, see [Bank Transfers errors](https://plaid.com/docs/errors/bank-transfers/). | [optional] 
**ErrorMessage** | Pointer to **string** | A human-readable description of the code associated with a failed transfer intent. | [optional] 

## Methods

### NewTransferIntentGetFailureReason

`func NewTransferIntentGetFailureReason() *TransferIntentGetFailureReason`

NewTransferIntentGetFailureReason instantiates a new TransferIntentGetFailureReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferIntentGetFailureReasonWithDefaults

`func NewTransferIntentGetFailureReasonWithDefaults() *TransferIntentGetFailureReason`

NewTransferIntentGetFailureReasonWithDefaults instantiates a new TransferIntentGetFailureReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorType

`func (o *TransferIntentGetFailureReason) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *TransferIntentGetFailureReason) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *TransferIntentGetFailureReason) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *TransferIntentGetFailureReason) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetErrorCode

`func (o *TransferIntentGetFailureReason) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TransferIntentGetFailureReason) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TransferIntentGetFailureReason) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *TransferIntentGetFailureReason) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *TransferIntentGetFailureReason) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TransferIntentGetFailureReason) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TransferIntentGetFailureReason) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TransferIntentGetFailureReason) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


