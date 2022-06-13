# LinkTokenCreateRequestUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSelectionEnabled** | Pointer to **bool** | If &#x60;true&#x60;, enables [update mode with Account Select](https://plaid.com/docs/link/update-mode/#using-update-mode-to-request-new-accounts). | [optional] [default to false]

## Methods

### NewLinkTokenCreateRequestUpdate

`func NewLinkTokenCreateRequestUpdate() *LinkTokenCreateRequestUpdate`

NewLinkTokenCreateRequestUpdate instantiates a new LinkTokenCreateRequestUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTokenCreateRequestUpdateWithDefaults

`func NewLinkTokenCreateRequestUpdateWithDefaults() *LinkTokenCreateRequestUpdate`

NewLinkTokenCreateRequestUpdateWithDefaults instantiates a new LinkTokenCreateRequestUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSelectionEnabled

`func (o *LinkTokenCreateRequestUpdate) GetAccountSelectionEnabled() bool`

GetAccountSelectionEnabled returns the AccountSelectionEnabled field if non-nil, zero value otherwise.

### GetAccountSelectionEnabledOk

`func (o *LinkTokenCreateRequestUpdate) GetAccountSelectionEnabledOk() (*bool, bool)`

GetAccountSelectionEnabledOk returns a tuple with the AccountSelectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSelectionEnabled

`func (o *LinkTokenCreateRequestUpdate) SetAccountSelectionEnabled(v bool)`

SetAccountSelectionEnabled sets AccountSelectionEnabled field to given value.

### HasAccountSelectionEnabled

`func (o *LinkTokenCreateRequestUpdate) HasAccountSelectionEnabled() bool`

HasAccountSelectionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


