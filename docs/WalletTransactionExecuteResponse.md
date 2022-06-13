# WalletTransactionExecuteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | A unique ID identifying the transaction | 
**Status** | [**WalletTransactionStatus**](WalletTransactionStatus.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewWalletTransactionExecuteResponse

`func NewWalletTransactionExecuteResponse(transactionId string, status WalletTransactionStatus, requestId string, ) *WalletTransactionExecuteResponse`

NewWalletTransactionExecuteResponse instantiates a new WalletTransactionExecuteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTransactionExecuteResponseWithDefaults

`func NewWalletTransactionExecuteResponseWithDefaults() *WalletTransactionExecuteResponse`

NewWalletTransactionExecuteResponseWithDefaults instantiates a new WalletTransactionExecuteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *WalletTransactionExecuteResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *WalletTransactionExecuteResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *WalletTransactionExecuteResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetStatus

`func (o *WalletTransactionExecuteResponse) GetStatus() WalletTransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WalletTransactionExecuteResponse) GetStatusOk() (*WalletTransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WalletTransactionExecuteResponse) SetStatus(v WalletTransactionStatus)`

SetStatus sets Status field to given value.


### GetRequestId

`func (o *WalletTransactionExecuteResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WalletTransactionExecuteResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WalletTransactionExecuteResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


