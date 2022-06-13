# IncomeBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**NullableIncomeBreakdownType**](IncomeBreakdownType.md) |  | 
**Rate** | **NullableFloat32** | The hourly rate at which the income is paid. | 
**Hours** | **NullableFloat32** | The number of hours logged for this income for this pay period. | 
**Total** | **NullableFloat32** | The total pay for this pay period. | 

## Methods

### NewIncomeBreakdown

`func NewIncomeBreakdown(type_ NullableIncomeBreakdownType, rate NullableFloat32, hours NullableFloat32, total NullableFloat32, ) *IncomeBreakdown`

NewIncomeBreakdown instantiates a new IncomeBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeBreakdownWithDefaults

`func NewIncomeBreakdownWithDefaults() *IncomeBreakdown`

NewIncomeBreakdownWithDefaults instantiates a new IncomeBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IncomeBreakdown) GetType() IncomeBreakdownType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncomeBreakdown) GetTypeOk() (*IncomeBreakdownType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncomeBreakdown) SetType(v IncomeBreakdownType)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *IncomeBreakdown) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *IncomeBreakdown) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetRate

`func (o *IncomeBreakdown) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *IncomeBreakdown) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *IncomeBreakdown) SetRate(v float32)`

SetRate sets Rate field to given value.


### SetRateNil

`func (o *IncomeBreakdown) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *IncomeBreakdown) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetHours

`func (o *IncomeBreakdown) GetHours() float32`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *IncomeBreakdown) GetHoursOk() (*float32, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *IncomeBreakdown) SetHours(v float32)`

SetHours sets Hours field to given value.


### SetHoursNil

`func (o *IncomeBreakdown) SetHoursNil(b bool)`

 SetHoursNil sets the value for Hours to be an explicit nil

### UnsetHours
`func (o *IncomeBreakdown) UnsetHours()`

UnsetHours ensures that no value is present for Hours, not even an explicit nil
### GetTotal

`func (o *IncomeBreakdown) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IncomeBreakdown) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IncomeBreakdown) SetTotal(v float32)`

SetTotal sets Total field to given value.


### SetTotalNil

`func (o *IncomeBreakdown) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *IncomeBreakdown) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


