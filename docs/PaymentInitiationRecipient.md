# PaymentInitiationRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientId** | **string** | The ID of the recipient. | 
**Name** | **string** | The name of the recipient. | 
**Address** | Pointer to [**NullablePaymentInitiationAddress**](PaymentInitiationAddress.md) |  | [optional] 
**Iban** | Pointer to **NullableString** | The International Bank Account Number (IBAN) for the recipient. | [optional] 
**Bacs** | Pointer to [**NullableRecipientBACSNullable**](RecipientBACSNullable.md) |  | [optional] 
**EmiRecipientId** | Pointer to **NullableString** | The EMI (E-Money Institution) recipient that this recipient is associated with, if any. This EMI recipient is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests. | [optional] 

## Methods

### NewPaymentInitiationRecipient

`func NewPaymentInitiationRecipient(recipientId string, name string, ) *PaymentInitiationRecipient`

NewPaymentInitiationRecipient instantiates a new PaymentInitiationRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationRecipientWithDefaults

`func NewPaymentInitiationRecipientWithDefaults() *PaymentInitiationRecipient`

NewPaymentInitiationRecipientWithDefaults instantiates a new PaymentInitiationRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientId

`func (o *PaymentInitiationRecipient) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *PaymentInitiationRecipient) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *PaymentInitiationRecipient) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.


### GetName

`func (o *PaymentInitiationRecipient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentInitiationRecipient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentInitiationRecipient) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *PaymentInitiationRecipient) GetAddress() PaymentInitiationAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PaymentInitiationRecipient) GetAddressOk() (*PaymentInitiationAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PaymentInitiationRecipient) SetAddress(v PaymentInitiationAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PaymentInitiationRecipient) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *PaymentInitiationRecipient) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *PaymentInitiationRecipient) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetIban

`func (o *PaymentInitiationRecipient) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentInitiationRecipient) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentInitiationRecipient) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *PaymentInitiationRecipient) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *PaymentInitiationRecipient) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *PaymentInitiationRecipient) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetBacs

`func (o *PaymentInitiationRecipient) GetBacs() RecipientBACSNullable`

GetBacs returns the Bacs field if non-nil, zero value otherwise.

### GetBacsOk

`func (o *PaymentInitiationRecipient) GetBacsOk() (*RecipientBACSNullable, bool)`

GetBacsOk returns a tuple with the Bacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacs

`func (o *PaymentInitiationRecipient) SetBacs(v RecipientBACSNullable)`

SetBacs sets Bacs field to given value.

### HasBacs

`func (o *PaymentInitiationRecipient) HasBacs() bool`

HasBacs returns a boolean if a field has been set.

### SetBacsNil

`func (o *PaymentInitiationRecipient) SetBacsNil(b bool)`

 SetBacsNil sets the value for Bacs to be an explicit nil

### UnsetBacs
`func (o *PaymentInitiationRecipient) UnsetBacs()`

UnsetBacs ensures that no value is present for Bacs, not even an explicit nil
### GetEmiRecipientId

`func (o *PaymentInitiationRecipient) GetEmiRecipientId() string`

GetEmiRecipientId returns the EmiRecipientId field if non-nil, zero value otherwise.

### GetEmiRecipientIdOk

`func (o *PaymentInitiationRecipient) GetEmiRecipientIdOk() (*string, bool)`

GetEmiRecipientIdOk returns a tuple with the EmiRecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmiRecipientId

`func (o *PaymentInitiationRecipient) SetEmiRecipientId(v string)`

SetEmiRecipientId sets EmiRecipientId field to given value.

### HasEmiRecipientId

`func (o *PaymentInitiationRecipient) HasEmiRecipientId() bool`

HasEmiRecipientId returns a boolean if a field has been set.

### SetEmiRecipientIdNil

`func (o *PaymentInitiationRecipient) SetEmiRecipientIdNil(b bool)`

 SetEmiRecipientIdNil sets the value for EmiRecipientId to be an explicit nil

### UnsetEmiRecipientId
`func (o *PaymentInitiationRecipient) UnsetEmiRecipientId()`

UnsetEmiRecipientId ensures that no value is present for EmiRecipientId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


