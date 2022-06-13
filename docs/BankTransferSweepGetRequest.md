# BankTransferSweepGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**SweepId** | **string** | Identifier of the sweep. | 

## Methods

### NewBankTransferSweepGetRequest

`func NewBankTransferSweepGetRequest(sweepId string, ) *BankTransferSweepGetRequest`

NewBankTransferSweepGetRequest instantiates a new BankTransferSweepGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferSweepGetRequestWithDefaults

`func NewBankTransferSweepGetRequestWithDefaults() *BankTransferSweepGetRequest`

NewBankTransferSweepGetRequestWithDefaults instantiates a new BankTransferSweepGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *BankTransferSweepGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BankTransferSweepGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BankTransferSweepGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BankTransferSweepGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *BankTransferSweepGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BankTransferSweepGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BankTransferSweepGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *BankTransferSweepGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSweepId

`func (o *BankTransferSweepGetRequest) GetSweepId() string`

GetSweepId returns the SweepId field if non-nil, zero value otherwise.

### GetSweepIdOk

`func (o *BankTransferSweepGetRequest) GetSweepIdOk() (*string, bool)`

GetSweepIdOk returns a tuple with the SweepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweepId

`func (o *BankTransferSweepGetRequest) SetSweepId(v string)`

SetSweepId sets SweepId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


