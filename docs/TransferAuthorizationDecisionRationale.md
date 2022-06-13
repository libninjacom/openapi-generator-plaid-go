# TransferAuthorizationDecisionRationale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A code representing the rationale for permitting or declining the proposed transfer. Possible values are:  &#x60;NSF&#x60; – Transaction likely to result in a return due to insufficient funds.  &#x60;RISK&#x60; - Transaction is high-risk.  &#x60;MANUALLY_VERIFIED_ITEM&#x60; – Item created via same-day micro deposits, limited information available. Plaid can only offer &#x60;permitted&#x60; as a transaction decision.  &#x60;LOGIN_REQUIRED&#x60; – Unable to collect the account information required for an authorization decision due to Item staleness. Can be rectified using Link update mode.  &#x60;ERROR&#x60; – Unable to collect the account information required for an authorization decision due to an error. | 
**Description** | **string** | A human-readable description of the code associated with a permitted transfer or transfer decline. | 

## Methods

### NewTransferAuthorizationDecisionRationale

`func NewTransferAuthorizationDecisionRationale(code string, description string, ) *TransferAuthorizationDecisionRationale`

NewTransferAuthorizationDecisionRationale instantiates a new TransferAuthorizationDecisionRationale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferAuthorizationDecisionRationaleWithDefaults

`func NewTransferAuthorizationDecisionRationaleWithDefaults() *TransferAuthorizationDecisionRationale`

NewTransferAuthorizationDecisionRationaleWithDefaults instantiates a new TransferAuthorizationDecisionRationale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TransferAuthorizationDecisionRationale) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TransferAuthorizationDecisionRationale) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TransferAuthorizationDecisionRationale) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *TransferAuthorizationDecisionRationale) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransferAuthorizationDecisionRationale) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransferAuthorizationDecisionRationale) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


