# SecurityOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isin** | Pointer to **string** | 12-character ISIN, a globally unique securities identifier. | [optional] 
**Cusip** | Pointer to **string** | 9-character CUSIP, an identifier assigned to North American securities. | [optional] 
**Sedol** | Pointer to **string** | 7-character SEDOL, an identifier assigned to securities in the UK. | [optional] 
**Name** | Pointer to **string** | A descriptive name for the security, suitable for display. | [optional] 
**TickerSymbol** | Pointer to **string** | The security’s trading symbol for publicly traded securities, and otherwise a short identifier if available. | [optional] 
**Currency** | Pointer to **string** | Either a valid &#x60;iso_currency_code&#x60; or &#x60;unofficial_currency_code&#x60; | [optional] 

## Methods

### NewSecurityOverride

`func NewSecurityOverride() *SecurityOverride`

NewSecurityOverride instantiates a new SecurityOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityOverrideWithDefaults

`func NewSecurityOverrideWithDefaults() *SecurityOverride`

NewSecurityOverrideWithDefaults instantiates a new SecurityOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsin

`func (o *SecurityOverride) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *SecurityOverride) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *SecurityOverride) SetIsin(v string)`

SetIsin sets Isin field to given value.

### HasIsin

`func (o *SecurityOverride) HasIsin() bool`

HasIsin returns a boolean if a field has been set.

### GetCusip

`func (o *SecurityOverride) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *SecurityOverride) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *SecurityOverride) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *SecurityOverride) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetSedol

`func (o *SecurityOverride) GetSedol() string`

GetSedol returns the Sedol field if non-nil, zero value otherwise.

### GetSedolOk

`func (o *SecurityOverride) GetSedolOk() (*string, bool)`

GetSedolOk returns a tuple with the Sedol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSedol

`func (o *SecurityOverride) SetSedol(v string)`

SetSedol sets Sedol field to given value.

### HasSedol

`func (o *SecurityOverride) HasSedol() bool`

HasSedol returns a boolean if a field has been set.

### GetName

`func (o *SecurityOverride) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityOverride) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityOverride) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityOverride) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTickerSymbol

`func (o *SecurityOverride) GetTickerSymbol() string`

GetTickerSymbol returns the TickerSymbol field if non-nil, zero value otherwise.

### GetTickerSymbolOk

`func (o *SecurityOverride) GetTickerSymbolOk() (*string, bool)`

GetTickerSymbolOk returns a tuple with the TickerSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerSymbol

`func (o *SecurityOverride) SetTickerSymbol(v string)`

SetTickerSymbol sets TickerSymbol field to given value.

### HasTickerSymbol

`func (o *SecurityOverride) HasTickerSymbol() bool`

HasTickerSymbol returns a boolean if a field has been set.

### GetCurrency

`func (o *SecurityOverride) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SecurityOverride) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SecurityOverride) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SecurityOverride) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


