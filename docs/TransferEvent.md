# TransferEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **int32** | Plaid’s unique identifier for this event. IDs are sequential unsigned 64-bit integers. | 
**Timestamp** | **time.Time** | The datetime when this event occurred. This will be of the form &#x60;2006-01-02T15:04:05Z&#x60;. | 
**EventType** | [**TransferEventType**](TransferEventType.md) |  | 
**AccountId** | **string** | The account ID associated with the transfer. | 
**TransferId** | **string** | Plaid’s unique identifier for a transfer. | 
**OriginationAccountId** | **NullableString** | The ID of the origination account that this balance belongs to. | 
**TransferType** | [**TransferType**](TransferType.md) |  | 
**TransferAmount** | **string** | The amount of the transfer (decimal string with two digits of precision e.g. \&quot;10.00\&quot;). | 
**FailureReason** | [**NullableTransferFailure**](TransferFailure.md) |  | 
**SweepId** | **NullableString** | Plaid’s unique identifier for a sweep. | 
**SweepAmount** | **NullableString** | A signed amount of how much was &#x60;swept&#x60; or &#x60;reverse_swept&#x60; (decimal string with two digits of precision e.g. \&quot;-5.50\&quot;). | 

## Methods

### NewTransferEvent

`func NewTransferEvent(eventId int32, timestamp time.Time, eventType TransferEventType, accountId string, transferId string, originationAccountId NullableString, transferType TransferType, transferAmount string, failureReason NullableTransferFailure, sweepId NullableString, sweepAmount NullableString, ) *TransferEvent`

NewTransferEvent instantiates a new TransferEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferEventWithDefaults

`func NewTransferEventWithDefaults() *TransferEvent`

NewTransferEventWithDefaults instantiates a new TransferEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *TransferEvent) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *TransferEvent) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *TransferEvent) SetEventId(v int32)`

SetEventId sets EventId field to given value.


### GetTimestamp

`func (o *TransferEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TransferEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TransferEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEventType

`func (o *TransferEvent) GetEventType() TransferEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TransferEvent) GetEventTypeOk() (*TransferEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TransferEvent) SetEventType(v TransferEventType)`

SetEventType sets EventType field to given value.


### GetAccountId

`func (o *TransferEvent) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferEvent) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferEvent) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTransferId

`func (o *TransferEvent) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *TransferEvent) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *TransferEvent) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.


### GetOriginationAccountId

`func (o *TransferEvent) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *TransferEvent) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *TransferEvent) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.


### SetOriginationAccountIdNil

`func (o *TransferEvent) SetOriginationAccountIdNil(b bool)`

 SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil

### UnsetOriginationAccountId
`func (o *TransferEvent) UnsetOriginationAccountId()`

UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
### GetTransferType

`func (o *TransferEvent) GetTransferType() TransferType`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *TransferEvent) GetTransferTypeOk() (*TransferType, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *TransferEvent) SetTransferType(v TransferType)`

SetTransferType sets TransferType field to given value.


### GetTransferAmount

`func (o *TransferEvent) GetTransferAmount() string`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *TransferEvent) GetTransferAmountOk() (*string, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *TransferEvent) SetTransferAmount(v string)`

SetTransferAmount sets TransferAmount field to given value.


### GetFailureReason

`func (o *TransferEvent) GetFailureReason() TransferFailure`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *TransferEvent) GetFailureReasonOk() (*TransferFailure, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *TransferEvent) SetFailureReason(v TransferFailure)`

SetFailureReason sets FailureReason field to given value.


### SetFailureReasonNil

`func (o *TransferEvent) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *TransferEvent) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetSweepId

`func (o *TransferEvent) GetSweepId() string`

GetSweepId returns the SweepId field if non-nil, zero value otherwise.

### GetSweepIdOk

`func (o *TransferEvent) GetSweepIdOk() (*string, bool)`

GetSweepIdOk returns a tuple with the SweepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweepId

`func (o *TransferEvent) SetSweepId(v string)`

SetSweepId sets SweepId field to given value.


### SetSweepIdNil

`func (o *TransferEvent) SetSweepIdNil(b bool)`

 SetSweepIdNil sets the value for SweepId to be an explicit nil

### UnsetSweepId
`func (o *TransferEvent) UnsetSweepId()`

UnsetSweepId ensures that no value is present for SweepId, not even an explicit nil
### GetSweepAmount

`func (o *TransferEvent) GetSweepAmount() string`

GetSweepAmount returns the SweepAmount field if non-nil, zero value otherwise.

### GetSweepAmountOk

`func (o *TransferEvent) GetSweepAmountOk() (*string, bool)`

GetSweepAmountOk returns a tuple with the SweepAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweepAmount

`func (o *TransferEvent) SetSweepAmount(v string)`

SetSweepAmount sets SweepAmount field to given value.


### SetSweepAmountNil

`func (o *TransferEvent) SetSweepAmountNil(b bool)`

 SetSweepAmountNil sets the value for SweepAmount to be an explicit nil

### UnsetSweepAmount
`func (o *TransferEvent) UnsetSweepAmount()`

UnsetSweepAmount ensures that no value is present for SweepAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


