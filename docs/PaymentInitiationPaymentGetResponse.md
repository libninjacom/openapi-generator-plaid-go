# PaymentInitiationPaymentGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** | The ID of the payment. Like all Plaid identifiers, the &#x60;payment_id&#x60; is case sensitive. | 
**Amount** | [**PaymentAmount**](PaymentAmount.md) |  | 
**Status** | [**PaymentInitiationPaymentStatus**](PaymentInitiationPaymentStatus.md) |  | 
**RecipientId** | **string** | The ID of the recipient | 
**Reference** | **string** | A reference for the payment. | 
**AdjustedReference** | Pointer to **NullableString** | The value of the reference sent to the bank after adjustment to pass bank validation rules. | [optional] 
**LastStatusUpdate** | **time.Time** | The date and time of the last time the &#x60;status&#x60; was updated, in IS0 8601 format | 
**Schedule** | Pointer to [**NullableExternalPaymentScheduleGet**](ExternalPaymentScheduleGet.md) |  | [optional] 
**RefundDetails** | Pointer to [**NullableExternalPaymentRefundDetails**](ExternalPaymentRefundDetails.md) |  | [optional] 
**Bacs** | [**NullableSenderBACSNullable**](SenderBACSNullable.md) |  | 
**Iban** | **NullableString** | The International Bank Account Number (IBAN) for the sender, if specified in the &#x60;/payment_initiation/payment/create&#x60; call. | 
**InitiatedRefunds** | Pointer to [**[]PaymentInitiationRefund**](PaymentInitiationRefund.md) | Initiated refunds associated with the payment. | [optional] 
**WalletId** | Pointer to **NullableString** | The EMI (E-Money Institution) wallet that this payment is associated with, if any. This wallet is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests. | [optional] 
**Scheme** | Pointer to [**NullablePaymentScheme**](PaymentScheme.md) |  | [optional] 
**AdjustedScheme** | Pointer to [**NullablePaymentScheme**](PaymentScheme.md) |  | [optional] 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewPaymentInitiationPaymentGetResponse

`func NewPaymentInitiationPaymentGetResponse(paymentId string, amount PaymentAmount, status PaymentInitiationPaymentStatus, recipientId string, reference string, lastStatusUpdate time.Time, bacs NullableSenderBACSNullable, iban NullableString, requestId string, ) *PaymentInitiationPaymentGetResponse`

NewPaymentInitiationPaymentGetResponse instantiates a new PaymentInitiationPaymentGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationPaymentGetResponseWithDefaults

`func NewPaymentInitiationPaymentGetResponseWithDefaults() *PaymentInitiationPaymentGetResponse`

NewPaymentInitiationPaymentGetResponseWithDefaults instantiates a new PaymentInitiationPaymentGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentInitiationPaymentGetResponse) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentInitiationPaymentGetResponse) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentInitiationPaymentGetResponse) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetAmount

`func (o *PaymentInitiationPaymentGetResponse) GetAmount() PaymentAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentInitiationPaymentGetResponse) GetAmountOk() (*PaymentAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentInitiationPaymentGetResponse) SetAmount(v PaymentAmount)`

SetAmount sets Amount field to given value.


### GetStatus

`func (o *PaymentInitiationPaymentGetResponse) GetStatus() PaymentInitiationPaymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentInitiationPaymentGetResponse) GetStatusOk() (*PaymentInitiationPaymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentInitiationPaymentGetResponse) SetStatus(v PaymentInitiationPaymentStatus)`

SetStatus sets Status field to given value.


### GetRecipientId

