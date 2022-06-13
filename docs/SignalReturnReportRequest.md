# SignalReturnReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**ClientTransactionId** | **string** | Must be the same as the &#x60;client_transaction_id&#x60; supplied when calling &#x60;/signal/evaluate&#x60; | 
**ReturnCode** | **string** | Must be a valid ACH return code (e.g. \&quot;R01\&quot;) | 

## Methods

### NewSignalReturnReportRequest

`func NewSignalReturnReportRequest(clientTransactionId string, returnCode string, ) *SignalReturnReportRequest`

NewSignalReturnReportRequest instantiates a new SignalReturnReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalReturnReportRequestWithDefaults

`func NewSignalReturnReportRequestWithDefaults() *SignalReturnReportRequest`

NewSignalReturnReportRequestWithDefaults instantiates a new SignalReturnReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SignalReturnReportRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SignalReturnReportRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SignalReturnReportRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SignalReturnReportRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *SignalReturnReportRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SignalReturnReportRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SignalReturnReportRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SignalReturnReportRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetClientTransactionId

`func (o *SignalReturnReportRequest) GetClientTransactionId() string`

GetClientTransactionId returns the ClientTransactionId field if non-nil, zero value otherwise.

### GetClientTransactionIdOk

`func (o *SignalReturnReportRequest) GetClientTransactionIdOk() (*string, bool)`

GetClientTransactionIdOk returns a tuple with the ClientTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTransactionId

`func (o *SignalReturnReportRequest) SetClientTransactionId(v string)`

SetClientTransactionId sets ClientTransactionId field to given value.


### GetReturnCode

`func (o *SignalReturnReportRequest) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *SignalReturnReportRequest) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *SignalReturnReportRequest) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


