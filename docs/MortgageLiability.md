# MortgageLiability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account that this liability belongs to. | 
**AccountNumber** | **string** | The account number of the loan. | 
**CurrentLateFee** | **NullableFloat32** | The current outstanding amount charged for late payment. | 
**EscrowBalance** | **NullableFloat32** | Total amount held in escrow to pay taxes and insurance on behalf of the borrower. | 
**HasPmi** | **NullableBool** | Indicates whether the borrower has private mortgage insurance in effect. | 
**HasPrepaymentPenalty** | **NullableBool** | Indicates whether the borrower will pay a penalty for early payoff of mortgage. | 
**InterestRate** | [**MortgageInterestRate**](MortgageInterestRate.md) |  | 
**LastPaymentAmount** | **NullableFloat32** | The amount of the last payment. | 
**LastPaymentDate** | **NullableString** | The date of the last payment. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). | 
**LoanTypeDescription** | **NullableString** | Description of the type of loan, for example &#x60;conventional&#x60;, &#x60;fixed&#x60;, or &#x60;variable&#x60;. This field is provided directly from the loan servicer and does not have an enumerated set of possible values. | 
**LoanTerm** | **NullableString** | Full duration of mortgage as at origination (e.g. &#x60;10 year&#x60;). | 
**MaturityDate** | **NullableString** | Original date on which mortgage is due in full. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). | 
**NextMonthlyPayment** | **NullableFloat32** | The amount of the next payment. | 
**NextPaymentDueDate** | **NullableString** | The due date for the next payment. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). | 
**OriginationDate** | **NullableString** | The date on which the loan was initially lent. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). | 
**OriginationPrincipalAmount** | **NullableFloat32** | The original principal balance of the mortgage. | 
**PastDueAmount** | **NullableFloat32** | Amount of loan (principal + interest) past due for payment. | 
**PropertyAddress** | [**MortgagePropertyAddress**](MortgagePropertyAddress.md) |  | 
**YtdInterestPaid** | **NullableFloat32** | The year to date (YTD) interest paid. | 
**YtdPrincipalPaid** | **NullableFloat32** | The YTD principal paid. | 

## Methods

### NewMortgageLiability

`func NewMortgageLiability(accountId string, accountNumber string, currentLateFee NullableFloat32, escrowBalance NullableFloat32, hasPmi NullableBool, hasPrepaymentPenalty NullableBool, interestRate MortgageInterestRate, lastPaymentAmount NullableFloat32, lastPaymentDate NullableString, loanTypeDescription NullableString, loanTerm NullableString, maturityDate NullableString, nextMonthlyPayment NullableFloat32, nextPaymentDueDate NullableString, originationDate NullableString, originationPrincipalAmount NullableFloat32, pastDueAmount NullableFloat32, propertyAddress MortgagePropertyAddress, ytdInterestPaid NullableFloat32, ytdPrincipalPaid NullableFloat32, ) *MortgageLiability`

NewMortgageLiability instantiates a new MortgageLiability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMortgageLiabilityWithDefaults

`func NewMortgageLiabilityWithDefaults() *MortgageLiability`

NewMortgageLiabilityWithDefaults instantiates a new MortgageLiability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *MortgageLiability) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MortgageLiability) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MortgageLiability) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountNumber

