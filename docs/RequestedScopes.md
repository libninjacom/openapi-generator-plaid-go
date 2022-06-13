# RequestedScopes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountFilters** | Pointer to [**AccountFilter**](AccountFilter.md) |  | [optional] 
**AccountSelectionCardinality** | [**AccountSelectionCardinality**](AccountSelectionCardinality.md) |  | 

## Methods

### NewRequestedScopes

`func NewRequestedScopes(accountSelectionCardinality AccountSelectionCardinality, ) *RequestedScopes`

NewRequestedScopes instantiates a new RequestedScopes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestedScopesWithDefaults

`func NewRequestedScopesWithDefaults() *RequestedScopes`

NewRequestedScopesWithDefaults instantiates a new RequestedScopes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountFilters

`func (o *RequestedScopes) GetAccountFilters() AccountFilter`

GetAccountFilters returns the AccountFilters field if non-nil, zero value otherwise.

### GetAccountFiltersOk

`func (o *RequestedScopes) GetAccountFiltersOk() (*AccountFilter, bool)`

GetAccountFiltersOk returns a tuple with the AccountFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFilters

`func (o *RequestedScopes) SetAccountFilters(v AccountFilter)`

SetAccountFilters sets AccountFilters field to given value.

### HasAccountFilters

`func (o *RequestedScopes) HasAccountFilters() bool`

HasAccountFilters returns a boolean if a field has been set.

### GetAccountSelectionCardinality

`func (o *RequestedScopes) GetAccountSelectionCardinality() AccountSelectionCardinality`

GetAccountSelectionCardinality returns the AccountSelectionCardinality field if non-nil, zero value otherwise.

### GetAccountSelectionCardinalityOk

`func (o *RequestedScopes) GetAccountSelectionCardinalityOk() (*AccountSelectionCardinality, bool)`

GetAccountSelectionCardinalityOk returns a tuple with the AccountSelectionCardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSelectionCardinality

`func (o *RequestedScopes) SetAccountSelectionCardinality(v AccountSelectionCardinality)`

SetAccountSelectionCardinality sets AccountSelectionCardinality field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


