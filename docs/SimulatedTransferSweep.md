# SimulatedTransferSweep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the sweep. | 
**Created** | **time.Time** | The datetime when the sweep occurred, in RFC 3339 format. | 
**Amount** | **string** | Signed decimal amount of the sweep as it appears on your sweep account ledger (e.g. \&quot;-10.00\&quot;)  If amount is not present, the sweep was net-settled to zero and outstanding debits and credits between the sweep account and Plaid are balanced. | 
**IsoCurrencyCode** | **string** | The currency of the sweep, e.g. \&quot;USD\&quot;. | 

## Methods

### NewSimulatedTransferSweep

`func NewSimulatedTransferSweep(id string, created time.Time, amount string, isoCurrencyCode string, ) *SimulatedTransferSweep`

NewSimulatedTransferSweep instantiates a new SimulatedTransferSweep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulatedTransferSweepWithDefaults

`func NewSimulatedTransferSweepWithDefaults() *SimulatedTransferSweep`

NewSimulatedTransferSweepWithDefaults instantiates a new SimulatedTransferSweep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimulatedTransferSweep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimulatedTransferSweep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimulatedTransferSweep) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *SimulatedTransferSweep) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SimulatedTransferSweep) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SimulatedTransferSweep) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAmount

`func (o *SimulatedTransferSweep) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SimulatedTransferSweep) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SimulatedTransferSweep) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIsoCurrencyCode

`func (o *SimulatedTransferSweep) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *SimulatedTransferSweep) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *SimulatedTransferSweep) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


