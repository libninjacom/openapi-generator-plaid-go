# Deductions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subtotals** | Pointer to [**[]Total**](Total.md) |  | [optional] 
**Breakdown** | [**[]DeductionsBreakdown**](DeductionsBreakdown.md) |  | 
**Totals** | Pointer to [**[]Total**](Total.md) |  | [optional] 
**Total** | [**DeductionsTotal**](DeductionsTotal.md) |  | 

## Methods

### NewDeductions

`func NewDeductions(breakdown []DeductionsBreakdown, total DeductionsTotal, ) *Deductions`

NewDeductions instantiates a new Deductions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeductionsWithDefaults

`func NewDeductionsWithDefaults() *Deductions`

NewDeductionsWithDefaults instantiates a new Deductions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtotals

`func (o *Deductions) GetSubtotals() []Total`

GetSubtotals returns the Subtotals field if non-nil, zero value otherwise.

### GetSubtotalsOk

`func (o *Deductions) GetSubtotalsOk() (*[]Total, bool)`

GetSubtotalsOk returns a tuple with the Subtotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotals

`func (o *Deductions) SetSubtotals(v []Total)`

SetSubtotals sets Subtotals field to given value.

### HasSubtotals

`func (o *Deductions) HasSubtotals() bool`

HasSubtotals returns a boolean if a field has been set.

### GetBreakdown

`func (o *Deductions) GetBreakdown() []DeductionsBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *Deductions) GetBreakdownOk() (*[]DeductionsBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *Deductions) SetBreakdown(v []DeductionsBreakdown)`

SetBreakdown sets Breakdown field to given value.


### GetTotals

`func (o *Deductions) GetTotals() []Total`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *Deductions) GetTotalsOk() (*[]Total, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *Deductions) SetTotals(v []Total)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *Deductions) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetTotal

`func (o *Deductions) GetTotal() DeductionsTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Deductions) GetTotalOk() (*DeductionsTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Deductions) SetTotal(v DeductionsTotal)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


