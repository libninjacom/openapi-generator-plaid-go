# NewAccountsAvailableWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | Pointer to **string** | &#x60;ITEM&#x60; | [optional] 
**WebhookCode** | Pointer to **string** | &#x60;NEW_ACCOUNTS_AVAILABLE&#x60; | [optional] 
**ItemId** | Pointer to **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | [optional] 
**Error** | Pointer to [**PlaidError**](PlaidError.md) |  | [optional] 

## Methods

### NewNewAccountsAvailableWebhook

`func NewNewAccountsAvailableWebhook() *NewAccountsAvailableWebhook`

NewNewAccountsAvailableWebhook instantiates a new NewAccountsAvailableWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAccountsAvailableWebhookWithDefaults

`func NewNewAccountsAvailableWebhookWithDefaults() *NewAccountsAvailableWebhook`

NewNewAccountsAvailableWebhookWithDefaults instantiates a new NewAccountsAvailableWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *NewAccountsAvailableWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *NewAccountsAvailableWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *NewAccountsAvailableWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.

### HasWebhookType

`func (o *NewAccountsAvailableWebhook) HasWebhookType() bool`

HasWebhookType returns a boolean if a field has been set.

### GetWebhookCode

`func (o *NewAccountsAvailableWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *NewAccountsAvailableWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *NewAccountsAvailableWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.

### HasWebhookCode

`func (o *NewAccountsAvailableWebhook) HasWebhookCode() bool`

HasWebhookCode returns a boolean if a field has been set.

### GetItemId

`func (o *NewAccountsAvailableWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *NewAccountsAvailableWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *NewAccountsAvailableWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *NewAccountsAvailableWebhook) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetError

`func (o *NewAccountsAvailableWebhook) GetError() PlaidError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NewAccountsAvailableWebhook) GetErrorOk() (*PlaidError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NewAccountsAvailableWebhook) SetError(v PlaidError)`

SetError sets Error field to given value.

### HasError

`func (o *NewAccountsAvailableWebhook) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


