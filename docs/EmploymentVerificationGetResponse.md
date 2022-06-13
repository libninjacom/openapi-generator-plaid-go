# EmploymentVerificationGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Employments** | [**[]EmploymentVerification**](EmploymentVerification.md) | A list of employment verification summaries. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewEmploymentVerificationGetResponse

`func NewEmploymentVerificationGetResponse(employments []EmploymentVerification, requestId string, ) *EmploymentVerificationGetResponse`

NewEmploymentVerificationGetResponse instantiates a new EmploymentVerificationGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmploymentVerificationGetResponseWithDefaults

`func NewEmploymentVerificationGetResponseWithDefaults() *EmploymentVerificationGetResponse`

NewEmploymentVerificationGetResponseWithDefaults instantiates a new EmploymentVerificationGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployments

`func (o *EmploymentVerificationGetResponse) GetEmployments() []EmploymentVerification`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *EmploymentVerificationGetResponse) GetEmploymentsOk() (*[]EmploymentVerification, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *EmploymentVerificationGetResponse) SetEmployments(v []EmploymentVerification)`

SetEmployments sets Employments field to given value.


### GetRequestId

`func (o *EmploymentVerificationGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EmploymentVerificationGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EmploymentVerificationGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


