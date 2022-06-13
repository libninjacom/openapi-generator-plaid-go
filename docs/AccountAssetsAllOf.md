# AccountAssetsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysAvailable** | **float32** | The duration of transaction history available for this Item, typically defined as the time since the date of the earliest transaction in that account. Only returned by Assets endpoints. | 
**Transactions** | [**[]AssetReportTransaction**](AssetReportTransaction.md) | Transaction history associated with the account. Only returned by Assets endpoints. Transaction history returned by endpoints such as &#x60;/transactions/get&#x60; or &#x60;/investments/transactions/get&#x60; will be returned in the top-level &#x60;transactions&#x60; field instead. | 
**Owners** | [**[]Owner**](Owner.md) | Data returned by the financial institution about the account owner or owners. Only returned by Identity or Assets endpoints. For business accounts, the name reported may be either the name of the individual or the name of the business, depending on the institution. Multiple owners on a single account will be represented in the same &#x60;owner&#x60; object, not in multiple owner objects within the array. In API versions 2018-05-22 and earlier, the &#x60;owners&#x60; object is not returned, and instead identity information is returned in the top level &#x60;identity&#x60; object. For more details, see [Plaid API versioning](https://plaid.com/docs/api/versioning/#version-2019-05-29) | 
**HistoricalBalances** | [**[]HistoricalBalance**](HistoricalBalance.md) | Calculated data about the historical balances on the account. Only returned by Assets endpoints and currently not supported by &#x60;brokerage&#x60; or &#x60;investment&#x60; accounts. | 

## Methods

### NewAccountAssetsAllOf

`func NewAccountAssetsAllOf(daysAvailable float32, transactions []AssetReportTransaction, owners []Owner, historicalBalances []HistoricalBalance, ) *AccountAssetsAllOf`

NewAccountAssetsAllOf instantiates a new AccountAssetsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAssetsAllOfWithDefaults

`func NewAccountAssetsAllOfWithDefaults() *AccountAssetsAllOf`

NewAccountAssetsAllOfWithDefaults instantiates a new AccountAssetsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysAvailable

`func (o *AccountAssetsAllOf) GetDaysAvailable() float32`

GetDaysAvailable returns the DaysAvailable field if non-nil, zero value otherwise.

### GetDaysAvailableOk

`func (o *AccountAssetsAllOf) GetDaysAvailableOk() (*float32, bool)`

GetDaysAvailableOk returns a tuple with the DaysAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysAvailable

`func (o *AccountAssetsAllOf) SetDaysAvailable(v float32)`

SetDaysAvailable sets DaysAvailable field to given value.


### GetTransactions

`func (o *AccountAssetsAllOf) GetTransactions() []AssetReportTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *AccountAssetsAllOf) GetTransactionsOk() (*[]AssetReportTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *AccountAssetsAllOf) SetTransactions(v []AssetReportTransaction)`

SetTransactions sets Transactions field to given value.


### GetOwners

`func (o *AccountAssetsAllOf) GetOwners() []Owner`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *AccountAssetsAllOf) GetOwnersOk() (*[]Owner, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *AccountAssetsAllOf) SetOwners(v []Owner)`

SetOwners sets Owners field to given value.


### GetHistoricalBalances

`func (o *AccountAssetsAllOf) GetHistoricalBalances() []HistoricalBalance`

GetHistoricalBalances returns the HistoricalBalances field if non-nil, zero value otherwise.

### GetHistoricalBalancesOk

`func (o *AccountAssetsAllOf) GetHistoricalBalancesOk() (*[]HistoricalBalance, bool)`

GetHistoricalBalancesOk returns a tuple with the HistoricalBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalBalances

`func (o *AccountAssetsAllOf) SetHistoricalBalances(v []HistoricalBalance)`

SetHistoricalBalances sets HistoricalBalances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


