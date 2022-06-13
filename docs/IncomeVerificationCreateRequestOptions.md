# IncomeVerificationCreateRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokens** | Pointer to **[]string** | An array of access tokens corresponding to the Items that will be cross-referenced with the product data. Plaid will attempt to correlate transaction history from these Items with data from the user&#39;s paystub, such as date and amount. The &#x60;verification&#x60; status of the paystub as returned by &#x60;/income/verification/paystubs/get&#x60; will indicate if the verification status was successful, or, if not, why it failed. If the &#x60;transactions&#x60; product was not initialized for the Items during Link, it will be initialized after this Link session. | [optional] 

## Methods

### NewIncomeVerificationCreateRequestOptions

`func NewIncomeVerificationCreateRequestOptions() *IncomeVerificationCreateRequestOptions`

NewIncomeVerificationCreateRequestOptions instantiates a new IncomeVerificationCreateRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationCreateRequestOptionsWithDefaults

`func NewIncomeVerificationCreateRequestOptionsWithDefaults() *IncomeVerificationCreateRequestOptions`

NewIncomeVerificationCreateRequestOptionsWithDefaults instantiates a new IncomeVerificationCreateRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokens

`func (o *IncomeVerificationCreateRequestOptions) GetAccessTokens() []string`

GetAccessTokens returns the AccessTokens field if non-nil, zero value otherwise.

### GetAccessTokensOk

`func (o *IncomeVerificationCreateRequestOptions) GetAccessTokensOk() (*[]string, bool)`

GetAccessTokensOk returns a tuple with the AccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokens

`func (o *IncomeVerificationCreateRequestOptions) SetAccessTokens(v []string)`

SetAccessTokens sets AccessTokens field to given value.

### HasAccessTokens

`func (o *IncomeVerificationCreateRequestOptions) HasAccessTokens() bool`

HasAccessTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


