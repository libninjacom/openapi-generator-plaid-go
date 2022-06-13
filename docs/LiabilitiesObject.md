# LiabilitiesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credit** | [**[]CreditCardLiability**](CreditCardLiability.md) | The credit accounts returned. | 
**Mortgage** | [**[]MortgageLiability**](MortgageLiability.md) | The mortgage accounts returned. | 
**Student** | [**[]StudentLoan**](StudentLoan.md) | The student loan accounts returned. | 

## Methods

### NewLiabilitiesObject

`func NewLiabilitiesObject(credit []CreditCardLiability, mortgage []MortgageLiability, student []StudentLoan, ) *LiabilitiesObject`

NewLiabilitiesObject instantiates a new LiabilitiesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiabilitiesObjectWithDefaults

`func NewLiabilitiesObjectWithDefaults() *LiabilitiesObject`

NewLiabilitiesObjectWithDefaults instantiates a new LiabilitiesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredit

`func (o *LiabilitiesObject) GetCredit() []CreditCardLiability`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *LiabilitiesObject) GetCreditOk() (*[]CreditCardLiability, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *LiabilitiesObject) SetCredit(v []CreditCardLiability)`

SetCredit sets Credit field to given value.


### SetCreditNil

`func (o *LiabilitiesObject) SetCreditNil(b bool)`

 SetCreditNil sets the value for Credit to be an explicit nil

### UnsetCredit
`func (o *LiabilitiesObject) UnsetCredit()`

UnsetCredit ensures that no value is present for Credit, not even an explicit nil
### GetMortgage

`func (o *LiabilitiesObject) GetMortgage() []MortgageLiability`

GetMortgage returns the Mortgage field if non-nil, zero value otherwise.

### GetMortgageOk

`func (o *LiabilitiesObject) GetMortgageOk() (*[]MortgageLiability, bool)`

GetMortgageOk returns a tuple with the Mortgage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMortgage

`func (o *LiabilitiesObject) SetMortgage(v []MortgageLiability)`

SetMortgage sets Mortgage field to given value.


### SetMortgageNil

`func (o *LiabilitiesObject) SetMortgageNil(b bool)`

 SetMortgageNil sets the value for Mortgage to be an explicit nil

### UnsetMortgage
`func (o *LiabilitiesObject) UnsetMortgage()`

UnsetMortgage ensures that no value is present for Mortgage, not even an explicit nil
### GetStudent

`func (o *LiabilitiesObject) GetStudent() []StudentLoan`

GetStudent returns the Student field if non-nil, zero value otherwise.

### GetStudentOk

`func (o *LiabilitiesObject) GetStudentOk() (*[]StudentLoan, bool)`

GetStudentOk returns a tuple with the Student field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudent

`func (o *LiabilitiesObject) SetStudent(v []StudentLoan)`

SetStudent sets Student field to given value.


### SetStudentNil

`func (o *LiabilitiesObject) SetStudentNil(b bool)`

 SetStudentNil sets the value for Student to be an explicit nil

### UnsetStudent
`func (o *LiabilitiesObject) UnsetStudent()`

UnsetStudent ensures that no value is present for Student, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


