# W2StateAndLocalWages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **NullableString** | State associated with the wage. | [optional] 
**EmployerStateIdNumber** | Pointer to **NullableString** | State identification number of the employer. | [optional] 
**StateWagesTips** | Pointer to **NullableString** | Wages and tips from the specified state. | [optional] 
**StateIncomeTax** | Pointer to **NullableString** | Income tax from the specified state. | [optional] 
**LocalWagesTips** | Pointer to **NullableString** | Wages and tips from the locality. | [optional] 
**LocalIncomeTax** | Pointer to **NullableString** | Income tax from the locality. | [optional] 
**LocalityName** | Pointer to **NullableString** | Name of the locality. | [optional] 

## Methods

### NewW2StateAndLocalWages

`func NewW2StateAndLocalWages() *W2StateAndLocalWages`

NewW2StateAndLocalWages instantiates a new W2StateAndLocalWages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewW2StateAndLocalWagesWithDefaults

`func NewW2StateAndLocalWagesWithDefaults() *W2StateAndLocalWages`

NewW2StateAndLocalWagesWithDefaults instantiates a new W2StateAndLocalWages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *W2StateAndLocalWages) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *W2StateAndLocalWages) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *W2StateAndLocalWages) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *W2StateAndLocalWages) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *W2StateAndLocalWages) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *W2StateAndLocalWages) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetEmployerStateIdNumber

`func (o *W2StateAndLocalWages) GetEmployerStateIdNumber() string`

GetEmployerStateIdNumber returns the EmployerStateIdNumber field if non-nil, zero value otherwise.

### GetEmployerStateIdNumberOk

`func (o *W2StateAndLocalWages) GetEmployerStateIdNumberOk() (*string, bool)`

GetEmployerStateIdNumberOk returns a tuple with the EmployerStateIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerStateIdNumber

`func (o *W2StateAndLocalWages) SetEmployerStateIdNumber(v string)`

SetEmployerStateIdNumber sets EmployerStateIdNumber field to given value.

### HasEmployerStateIdNumber

`func (o *W2StateAndLocalWages) HasEmployerStateIdNumber() bool`

HasEmployerStateIdNumber returns a boolean if a field has been set.

### SetEmployerStateIdNumberNil

`func (o *W2StateAndLocalWages) SetEmployerStateIdNumberNil(b bool)`

 SetEmployerStateIdNumberNil sets the value for EmployerStateIdNumber to be an explicit nil

### UnsetEmployerStateIdNumber
`func (o *W2StateAndLocalWages) UnsetEmployerStateIdNumber()`

UnsetEmployerStateIdNumber ensures that no value is present for EmployerStateIdNumber, not even an explicit nil
### GetStateWagesTips

`func (o *W2StateAndLocalWages) GetStateWagesTips() string`

GetStateWagesTips returns the StateWagesTips field if non-nil, zero value otherwise.

### GetStateWagesTipsOk

`func (o *W2StateAndLocalWages) GetStateWagesTipsOk() (*string, bool)`

GetStateWagesTipsOk returns a tuple with the StateWagesTips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateWagesTips

`func (o *W2StateAndLocalWages) SetStateWagesTips(v string)`

SetStateWagesTips sets StateWagesTips field to given value.

### HasStateWagesTips

`func (o *W2StateAndLocalWages) HasStateWagesTips() bool`

HasStateWagesTips returns a boolean if a field has been set.

### SetStateWagesTipsNil

`func (o *W2StateAndLocalWages) SetStateWagesTipsNil(b bool)`

 SetStateWagesTipsNil sets the value for StateWagesTips to be an explicit nil

### UnsetStateWagesTips
`func (o *W2StateAndLocalWages) UnsetStateWagesTips()`

UnsetStateWagesTips ensures that no value is present for StateWagesTips, not even an explicit nil
### GetStateIncomeTax

`func (o *W2StateAndLocalWages) GetStateIncomeTax() string`

GetStateIncomeTax returns the StateIncomeTax field if non-nil, zero value otherwise.

### GetStateIncomeTaxOk

`func (o *W2StateAndLocalWages) GetStateIncomeTaxOk() (*string, bool)`

GetStateIncomeTaxOk returns a tuple with the StateIncomeTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateIncomeTax

