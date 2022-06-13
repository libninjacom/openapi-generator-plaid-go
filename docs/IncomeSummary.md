# IncomeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployerName** | [**EmployerIncomeSummaryFieldString**](EmployerIncomeSummaryFieldString.md) |  | 
**EmployeeName** | [**EmployeeIncomeSummaryFieldString**](EmployeeIncomeSummaryFieldString.md) |  | 
**YtdGrossIncome** | [**YTDGrossIncomeSummaryFieldNumber**](YTDGrossIncomeSummaryFieldNumber.md) |  | 
**YtdNetIncome** | [**YTDNetIncomeSummaryFieldNumber**](YTDNetIncomeSummaryFieldNumber.md) |  | 
**PayFrequency** | [**NullablePayFrequency**](PayFrequency.md) |  | 
**ProjectedWage** | [**ProjectedIncomeSummaryFieldNumber**](ProjectedIncomeSummaryFieldNumber.md) |  | 
**VerifiedTransaction** | [**NullableTransactionData**](TransactionData.md) |  | 

## Methods

### NewIncomeSummary

`func NewIncomeSummary(employerName EmployerIncomeSummaryFieldString, employeeName EmployeeIncomeSummaryFieldString, ytdGrossIncome YTDGrossIncomeSummaryFieldNumber, ytdNetIncome YTDNetIncomeSummaryFieldNumber, payFrequency NullablePayFrequency, projectedWage ProjectedIncomeSummaryFieldNumber, verifiedTransaction NullableTransactionData, ) *IncomeSummary`

NewIncomeSummary instantiates a new IncomeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeSummaryWithDefaults

`func NewIncomeSummaryWithDefaults() *IncomeSummary`

NewIncomeSummaryWithDefaults instantiates a new IncomeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployerName

`func (o *IncomeSummary) GetEmployerName() EmployerIncomeSummaryFieldString`

GetEmployerName returns the EmployerName field if non-nil, zero value otherwise.

### GetEmployerNameOk

`func (o *IncomeSummary) GetEmployerNameOk() (*EmployerIncomeSummaryFieldString, bool)`

GetEmployerNameOk returns a tuple with the EmployerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerName

`func (o *IncomeSummary) SetEmployerName(v EmployerIncomeSummaryFieldString)`

SetEmployerName sets EmployerName field to given value.


### GetEmployeeName

`func (o *IncomeSummary) GetEmployeeName() EmployeeIncomeSummaryFieldString`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *IncomeSummary) GetEmployeeNameOk() (*EmployeeIncomeSummaryFieldString, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *IncomeSummary) SetEmployeeName(v EmployeeIncomeSummaryFieldString)`

SetEmployeeName sets EmployeeName field to given value.


### GetYtdGrossIncome

`func (o *IncomeSummary) GetYtdGrossIncome() YTDGrossIncomeSummaryFieldNumber`

GetYtdGrossIncome returns the YtdGrossIncome field if non-nil, zero value otherwise.

### GetYtdGrossIncomeOk

`func (o *IncomeSummary) GetYtdGrossIncomeOk() (*YTDGrossIncomeSummaryFieldNumber, bool)`

GetYtdGrossIncomeOk returns a tuple with the YtdGrossIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdGrossIncome

`func (o *IncomeSummary) SetYtdGrossIncome(v YTDGrossIncomeSummaryFieldNumber)`

SetYtdGrossIncome sets YtdGrossIncome field to given value.


### GetYtdNetIncome

`func (o *IncomeSummary) GetYtdNetIncome() YTDNetIncomeSummaryFieldNumber`

GetYtdNetIncome returns the YtdNetIncome field if non-nil, zero value otherwise.

### GetYtdNetIncomeOk

`func (o *IncomeSummary) GetYtdNetIncomeOk() (*YTDNetIncomeSummaryFieldNumber, bool)`

GetYtdNetIncomeOk returns a tuple with the YtdNetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdNetIncome

`func (o *IncomeSummary) SetYtdNetIncome(v YTDNetIncomeSummaryFieldNumber)`

SetYtdNetIncome sets YtdNetIncome field to given value.


### GetPayFrequency

`func (o *IncomeSummary) GetPayFrequency() PayFrequency`

GetPayFrequency returns the PayFrequency field if non-nil, zero value otherwise.

### GetPayFrequencyOk

`func (o *IncomeSummary) GetPayFrequencyOk() (*PayFrequency, bool)`

GetPayFrequencyOk returns a tuple with the PayFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayFrequency

`func (o *IncomeSummary) SetPayFrequency(v PayFrequency)`

SetPayFrequency sets PayFrequency field to given value.


### SetPayFrequencyNil

`func (o *IncomeSummary) SetPayFrequencyNil(b bool)`

 SetPayFrequencyNil sets the value for PayFrequency to be an explicit nil

### UnsetPayFrequency
`func (o *IncomeSummary) UnsetPayFrequency()`

UnsetPayFrequency ensures that no value is present for PayFrequency, not even an explicit nil
### GetProjectedWage

`func (o *IncomeSummary) GetProjectedWage() ProjectedIncomeSummaryFieldNumber`

GetProjectedWage returns the ProjectedWage field if non-nil, zero value otherwise.

### GetProjectedWageOk

`func (o *IncomeSummary) GetProjectedWageOk() (*ProjectedIncomeSummaryFieldNumber, bool)`

GetProjectedWageOk returns a tuple with the ProjectedWage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedWage

`func (o *IncomeSummary) SetProjectedWage(v ProjectedIncomeSummaryFieldNumber)`

SetProjectedWage sets ProjectedWage field to given value.


### GetVerifiedTransaction

`func (o *IncomeSummary) GetVerifiedTransaction() TransactionData`

GetVerifiedTransaction returns the VerifiedTransaction field if non-nil, zero value otherwise.

### GetVerifiedTransactionOk

`func (o *IncomeSummary) GetVerifiedTransactionOk() (*TransactionData, bool)`

GetVerifiedTransactionOk returns a tuple with the VerifiedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedTransaction

`func (o *IncomeSummary) SetVerifiedTransaction(v TransactionData)`

SetVerifiedTransaction sets VerifiedTransaction field to given value.


### SetVerifiedTransactionNil

`func (o *IncomeSummary) SetVerifiedTransactionNil(b bool)`

 SetVerifiedTransactionNil sets the value for VerifiedTransaction to be an explicit nil

### UnsetVerifiedTransaction
`func (o *IncomeSummary) UnsetVerifiedTransaction()`

UnsetVerifiedTransaction ensures that no value is present for VerifiedTransaction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


