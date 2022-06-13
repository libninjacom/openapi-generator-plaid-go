# IncomeVerificationPrecheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrecheckId** | **string** | ID of the precheck. Provide this value when calling &#x60;/link/token/create&#x60; in order to optimize Link conversion. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 
**Confidence** | [**IncomeVerificationPrecheckConfidence**](IncomeVerificationPrecheckConfidence.md) |  | 

## Methods

### NewIncomeVerificationPrecheckResponse

`func NewIncomeVerificationPrecheckResponse(precheckId string, requestId string, confidence IncomeVerificationPrecheckConfidence, ) *IncomeVerificationPrecheckResponse`

NewIncomeVerificationPrecheckResponse instantiates a new IncomeVerificationPrecheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationPrecheckResponseWithDefaults

`func NewIncomeVerificationPrecheckResponseWithDefaults() *IncomeVerificationPrecheckResponse`

NewIncomeVerificationPrecheckResponseWithDefaults instantiates a new IncomeVerificationPrecheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrecheckId

`func (o *IncomeVerificationPrecheckResponse) GetPrecheckId() string`

GetPrecheckId returns the PrecheckId field if non-nil, zero value otherwise.

### GetPrecheckIdOk

`func (o *IncomeVerificationPrecheckResponse) GetPrecheckIdOk() (*string, bool)`

GetPrecheckIdOk returns a tuple with the PrecheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecheckId

`func (o *IncomeVerificationPrecheckResponse) SetPrecheckId(v string)`

SetPrecheckId sets PrecheckId field to given value.


### GetRequestId

`func (o *IncomeVerificationPrecheckResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *IncomeVerificationPrecheckResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *IncomeVerificationPrecheckResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetConfidence

`func (o *IncomeVerificationPrecheckResponse) GetConfidence() IncomeVerificationPrecheckConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *IncomeVerificationPrecheckResponse) GetConfidenceOk() (*IncomeVerificationPrecheckConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *IncomeVerificationPrecheckResponse) SetConfidence(v IncomeVerificationPrecheckConfidence)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