`func (o *W2StateAndLocalWages) SetStateIncomeTax(v string)`

SetStateIncomeTax sets StateIncomeTax field to given value.

### HasStateIncomeTax

`func (o *W2StateAndLocalWages) HasStateIncomeTax() bool`

HasStateIncomeTax returns a boolean if a field has been set.

### SetStateIncomeTaxNil

`func (o *W2StateAndLocalWages) SetStateIncomeTaxNil(b bool)`

 SetStateIncomeTaxNil sets the value for StateIncomeTax to be an explicit nil

### UnsetStateIncomeTax
`func (o *W2StateAndLocalWages) UnsetStateIncomeTax()`

UnsetStateIncomeTax ensures that no value is present for StateIncomeTax, not even an explicit nil
### GetLocalWagesTips

`func (o *W2StateAndLocalWages) GetLocalWagesTips() string`

GetLocalWagesTips returns the LocalWagesTips field if non-nil, zero value otherwise.

### GetLocalWagesTipsOk

`func (o *W2StateAndLocalWages) GetLocalWagesTipsOk() (*string, bool)`

GetLocalWagesTipsOk returns a tuple with the LocalWagesTips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalWagesTips

`func (o *W2StateAndLocalWages) SetLocalWagesTips(v string)`

SetLocalWagesTips sets LocalWagesTips field to given value.

### HasLocalWagesTips

`func (o *W2StateAndLocalWages) HasLocalWagesTips() bool`

HasLocalWagesTips returns a boolean if a field has been set.

### SetLocalWagesTipsNil

`func (o *W2StateAndLocalWages) SetLocalWagesTipsNil(b bool)`

 SetLocalWagesTipsNil sets the value for LocalWagesTips to be an explicit nil

### UnsetLocalWagesTips
`func (o *W2StateAndLocalWages) UnsetLocalWagesTips()`

UnsetLocalWagesTips ensures that no value is present for LocalWagesTips, not even an explicit nil
### GetLocalIncomeTax

`func (o *W2StateAndLocalWages) GetLocalIncomeTax() string`

GetLocalIncomeTax returns the LocalIncomeTax field if non-nil, zero value otherwise.

### GetLocalIncomeTaxOk

`func (o *W2StateAndLocalWages) GetLocalIncomeTaxOk() (*string, bool)`

GetLocalIncomeTaxOk returns a tuple with the LocalIncomeTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIncomeTax

`func (o *W2StateAndLocalWages) SetLocalIncomeTax(v string)`

SetLocalIncomeTax sets LocalIncomeTax field to given value.

### HasLocalIncomeTax

`func (o *W2StateAndLocalWages) HasLocalIncomeTax() bool`

HasLocalIncomeTax returns a boolean if a field has been set.

### SetLocalIncomeTaxNil

`func (o *W2StateAndLocalWages) SetLocalIncomeTaxNil(b bool)`

 SetLocalIncomeTaxNil sets the value for LocalIncomeTax to be an explicit nil

### UnsetLocalIncomeTax
`func (o *W2StateAndLocalWages) UnsetLocalIncomeTax()`

UnsetLocalIncomeTax ensures that no value is present for LocalIncomeTax, not even an explicit nil
### GetLocalityName

`func (o *W2StateAndLocalWages) GetLocalityName() string`

GetLocalityName returns the LocalityName field if non-nil, zero value otherwise.

### GetLocalityNameOk

`func (o *W2StateAndLocalWages) GetLocalityNameOk() (*string, bool)`

GetLocalityNameOk returns a tuple with the LocalityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalityName

`func (o *W2StateAndLocalWages) SetLocalityName(v string)`

SetLocalityName sets LocalityName field to given value.

### HasLocalityName

`func (o *W2StateAndLocalWages) HasLocalityName() bool`

HasLocalityName returns a boolean if a field has been set.

### SetLocalityNameNil

`func (o *W2StateAndLocalWages) SetLocalityNameNil(b bool)`

 SetLocalityNameNil sets the value for LocalityName to be an explicit nil

### UnsetLocalityName
`func (o *W2StateAndLocalWages) UnsetLocalityName()`

UnsetLocalityName ensures that no value is present for LocalityName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


