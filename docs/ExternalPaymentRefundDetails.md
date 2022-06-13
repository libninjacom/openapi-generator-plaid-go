# ExternalPaymentRefundDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the account holder. | 
**Iban** | **NullableString** | The International Bank Account Number (IBAN) for the account. | 
**Bacs** | [**NullableRecipientBACSNullable**](RecipientBACSNullable.md) |  | 

## Methods

### NewExternalPaymentRefundDetails

`func NewExternalPaymentRefundDetails(name string, iban NullableString, bacs NullableRecipientBACSNullable, ) *ExternalPaymentRefundDetails`

NewExternalPaymentRefundDetails instantiates a new ExternalPaymentRefundDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPaymentRefundDetailsWithDefaults

`func NewExternalPaymentRefundDetailsWithDefaults() *ExternalPaymentRefundDetails`

NewExternalPaymentRefundDetailsWithDefaults instantiates a new ExternalPaymentRefundDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalPaymentRefundDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalPaymentRefundDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalPaymentRefundDetails) SetName(v string)`

SetName sets Name field to given value.


### GetIban

`func (o *ExternalPaymentRefundDetails) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *ExternalPaymentRefundDetails) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *ExternalPaymentRefundDetails) SetIban(v string)`

SetIban sets Iban field to given value.


### SetIbanNil

`func (o *ExternalPaymentRefundDetails) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *ExternalPaymentRefundDetails) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetBacs

`func (o *ExternalPaymentRefundDetails) GetBacs() RecipientBACSNullable`

GetBacs returns the Bacs field if non-nil, zero value otherwise.

### GetBacsOk

`func (o *ExternalPaymentRefundDetails) GetBacsOk() (*RecipientBACSNullable, bool)`

GetBacsOk returns a tuple with the Bacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacs

`func (o *ExternalPaymentRefundDetails) SetBacs(v RecipientBACSNullable)`

SetBacs sets Bacs field to given value.


### SetBacsNil

`func (o *ExternalPaymentRefundDetails) SetBacsNil(b bool)`

 SetBacsNil sets the value for Bacs to be an explicit nil

### UnsetBacs
`func (o *ExternalPaymentRefundDetails) UnsetBacs()`

UnsetBacs ensures that no value is present for Bacs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


