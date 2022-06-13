# PaymentInitiationPayment

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

## Methods

### NewPaymentInitiationPayment

`func NewPaymentInitiationPayment(paymentId string, amount PaymentAmount, status PaymentInitiationPaymentStatus, recipientId string, reference string, lastStatusUpdate time.Time, bacs NullableSenderBACSNullable, iban NullableString, ) *PaymentInitiationPayment`

NewPaymentInitiationPayment instantiates a new PaymentInitiationPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationPaymentWithDefaults

`func NewPaymentInitiationPaymentWithDefaults() *PaymentInitiationPayment`

NewPaymentInitiationPaymentWithDefaults instantiates a new PaymentInitiationPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentInitiationPayment) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentInitiationPayment) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentInitiationPayment) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetAmount

`func (o *PaymentInitiationPayment) GetAmount() PaymentAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentInitiationPayment) GetAmountOk() (*PaymentAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentInitiationPayment) SetAmount(v PaymentAmount)`

SetAmount sets Amount field to given value.


### GetStatus

`func (o *PaymentInitiationPayment) GetStatus() PaymentInitiationPaymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentInitiationPayment) GetStatusOk() (*PaymentInitiationPaymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentInitiationPayment) SetStatus(v PaymentInitiationPaymentStatus)`

SetStatus sets Status field to given value.


### GetRecipientId

`func (o *PaymentInitiationPayment) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *PaymentInitiationPayment) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *PaymentInitiationPayment) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.


### GetReference

`func (o *PaymentInitiationPayment) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentInitiationPayment) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentInitiationPayment) SetReference(v string)`

SetReference sets Reference field to given value.


### GetAdjustedReference

`func (o *PaymentInitiationPayment) GetAdjustedReference() string`

GetAdjustedReference returns the AdjustedReference field if non-nil, zero value otherwise.

### GetAdjustedReferenceOk

`func (o *PaymentInitiationPayment) GetAdjustedReferenceOk() (*string, bool)`

GetAdjustedReferenceOk returns a tuple with the AdjustedReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedReference

`func (o *PaymentInitiationPayment) SetAdjustedReference(v string)`

SetAdjustedReference sets AdjustedReference field to given value.

### HasAdjustedReference

`func (o *PaymentInitiationPayment) HasAdjustedReference() bool`

HasAdjustedReference returns a boolean if a field has been set.

### SetAdjustedReferenceNil

`func (o *PaymentInitiationPayment) SetAdjustedReferenceNil(b bool)`

 SetAdjustedReferenceNil sets the value for AdjustedReference to be an explicit nil

### UnsetAdjustedReference
`func (o *PaymentInitiationPayment) UnsetAdjustedReference()`

UnsetAdjustedReference ensures that no value is present for AdjustedReference, not even an explicit nil
### GetLastStatusUpdate

`func (o *PaymentInitiationPayment) GetLastStatusUpdate() time.Time`

GetLastStatusUpdate returns the LastStatusUpdate field if non-nil, zero value otherwise.

### GetLastStatusUpdateOk

`func (o *PaymentInitiationPayment) GetLastStatusUpdateOk() (*time.Time, bool)`

GetLastStatusUpdateOk returns a tuple with the LastStatusUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusUpdate

`func (o *PaymentInitiationPayment) SetLastStatusUpdate(v time.Time)`

SetLastStatusUpdate sets LastStatusUpdate field to given value.


### GetSchedule

`func (o *PaymentInitiationPayment) GetSchedule() ExternalPaymentScheduleGet`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *PaymentInitiationPayment) GetScheduleOk() (*ExternalPaymentScheduleGet, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *PaymentInitiationPayment) SetSchedule(v ExternalPaymentScheduleGet)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *PaymentInitiationPayment) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *PaymentInitiationPayment) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *PaymentInitiationPayment) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetRefundDetails

`func (o *PaymentInitiationPayment) GetRefundDetails() ExternalPaymentRefundDetails`

GetRefundDetails returns the RefundDetails field if non-nil, zero value otherwise.

### GetRefundDetailsOk

`func (o *PaymentInitiationPayment) GetRefundDetailsOk() (*ExternalPaymentRefundDetails, bool)`

GetRefundDetailsOk returns a tuple with the RefundDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundDetails

`func (o *PaymentInitiationPayment) SetRefundDetails(v ExternalPaymentRefundDetails)`

SetRefundDetails sets RefundDetails field to given value.

### HasRefundDetails

`func (o *PaymentInitiationPayment) HasRefundDetails() bool`

HasRefundDetails returns a boolean if a field has been set.

### SetRefundDetailsNil

`func (o *PaymentInitiationPayment) SetRefundDetailsNil(b bool)`

 SetRefundDetailsNil sets the value for RefundDetails to be an explicit nil

### UnsetRefundDetails
`func (o *PaymentInitiationPayment) UnsetRefundDetails()`

UnsetRefundDetails ensures that no value is present for RefundDetails, not even an explicit nil
### GetBacs

`func (o *PaymentInitiationPayment) GetBacs() SenderBACSNullable`

GetBacs returns the Bacs field if non-nil, zero value otherwise.

### GetBacsOk

`func (o *PaymentInitiationPayment) GetBacsOk() (*SenderBACSNullable, bool)`

GetBacsOk returns a tuple with the Bacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacs

`func (o *PaymentInitiationPayment) SetBacs(v SenderBACSNullable)`

SetBacs sets Bacs field to given value.


### SetBacsNil

`func (o *PaymentInitiationPayment) SetBacsNil(b bool)`

 SetBacsNil sets the value for Bacs to be an explicit nil

### UnsetBacs
`func (o *PaymentInitiationPayment) UnsetBacs()`

UnsetBacs ensures that no value is present for Bacs, not even an explicit nil
### GetIban

`func (o *PaymentInitiationPayment) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentInitiationPayment) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentInitiationPayment) SetIban(v string)`

SetIban sets Iban field to given value.


### SetIbanNil

`func (o *PaymentInitiationPayment) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *PaymentInitiationPayment) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetInitiatedRefunds

