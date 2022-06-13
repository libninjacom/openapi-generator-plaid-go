# DistributionBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **NullableString** | Name of the account for the given distribution. | [optional] 
**BankName** | Pointer to **NullableString** | The name of the bank that the payment is being deposited to. | [optional] 
**CurrentAmount** | Pointer to **NullableFloat32** | The amount distributed to this account. | [optional] 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the net pay. Always &#x60;null&#x60; if &#x60;unofficial_currency_code&#x60; is non-null. | [optional] 
**Mask** | Pointer to **NullableString** | The last 2-4 alphanumeric characters of an account&#39;s official account number. | [optional] 
**Type** | Pointer to **NullableString** | Type of the account that the paystub was sent to (e.g. &#39;checking&#39;). | [optional] 
**UnofficialCurrencyCode** | Pointer to **NullableString** | The unofficial currency code associated with the net pay. Always &#x60;null&#x60; if &#x60;iso_currency_code&#x60; is non-&#x60;null&#x60;. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;iso_currency_code&#x60;s. | [optional] 
**CurrentPay** | Pointer to [**Pay**](Pay.md) |  | [optional] 

## Methods

### NewDistributionBreakdown

`func NewDistributionBreakdown() *DistributionBreakdown`

NewDistributionBreakdown instantiates a new DistributionBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionBreakdownWithDefaults

`func NewDistributionBreakdownWithDefaults() *DistributionBreakdown`

NewDistributionBreakdownWithDefaults instantiates a new DistributionBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *DistributionBreakdown) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *DistributionBreakdown) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *DistributionBreakdown) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *DistributionBreakdown) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *DistributionBreakdown) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *DistributionBreakdown) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetBankName

`func (o *DistributionBreakdown) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *DistributionBreakdown) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *DistributionBreakdown) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *DistributionBreakdown) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### SetBankNameNil

`func (o *DistributionBreakdown) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *DistributionBreakdown) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetCurrentAmount

`func (o *DistributionBreakdown) GetCurrentAmount() float32`

GetCurrentAmount returns the CurrentAmount field if non-nil, zero value otherwise.

### GetCurrentAmountOk

`func (o *DistributionBreakdown) GetCurrentAmountOk() (*float32, bool)`

GetCurrentAmountOk returns a tuple with the CurrentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmount

`func (o *DistributionBreakdown) SetCurrentAmount(v float32)`

SetCurrentAmount sets CurrentAmount field to given value.

### HasCurrentAmount

`func (o *DistributionBreakdown) HasCurrentAmount() bool`

HasCurrentAmount returns a boolean if a field has been set.

### SetCurrentAmountNil

`func (o *DistributionBreakdown) SetCurrentAmountNil(b bool)`

 SetCurrentAmountNil sets the value for CurrentAmount to be an explicit nil

### UnsetCurrentAmount
`func (o *DistributionBreakdown) UnsetCurrentAmount()`

UnsetCurrentAmount ensures that no value is present for CurrentAmount, not even an explicit nil
### GetIsoCurrencyCode

`func (o *DistributionBreakdown) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *DistributionBreakdown) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *DistributionBreakdown) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *DistributionBreakdown) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *DistributionBreakdown) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *DistributionBreakdown) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetMask

`func (o *DistributionBreakdown) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *DistributionBreakdown) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *DistributionBreakdown) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *DistributionBreakdown) HasMask() bool`

HasMask returns a boolean if a field has been set.

### SetMaskNil

`func (o *DistributionBreakdown) SetMaskNil(b bool)`

 SetMaskNil sets the value for Mask to be an explicit nil

### UnsetMask
`func (o *DistributionBreakdown) UnsetMask()`

UnsetMask ensures that no value is present for Mask, not even an explicit nil
### GetType

`func (o *DistributionBreakdown) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DistributionBreakdown) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DistributionBreakdown) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DistributionBreakdown) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DistributionBreakdown) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DistributionBreakdown) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *DistributionBreakdown) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *DistributionBreakdown) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *DistributionBreakdown) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.

### HasUnofficialCurrencyCode

`func (o *DistributionBreakdown) HasUnofficialCurrencyCode() bool`

HasUnofficialCurrencyCode returns a boolean if a field has been set.

### SetUnofficialCurrencyCodeNil

`func (o *DistributionBreakdown) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *DistributionBreakdown) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
### GetCurrentPay

`func (o *DistributionBreakdown) GetCurrentPay() Pay`

GetCurrentPay returns the CurrentPay field if non-nil, zero value otherwise.

### GetCurrentPayOk

`func (o *DistributionBreakdown) GetCurrentPayOk() (*Pay, bool)`

GetCurrentPayOk returns a tuple with the CurrentPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPay

`func (o *DistributionBreakdown) SetCurrentPay(v Pay)`

SetCurrentPay sets CurrentPay field to given value.

### HasCurrentPay

`func (o *DistributionBreakdown) HasCurrentPay() bool`

HasCurrentPay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


