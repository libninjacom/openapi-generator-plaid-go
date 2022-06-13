# LinkTokenCreateRequestIncomeVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomeVerificationId** | Pointer to **string** | The &#x60;income_verification_id&#x60; of the verification instance, as provided by &#x60;/income/verification/create&#x60;. | [optional] 
**AssetReportId** | Pointer to **string** | The &#x60;asset_report_id&#x60; of an asset report associated with the user, as provided by &#x60;/asset_report/create&#x60;. Providing an &#x60;asset_report_id&#x60; is optional and can be used to verify the user through a streamlined flow. If provided, the bank linking flow will be skipped. | [optional] 
**PrecheckId** | Pointer to **string** | The ID of a precheck created with &#x60;/income/verification/precheck&#x60;. Will be used to improve conversion of the income verification flow by streamlining the Link interface presented to the end user. | [optional] 
**AccessTokens** | Pointer to **[]string** | An array of access tokens corresponding to the Items that will be cross-referenced with the product data. If the &#x60;transactions&#x60; product was not initialized for the Items during link, it will be initialized after this Link session. | [optional] 

## Methods

### NewLinkTokenCreateRequestIncomeVerification

`func NewLinkTokenCreateRequestIncomeVerification() *LinkTokenCreateRequestIncomeVerification`

NewLinkTokenCreateRequestIncomeVerification instantiates a new LinkTokenCreateRequestIncomeVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTokenCreateRequestIncomeVerificationWithDefaults

`func NewLinkTokenCreateRequestIncomeVerificationWithDefaults() *LinkTokenCreateRequestIncomeVerification`

NewLinkTokenCreateRequestIncomeVerificationWithDefaults instantiates a new LinkTokenCreateRequestIncomeVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomeVerificationId

`func (o *LinkTokenCreateRequestIncomeVerification) GetIncomeVerificationId() string`

GetIncomeVerificationId returns the IncomeVerificationId field if non-nil, zero value otherwise.

### GetIncomeVerificationIdOk

`func (o *LinkTokenCreateRequestIncomeVerification) GetIncomeVerificationIdOk() (*string, bool)`

GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeVerificationId

`func (o *LinkTokenCreateRequestIncomeVerification) SetIncomeVerificationId(v string)`

SetIncomeVerificationId sets IncomeVerificationId field to given value.

### HasIncomeVerificationId

`func (o *LinkTokenCreateRequestIncomeVerification) HasIncomeVerificationId() bool`

HasIncomeVerificationId returns a boolean if a field has been set.

### GetAssetReportId

`func (o *LinkTokenCreateRequestIncomeVerification) GetAssetReportId() string`

GetAssetReportId returns the AssetReportId field if non-nil, zero value otherwise.

### GetAssetReportIdOk

`func (o *LinkTokenCreateRequestIncomeVerification) GetAssetReportIdOk() (*string, bool)`

GetAssetReportIdOk returns a tuple with the AssetReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportId

`func (o *LinkTokenCreateRequestIncomeVerification) SetAssetReportId(v string)`

SetAssetReportId sets AssetReportId field to given value.

### HasAssetReportId

`func (o *LinkTokenCreateRequestIncomeVerification) HasAssetReportId() bool`

HasAssetReportId returns a boolean if a field has been set.

### GetPrecheckId

`func (o *LinkTokenCreateRequestIncomeVerification) GetPrecheckId() string`

GetPrecheckId returns the PrecheckId field if non-nil, zero value otherwise.

### GetPrecheckIdOk

`func (o *LinkTokenCreateRequestIncomeVerification) GetPrecheckIdOk() (*string, bool)`

GetPrecheckIdOk returns a tuple with the PrecheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecheckId

`func (o *LinkTokenCreateRequestIncomeVerification) SetPrecheckId(v string)`

SetPrecheckId sets PrecheckId field to given value.

### HasPrecheckId

`func (o *LinkTokenCreateRequestIncomeVerification) HasPrecheckId() bool`

HasPrecheckId returns a boolean if a field has been set.

### GetAccessTokens

`func (o *LinkTokenCreateRequestIncomeVerification) GetAccessTokens() []string`

GetAccessTokens returns the AccessTokens field if non-nil, zero value otherwise.

### GetAccessTokensOk

`func (o *LinkTokenCreateRequestIncomeVerification) GetAccessTokensOk() (*[]string, bool)`

GetAccessTokensOk returns a tuple with the AccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokens

`func (o *LinkTokenCreateRequestIncomeVerification) SetAccessTokens(v []string)`

SetAccessTokens sets AccessTokens field to given value.

### HasAccessTokens

`func (o *LinkTokenCreateRequestIncomeVerification) HasAccessTokens() bool`

HasAccessTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


