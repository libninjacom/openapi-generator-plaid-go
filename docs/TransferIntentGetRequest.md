# TransferIntentGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | 
**Secret** | **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | 
**TransferIntentId** | **string** | Plaid&#39;s unique identifier for a transfer intent object. | 

## Methods

### NewTransferIntentGetRequest

`func NewTransferIntentGetRequest(clientId string, secret string, transferIntentId string, ) *TransferIntentGetRequest`

NewTransferIntentGetRequest instantiates a new TransferIntentGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferIntentGetRequestWithDefaults

`func NewTransferIntentGetRequestWithDefaults() *TransferIntentGetRequest`

NewTransferIntentGetRequestWithDefaults instantiates a new TransferIntentGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TransferIntentGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TransferIntentGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TransferIntentGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetSecret

`func (o *TransferIntentGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TransferIntentGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TransferIntentGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetTransferIntentId

`func (o *TransferIntentGetRequest) GetTransferIntentId() string`

GetTransferIntentId returns the TransferIntentId field if non-nil, zero value otherwise.

### GetTransferIntentIdOk

`func (o *TransferIntentGetRequest) GetTransferIntentIdOk() (*string, bool)`

GetTransferIntentIdOk returns a tuple with the TransferIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferIntentId

`func (o *TransferIntentGetRequest) SetTransferIntentId(v string)`

SetTransferIntentId sets TransferIntentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


