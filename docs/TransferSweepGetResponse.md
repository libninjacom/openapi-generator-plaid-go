# TransferSweepGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sweep** | [**TransferSweep**](TransferSweep.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransferSweepGetResponse

`func NewTransferSweepGetResponse(sweep TransferSweep, requestId string, ) *TransferSweepGetResponse`

NewTransferSweepGetResponse instantiates a new TransferSweepGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferSweepGetResponseWithDefaults

`func NewTransferSweepGetResponseWithDefaults() *TransferSweepGetResponse`

NewTransferSweepGetResponseWithDefaults instantiates a new TransferSweepGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSweep

`func (o *TransferSweepGetResponse) GetSweep() TransferSweep`

GetSweep returns the Sweep field if non-nil, zero value otherwise.

### GetSweepOk

`func (o *TransferSweepGetResponse) GetSweepOk() (*TransferSweep, bool)`

GetSweepOk returns a tuple with the Sweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweep

`func (o *TransferSweepGetResponse) SetSweep(v TransferSweep)`

SetSweep sets Sweep field to given value.


### GetRequestId

`func (o *TransferSweepGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransferSweepGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransferSweepGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


