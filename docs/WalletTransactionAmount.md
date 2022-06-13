# WalletTransactionAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsoCurrencyCode** | **string** | The ISO-4217 currency code of the transaction. Currently, only &#x60;\&quot;GBP\&quot;&#x60; is supported. | 
**Value** | **float32** | The amount of the transaction. Must contain at most two digits of precision e.g. &#x60;1.23&#x60;. | 

## Methods

### NewWalletTransactionAmount

`func NewWalletTransactionAmount(isoCurrencyCode string, value float32, ) *WalletTransactionAmount`

NewWalletTransactionAmount instantiates a new WalletTransactionAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionAmountWithDefaults

`func NewWalletTransactionAmountWithDefaults() *WalletTransactionAmount`

NewWalletTransactionAmountWithDefaults instantiates a new WalletTransactionAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsoCurrencyCode

`func (o *WalletTransactionAmount) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *WalletTransactionAmount) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *WalletTransactionAmount) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### GetValue

`func (o *WalletTransactionAmount) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WalletTransactionAmount) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WalletTransactionAmount) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


