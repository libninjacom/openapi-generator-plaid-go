# TransferIntentGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferIntent** | [**TransferIntentGet**](TransferIntentGet.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransferIntentGetResponse

`func NewTransferIntentGetResponse(transferIntent TransferIntentGet, requestId string, ) *TransferIntentGetResponse`

NewTransferIntentGetResponse instantiates a new TransferIntentGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferIntentGetResponseWithDefaults

`func NewTransferIntentGetResponseWithDefaults() *TransferIntentGetResponse`

NewTransferIntentGetResponseWithDefaults instantiates a new TransferIntentGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferIntent

`func (o *TransferIntentGetResponse) GetTransferIntent() TransferIntentGet`

GetTransferIntent returns the TransferIntent field if non-nil, zero value otherwise.

### GetTransferIntentOk

`func (o *TransferIntentGetResponse) GetTransferIntentOk() (*TransferIntentGet, bool)`

GetTransferIntentOk returns a tuple with the TransferIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferIntent

`func (o *TransferIntentGetResponse) SetTransferIntent(v TransferIntentGet)`

SetTransferIntent sets TransferIntent field to given value.


### GetRequestId

`func (o *TransferIntentGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransferIntentGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransferIntentGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


