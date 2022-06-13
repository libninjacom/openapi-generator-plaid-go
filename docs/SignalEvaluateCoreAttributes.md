# SignalEvaluateCoreAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnauthorizedTransactionsCount7d** | Pointer to **int32** | We parse and analyze historical transaction metadata to identify the number of possible past returns due to unauthorized transactions over the past 7 days from the account that will be debited. | [optional] 
**UnauthorizedTransactionsCount30d** | Pointer to **int32** | We parse and analyze historical transaction metadata to identify the number of possible past returns due to unauthorized transactions over the past 30 days from the account that will be debited. | [optional] 
**UnauthorizedTransactionsCount60d** | Pointer to **int32** | We parse and analyze historical transaction metadata to identify the number of possible past returns due to unauthorized transactions over the past 60 days from the account that will be debited. | [optional] 
**UnauthorizedTransactionsCount90d** | Pointer to **int32** | We parse and analyze historical transaction metadata to identify the number of possible past returns due to unauthorized transactions over the past 90 days from the account that will be debited. | [optional] 
**NsfOverdraftTransactionsCount7d** | Pointer to **int32** | We parse and analyze historical transaction metadata to identify the number of possible past returns due to non-sufficient funds/overdrafts over the past 7 days from the account that will be debited. | [optional] 
**NsfOverdraftTransactionsCount30d** | Pointer to **int32** | We parse and analyze historical transaction metadata to identify the number of possible past returns due to non-sufficient funds/overdrafts over the past 30 days from the account that will be debited. | [optional] 
**NsfOverdraftTransactionsCount60d** | Pointer to **int32** | We parse and analyze historical transaction metadata to identify the number of possible past returns due to non-sufficient funds/overdrafts over the past 60 days from the account that will be debited. | [optional] 
**NsfOverdraftTransactionsCount90d** | Pointer to **int32** | We parse and analyze historical transaction metadata to identify the number of possible past returns due to non-sufficient funds/overdrafts over the past 90 days from the account that will be debited. | [optional] 
**DaysSinceFirstPlaidConnection** | Pointer to **NullableInt32** | The number of days since the first time the Item was connected to an application via Plaid | [optional] 
**PlaidConnectionsCount7d** | Pointer to **NullableInt32** | The number of times the Item has been connected to applications via Plaid over the past 7 days | [optional] 
**PlaidConnectionsCount30d** | Pointer to **NullableInt32** | The number of times the Item has been connected to applications via Plaid over the past 30 days | [optional] 
**TotalPlaidConnectionsCount** | Pointer to **NullableInt32** | The total number of times the Item has been connected to applications via Plaid | [optional] 
**IsSavingsOrMoneyMarketAccount** | Pointer to **bool** | Indicates if the ACH transaction funding account is a savings/money market account | [optional] 
**TotalCreditTransactionsAmount10d** | Pointer to **float32** | The total credit (inflow) transaction amount over the past 10 days from the account that will be debited | [optional] 
**TotalDebitTransactionsAmount10d** | Pointer to **float32** | The total debit (outflow) transaction amount over the past 10 days from the account that will be debited | [optional] 
**P50CreditTransactionsAmount28d** | Pointer to **NullableFloat32** | The 50th percentile of all credit (inflow) transaction amounts over the past 28 days from the account that will be debited | [optional] 
**P50DebitTransactionsAmount28d** | Pointer to **NullableFloat32** | The 50th percentile of all debit (outflow) transaction amounts over the past 28 days from the account that will be debited | [optional] 
**P95CreditTransactionsAmount28d** | Pointer to **NullableFloat32** | The 95th percentile of all credit (inflow) transaction amounts over the past 28 days from the account that will be debited | [optional] 
**P95DebitTransactionsAmount28d** | Pointer to **NullableFloat32** | The 95th percentile of all debit (outflow) transaction amounts over the past 28 days from the account that will be debited | [optional] 
**DaysWithNegativeBalanceCount90d** | Pointer to **NullableInt32** | The number of days within the past 90 days when the account that will be debited had a negative end-of-day available balance | [optional] 
**P90EodBalance30d** | Pointer to **NullableFloat32** | The 90th percentile of the end-of-day available balance over the past 30 days of the account that will be debited | [optional] 
**P90EodBalance60d** | Pointer to **NullableFloat32** | The 90th percentile of the end-of-day available balance over the past 60 days of the account that will be debited | [optional] 
**P90EodBalance90d** | Pointer to **NullableFloat32** | The 90th percentile of the end-of-day available balance over the past 90 days of the account that will be debited | [optional] 
**P10EodBalance30d** | Pointer to **NullableFloat32** | The 10th percentile of the end-of-day available balance over the past 30 days of the account that will be debited | [optional] 
**P10EodBalance60d** | Pointer to **NullableFloat32** | The 10th percentile of the end-of-day available balance over the past 60 days of the account that will be debited | [optional] 
**P10EodBalance90d** | Pointer to **NullableFloat32** | The 10th percentile of the end-of-day available balance over the past 90 days of the account that will be debited | [optional] 
**AvailableBalance** | Pointer to **NullableFloat32** | Available balance, as of the &#x60;balance_last_updated&#x60; time. The available balance is the current balance less any outstanding holds or debits that have not yet posted to the account. | [optional] 
**CurrentBalance** | Pointer to **NullableFloat32** | Current balance, as of the &#x60;balance_last_updated&#x60; time. The current balance is the total amount of funds in the account. | [optional] 
**BalanceLastUpdated** | Pointer to **NullableTime** | Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DDTHH:mm:ssZ) indicating the last time that the balance for the given account has been updated. | [optional] 
**PhoneChangeCount28d** | Pointer to **NullableInt32** | The number of times the account&#39;s phone numbers on file have changed over the past 28 days | [optional] 
**PhoneChangeCount90d** | Pointer to **NullableInt32** | The number of times the account&#39;s phone numbers on file have changed over the past 90 days | [optional] 
**EmailChangeCount28d** | Pointer to **NullableInt32** | The number of times the account&#39;s email addresses on file have changed over the past 28 days | [optional] 
**EmailChangeCount90d** | Pointer to **NullableInt32** | The number of times the account&#39;s email addresses on file have changed over the past 90 days | [optional] 
**AddressChangeCount28d** | Pointer to **NullableInt32** | The number of times the account&#39;s addresses on file have changed over the past 28 days | [optional] 
**AddressChangeCount90d** | Pointer to **NullableInt32** | The number of times the account&#39;s addresses on file have changed over the past 90 days | [optional] 

## Methods

### NewSignalEvaluateCoreAttributes