`func (o *MortgageLiability) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *MortgageLiability) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *MortgageLiability) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetCurrentLateFee

`func (o *MortgageLiability) GetCurrentLateFee() float32`

GetCurrentLateFee returns the CurrentLateFee field if non-nil, zero value otherwise.

### GetCurrentLateFeeOk

`func (o *MortgageLiability) GetCurrentLateFeeOk() (*float32, bool)`

GetCurrentLateFeeOk returns a tuple with the CurrentLateFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLateFee

`func (o *MortgageLiability) SetCurrentLateFee(v float32)`

SetCurrentLateFee sets CurrentLateFee field to given value.


### SetCurrentLateFeeNil

`func (o *MortgageLiability) SetCurrentLateFeeNil(b bool)`

 SetCurrentLateFeeNil sets the value for CurrentLateFee to be an explicit nil

### UnsetCurrentLateFee
`func (o *MortgageLiability) UnsetCurrentLateFee()`

UnsetCurrentLateFee ensures that no value is present for CurrentLateFee, not even an explicit nil
### GetEscrowBalance

`func (o *MortgageLiability) GetEscrowBalance() float32`

GetEscrowBalance returns the EscrowBalance field if non-nil, zero value otherwise.

### GetEscrowBalanceOk

`func (o *MortgageLiability) GetEscrowBalanceOk() (*float32, bool)`

GetEscrowBalanceOk returns a tuple with the EscrowBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscrowBalance

`func (o *MortgageLiability) SetEscrowBalance(v float32)`

SetEscrowBalance sets EscrowBalance field to given value.


### SetEscrowBalanceNil

`func (o *MortgageLiability) SetEscrowBalanceNil(b bool)`

 SetEscrowBalanceNil sets the value for EscrowBalance to be an explicit nil

### UnsetEscrowBalance
`func (o *MortgageLiability) UnsetEscrowBalance()`

UnsetEscrowBalance ensures that no value is present for EscrowBalance, not even an explicit nil
### GetHasPmi

`func (o *MortgageLiability) GetHasPmi() bool`

GetHasPmi returns the HasPmi field if non-nil, zero value otherwise.

### GetHasPmiOk

`func (o *MortgageLiability) GetHasPmiOk() (*bool, bool)`

GetHasPmiOk returns a tuple with the HasPmi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPmi

`func (o *MortgageLiability) SetHasPmi(v bool)`

SetHasPmi sets HasPmi field to given value.


### SetHasPmiNil

`func (o *MortgageLiability) SetHasPmiNil(b bool)`

 SetHasPmiNil sets the value for HasPmi to be an explicit nil

### UnsetHasPmi
`func (o *MortgageLiability) UnsetHasPmi()`

UnsetHasPmi ensures that no value is present for HasPmi, not even an explicit nil
### GetHasPrepaymentPenalty

`func (o *MortgageLiability) GetHasPrepaymentPenalty() bool`

GetHasPrepaymentPenalty returns the HasPrepaymentPenalty field if non-nil, zero value otherwise.

### GetHasPrepaymentPenaltyOk

`func (o *MortgageLiability) GetHasPrepaymentPenaltyOk() (*bool, bool)`

GetHasPrepaymentPenaltyOk returns a tuple with the HasPrepaymentPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrepaymentPenalty

`func (o *MortgageLiability) SetHasPrepaymentPenalty(v bool)`

SetHasPrepaymentPenalty sets HasPrepaymentPenalty field to given value.


### SetHasPrepaymentPenaltyNil

`func (o *MortgageLiability) SetHasPrepaymentPenaltyNil(b bool)`

 SetHasPrepaymentPenaltyNil sets the value for HasPrepaymentPenalty to be an explicit nil

### UnsetHasPrepaymentPenalty
`func (o *MortgageLiability) UnsetHasPrepaymentPenalty()`

UnsetHasPrepaymentPenalty ensures that no value is present for HasPrepaymentPenalty, not even an explicit nil
### GetInterestRate

`func (o *MortgageLiability) GetInterestRate() MortgageInterestRate`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *MortgageLiability) GetInterestRateOk() (*MortgageInterestRate, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *MortgageLiability) SetInterestRate(v MortgageInterestRate)`

SetInterestRate sets InterestRate field to given value.


### GetLastPaymentAmount

`func (o *MortgageLiability) GetLastPaymentAmount() float32`

GetLastPaymentAmount returns the LastPaymentAmount field if non-nil, zero value otherwise.

### GetLastPaymentAmountOk

`func (o *MortgageLiability) GetLastPaymentAmountOk() (*float32, bool)`

