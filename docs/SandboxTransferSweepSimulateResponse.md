# SandboxTransferSweepSimulateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sweep** | Pointer to [**SimulatedTransferSweep**](SimulatedTransferSweep.md) |  | [optional] 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewSandboxTransferSweepSimulateResponse

`func NewSandboxTransferSweepSimulateResponse(requestId string, ) *SandboxTransferSweepSimulateResponse`

NewSandboxTransferSweepSimulateResponse instantiates a new SandboxTransferSweepSimulateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxTransferSweepSimulateResponseWithDefaults

`func NewSandboxTransferSweepSimulateResponseWithDefaults() *SandboxTransferSweepSimulateResponse`

NewSandboxTransferSweepSimulateResponseWithDefaults instantiates a new SandboxTransferSweepSimulateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSweep

`func (o *SandboxTransferSweepSimulateResponse) GetSweep() SimulatedTransferSweep`

GetSweep returns the Sweep field if non-nil, zero value otherwise.

### GetSweepOk

`func (o *SandboxTransferSweepSimulateResponse) GetSweepOk() (*SimulatedTransferSweep, bool)`

GetSweepOk returns a tuple with the Sweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweep

`func (o *SandboxTransferSweepSimulateResponse) SetSweep(v SimulatedTransferSweep)`

SetSweep sets Sweep field to given value.

### HasSweep

`func (o *SandboxTransferSweepSimulateResponse) HasSweep() bool`

HasSweep returns a boolean if a field has been set.

### GetRequestId

`func (o *SandboxTransferSweepSimulateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SandboxTransferSweepSimulateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SandboxTransferSweepSimulateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