`func NewSignalEvaluateCoreAttributes() *SignalEvaluateCoreAttributes`

NewSignalEvaluateCoreAttributes instantiates a new SignalEvaluateCoreAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalEvaluateCoreAttributesWithDefaults

`func NewSignalEvaluateCoreAttributesWithDefaults() *SignalEvaluateCoreAttributes`

NewSignalEvaluateCoreAttributesWithDefaults instantiates a new SignalEvaluateCoreAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnauthorizedTransactionsCount7d

`func (o *SignalEvaluateCoreAttributes) GetUnauthorizedTransactionsCount7d() int32`

GetUnauthorizedTransactionsCount7d returns the UnauthorizedTransactionsCount7d field if non-nil, zero value otherwise.

### GetUnauthorizedTransactionsCount7dOk

`func (o *SignalEvaluateCoreAttributes) GetUnauthorizedTransactionsCount7dOk() (*int32, bool)`

GetUnauthorizedTransactionsCount7dOk returns a tuple with the UnauthorizedTransactionsCount7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthorizedTransactionsCount7d

`func (o *SignalEvaluateCoreAttributes) SetUnauthorizedTransactionsCount7d(v int32)`

SetUnauthorizedTransactionsCount7d sets UnauthorizedTransactionsCount7d field to given value.

### HasUnauthorizedTransactionsCount7d

`func (o *SignalEvaluateCoreAttributes) HasUnauthorizedTransactionsCount7d() bool`

HasUnauthorizedTransactionsCount7d returns a boolean if a field has been set.

### GetUnauthorizedTransactionsCount30d

`func (o *SignalEvaluateCoreAttributes) GetUnauthorizedTransactionsCount30d() int32`

GetUnauthorizedTransactionsCount30d returns the UnauthorizedTransactionsCount30d field if non-nil, zero value otherwise.

### GetUnauthorizedTransactionsCount30dOk

`func (o *SignalEvaluateCoreAttributes) GetUnauthorizedTransactionsCount30dOk() (*int32, bool)`

GetUnauthorizedTransactionsCount30dOk returns a tuple with the UnauthorizedTransactionsCount30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthorizedTransactionsCount30d

`func (o *SignalEvaluateCoreAttributes) SetUnauthorizedTransactionsCount30d(v int32)`

SetUnauthorizedTransactionsCount30d sets UnauthorizedTransactionsCount30d field to given value.

### HasUnauthorizedTransactionsCount30d

`func (o *SignalEvaluateCoreAttributes) HasUnauthorizedTransactionsCount30d() bool`

HasUnauthorizedTransactionsCount30d returns a boolean if a field has been set.

### GetUnauthorizedTransactionsCount60d

`func (o *SignalEvaluateCoreAttributes) GetUnauthorizedTransactionsCount60d() int32`

GetUnauthorizedTransactionsCount60d returns the UnauthorizedTransactionsCount60d field if non-nil, zero value otherwise.

### GetUnauthorizedTransactionsCount60dOk

`func (o *SignalEvaluateCoreAttributes) GetUnauthorizedTransactionsCount60dOk() (*int32, bool)`

GetUnauthorizedTransactionsCount60dOk returns a tuple with the UnauthorizedTransactionsCount60d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthorizedTransactionsCount60d

`func (o *SignalEvaluateCoreAttributes) SetUnauthorizedTransactionsCount60d(v int32)`

SetUnauthorizedTransactionsCount60d sets UnauthorizedTransactionsCount60d field to given value.

### HasUnauthorizedTransactionsCount60d

`func (o *SignalEvaluateCoreAttributes) HasUnauthorizedTransactionsCount60d() bool`

HasUnauthorizedTransactionsCount60d returns a boolean if a field has been set.

### GetUnauthorizedTransactionsCount90d

`func (o *SignalEvaluateCoreAttributes) GetUnauthorizedTransactionsCount90d() int32`

GetUnauthorizedTransactionsCount90d returns the UnauthorizedTransactionsCount90d field if non-nil, zero value otherwise.

### GetUnauthorizedTransactionsCount90dOk

`func (o *SignalEvaluateCoreAttributes) GetUnauthorizedTransactionsCount90dOk() (*int32, bool)`

GetUnauthorizedTransactionsCount90dOk returns a tuple with the UnauthorizedTransactionsCount90d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthorizedTransactionsCount90d

`func (o *SignalEvaluateCoreAttributes) SetUnauthorizedTransactionsCount90d(v int32)`

SetUnauthorizedTransactionsCount90d sets UnauthorizedTransactionsCount90d field to given value.

### HasUnauthorizedTransactionsCount90d

`func (o *SignalEvaluateCoreAttributes) HasUnauthorizedTransactionsCount90d() bool`

HasUnauthorizedTransactionsCount90d returns a boolean if a field has been set.

### GetNsfOverdraftTransactionsCount7d

`func (o *SignalEvaluateCoreAttributes) GetNsfOverdraftTransactionsCount7d() int32`

GetNsfOverdraftTransactionsCount7d returns the NsfOverdraftTransactionsCount7d field if non-nil, zero value otherwise.

### GetNsfOverdraftTransactionsCount7dOk

`func (o *SignalEvaluateCoreAttributes) GetNsfOverdraftTransactionsCount7dOk() (*int32, bool)`

GetNsfOverdraftTransactionsCount7dOk returns a tuple with the NsfOverdraftTransactionsCount7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfOverdraftTransactionsCount7d

`func (o *SignalEvaluateCoreAttributes) SetNsfOverdraftTransactionsCount7d(v int32)`

SetNsfOverdraftTransactionsCount7d sets NsfOverdraftTransactionsCount7d field to given value.

### HasNsfOverdraftTransactionsCount7d

`func (o *SignalEvaluateCoreAttributes) HasNsfOverdraftTransactionsCount7d() bool`

HasNsfOverdraftTransactionsCount7d returns a boolean if a field has been set.

### GetNsfOverdraftTransactionsCount30d

`func (o *SignalEvaluateCoreAttributes) GetNsfOverdraftTransactionsCount30d() int32`

GetNsfOverdraftTransactionsCount30d returns the NsfOverdraftTransactionsCount30d field if non-nil, zero value otherwise.

### GetNsfOverdraftTransactionsCount30dOk

`func (o *SignalEvaluateCoreAttributes) GetNsfOverdraftTransactionsCount30dOk() (*int32, bool)`

GetNsfOverdraftTransactionsCount30dOk returns a tuple with the NsfOverdraftTransactionsCount30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfOverdraftTransactionsCount30d

`func (o *SignalEvaluateCoreAttributes) SetNsfOverdraftTransactionsCount30d(v int32)`

