# PaymentInitiationRecipientGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientId** | **string** | The ID of the recipient. | 
**Name** | **string** | The name of the recipient. | 
**Address** | Pointer to [**NullablePaymentInitiationAddress**](PaymentInitiationAddress.md) |  | [optional] 
**Iban** | Pointer to **NullableString** | The International Bank Account Number (IBAN) for the recipient. | [optional] 
**Bacs** | Pointer to [**NullableRecipientBACSNullable**](RecipientBACSNullable.md) |  | [optional] 
**EmiRecipientId** | Pointer to **NullableString** | The EMI (E-Money Institution) recipient that this recipient is associated with, if any. This EMI recipient is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests. | [optional] 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewPaymentInitiationRecipientGetResponse

`func NewPaymentInitiationRecipientGetResponse(recipientId string, name string, requestId string, ) *PaymentInitiationRecipientGetResponse`

NewPaymentInitiationRecipientGetResponse instantiates a new PaymentInitiationRecipientGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationRecipientGetResponseWithDefaults

`func NewPaymentInitiationRecipientGetResponseWithDefaults() *PaymentInitiationRecipientGetResponse`

NewPaymentInitiationRecipientGetResponseWithDefaults instantiates a new PaymentInitiationRecipientGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientId

`func (o *PaymentInitiationRecipientGetResponse) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *PaymentInitiationRecipientGetResponse) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *PaymentInitiationRecipientGetResponse) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.


### GetName

`func (o *PaymentInitiationRecipientGetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentInitiationRecipientGetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentInitiationRecipientGetResponse) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *PaymentInitiationRecipientGetResponse) GetAddress() PaymentInitiationAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PaymentInitiationRecipientGetResponse) GetAddressOk() (*PaymentInitiationAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PaymentInitiationRecipientGetResponse) SetAddress(v PaymentInitiationAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PaymentInitiationRecipientGetResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *PaymentInitiationRecipientGetResponse) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *PaymentInitiationRecipientGetResponse) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetIban

`func (o *PaymentInitiationRecipientGetResponse) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentInitiationRecipientGetResponse) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentInitiationRecipientGetResponse) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *PaymentInitiationRecipientGetResponse) HasIban() bool`

HasIban returns a boolean if a field has been set.

### SetIbanNil

`func (o *PaymentInitiationRecipientGetResponse) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *PaymentInitiationRecipientGetResponse) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetBacs

`func (o *PaymentInitiationRecipientGetResponse) GetBacs() RecipientBACSNullable`

GetBacs returns the Bacs field if non-nil, zero value otherwise.

### GetBacsOk

`func (o *PaymentInitiationRecipientGetResponse) GetBacsOk() (*RecipientBACSNullable, bool)`

GetBacsOk returns a tuple with the Bacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacs

`func (o *PaymentInitiationRecipientGetResponse) SetBacs(v RecipientBACSNullable)`

SetBacs sets Bacs field to given value.

### HasBacs

`func (o *PaymentInitiationRecipientGetResponse) HasBacs() bool`

HasBacs returns a boolean if a field has been set.

### SetBacsNil

`func (o *PaymentInitiationRecipientGetResponse) SetBacsNil(b bool)`

 SetBacsNil sets the value for Bacs to be an explicit nil

### UnsetBacs
`func (o *PaymentInitiationRecipientGetResponse) UnsetBacs()`

UnsetBacs ensures that no value is present for Bacs, not even an explicit nil
### GetEmiRecipientId

`func (o *PaymentInitiationRecipientGetResponse) GetEmiRecipientId() string`

GetEmiRecipientId returns the EmiRecipientId field if non-nil, zero value otherwise.

### GetEmiRecipientIdOk

`func (o *PaymentInitiationRecipientGetResponse) GetEmiRecipientIdOk() (*string, bool)`

GetEmiRecipientIdOk returns a tuple with the EmiRecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmiRecipientId

`func (o *PaymentInitiationRecipientGetResponse) SetEmiRecipientId(v string)`

SetEmiRecipientId sets EmiRecipientId field to given value.

### HasEmiRecipientId

`func (o *PaymentInitiationRecipientGetResponse) HasEmiRecipientId() bool`

HasEmiRecipientId returns a boolean if a field has been set.

### SetEmiRecipientIdNil

`func (o *PaymentInitiationRecipientGetResponse) SetEmiRecipientIdNil(b bool)`

 SetEmiRecipientIdNil sets the value for EmiRecipientId to be an explicit nil

### UnsetEmiRecipientId
`func (o *PaymentInitiationRecipientGetResponse) UnsetEmiRecipientId()`

UnsetEmiRecipientId ensures that no value is present for EmiRecipientId, not even an explicit nil
### GetRequestId

`func (o *PaymentInitiationRecipientGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentInitiationRecipientGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentInitiationRecipientGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


