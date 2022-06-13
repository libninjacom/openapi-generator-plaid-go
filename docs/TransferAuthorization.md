# TransferAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Plaid’s unique identifier for a transfer authorization. | 
**Created** | **time.Time** | The datetime representing when the authorization was created, in the format &#x60;2006-01-02T15:04:05Z&#x60;. | 
**Decision** | **string** |  A decision regarding the proposed transfer.  &#x60;approved&#x60; – The proposed transfer has received the end user&#39;s consent and has been approved for processing. Plaid has also reviewed the proposed transfer and has approved it for processing.   &#x60;permitted&#x60; – Plaid was unable to fetch the information required to approve or decline the proposed transfer. You may proceed with the transfer, but further review is recommended. Plaid is awaiting further instructions from the client.  &#x60;declined&#x60; – Plaid reviewed the proposed transfer and declined processing. Refer to the &#x60;code&#x60; field in the &#x60;decision_rationale&#x60; object for details. | 
**DecisionRationale** | [**NullableTransferAuthorizationDecisionRationale**](TransferAuthorizationDecisionRationale.md) |  | 
**GuaranteeDecision** | [**NullableTransferAuthorizationGuaranteeDecision**](TransferAuthorizationGuaranteeDecision.md) |  | 
**GuaranteeDecisionRationale** | [**NullableTransferAuthorizationGuaranteeDecisionRationale**](TransferAuthorizationGuaranteeDecisionRationale.md) |  | 
**ProposedTransfer** | [**TransferAuthorizationProposedTransfer**](TransferAuthorizationProposedTransfer.md) |  | 

## Methods

### NewTransferAuthorization

`func NewTransferAuthorization(id string, created time.Time, decision string, decisionRationale NullableTransferAuthorizationDecisionRationale, guaranteeDecision NullableTransferAuthorizationGuaranteeDecision, guaranteeDecisionRationale NullableTransferAuthorizationGuaranteeDecisionRationale, proposedTransfer TransferAuthorizationProposedTransfer, ) *TransferAuthorization`

NewTransferAuthorization instantiates a new TransferAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferAuthorizationWithDefaults

`func NewTransferAuthorizationWithDefaults() *TransferAuthorization`

NewTransferAuthorizationWithDefaults instantiates a new TransferAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransferAuthorization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferAuthorization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferAuthorization) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *TransferAuthorization) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TransferAuthorization) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TransferAuthorization) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDecision

`func (o *TransferAuthorization) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *TransferAuthorization) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *TransferAuthorization) SetDecision(v string)`

SetDecision sets Decision field to given value.


### GetDecisionRationale

`func (o *TransferAuthorization) GetDecisionRationale() TransferAuthorizationDecisionRationale`

GetDecisionRationale returns the DecisionRationale field if non-nil, zero value otherwise.

### GetDecisionRationaleOk

`func (o *TransferAuthorization) GetDecisionRationaleOk() (*TransferAuthorizationDecisionRationale, bool)`

GetDecisionRationaleOk returns a tuple with the DecisionRationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionRationale

`func (o *TransferAuthorization) SetDecisionRationale(v TransferAuthorizationDecisionRationale)`

SetDecisionRationale sets DecisionRationale field to given value.


### SetDecisionRationaleNil

`func (o *TransferAuthorization) SetDecisionRationaleNil(b bool)`

 SetDecisionRationaleNil sets the value for DecisionRationale to be an explicit nil

### UnsetDecisionRationale
`func (o *TransferAuthorization) UnsetDecisionRationale()`

UnsetDecisionRationale ensures that no value is present for DecisionRationale, not even an explicit nil
### GetGuaranteeDecision

`func (o *TransferAuthorization) GetGuaranteeDecision() TransferAuthorizationGuaranteeDecision`

GetGuaranteeDecision returns the GuaranteeDecision field if non-nil, zero value otherwise.

### GetGuaranteeDecisionOk

`func (o *TransferAuthorization) GetGuaranteeDecisionOk() (*TransferAuthorizationGuaranteeDecision, bool)`

GetGuaranteeDecisionOk returns a tuple with the GuaranteeDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeDecision

`func (o *TransferAuthorization) SetGuaranteeDecision(v TransferAuthorizationGuaranteeDecision)`

SetGuaranteeDecision sets GuaranteeDecision field to given value.


### SetGuaranteeDecisionNil

`func (o *TransferAuthorization) SetGuaranteeDecisionNil(b bool)`

 SetGuaranteeDecisionNil sets the value for GuaranteeDecision to be an explicit nil

### UnsetGuaranteeDecision
`func (o *TransferAuthorization) UnsetGuaranteeDecision()`

UnsetGuaranteeDecision ensures that no value is present for GuaranteeDecision, not even an explicit nil
### GetGuaranteeDecisionRationale

`func (o *TransferAuthorization) GetGuaranteeDecisionRationale() TransferAuthorizationGuaranteeDecisionRationale`

GetGuaranteeDecisionRationale returns the GuaranteeDecisionRationale field if non-nil, zero value otherwise.

### GetGuaranteeDecisionRationaleOk

`func (o *TransferAuthorization) GetGuaranteeDecisionRationaleOk() (*TransferAuthorizationGuaranteeDecisionRationale, bool)`

GetGuaranteeDecisionRationaleOk returns a tuple with the GuaranteeDecisionRationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeDecisionRationale

`func (o *TransferAuthorization) SetGuaranteeDecisionRationale(v TransferAuthorizationGuaranteeDecisionRationale)`

SetGuaranteeDecisionRationale sets GuaranteeDecisionRationale field to given value.


### SetGuaranteeDecisionRationaleNil

`func (o *TransferAuthorization) SetGuaranteeDecisionRationaleNil(b bool)`

 SetGuaranteeDecisionRationaleNil sets the value for GuaranteeDecisionRationale to be an explicit nil

### UnsetGuaranteeDecisionRationale
`func (o *TransferAuthorization) UnsetGuaranteeDecisionRationale()`

UnsetGuaranteeDecisionRationale ensures that no value is present for GuaranteeDecisionRationale, not even an explicit nil
### GetProposedTransfer

`func (o *TransferAuthorization) GetProposedTransfer() TransferAuthorizationProposedTransfer`

GetProposedTransfer returns the ProposedTransfer field if non-nil, zero value otherwise.

### GetProposedTransferOk

`func (o *TransferAuthorization) GetProposedTransferOk() (*TransferAuthorizationProposedTransfer, bool)`

GetProposedTransferOk returns a tuple with the ProposedTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedTransfer

`func (o *TransferAuthorization) SetProposedTransfer(v TransferAuthorizationProposedTransfer)`

SetProposedTransfer sets ProposedTransfer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


