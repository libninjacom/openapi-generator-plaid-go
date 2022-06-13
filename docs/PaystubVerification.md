# PaystubVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationStatus** | [**NullablePaystubVerificationStatus**](PaystubVerificationStatus.md) |  | 
**VerificationAttributes** | [**[]VerificationAttribute**](VerificationAttribute.md) |  | 

## Methods

### NewPaystubVerification

`func NewPaystubVerification(verificationStatus NullablePaystubVerificationStatus, verificationAttributes []VerificationAttribute, ) *PaystubVerification`

NewPaystubVerification instantiates a new PaystubVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaystubVerificationWithDefaults

`func NewPaystubVerificationWithDefaults() *PaystubVerification`

NewPaystubVerificationWithDefaults instantiates a new PaystubVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationStatus

`func (o *PaystubVerification) GetVerificationStatus() PaystubVerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *PaystubVerification) GetVerificationStatusOk() (*PaystubVerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *PaystubVerification) SetVerificationStatus(v PaystubVerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.


### SetVerificationStatusNil

`func (o *PaystubVerification) SetVerificationStatusNil(b bool)`

 SetVerificationStatusNil sets the value for VerificationStatus to be an explicit nil

### UnsetVerificationStatus
`func (o *PaystubVerification) UnsetVerificationStatus()`

UnsetVerificationStatus ensures that no value is present for VerificationStatus, not even an explicit nil
### GetVerificationAttributes

`func (o *PaystubVerification) GetVerificationAttributes() []VerificationAttribute`

GetVerificationAttributes returns the VerificationAttributes field if non-nil, zero value otherwise.

### GetVerificationAttributesOk

`func (o *PaystubVerification) GetVerificationAttributesOk() (*[]VerificationAttribute, bool)`

GetVerificationAttributesOk returns a tuple with the VerificationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationAttributes

`func (o *PaystubVerification) SetVerificationAttributes(v []VerificationAttribute)`

SetVerificationAttributes sets VerificationAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


