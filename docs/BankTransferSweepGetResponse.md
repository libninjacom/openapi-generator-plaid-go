# BankTransferSweepGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sweep** | [**BankTransferSweep**](BankTransferSweep.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewBankTransferSweepGetResponse

`func NewBankTransferSweepGetResponse(sweep BankTransferSweep, requestId string, ) *BankTransferSweepGetResponse`

NewBankTransferSweepGetResponse instantiates a new BankTransferSweepGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferSweepGetResponseWithDefaults

`func NewBankTransferSweepGetResponseWithDefaults() *BankTransferSweepGetResponse`

NewBankTransferSweepGetResponseWithDefaults instantiates a new BankTransferSweepGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSweep

`func (o *BankTransferSweepGetResponse) GetSweep() BankTransferSweep`

GetSweep returns the Sweep field if non-nil, zero value otherwise.

### GetSweepOk

`func (o *BankTransferSweepGetResponse) GetSweepOk() (*BankTransferSweep, bool)`

GetSweepOk returns a tuple with the Sweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweep

`func (o *BankTransferSweepGetResponse) SetSweep(v BankTransferSweep)`

SetSweep sets Sweep field to given value.


### GetRequestId

`func (o *BankTransferSweepGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BankTransferSweepGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BankTransferSweepGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


