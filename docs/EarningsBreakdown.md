# EarningsBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanonicalDescription** | Pointer to [**NullableEarningsBreakdownCanonicalDescription**](EarningsBreakdownCanonicalDescription.md) |  | [optional] 
**CurrentAmount** | Pointer to **NullableFloat32** | Raw amount of the earning line item. | [optional] 
**Description** | Pointer to **NullableString** | Description of the earning line item. | [optional] 
**Hours** | Pointer to **NullableFloat32** | Number of hours applicable for this earning. | [optional] 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the line item. Always &#x60;null&#x60; if &#x60;unofficial_currency_code&#x60; is non-null. | [optional] 
**Rate** | Pointer to **NullableFloat32** | Hourly rate applicable for this earning. | [optional] 
**UnofficialCurrencyCode** | Pointer to **NullableString** | The unofficial currency code associated with the line item. Always &#x60;null&#x60; if &#x60;iso_currency_code&#x60; is non-&#x60;null&#x60;. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;iso_currency_code&#x60;s. | [optional] 
**YtdAmount** | Pointer to **NullableFloat32** | The year-to-date amount of the deduction. | [optional] 

## Methods

### NewEarningsBreakdown

`func NewEarningsBreakdown() *EarningsBreakdown`

NewEarningsBreakdown instantiates a new EarningsBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarningsBreakdownWithDefaults

`func NewEarningsBreakdownWithDefaults() *EarningsBreakdown`

NewEarningsBreakdownWithDefaults instantiates a new EarningsBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonicalDescription

`func (o *EarningsBreakdown) GetCanonicalDescription() EarningsBreakdownCanonicalDescription`

GetCanonicalDescription returns the CanonicalDescription field if non-nil, zero value otherwise.

### GetCanonicalDescriptionOk

`func (o *EarningsBreakdown) GetCanonicalDescriptionOk() (*EarningsBreakdownCanonicalDescription, bool)`

GetCanonicalDescriptionOk returns a tuple with the CanonicalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalDescription

`func (o *EarningsBreakdown) SetCanonicalDescription(v EarningsBreakdownCanonicalDescription)`

SetCanonicalDescription sets CanonicalDescription field to given value.

### HasCanonicalDescription

`func (o *EarningsBreakdown) HasCanonicalDescription() bool`

HasCanonicalDescription returns a boolean if a field has been set.

### SetCanonicalDescriptionNil

`func (o *EarningsBreakdown) SetCanonicalDescriptionNil(b bool)`

 SetCanonicalDescriptionNil sets the value for CanonicalDescription to be an explicit nil

### UnsetCanonicalDescription
`func (o *EarningsBreakdown) UnsetCanonicalDescription()`

UnsetCanonicalDescription ensures that no value is present for CanonicalDescription, not even an explicit nil
### GetCurrentAmount

`func (o *EarningsBreakdown) GetCurrentAmount() float32`

GetCurrentAmount returns the CurrentAmount field if non-nil, zero value otherwise.

### GetCurrentAmountOk

`func (o *EarningsBreakdown) GetCurrentAmountOk() (*float32, bool)`

GetCurrentAmountOk returns a tuple with the CurrentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmount

`func (o *EarningsBreakdown) SetCurrentAmount(v float32)`

SetCurrentAmount sets CurrentAmount field to given value.

### HasCurrentAmount

`func (o *EarningsBreakdown) HasCurrentAmount() bool`

HasCurrentAmount returns a boolean if a field has been set.

### SetCurrentAmountNil

`func (o *EarningsBreakdown) SetCurrentAmountNil(b bool)`

 SetCurrentAmountNil sets the value for CurrentAmount to be an explicit nil

### UnsetCurrentAmount
`func (o *EarningsBreakdown) UnsetCurrentAmount()`

UnsetCurrentAmount ensures that no value is present for CurrentAmount, not even an explicit nil
### GetDescription

`func (o *EarningsBreakdown) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EarningsBreakdown) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EarningsBreakdown) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EarningsBreakdown) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EarningsBreakdown) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EarningsBreakdown) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHours

`func (o *EarningsBreakdown) GetHours() float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *EarningsBreakdown) GetHoursOk() (*float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *EarningsBreakdown) SetHours(v float32)`

SetHours sets Hours field to given value.

### HasHours

`func (o *EarningsBreakdown) HasHours() bool`

HasHours returns a boolean if a field has been set.

### SetHoursNil

`func (o *EarningsBreakdown) SetHoursNil(b bool)`

 SetHoursNil sets the value for Hours to be an explicit nil

### UnsetHours
`func (o *EarningsBreakdown) UnsetHours()`

UnsetHours ensures that no value is present for Hours, not even an explicit nil
### GetIsoCurrencyCode

`func (o *EarningsBreakdown) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *EarningsBreakdown) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *EarningsBreakdown) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *EarningsBreakdown) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *EarningsBreakdown) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *EarningsBreakdown) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetRate

`func (o *EarningsBreakdown) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *EarningsBreakdown) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *EarningsBreakdown) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *EarningsBreakdown) HasRate() bool`

HasRate returns a boolean if a field has been set.

### SetRateNil

`func (o *EarningsBreakdown) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *EarningsBreakdown) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *EarningsBreakdown) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *EarningsBreakdown) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *EarningsBreakdown) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.

### HasUnofficialCurrencyCode

`func (o *EarningsBreakdown) HasUnofficialCurrencyCode() bool`

HasUnofficialCurrencyCode returns a boolean if a field has been set.

### SetUnofficialCurrencyCodeNil

`func (o *EarningsBreakdown) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *EarningsBreakdown) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
### GetYtdAmount

`func (o *EarningsBreakdown) GetYtdAmount() float32`

GetYtdAmount returns the YtdAmount field if non-nil, zero value otherwise.

### GetYtdAmountOk

`func (o *EarningsBreakdown) GetYtdAmountOk() (*float32, bool)`

GetYtdAmountOk returns a tuple with the YtdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdAmount

`func (o *EarningsBreakdown) SetYtdAmount(v float32)`

SetYtdAmount sets YtdAmount field to given value.

### HasYtdAmount

`func (o *EarningsBreakdown) HasYtdAmount() bool`

HasYtdAmount returns a boolean if a field has been set.

### SetYtdAmountNil

`func (o *EarningsBreakdown) SetYtdAmountNil(b bool)`

 SetYtdAmountNil sets the value for YtdAmount to be an explicit nil

### UnsetYtdAmount
`func (o *EarningsBreakdown) UnsetYtdAmount()`

UnsetYtdAmount ensures that no value is present for YtdAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


