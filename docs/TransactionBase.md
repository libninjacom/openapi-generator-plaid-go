# TransactionBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionType** | Pointer to **string** | Please use the &#x60;payment_channel&#x60; field, &#x60;transaction_type&#x60; will be deprecated in the future.  &#x60;digital:&#x60; transactions that took place online.  &#x60;place:&#x60; transactions that were made at a physical location.  &#x60;special:&#x60; transactions that relate to banks, e.g. fees or deposits.  &#x60;unresolved:&#x60; transactions that do not fit into the other three types.  | [optional] 
**PendingTransactionId** | Pointer to **NullableString** | The ID of a posted transaction&#39;s associated pending transaction, where applicable. | [optional] 
**CategoryId** | Pointer to **NullableString** | The ID of the category to which this transaction belongs. For a full list of categories, see [&#x60;/categories/get&#x60;](https://plaid.com/docs/api/products/#categoriesget).  If the &#x60;transactions&#x60; object was returned by an Assets endpoint such as &#x60;/asset_report/get/&#x60; or &#x60;/asset_report/pdf/get&#x60;, this field will only appear in an Asset Report with Insights. | [optional] 
**Category** | Pointer to **[]string** | A hierarchical array of the categories to which this transaction belongs. For a full list of categories, see [&#x60;/categories/get&#x60;](https://plaid.com/docs/api/products/#categoriesget).  If the &#x60;transactions&#x60; object was returned by an Assets endpoint such as &#x60;/asset_report/get/&#x60; or &#x60;/asset_report/pdf/get&#x60;, this field will only appear in an Asset Report with Insights. | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**PaymentMeta** | Pointer to [**PaymentMeta**](PaymentMeta.md) |  | [optional] 
**AccountOwner** | Pointer to **NullableString** | The name of the account owner. This field is not typically populated and only relevant when dealing with sub-accounts. | [optional] 
**Name** | Pointer to **string** | The merchant name or transaction description.  If the &#x60;transactions&#x60; object was returned by a Transactions endpoint such as &#x60;/transactions/get&#x60;, this field will always appear. If the &#x60;transactions&#x60; object was returned by an Assets endpoint such as &#x60;/asset_report/get/&#x60; or &#x60;/asset_report/pdf/get&#x60;, this field will only appear in an Asset Report with Insights. | [optional] 
**OriginalDescription** | Pointer to **NullableString** | The string returned by the financial institution to describe the transaction. For transactions returned by &#x60;/transactions/get&#x60;, this field is in beta and will be omitted unless the client is both enrolled in the closed beta program and has set &#x60;options.include_original_description&#x60; to &#x60;true&#x60;. | [optional] 
**AccountId** | **string** | The ID of the account in which this transaction occurred. | 
**Amount** | **float32** | The settled value of the transaction, denominated in the account&#39;s currency, as stated in &#x60;iso_currency_code&#x60; or &#x60;unofficial_currency_code&#x60;. Positive values when money moves out of the account; negative values when money moves in. For example, debit card purchases are positive; credit card payments, direct deposits, and refunds are negative. | 
**IsoCurrencyCode** | **NullableString** | The ISO-4217 currency code of the transaction. Always &#x60;null&#x60; if &#x60;unofficial_currency_code&#x60; is non-null. | 
**UnofficialCurrencyCode** | **NullableString** | The unofficial currency code associated with the transaction. Always &#x60;null&#x60; if &#x60;iso_currency_code&#x60; is non-&#x60;null&#x60;. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;iso_currency_code&#x60;s. | 
**Date** | **string** | For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted. Both dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( &#x60;YYYY-MM-DD&#x60; ). | 
**Pending** | **bool** | When &#x60;true&#x60;, identifies the transaction as pending or unsettled. Pending transaction details (name, type, amount, category ID) may change before they are settled. | 
**TransactionId** | **string** | The unique ID of the transaction. Like all Plaid identifiers, the &#x60;transaction_id&#x60; is case sensitive. | 
**MerchantName** | Pointer to **NullableString** | The merchant name, as extracted by Plaid from the &#x60;name&#x60; field. | [optional] 
**CheckNumber** | Pointer to **NullableString** | The check number of the transaction. This field is only populated for check transactions. | [optional] 

## Methods

### NewTransactionBase

`func NewTransactionBase(accountId string, amount float32, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString, date string, pending bool, transactionId string, ) *TransactionBase`

NewTransactionBase instantiates a new TransactionBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionBaseWithDefaults

`func NewTransactionBaseWithDefaults() *TransactionBase`

NewTransactionBaseWithDefaults instantiates a new TransactionBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionType

`func (o *TransactionBase) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *TransactionBase) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *TransactionBase) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *TransactionBase) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetPendingTransactionId

`func (o *TransactionBase) GetPendingTransactionId() string`

GetPendingTransactionId returns the PendingTransactionId field if non-nil, zero value otherwise.

### GetPendingTransactionIdOk

`func (o *TransactionBase) GetPendingTransactionIdOk() (*string, bool)`

GetPendingTransactionIdOk returns a tuple with the PendingTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingTransactionId

`func (o *TransactionBase) SetPendingTransactionId(v string)`

SetPendingTransactionId sets PendingTransactionId field to given value.

### HasPendingTransactionId

`func (o *TransactionBase) HasPendingTransactionId() bool`

HasPendingTransactionId returns a boolean if a field has been set.

### SetPendingTransactionIdNil

`func (o *TransactionBase) SetPendingTransactionIdNil(b bool)`

 SetPendingTransactionIdNil sets the value for PendingTransactionId to be an explicit nil

### UnsetPendingTransactionId
`func (o *TransactionBase) UnsetPendingTransactionId()`

UnsetPendingTransactionId ensures that no value is present for PendingTransactionId, not even an explicit nil
### GetCategoryId

`func (o *TransactionBase) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *TransactionBase) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *TransactionBase) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *TransactionBase) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### SetCategoryIdNil

