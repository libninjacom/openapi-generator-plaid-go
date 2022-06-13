# AccountBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **NullableFloat32** | The amount of funds available to be withdrawn from the account, as determined by the financial institution.  For &#x60;credit&#x60;-type accounts, the &#x60;available&#x60; balance typically equals the &#x60;limit&#x60; less the &#x60;current&#x60; balance, less any pending outflows plus any pending inflows.  For &#x60;depository&#x60;-type accounts, the &#x60;available&#x60; balance typically equals the &#x60;current&#x60; balance less any pending outflows plus any pending inflows. For &#x60;depository&#x60;-type accounts, the &#x60;available&#x60; balance does not include the overdraft limit.  For &#x60;investment&#x60;-type accounts (or &#x60;brokerage&#x60;-type accounts for API versions 2018-05-22 and earlier), the &#x60;available&#x60; balance is the total cash available to withdraw as presented by the institution.  Note that not all institutions calculate the &#x60;available&#x60;  balance. In the event that &#x60;available&#x60; balance is unavailable, Plaid will return an &#x60;available&#x60; balance value of &#x60;null&#x60;.  Available balance may be cached and is not guaranteed to be up-to-date in realtime unless the value was returned by &#x60;/accounts/balance/get&#x60;.  If &#x60;current&#x60; is &#x60;null&#x60; this field is guaranteed not to be &#x60;null&#x60;. | 
**Current** | **NullableFloat32** | The total amount of funds in or owed by the account.  For &#x60;credit&#x60;-type accounts, a positive balance indicates the amount owed; a negative amount indicates the lender owing the account holder.  For &#x60;loan&#x60;-type accounts, the current balance is the principal remaining on the loan, except in the case of student loan accounts at Sallie Mae (&#x60;ins_116944&#x60;). For Sallie Mae student loans, the account&#39;s balance includes both principal and any outstanding interest.  For &#x60;investment&#x60;-type accounts (or &#x60;brokerage&#x60;-type accounts for API versions 2018-05-22 and earlier), the current balance is the total value of assets as presented by the institution.  Note that balance information may be cached unless the value was returned by &#x60;/accounts/balance/get&#x60;; if the Item is enabled for Transactions, the balance will be at least as recent as the most recent Transaction update. If you require realtime balance information, use the &#x60;available&#x60; balance as provided by &#x60;/accounts/balance/get&#x60;.  When returned by &#x60;/accounts/balance/get&#x60;, this field may be &#x60;null&#x60;. When this happens, &#x60;available&#x60; is guaranteed not to be &#x60;null&#x60;. | 
**Limit** | **NullableFloat32** | For &#x60;credit&#x60;-type accounts, this represents the credit limit.  For &#x60;depository&#x60;-type accounts, this represents the pre-arranged overdraft limit, which is common for current (checking) accounts in Europe.  In North America, this field is typically only available for &#x60;credit&#x60;-type accounts. | 
**IsoCurrencyCode** | **NullableString** | The ISO-4217 currency code of the balance. Always null if &#x60;unofficial_currency_code&#x60; is non-null. | 
**UnofficialCurrencyCode** | **NullableString** | The unofficial currency code associated with the balance. Always null if &#x60;iso_currency_code&#x60; is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;unofficial_currency_code&#x60;s. | 
**LastUpdatedDatetime** | Pointer to **NullableTime** | Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (&#x60;YYYY-MM-DDTHH:mm:ssZ&#x60;) indicating the last time that the balance for the given account has been updated  This is currently only provided when the &#x60;min_last_updated_datetime&#x60; is passed when calling &#x60;/accounts/balance/get&#x60; for &#x60;ins_128026&#x60; (Capital One). | [optional] 

## Methods

### NewAccountBalance

`func NewAccountBalance(available NullableFloat32, current NullableFloat32, limit NullableFloat32, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString, ) *AccountBalance`

NewAccountBalance instantiates a new AccountBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBalanceWithDefaults

`func NewAccountBalanceWithDefaults() *AccountBalance`

NewAccountBalanceWithDefaults instantiates a new AccountBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *AccountBalance) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AccountBalance) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AccountBalance) SetAvailable(v float32)`

SetAvailable sets Available field to given value.


### SetAvailableNil

`func (o *AccountBalance) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *AccountBalance) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetCurrent

`func (o *AccountBalance) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *AccountBalance) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *AccountBalance) SetCurrent(v float32)`

SetCurrent sets Current field to given value.


### SetCurrentNil

`func (o *AccountBalance) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *AccountBalance) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil
### GetLimit

`func (o *AccountBalance) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AccountBalance) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AccountBalance) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *AccountBalance) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *AccountBalance) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetIsoCurrencyCode

`func (o *AccountBalance) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *AccountBalance) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *AccountBalance) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### SetIsoCurrencyCodeNil

`func (o *AccountBalance) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *AccountBalance) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *AccountBalance) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *AccountBalance) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *AccountBalance) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.


### SetUnofficialCurrencyCodeNil

`func (o *AccountBalance) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *AccountBalance) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
### GetLastUpdatedDatetime

`func (o *AccountBalance) GetLastUpdatedDatetime() time.Time`

GetLastUpdatedDatetime returns the LastUpdatedDatetime field if non-nil, zero value otherwise.

### GetLastUpdatedDatetimeOk

`func (o *AccountBalance) GetLastUpdatedDatetimeOk() (*time.Time, bool)`

GetLastUpdatedDatetimeOk returns a tuple with the LastUpdatedDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDatetime

`func (o *AccountBalance) SetLastUpdatedDatetime(v time.Time)`

SetLastUpdatedDatetime sets LastUpdatedDatetime field to given value.

### HasLastUpdatedDatetime

`func (o *AccountBalance) HasLastUpdatedDatetime() bool`

HasLastUpdatedDatetime returns a boolean if a field has been set.

### SetLastUpdatedDatetimeNil

`func (o *AccountBalance) SetLastUpdatedDatetimeNil(b bool)`

 SetLastUpdatedDatetimeNil sets the value for LastUpdatedDatetime to be an explicit nil

### UnsetLastUpdatedDatetime
`func (o *AccountBalance) UnsetLastUpdatedDatetime()`

UnsetLastUpdatedDatetime ensures that no value is present for LastUpdatedDatetime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


