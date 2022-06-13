# Paystub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deductions** | [**Deductions**](Deductions.md) |  | 
**DocId** | **string** | An identifier of the document referenced by the document metadata. | 
**Earnings** | [**Earnings**](Earnings.md) |  | 
**Employee** | [**Employee**](Employee.md) |  | 
**Employer** | [**PaystubEmployer**](PaystubEmployer.md) |  | 
**EmploymentDetails** | Pointer to [**EmploymentDetails**](EmploymentDetails.md) |  | [optional] 
**NetPay** | [**NetPay**](NetPay.md) |  | 
**PayPeriodDetails** | [**PayPeriodDetails**](PayPeriodDetails.md) |  | 
**PaystubDetails** | Pointer to [**PaystubDetails**](PaystubDetails.md) |  | [optional] 
**IncomeBreakdown** | Pointer to [**[]IncomeBreakdown**](IncomeBreakdown.md) |  | [optional] 
**YtdEarnings** | Pointer to [**PaystubYTDDetails**](PaystubYTDDetails.md) |  | [optional] 
**Verification** | [**NullablePaystubVerification**](PaystubVerification.md) |  | 

## Methods

### NewPaystub

`func NewPaystub(deductions Deductions, docId string, earnings Earnings, employee Employee, employer PaystubEmployer, netPay NetPay, payPeriodDetails PayPeriodDetails, verification NullablePaystubVerification, ) *Paystub`

NewPaystub instantiates a new Paystub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaystubWithDefaults

`func NewPaystubWithDefaults() *Paystub`

NewPaystubWithDefaults instantiates a new Paystub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeductions

`func (o *Paystub) GetDeductions() Deductions`

GetDeductions returns the Deductions field if non-nil, zero value otherwise.

### GetDeductionsOk

`func (o *Paystub) GetDeductionsOk() (*Deductions, bool)`

GetDeductionsOk returns a tuple with the Deductions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductions

`func (o *Paystub) SetDeductions(v Deductions)`

SetDeductions sets Deductions field to given value.


### GetDocId

`func (o *Paystub) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *Paystub) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *Paystub) SetDocId(v string)`

SetDocId sets DocId field to given value.


### GetEarnings

`func (o *Paystub) GetEarnings() Earnings`

GetEarnings returns the Earnings field if non-nil, zero value otherwise.

### GetEarningsOk

`func (o *Paystub) GetEarningsOk() (*Earnings, bool)`

GetEarningsOk returns a tuple with the Earnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnings

`func (o *Paystub) SetEarnings(v Earnings)`

SetEarnings sets Earnings field to given value.


### GetEmployee

`func (o *Paystub) GetEmployee() Employee`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *Paystub) GetEmployeeOk() (*Employee, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *Paystub) SetEmployee(v Employee)`

SetEmployee sets Employee field to given value.


### GetEmployer

`func (o *Paystub) GetEmployer() PaystubEmployer`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *Paystub) GetEmployerOk() (*PaystubEmployer, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *Paystub) SetEmployer(v PaystubEmployer)`

SetEmployer sets Employer field to given value.


### GetEmploymentDetails

`func (o *Paystub) GetEmploymentDetails() EmploymentDetails`

GetEmploymentDetails returns the EmploymentDetails field if non-nil, zero value otherwise.

### GetEmploymentDetailsOk

`func (o *Paystub) GetEmploymentDetailsOk() (*EmploymentDetails, bool)`

GetEmploymentDetailsOk returns a tuple with the EmploymentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentDetails

`func (o *Paystub) SetEmploymentDetails(v EmploymentDetails)`

SetEmploymentDetails sets EmploymentDetails field to given value.

### HasEmploymentDetails

`func (o *Paystub) HasEmploymentDetails() bool`

HasEmploymentDetails returns a boolean if a field has been set.

### GetNetPay

`func (o *Paystub) GetNetPay() NetPay`

GetNetPay returns the NetPay field if non-nil, zero value otherwise.

### GetNetPayOk

`func (o *Paystub) GetNetPayOk() (*NetPay, bool)`

GetNetPayOk returns a tuple with the NetPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPay

`func (o *Paystub) SetNetPay(v NetPay)`

SetNetPay sets NetPay field to given value.


### GetPayPeriodDetails

`func (o *Paystub) GetPayPeriodDetails() PayPeriodDetails`

GetPayPeriodDetails returns the PayPeriodDetails field if non-nil, zero value otherwise.

### GetPayPeriodDetailsOk

`func (o *Paystub) GetPayPeriodDetailsOk() (*PayPeriodDetails, bool)`

GetPayPeriodDetailsOk returns a tuple with the PayPeriodDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriodDetails

`func (o *Paystub) SetPayPeriodDetails(v PayPeriodDetails)`

SetPayPeriodDetails sets PayPeriodDetails field to given value.


### GetPaystubDetails

`func (o *Paystub) GetPaystubDetails() PaystubDetails`

GetPaystubDetails returns the PaystubDetails field if non-nil, zero value otherwise.

### GetPaystubDetailsOk

`func (o *Paystub) GetPaystubDetailsOk() (*PaystubDetails, bool)`

GetPaystubDetailsOk returns a tuple with the PaystubDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaystubDetails

`func (o *Paystub) SetPaystubDetails(v PaystubDetails)`

SetPaystubDetails sets PaystubDetails field to given value.

### HasPaystubDetails

`func (o *Paystub) HasPaystubDetails() bool`

HasPaystubDetails returns a boolean if a field has been set.

### GetIncomeBreakdown

`func (o *Paystub) GetIncomeBreakdown() []IncomeBreakdown`

GetIncomeBreakdown returns the IncomeBreakdown field if non-nil, zero value otherwise.

### GetIncomeBreakdownOk

`func (o *Paystub) GetIncomeBreakdownOk() (*[]IncomeBreakdown, bool)`

GetIncomeBreakdownOk returns a tuple with the IncomeBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeBreakdown

`func (o *Paystub) SetIncomeBreakdown(v []IncomeBreakdown)`

SetIncomeBreakdown sets IncomeBreakdown field to given value.

### HasIncomeBreakdown

`func (o *Paystub) HasIncomeBreakdown() bool`

HasIncomeBreakdown returns a boolean if a field has been set.

### GetYtdEarnings

`func (o *Paystub) GetYtdEarnings() PaystubYTDDetails`

GetYtdEarnings returns the YtdEarnings field if non-nil, zero value otherwise.

### GetYtdEarningsOk

`func (o *Paystub) GetYtdEarningsOk() (*PaystubYTDDetails, bool)`

GetYtdEarningsOk returns a tuple with the YtdEarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdEarnings

`func (o *Paystub) SetYtdEarnings(v PaystubYTDDetails)`

SetYtdEarnings sets YtdEarnings field to given value.

### HasYtdEarnings

`func (o *Paystub) HasYtdEarnings() bool`

HasYtdEarnings returns a boolean if a field has been set.

### GetVerification

`func (o *Paystub) GetVerification() PaystubVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *Paystub) GetVerificationOk() (*PaystubVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *Paystub) SetVerification(v PaystubVerification)`

SetVerification sets Verification field to given value.


### SetVerificationNil

`func (o *Paystub) SetVerificationNil(b bool)`

 SetVerificationNil sets the value for Verification to be an explicit nil

### UnsetVerification
`func (o *Paystub) UnsetVerification()`

UnsetVerification ensures that no value is present for Verification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


