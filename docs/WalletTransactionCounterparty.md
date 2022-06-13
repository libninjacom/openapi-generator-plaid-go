# WalletTransactionCounterparty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the counterparty | 
**Numbers** | [**WalletTransactionCounterpartyNumbers**](WalletTransactionCounterpartyNumbers.md) |  | 

## Methods

### NewWalletTransactionCounterparty

`func NewWalletTransactionCounterparty(name string, numbers WalletTransactionCounterpartyNumbers, ) *WalletTransactionCounterparty`

NewWalletTransactionCounterparty instantiates a new WalletTransactionCounterparty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionCounterpartyWithDefaults

`func NewWalletTransactionCounterpartyWithDefaults() *WalletTransactionCounterparty`

NewWalletTransactionCounterpartyWithDefaults instantiates a new WalletTransactionCounterparty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WalletTransactionCounterparty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletTransactionCounterparty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletTransactionCounterparty) SetName(v string)`

SetName sets Name field to given value.


### GetNumbers

`func (o *WalletTransactionCounterparty) GetNumbers() WalletTransactionCounterpartyNumbers`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *WalletTransactionCounterparty) GetNumbersOk() (*WalletTransactionCounterpartyNumbers, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *WalletTransactionCounterparty) SetNumbers(v WalletTransactionCounterpartyNumbers)`

SetNumbers sets Numbers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


