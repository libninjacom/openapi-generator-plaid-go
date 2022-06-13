# ExternalPaymentOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestRefundDetails** | Pointer to **NullableBool** | When &#x60;true&#x60;, Plaid will attempt to request refund details from the payee&#39;s financial institution.  Support varies between financial institutions and will not always be available.  If refund details could be retrieved, they will be available in the &#x60;/payment_initiation/payment/get&#x60; response. | [optional] 
**Iban** | Pointer to **NullableString** | The International Bank Account Number (IBAN) for the payer&#39;s account. If provided, the end user will be able to send payments only from the specified bank account. | [optional] 
**Bacs** | Pointer to [**NullablePaymentInitiationOptionalRestrictionBacs**](PaymentInitiationOptionalRestrictionBacs.md) |  | [optional] 
**WalletId** | Pointer to **NullableString** | The EMI (E-Money Institution) wallet that this payment is associated with, if any. This wallet is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests. | [optional] 
**Scheme** | Pointer to [**NullablePaymentScheme**](PaymentScheme.md) |  | [optional] 

## Methods

### NewExternalPaymentOptions

`func NewExternalPaymentOptions() *ExternalPaymentOptions`

NewExternalPaymentOptions instantiates a new ExternalPaymentOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPaymentOptionsWithDefaults

`func NewExternalPaymentOptionsWithDefaults() *ExternalPaymentOptions`

NewExternalPaymentOptionsWithDefaults instantiates a new ExternalPaymentOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestRefundDetails

`func (o *ExternalPaymentOptions) GetRequestRefundDetails() bool`

GetRequestRefundDetails returns the RequestRefundDetails field if non-nil, zero value otherwise.

### GetRequestRefundDetailsOk

`func (o *ExternalPaymentOptions) GetRequestRefundDetailsOk() (*bool, bool)`

GetRequestRefundDetailsOk returns a tuple with the RequestRefundDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRefundDetails

`func (o *ExternalPaymentOptions) SetRequestRefundDetails(v bool)`

SetRequestRefundDetails sets RequestRefundDetails field to given value.

### HasRequestRefundDetails

`func (o *ExternalPaymentOptions) HasRequestRefundDetails() bool`

HasRequestRefundDetails returns a boolean if a field has been set.

### SetRequestRefundDetailsNil

`func (o *ExternalPaymentOptions) SetRequestRefundDetailsNil(b bool)`

 SetRequestRefundDetailsNil sets the value for RequestRefundDetails to be an explicit nil

### UnsetRequestRefundDetails
`func (o *ExternalPaymentOptions) UnsetRequestRefundDetails()`

UnsetRequestRefundDetails ensures that no value is present for RequestRefundDetails, not even an explicit nil
### GetIban

`func (o *ExternalPaymentOptions) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *ExternalPaymentOptions) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *ExternalPaymentOptions) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *ExternalPaymentOptions) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *ExternalPaymentOptions) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *ExternalPaymentOptions) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetBacs

`func (o *ExternalPaymentOptions) GetBacs() PaymentInitiationOptionalRestrictionBacs`

GetBacs returns the Bacs field if non-nil, zero value otherwise.

### GetBacsOk

`func (o *ExternalPaymentOptions) GetBacsOk() (*PaymentInitiationOptionalRestrictionBacs, bool)`

GetBacsOk returns a tuple with the Bacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacs

`func (o *ExternalPaymentOptions) SetBacs(v PaymentInitiationOptionalRestrictionBacs)`

SetBacs sets Bacs field to given value.

### HasBacs

`func (o *ExternalPaymentOptions) HasBacs() bool`

HasBacs returns a boolean if a field has been set.

### SetBacsNil

`func (o *ExternalPaymentOptions) SetBacsNil(b bool)`

 SetBacsNil sets the value for Bacs to be an explicit nil

### UnsetBacs
`func (o *ExternalPaymentOptions) UnsetBacs()`

UnsetBacs ensures that no value is present for Bacs, not even an explicit nil
### GetWalletId

`func (o *ExternalPaymentOptions) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ExternalPaymentOptions) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ExternalPaymentOptions) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *ExternalPaymentOptions) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *ExternalPaymentOptions) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *ExternalPaymentOptions) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetScheme

`func (o *ExternalPaymentOptions) GetScheme() PaymentScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ExternalPaymentOptions) GetSchemeOk() (*PaymentScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ExternalPaymentOptions) SetScheme(v PaymentScheme)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ExternalPaymentOptions) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *ExternalPaymentOptions) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *ExternalPaymentOptions) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


