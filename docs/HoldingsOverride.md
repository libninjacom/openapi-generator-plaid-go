# HoldingsOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstitutionPrice** | **float32** | The last price given by the institution for this security | 
**InstitutionPriceAsOf** | Pointer to **string** | The date at which &#x60;institution_price&#x60; was current. Must be formatted as an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) date. | [optional] 
**CostBasis** | Pointer to **float32** | The average original value of the holding. Multiple cost basis values for the same security purchased at different prices are not supported. | [optional] 
**Quantity** | **float32** | The total quantity of the asset held, as reported by the financial institution. | 
**Currency** | **string** | Either a valid &#x60;iso_currency_code&#x60; or &#x60;unofficial_currency_code&#x60; | 
**Security** | [**SecurityOverride**](SecurityOverride.md) |  | 

## Methods

### NewHoldingsOverride

`func NewHoldingsOverride(institutionPrice float32, quantity float32, currency string, security SecurityOverride, ) *HoldingsOverride`

NewHoldingsOverride instantiates a new HoldingsOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldingsOverrideWithDefaults

`func NewHoldingsOverrideWithDefaults() *HoldingsOverride`

NewHoldingsOverrideWithDefaults instantiates a new HoldingsOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutionPrice

`func (o *HoldingsOverride) GetInstitutionPrice() float32`

GetInstitutionPrice returns the InstitutionPrice field if non-nil, zero value otherwise.

### GetInstitutionPriceOk

`func (o *HoldingsOverride) GetInstitutionPriceOk() (*float32, bool)`

GetInstitutionPriceOk returns a tuple with the InstitutionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionPrice

`func (o *HoldingsOverride) SetInstitutionPrice(v float32)`

SetInstitutionPrice sets InstitutionPrice field to given value.


### GetInstitutionPriceAsOf

`func (o *HoldingsOverride) GetInstitutionPriceAsOf() string`

GetInstitutionPriceAsOf returns the InstitutionPriceAsOf field if non-nil, zero value otherwise.

### GetInstitutionPriceAsOfOk

`func (o *HoldingsOverride) GetInstitutionPriceAsOfOk() (*string, bool)`

GetInstitutionPriceAsOfOk returns a tuple with the InstitutionPriceAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionPriceAsOf

`func (o *HoldingsOverride) SetInstitutionPriceAsOf(v string)`

SetInstitutionPriceAsOf sets InstitutionPriceAsOf field to given value.

### HasInstitutionPriceAsOf

`func (o *HoldingsOverride) HasInstitutionPriceAsOf() bool`

HasInstitutionPriceAsOf returns a boolean if a field has been set.

### GetCostBasis

`func (o *HoldingsOverride) GetCostBasis() float32`

GetCostBasis returns the CostBasis field if non-nil, zero value otherwise.

### GetCostBasisOk

`func (o *HoldingsOverride) GetCostBasisOk() (*float32, bool)`

GetCostBasisOk returns a tuple with the CostBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBasis

`func (o *HoldingsOverride) SetCostBasis(v float32)`

SetCostBasis sets CostBasis field to given value.

### HasCostBasis

`func (o *HoldingsOverride) HasCostBasis() bool`

HasCostBasis returns a boolean if a field has been set.

### GetQuantity

`func (o *HoldingsOverride) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *HoldingsOverride) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *HoldingsOverride) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetCurrency

`func (o *HoldingsOverride) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *HoldingsOverride) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *HoldingsOverride) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSecurity

`func (o *HoldingsOverride) GetSecurity() SecurityOverride`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *HoldingsOverride) GetSecurityOk() (*SecurityOverride, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *HoldingsOverride) SetSecurity(v SecurityOverride)`

SetSecurity sets Security field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