`func (o *TransactionBase) SetCategoryIdNil(b bool)`

 SetCategoryIdNil sets the value for CategoryId to be an explicit nil

### UnsetCategoryId
`func (o *TransactionBase) UnsetCategoryId()`

UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
### GetCategory

`func (o *TransactionBase) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TransactionBase) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TransactionBase) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TransactionBase) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *TransactionBase) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *TransactionBase) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetLocation

`func (o *TransactionBase) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TransactionBase) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TransactionBase) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TransactionBase) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPaymentMeta

`func (o *TransactionBase) GetPaymentMeta() PaymentMeta`

GetPaymentMeta returns the PaymentMeta field if non-nil, zero value otherwise.

### GetPaymentMetaOk

`func (o *TransactionBase) GetPaymentMetaOk() (*PaymentMeta, bool)`

GetPaymentMetaOk returns a tuple with the PaymentMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMeta

`func (o *TransactionBase) SetPaymentMeta(v PaymentMeta)`

SetPaymentMeta sets PaymentMeta field to given value.

### HasPaymentMeta

`func (o *TransactionBase) HasPaymentMeta() bool`

HasPaymentMeta returns a boolean if a field has been set.

### GetAccountOwner

`func (o *TransactionBase) GetAccountOwner() string`

GetAccountOwner returns the AccountOwner field if non-nil, zero value otherwise.

### GetAccountOwnerOk

`func (o *TransactionBase) GetAccountOwnerOk() (*string, bool)`

GetAccountOwnerOk returns a tuple with the AccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOwner

`func (o *TransactionBase) SetAccountOwner(v string)`

SetAccountOwner sets AccountOwner field to given value.

### HasAccountOwner

`func (o *TransactionBase) HasAccountOwner() bool`

HasAccountOwner returns a boolean if a field has been set.

### SetAccountOwnerNil

`func (o *TransactionBase) SetAccountOwnerNil(b bool)`

 SetAccountOwnerNil sets the value for AccountOwner to be an explicit nil

### UnsetAccountOwner
`func (o *TransactionBase) UnsetAccountOwner()`

UnsetAccountOwner ensures that no value is present for AccountOwner, not even an explicit nil
### GetName

`func (o *TransactionBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransactionBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransactionBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransactionBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginalDescription

`func (o *TransactionBase) GetOriginalDescription() string`

GetOriginalDescription returns the OriginalDescription field if non-nil, zero value otherwise.

### GetOriginalDescriptionOk

`func (o *TransactionBase) GetOriginalDescriptionOk() (*string, bool)`

GetOriginalDescriptionOk returns a tuple with the OriginalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDescription

`func (o *TransactionBase) SetOriginalDescription(v string)`

SetOriginalDescription sets OriginalDescription field to given value.

### HasOriginalDescription

`func (o *TransactionBase) HasOriginalDescription() bool`

HasOriginalDescription returns a boolean if a field has been set.

### SetOriginalDescriptionNil

`func (o *TransactionBase) SetOriginalDescriptionNil(b bool)`

 SetOriginalDescriptionNil sets the value for OriginalDescription to be an explicit nil

### UnsetOriginalDescription
`func (o *TransactionBase) UnsetOriginalDescription()`

UnsetOriginalDescription ensures that no value is present for OriginalDescription, not even an explicit nil
### GetAccountId

`func (o *TransactionBase) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransactionBase) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransactionBase) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *TransactionBase) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionBase) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionBase) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetIsoCurrencyCode

`func (o *TransactionBase) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *TransactionBase) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *TransactionBase) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### SetIsoCurrencyCodeNil

`func (o *TransactionBase) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *TransactionBase) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *TransactionBase) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *TransactionBase) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *TransactionBase) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.


### SetUnofficialCurrencyCodeNil

`func (o *TransactionBase) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *TransactionBase) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
### GetDate

`func (o *TransactionBase) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionBase) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionBase) SetDate(v string)`

SetDate sets Date field to given value.


### GetPending

`func (o *TransactionBase) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *TransactionBase) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *TransactionBase) SetPending(v bool)`

SetPending sets Pending field to given value.


### GetTransactionId

`func (o *TransactionBase) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransactionBase) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransactionBase) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetMerchantName

`func (o *TransactionBase) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *TransactionBase) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *TransactionBase) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *TransactionBase) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### SetMerchantNameNil

`func (o *TransactionBase) SetMerchantNameNil(b bool)`

 SetMerchantNameNil sets the value for MerchantName to be an explicit nil

### UnsetMerchantName
`func (o *TransactionBase) UnsetMerchantName()`

UnsetMerchantName ensures that no value is present for MerchantName, not even an explicit nil
### GetCheckNumber

`func (o *TransactionBase) GetCheckNumber() string`

GetCheckNumber returns the CheckNumber field if non-nil, zero value otherwise.

### GetCheckNumberOk

`func (o *TransactionBase) GetCheckNumberOk() (*string, bool)`

GetCheckNumberOk returns a tuple with the CheckNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNumber

`func (o *TransactionBase) SetCheckNumber(v string)`

SetCheckNumber sets CheckNumber field to given value.

### HasCheckNumber

`func (o *TransactionBase) HasCheckNumber() bool`

HasCheckNumber returns a boolean if a field has been set.

### SetCheckNumberNil

`func (o *TransactionBase) SetCheckNumberNil(b bool)`

 SetCheckNumberNil sets the value for CheckNumber to be an explicit nil

### UnsetCheckNumber
`func (o *TransactionBase) UnsetCheckNumber()`

UnsetCheckNumber ensures that no value is present for CheckNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


