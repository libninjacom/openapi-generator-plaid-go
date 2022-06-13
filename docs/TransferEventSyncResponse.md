# TransferEventSyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferEvents** | [**[]TransferEvent**](TransferEvent.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransferEventSyncResponse

`func NewTransferEventSyncResponse(transferEvents []TransferEvent, requestId string, ) *TransferEventSyncResponse`

NewTransferEventSyncResponse instantiates a new TransferEventSyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferEventSyncResponseWithDefaults

`func NewTransferEventSyncResponseWithDefaults() *TransferEventSyncResponse`

NewTransferEventSyncResponseWithDefaults instantiates a new TransferEventSyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferEvents

`func (o *TransferEventSyncResponse) GetTransferEvents() []TransferEvent`

GetTransferEvents returns the TransferEvents field if non-nil, zero value otherwise.

### GetTransferEventsOk

`func (o *TransferEventSyncResponse) GetTransferEventsOk() (*[]TransferEvent, bool)`

GetTransferEventsOk returns a tuple with the TransferEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferEvents

`func (o *TransferEventSyncResponse) SetTransferEvents(v []TransferEvent)`

SetTransferEvents sets TransferEvents field to given value.


### GetRequestId

`func (o *TransferEventSyncResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransferEventSyncResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransferEventSyncResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


