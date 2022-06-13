# OverrideAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OverrideAccountType**](OverrideAccountType.md) |  | 
**Subtype** | [**NullableAccountSubtype**](AccountSubtype.md) |  | 
**StartingBalance** | **float32** | If provided, the account will start with this amount as the current balance.  | 
**ForceAvailableBalance** | **float32** | If provided, the account will always have this amount as its  available balance, regardless of current balance or changes in transactions over time. | 
**Currency** | **string** | ISO-4217 currency code. If provided, the account will be denominated in the given currency. Transactions will also be in this currency by default. | 
**Meta** | [**Meta**](Meta.md) |  | 
**Numbers** | [**Numbers**](Numbers.md) |  | 
**Transactions** | [**[]TransactionOverride**](TransactionOverride.md) | Specify the list of transactions on the account. | 
**Holdings** | Pointer to [**HoldingsOverride**](HoldingsOverride.md) |  | [optional] 
**InvestmentTransactions** | Pointer to [**InvestmentsTransactionsOverride**](InvestmentsTransactionsOverride.md) |  | [optional] 
**Identity** | [**OwnerOverride**](OwnerOverride.md) |  | 
**Liability** | [**LiabilityOverride**](LiabilityOverride.md) |  | 
**InflowModel** | [**InflowModel**](InflowModel.md) |  | 
**Income** | Pointer to [**IncomeOverride**](IncomeOverride.md) |  | [optional] 

## Methods

### NewOverrideAccounts

`func NewOverrideAccounts(type_ OverrideAccountType, subtype NullableAccountSubtype, startingBalance float32, forceAvailableBalance float32, currency string, meta Meta, numbers Numbers, transactions []TransactionOverride, identity OwnerOverride, liability LiabilityOverride, inflowModel InflowModel, ) *OverrideAccounts`

NewOverrideAccounts instantiates a new OverrideAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverrideAccountsWithDefaults

`func NewOverrideAccountsWithDefaults() *OverrideAccounts`

NewOverrideAccountsWithDefaults instantiates a new OverrideAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OverrideAccounts) GetType() OverrideAccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverrideAccounts) GetTypeOk() (*OverrideAccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverrideAccounts) SetType(v OverrideAccountType)`

SetType sets Type field to given value.


### GetSubtype

`func (o *OverrideAccounts) GetSubtype() AccountSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *OverrideAccounts) GetSubtypeOk() (*AccountSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *OverrideAccounts) SetSubtype(v AccountSubtype)`

SetSubtype sets Subtype field to given value.


### SetSubtypeNil

`func (o *OverrideAccounts) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *OverrideAccounts) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetStartingBalance

`func (o *OverrideAccounts) GetStartingBalance() float32`

GetStartingBalance returns the StartingBalance field if non-nil, zero value otherwise.

### GetStartingBalanceOk

`func (o *OverrideAccounts) GetStartingBalanceOk() (*float32, bool)`

GetStartingBalanceOk returns a tuple with the StartingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingBalance

`func (o *OverrideAccounts) SetStartingBalance(v float32)`

SetStartingBalance sets StartingBalance field to given value.


### GetForceAvailableBalance

`func (o *OverrideAccounts) GetForceAvailableBalance() float32`

GetForceAvailableBalance returns the ForceAvailableBalance field if non-nil, zero value otherwise.

### GetForceAvailableBalanceOk

`func (o *OverrideAccounts) GetForceAvailableBalanceOk() (*float32, bool)`

GetForceAvailableBalanceOk returns a tuple with the ForceAvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceAvailableBalance

`func (o *OverrideAccounts) SetForceAvailableBalance(v float32)`

SetForceAvailableBalance sets ForceAvailableBalance field to given value.


### GetCurrency

`func (o *OverrideAccounts) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OverrideAccounts) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OverrideAccounts) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMeta

`func (o *OverrideAccounts) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *OverrideAccounts) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *OverrideAccounts) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetNumbers

`func (o *OverrideAccounts) GetNumbers() Numbers`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *OverrideAccounts) GetNumbersOk() (*Numbers, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *OverrideAccounts) SetNumbers(v Numbers)`

SetNumbers sets Numbers field to given value.


### GetTransactions

`func (o *OverrideAccounts) GetTransactions() []TransactionOverride`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *OverrideAccounts) GetTransactionsOk() (*[]TransactionOverride, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *OverrideAccounts) SetTransactions(v []TransactionOverride)`

SetTransactions sets Transactions field to given value.


### GetHoldings

`func (o *OverrideAccounts) GetHoldings() HoldingsOverride`

GetHoldings returns the Holdings field if non-nil, zero value otherwise.

### GetHoldingsOk

`func (o *OverrideAccounts) GetHoldingsOk() (*HoldingsOverride, bool)`

GetHoldingsOk returns a tuple with the Holdings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldings

`func (o *OverrideAccounts) SetHoldings(v HoldingsOverride)`

SetHoldings sets Holdings field to given value.

### HasHoldings

`func (o *OverrideAccounts) HasHoldings() bool`

HasHoldings returns a boolean if a field has been set.

### GetInvestmentTransactions

`func (o *OverrideAccounts) GetInvestmentTransactions() InvestmentsTransactionsOverride`

GetInvestmentTransactions returns the InvestmentTransactions field if non-nil, zero value otherwise.

### GetInvestmentTransactionsOk

`func (o *OverrideAccounts) GetInvestmentTransactionsOk() (*InvestmentsTransactionsOverride, bool)`

GetInvestmentTransactionsOk returns a tuple with the InvestmentTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentTransactions

`func (o *OverrideAccounts) SetInvestmentTransactions(v InvestmentsTransactionsOverride)`

SetInvestmentTransactions sets InvestmentTransactions field to given value.

### HasInvestmentTransactions

`func (o *OverrideAccounts) HasInvestmentTransactions() bool`

HasInvestmentTransactions returns a boolean if a field has been set.

### GetIdentity

`func (o *OverrideAccounts) GetIdentity() OwnerOverride`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *OverrideAccounts) GetIdentityOk() (*OwnerOverride, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *OverrideAccounts) SetIdentity(v OwnerOverride)`

SetIdentity sets Identity field to given value.


### GetLiability

`func (o *OverrideAccounts) GetLiability() LiabilityOverride`

GetLiability returns the Liability field if non-nil, zero value otherwise.

### GetLiabilityOk

`func (o *OverrideAccounts) GetLiabilityOk() (*LiabilityOverride, bool)`

GetLiabilityOk returns a tuple with the Liability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiability

`func (o *OverrideAccounts) SetLiability(v LiabilityOverride)`

SetLiability sets Liability field to given value.


### GetInflowModel

`func (o *OverrideAccounts) GetInflowModel() InflowModel`

GetInflowModel returns the InflowModel field if non-nil, zero value otherwise.

### GetInflowModelOk

`func (o *OverrideAccounts) GetInflowModelOk() (*InflowModel, bool)`

GetInflowModelOk returns a tuple with the InflowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflowModel

`func (o *OverrideAccounts) SetInflowModel(v InflowModel)`

SetInflowModel sets InflowModel field to given value.


### GetIncome

`func (o *OverrideAccounts) GetIncome() IncomeOverride`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *OverrideAccounts) GetIncomeOk() (*IncomeOverride, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *OverrideAccounts) SetIncome(v IncomeOverride)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *OverrideAccounts) HasIncome() bool`

HasIncome returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


