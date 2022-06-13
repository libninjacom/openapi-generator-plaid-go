# IncomeVerificationPrecheckMilitaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActiveDuty** | Pointer to **NullableBool** | Is the user currently active duty in the US military | [optional] 
**Branch** | Pointer to **NullableString** | If the user is currently serving in the US military, the branch of the military they are serving in | [optional] 

## Methods

### NewIncomeVerificationPrecheckMilitaryInfo

`func NewIncomeVerificationPrecheckMilitaryInfo() *IncomeVerificationPrecheckMilitaryInfo`

NewIncomeVerificationPrecheckMilitaryInfo instantiates a new IncomeVerificationPrecheckMilitaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationPrecheckMilitaryInfoWithDefaults

`func NewIncomeVerificationPrecheckMilitaryInfoWithDefaults() *IncomeVerificationPrecheckMilitaryInfo`

NewIncomeVerificationPrecheckMilitaryInfoWithDefaults instantiates a new IncomeVerificationPrecheckMilitaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActiveDuty

`func (o *IncomeVerificationPrecheckMilitaryInfo) GetIsActiveDuty() bool`

GetIsActiveDuty returns the IsActiveDuty field if non-nil, zero value otherwise.

### GetIsActiveDutyOk

`func (o *IncomeVerificationPrecheckMilitaryInfo) GetIsActiveDutyOk() (*bool, bool)`

GetIsActiveDutyOk returns a tuple with the IsActiveDuty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActiveDuty

`func (o *IncomeVerificationPrecheckMilitaryInfo) SetIsActiveDuty(v bool)`

SetIsActiveDuty sets IsActiveDuty field to given value.

### HasIsActiveDuty

`func (o *IncomeVerificationPrecheckMilitaryInfo) HasIsActiveDuty() bool`

HasIsActiveDuty returns a boolean if a field has been set.

### SetIsActiveDutyNil

`func (o *IncomeVerificationPrecheckMilitaryInfo) SetIsActiveDutyNil(b bool)`

 SetIsActiveDutyNil sets the value for IsActiveDuty to be an explicit nil

### UnsetIsActiveDuty
`func (o *IncomeVerificationPrecheckMilitaryInfo) UnsetIsActiveDuty()`

UnsetIsActiveDuty ensures that no value is present for IsActiveDuty, not even an explicit nil
### GetBranch

`func (o *IncomeVerificationPrecheckMilitaryInfo) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *IncomeVerificationPrecheckMilitaryInfo) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *IncomeVerificationPrecheckMilitaryInfo) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *IncomeVerificationPrecheckMilitaryInfo) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### SetBranchNil

`func (o *IncomeVerificationPrecheckMilitaryInfo) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *IncomeVerificationPrecheckMilitaryInfo) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


