# TransactionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentChannel** | **string** | The channel used to make a payment. &#x60;online:&#x60; transactions that took place online.  &#x60;in store:&#x60; transactions that were made at a physical location.  &#x60;other:&#x60; transactions that relate to banks, e.g. fees or deposits.  This field replaces the &#x60;transaction_type&#x60; field.  | 
**AuthorizedDate** | **NullableString** | The date that the transaction was authorized. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( &#x60;YYYY-MM-DD&#x60; ). The &#x60;authorized_date&#x60; field uses machine learning to determine a transaction date for transactions where the &#x60;date_transacted&#x60; is not available. If the &#x60;date_transacted&#x60; field is present and not &#x60;null&#x60;, the &#x60;authorized_date&#x60; field will have the same value as the &#x60;date_transacted&#x60; field. | 
**AuthorizedDatetime** | **NullableTime** | Date and time when a transaction was authorized in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( &#x60;YYYY-MM-DDTHH:mm:ssZ&#x60; ).  This field is returned for select financial institutions and comes as provided by the institution. It may contain default time values (such as 00:00:00). | 
**Datetime** | **NullableTime** | Date and time when a transaction was posted in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( &#x60;YYYY-MM-DDTHH:mm:ssZ&#x60; ).  This field is returned for select financial institutions and comes as provided by the institution. It may contain default time values (such as 00:00:00). | 
**TransactionCode** | [**NullableTransactionCode**](TransactionCode.md) |  | 
**PersonalFinanceCategory** | Pointer to [**NullablePersonalFinanceCategory**](PersonalFinanceCategory.md) |  | [optional] 

## Methods

### NewTransactionAllOf

`func NewTransactionAllOf(paymentChannel string, authorizedDate NullableString, authorizedDatetime NullableTime, datetime NullableTime, transactionCode NullableTransactionCode, ) *TransactionAllOf`

NewTransactionAllOf instantiates a new TransactionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionAllOfWithDefaults

`func NewTransactionAllOfWithDefaults() *TransactionAllOf`

NewTransactionAllOfWithDefaults instantiates a new TransactionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentChannel

`func (o *TransactionAllOf) GetPaymentChannel() string`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *TransactionAllOf) GetPaymentChannelOk() (*string, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *TransactionAllOf) SetPaymentChannel(v string)`

SetPaymentChannel sets PaymentChannel field to given value.


### GetAuthorizedDate

`func (o *TransactionAllOf) GetAuthorizedDate() string`

GetAuthorizedDate returns the AuthorizedDate field if non-nil, zero value otherwise.

### GetAuthorizedDateOk

`func (o *TransactionAllOf) GetAuthorizedDateOk() (*string, bool)`

GetAuthorizedDateOk returns a tuple with the AuthorizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedDate

`func (o *TransactionAllOf) SetAuthorizedDate(v string)`

SetAuthorizedDate sets AuthorizedDate field to given value.


### SetAuthorizedDateNil

`func (o *TransactionAllOf) SetAuthorizedDateNil(b bool)`

 SetAuthorizedDateNil sets the value for AuthorizedDate to be an explicit nil

### UnsetAuthorizedDate
`func (o *TransactionAllOf) UnsetAuthorizedDate()`

UnsetAuthorizedDate ensures that no value is present for AuthorizedDate, not even an explicit nil
### GetAuthorizedDatetime

`func (o *TransactionAllOf) GetAuthorizedDatetime() time.Time`

GetAuthorizedDatetime returns the AuthorizedDatetime field if non-nil, zero value otherwise.

### GetAuthorizedDatetimeOk

`func (o *TransactionAllOf) GetAuthorizedDatetimeOk() (*time.Time, bool)`

GetAuthorizedDatetimeOk returns a tuple with the AuthorizedDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedDatetime

`func (o *TransactionAllOf) SetAuthorizedDatetime(v time.Time)`

SetAuthorizedDatetime sets AuthorizedDatetime field to given value.


### SetAuthorizedDatetimeNil

`func (o *TransactionAllOf) SetAuthorizedDatetimeNil(b bool)`

 SetAuthorizedDatetimeNil sets the value for AuthorizedDatetime to be an explicit nil

### UnsetAuthorizedDatetime
`func (o *TransactionAllOf) UnsetAuthorizedDatetime()`

UnsetAuthorizedDatetime ensures that no value is present for AuthorizedDatetime, not even an explicit nil
### GetDatetime

`func (o *TransactionAllOf) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *TransactionAllOf) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *TransactionAllOf) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### SetDatetimeNil

`func (o *TransactionAllOf) SetDatetimeNil(b bool)`

 SetDatetimeNil sets the value for Datetime to be an explicit nil

### UnsetDatetime
`func (o *TransactionAllOf) UnsetDatetime()`

UnsetDatetime ensures that no value is present for Datetime, not even an explicit nil
### GetTransactionCode

`func (o *TransactionAllOf) GetTransactionCode() TransactionCode`

GetTransactionCode returns the TransactionCode field if non-nil, zero value otherwise.

### GetTransactionCodeOk

`func (o *TransactionAllOf) GetTransactionCodeOk() (*TransactionCode, bool)`

GetTransactionCodeOk returns a tuple with the TransactionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCode

`func (o *TransactionAllOf) SetTransactionCode(v TransactionCode)`

SetTransactionCode sets TransactionCode field to given value.


### SetTransactionCodeNil

`func (o *TransactionAllOf) SetTransactionCodeNil(b bool)`

 SetTransactionCodeNil sets the value for TransactionCode to be an explicit nil

### UnsetTransactionCode
`func (o *TransactionAllOf) UnsetTransactionCode()`

UnsetTransactionCode ensures that no value is present for TransactionCode, not even an explicit nil
### GetPersonalFinanceCategory

`func (o *TransactionAllOf) GetPersonalFinanceCategory() PersonalFinanceCategory`

GetPersonalFinanceCategory returns the PersonalFinanceCategory field if non-nil, zero value otherwise.

### GetPersonalFinanceCategoryOk

`func (o *TransactionAllOf) GetPersonalFinanceCategoryOk() (*PersonalFinanceCategory, bool)`

GetPersonalFinanceCategoryOk returns a tuple with the PersonalFinanceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalFinanceCategory

`func (o *TransactionAllOf) SetPersonalFinanceCategory(v PersonalFinanceCategory)`

SetPersonalFinanceCategory sets PersonalFinanceCategory field to given value.

### HasPersonalFinanceCategory

`func (o *TransactionAllOf) HasPersonalFinanceCategory() bool`

HasPersonalFinanceCategory returns a boolean if a field has been set.

### SetPersonalFinanceCategoryNil

`func (o *TransactionAllOf) SetPersonalFinanceCategoryNil(b bool)`

 SetPersonalFinanceCategoryNil sets the value for PersonalFinanceCategory to be an explicit nil

### UnsetPersonalFinanceCategory
`func (o *TransactionAllOf) UnsetPersonalFinanceCategory()`

UnsetPersonalFinanceCategory ensures that no value is present for PersonalFinanceCategory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


