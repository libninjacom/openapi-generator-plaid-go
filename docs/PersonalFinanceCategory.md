# PersonalFinanceCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | **string** | A high level category that communicates the broad category of the transaction. | 
**Detailed** | **string** | Provides additional granularity to the primary categorization. | 

## Methods

### NewPersonalFinanceCategory

`func NewPersonalFinanceCategory(primary string, detailed string, ) *PersonalFinanceCategory`

NewPersonalFinanceCategory instantiates a new PersonalFinanceCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalFinanceCategoryWithDefaults

`func NewPersonalFinanceCategoryWithDefaults() *PersonalFinanceCategory`

NewPersonalFinanceCategoryWithDefaults instantiates a new PersonalFinanceCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *PersonalFinanceCategory) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *PersonalFinanceCategory) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *PersonalFinanceCategory) SetPrimary(v string)`

SetPrimary sets Primary field to given value.


### GetDetailed

`func (o *PersonalFinanceCategory) GetDetailed() string`

GetDetailed returns the Detailed field if non-nil, zero value otherwise.

### GetDetailedOk

`func (o *PersonalFinanceCategory) GetDetailedOk() (*string, bool)`

GetDetailedOk returns a tuple with the Detailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailed

`func (o *PersonalFinanceCategory) SetDetailed(v string)`

SetDetailed sets Detailed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


