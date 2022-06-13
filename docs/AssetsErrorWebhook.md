# AssetsErrorWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;ASSETS&#x60; | 
**WebhookCode** | **string** | &#x60;ERROR&#x60; | 
**Error** | [**PlaidError**](PlaidError.md) |  | 
**AssetReportId** | **string** | The ID associated with the Asset Report. | 

## Methods

### NewAssetsErrorWebhook

`func NewAssetsErrorWebhook(webhookType string, webhookCode string, error_ PlaidError, assetReportId string, ) *AssetsErrorWebhook`

NewAssetsErrorWebhook instantiates a new AssetsErrorWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsErrorWebhookWithDefaults

`func NewAssetsErrorWebhookWithDefaults() *AssetsErrorWebhook`

NewAssetsErrorWebhookWithDefaults instantiates a new AssetsErrorWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *AssetsErrorWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *AssetsErrorWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *AssetsErrorWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *AssetsErrorWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *AssetsErrorWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *AssetsErrorWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetError

`func (o *AssetsErrorWebhook) GetError() PlaidError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AssetsErrorWebhook) GetErrorOk() (*PlaidError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AssetsErrorWebhook) SetError(v PlaidError)`

SetError sets Error field to given value.


### GetAssetReportId

`func (o *AssetsErrorWebhook) GetAssetReportId() string`

GetAssetReportId returns the AssetReportId field if non-nil, zero value otherwise.

### GetAssetReportIdOk

`func (o *AssetsErrorWebhook) GetAssetReportIdOk() (*string, bool)`

GetAssetReportIdOk returns a tuple with the AssetReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportId

`func (o *AssetsErrorWebhook) SetAssetReportId(v string)`

SetAssetReportId sets AssetReportId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