SetNsfOverdraftTransactionsCount30d sets NsfOverdraftTransactionsCount30d field to given value.

### HasNsfOverdraftTransactionsCount30d

`func (o *SignalEvaluateCoreAttributes) HasNsfOverdraftTransactionsCount30d() bool`

HasNsfOverdraftTransactionsCount30d returns a boolean if a field has been set.

### GetNsfOverdraftTransactionsCount60d

`func (o *SignalEvaluateCoreAttributes) GetNsfOverdraftTransactionsCount60d() int32`

GetNsfOverdraftTransactionsCount60d returns the NsfOverdraftTransactionsCount60d field if non-nil, zero value otherwise.

### GetNsfOverdraftTransactionsCount60dOk

`func (o *SignalEvaluateCoreAttributes) GetNsfOverdraftTransactionsCount60dOk() (*int32, bool)`

GetNsfOverdraftTransactionsCount60dOk returns a tuple with the NsfOverdraftTransactionsCount60d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfOverdraftTransactionsCount60d

`func (o *SignalEvaluateCoreAttributes) SetNsfOverdraftTransactionsCount60d(v int32)`

SetNsfOverdraftTransactionsCount60d sets NsfOverdraftTransactionsCount60d field to given value.

### HasNsfOverdraftTransactionsCount60d

`func (o *SignalEvaluateCoreAttributes) HasNsfOverdraftTransactionsCount60d() bool`

HasNsfOverdraftTransactionsCount60d returns a boolean if a field has been set.

### GetNsfOverdraftTransactionsCount90d

`func (o *SignalEvaluateCoreAttributes) GetNsfOverdraftTransactionsCount90d() int32`

GetNsfOverdraftTransactionsCount90d returns the NsfOverdraftTransactionsCount90d field if non-nil, zero value otherwise.

### GetNsfOverdraftTransactionsCount90dOk

`func (o *SignalEvaluateCoreAttributes) GetNsfOverdraftTransactionsCount90dOk() (*int32, bool)`

GetNsfOverdraftTransactionsCount90dOk returns a tuple with the NsfOverdraftTransactionsCount90d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfOverdraftTransactionsCount90d

`func (o *SignalEvaluateCoreAttributes) SetNsfOverdraftTransactionsCount90d(v int32)`

SetNsfOverdraftTransactionsCount90d sets NsfOverdraftTransactionsCount90d field to given value.

### HasNsfOverdraftTransactionsCount90d

`func (o *SignalEvaluateCoreAttributes) HasNsfOverdraftTransactionsCount90d() bool`

HasNsfOverdraftTransactionsCount90d returns a boolean if a field has been set.

### GetDaysSinceFirstPlaidConnection

`func (o *SignalEvaluateCoreAttributes) GetDaysSinceFirstPlaidConnection() int32`

GetDaysSinceFirstPlaidConnection returns the DaysSinceFirstPlaidConnection field if non-nil, zero value otherwise.

### GetDaysSinceFirstPlaidConnectionOk

`func (o *SignalEvaluateCoreAttributes) GetDaysSinceFirstPlaidConnectionOk() (*int32, bool)`

GetDaysSinceFirstPlaidConnectionOk returns a tuple with the DaysSinceFirstPlaidConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysSinceFirstPlaidConnection

`func (o *SignalEvaluateCoreAttributes) SetDaysSinceFirstPlaidConnection(v int32)`

SetDaysSinceFirstPlaidConnection sets DaysSinceFirstPlaidConnection field to given value.

### HasDaysSinceFirstPlaidConnection

`func (o *SignalEvaluateCoreAttributes) HasDaysSinceFirstPlaidConnection() bool`

HasDaysSinceFirstPlaidConnection returns a boolean if a field has been set.

### SetDaysSinceFirstPlaidConnectionNil

`func (o *SignalEvaluateCoreAttributes) SetDaysSinceFirstPlaidConnectionNil(b bool)`

 SetDaysSinceFirstPlaidConnectionNil sets the value for DaysSinceFirstPlaidConnection to be an explicit nil

### UnsetDaysSinceFirstPlaidConnection
`func (o *SignalEvaluateCoreAttributes) UnsetDaysSinceFirstPlaidConnection()`

UnsetDaysSinceFirstPlaidConnection ensures that no value is present for DaysSinceFirstPlaidConnection, not even an explicit nil
### GetPlaidConnectionsCount7d

`func (o *SignalEvaluateCoreAttributes) GetPlaidConnectionsCount7d() int32`

GetPlaidConnectionsCount7d returns the PlaidConnectionsCount7d field if non-nil, zero value otherwise.

### GetPlaidConnectionsCount7dOk

`func (o *SignalEvaluateCoreAttributes) GetPlaidConnectionsCount7dOk() (*int32, bool)`

GetPlaidConnectionsCount7dOk returns a tuple with the PlaidConnectionsCount7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaidConnectionsCount7d

`func (o *SignalEvaluateCoreAttributes) SetPlaidConnectionsCount7d(v int32)`

SetPlaidConnectionsCount7d sets PlaidConnectionsCount7d field to given value.

### HasPlaidConnectionsCount7d

`func (o *SignalEvaluateCoreAttributes) HasPlaidConnectionsCount7d() bool`

HasPlaidConnectionsCount7d returns a boolean if a field has been set.

### SetPlaidConnectionsCount7dNil

`func (o *SignalEvaluateCoreAttributes) SetPlaidConnectionsCount7dNil(b bool)`

 SetPlaidConnectionsCount7dNil sets the value for PlaidConnectionsCount7d to be an explicit nil

### UnsetPlaidConnectionsCount7d
`func (o *SignalEvaluateCoreAttributes) UnsetPlaidConnectionsCount7d()`

UnsetPlaidConnectionsCount7d ensures that no value is present for PlaidConnectionsCount7d, not even an explicit nil
### GetPlaidConnectionsCount30d

`func (o *SignalEvaluateCoreAttributes) GetPlaidConnectionsCount30d() int32`

GetPlaidConnectionsCount30d returns the PlaidConnectionsCount30d field if non-nil, zero value otherwise.

### GetPlaidConnectionsCount30dOk

`func (o *SignalEvaluateCoreAttributes) GetPlaidConnectionsCount30dOk() (*int32, bool)`

GetPlaidConnectionsCount30dOk returns a tuple with the PlaidConnectionsCount30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaidConnectionsCount30d

`func (o *SignalEvaluateCoreAttributes) SetPlaidConnectionsCount30d(v int32)`

SetPlaidConnectionsCount30d sets PlaidConnectionsCount30d field to given value.

