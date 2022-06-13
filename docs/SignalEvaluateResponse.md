# SignalEvaluateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 
**Scores** | [**SignalScores**](SignalScores.md) |  | 
**CoreAttributes** | Pointer to [**SignalEvaluateCoreAttributes**](SignalEvaluateCoreAttributes.md) |  | [optional] 

## Methods

### NewSignalEvaluateResponse

`func NewSignalEvaluateResponse(requestId string, scores SignalScores, ) *SignalEvaluateResponse`

NewSignalEvaluateResponse instantiates a new SignalEvaluateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalEvaluateResponseWithDefaults

`func NewSignalEvaluateResponseWithDefaults() *SignalEvaluateResponse`

NewSignalEvaluateResponseWithDefaults instantiates a new SignalEvaluateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *SignalEvaluateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SignalEvaluateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SignalEvaluateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetScores

`func (o *SignalEvaluateResponse) GetScores() SignalScores`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *SignalEvaluateResponse) GetScoresOk() (*SignalScores, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *SignalEvaluateResponse) SetScores(v SignalScores)`

SetScores sets Scores field to given value.


### GetCoreAttributes

`func (o *SignalEvaluateResponse) GetCoreAttributes() SignalEvaluateCoreAttributes`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *SignalEvaluateResponse) GetCoreAttributesOk() (*SignalEvaluateCoreAttributes, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *SignalEvaluateResponse) SetCoreAttributes(v SignalEvaluateCoreAttributes)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *SignalEvaluateResponse) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