`func (o *PaymentInitiationPayment) GetInitiatedRefunds() []PaymentInitiationRefund`

GetInitiatedRefunds returns the InitiatedRefunds field if non-nil, zero value otherwise.

### GetInitiatedRefundsOk

`func (o *PaymentInitiationPayment) GetInitiatedRefundsOk() (*[]PaymentInitiationRefund, bool)`

GetInitiatedRefundsOk returns a tuple with the InitiatedRefunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedRefunds

`func (o *PaymentInitiationPayment) SetInitiatedRefunds(v []PaymentInitiationRefund)`

SetInitiatedRefunds sets InitiatedRefunds field to given value.

### HasInitiatedRefunds

`func (o *PaymentInitiationPayment) HasInitiatedRefunds() bool`

HasInitiatedRefunds returns a boolean if a field has been set.

### GetWalletId

`func (o *PaymentInitiationPayment) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *PaymentInitiationPayment) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *PaymentInitiationPayment) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *PaymentInitiationPayment) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *PaymentInitiationPayment) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *PaymentInitiationPayment) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetScheme

`func (o *PaymentInitiationPayment) GetScheme() PaymentScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PaymentInitiationPayment) GetSchemeOk() (*PaymentScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PaymentInitiationPayment) SetScheme(v PaymentScheme)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PaymentInitiationPayment) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *PaymentInitiationPayment) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *PaymentInitiationPayment) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetAdjustedScheme

`func (o *PaymentInitiationPayment) GetAdjustedScheme() PaymentScheme`

GetAdjustedScheme returns the AdjustedScheme field if non-nil, zero value otherwise.

### GetAdjustedSchemeOk

`func (o *PaymentInitiationPayment) GetAdjustedSchemeOk() (*PaymentScheme, bool)`

GetAdjustedSchemeOk returns a tuple with the AdjustedScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedScheme

`func (o *PaymentInitiationPayment) SetAdjustedScheme(v PaymentScheme)`

SetAdjustedScheme sets AdjustedScheme field to given value.

### HasAdjustedScheme

`func (o *PaymentInitiationPayment) HasAdjustedScheme() bool`

HasAdjustedScheme returns a boolean if a field has been set.

### SetAdjustedSchemeNil

`func (o *PaymentInitiationPayment) SetAdjustedSchemeNil(b bool)`

 SetAdjustedSchemeNil sets the value for AdjustedScheme to be an explicit nil

### UnsetAdjustedScheme
`func (o *PaymentInitiationPayment) UnsetAdjustedScheme()`

UnsetAdjustedScheme ensures that no value is present for AdjustedScheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


