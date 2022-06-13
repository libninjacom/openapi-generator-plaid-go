# SignalDecisionReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**ClientTransactionId** | **string** | Must be the same as the &#x60;client_transaction_id&#x60; supplied when calling &#x60;/signal/evaluate&#x60; | 
**Initiated** | **bool** | &#x60;true&#x60; if the ACH transaction was initiated, &#x60;false&#x60; otherwise. | 
**DaysFundsOnHold** | Pointer to **NullableInt32** | The actual number of days (hold time) since the ACH debit transaction that you wait before making funds available to your customers. The holding time could affect the ACH return rate. For example, use 0 if you make funds available to your customers instantly or the same day following the debit transaction, or 1 if you make funds available the next day following the debit initialization. | [optional] 

## Methods

### NewSignalDecisionReportRequest

`func NewSignalDecisionReportRequest(clientTransactionId string, initiated bool, ) *SignalDecisionReportRequest`

NewSignalDecisionReportRequest instantiates a new SignalDecisionReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalDecisionReportRequestWithDefaults

`func NewSignalDecisionReportRequestWithDefaults() *SignalDecisionReportRequest`

NewSignalDecisionReportRequestWithDefaults instantiates a new SignalDecisionReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SignalDecisionReportRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SignalDecisionReportRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SignalDecisionReportRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SignalDecisionReportRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *SignalDecisionReportRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SignalDecisionReportRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SignalDecisionReportRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SignalDecisionReportRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetClientTransactionId

`func (o *SignalDecisionReportRequest) GetClientTransactionId() string`

GetClientTransactionId returns the ClientTransactionId field if non-nil, zero value otherwise.

### GetClientTransactionIdOk

`func (o *SignalDecisionReportRequest) GetClientTransactionIdOk() (*string, bool)`

GetClientTransactionIdOk returns a tuple with the ClientTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTransactionId

`func (o *SignalDecisionReportRequest) SetClientTransactionId(v string)`

SetClientTransactionId sets ClientTransactionId field to given value.


### GetInitiated

`func (o *SignalDecisionReportRequest) GetInitiated() bool`

GetInitiated returns the Initiated field if non-nil, zero value otherwise.

### GetInitiatedOk

`func (o *SignalDecisionReportRequest) GetInitiatedOk() (*bool, bool)`

GetInitiatedOk returns a tuple with the Initiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiated

`func (o *SignalDecisionReportRequest) SetInitiated(v bool)`

SetInitiated sets Initiated field to given value.


### GetDaysFundsOnHold

`func (o *SignalDecisionReportRequest) GetDaysFundsOnHold() int32`

GetDaysFundsOnHold returns the DaysFundsOnHold field if non-nil, zero value otherwise.

### GetDaysFundsOnHoldOk

`func (o *SignalDecisionReportRequest) GetDaysFundsOnHoldOk() (*int32, bool)`

GetDaysFundsOnHoldOk returns a tuple with the DaysFundsOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysFundsOnHold

`func (o *SignalDecisionReportRequest) SetDaysFundsOnHold(v int32)`

SetDaysFundsOnHold sets DaysFundsOnHold field to given value.

### HasDaysFundsOnHold

`func (o *SignalDecisionReportRequest) HasDaysFundsOnHold() bool`

HasDaysFundsOnHold returns a boolean if a field has been set.

### SetDaysFundsOnHoldNil

`func (o *SignalDecisionReportRequest) SetDaysFundsOnHoldNil(b bool)`

 SetDaysFundsOnHoldNil sets the value for DaysFundsOnHold to be an explicit nil

### UnsetDaysFundsOnHold
`func (o *SignalDecisionReportRequest) UnsetDaysFundsOnHold()`

UnsetDaysFundsOnHold ensures that no value is present for DaysFundsOnHold, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


