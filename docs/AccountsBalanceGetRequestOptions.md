# AccountsBalanceGetRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIds** | Pointer to **[]string** | A list of &#x60;account_ids&#x60; to retrieve for the Item. The default value is &#x60;null&#x60;.  Note: An error will be returned if a provided &#x60;account_id&#x60; is not associated with the Item. | [optional] 
**MinLastUpdatedDatetime** | Pointer to **time.Time** | Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (&#x60;YYYY-MM-DDTHH:mm:ssZ&#x60;) indicating the oldest acceptable balance when making a request to &#x60;/accounts/balance/get&#x60;.  If the balance that is pulled for &#x60;ins_128026&#x60; (Capital One) is older than the given timestamp, an &#x60;INVALID_REQUEST&#x60; error with the code of &#x60;LAST_UPDATED_DATETIME_OUT_OF_RANGE&#x60; will be returned with the most recent timestamp for the requested account contained in the response.  This field is only used when the institution is &#x60;ins_128026&#x60; (Capital One), in which case a value must be provided or an &#x60;INVALID_REQUEST&#x60; error with the code of &#x60;INVALID_FIELD&#x60; will be returned. For all other institutions, this field is ignored. | [optional] 

## Methods

### NewAccountsBalanceGetRequestOptions

`func NewAccountsBalanceGetRequestOptions() *AccountsBalanceGetRequestOptions`

NewAccountsBalanceGetRequestOptions instantiates a new AccountsBalanceGetRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsBalanceGetRequestOptionsWithDefaults

`func NewAccountsBalanceGetRequestOptionsWithDefaults() *AccountsBalanceGetRequestOptions`

NewAccountsBalanceGetRequestOptionsWithDefaults instantiates a new AccountsBalanceGetRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIds

`func (o *AccountsBalanceGetRequestOptions) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *AccountsBalanceGetRequestOptions) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *AccountsBalanceGetRequestOptions) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.

### HasAccountIds

`func (o *AccountsBalanceGetRequestOptions) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.

### GetMinLastUpdatedDatetime

`func (o *AccountsBalanceGetRequestOptions) GetMinLastUpdatedDatetime() time.Time`

GetMinLastUpdatedDatetime returns the MinLastUpdatedDatetime field if non-nil, zero value otherwise.

### GetMinLastUpdatedDatetimeOk

`func (o *AccountsBalanceGetRequestOptions) GetMinLastUpdatedDatetimeOk() (*time.Time, bool)`

GetMinLastUpdatedDatetimeOk returns a tuple with the MinLastUpdatedDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLastUpdatedDatetime

`func (o *AccountsBalanceGetRequestOptions) SetMinLastUpdatedDatetime(v time.Time)`

SetMinLastUpdatedDatetime sets MinLastUpdatedDatetime field to given value.

### HasMinLastUpdatedDatetime

`func (o *AccountsBalanceGetRequestOptions) HasMinLastUpdatedDatetime() bool`

HasMinLastUpdatedDatetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


