# DeductionsBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAmount** | Pointer to **NullableFloat32** | Raw amount of the deduction | [optional] 
**Description** | Pointer to **NullableString** | Description of the deduction line item | [optional] 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the line item. Always &#x60;null&#x60; if &#x60;unofficial_currency_code&#x60; is non-null. | [optional] 
**UnofficialCurrencyCode** | Pointer to **NullableString** | The unofficial currency code associated with the line item. Always &#x60;null&#x60; if &#x60;iso_currency_code&#x60; is non-&#x60;null&#x60;. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;iso_currency_code&#x60;s. | [optional] 
**YtdAmount** | Pointer to **NullableFloat32** | The year-to-date amount of the deduction | [optional] 

## Methods

### NewDeductionsBreakdown

`func NewDeductionsBreakdown() *DeductionsBreakdown`

NewDeductionsBreakdown instantiates a new DeductionsBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeductionsBreakdownWithDefaults

`func NewDeductionsBreakdownWithDefaults() *DeductionsBreakdown`

NewDeductionsBreakdownWithDefaults instantiates a new DeductionsBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentAmount

`func (o *DeductionsBreakdown) GetCurrentAmount() float32`

GetCurrentAmount returns the CurrentAmount field if non-nil, zero value otherwise.

### GetCurrentAmountOk

`func (o *DeductionsBreakdown) GetCurrentAmountOk() (*float32, bool)`

GetCurrentAmountOk returns a tuple with the CurrentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmount

`func (o *DeductionsBreakdown) SetCurrentAmount(v float32)`

SetCurrentAmount sets CurrentAmount field to given value.

### HasCurrentAmount

`func (o *DeductionsBreakdown) HasCurrentAmount() bool`

HasCurrentAmount returns a boolean if a field has been set.

### SetCurrentAmountNil

`func (o *DeductionsBreakdown) SetCurrentAmountNil(b bool)`

 SetCurrentAmountNil sets the value for CurrentAmount to be an explicit nil

### UnsetCurrentAmount
`func (o *DeductionsBreakdown) UnsetCurrentAmount()`

UnsetCurrentAmount ensures that no value is present for CurrentAmount, not even an explicit nil
### GetDescription

`func (o *DeductionsBreakdown) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeductionsBreakdown) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeductionsBreakdown) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeductionsBreakdown) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DeductionsBreakdown) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DeductionsBreakdown) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsoCurrencyCode

`func (o *DeductionsBreakdown) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *DeductionsBreakdown) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *DeductionsBreakdown) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *DeductionsBreakdown) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *DeductionsBreakdown) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *DeductionsBreakdown) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *DeductionsBreakdown) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *DeductionsBreakdown) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *DeductionsBreakdown) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.

### HasUnofficialCurrencyCode

`func (o *DeductionsBreakdown) HasUnofficialCurrencyCode() bool`

HasUnofficialCurrencyCode returns a boolean if a field has been set.

### SetUnofficialCurrencyCodeNil

`func (o *DeductionsBreakdown) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *DeductionsBreakdown) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
### GetYtdAmount

`func (o *DeductionsBreakdown) GetYtdAmount() float32`

GetYtdAmount returns the YtdAmount field if non-nil, zero value otherwise.

### GetYtdAmountOk

`func (o *DeductionsBreakdown) GetYtdAmountOk() (*float32, bool)`

GetYtdAmountOk returns a tuple with the YtdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdAmount

`func (o *DeductionsBreakdown) SetYtdAmount(v float32)`

SetYtdAmount sets YtdAmount field to given value.

### HasYtdAmount

`func (o *DeductionsBreakdown) HasYtdAmount() bool`

HasYtdAmount returns a boolean if a field has been set.

### SetYtdAmountNil

`func (o *DeductionsBreakdown) SetYtdAmountNil(b bool)`

 SetYtdAmountNil sets the value for YtdAmount to be an explicit nil

### UnsetYtdAmount
`func (o *DeductionsBreakdown) UnsetYtdAmount()`

UnsetYtdAmount ensures that no value is present for YtdAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


