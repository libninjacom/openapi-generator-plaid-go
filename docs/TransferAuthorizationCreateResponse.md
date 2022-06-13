# TransferAuthorizationCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | [**TransferAuthorization**](TransferAuthorization.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransferAuthorizationCreateResponse

`func NewTransferAuthorizationCreateResponse(authorization TransferAuthorization, requestId string, ) *TransferAuthorizationCreateResponse`

NewTransferAuthorizationCreateResponse instantiates a new TransferAuthorizationCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferAuthorizationCreateResponseWithDefaults

`func NewTransferAuthorizationCreateResponseWithDefaults() *TransferAuthorizationCreateResponse`

NewTransferAuthorizationCreateResponseWithDefaults instantiates a new TransferAuthorizationCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorization

`func (o *TransferAuthorizationCreateResponse) GetAuthorization() TransferAuthorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *TransferAuthorizationCreateResponse) GetAuthorizationOk() (*TransferAuthorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *TransferAuthorizationCreateResponse) SetAuthorization(v TransferAuthorization)`

SetAuthorization sets Authorization field to given value.


### GetRequestId

`func (o *TransferAuthorizationCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransferAuthorizationCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransferAuthorizationCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


