# TransferAuthorizationGuaranteeDecisionRationale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A code representing the reason Plaid declined to guarantee this transfer:  &#x60;RETURN_BANK&#x60;: The risk of a bank-initiated return (for example, an R01/NSF) is too high to guarantee this transfer.  &#x60;RETURN_CUSTOMER&#x60;: The risk of a customer-initiated return (for example, a R10/Unauthorized) is too high to guarantee this transfer.  &#x60;GUARANTEE_LIMIT_REACHED&#x60;: This transfer is low-risk, but Guaranteed ACH has exhausted an internal limit on the number or rate of guarantees that applies to this transfer.  &#x60;RISK_ESTIMATE_UNAVAILABLE&#x60;: A risk estimate is unavailable for this Item. | 
**Description** | **string** | A human-readable description of why the transfer cannot be guaranteed. | 

## Methods

### NewTransferAuthorizationGuaranteeDecisionRationale

`func NewTransferAuthorizationGuaranteeDecisionRationale(code string, description string, ) *TransferAuthorizationGuaranteeDecisionRationale`

NewTransferAuthorizationGuaranteeDecisionRationale instantiates a new TransferAuthorizationGuaranteeDecisionRationale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferAuthorizationGuaranteeDecisionRationaleWithDefaults

`func NewTransferAuthorizationGuaranteeDecisionRationaleWithDefaults() *TransferAuthorizationGuaranteeDecisionRationale`

NewTransferAuthorizationGuaranteeDecisionRationaleWithDefaults instantiates a new TransferAuthorizationGuaranteeDecisionRationale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TransferAuthorizationGuaranteeDecisionRationale) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TransferAuthorizationGuaranteeDecisionRationale) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TransferAuthorizationGuaranteeDecisionRationale) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *TransferAuthorizationGuaranteeDecisionRationale) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransferAuthorizationGuaranteeDecisionRationale) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransferAuthorizationGuaranteeDecisionRationale) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


