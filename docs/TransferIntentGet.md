# TransferIntentGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Plaid&#39;s unique identifier for a transfer intent object. | 
**Created** | **time.Time** | The datetime the transfer was created. This will be of the form &#x60;2006-01-02T15:04:05Z&#x60;. | 
**Status** | **string** | The status of the transfer intent.  - &#x60;PENDING&#x60; – The transfer intent is pending. - &#x60;SUCCEEDED&#x60; – The transfer intent was successfully created. - &#x60;FAILED&#x60; – The transfer intent was unable to be created. | 
**TransferId** | **NullableString** | Plaid&#39;s unique identifier for the transfer created through the UI. Returned only if the transfer was successfully created. Null value otherwise. | 
**FailureReason** | [**NullableTransferIntentGetFailureReason**](TransferIntentGetFailureReason.md) |  | 
**AuthorizationDecision** | **NullableString** |  A decision regarding the proposed transfer.  &#x60;APPROVED&#x60; – The proposed transfer has received the end user&#39;s consent and has been approved for processing. Plaid has also reviewed the proposed transfer and has approved it for processing.   &#x60;PERMITTED&#x60; – Plaid was unable to fetch the information required to approve or decline the proposed transfer. You may proceed with the transfer, but further review is recommended. Plaid is awaiting further instructions from the client.  &#x60;DECLINED&#x60; – Plaid reviewed the proposed transfer and declined processing. Refer to the &#x60;code&#x60; field in the &#x60;decision_rationale&#x60; object for details. Null value otherwise. | 
**AuthorizationDecisionRationale** | [**NullableTransferAuthorizationDecisionRationale**](TransferAuthorizationDecisionRationale.md) |  | 
**AccountId** | Pointer to **NullableString** | The Plaid &#x60;account_id&#x60; for the account that will be debited or credited. Returned only if &#x60;account_id&#x60; was set on intent creation. | [optional] 
**OriginationAccountId** | **string** | Plaid’s unique identifier for the origination account used for the transfer. | 
**Amount** | **string** | The amount of the transfer (decimal string with two digits of precision e.g. \&quot;10.00\&quot;). | 
**Mode** | [**TransferIntentCreateMode**](TransferIntentCreateMode.md) |  | 
**AchClass** | [**ACHClass**](ACHClass.md) |  | 
**User** | [**TransferUserInResponse**](TransferUserInResponse.md) |  | 
**Description** | **string** | A description for the underlying transfer. Maximum of 8 characters. | 
**Metadata** | Pointer to **map[string]string** | The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters  | [optional] 
**IsoCurrencyCode** | **string** | The currency of the transfer amount, e.g. \&quot;USD\&quot; | 

## Methods

### NewTransferIntentGet

`func NewTransferIntentGet(id string, created time.Time, status string, transferId NullableString, failureReason NullableTransferIntentGetFailureReason, authorizationDecision NullableString, authorizationDecisionRationale NullableTransferAuthorizationDecisionRationale, originationAccountId string, amount string, mode TransferIntentCreateMode, achClass ACHClass, user TransferUserInResponse, description string, isoCurrencyCode string, ) *TransferIntentGet`

NewTransferIntentGet instantiates a new TransferIntentGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferIntentGetWithDefaults

`func NewTransferIntentGetWithDefaults() *TransferIntentGet`

NewTransferIntentGetWithDefaults instantiates a new TransferIntentGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransferIntentGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferIntentGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferIntentGet) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *TransferIntentGet) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TransferIntentGet) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TransferIntentGet) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetStatus

`func (o *TransferIntentGet) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferIntentGet) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferIntentGet) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransferId

`func (o *TransferIntentGet) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *TransferIntentGet) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *TransferIntentGet) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.


### SetTransferIdNil

`func (o *TransferIntentGet) SetTransferIdNil(b bool)`

 SetTransferIdNil sets the value for TransferId to be an explicit nil

### UnsetTransferId
`func (o *TransferIntentGet) UnsetTransferId()`

UnsetTransferId ensures that no value is present for TransferId, not even an explicit nil
### GetFailureReason

`func (o *TransferIntentGet) GetFailureReason() TransferIntentGetFailureReason`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *TransferIntentGet) GetFailureReasonOk() (*TransferIntentGetFailureReason, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *TransferIntentGet) SetFailureReason(v TransferIntentGetFailureReason)`

SetFailureReason sets FailureReason field to given value.


### SetFailureReasonNil

`func (o *TransferIntentGet) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *TransferIntentGet) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetAuthorizationDecision

`func (o *TransferIntentGet) GetAuthorizationDecision() string`

GetAuthorizationDecision returns the AuthorizationDecision field if non-nil, zero value otherwise.

### GetAuthorizationDecisionOk

`func (o *TransferIntentGet) GetAuthorizationDecisionOk() (*string, bool)`

GetAuthorizationDecisionOk returns a tuple with the AuthorizationDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDecision

`func (o *TransferIntentGet) SetAuthorizationDecision(v string)`

SetAuthorizationDecision sets AuthorizationDecision field to given value.


### SetAuthorizationDecisionNil

`func (o *TransferIntentGet) SetAuthorizationDecisionNil(b bool)`

 SetAuthorizationDecisionNil sets the value for AuthorizationDecision to be an explicit nil

### UnsetAuthorizationDecision
`func (o *TransferIntentGet) UnsetAuthorizationDecision()`

UnsetAuthorizationDecision ensures that no value is present for AuthorizationDecision, not even an explicit nil
### GetAuthorizationDecisionRationale

`func (o *TransferIntentGet) GetAuthorizationDecisionRationale() TransferAuthorizationDecisionRationale`

GetAuthorizationDecisionRationale returns the AuthorizationDecisionRationale field if non-nil, zero value otherwise.

### GetAuthorizationDecisionRationaleOk

`func (o *TransferIntentGet) GetAuthorizationDecisionRationaleOk() (*TransferAuthorizationDecisionRationale, bool)`

GetAuthorizationDecisionRationaleOk returns a tuple with the AuthorizationDecisionRationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDecisionRationale

`func (o *TransferIntentGet) SetAuthorizationDecisionRationale(v TransferAuthorizationDecisionRationale)`

SetAuthorizationDecisionRationale sets AuthorizationDecisionRationale field to given value.


### SetAuthorizationDecisionRationaleNil

`func (o *TransferIntentGet) SetAuthorizationDecisionRationaleNil(b bool)`

 SetAuthorizationDecisionRationaleNil sets the value for AuthorizationDecisionRationale to be an explicit nil

### UnsetAuthorizationDecisionRationale
`func (o *TransferIntentGet) UnsetAuthorizationDecisionRationale()`

UnsetAuthorizationDecisionRationale ensures that no value is present for AuthorizationDecisionRationale, not even an explicit nil
### GetAccountId

`func (o *TransferIntentGet) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferIntentGet) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferIntentGet) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TransferIntentGet) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *TransferIntentGet) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *TransferIntentGet) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetOriginationAccountId

`func (o *TransferIntentGet) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *TransferIntentGet) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *TransferIntentGet) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.


### GetAmount

`func (o *TransferIntentGet) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferIntentGet) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferIntentGet) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMode

`func (o *TransferIntentGet) GetMode() TransferIntentCreateMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TransferIntentGet) GetModeOk() (*TransferIntentCreateMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TransferIntentGet) SetMode(v TransferIntentCreateMode)`

SetMode sets Mode field to given value.


### GetAchClass

`func (o *TransferIntentGet) GetAchClass() ACHClass`

GetAchClass returns the AchClass field if non-nil, zero value otherwise.

### GetAchClassOk

`func (o *TransferIntentGet) GetAchClassOk() (*ACHClass, bool)`

GetAchClassOk returns a tuple with the AchClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchClass

`func (o *TransferIntentGet) SetAchClass(v ACHClass)`

SetAchClass sets AchClass field to given value.


### GetUser

`func (o *TransferIntentGet) GetUser() TransferUserInResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TransferIntentGet) GetUserOk() (*TransferUserInResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TransferIntentGet) SetUser(v TransferUserInResponse)`

SetUser sets User field to given value.


### GetDescription

`func (o *TransferIntentGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransferIntentGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransferIntentGet) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMetadata

`func (o *TransferIntentGet) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TransferIntentGet) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TransferIntentGet) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TransferIntentGet) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *TransferIntentGet) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *TransferIntentGet) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetIsoCurrencyCode

`func (o *TransferIntentGet) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *TransferIntentGet) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *TransferIntentGet) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