`func (o *PaymentInitiationPaymentGetResponse) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *PaymentInitiationPaymentGetResponse) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *PaymentInitiationPaymentGetResponse) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.


### GetReference

`func (o *PaymentInitiationPaymentGetResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentInitiationPaymentGetResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentInitiationPaymentGetResponse) SetReference(v string)`

SetReference sets Reference field to given value.


### GetAdjustedReference

`func (o *PaymentInitiationPaymentGetResponse) GetAdjustedReference() string`

GetAdjustedReference returns the AdjustedReference field if non-nil, zero value otherwise.

### GetAdjustedReferenceOk

`func (o *PaymentInitiationPaymentGetResponse) GetAdjustedReferenceOk() (*string, bool)`

GetAdjustedReferenceOk returns a tuple with the AdjustedReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedReference

`func (o *PaymentInitiationPaymentGetResponse) SetAdjustedReference(v string)`

SetAdjustedReference sets AdjustedReference field to given value.

### HasAdjustedReference

`func (o *PaymentInitiationPaymentGetResponse) HasAdjustedReference() bool`

HasAdjustedReference returns a boolean if a field has been set.

### SetAdjustedReferenceNil

`func (o *PaymentInitiationPaymentGetResponse) SetAdjustedReferenceNil(b bool)`

 SetAdjustedReferenceNil sets the value for AdjustedReference to be an explicit nil

### UnsetAdjustedReference
`func (o *PaymentInitiationPaymentGetResponse) UnsetAdjustedReference()`

UnsetAdjustedReference ensures that no value is present for AdjustedReference, not even an explicit nil
### GetLastStatusUpdate

`func (o *PaymentInitiationPaymentGetResponse) GetLastStatusUpdate() time.Time`

GetLastStatusUpdate returns the LastStatusUpdate field if non-nil, zero value otherwise.

### GetLastStatusUpdateOk

`func (o *PaymentInitiationPaymentGetResponse) GetLastStatusUpdateOk() (*time.Time, bool)`

GetLastStatusUpdateOk returns a tuple with the LastStatusUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusUpdate

`func (o *PaymentInitiationPaymentGetResponse) SetLastStatusUpdate(v time.Time)`

SetLastStatusUpdate sets LastStatusUpdate field to given value.


### GetSchedule

`func (o *PaymentInitiationPaymentGetResponse) GetSchedule() ExternalPaymentScheduleGet`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *PaymentInitiationPaymentGetResponse) GetScheduleOk() (*ExternalPaymentScheduleGet, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *PaymentInitiationPaymentGetResponse) SetSchedule(v ExternalPaymentScheduleGet)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *PaymentInitiationPaymentGetResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *PaymentInitiationPaymentGetResponse) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *PaymentInitiationPaymentGetResponse) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetRefundDetails

`func (o *PaymentInitiationPaymentGetResponse) GetRefundDetails() ExternalPaymentRefundDetails`

GetRefundDetails returns the RefundDetails field if non-nil, zero value otherwise.

### GetRefundDetailsOk

`func (o *PaymentInitiationPaymentGetResponse) GetRefundDetailsOk() (*ExternalPaymentRefundDetails, bool)`

GetRefundDetailsOk returns a tuple with the RefundDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundDetails

`func (o *PaymentInitiationPaymentGetResponse) SetRefundDetails(v ExternalPaymentRefundDetails)`

SetRefundDetails sets RefundDetails field to given value.

### HasRefundDetails

`func (o *PaymentInitiationPaymentGetResponse) HasRefundDetails() bool`

HasRefundDetails returns a boolean if a field has been set.

### SetRefundDetailsNil

`func (o *PaymentInitiationPaymentGetResponse) SetRefundDetailsNil(b bool)`

 SetRefundDetailsNil sets the value for RefundDetails to be an explicit nil

### UnsetRefundDetails
`func (o *PaymentInitiationPaymentGetResponse) UnsetRefundDetails()`

UnsetRefundDetails ensures that no value is present for RefundDetails, not even an explicit nil
### GetBacs

`func (o *PaymentInitiationPaymentGetResponse) GetBacs() SenderBACSNullable`

GetBacs returns the Bacs field if non-nil, zero value otherwise.

### GetBacsOk

`func (o *PaymentInitiationPaymentGetResponse) GetBacsOk() (*SenderBACSNullable, bool)`

GetBacsOk returns a tuple with the Bacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacs

`func (o *PaymentInitiationPaymentGetResponse) SetBacs(v SenderBACSNullable)`

SetBacs sets Bacs field to given value.


### SetBacsNil

`func (o *PaymentInitiationPaymentGetResponse) SetBacsNil(b bool)`

 SetBacsNil sets the value for Bacs to be an explicit nil

### UnsetBacs
`func (o *PaymentInitiationPaymentGetResponse) UnsetBacs()`

UnsetBacs ensures that no value is present for Bacs, not even an explicit nil
### GetIban

`func (o *PaymentInitiationPaymentGetResponse) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentInitiationPaymentGetResponse) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentInitiationPaymentGetResponse) SetIban(v string)`

SetIban sets Iban field to given value.


### SetIbanNil

`func (o *PaymentInitiationPaymentGetResponse) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *PaymentInitiationPaymentGetResponse) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetInitiatedRefunds

