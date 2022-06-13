# LinkTokenCreateLoanFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSubtypes** | Pointer to [**[]LoanAccountSubtype**](LoanAccountSubtype.md) | An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#account-type-schema).  | [optional] 

## Methods

### NewLinkTokenCreateLoanFilter

`func NewLinkTokenCreateLoanFilter() *LinkTokenCreateLoanFilter`

NewLinkTokenCreateLoanFilter instantiates a new LinkTokenCreateLoanFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTokenCreateLoanFilterWithDefaults

`func NewLinkTokenCreateLoanFilterWithDefaults() *LinkTokenCreateLoanFilter`

NewLinkTokenCreateLoanFilterWithDefaults instantiates a new LinkTokenCreateLoanFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSubtypes

`func (o *LinkTokenCreateLoanFilter) GetAccountSubtypes() []LoanAccountSubtype`

GetAccountSubtypes returns the AccountSubtypes field if non-nil, zero value otherwise.

### GetAccountSubtypesOk

`func (o *LinkTokenCreateLoanFilter) GetAccountSubtypesOk() (*[]LoanAccountSubtype, bool)`

GetAccountSubtypesOk returns a tuple with the AccountSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSubtypes

`func (o *LinkTokenCreateLoanFilter) SetAccountSubtypes(v []LoanAccountSubtype)`

SetAccountSubtypes sets AccountSubtypes field to given value.

### HasAccountSubtypes

`func (o *LinkTokenCreateLoanFilter) HasAccountSubtypes() bool`

HasAccountSubtypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


