# TransferCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transfer** | [**Transfer**](Transfer.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransferCreateResponse

`func NewTransferCreateResponse(transfer Transfer, requestId string, ) *TransferCreateResponse`

NewTransferCreateResponse instantiates a new TransferCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferCreateResponseWithDefaults

`func NewTransferCreateResponseWithDefaults() *TransferCreateResponse`

NewTransferCreateResponseWithDefaults instantiates a new TransferCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransfer

`func (o *TransferCreateResponse) GetTransfer() Transfer`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *TransferCreateResponse) GetTransferOk() (*Transfer, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *TransferCreateResponse) SetTransfer(v Transfer)`

SetTransfer sets Transfer field to given value.


### GetRequestId

`func (o *TransferCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransferCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransferCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


