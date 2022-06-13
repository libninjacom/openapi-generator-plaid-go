# IncomeOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paystubs** | Pointer to [**[]PaystubOverride**](PaystubOverride.md) | A list of paystubs associated with the account. | [optional] 

## Methods

### NewIncomeOverride

`func NewIncomeOverride() *IncomeOverride`

NewIncomeOverride instantiates a new IncomeOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeOverrideWithDefaults

`func NewIncomeOverrideWithDefaults() *IncomeOverride`

NewIncomeOverrideWithDefaults instantiates a new IncomeOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaystubs

`func (o *IncomeOverride) GetPaystubs() []PaystubOverride`

GetPaystubs returns the Paystubs field if non-nil, zero value otherwise.

### GetPaystubsOk

`func (o *IncomeOverride) GetPaystubsOk() (*[]PaystubOverride, bool)`

GetPaystubsOk returns a tuple with the Paystubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaystubs

`func (o *IncomeOverride) SetPaystubs(v []PaystubOverride)`

SetPaystubs sets Paystubs field to given value.

### HasPaystubs

`func (o *IncomeOverride) HasPaystubs() bool`

HasPaystubs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


