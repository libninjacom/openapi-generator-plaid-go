# TransferRepaymentReturn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferId** | **string** | The unique identifier of the guaranteed transfer that was returned. | 
**EventId** | **int32** | The unique identifier of the corresponding &#x60;reversed&#x60; transfer event. | 
**Amount** | **string** | The value of the returned transfer. | 
**IsoCurrencyCode** | **string** | The currency of the repayment, e.g. \&quot;USD\&quot;. | 

## Methods

### NewTransferRepaymentReturn

`func NewTransferRepaymentReturn(transferId string, eventId int32, amount string, isoCurrencyCode string, ) *TransferRepaymentReturn`

NewTransferRepaymentReturn instantiates a new TransferRepaymentReturn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRepaymentReturnWithDefaults

`func NewTransferRepaymentReturnWithDefaults() *TransferRepaymentReturn`

NewTransferRepaymentReturnWithDefaults instantiates a new TransferRepaymentReturn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferId

`func (o *TransferRepaymentReturn) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *TransferRepaymentReturn) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *TransferRepaymentReturn) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.


### GetEventId

`func (o *TransferRepaymentReturn) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *TransferRepaymentReturn) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *TransferRepaymentReturn) SetEventId(v int32)`

SetEventId sets EventId field to given value.


### GetAmount

`func (o *TransferRepaymentReturn) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferRepaymentReturn) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferRepaymentReturn) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIsoCurrencyCode

`func (o *TransferRepaymentReturn) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *TransferRepaymentReturn) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *TransferRepaymentReturn) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


