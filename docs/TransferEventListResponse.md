# TransferEventListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferEvents** | [**[]TransferEvent**](TransferEvent.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransferEventListResponse

`func NewTransferEventListResponse(transferEvents []TransferEvent, requestId string, ) *TransferEventListResponse`

NewTransferEventListResponse instantiates a new TransferEventListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferEventListResponseWithDefaults

`func NewTransferEventListResponseWithDefaults() *TransferEventListResponse`

NewTransferEventListResponseWithDefaults instantiates a new TransferEventListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferEvents

`func (o *TransferEventListResponse) GetTransferEvents() []TransferEvent`

GetTransferEvents returns the TransferEvents field if non-nil, zero value otherwise.

### GetTransferEventsOk

`func (o *TransferEventListResponse) GetTransferEventsOk() (*[]TransferEvent, bool)`

GetTransferEventsOk returns a tuple with the TransferEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferEvents

`func (o *TransferEventListResponse) SetTransferEvents(v []TransferEvent)`

SetTransferEvents sets TransferEvents field to given value.


### GetRequestId

`func (o *TransferEventListResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransferEventListResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransferEventListResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


