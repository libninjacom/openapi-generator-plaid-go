# BankInitiatedReturnRisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | **int32** | A score from 0-99 that indicates the transaction return risk: a higher risk score suggests a higher return likelihood. | 
**RiskTier** | **int32** | In the &#x60;bank_initiated_return_risk&#x60; object, there are eight risk tiers corresponding to the scores:   1: Predicted bank-initiated return incidence rate between 0.0% - 0.5%   2: Predicted bank-initiated return incidence rate between 0.5% - 1.5%   3: Predicted bank-initiated return incidence rate between 1.5% - 3%   4: Predicted bank-initiated return incidence rate between 3% - 5%   5: Predicted bank-initiated return incidence rate between 5% - 10%   6: Predicted bank-initiated return incidence rate between 10% - 15%   7: Predicted bank-initiated return incidence rate between 15% and 50%   8: Predicted bank-initiated return incidence rate greater than 50%  | 

## Methods

### NewBankInitiatedReturnRisk

`func NewBankInitiatedReturnRisk(score int32, riskTier int32, ) *BankInitiatedReturnRisk`

NewBankInitiatedReturnRisk instantiates a new BankInitiatedReturnRisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankInitiatedReturnRiskWithDefaults

`func NewBankInitiatedReturnRiskWithDefaults() *BankInitiatedReturnRisk`

NewBankInitiatedReturnRiskWithDefaults instantiates a new BankInitiatedReturnRisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *BankInitiatedReturnRisk) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *BankInitiatedReturnRisk) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *BankInitiatedReturnRisk) SetScore(v int32)`

SetScore sets Score field to given value.


### GetRiskTier

`func (o *BankInitiatedReturnRisk) GetRiskTier() int32`

GetRiskTier returns the RiskTier field if non-nil, zero value otherwise.

### GetRiskTierOk

`func (o *BankInitiatedReturnRisk) GetRiskTierOk() (*int32, bool)`

GetRiskTierOk returns a tuple with the RiskTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTier

`func (o *BankInitiatedReturnRisk) SetRiskTier(v int32)`

SetRiskTier sets RiskTier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


