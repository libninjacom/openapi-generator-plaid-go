# IncomeVerificationRefreshResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 
**VerificationRefreshStatus** | [**VerificationRefreshStatus**](VerificationRefreshStatus.md) |  | 

## Methods

### NewIncomeVerificationRefreshResponse

`func NewIncomeVerificationRefreshResponse(requestId string, verificationRefreshStatus VerificationRefreshStatus, ) *IncomeVerificationRefreshResponse`

NewIncomeVerificationRefreshResponse instantiates a new IncomeVerificationRefreshResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationRefreshResponseWithDefaults

`func NewIncomeVerificationRefreshResponseWithDefaults() *IncomeVerificationRefreshResponse`

NewIncomeVerificationRefreshResponseWithDefaults instantiates a new IncomeVerificationRefreshResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *IncomeVerificationRefreshResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *IncomeVerificationRefreshResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *IncomeVerificationRefreshResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetVerificationRefreshStatus

`func (o *IncomeVerificationRefreshResponse) GetVerificationRefreshStatus() VerificationRefreshStatus`

GetVerificationRefreshStatus returns the VerificationRefreshStatus field if non-nil, zero value otherwise.

### GetVerificationRefreshStatusOk

`func (o *IncomeVerificationRefreshResponse) GetVerificationRefreshStatusOk() (*VerificationRefreshStatus, bool)`

GetVerificationRefreshStatusOk returns a tuple with the VerificationRefreshStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRefreshStatus

`func (o *IncomeVerificationRefreshResponse) SetVerificationRefreshStatus(v VerificationRefreshStatus)`

SetVerificationRefreshStatus sets VerificationRefreshStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