### HasPlaidConnectionsCount30d

`func (o *SignalEvaluateCoreAttributes) HasPlaidConnectionsCount30d() bool`

HasPlaidConnectionsCount30d returns a boolean if a field has been set.

### SetPlaidConnectionsCount30dNil

`func (o *SignalEvaluateCoreAttributes) SetPlaidConnectionsCount30dNil(b bool)`

 SetPlaidConnectionsCount30dNil sets the value for PlaidConnectionsCount30d to be an explicit nil

### UnsetPlaidConnectionsCount30d
`func (o *SignalEvaluateCoreAttributes) UnsetPlaidConnectionsCount30d()`

UnsetPlaidConnectionsCount30d ensures that no value is present for PlaidConnectionsCount30d, not even an explicit nil
### GetTotalPlaidConnectionsCount

`func (o *SignalEvaluateCoreAttributes) GetTotalPlaidConnectionsCount() int32`

GetTotalPlaidConnectionsCount returns the TotalPlaidConnectionsCount field if non-nil, zero value otherwise.

### GetTotalPlaidConnectionsCountOk

`func (o *SignalEvaluateCoreAttributes) GetTotalPlaidConnectionsCountOk() (*int32, bool)`

GetTotalPlaidConnectionsCountOk returns a tuple with the TotalPlaidConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPlaidConnectionsCount

`func (o *SignalEvaluateCoreAttributes) SetTotalPlaidConnectionsCount(v int32)`

SetTotalPlaidConnectionsCount sets TotalPlaidConnectionsCount field to given value.

### HasTotalPlaidConnectionsCount

`func (o *SignalEvaluateCoreAttributes) HasTotalPlaidConnectionsCount() bool`

HasTotalPlaidConnectionsCount returns a boolean if a field has been set.

### SetTotalPlaidConnectionsCountNil

`func (o *SignalEvaluateCoreAttributes) SetTotalPlaidConnectionsCountNil(b bool)`

 SetTotalPlaidConnectionsCountNil sets the value for TotalPlaidConnectionsCount to be an explicit nil

### UnsetTotalPlaidConnectionsCount
`func (o *SignalEvaluateCoreAttributes) UnsetTotalPlaidConnectionsCount()`

UnsetTotalPlaidConnectionsCount ensures that no value is present for TotalPlaidConnectionsCount, not even an explicit nil
### GetIsSavingsOrMoneyMarketAccount

`func (o *SignalEvaluateCoreAttributes) GetIsSavingsOrMoneyMarketAccount() bool`

GetIsSavingsOrMoneyMarketAccount returns the IsSavingsOrMoneyMarketAccount field if non-nil, zero value otherwise.

### GetIsSavingsOrMoneyMarketAccountOk

`func (o *SignalEvaluateCoreAttributes) GetIsSavingsOrMoneyMarketAccountOk() (*bool, bool)`

GetIsSavingsOrMoneyMarketAccountOk returns a tuple with the IsSavingsOrMoneyMarketAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSavingsOrMoneyMarketAccount

`func (o *SignalEvaluateCoreAttributes) SetIsSavingsOrMoneyMarketAccount(v bool)`

SetIsSavingsOrMoneyMarketAccount sets IsSavingsOrMoneyMarketAccount field to given value.

### HasIsSavingsOrMoneyMarketAccount

`func (o *SignalEvaluateCoreAttributes) HasIsSavingsOrMoneyMarketAccount() bool`

HasIsSavingsOrMoneyMarketAccount returns a boolean if a field has been set.

### GetTotalCreditTransactionsAmount10d

`func (o *SignalEvaluateCoreAttributes) GetTotalCreditTransactionsAmount10d() float32`

GetTotalCreditTransactionsAmount10d returns the TotalCreditTransactionsAmount10d field if non-nil, zero value otherwise.

### GetTotalCreditTransactionsAmount10dOk

`func (o *SignalEvaluateCoreAttributes) GetTotalCreditTransactionsAmount10dOk() (*float32, bool)`

GetTotalCreditTransactionsAmount10dOk returns a tuple with the TotalCreditTransactionsAmount10d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditTransactionsAmount10d

`func (o *SignalEvaluateCoreAttributes) SetTotalCreditTransactionsAmount10d(v float32)`

SetTotalCreditTransactionsAmount10d sets TotalCreditTransactionsAmount10d field to given value.

### HasTotalCreditTransactionsAmount10d

`func (o *SignalEvaluateCoreAttributes) HasTotalCreditTransactionsAmount10d() bool`

HasTotalCreditTransactionsAmount10d returns a boolean if a field has been set.

### GetTotalDebitTransactionsAmount10d

`func (o *SignalEvaluateCoreAttributes) GetTotalDebitTransactionsAmount10d() float32`

GetTotalDebitTransactionsAmount10d returns the TotalDebitTransactionsAmount10d field if non-nil, zero value otherwise.

### GetTotalDebitTransactionsAmount10dOk

`func (o *SignalEvaluateCoreAttributes) GetTotalDebitTransactionsAmount10dOk() (*float32, bool)`

GetTotalDebitTransactionsAmount10dOk returns a tuple with the TotalDebitTransactionsAmount10d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebitTransactionsAmount10d

`func (o *SignalEvaluateCoreAttributes) SetTotalDebitTransactionsAmount10d(v float32)`

SetTotalDebitTransactionsAmount10d sets TotalDebitTransactionsAmount10d field to given value.

### HasTotalDebitTransactionsAmount10d

`func (o *SignalEvaluateCoreAttributes) HasTotalDebitTransactionsAmount10d() bool`

HasTotalDebitTransactionsAmount10d returns a boolean if a field has been set.

### GetP50CreditTransactionsAmount28d

`func (o *SignalEvaluateCoreAttributes) GetP50CreditTransactionsAmount28d() float32`

GetP50CreditTransactionsAmount28d returns the P50CreditTransactionsAmount28d field if non-nil, zero value otherwise.

### GetP50CreditTransactionsAmount28dOk

`func (o *SignalEvaluateCoreAttributes) GetP50CreditTransactionsAmount28dOk() (*float32, bool)`

GetP50CreditTransactionsAmount28dOk returns a tuple with the P50CreditTransactionsAmount28d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP50CreditTransactionsAmount28d

`func (o *SignalEvaluateCoreAttributes) SetP50CreditTransactionsAmount28d(v float32)`

SetP50CreditTransactionsAmount28d sets P50CreditTransactionsAmount28d field to given value.

### HasP50CreditTransactionsAmount28d

`func (o *SignalEvaluateCoreAttributes) HasP50CreditTransactionsAmount28d() bool`

