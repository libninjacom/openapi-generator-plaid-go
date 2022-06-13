# PaymentInitiationMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsInternationalPayments** | **bool** | Indicates whether the institution supports payments from a different country. | 
**MaximumPaymentAmount** | **map[string]string** | A mapping of currency to maximum payment amount (denominated in the smallest unit of currency) supported by the institution.  Example: &#x60;{\&quot;GBP\&quot;: \&quot;10000\&quot;}&#x60;  | 
**SupportsRefundDetails** | **bool** | Indicates whether the institution supports returning refund details when initiating a payment. | 
**StandingOrderMetadata** | [**NullablePaymentInitiationStandingOrderMetadata**](PaymentInitiationStandingOrderMetadata.md) |  | 

## Methods

### NewPaymentInitiationMetadata

`func NewPaymentInitiationMetadata(supportsInternationalPayments bool, maximumPaymentAmount map[string]string, supportsRefundDetails bool, standingOrderMetadata NullablePaymentInitiationStandingOrderMetadata, ) *PaymentInitiationMetadata`

NewPaymentInitiationMetadata instantiates a new PaymentInitiationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationMetadataWithDefaults

`func NewPaymentInitiationMetadataWithDefaults() *PaymentInitiationMetadata`

NewPaymentInitiationMetadataWithDefaults instantiates a new PaymentInitiationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportsInternationalPayments

`func (o *PaymentInitiationMetadata) GetSupportsInternationalPayments() bool`

GetSupportsInternationalPayments returns the SupportsInternationalPayments field if non-nil, zero value otherwise.

### GetSupportsInternationalPaymentsOk

`func (o *PaymentInitiationMetadata) GetSupportsInternationalPaymentsOk() (*bool, bool)`

GetSupportsInternationalPaymentsOk returns a tuple with the SupportsInternationalPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsInternationalPayments

`func (o *PaymentInitiationMetadata) SetSupportsInternationalPayments(v bool)`

SetSupportsInternationalPayments sets SupportsInternationalPayments field to given value.


### GetMaximumPaymentAmount

`func (o *PaymentInitiationMetadata) GetMaximumPaymentAmount() map[string]string`

GetMaximumPaymentAmount returns the MaximumPaymentAmount field if non-nil, zero value otherwise.

### GetMaximumPaymentAmountOk

`func (o *PaymentInitiationMetadata) GetMaximumPaymentAmountOk() (*map[string]string, bool)`

GetMaximumPaymentAmountOk returns a tuple with the MaximumPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPaymentAmount

`func (o *PaymentInitiationMetadata) SetMaximumPaymentAmount(v map[string]string)`

SetMaximumPaymentAmount sets MaximumPaymentAmount field to given value.


### GetSupportsRefundDetails

`func (o *PaymentInitiationMetadata) GetSupportsRefundDetails() bool`

GetSupportsRefundDetails returns the SupportsRefundDetails field if non-nil, zero value otherwise.

### GetSupportsRefundDetailsOk

`func (o *PaymentInitiationMetadata) GetSupportsRefundDetailsOk() (*bool, bool)`

GetSupportsRefundDetailsOk returns a tuple with the SupportsRefundDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRefundDetails

`func (o *PaymentInitiationMetadata) SetSupportsRefundDetails(v bool)`

SetSupportsRefundDetails sets SupportsRefundDetails field to given value.


### GetStandingOrderMetadata

`func (o *PaymentInitiationMetadata) GetStandingOrderMetadata() PaymentInitiationStandingOrderMetadata`

GetStandingOrderMetadata returns the StandingOrderMetadata field if non-nil, zero value otherwise.

### GetStandingOrderMetadataOk

`func (o *PaymentInitiationMetadata) GetStandingOrderMetadataOk() (*PaymentInitiationStandingOrderMetadata, bool)`

GetStandingOrderMetadataOk returns a tuple with the StandingOrderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandingOrderMetadata

`func (o *PaymentInitiationMetadata) SetStandingOrderMetadata(v PaymentInitiationStandingOrderMetadata)`

SetStandingOrderMetadata sets StandingOrderMetadata field to given value.


### SetStandingOrderMetadataNil

`func (o *PaymentInitiationMetadata) SetStandingOrderMetadataNil(b bool)`

 SetStandingOrderMetadataNil sets the value for StandingOrderMetadata to be an explicit nil

### UnsetStandingOrderMetadata
`func (o *PaymentInitiationMetadata) UnsetStandingOrderMetadata()`

UnsetStandingOrderMetadata ensures that no value is present for StandingOrderMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


