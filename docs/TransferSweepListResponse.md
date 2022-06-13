# TransferSweepListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sweeps** | [**[]TransferSweep**](TransferSweep.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransferSweepListResponse

`func NewTransferSweepListResponse(sweeps []TransferSweep, requestId string, ) *TransferSweepListResponse`

NewTransferSweepListResponse instantiates a new TransferSweepListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferSweepListResponseWithDefaults

`func NewTransferSweepListResponseWithDefaults() *TransferSweepListResponse`

NewTransferSweepListResponseWithDefaults instantiates a new TransferSweepListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSweeps

`func (o *TransferSweepListResponse) GetSweeps() []TransferSweep`

GetSweeps returns the Sweeps field if non-nil, zero value otherwise.

### GetSweepsOk

`func (o *TransferSweepListResponse) GetSweepsOk() (*[]TransferSweep, bool)`

GetSweepsOk returns a tuple with the Sweeps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweeps

`func (o *TransferSweepListResponse) SetSweeps(v []TransferSweep)`

SetSweeps sets Sweeps field to given value.


### GetRequestId

`func (o *TransferSweepListResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransferSweepListResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransferSweepListResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


