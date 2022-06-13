# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Plaid’s unique identifier for a transfer. | 
**AchClass** | [**ACHClass**](ACHClass.md) |  | 
**AccountId** | **string** | The account ID that should be credited/debited for this transfer. | 
**Type** | [**TransferType**](TransferType.md) |  | 
**User** | [**TransferUserInResponse**](TransferUserInResponse.md) |  | 
**Amount** | **string** | The amount of the transfer (decimal string with two digits of precision e.g. \&quot;10.00\&quot;). | 
**Description** | **string** | The description of the transfer. | 
**Created** | **time.Time** | The datetime when this transfer was created. This will be of the form &#x60;2006-01-02T15:04:05Z&#x60; | 
**Status** | [**TransferStatus**](TransferStatus.md) |  | 
**SweepStatus** | Pointer to [**NullableTransferSweepStatus**](TransferSweepStatus.md) |  | [optional] 
**Network** | [**TransferNetwork**](TransferNetwork.md) |  | 
**Cancellable** | **bool** | When &#x60;true&#x60;, you can still cancel this transfer. | 
**FailureReason** | [**NullableTransferFailure**](TransferFailure.md) |  | 
**Metadata** | **map[string]string** | The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters  | 
**OriginationAccountId** | **string** | Plaid’s unique identifier for the origination account that was used for this transfer. | 
**GuaranteeDecision** | [**NullableTransferAuthorizationGuaranteeDecision**](TransferAuthorizationGuaranteeDecision.md) |  | 
**GuaranteeDecisionRationale** | [**NullableTransferAuthorizationGuaranteeDecisionRationale**](TransferAuthorizationGuaranteeDecisionRationale.md) |  | 
**IsoCurrencyCode** | **string** | The currency of the transfer amount, e.g. \&quot;USD\&quot; | 

## Methods

### NewTransfer

`func NewTransfer(id string, achClass ACHClass, accountId string, type_ TransferType, user TransferUserInResponse, amount string, description string, created time.Time, status TransferStatus, network TransferNetwork, cancellable bool, failureReason NullableTransferFailure, metadata map[string]string, originationAccountId string, guaranteeDecision NullableTransferAuthorizationGuaranteeDecision, guaranteeDecisionRationale NullableTransferAuthorizationGuaranteeDecisionRationale, isoCurrencyCode string, ) *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transfer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transfer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transfer) SetId(v string)`

SetId sets Id field to given value.


### GetAchClass

`func (o *Transfer) GetAchClass() ACHClass`

GetAchClass returns the AchClass field if non-nil, zero value otherwise.

### GetAchClassOk

`func (o *Transfer) GetAchClassOk() (*ACHClass, bool)`

GetAchClassOk returns a tuple with the AchClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchClass

`func (o *Transfer) SetAchClass(v ACHClass)`

SetAchClass sets AchClass field to given value.


### GetAccountId

`func (o *Transfer) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Transfer) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Transfer) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetType

`func (o *Transfer) GetType() TransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transfer) GetTypeOk() (*TransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transfer) SetType(v TransferType)`

SetType sets Type field to given value.


### GetUser

`func (o *Transfer) GetUser() TransferUserInResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Transfer) GetUserOk() (*TransferUserInResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Transfer) SetUser(v TransferUserInResponse)`

SetUser sets User field to given value.


### GetAmount

`func (o *Transfer) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transfer) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transfer) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *Transfer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transfer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transfer) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreated

`func (o *Transfer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Transfer) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Transfer) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetStatus

`func (o *Transfer) GetStatus() TransferStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transfer) GetStatusOk() (*TransferStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transfer) SetStatus(v TransferStatus)`

SetStatus sets Status field to given value.


### GetSweepStatus

`func (o *Transfer) GetSweepStatus() TransferSweepStatus`

GetSweepStatus returns the SweepStatus field if non-nil, zero value otherwise.

### GetSweepStatusOk

`func (o *Transfer) GetSweepStatusOk() (*TransferSweepStatus, bool)`

GetSweepStatusOk returns a tuple with the SweepStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweepStatus

`func (o *Transfer) SetSweepStatus(v TransferSweepStatus)`

SetSweepStatus sets SweepStatus field to given value.

### HasSweepStatus

`func (o *Transfer) HasSweepStatus() bool`

HasSweepStatus returns a boolean if a field has been set.

### SetSweepStatusNil

`func (o *Transfer) SetSweepStatusNil(b bool)`

 SetSweepStatusNil sets the value for SweepStatus to be an explicit nil

### UnsetSweepStatus
`func (o *Transfer) UnsetSweepStatus()`

UnsetSweepStatus ensures that no value is present for SweepStatus, not even an explicit nil
### GetNetwork

`func (o *Transfer) GetNetwork() TransferNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Transfer) GetNetworkOk() (*TransferNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Transfer) SetNetwork(v TransferNetwork)`

SetNetwork sets Network field to given value.


### GetCancellable

`func (o *Transfer) GetCancellable() bool`

GetCancellable returns the Cancellable field if non-nil, zero value otherwise.

### GetCancellableOk

`func (o *Transfer) GetCancellableOk() (*bool, bool)`

GetCancellableOk returns a tuple with the Cancellable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellable

`func (o *Transfer) SetCancellable(v bool)`

SetCancellable sets Cancellable field to given value.


### GetFailureReason

`func (o *Transfer) GetFailureReason() TransferFailure`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *Transfer) GetFailureReasonOk() (*TransferFailure, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *Transfer) SetFailureReason(v TransferFailure)`

SetFailureReason sets FailureReason field to given value.


### SetFailureReasonNil

`func (o *Transfer) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *Transfer) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetMetadata

`func (o *Transfer) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Transfer) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Transfer) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *Transfer) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Transfer) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOriginationAccountId

`func (o *Transfer) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *Transfer) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *Transfer) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.


### GetGuaranteeDecision

`func (o *Transfer) GetGuaranteeDecision() TransferAuthorizationGuaranteeDecision`

GetGuaranteeDecision returns the GuaranteeDecision field if non-nil, zero value otherwise.

### GetGuaranteeDecisionOk

`func (o *Transfer) GetGuaranteeDecisionOk() (*TransferAuthorizationGuaranteeDecision, bool)`

GetGuaranteeDecisionOk returns a tuple with the GuaranteeDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeDecision

`func (o *Transfer) SetGuaranteeDecision(v TransferAuthorizationGuaranteeDecision)`

SetGuaranteeDecision sets GuaranteeDecision field to given value.


### SetGuaranteeDecisionNil

`func (o *Transfer) SetGuaranteeDecisionNil(b bool)`

 SetGuaranteeDecisionNil sets the value for GuaranteeDecision to be an explicit nil

### UnsetGuaranteeDecision
`func (o *Transfer) UnsetGuaranteeDecision()`

UnsetGuaranteeDecision ensures that no value is present for GuaranteeDecision, not even an explicit nil
### GetGuaranteeDecisionRationale

`func (o *Transfer) GetGuaranteeDecisionRationale() TransferAuthorizationGuaranteeDecisionRationale`

GetGuaranteeDecisionRationale returns the GuaranteeDecisionRationale field if non-nil, zero value otherwise.

### GetGuaranteeDecisionRationaleOk

`func (o *Transfer) GetGuaranteeDecisionRationaleOk() (*TransferAuthorizationGuaranteeDecisionRationale, bool)`

GetGuaranteeDecisionRationaleOk returns a tuple with the GuaranteeDecisionRationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeDecisionRationale

`func (o *Transfer) SetGuaranteeDecisionRationale(v TransferAuthorizationGuaranteeDecisionRationale)`

SetGuaranteeDecisionRationale sets GuaranteeDecisionRationale field to given value.


### SetGuaranteeDecisionRationaleNil

`func (o *Transfer) SetGuaranteeDecisionRationaleNil(b bool)`

 SetGuaranteeDecisionRationaleNil sets the value for GuaranteeDecisionRationale to be an explicit nil

### UnsetGuaranteeDecisionRationale
`func (o *Transfer) UnsetGuaranteeDecisionRationale()`

UnsetGuaranteeDecisionRationale ensures that no value is present for GuaranteeDecisionRationale, not even an explicit nil
### GetIsoCurrencyCode

`func (o *Transfer) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *Transfer) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *Transfer) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


