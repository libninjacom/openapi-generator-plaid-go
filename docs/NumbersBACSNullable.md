# NumbersBACSNullable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The Plaid account ID associated with the account numbers | 
**Account** | **string** | The BACS account number for the account | 
**SortCode** | **string** | The BACS sort code for the account | 

## Methods

### NewNumbersBACSNullable

`func NewNumbersBACSNullable(accountId string, account string, sortCode string, ) *NumbersBACSNullable`

NewNumbersBACSNullable instantiates a new NumbersBACSNullable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumbersBACSNullableWithDefaults

`func NewNumbersBACSNullableWithDefaults() *NumbersBACSNullable`

NewNumbersBACSNullableWithDefaults instantiates a new NumbersBACSNullable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NumbersBACSNullable) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NumbersBACSNullable) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NumbersBACSNullable) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccount

`func (o *NumbersBACSNullable) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NumbersBACSNullable) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NumbersBACSNullable) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetSortCode

`func (o *NumbersBACSNullable) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *NumbersBACSNullable) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *NumbersBACSNullable) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


