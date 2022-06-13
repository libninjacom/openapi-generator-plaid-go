# DepositSwitchCreateRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Webhook** | Pointer to **NullableString** | The URL registered to receive webhooks when the status of a deposit switch request has changed.  | [optional] 
**TransactionItemAccessTokens** | Pointer to **[]string** | An array of access tokens corresponding to transaction items to use when attempting to match the user to their Payroll Provider. These tokens must be created by the same client id as the one creating the switch, and have access to the transactions product. | [optional] 

## Methods

### NewDepositSwitchCreateRequestOptions

`func NewDepositSwitchCreateRequestOptions() *DepositSwitchCreateRequestOptions`

NewDepositSwitchCreateRequestOptions instantiates a new DepositSwitchCreateRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchCreateRequestOptionsWithDefaults

`func NewDepositSwitchCreateRequestOptionsWithDefaults() *DepositSwitchCreateRequestOptions`

NewDepositSwitchCreateRequestOptionsWithDefaults instantiates a new DepositSwitchCreateRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhook

`func (o *DepositSwitchCreateRequestOptions) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *DepositSwitchCreateRequestOptions) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *DepositSwitchCreateRequestOptions) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *DepositSwitchCreateRequestOptions) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### SetWebhookNil

`func (o *DepositSwitchCreateRequestOptions) SetWebhookNil(b bool)`

 SetWebhookNil sets the value for Webhook to be an explicit nil

### UnsetWebhook
`func (o *DepositSwitchCreateRequestOptions) UnsetWebhook()`

UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
### GetTransactionItemAccessTokens

`func (o *DepositSwitchCreateRequestOptions) GetTransactionItemAccessTokens() []string`

GetTransactionItemAccessTokens returns the TransactionItemAccessTokens field if non-nil, zero value otherwise.

### GetTransactionItemAccessTokensOk

`func (o *DepositSwitchCreateRequestOptions) GetTransactionItemAccessTokensOk() (*[]string, bool)`

GetTransactionItemAccessTokensOk returns a tuple with the TransactionItemAccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionItemAccessTokens

`func (o *DepositSwitchCreateRequestOptions) SetTransactionItemAccessTokens(v []string)`

SetTransactionItemAccessTokens sets TransactionItemAccessTokens field to given value.

### HasTransactionItemAccessTokens

`func (o *DepositSwitchCreateRequestOptions) HasTransactionItemAccessTokens() bool`

HasTransactionItemAccessTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


