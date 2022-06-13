# Earnings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subtotals** | Pointer to [**[]EarningsTotal**](EarningsTotal.md) |  | [optional] 
**Totals** | Pointer to [**[]EarningsTotal**](EarningsTotal.md) |  | [optional] 
**Breakdown** | Pointer to [**[]EarningsBreakdown**](EarningsBreakdown.md) |  | [optional] 
**Total** | Pointer to [**EarningsTotal**](EarningsTotal.md) |  | [optional] 

## Methods

### NewEarnings

`func NewEarnings() *Earnings`

NewEarnings instantiates a new Earnings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarningsWithDefaults

`func NewEarningsWithDefaults() *Earnings`

NewEarningsWithDefaults instantiates a new Earnings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtotals

`func (o *Earnings) GetSubtotals() []EarningsTotal`

GetSubtotals returns the Subtotals field if non-nil, zero value otherwise.

### GetSubtotalsOk

`func (o *Earnings) GetSubtotalsOk() (*[]EarningsTotal, bool)`

GetSubtotalsOk returns a tuple with the Subtotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotals

`func (o *Earnings) SetSubtotals(v []EarningsTotal)`

SetSubtotals sets Subtotals field to given value.

### HasSubtotals

`func (o *Earnings) HasSubtotals() bool`

HasSubtotals returns a boolean if a field has been set.

### GetTotals

`func (o *Earnings) GetTotals() []EarningsTotal`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *Earnings) GetTotalsOk() (*[]EarningsTotal, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *Earnings) SetTotals(v []EarningsTotal)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *Earnings) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetBreakdown

`func (o *Earnings) GetBreakdown() []EarningsBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *Earnings) GetBreakdownOk() (*[]EarningsBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *Earnings) SetBreakdown(v []EarningsBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *Earnings) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetTotal

`func (o *Earnings) GetTotal() EarningsTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Earnings) GetTotalOk() (*EarningsTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Earnings) SetTotal(v EarningsTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Earnings) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


