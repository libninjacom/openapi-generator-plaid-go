# BankTransferSweep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the sweep. | 
**CreatedAt** | **time.Time** | The datetime when the sweep occurred, in RFC 3339 format. | 
**Amount** | **string** | The amount of the sweep. | 
**IsoCurrencyCode** | **string** | The currency of the sweep, e.g. \&quot;USD\&quot;. | 

## Methods

### NewBankTransferSweep

`func NewBankTransferSweep(id string, createdAt time.Time, amount string, isoCurrencyCode string, ) *BankTransferSweep`

NewBankTransferSweep instantiates a new BankTransferSweep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferSweepWithDefaults

`func NewBankTransferSweepWithDefaults() *BankTransferSweep`

NewBankTransferSweepWithDefaults instantiates a new BankTransferSweep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankTransferSweep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankTransferSweep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankTransferSweep) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *BankTransferSweep) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BankTransferSweep) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BankTransferSweep) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetAmount

`func (o *BankTransferSweep) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BankTransferSweep) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BankTransferSweep) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIsoCurrencyCode

`func (o *BankTransferSweep) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *BankTransferSweep) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *BankTransferSweep) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


