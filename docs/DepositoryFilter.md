# DepositoryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSubtypes** | [**[]DepositoryAccountSubtype**](DepositoryAccountSubtype.md) | An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#account-type-schema).  | 

## Methods

### NewDepositoryFilter

`func NewDepositoryFilter(accountSubtypes []DepositoryAccountSubtype, ) *DepositoryFilter`

NewDepositoryFilter instantiates a new DepositoryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositoryFilterWithDefaults

`func NewDepositoryFilterWithDefaults() *DepositoryFilter`

NewDepositoryFilterWithDefaults instantiates a new DepositoryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSubtypes

`func (o *DepositoryFilter) GetAccountSubtypes() []DepositoryAccountSubtype`

GetAccountSubtypes returns the AccountSubtypes field if non-nil, zero value otherwise.

### GetAccountSubtypesOk

`func (o *DepositoryFilter) GetAccountSubtypesOk() (*[]DepositoryAccountSubtype, bool)`

GetAccountSubtypesOk returns a tuple with the AccountSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSubtypes

`func (o *DepositoryFilter) SetAccountSubtypes(v []DepositoryAccountSubtype)`

SetAccountSubtypes sets AccountSubtypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