HasP50CreditTransactionsAmount28d returns a boolean if a field has been set.

### SetP50CreditTransactionsAmount28dNil

`func (o *SignalEvaluateCoreAttributes) SetP50CreditTransactionsAmount28dNil(b bool)`

 SetP50CreditTransactionsAmount28dNil sets the value for P50CreditTransactionsAmount28d to be an explicit nil

### UnsetP50CreditTransactionsAmount28d
`func (o *SignalEvaluateCoreAttributes) UnsetP50CreditTransactionsAmount28d()`

UnsetP50CreditTransactionsAmount28d ensures that no value is present for P50CreditTransactionsAmount28d, not even an explicit nil
### GetP50DebitTransactionsAmount28d

`func (o *SignalEvaluateCoreAttributes) GetP50DebitTransactionsAmount28d() float32`

GetP50DebitTransactionsAmount28d returns the P50DebitTransactionsAmount28d field if non-nil, zero value otherwise.

### GetP50DebitTransactionsAmount28dOk

`func (o *SignalEvaluateCoreAttributes) GetP50DebitTransactionsAmount28dOk() (*float32, bool)`

GetP50DebitTransactionsAmount28dOk returns a tuple with the P50DebitTransactionsAmount28d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP50DebitTransactionsAmount28d

`func (o *SignalEvaluateCoreAttributes) SetP50DebitTransactionsAmount28d(v float32)`

SetP50DebitTransactionsAmount28d sets P50DebitTransactionsAmount28d field to given value.

### HasP50DebitTransactionsAmount28d

`func (o *SignalEvaluateCoreAttributes) HasP50DebitTransactionsAmount28d() bool`

HasP50DebitTransactionsAmount28d returns a boolean if a field has been set.

### SetP50DebitTransactionsAmount28dNil

`func (o *SignalEvaluateCoreAttributes) SetP50DebitTransactionsAmount28dNil(b bool)`

 SetP50DebitTransactionsAmount28dNil sets the value for P50DebitTransactionsAmount28d to be an explicit nil

### UnsetP50DebitTransactionsAmount28d
`func (o *SignalEvaluateCoreAttributes) UnsetP50DebitTransactionsAmount28d()`

UnsetP50DebitTransactionsAmount28d ensures that no value is present for P50DebitTransactionsAmount28d, not even an explicit nil
### GetP95CreditTransactionsAmount28d

`func (o *SignalEvaluateCoreAttributes) GetP95CreditTransactionsAmount28d() float32`

GetP95CreditTransactionsAmount28d returns the P95CreditTransactionsAmount28d field if non-nil, zero value otherwise.

### GetP95CreditTransactionsAmount28dOk

`func (o *SignalEvaluateCoreAttributes) GetP95CreditTransactionsAmount28dOk() (*float32, bool)`

GetP95CreditTransactionsAmount28dOk returns a tuple with the P95CreditTransactionsAmount28d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95CreditTransactionsAmount28d

`func (o *SignalEvaluateCoreAttributes) SetP95CreditTransactionsAmount28d(v float32)`

SetP95CreditTransactionsAmount28d sets P95CreditTransactionsAmount28d field to given value.

### HasP95CreditTransactionsAmount28d

`func (o *SignalEvaluateCoreAttributes) HasP95CreditTransactionsAmount28d() bool`

HasP95CreditTransactionsAmount28d returns a boolean if a field has been set.

### SetP95CreditTransactionsAmount28dNil

`func (o *SignalEvaluateCoreAttributes) SetP95CreditTransactionsAmount28dNil(b bool)`

 SetP95CreditTransactionsAmount28dNil sets the value for P95CreditTransactionsAmount28d to be an explicit nil

### UnsetP95CreditTransactionsAmount28d
`func (o *SignalEvaluateCoreAttributes) UnsetP95CreditTransactionsAmount28d()`

UnsetP95CreditTransactionsAmount28d ensures that no value is present for P95CreditTransactionsAmount28d, not even an explicit nil
### GetP95DebitTransactionsAmount28d

`func (o *SignalEvaluateCoreAttributes) GetP95DebitTransactionsAmount28d() float32`

GetP95DebitTransactionsAmount28d returns the P95DebitTransactionsAmount28d field if non-nil, zero value otherwise.

### GetP95DebitTransactionsAmount28dOk

`func (o *SignalEvaluateCoreAttributes) GetP95DebitTransactionsAmount28dOk() (*float32, bool)`

GetP95DebitTransactionsAmount28dOk returns a tuple with the P95DebitTransactionsAmount28d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95DebitTransactionsAmount28d

`func (o *SignalEvaluateCoreAttributes) SetP95DebitTransactionsAmount28d(v float32)`

SetP95DebitTransactionsAmount28d sets P95DebitTransactionsAmount28d field to given value.

### HasP95DebitTransactionsAmount28d

`func (o *SignalEvaluateCoreAttributes) HasP95DebitTransactionsAmount28d() bool`

HasP95DebitTransactionsAmount28d returns a boolean if a field has been set.

### SetP95DebitTransactionsAmount28dNil

`func (o *SignalEvaluateCoreAttributes) SetP95DebitTransactionsAmount28dNil(b bool)`

 SetP95DebitTransactionsAmount28dNil sets the value for P95DebitTransactionsAmount28d to be an explicit nil

### UnsetP95DebitTransactionsAmount28d
`func (o *SignalEvaluateCoreAttributes) UnsetP95DebitTransactionsAmount28d()`

UnsetP95DebitTransactionsAmount28d ensures that no value is present for P95DebitTransactionsAmount28d, not even an explicit nil
### GetDaysWithNegativeBalanceCount90d

`func (o *SignalEvaluateCoreAttributes) GetDaysWithNegativeBalanceCount90d() int32`

GetDaysWithNegativeBalanceCount90d returns the DaysWithNegativeBalanceCount90d field if non-nil, zero value otherwise.

### GetDaysWithNegativeBalanceCount90dOk

`func (o *SignalEvaluateCoreAttributes) GetDaysWithNegativeBalanceCount90dOk() (*int32, bool)`

GetDaysWithNegativeBalanceCount90dOk returns a tuple with the DaysWithNegativeBalanceCount90d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysWithNegativeBalanceCount90d

`func (o *SignalEvaluateCoreAttributes) SetDaysWithNegativeBalanceCount90d(v int32)`

SetDaysWithNegativeBalanceCount90d sets DaysWithNegativeBalanceCount90d field to given value.

### HasDaysWithNegativeBalanceCount90d

