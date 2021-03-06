# TransferRepaymentReturnListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepaymentReturns** | [**[]TransferRepaymentReturn**](TransferRepaymentReturn.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransferRepaymentReturnListResponse

`func NewTransferRepaymentReturnListResponse(repaymentReturns []TransferRepaymentReturn, requestId string, ) *TransferRepaymentReturnListResponse`

NewTransferRepaymentReturnListResponse instantiates a new TransferRepaymentReturnListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRepaymentReturnListResponseWithDefaults

`func NewTransferRepaymentReturnListResponseWithDefaults() *TransferRepaymentReturnListResponse`

NewTransferRepaymentReturnListResponseWithDefaults instantiates a new TransferRepaymentReturnListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepaymentReturns

`func (o *TransferRepaymentReturnListResponse) GetRepaymentReturns() []TransferRepaymentReturn`

GetRepaymentReturns returns the RepaymentReturns field if non-nil, zero value otherwise.

### GetRepaymentReturnsOk

`func (o *TransferRepaymentReturnListResponse) GetRepaymentReturnsOk() (*[]TransferRepaymentReturn, bool)`

GetRepaymentReturnsOk returns a tuple with the RepaymentReturns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepaymentReturns

`func (o *TransferRepaymentReturnListResponse) SetRepaymentReturns(v []TransferRepaymentReturn)`

SetRepaymentReturns sets RepaymentReturns field to given value.


### GetRequestId

`func (o *TransferRepaymentReturnListResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransferRepaymentReturnListResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransferRepaymentReturnListResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


