# WalletTransactionCounterpartyBACS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | The account number of the account. Maximum of 10 characters. | [optional] 
**SortCode** | Pointer to **string** | The 6-character sort code of the account. | [optional] 

## Methods

### NewWalletTransactionCounterpartyBACS

`func NewWalletTransactionCounterpartyBACS() *WalletTransactionCounterpartyBACS`

NewWalletTransactionCounterpartyBACS instantiates a new WalletTransactionCounterpartyBACS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionCounterpartyBACSWithDefaults

`func NewWalletTransactionCounterpartyBACSWithDefaults() *WalletTransactionCounterpartyBACS`

NewWalletTransactionCounterpartyBACSWithDefaults instantiates a new WalletTransactionCounterpartyBACS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *WalletTransactionCounterpartyBACS) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WalletTransactionCounterpartyBACS) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WalletTransactionCounterpartyBACS) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *WalletTransactionCounterpartyBACS) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSortCode

`func (o *WalletTransactionCounterpartyBACS) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *WalletTransactionCounterpartyBACS) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *WalletTransactionCounterpartyBACS) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.

### HasSortCode

`func (o *WalletTransactionCounterpartyBACS) HasSortCode() bool`

HasSortCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