GetLastPaymentAmountOk returns a tuple with the LastPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentAmount

`func (o *MortgageLiability) SetLastPaymentAmount(v float32)`

SetLastPaymentAmount sets LastPaymentAmount field to given value.


### SetLastPaymentAmountNil

`func (o *MortgageLiability) SetLastPaymentAmountNil(b bool)`

 SetLastPaymentAmountNil sets the value for LastPaymentAmount to be an explicit nil

### UnsetLastPaymentAmount
`func (o *MortgageLiability) UnsetLastPaymentAmount()`

UnsetLastPaymentAmount ensures that no value is present for LastPaymentAmount, not even an explicit nil
### GetLastPaymentDate

`func (o *MortgageLiability) GetLastPaymentDate() string`

GetLastPaymentDate returns the LastPaymentDate field if non-nil, zero value otherwise.

### GetLastPaymentDateOk

`func (o *MortgageLiability) GetLastPaymentDateOk() (*string, bool)`

GetLastPaymentDateOk returns a tuple with the LastPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentDate

`func (o *MortgageLiability) SetLastPaymentDate(v string)`

SetLastPaymentDate sets LastPaymentDate field to given value.


### SetLastPaymentDateNil

`func (o *MortgageLiability) SetLastPaymentDateNil(b bool)`

 SetLastPaymentDateNil sets the value for LastPaymentDate to be an explicit nil

### UnsetLastPaymentDate
`func (o *MortgageLiability) UnsetLastPaymentDate()`

UnsetLastPaymentDate ensures that no value is present for LastPaymentDate, not even an explicit nil
### GetLoanTypeDescription

`func (o *MortgageLiability) GetLoanTypeDescription() string`

GetLoanTypeDescription returns the LoanTypeDescription field if non-nil, zero value otherwise.

### GetLoanTypeDescriptionOk

`func (o *MortgageLiability) GetLoanTypeDescriptionOk() (*string, bool)`

GetLoanTypeDescriptionOk returns a tuple with the LoanTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanTypeDescription

`func (o *MortgageLiability) SetLoanTypeDescription(v string)`

SetLoanTypeDescription sets LoanTypeDescription field to given value.


### SetLoanTypeDescriptionNil

`func (o *MortgageLiability) SetLoanTypeDescriptionNil(b bool)`

 SetLoanTypeDescriptionNil sets the value for LoanTypeDescription to be an explicit nil

### UnsetLoanTypeDescription
`func (o *MortgageLiability) UnsetLoanTypeDescription()`

UnsetLoanTypeDescription ensures that no value is present for LoanTypeDescription, not even an explicit nil
### GetLoanTerm

`func (o *MortgageLiability) GetLoanTerm() string`

GetLoanTerm returns the LoanTerm field if non-nil, zero value otherwise.

### GetLoanTermOk

`func (o *MortgageLiability) GetLoanTermOk() (*string, bool)`

GetLoanTermOk returns a tuple with the LoanTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanTerm

`func (o *MortgageLiability) SetLoanTerm(v string)`

SetLoanTerm sets LoanTerm field to given value.


### SetLoanTermNil

`func (o *MortgageLiability) SetLoanTermNil(b bool)`

 SetLoanTermNil sets the value for LoanTerm to be an explicit nil

### UnsetLoanTerm
`func (o *MortgageLiability) UnsetLoanTerm()`

UnsetLoanTerm ensures that no value is present for LoanTerm, not even an explicit nil
### GetMaturityDate

