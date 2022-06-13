# CustomerInitiatedReturnRisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | **int32** | A score from 0-99 that indicates the transaction return risk: a higher risk score suggests a higher return likelihood. | 
**RiskTier** | **int32** | A tier corresponding to the projected likelihood that the transaction, if initiated, will be subject to a return.  In the &#x60;customer_initiated_return_risk&#x60; object, there are five risk tiers corresponding to the scores:   1: Predicted customer-initiated return incidence rate between 0.00% - 0.02%   2: Predicted customer-initiated return incidence rate between 0.02% - 0.05%   3: Predicted customer-initiated return incidence rate between 0.05% - 0.1%   4: Predicted customer-initiated return incidence rate between 0.1% - 0.5%   5: Predicted customer-initiated return incidence rate greater than 0.5%  | 

## Methods

### NewCustomerInitiatedReturnRisk

`func NewCustomerInitiatedReturnRisk(score int32, riskTier int32, ) *CustomerInitiatedReturnRisk`

NewCustomerInitiatedReturnRisk instantiates a new CustomerInitiatedReturnRisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerInitiatedReturnRiskWithDefaults

`func NewCustomerInitiatedReturnRiskWithDefaults() *CustomerInitiatedReturnRisk`

NewCustomerInitiatedReturnRiskWithDefaults instantiates a new CustomerInitiatedReturnRisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *CustomerInitiatedReturnRisk) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CustomerInitiatedReturnRisk) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CustomerInitiatedReturnRisk) SetScore(v int32)`

SetScore sets Score field to given value.


### GetRiskTier

`func (o *CustomerInitiatedReturnRisk) GetRiskTier() int32`

GetRiskTier returns the RiskTier field if non-nil, zero value otherwise.

### GetRiskTierOk

`func (o *CustomerInitiatedReturnRisk) GetRiskTierOk() (*int32, bool)`

GetRiskTierOk returns a tuple with the RiskTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTier

`func (o *CustomerInitiatedReturnRisk) SetRiskTier(v int32)`

SetRiskTier sets RiskTier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