`func (o *SignalEvaluateCoreAttributes) HasDaysWithNegativeBalanceCount90d() bool`

HasDaysWithNegativeBalanceCount90d returns a boolean if a field has been set.

### SetDaysWithNegativeBalanceCount90dNil

`func (o *SignalEvaluateCoreAttributes) SetDaysWithNegativeBalanceCount90dNil(b bool)`

 SetDaysWithNegativeBalanceCount90dNil sets the value for DaysWithNegativeBalanceCount90d to be an explicit nil

### UnsetDaysWithNegativeBalanceCount90d
`func (o *SignalEvaluateCoreAttributes) UnsetDaysWithNegativeBalanceCount90d()`

UnsetDaysWithNegativeBalanceCount90d ensures that no value is present for DaysWithNegativeBalanceCount90d, not even an explicit nil
### GetP90EodBalance30d

`func (o *SignalEvaluateCoreAttributes) GetP90EodBalance30d() float32`

GetP90EodBalance30d returns the P90EodBalance30d field if non-nil, zero value otherwise.

### GetP90EodBalance30dOk

`func (o *SignalEvaluateCoreAttributes) GetP90EodBalance30dOk() (*float32, bool)`

GetP90EodBalance30dOk returns a tuple with the P90EodBalance30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP90EodBalance30d

`func (o *SignalEvaluateCoreAttributes) SetP90EodBalance30d(v float32)`

SetP90EodBalance30d sets P90EodBalance30d field to given value.

### HasP90EodBalance30d

`func (o *SignalEvaluateCoreAttributes) HasP90EodBalance30d() bool`

HasP90EodBalance30d returns a boolean if a field has been set.

### SetP90EodBalance30dNil

`func (o *SignalEvaluateCoreAttributes) SetP90EodBalance30dNil(b bool)`

 SetP90EodBalance30dNil sets the value for P90EodBalance30d to be an explicit nil

### UnsetP90EodBalance30d
`func (o *SignalEvaluateCoreAttributes) UnsetP90EodBalance30d()`

UnsetP90EodBalance30d ensures that no value is present for P90EodBalance30d, not even an explicit nil
### GetP90EodBalance60d

`func (o *SignalEvaluateCoreAttributes) GetP90EodBalance60d() float32`

GetP90EodBalance60d returns the P90EodBalance60d field if non-nil, zero value otherwise.

### GetP90EodBalance60dOk

`func (o *SignalEvaluateCoreAttributes) GetP90EodBalance60dOk() (*float32, bool)`

GetP90EodBalance60dOk returns a tuple with the P90EodBalance60d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP90EodBalance60d

`func (o *SignalEvaluateCoreAttributes) SetP90EodBalance60d(v float32)`

SetP90EodBalance60d sets P90EodBalance60d field to given value.

### HasP90EodBalance60d

`func (o *SignalEvaluateCoreAttributes) HasP90EodBalance60d() bool`

HasP90EodBalance60d returns a boolean if a field has been set.

### SetP90EodBalance60dNil

`func (o *SignalEvaluateCoreAttributes) SetP90EodBalance60dNil(b bool)`

 SetP90EodBalance60dNil sets the value for P90EodBalance60d to be an explicit nil

### UnsetP90EodBalance60d
`func (o *SignalEvaluateCoreAttributes) UnsetP90EodBalance60d()`

UnsetP90EodBalance60d ensures that no value is present for P90EodBalance60d, not even an explicit nil
### GetP90EodBalance90d

`func (o *SignalEvaluateCoreAttributes) GetP90EodBalance90d() float32`

GetP90EodBalance90d returns the P90EodBalance90d field if non-nil, zero value otherwise.

### GetP90EodBalance90dOk

`func (o *SignalEvaluateCoreAttributes) GetP90EodBalance90dOk() (*float32, bool)`

GetP90EodBalance90dOk returns a tuple with the P90EodBalance90d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP90EodBalance90d

`func (o *SignalEvaluateCoreAttributes) SetP90EodBalance90d(v float32)`

SetP90EodBalance90d sets P90EodBalance90d field to given value.

### HasP90EodBalance90d

`func (o *SignalEvaluateCoreAttributes) HasP90EodBalance90d() bool`

HasP90EodBalance90d returns a boolean if a field has been set.

### SetP90EodBalance90dNil

`func (o *SignalEvaluateCoreAttributes) SetP90EodBalance90dNil(b bool)`

 SetP90EodBalance90dNil sets the value for P90EodBalance90d to be an explicit nil

### UnsetP90EodBalance90d
`func (o *SignalEvaluateCoreAttributes) UnsetP90EodBalance90d()`

UnsetP90EodBalance90d ensures that no value is present for P90EodBalance90d, not even an explicit nil
### GetP10EodBalance30d

`func (o *SignalEvaluateCoreAttributes) GetP10EodBalance30d() float32`

GetP10EodBalance30d returns the P10EodBalance30d field if non-nil, zero value otherwise.

### GetP10EodBalance30dOk

`func (o *SignalEvaluateCoreAttributes) GetP10EodBalance30dOk() (*float32, bool)`

GetP10EodBalance30dOk returns a tuple with the P10EodBalance30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP10EodBalance30d

`func (o *SignalEvaluateCoreAttributes) SetP10EodBalance30d(v float32)`

SetP10EodBalance30d sets P10EodBalance30d field to given value.

### HasP10EodBalance30d

`func (o *SignalEvaluateCoreAttributes) HasP10EodBalance30d() bool`

HasP10EodBalance30d returns a boolean if a field has been set.

### SetP10EodBalance30dNil

`func (o *SignalEvaluateCoreAttributes) SetP10EodBalance30dNil(b bool)`

 SetP10EodBalance30dNil sets the value for P10EodBalance30d to be an explicit nil

### UnsetP10EodBalance30d
`func (o *SignalEvaluateCoreAttributes) UnsetP10EodBalance30d()`

UnsetP10EodBalance30d ensures that no value is present for P10EodBalance30d, not even an explicit nil
### GetP10EodBalance60d

`func (o *SignalEvaluateCoreAttributes) GetP10EodBalance60d() float32`

GetP10EodBalance60d returns the P10EodBalance60d field if non-nil, zero value otherwise.

### GetP10EodBalance60dOk

`func (o *SignalEvaluateCoreAttributes) GetP10EodBalance60dOk() (*float32, bool)`

GetP10EodBalance60dOk returns a tuple with the P10EodBalance60d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP10EodBalance60d

`func (o *SignalEvaluateCoreAttributes) SetP10EodBalance60d(v float32)`

SetP10EodBalance60d sets P10EodBalance60d field to given value.

### HasP10EodBalance60d

