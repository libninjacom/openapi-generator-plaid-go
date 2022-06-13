# EarningsTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAmount** | Pointer to **NullableFloat32** | Total amount of the earnings for this pay period | [optional] 
**CurrentPay** | Pointer to [**Pay**](Pay.md) |  | [optional] 
**YtdPay** | Pointer to [**Pay**](Pay.md) |  | [optional] 
**Hours** | Pointer to **NullableFloat32** | Total number of hours worked for this pay period | [optional] 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the line item. Always &#x60;null&#x60; if &#x60;unofficial_currency_code&#x60; is non-null. | [optional] 
**UnofficialCurrencyCode** | Pointer to **NullableString** | The unofficial currency code associated with the security. Always &#x60;null&#x60; if &#x60;iso_currency_code&#x60; is non-&#x60;null&#x60;. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;iso_currency_code&#x60;s. | [optional] 
**YtdAmount** | Pointer to **NullableFloat32** | The total year-to-date amount of the earnings | [optional] 

## Methods

### NewEarningsTotal

`func NewEarningsTotal() *EarningsTotal`

NewEarningsTotal instantiates a new EarningsTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarningsTotalWithDefaults

`func NewEarningsTotalWithDefaults() *EarningsTotal`

NewEarningsTotalWithDefaults instantiates a new EarningsTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentAmount

`func (o *EarningsTotal) GetCurrentAmount() float32`

GetCurrentAmount returns the CurrentAmount field if non-nil, zero value otherwise.

### GetCurrentAmountOk

`func (o *EarningsTotal) GetCurrentAmountOk() (*float32, bool)`

GetCurrentAmountOk returns a tuple with the CurrentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmount

`func (o *EarningsTotal) SetCurrentAmount(v float32)`

SetCurrentAmount sets CurrentAmount field to given value.

### HasCurrentAmount

`func (o *EarningsTotal) HasCurrentAmount() bool`

HasCurrentAmount returns a boolean if a field has been set.

### SetCurrentAmountNil

`func (o *EarningsTotal) SetCurrentAmountNil(b bool)`

 SetCurrentAmountNil sets the value for CurrentAmount to be an explicit nil

### UnsetCurrentAmount
`func (o *EarningsTotal) UnsetCurrentAmount()`

UnsetCurrentAmount ensures that no value is present for CurrentAmount, not even an explicit nil
### GetCurrentPay

`func (o *EarningsTotal) GetCurrentPay() Pay`

GetCurrentPay returns the CurrentPay field if non-nil, zero value otherwise.

### GetCurrentPayOk

`func (o *EarningsTotal) GetCurrentPayOk() (*Pay, bool)`

GetCurrentPayOk returns a tuple with the CurrentPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPay

`func (o *EarningsTotal) SetCurrentPay(v Pay)`

SetCurrentPay sets CurrentPay field to given value.

### HasCurrentPay

`func (o *EarningsTotal) HasCurrentPay() bool`

HasCurrentPay returns a boolean if a field has been set.

### GetYtdPay

`func (o *EarningsTotal) GetYtdPay() Pay`

GetYtdPay returns the YtdPay field if non-nil, zero value otherwise.

### GetYtdPayOk

`func (o *EarningsTotal) GetYtdPayOk() (*Pay, bool)`

GetYtdPayOk returns a tuple with the YtdPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdPay

`func (o *EarningsTotal) SetYtdPay(v Pay)`

SetYtdPay sets YtdPay field to given value.

### HasYtdPay

`func (o *EarningsTotal) HasYtdPay() bool`

HasYtdPay returns a boolean if a field has been set.

### GetHours

`func (o *EarningsTotal) GetHours() float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *EarningsTotal) GetHoursOk() (*float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *EarningsTotal) SetHours(v float32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *EarningsTotal) HasHours() bool`

HasHours returns a boolean if a field has been set.

### SetHoursNil

`func (o *EarningsTotal) SetHoursNil(b bool)`

 SetHoursNil sets the value for Hours to be an explicit nil

### UnsetHours
`func (o *EarningsTotal) UnsetHours()`

UnsetHours ensures that no value is present for Hours, not even an explicit nil
### GetIsoCurrencyCode

`func (o *EarningsTotal) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *EarningsTotal) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *EarningsTotal) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *EarningsTotal) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *EarningsTotal) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *EarningsTotal) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *EarningsTotal) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *EarningsTotal) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *EarningsTotal) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.

### HasUnofficialCurrencyCode

`func (o *EarningsTotal) HasUnofficialCurrencyCode() bool`

HasUnofficialCurrencyCode returns a boolean if a field has been set.

### SetUnofficialCurrencyCodeNil

`func (o *EarningsTotal) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *EarningsTotal) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
### GetYtdAmount

`func (o *EarningsTotal) GetYtdAmount() float32`

GetYtdAmount returns the YtdAmount field if non-nil, zero value otherwise.

### GetYtdAmountOk

`func (o *EarningsTotal) GetYtdAmountOk() (*float32, bool)`

GetYtdAmountOk returns a tuple with the YtdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdAmount

`func (o *EarningsTotal) SetYtdAmount(v float32)`

SetYtdAmount sets YtdAmount field to given value.

### HasYtdAmount

`func (o *EarningsTotal) HasYtdAmount() bool`

HasYtdAmount returns a boolean if a field has been set.

### SetYtdAmountNil

`func (o *EarningsTotal) SetYtdAmountNil(b bool)`

 SetYtdAmountNil sets the value for YtdAmount to be an explicit nil

### UnsetYtdAmount
`func (o *EarningsTotal) UnsetYtdAmount()`

UnsetYtdAmount ensures that no value is present for YtdAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


