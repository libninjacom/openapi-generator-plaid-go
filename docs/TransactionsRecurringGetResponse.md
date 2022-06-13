# TransactionsRecurringGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InflowStreams** | [**[]TransactionStream**](TransactionStream.md) | An array of depository transaction streams. | 
**OutflowStreams** | [**[]TransactionStream**](TransactionStream.md) | An array of expense transaction streams. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransactionsRecurringGetResponse

`func NewTransactionsRecurringGetResponse(inflowStreams []TransactionStream, outflowStreams []TransactionStream, requestId string, ) *TransactionsRecurringGetResponse`

NewTransactionsRecurringGetResponse instantiates a new TransactionsRecurringGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsRecurringGetResponseWithDefaults

`func NewTransactionsRecurringGetResponseWithDefaults() *TransactionsRecurringGetResponse`

NewTransactionsRecurringGetResponseWithDefaults instantiates a new TransactionsRecurringGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInflowStreams

`func (o *TransactionsRecurringGetResponse) GetInflowStreams() []TransactionStream`

GetInflowStreams returns the InflowStreams field if non-nil, zero value otherwise.

### GetInflowStreamsOk

`func (o *TransactionsRecurringGetResponse) GetInflowStreamsOk() (*[]TransactionStream, bool)`

GetInflowStreamsOk returns a tuple with the InflowStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflowStreams

`func (o *TransactionsRecurringGetResponse) SetInflowStreams(v []TransactionStream)`

SetInflowStreams sets InflowStreams field to given value.


### GetOutflowStreams

`func (o *TransactionsRecurringGetResponse) GetOutflowStreams() []TransactionStream`

GetOutflowStreams returns the OutflowStreams field if non-nil, zero value otherwise.

### GetOutflowStreamsOk

`func (o *TransactionsRecurringGetResponse) GetOutflowStreamsOk() (*[]TransactionStream, bool)`

GetOutflowStreamsOk returns a tuple with the OutflowStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutflowStreams

`func (o *TransactionsRecurringGetResponse) SetOutflowStreams(v []TransactionStream)`

SetOutflowStreams sets OutflowStreams field to given value.


### GetRequestId

`func (o *TransactionsRecurringGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransactionsRecurringGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransactionsRecurringGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


