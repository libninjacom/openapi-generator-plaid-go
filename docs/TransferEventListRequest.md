# TransferEventListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**StartDate** | Pointer to **NullableTime** | The start datetime of transfers to list. This should be in RFC 3339 format (i.e. &#x60;2019-12-06T22:35:49Z&#x60;) | [optional] 
**EndDate** | Pointer to **NullableTime** | The end datetime of transfers to list. This should be in RFC 3339 format (i.e. &#x60;2019-12-06T22:35:49Z&#x60;) | [optional] 
**TransferId** | Pointer to **NullableString** | Plaid’s unique identifier for a transfer. | [optional] 
**AccountId** | Pointer to **NullableString** | The account ID to get events for all transactions to/from an account. | [optional] 
**TransferType** | Pointer to **NullableString** | The type of transfer. This will be either &#x60;debit&#x60; or &#x60;credit&#x60;.  A &#x60;debit&#x60; indicates a transfer of money into your origination account; a &#x60;credit&#x60; indicates a transfer of money out of your origination account. | [optional] 
**EventTypes** | Pointer to [**[]TransferEventType**](TransferEventType.md) | Filter events by event type. | [optional] 
**SweepId** | Pointer to **string** | Plaid’s unique identifier for a sweep. | [optional] 
**Count** | Pointer to **NullableInt32** | The maximum number of transfer events to return. If the number of events matching the above parameters is greater than &#x60;count&#x60;, the most recent events will be returned. | [optional] [default to 25]
**Offset** | Pointer to **NullableInt32** | The offset into the list of transfer events. When &#x60;count&#x60;&#x3D;25 and &#x60;offset&#x60;&#x3D;0, the first 25 events will be returned. When &#x60;count&#x60;&#x3D;25 and &#x60;offset&#x60;&#x3D;25, the next 25 bank transfer events will be returned. | [optional] [default to 0]
**OriginationAccountId** | Pointer to **NullableString** | The origination account ID to get events for transfers from a specific origination account. | [optional] 

## Methods

### NewTransferEventListRequest

`func NewTransferEventListRequest() *TransferEventListRequest`

NewTransferEventListRequest instantiates a new TransferEventListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferEventListRequestWithDefaults

`func NewTransferEventListRequestWithDefaults() *TransferEventListRequest`

NewTransferEventListRequestWithDefaults instantiates a new TransferEventListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TransferEventListRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TransferEventListRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TransferEventListRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TransferEventListRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *TransferEventListRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TransferEventListRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TransferEventListRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *TransferEventListRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStartDate

`func (o *TransferEventListRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TransferEventListRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TransferEventListRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TransferEventListRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *TransferEventListRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *TransferEventListRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *TransferEventListRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TransferEventListRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TransferEventListRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TransferEventListRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *TransferEventListRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *TransferEventListRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetTransferId

`func (o *TransferEventListRequest) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *TransferEventListRequest) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *TransferEventListRequest) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.

### HasTransferId

`func (o *TransferEventListRequest) HasTransferId() bool`

HasTransferId returns a boolean if a field has been set.

### SetTransferIdNil

`func (o *TransferEventListRequest) SetTransferIdNil(b bool)`

 SetTransferIdNil sets the value for TransferId to be an explicit nil

### UnsetTransferId
`func (o *TransferEventListRequest) UnsetTransferId()`

UnsetTransferId ensures that no value is present for TransferId, not even an explicit nil
### GetAccountId

`func (o *TransferEventListRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferEventListRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferEventListRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TransferEventListRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *TransferEventListRequest) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *TransferEventListRequest) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetTransferType

`func (o *TransferEventListRequest) GetTransferType() string`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *TransferEventListRequest) GetTransferTypeOk() (*string, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *TransferEventListRequest) SetTransferType(v string)`

SetTransferType sets TransferType field to given value.

### HasTransferType

`func (o *TransferEventListRequest) HasTransferType() bool`

HasTransferType returns a boolean if a field has been set.

### SetTransferTypeNil

`func (o *TransferEventListRequest) SetTransferTypeNil(b bool)`

 SetTransferTypeNil sets the value for TransferType to be an explicit nil

### UnsetTransferType
`func (o *TransferEventListRequest) UnsetTransferType()`

UnsetTransferType ensures that no value is present for TransferType, not even an explicit nil
### GetEventTypes

`func (o *TransferEventListRequest) GetEventTypes() []TransferEventType`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *TransferEventListRequest) GetEventTypesOk() (*[]TransferEventType, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *TransferEventListRequest) SetEventTypes(v []TransferEventType)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *TransferEventListRequest) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetSweepId

`func (o *TransferEventListRequest) GetSweepId() string`

GetSweepId returns the SweepId field if non-nil, zero value otherwise.

### GetSweepIdOk

`func (o *TransferEventListRequest) GetSweepIdOk() (*string, bool)`

GetSweepIdOk returns a tuple with the SweepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweepId

`func (o *TransferEventListRequest) SetSweepId(v string)`

SetSweepId sets SweepId field to given value.

### HasSweepId

`func (o *TransferEventListRequest) HasSweepId() bool`

HasSweepId returns a boolean if a field has been set.

### GetCount

`func (o *TransferEventListRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TransferEventListRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TransferEventListRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TransferEventListRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *TransferEventListRequest) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *TransferEventListRequest) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetOffset

`func (o *TransferEventListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TransferEventListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TransferEventListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TransferEventListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *TransferEventListRequest) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *TransferEventListRequest) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetOriginationAccountId

`func (o *TransferEventListRequest) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *TransferEventListRequest) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *TransferEventListRequest) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.

### HasOriginationAccountId

`func (o *TransferEventListRequest) HasOriginationAccountId() bool`

HasOriginationAccountId returns a boolean if a field has been set.

### SetOriginationAccountIdNil

`func (o *TransferEventListRequest) SetOriginationAccountIdNil(b bool)`

 SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil

### UnsetOriginationAccountId
`func (o *TransferEventListRequest) UnsetOriginationAccountId()`

UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


