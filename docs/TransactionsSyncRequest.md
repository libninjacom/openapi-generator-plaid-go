# TransactionsSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**Cursor** | Pointer to **string** | The cursor value represents the last update requested. Providing it will cause the response to only return changes after this update. If omitted, the entire history of updates will be returned, starting with the first-added transactions on the item. Note: The upper-bound length of this cursor is 256 characters of base64. | [optional] 
**Count** | Pointer to **int32** | The number of transaction updates to fetch. | [optional] [default to 100]

## Methods

### NewTransactionsSyncRequest

`func NewTransactionsSyncRequest(accessToken string, ) *TransactionsSyncRequest`

NewTransactionsSyncRequest instantiates a new TransactionsSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsSyncRequestWithDefaults

`func NewTransactionsSyncRequestWithDefaults() *TransactionsSyncRequest`

NewTransactionsSyncRequestWithDefaults instantiates a new TransactionsSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TransactionsSyncRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TransactionsSyncRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TransactionsSyncRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TransactionsSyncRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAccessToken

`func (o *TransactionsSyncRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TransactionsSyncRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TransactionsSyncRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetSecret

`func (o *TransactionsSyncRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TransactionsSyncRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TransactionsSyncRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *TransactionsSyncRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetCursor

`func (o *TransactionsSyncRequest) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *TransactionsSyncRequest) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *TransactionsSyncRequest) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *TransactionsSyncRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetCount

`func (o *TransactionsSyncRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TransactionsSyncRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TransactionsSyncRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TransactionsSyncRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


