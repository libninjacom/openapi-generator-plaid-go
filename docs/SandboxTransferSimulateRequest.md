# SandboxTransferSimulateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**TransferId** | **string** | Plaid’s unique identifier for a transfer. | 
**EventType** | **string** | The asynchronous event to be simulated. May be: &#x60;posted&#x60;, &#x60;failed&#x60;, or &#x60;reversed&#x60;.  An error will be returned if the event type is incompatible with the current transfer status. Compatible status --&gt; event type transitions include:  &#x60;pending&#x60; --&gt; &#x60;failed&#x60;  &#x60;pending&#x60; --&gt; &#x60;posted&#x60;  &#x60;posted&#x60; --&gt; &#x60;reversed&#x60;  | 
**FailureReason** | Pointer to [**NullableTransferFailure**](TransferFailure.md) |  | [optional] 

## Methods

### NewSandboxTransferSimulateRequest

`func NewSandboxTransferSimulateRequest(transferId string, eventType string, ) *SandboxTransferSimulateRequest`

NewSandboxTransferSimulateRequest instantiates a new SandboxTransferSimulateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxTransferSimulateRequestWithDefaults

`func NewSandboxTransferSimulateRequestWithDefaults() *SandboxTransferSimulateRequest`

NewSandboxTransferSimulateRequestWithDefaults instantiates a new SandboxTransferSimulateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SandboxTransferSimulateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SandboxTransferSimulateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SandboxTransferSimulateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SandboxTransferSimulateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *SandboxTransferSimulateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SandboxTransferSimulateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SandboxTransferSimulateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SandboxTransferSimulateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTransferId

`func (o *SandboxTransferSimulateRequest) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *SandboxTransferSimulateRequest) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *SandboxTransferSimulateRequest) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.


### GetEventType

`func (o *SandboxTransferSimulateRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SandboxTransferSimulateRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SandboxTransferSimulateRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetFailureReason

`func (o *SandboxTransferSimulateRequest) GetFailureReason() TransferFailure`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *SandboxTransferSimulateRequest) GetFailureReasonOk() (*TransferFailure, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *SandboxTransferSimulateRequest) SetFailureReason(v TransferFailure)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *SandboxTransferSimulateRequest) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *SandboxTransferSimulateRequest) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *SandboxTransferSimulateRequest) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