`func (o *SignalEvaluateCoreAttributes) HasP10EodBalance60d() bool`

HasP10EodBalance60d returns a boolean if a field has been set.

### SetP10EodBalance60dNil

`func (o *SignalEvaluateCoreAttributes) SetP10EodBalance60dNil(b bool)`

 SetP10EodBalance60dNil sets the value for P10EodBalance60d to be an explicit nil

### UnsetP10EodBalance60d
`func (o *SignalEvaluateCoreAttributes) UnsetP10EodBalance60d()`

UnsetP10EodBalance60d ensures that no value is present for P10EodBalance60d, not even an explicit nil
### GetP10EodBalance90d

`func (o *SignalEvaluateCoreAttributes) GetP10EodBalance90d() float32`

GetP10EodBalance90d returns the P10EodBalance90d field if non-nil, zero value otherwise.

### GetP10EodBalance90dOk

`func (o *SignalEvaluateCoreAttributes) GetP10EodBalance90dOk() (*float32, bool)`

GetP10EodBalance90dOk returns a tuple with the P10EodBalance90d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP10EodBalance90d

`func (o *SignalEvaluateCoreAttributes) SetP10EodBalance90d(v float32)`

SetP10EodBalance90d sets P10EodBalance90d field to given value.

### HasP10EodBalance90d

`func (o *SignalEvaluateCoreAttributes) HasP10EodBalance90d() bool`

HasP10EodBalance90d returns a boolean if a field has been set.

### SetP10EodBalance90dNil

`func (o *SignalEvaluateCoreAttributes) SetP10EodBalance90dNil(b bool)`

 SetP10EodBalance90dNil sets the value for P10EodBalance90d to be an explicit nil

### UnsetP10EodBalance90d
`func (o *SignalEvaluateCoreAttributes) UnsetP10EodBalance90d()`

UnsetP10EodBalance90d ensures that no value is present for P10EodBalance90d, not even an explicit nil
### GetAvailableBalance

`func (o *SignalEvaluateCoreAttributes) GetAvailableBalance() float32`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *SignalEvaluateCoreAttributes) GetAvailableBalanceOk() (*float32, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *SignalEvaluateCoreAttributes) SetAvailableBalance(v float32)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *SignalEvaluateCoreAttributes) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### SetAvailableBalanceNil

`func (o *SignalEvaluateCoreAttributes) SetAvailableBalanceNil(b bool)`

 SetAvailableBalanceNil sets the value for AvailableBalance to be an explicit nil

### UnsetAvailableBalance
`func (o *SignalEvaluateCoreAttributes) UnsetAvailableBalance()`

UnsetAvailableBalance ensures that no value is present for AvailableBalance, not even an explicit nil
### GetCurrentBalance