`func (o *PaymentInitiationPaymentGetResponse) GetInitiatedRefunds() []PaymentInitiationRefund`

GetInitiatedRefunds returns the InitiatedRefunds field if non-nil, zero value otherwise.

### GetInitiatedRefundsOk

`func (o *PaymentInitiationPaymentGetResponse) GetInitiatedRefundsOk() (*[]PaymentInitiationRefund, bool)`

GetInitiatedRefundsOk returns a tuple with the InitiatedRefunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedRefunds

`func (o *PaymentInitiationPaymentGetResponse) SetInitiatedRefunds(v []PaymentInitiationRefund)`

SetInitiatedRefunds sets InitiatedRefunds field to given value.

### HasInitiatedRefunds

`func (o *PaymentInitiationPaymentGetResponse) HasInitiatedRefunds() bool`

HasInitiatedRefunds returns a boolean if a field has been set.

### GetWalletId

`func (o *PaymentInitiationPaymentGetResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *PaymentInitiationPaymentGetResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *PaymentInitiationPaymentGetResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *PaymentInitiationPaymentGetResponse) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *PaymentInitiationPaymentGetResponse) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *PaymentInitiationPaymentGetResponse) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetScheme

`func (o *PaymentInitiationPaymentGetResponse) GetScheme() PaymentScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PaymentInitiationPaymentGetResponse) GetSchemeOk() (*PaymentScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PaymentInitiationPaymentGetResponse) SetScheme(v PaymentScheme)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PaymentInitiationPaymentGetResponse) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *PaymentInitiationPaymentGetResponse) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *PaymentInitiationPaymentGetResponse) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetAdjustedScheme

`func (o *PaymentInitiationPaymentGetResponse) GetAdjustedScheme() PaymentScheme`

GetAdjustedScheme returns the AdjustedScheme field if non-nil, zero value otherwise.

### GetAdjustedSchemeOk

`func (o *PaymentInitiationPaymentGetResponse) GetAdjustedSchemeOk() (*PaymentScheme, bool)`

GetAdjustedSchemeOk returns a tuple with the AdjustedScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedScheme

`func (o *PaymentInitiationPaymentGetResponse) SetAdjustedScheme(v PaymentScheme)`

SetAdjustedScheme sets AdjustedScheme field to given value.

### HasAdjustedScheme

`func (o *PaymentInitiationPaymentGetResponse) HasAdjustedScheme() bool`

HasAdjustedScheme returns a boolean if a field has been set.

### SetAdjustedSchemeNil

`func (o *PaymentInitiationPaymentGetResponse) SetAdjustedSchemeNil(b bool)`

 SetAdjustedSchemeNil sets the value for AdjustedScheme to be an explicit nil

### UnsetAdjustedScheme
`func (o *PaymentInitiationPaymentGetResponse) UnsetAdjustedScheme()`

UnsetAdjustedScheme ensures that no value is present for AdjustedScheme, not even an explicit nil
### GetRequestId

`func (o *PaymentInitiationPaymentGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentInitiationPaymentGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentInitiationPaymentGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


