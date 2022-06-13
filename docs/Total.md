# Total

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanonicalDescription** | Pointer to [**NullableTotalCanonicalDescription**](TotalCanonicalDescription.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Text of the line item as printed on the paystub. | [optional] 
**CurrentPay** | Pointer to [**Pay**](Pay.md) |  | [optional] 
**YtdPay** | Pointer to [**Pay**](Pay.md) |  | [optional] 

## Methods

### NewTotal

`func NewTotal() *Total`

NewTotal instantiates a new Total object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTotalWithDefaults

`func NewTotalWithDefaults() *Total`

NewTotalWithDefaults instantiates a new Total object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanonicalDescription

`func (o *Total) GetCanonicalDescription() TotalCanonicalDescription`

GetCanonicalDescription returns the CanonicalDescription field if non-nil, zero value otherwise.

### GetCanonicalDescriptionOk

`func (o *Total) GetCanonicalDescriptionOk() (*TotalCanonicalDescription, bool)`

GetCanonicalDescriptionOk returns a tuple with the CanonicalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalDescription

`func (o *Total) SetCanonicalDescription(v TotalCanonicalDescription)`

SetCanonicalDescription sets CanonicalDescription field to given value.

### HasCanonicalDescription

`func (o *Total) HasCanonicalDescription() bool`

HasCanonicalDescription returns a boolean if a field has been set.

### SetCanonicalDescriptionNil

`func (o *Total) SetCanonicalDescriptionNil(b bool)`

 SetCanonicalDescriptionNil sets the value for CanonicalDescription to be an explicit nil

### UnsetCanonicalDescription
`func (o *Total) UnsetCanonicalDescription()`

UnsetCanonicalDescription ensures that no value is present for CanonicalDescription, not even an explicit nil
### GetDescription

`func (o *Total) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Total) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Total) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Total) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Total) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Total) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCurrentPay

`func (o *Total) GetCurrentPay() Pay`

GetCurrentPay returns the CurrentPay field if non-nil, zero value otherwise.

### GetCurrentPayOk

`func (o *Total) GetCurrentPayOk() (*Pay, bool)`

GetCurrentPayOk returns a tuple with the CurrentPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPay

`func (o *Total) SetCurrentPay(v Pay)`

SetCurrentPay sets CurrentPay field to given value.

### HasCurrentPay

`func (o *Total) HasCurrentPay() bool`

HasCurrentPay returns a boolean if a field has been set.

### GetYtdPay

`func (o *Total) GetYtdPay() Pay`

GetYtdPay returns the YtdPay field if non-nil, zero value otherwise.

### GetYtdPayOk

`func (o *Total) GetYtdPayOk() (*Pay, bool)`

GetYtdPayOk returns a tuple with the YtdPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdPay

`func (o *Total) SetYtdPay(v Pay)`

SetYtdPay sets YtdPay field to given value.

### HasYtdPay

`func (o *Total) HasYtdPay() bool`

HasYtdPay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


