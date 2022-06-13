# DepositSwitchStateUpdateWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | Pointer to **string** | &#x60;\&quot;DEPOSIT_SWITCH\&quot;&#x60; | [optional] 
**WebhookCode** | Pointer to **string** | &#x60;\&quot;SWITCH_STATE_UPDATE\&quot;&#x60; | [optional] 
**State** | Pointer to **string** |  The state, or status, of the deposit switch.  &#x60;initialized&#x60;: The deposit switch has been initialized with the user entering the information required to submit the deposit switch request.  &#x60;processing&#x60;: The deposit switch request has been submitted and is being processed.  &#x60;completed&#x60;: The user&#39;s employer has fulfilled and completed the deposit switch request.  &#x60;error&#x60;: There was an error processing the deposit switch request.  For more information, see the [Deposit Switch API reference](/docs/deposit-switch/reference#deposit_switchget). | [optional] 
**DepositSwitchId** | Pointer to **string** | The ID of the deposit switch. | [optional] 

## Methods

### NewDepositSwitchStateUpdateWebhook

`func NewDepositSwitchStateUpdateWebhook() *DepositSwitchStateUpdateWebhook`

NewDepositSwitchStateUpdateWebhook instantiates a new DepositSwitchStateUpdateWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchStateUpdateWebhookWithDefaults

`func NewDepositSwitchStateUpdateWebhookWithDefaults() *DepositSwitchStateUpdateWebhook`

NewDepositSwitchStateUpdateWebhookWithDefaults instantiates a new DepositSwitchStateUpdateWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *DepositSwitchStateUpdateWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *DepositSwitchStateUpdateWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *DepositSwitchStateUpdateWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.

### HasWebhookType

`func (o *DepositSwitchStateUpdateWebhook) HasWebhookType() bool`

HasWebhookType returns a boolean if a field has been set.

### GetWebhookCode

`func (o *DepositSwitchStateUpdateWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *DepositSwitchStateUpdateWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *DepositSwitchStateUpdateWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.

### HasWebhookCode

`func (o *DepositSwitchStateUpdateWebhook) HasWebhookCode() bool`

HasWebhookCode returns a boolean if a field has been set.

### GetState

`func (o *DepositSwitchStateUpdateWebhook) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DepositSwitchStateUpdateWebhook) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DepositSwitchStateUpdateWebhook) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DepositSwitchStateUpdateWebhook) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDepositSwitchId

`func (o *DepositSwitchStateUpdateWebhook) GetDepositSwitchId() string`

GetDepositSwitchId returns the DepositSwitchId field if non-nil, zero value otherwise.

### GetDepositSwitchIdOk

`func (o *DepositSwitchStateUpdateWebhook) GetDepositSwitchIdOk() (*string, bool)`

GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositSwitchId

`func (o *DepositSwitchStateUpdateWebhook) SetDepositSwitchId(v string)`

SetDepositSwitchId sets DepositSwitchId field to given value.

### HasDepositSwitchId

`func (o *DepositSwitchStateUpdateWebhook) HasDepositSwitchId() bool`

HasDepositSwitchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