`func (o *SignalEvaluateCoreAttributes) GetCurrentBalance() float32`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *SignalEvaluateCoreAttributes) GetCurrentBalanceOk() (*float32, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalance

`func (o *SignalEvaluateCoreAttributes) SetCurrentBalance(v float32)`

SetCurrentBalance sets CurrentBalance field to given value.

### HasCurrentBalance

`func (o *SignalEvaluateCoreAttributes) HasCurrentBalance() bool`

HasCurrentBalance returns a boolean if a field has been set.

### SetCurrentBalanceNil

`func (o *SignalEvaluateCoreAttributes) SetCurrentBalanceNil(b bool)`

 SetCurrentBalanceNil sets the value for CurrentBalance to be an explicit nil

### UnsetCurrentBalance
`func (o *SignalEvaluateCoreAttributes) UnsetCurrentBalance()`

UnsetCurrentBalance ensures that no value is present for CurrentBalance, not even an explicit nil
### GetBalanceLastUpdated

`func (o *SignalEvaluateCoreAttributes) GetBalanceLastUpdated() time.Time`

GetBalanceLastUpdated returns the BalanceLastUpdated field if non-nil, zero value otherwise.

### GetBalanceLastUpdatedOk

`func (o *SignalEvaluateCoreAttributes) GetBalanceLastUpdatedOk() (*time.Time, bool)`

GetBalanceLastUpdatedOk returns a tuple with the BalanceLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceLastUpdated

`func (o *SignalEvaluateCoreAttributes) SetBalanceLastUpdated(v time.Time)`

SetBalanceLastUpdated sets BalanceLastUpdated field to given value.

### HasBalanceLastUpdated

`func (o *SignalEvaluateCoreAttributes) HasBalanceLastUpdated() bool`

HasBalanceLastUpdated returns a boolean if a field has been set.

### SetBalanceLastUpdatedNil

`func (o *SignalEvaluateCoreAttributes) SetBalanceLastUpdatedNil(b bool)`

 SetBalanceLastUpdatedNil sets the value for BalanceLastUpdated to be an explicit nil

### UnsetBalanceLastUpdated
`func (o *SignalEvaluateCoreAttributes) UnsetBalanceLastUpdated()`

UnsetBalanceLastUpdated ensures that no value is present for BalanceLastUpdated, not even an explicit nil
### GetPhoneChangeCount28d

`func (o *SignalEvaluateCoreAttributes) GetPhoneChangeCount28d() int32`

GetPhoneChangeCount28d returns the PhoneChangeCount28d field if non-nil, zero value otherwise.

### GetPhoneChangeCount28dOk

`func (o *SignalEvaluateCoreAttributes) GetPhoneChangeCount28dOk() (*int32, bool)`

GetPhoneChangeCount28dOk returns a tuple with the PhoneChangeCount28d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneChangeCount28d

`func (o *SignalEvaluateCoreAttributes) SetPhoneChangeCount28d(v int32)`

SetPhoneChangeCount28d sets PhoneChangeCount28d field to given value.

### HasPhoneChangeCount28d

`func (o *SignalEvaluateCoreAttributes) HasPhoneChangeCount28d() bool`

HasPhoneChangeCount28d returns a boolean if a field has been set.

### SetPhoneChangeCount28dNil

`func (o *SignalEvaluateCoreAttributes) SetPhoneChangeCount28dNil(b bool)`

 SetPhoneChangeCount28dNil sets the value for PhoneChangeCount28d to be an explicit nil

### UnsetPhoneChangeCount28d
`func (o *SignalEvaluateCoreAttributes) UnsetPhoneChangeCount28d()`

UnsetPhoneChangeCount28d ensures that no value is present for PhoneChangeCount28d, not even an explicit nil
### GetPhoneChangeCount90d

`func (o *SignalEvaluateCoreAttributes) GetPhoneChangeCount90d() int32`

GetPhoneChangeCount90d returns the PhoneChangeCount90d field if non-nil, zero value otherwise.

### GetPhoneChangeCount90dOk

`func (o *SignalEvaluateCoreAttributes) GetPhoneChangeCount90dOk() (*int32, bool)`

GetPhoneChangeCount90dOk returns a tuple with the PhoneChangeCount90d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneChangeCount90d

`func (o *SignalEvaluateCoreAttributes) SetPhoneChangeCount90d(v int32)`

SetPhoneChangeCount90d sets PhoneChangeCount90d field to given value.

### HasPhoneChangeCount90d

`func (o *SignalEvaluateCoreAttributes) HasPhoneChangeCount90d() bool`

HasPhoneChangeCount90d returns a boolean if a field has been set.

### SetPhoneChangeCount90dNil

`func (o *SignalEvaluateCoreAttributes) SetPhoneChangeCount90dNil(b bool)`

 SetPhoneChangeCount90dNil sets the value for PhoneChangeCount90d to be an explicit nil

### UnsetPhoneChangeCount90d
`func (o *SignalEvaluateCoreAttributes) UnsetPhoneChangeCount90d()`

UnsetPhoneChangeCount90d ensures that no value is present for PhoneChangeCount90d, not even an explicit nil
### GetEmailChangeCount28d

`func (o *SignalEvaluateCoreAttributes) GetEmailChangeCount28d() int32`

GetEmailChangeCount28d returns the EmailChangeCount28d field if non-nil, zero value otherwise.

### GetEmailChangeCount28dOk

`func (o *SignalEvaluateCoreAttributes) GetEmailChangeCount28dOk() (*int32, bool)`

GetEmailChangeCount28dOk returns a tuple with the EmailChangeCount28d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailChangeCount28d

`func (o *SignalEvaluateCoreAttributes) SetEmailChangeCount28d(v int32)`

SetEmailChangeCount28d sets EmailChangeCount28d field to given value.

### HasEmailChangeCount28d

`func (o *SignalEvaluateCoreAttributes) HasEmailChangeCount28d() bool`

HasEmailChangeCount28d returns a boolean if a field has been set.

### SetEmailChangeCount28dNil

`func (o *SignalEvaluateCoreAttributes) SetEmailChangeCount28dNil(b bool)`

 SetEmailChangeCount28dNil sets the value for EmailChangeCount28d to be an explicit nil

### UnsetEmailChangeCount28d
`func (o *SignalEvaluateCoreAttributes) UnsetEmailChangeCount28d()`

UnsetEmailChangeCount28d ensures that no value is present for EmailChangeCount28d, not even an explicit nil
### GetEmailChangeCount90d

`func (o *SignalEvaluateCoreAttributes) GetEmailChangeCount90d() int32`

GetEmailChangeCount90d returns the EmailChangeCount90d field if non-nil, zero value otherwise.

### GetEmailChangeCount90dOk

`func (o *SignalEvaluateCoreAttributes) GetEmailChangeCount90dOk() (*int32, bool)`

GetEmailChangeCount90dOk returns a tuple with the EmailChangeCount90d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailChangeCount90d

`func (o *SignalEvaluateCoreAttributes) SetEmailChangeCount90d(v int32)`

SetEmailChangeCount90d sets EmailChangeCount90d field to given value.

### HasEmailChangeCount90d

`func (o *SignalEvaluateCoreAttributes) HasEmailChangeCount90d() bool`

HasEmailChangeCount90d returns a boolean if a field has been set.

### SetEmailChangeCount90dNil

`func (o *SignalEvaluateCoreAttributes) SetEmailChangeCount90dNil(b bool)`

 SetEmailChangeCount90dNil sets the value for EmailChangeCount90d to be an explicit nil

### UnsetEmailChangeCount90d
`func (o *SignalEvaluateCoreAttributes) UnsetEmailChangeCount90d()`

UnsetEmailChangeCount90d ensures that no value is present for EmailChangeCount90d, not even an explicit nil
### GetAddressChangeCount28d

`func (o *SignalEvaluateCoreAttributes) GetAddressChangeCount28d() int32`

GetAddressChangeCount28d returns the AddressChangeCount28d field if non-nil, zero value otherwise.

### GetAddressChangeCount28dOk

`func (o *SignalEvaluateCoreAttributes) GetAddressChangeCount28dOk() (*int32, bool)`

GetAddressChangeCount28dOk returns a tuple with the AddressChangeCount28d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressChangeCount28d

`func (o *SignalEvaluateCoreAttributes) SetAddressChangeCount28d(v int32)`

SetAddressChangeCount28d sets AddressChangeCount28d field to given value.

### HasAddressChangeCount28d

`func (o *SignalEvaluateCoreAttributes) HasAddressChangeCount28d() bool`

HasAddressChangeCount28d returns a boolean if a field has been set.

### SetAddressChangeCount28dNil

`func (o *SignalEvaluateCoreAttributes) SetAddressChangeCount28dNil(b bool)`

 SetAddressChangeCount28dNil sets the value for AddressChangeCount28d to be an explicit nil

### UnsetAddressChangeCount28d
`func (o *SignalEvaluateCoreAttributes) UnsetAddressChangeCount28d()`

UnsetAddressChangeCount28d ensures that no value is present for AddressChangeCount28d, not even an explicit nil
### GetAddressChangeCount90d

`func (o *SignalEvaluateCoreAttributes) GetAddressChangeCount90d() int32`

GetAddressChangeCount90d returns the AddressChangeCount90d field if non-nil, zero value otherwise.

### GetAddressChangeCount90dOk

`func (o *SignalEvaluateCoreAttributes) GetAddressChangeCount90dOk() (*int32, bool)`

GetAddressChangeCount90dOk returns a tuple with the AddressChangeCount90d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressChangeCount90d

`func (o *SignalEvaluateCoreAttributes) SetAddressChangeCount90d(v int32)`

SetAddressChangeCount90d sets AddressChangeCount90d field to given value.

### HasAddressChangeCount90d

`func (o *SignalEvaluateCoreAttributes) HasAddressChangeCount90d() bool`

HasAddressChangeCount90d returns a boolean if a field has been set.

### SetAddressChangeCount90dNil

`func (o *SignalEvaluateCoreAttributes) SetAddressChangeCount90dNil(b bool)`

 SetAddressChangeCount90dNil sets the value for AddressChangeCount90d to be an explicit nil

### UnsetAddressChangeCount90d
`func (o *SignalEvaluateCoreAttributes) UnsetAddressChangeCount90d()`

UnsetAddressChangeCount90d ensures that no value is present for AddressChangeCount90d, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


