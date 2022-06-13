# IncomeVerificationStatusWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;\&quot;INCOME\&quot;&#x60; | 
**WebhookCode** | **string** | &#x60;income_verification&#x60; | 
**IncomeVerificationId** | **string** | The &#x60;income_verification_id&#x60; of the verification instance whose status is being reported. | 
**ItemId** | **string** | The Item ID associated with the verification. | 
**VerificationStatus** | **string** | &#x60;VERIFICATION_STATUS_PROCESSING_COMPLETE&#x60;: The income verification status processing has completed. If the user uploaded multiple documents, this webhook will fire when all documents have finished processing. Call the &#x60;/income/verification/paystubs/get&#x60; endpoint and check the document metadata to see which documents were successfully parsed.  &#x60;VERIFICATION_STATUS_PROCESSING_FAILED&#x60;: A failure occurred when attempting to process the verification documentation.  &#x60;VERIFICATION_STATUS_PENDING_APPROVAL&#x60;: The income verification has been sent to the user for review. | 

## Methods

### NewIncomeVerificationStatusWebhook

`func NewIncomeVerificationStatusWebhook(webhookType string, webhookCode string, incomeVerificationId string, itemId string, verificationStatus string, ) *IncomeVerificationStatusWebhook`

NewIncomeVerificationStatusWebhook instantiates a new IncomeVerificationStatusWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationStatusWebhookWithDefaults

`func NewIncomeVerificationStatusWebhookWithDefaults() *IncomeVerificationStatusWebhook`

NewIncomeVerificationStatusWebhookWithDefaults instantiates a new IncomeVerificationStatusWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *IncomeVerificationStatusWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *IncomeVerificationStatusWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *IncomeVerificationStatusWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *IncomeVerificationStatusWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *IncomeVerificationStatusWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *IncomeVerificationStatusWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetIncomeVerificationId

`func (o *IncomeVerificationStatusWebhook) GetIncomeVerificationId() string`

GetIncomeVerificationId returns the IncomeVerificationId field if non-nil, zero value otherwise.

### GetIncomeVerificationIdOk

`func (o *IncomeVerificationStatusWebhook) GetIncomeVerificationIdOk() (*string, bool)`

GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeVerificationId

`func (o *IncomeVerificationStatusWebhook) SetIncomeVerificationId(v string)`

SetIncomeVerificationId sets IncomeVerificationId field to given value.


### GetItemId

`func (o *IncomeVerificationStatusWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *IncomeVerificationStatusWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *IncomeVerificationStatusWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetVerificationStatus

`func (o *IncomeVerificationStatusWebhook) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *IncomeVerificationStatusWebhook) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *IncomeVerificationStatusWebhook) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


