# DeductionsTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAmount** | Pointer to **NullableFloat32** | Raw amount of the deduction | [optional] 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the line item. Always &#x60;null&#x60; if &#x60;unofficial_currency_code&#x60; is non-null. | [optional] 
**UnofficialCurrencyCode** | Pointer to **NullableString** | The unofficial currency code associated with the line item. Always &#x60;null&#x60; if &#x60;iso_currency_code&#x60; is non-&#x60;null&#x60;. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;iso_currency_code&#x60;s. | [optional] 
**YtdAmount** | Pointer to **NullableFloat32** | The year-to-date total amount of the deductions | [optional] 

## Methods

### NewDeductionsTotal

`func NewDeductionsTotal() *DeductionsTotal`

NewDeductionsTotal instantiates a new DeductionsTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeductionsTotalWithDefaults

`func NewDeductionsTotalWithDefaults() *DeductionsTotal`

NewDeductionsTotalWithDefaults instantiates a new DeductionsTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentAmount

`func (o *DeductionsTotal) GetCurrentAmount() float32`

GetCurrentAmount returns the CurrentAmount field if non-nil, zero value otherwise.

### GetCurrentAmountOk

`func (o *DeductionsTotal) GetCurrentAmountOk() (*float32, bool)`

GetCurrentAmountOk returns a tuple with the CurrentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmount

`func (o *DeductionsTotal) SetCurrentAmount(v float32)`

SetCurrentAmount sets CurrentAmount field to given value.

### HasCurrentAmount

`func (o *DeductionsTotal) HasCurrentAmount() bool`

HasCurrentAmount returns a boolean if a field has been set.

### SetCurrentAmountNil

`func (o *DeductionsTotal) SetCurrentAmountNil(b bool)`

 SetCurrentAmountNil sets the value for CurrentAmount to be an explicit nil

### UnsetCurrentAmount
`func (o *DeductionsTotal) UnsetCurrentAmount()`

UnsetCurrentAmount ensures that no value is present for CurrentAmount, not even an explicit nil
### GetIsoCurrencyCode

`func (o *DeductionsTotal) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *DeductionsTotal) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *DeductionsTotal) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *DeductionsTotal) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *DeductionsTotal) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *DeductionsTotal) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *DeductionsTotal) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *DeductionsTotal) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *DeductionsTotal) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.

### HasUnofficialCurrencyCode

`func (o *DeductionsTotal) HasUnofficialCurrencyCode() bool`

HasUnofficialCurrencyCode returns a boolean if a field has been set.

### SetUnofficialCurrencyCodeNil

`func (o *DeductionsTotal) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *DeductionsTotal) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
### GetYtdAmount

`func (o *DeductionsTotal) GetYtdAmount() float32`

GetYtdAmount returns the YtdAmount field if non-nil, zero value otherwise.

### GetYtdAmountOk

`func (o *DeductionsTotal) GetYtdAmountOk() (*float32, bool)`

GetYtdAmountOk returns a tuple with the YtdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdAmount

`func (o *DeductionsTotal) SetYtdAmount(v float32)`

SetYtdAmount sets YtdAmount field to given value.

### HasYtdAmount

`func (o *DeductionsTotal) HasYtdAmount() bool`

HasYtdAmount returns a boolean if a field has been set.

### SetYtdAmountNil

`func (o *DeductionsTotal) SetYtdAmountNil(b bool)`

 SetYtdAmountNil sets the value for YtdAmount to be an explicit nil

### UnsetYtdAmount
`func (o *DeductionsTotal) UnsetYtdAmount()`

UnsetYtdAmount ensures that no value is present for YtdAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


