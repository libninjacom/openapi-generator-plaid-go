# TransferRepayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepaymentId** | **string** | Identifier of the repayment. | 
**Created** | **time.Time** | The datetime when the repayment occurred, in RFC 3339 format. | 
**Amount** | **string** | Decimal amount of the repayment as it appears on your account ledger. | 
**IsoCurrencyCode** | **string** | The currency of the repayment, e.g. \&quot;USD\&quot;. | 

## Methods

### NewTransferRepayment

`func NewTransferRepayment(repaymentId string, created time.Time, amount string, isoCurrencyCode string, ) *TransferRepayment`

NewTransferRepayment instantiates a new TransferRepayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRepaymentWithDefaults

`func NewTransferRepaymentWithDefaults() *TransferRepayment`

NewTransferRepaymentWithDefaults instantiates a new TransferRepayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepaymentId

`func (o *TransferRepayment) GetRepaymentId() string`

GetRepaymentId returns the RepaymentId field if non-nil, zero value otherwise.

### GetRepaymentIdOk

`func (o *TransferRepayment) GetRepaymentIdOk() (*string, bool)`

GetRepaymentIdOk returns a tuple with the RepaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepaymentId

`func (o *TransferRepayment) SetRepaymentId(v string)`

SetRepaymentId sets RepaymentId field to given value.


### GetCreated

`func (o *TransferRepayment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TransferRepayment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TransferRepayment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAmount

`func (o *TransferRepayment) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferRepayment) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferRepayment) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIsoCurrencyCode

`func (o *TransferRepayment) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *TransferRepayment) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *TransferRepayment) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


