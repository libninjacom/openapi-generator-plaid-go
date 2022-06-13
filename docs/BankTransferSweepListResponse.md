# BankTransferSweepListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sweeps** | [**[]BankTransferSweep**](BankTransferSweep.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewBankTransferSweepListResponse

`func NewBankTransferSweepListResponse(sweeps []BankTransferSweep, requestId string, ) *BankTransferSweepListResponse`

NewBankTransferSweepListResponse instantiates a new BankTransferSweepListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferSweepListResponseWithDefaults

`func NewBankTransferSweepListResponseWithDefaults() *BankTransferSweepListResponse`

NewBankTransferSweepListResponseWithDefaults instantiates a new BankTransferSweepListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSweeps

`func (o *BankTransferSweepListResponse) GetSweeps() []BankTransferSweep`

GetSweeps returns the Sweeps field if non-nil, zero value otherwise.

### GetSweepsOk

`func (o *BankTransferSweepListResponse) GetSweepsOk() (*[]BankTransferSweep, bool)`

GetSweepsOk returns a tuple with the Sweeps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweeps

`func (o *BankTransferSweepListResponse) SetSweeps(v []BankTransferSweep)`

SetSweeps sets Sweeps field to given value.


### GetRequestId

`func (o *BankTransferSweepListResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BankTransferSweepListResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BankTransferSweepListResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


