# WalletBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsoCurrencyCode** | **string** | The ISO-4217 currency code of the balance | 
**Current** | **float32** | The total amount of funds in the account | 

## Methods

### NewWalletBalance

`func NewWalletBalance(isoCurrencyCode string, current float32, ) *WalletBalance`

NewWalletBalance instantiates a new WalletBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletBalanceWithDefaults

`func NewWalletBalanceWithDefaults() *WalletBalance`

NewWalletBalanceWithDefaults instantiates a new WalletBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsoCurrencyCode

`func (o *WalletBalance) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *WalletBalance) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *WalletBalance) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### GetCurrent

`func (o *WalletBalance) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *WalletBalance) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *WalletBalance) SetCurrent(v float32)`

SetCurrent sets Current field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


