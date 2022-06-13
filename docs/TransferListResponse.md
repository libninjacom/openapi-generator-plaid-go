# TransferListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transfers** | [**[]Transfer**](Transfer.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransferListResponse

`func NewTransferListResponse(transfers []Transfer, requestId string, ) *TransferListResponse`

NewTransferListResponse instantiates a new TransferListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferListResponseWithDefaults

`func NewTransferListResponseWithDefaults() *TransferListResponse`

NewTransferListResponseWithDefaults instantiates a new TransferListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransfers

`func (o *TransferListResponse) GetTransfers() []Transfer`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *TransferListResponse) GetTransfersOk() (*[]Transfer, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *TransferListResponse) SetTransfers(v []Transfer)`

SetTransfers sets Transfers field to given value.


### GetRequestId

`func (o *TransferListResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransferListResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransferListResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


