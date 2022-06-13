# InvestmentsTransactionsOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Posting date for the transaction. Must be formatted as an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) date. | 
**Name** | **string** | The institution&#39;s description of the transaction. | 
**Quantity** | **float32** | The number of units of the security involved in this transaction. Must be positive if the type is a buy and negative if the type is a sell. | 
**Price** | **float32** | The price of the security at which this transaction occurred. | 
**Fees** | Pointer to **float32** | The combined value of all fees applied to this transaction. | [optional] 
**Type** | **string** | The type of the investment transaction. Possible values are: &#x60;buy&#x60;: Buying an investment &#x60;sell&#x60;: Selling an investment &#x60;cash&#x60;: Activity that modifies a cash position &#x60;fee&#x60;: A fee on the account &#x60;transfer&#x60;: Activity that modifies a position, but not through buy/sell activity e.g. options exercise, portfolio transfer | 
**Currency** | **string** | Either a valid &#x60;iso_currency_code&#x60; or &#x60;unofficial_currency_code&#x60; | 
**Security** | Pointer to [**SecurityOverride**](SecurityOverride.md) |  | [optional] 

## Methods

### NewInvestmentsTransactionsOverride

`func NewInvestmentsTransactionsOverride(date string, name string, quantity float32, price float32, type_ string, currency string, ) *InvestmentsTransactionsOverride`

NewInvestmentsTransactionsOverride instantiates a new InvestmentsTransactionsOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentsTransactionsOverrideWithDefaults

`func NewInvestmentsTransactionsOverrideWithDefaults() *InvestmentsTransactionsOverride`

NewInvestmentsTransactionsOverrideWithDefaults instantiates a new InvestmentsTransactionsOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *InvestmentsTransactionsOverride) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InvestmentsTransactionsOverride) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InvestmentsTransactionsOverride) SetDate(v string)`

SetDate sets Date field to given value.


### GetName

`func (o *InvestmentsTransactionsOverride) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvestmentsTransactionsOverride) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvestmentsTransactionsOverride) SetName(v string)`

SetName sets Name field to given value.


### GetQuantity

`func (o *InvestmentsTransactionsOverride) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvestmentsTransactionsOverride) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvestmentsTransactionsOverride) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *InvestmentsTransactionsOverride) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvestmentsTransactionsOverride) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvestmentsTransactionsOverride) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetFees

`func (o *InvestmentsTransactionsOverride) GetFees() float32`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *InvestmentsTransactionsOverride) GetFeesOk() (*float32, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *InvestmentsTransactionsOverride) SetFees(v float32)`

SetFees sets Fees field to given value.

### HasFees

`func (o *InvestmentsTransactionsOverride) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetType

`func (o *InvestmentsTransactionsOverride) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvestmentsTransactionsOverride) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvestmentsTransactionsOverride) SetType(v string)`

SetType sets Type field to given value.


### GetCurrency

`func (o *InvestmentsTransactionsOverride) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvestmentsTransactionsOverride) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvestmentsTransactionsOverride) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSecurity

`func (o *InvestmentsTransactionsOverride) GetSecurity() SecurityOverride`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *InvestmentsTransactionsOverride) GetSecurityOk() (*SecurityOverride, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *InvestmentsTransactionsOverride) SetSecurity(v SecurityOverride)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *InvestmentsTransactionsOverride) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


