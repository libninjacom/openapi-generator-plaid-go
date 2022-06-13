# AccountFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depository** | Pointer to **[]string** | A list of account subtypes to be filtered. | [optional] 
**Credit** | Pointer to **[]string** | A list of account subtypes to be filtered. | [optional] 
**Loan** | Pointer to **[]string** | A list of account subtypes to be filtered. | [optional] 
**Investment** | Pointer to **[]string** | A list of account subtypes to be filtered. | [optional] 

## Methods

### NewAccountFilter

`func NewAccountFilter() *AccountFilter`

NewAccountFilter instantiates a new AccountFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountFilterWithDefaults

`func NewAccountFilterWithDefaults() *AccountFilter`

NewAccountFilterWithDefaults instantiates a new AccountFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepository

`func (o *AccountFilter) GetDepository() []string`

GetDepository returns the Depository field if non-nil, zero value otherwise.

### GetDepositoryOk

`func (o *AccountFilter) GetDepositoryOk() (*[]string, bool)`

GetDepositoryOk returns a tuple with the Depository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepository

`func (o *AccountFilter) SetDepository(v []string)`

SetDepository sets Depository field to given value.

### HasDepository

`func (o *AccountFilter) HasDepository() bool`

HasDepository returns a boolean if a field has been set.

### GetCredit

`func (o *AccountFilter) GetCredit() []string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *AccountFilter) GetCreditOk() (*[]string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *AccountFilter) SetCredit(v []string)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *AccountFilter) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetLoan

`func (o *AccountFilter) GetLoan() []string`

GetLoan returns the Loan field if non-nil, zero value otherwise.

### GetLoanOk

`func (o *AccountFilter) GetLoanOk() (*[]string, bool)`

GetLoanOk returns a tuple with the Loan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoan

`func (o *AccountFilter) SetLoan(v []string)`

SetLoan sets Loan field to given value.

### HasLoan

`func (o *AccountFilter) HasLoan() bool`

HasLoan returns a boolean if a field has been set.

### GetInvestment

`func (o *AccountFilter) GetInvestment() []string`

GetInvestment returns the Investment field if non-nil, zero value otherwise.

### GetInvestmentOk

`func (o *AccountFilter) GetInvestmentOk() (*[]string, bool)`

GetInvestmentOk returns a tuple with the Investment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestment

`func (o *AccountFilter) SetInvestment(v []string)`

SetInvestment sets Investment field to given value.

### HasInvestment

`func (o *AccountFilter) HasInvestment() bool`

HasInvestment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


