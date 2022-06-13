# LiabilitiesDefaultUpdateWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;LIABILITIES&#x60; | 
**WebhookCode** | **string** | &#x60;DEFAULT_UPDATE&#x60; | 
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 
**Error** | [**PlaidError**](PlaidError.md) |  | 
**AccountIdsWithNewLiabilities** | **[]string** | An array of &#x60;account_id&#x60;&#39;s for accounts that contain new liabilities.&#39; | 
**AccountIdsWithUpdatedLiabilities** | **map[string][]string** | An object with keys of &#x60;account_id&#x60;&#39;s that are mapped to their respective liabilities fields that changed.  Example: &#x60;{ \&quot;XMBvvyMGQ1UoLbKByoMqH3nXMj84ALSdE5B58\&quot;: [\&quot;past_amount_due\&quot;] }&#x60;  | 

## Methods

### NewLiabilitiesDefaultUpdateWebhook

`func NewLiabilitiesDefaultUpdateWebhook(webhookType string, webhookCode string, itemId string, error_ PlaidError, accountIdsWithNewLiabilities []string, accountIdsWithUpdatedLiabilities map[string][]string, ) *LiabilitiesDefaultUpdateWebhook`

NewLiabilitiesDefaultUpdateWebhook instantiates a new LiabilitiesDefaultUpdateWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiabilitiesDefaultUpdateWebhookWithDefaults

`func NewLiabilitiesDefaultUpdateWebhookWithDefaults() *LiabilitiesDefaultUpdateWebhook`

NewLiabilitiesDefaultUpdateWebhookWithDefaults instantiates a new LiabilitiesDefaultUpdateWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *LiabilitiesDefaultUpdateWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *LiabilitiesDefaultUpdateWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *LiabilitiesDefaultUpdateWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *LiabilitiesDefaultUpdateWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *LiabilitiesDefaultUpdateWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *LiabilitiesDefaultUpdateWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetItemId

`func (o *LiabilitiesDefaultUpdateWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *LiabilitiesDefaultUpdateWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *LiabilitiesDefaultUpdateWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetError

`func (o *LiabilitiesDefaultUpdateWebhook) GetError() PlaidError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LiabilitiesDefaultUpdateWebhook) GetErrorOk() (*PlaidError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LiabilitiesDefaultUpdateWebhook) SetError(v PlaidError)`

SetError sets Error field to given value.


### GetAccountIdsWithNewLiabilities

`func (o *LiabilitiesDefaultUpdateWebhook) GetAccountIdsWithNewLiabilities() []string`

GetAccountIdsWithNewLiabilities returns the AccountIdsWithNewLiabilities field if non-nil, zero value otherwise.

### GetAccountIdsWithNewLiabilitiesOk

`func (o *LiabilitiesDefaultUpdateWebhook) GetAccountIdsWithNewLiabilitiesOk() (*[]string, bool)`

GetAccountIdsWithNewLiabilitiesOk returns a tuple with the AccountIdsWithNewLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdsWithNewLiabilities

`func (o *LiabilitiesDefaultUpdateWebhook) SetAccountIdsWithNewLiabilities(v []string)`

SetAccountIdsWithNewLiabilities sets AccountIdsWithNewLiabilities field to given value.


### GetAccountIdsWithUpdatedLiabilities

`func (o *LiabilitiesDefaultUpdateWebhook) GetAccountIdsWithUpdatedLiabilities() map[string][]string`

GetAccountIdsWithUpdatedLiabilities returns the AccountIdsWithUpdatedLiabilities field if non-nil, zero value otherwise.

### GetAccountIdsWithUpdatedLiabilitiesOk

`func (o *LiabilitiesDefaultUpdateWebhook) GetAccountIdsWithUpdatedLiabilitiesOk() (*map[string][]string, bool)`

GetAccountIdsWithUpdatedLiabilitiesOk returns a tuple with the AccountIdsWithUpdatedLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdsWithUpdatedLiabilities

`func (o *LiabilitiesDefaultUpdateWebhook) SetAccountIdsWithUpdatedLiabilities(v map[string][]string)`

SetAccountIdsWithUpdatedLiabilities sets AccountIdsWithUpdatedLiabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


