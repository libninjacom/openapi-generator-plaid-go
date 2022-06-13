# Numbers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Will be used for the account number. | [optional] 
**AchRouting** | Pointer to **string** | Must be a valid ACH routing number. | [optional] 
**AchWireRouting** | Pointer to **string** | Must be a valid wire transfer routing number. | [optional] 
**EftInstitution** | Pointer to **string** | EFT institution number. Must be specified alongside &#x60;eft_branch&#x60;. | [optional] 
**EftBranch** | Pointer to **string** | EFT branch number. Must be specified alongside &#x60;eft_institution&#x60;. | [optional] 
**InternationalBic** | Pointer to **string** | Bank identifier code (BIC). Must be specified alongside &#x60;international_iban&#x60;. | [optional] 
**InternationalIban** | Pointer to **string** | International bank account number (IBAN). If no account number is specified via &#x60;account&#x60;, will also be used as the account number by default. Must be specified alongside &#x60;international_bic&#x60;. | [optional] 
**BacsSortCode** | Pointer to **string** | BACS sort code | [optional] 

## Methods

### NewNumbers

`func NewNumbers() *Numbers`

NewNumbers instantiates a new Numbers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumbersWithDefaults

`func NewNumbersWithDefaults() *Numbers`

NewNumbersWithDefaults instantiates a new Numbers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Numbers) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Numbers) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Numbers) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Numbers) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAchRouting

`func (o *Numbers) GetAchRouting() string`

GetAchRouting returns the AchRouting field if non-nil, zero value otherwise.

### GetAchRoutingOk

`func (o *Numbers) GetAchRoutingOk() (*string, bool)`

GetAchRoutingOk returns a tuple with the AchRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchRouting

`func (o *Numbers) SetAchRouting(v string)`

SetAchRouting sets AchRouting field to given value.

### HasAchRouting

`func (o *Numbers) HasAchRouting() bool`

HasAchRouting returns a boolean if a field has been set.

### GetAchWireRouting

`func (o *Numbers) GetAchWireRouting() string`

GetAchWireRouting returns the AchWireRouting field if non-nil, zero value otherwise.

### GetAchWireRoutingOk

`func (o *Numbers) GetAchWireRoutingOk() (*string, bool)`

GetAchWireRoutingOk returns a tuple with the AchWireRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchWireRouting

`func (o *Numbers) SetAchWireRouting(v string)`

SetAchWireRouting sets AchWireRouting field to given value.

### HasAchWireRouting

`func (o *Numbers) HasAchWireRouting() bool`

HasAchWireRouting returns a boolean if a field has been set.

### GetEftInstitution

`func (o *Numbers) GetEftInstitution() string`

GetEftInstitution returns the EftInstitution field if non-nil, zero value otherwise.

### GetEftInstitutionOk

`func (o *Numbers) GetEftInstitutionOk() (*string, bool)`

GetEftInstitutionOk returns a tuple with the EftInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEftInstitution

`func (o *Numbers) SetEftInstitution(v string)`

SetEftInstitution sets EftInstitution field to given value.

### HasEftInstitution

`func (o *Numbers) HasEftInstitution() bool`

HasEftInstitution returns a boolean if a field has been set.

### GetEftBranch

`func (o *Numbers) GetEftBranch() string`

GetEftBranch returns the EftBranch field if non-nil, zero value otherwise.

### GetEftBranchOk

`func (o *Numbers) GetEftBranchOk() (*string, bool)`

GetEftBranchOk returns a tuple with the EftBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEftBranch

`func (o *Numbers) SetEftBranch(v string)`

SetEftBranch sets EftBranch field to given value.

### HasEftBranch

`func (o *Numbers) HasEftBranch() bool`

HasEftBranch returns a boolean if a field has been set.

### GetInternationalBic

`func (o *Numbers) GetInternationalBic() string`

GetInternationalBic returns the InternationalBic field if non-nil, zero value otherwise.

### GetInternationalBicOk

`func (o *Numbers) GetInternationalBicOk() (*string, bool)`

GetInternationalBicOk returns a tuple with the InternationalBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalBic

`func (o *Numbers) SetInternationalBic(v string)`

SetInternationalBic sets InternationalBic field to given value.

### HasInternationalBic

`func (o *Numbers) HasInternationalBic() bool`

HasInternationalBic returns a boolean if a field has been set.

### GetInternationalIban

`func (o *Numbers) GetInternationalIban() string`

GetInternationalIban returns the InternationalIban field if non-nil, zero value otherwise.

### GetInternationalIbanOk

`func (o *Numbers) GetInternationalIbanOk() (*string, bool)`

GetInternationalIbanOk returns a tuple with the InternationalIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalIban

`func (o *Numbers) SetInternationalIban(v string)`

SetInternationalIban sets InternationalIban field to given value.

### HasInternationalIban

`func (o *Numbers) HasInternationalIban() bool`

HasInternationalIban returns a boolean if a field has been set.

### GetBacsSortCode

`func (o *Numbers) GetBacsSortCode() string`

GetBacsSortCode returns the BacsSortCode field if non-nil, zero value otherwise.

### GetBacsSortCodeOk

`func (o *Numbers) GetBacsSortCodeOk() (*string, bool)`

GetBacsSortCodeOk returns a tuple with the BacsSortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacsSortCode

`func (o *Numbers) SetBacsSortCode(v string)`

SetBacsSortCode sets BacsSortCode field to given value.

### HasBacsSortCode

`func (o *Numbers) HasBacsSortCode() bool`

HasBacsSortCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


