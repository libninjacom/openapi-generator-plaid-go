# NumbersACHNullable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The Plaid account ID associated with the account numbers | 
**Account** | **string** | The ACH account number for the account.  Note that when using OAuth with Chase Bank (&#x60;ins_56&#x60;), Chase will issue \&quot;tokenized\&quot; routing and account numbers, which are not the user&#39;s actual account and routing numbers. These tokenized numbers should work identically to normal account and routing numbers. The digits returned in the &#x60;mask&#x60; field will continue to reflect the actual account number, rather than the tokenized account number; for this reason, when displaying account numbers to the user to help them identify their account in your UI, always use the &#x60;mask&#x60; rather than truncating the &#x60;account&#x60; number. If a user revokes their permissions to your app, the tokenized numbers will continue to work for ACH deposits, but not withdrawals. | 
**Routing** | **string** | The ACH routing number for the account. If the institution is &#x60;ins_56&#x60;, this may be a tokenized routing number. For more information, see the description of the &#x60;account&#x60; field. | 
**WireRouting** | **NullableString** | The wire transfer routing number for the account, if available | 

## Methods

### NewNumbersACHNullable

`func NewNumbersACHNullable(accountId string, account string, routing string, wireRouting NullableString, ) *NumbersACHNullable`

NewNumbersACHNullable instantiates a new NumbersACHNullable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumbersACHNullableWithDefaults

`func NewNumbersACHNullableWithDefaults() *NumbersACHNullable`

NewNumbersACHNullableWithDefaults instantiates a new NumbersACHNullable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NumbersACHNullable) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NumbersACHNullable) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NumbersACHNullable) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccount

`func (o *NumbersACHNullable) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NumbersACHNullable) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NumbersACHNullable) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetRouting

`func (o *NumbersACHNullable) GetRouting() string`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *NumbersACHNullable) GetRoutingOk() (*string, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *NumbersACHNullable) SetRouting(v string)`

SetRouting sets Routing field to given value.


### GetWireRouting

`func (o *NumbersACHNullable) GetWireRouting() string`

GetWireRouting returns the WireRouting field if non-nil, zero value otherwise.

### GetWireRoutingOk

`func (o *NumbersACHNullable) GetWireRoutingOk() (*string, bool)`

GetWireRoutingOk returns a tuple with the WireRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireRouting

`func (o *NumbersACHNullable) SetWireRouting(v string)`

SetWireRouting sets WireRouting field to given value.


### SetWireRoutingNil

`func (o *NumbersACHNullable) SetWireRoutingNil(b bool)`

 SetWireRoutingNil sets the value for WireRouting to be an explicit nil

### UnsetWireRouting
`func (o *NumbersACHNullable) UnsetWireRouting()`

UnsetWireRouting ensures that no value is present for WireRouting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