`func (o *MortgageLiability) GetMaturityDate() string`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *MortgageLiability) GetMaturityDateOk() (*string, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *MortgageLiability) SetMaturityDate(v string)`

SetMaturityDate sets MaturityDate field to given value.


### SetMaturityDateNil

`func (o *MortgageLiability) SetMaturityDateNil(b bool)`

 SetMaturityDateNil sets the value for MaturityDate to be an explicit nil

### UnsetMaturityDate
`func (o *MortgageLiability) UnsetMaturityDate()`

UnsetMaturityDate ensures that no value is present for MaturityDate, not even an explicit nil
### GetNextMonthlyPayment

`func (o *MortgageLiability) GetNextMonthlyPayment() float32`

GetNextMonthlyPayment returns the NextMonthlyPayment field if non-nil, zero value otherwise.

### GetNextMonthlyPaymentOk

`func (o *MortgageLiability) GetNextMonthlyPaymentOk() (*float32, bool)`

GetNextMonthlyPaymentOk returns a tuple with the NextMonthlyPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextMonthlyPayment

`func (o *MortgageLiability) SetNextMonthlyPayment(v float32)`

SetNextMonthlyPayment sets NextMonthlyPayment field to given value.


### SetNextMonthlyPaymentNil

`func (o *MortgageLiability) SetNextMonthlyPaymentNil(b bool)`

 SetNextMonthlyPaymentNil sets the value for NextMonthlyPayment to be an explicit nil

### UnsetNextMonthlyPayment
`func (o *MortgageLiability) UnsetNextMonthlyPayment()`

UnsetNextMonthlyPayment ensures that no value is present for NextMonthlyPayment, not even an explicit nil
### GetNextPaymentDueDate

`func (o *MortgageLiability) GetNextPaymentDueDate() string`

GetNextPaymentDueDate returns the NextPaymentDueDate field if non-nil, zero value otherwise.

### GetNextPaymentDueDateOk

`func (o *MortgageLiability) GetNextPaymentDueDateOk() (*string, bool)`

GetNextPaymentDueDateOk returns a tuple with the NextPaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPaymentDueDate

`func (o *MortgageLiability) SetNextPaymentDueDate(v string)`

SetNextPaymentDueDate sets NextPaymentDueDate field to given value.


### SetNextPaymentDueDateNil

`func (o *MortgageLiability) SetNextPaymentDueDateNil(b bool)`

 SetNextPaymentDueDateNil sets the value for NextPaymentDueDate to be an explicit nil

### UnsetNextPaymentDueDate
`func (o *MortgageLiability) UnsetNextPaymentDueDate()`

UnsetNextPaymentDueDate ensures that no value is present for NextPaymentDueDate, not even an explicit nil
### GetOriginationDate

`func (o *MortgageLiability) GetOriginationDate() string`

GetOriginationDate returns the OriginationDate field if non-nil, zero value otherwise.

### GetOriginationDateOk

`func (o *MortgageLiability) GetOriginationDateOk() (*string, bool)`

GetOriginationDateOk returns a tuple with the OriginationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationDate

`func (o *MortgageLiability) SetOriginationDate(v string)`

SetOriginationDate sets OriginationDate field to given value.


### SetOriginationDateNil

`func (o *MortgageLiability) SetOriginationDateNil(b bool)`

 SetOriginationDateNil sets the value for OriginationDate to be an explicit nil

### UnsetOriginationDate
`func (o *MortgageLiability) UnsetOriginationDate()`

UnsetOriginationDate ensures that no value is present for OriginationDate, not even an explicit nil
### GetOriginationPrincipalAmount

`func (o *MortgageLiability) GetOriginationPrincipalAmount() float32`

GetOriginationPrincipalAmount returns the OriginationPrincipalAmount field if non-nil, zero value otherwise.

### GetOriginationPrincipalAmountOk

`func (o *MortgageLiability) GetOriginationPrincipalAmountOk() (*float32, bool)`

GetOriginationPrincipalAmountOk returns a tuple with the OriginationPrincipalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationPrincipalAmount

`func (o *MortgageLiability) SetOriginationPrincipalAmount(v float32)`

SetOriginationPrincipalAmount sets OriginationPrincipalAmount field to given value.


### SetOriginationPrincipalAmountNil

`func (o *MortgageLiability) SetOriginationPrincipalAmountNil(b bool)`

 SetOriginationPrincipalAmountNil sets the value for OriginationPrincipalAmount to be an explicit nil

### UnsetOriginationPrincipalAmount
`func (o *MortgageLiability) UnsetOriginationPrincipalAmount()`

UnsetOriginationPrincipalAmount ensures that no value is present for OriginationPrincipalAmount, not even an explicit nil
### GetPastDueAmount

`func (o *MortgageLiability) GetPastDueAmount() float32`

GetPastDueAmount returns the PastDueAmount field if non-nil, zero value otherwise.

### GetPastDueAmountOk

`func (o *MortgageLiability) GetPastDueAmountOk() (*float32, bool)`

GetPastDueAmountOk returns a tuple with the PastDueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastDueAmount

`func (o *MortgageLiability) SetPastDueAmount(v float32)`

SetPastDueAmount sets PastDueAmount field to given value.


### SetPastDueAmountNil

`func (o *MortgageLiability) SetPastDueAmountNil(b bool)`

 SetPastDueAmountNil sets the value for PastDueAmount to be an explicit nil

### UnsetPastDueAmount
`func (o *MortgageLiability) UnsetPastDueAmount()`

UnsetPastDueAmount ensures that no value is present for PastDueAmount, not even an explicit nil
### GetPropertyAddress

`func (o *MortgageLiability) GetPropertyAddress() MortgagePropertyAddress`

GetPropertyAddress returns the PropertyAddress field if non-nil, zero value otherwise.

### GetPropertyAddressOk

`func (o *MortgageLiability) GetPropertyAddressOk() (*MortgagePropertyAddress, bool)`

GetPropertyAddressOk returns a tuple with the PropertyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyAddress

`func (o *MortgageLiability) SetPropertyAddress(v MortgagePropertyAddress)`

SetPropertyAddress sets PropertyAddress field to given value.


### GetYtdInterestPaid

`func (o *MortgageLiability) GetYtdInterestPaid() float32`

GetYtdInterestPaid returns the YtdInterestPaid field if non-nil, zero value otherwise.

### GetYtdInterestPaidOk

`func (o *MortgageLiability) GetYtdInterestPaidOk() (*float32, bool)`

GetYtdInterestPaidOk returns a tuple with the YtdInterestPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdInterestPaid

`func (o *MortgageLiability) SetYtdInterestPaid(v float32)`

SetYtdInterestPaid sets YtdInterestPaid field to given value.


### SetYtdInterestPaidNil

`func (o *MortgageLiability) SetYtdInterestPaidNil(b bool)`

 SetYtdInterestPaidNil sets the value for YtdInterestPaid to be an explicit nil

### UnsetYtdInterestPaid
`func (o *MortgageLiability) UnsetYtdInterestPaid()`

UnsetYtdInterestPaid ensures that no value is present for YtdInterestPaid, not even an explicit nil
### GetYtdPrincipalPaid

`func (o *MortgageLiability) GetYtdPrincipalPaid() float32`

GetYtdPrincipalPaid returns the YtdPrincipalPaid field if non-nil, zero value otherwise.

### GetYtdPrincipalPaidOk

`func (o *MortgageLiability) GetYtdPrincipalPaidOk() (*float32, bool)`

GetYtdPrincipalPaidOk returns a tuple with the YtdPrincipalPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdPrincipalPaid

`func (o *MortgageLiability) SetYtdPrincipalPaid(v float32)`

SetYtdPrincipalPaid sets YtdPrincipalPaid field to given value.


### SetYtdPrincipalPaidNil

`func (o *MortgageLiability) SetYtdPrincipalPaidNil(b bool)`

 SetYtdPrincipalPaidNil sets the value for YtdPrincipalPaid to be an explicit nil

### UnsetYtdPrincipalPaid
`func (o *MortgageLiability) UnsetYtdPrincipalPaid()`

UnsetYtdPrincipalPaid ensures that no value is present for YtdPrincipalPaid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


